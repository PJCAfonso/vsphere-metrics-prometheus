/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vsphere

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

var (
	metricsMapEsx = make(map[int]*prometheus.GaugeVec)
)

func (c *Client) registerEsxMetrics() error {
	log.Debugln("registerEsxMetrics ENTER")

	// Create client
	err := c.getClient()
	if err != nil {
		log.Errorln("getClient failed:", err)
		log.Debugln("registerEsxMetrics LEAVE")

		return err
	}

	var performanceManager mo.PerformanceManager
	err = c.vClient.RetrieveOne(*c.ctx, *c.vClient.ServiceContent.PerfManager, nil, &performanceManager)
	if err != nil {
		log.Errorln("RetrieveOne failed:", err)
		log.Debugln("registerEsxMetrics LEAVE")

		return err
	}

	// As outline in https://code.vmware.com/doc/preview?id=6784#/doc/vim.PerformanceManager.CounterInfo.html
	for _, perfCounterInfo := range performanceManager.PerfCounter {
		nameInfo := perfCounterInfo.NameInfo.GetElementDescription()
		keyTmp := strings.Join(strings.Split(nameInfo.Key, "."), "_")
		metricName := fmt.Sprintf("%d_%s", perfCounterInfo.Key, strcase.ToSnake(keyTmp))
		log.Debugln("Key:", metricName)

		myMetric := prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "vsphere",
				Subsystem: "esx",
				Name:      metricName,
				Help:      metricName,
			},
			[]string{"datacenter", "host"},
		)
		metricsMapEsx[int(perfCounterInfo.Key)] = myMetric
		prometheus.MustRegister(myMetric)
	}

	log.Debugln("registerEsxMetrics Succeeded")
	log.Debugln("registerEsxMetrics LEAVE")

	return nil
}

//GetVSphereEsxStats gets stats for an individual ESX host
func (c *Client) GetVSphereEsxStats(w http.ResponseWriter, r *http.Request) error {
	log.Debugln("GetVSphereEsxStats ENTER")

	vars := mux.Vars(r)
	datacenterStr := vars["datacenter"]
	log.Infoln("Datacenter:", datacenterStr)
	hostStr := vars["host"]
	log.Infoln("Host:", hostStr)

	// Create client
	err := c.getClient()
	if err != nil {
		http.Error(w, "Unable connect to the vCenter Server", http.StatusGone)
		log.Errorln("getClient failed:", err)
		log.Debugln("GetVSphereEsxStats LEAVE")

		return err
	}

	//find our objects
	finder := find.NewFinder(c.vClient.Client, false)

	dc, err := finder.Datacenter(*c.ctx, datacenterStr)
	if err != nil {
		http.Error(w, "Unable find the Datacener", http.StatusGone)
		log.Errorln("finder.Datacenter(", datacenterStr, "):", err)
		log.Debugln("GetVSphereEsxStats LEAVE")
		return err
	}
	finder.SetDatacenter(dc)

	host, err := finder.HostSystem(*c.ctx, hostStr)
	if err != nil {
		http.Error(w, "Unable find the HostSystem", http.StatusGone)
		log.Errorln("finder.HostSystem(", hostStr, "):", err)
		log.Debugln("GetVSphereEsxStats LEAVE")
		return err
	}

	log.Infoln("Host:", host.Name())
	log.Infoln("Host:", host.InventoryPath)

	var oHost mo.HostSystem
	err = host.Properties(*c.ctx, host.Reference(), []string{"summary"}, &oHost)
	if err != nil {
		http.Error(w, "Unable get the HostSystem properties", http.StatusBadRequest)
		log.Errorln("host.Properties(", hostStr, "):", err)
		log.Debugln("GetVSphereEsxStats LEAVE")
		return err
	}

	log.Infoln(oHost.Self.Value)
	log.Infoln(oHost.Summary.Config.Name)
	log.Infoln(oHost.Summary.Config.Product.Version)
	log.Infoln(oHost.Summary.Config.Product.Build)
	log.Infoln(string(oHost.Summary.OverallStatus))
	log.Infoln(string(oHost.OverallStatus))

	var performanceManager mo.PerformanceManager
	err = c.vClient.RetrieveOne(*c.ctx, *c.vClient.ServiceContent.PerfManager, nil, &performanceManager)
	if err != nil {
		log.Errorln("RetrieveOne failed:", err)
		log.Debugln("GetVSphereEsxStats LEAVE")

		return err
	}

	querySpec := types.PerfQuerySpec{
		Entity:     host.Reference(),
		MaxSample:  1,
		IntervalId: 20,
	}
	query := types.QueryPerf{
		This:      *c.vClient.ServiceContent.PerfManager,
		QuerySpec: []types.PerfQuerySpec{querySpec},
	}

	response, err := methods.QueryPerf(*c.ctx, c.vClient, &query)
	if err != nil {
		log.Errorln("QueryPerf failed:", err)
		log.Debugln("GetVSphereEsxStats LEAVE")

		return err
	}

	for _, base := range response.Returnval {
		metric := base.(*types.PerfEntityMetric)
		for _, baseSeries := range metric.Value {
			series := baseSeries.(*types.PerfMetricIntSeries)
			myMetric := metricsMapEsx[int(series.Id.CounterId)]

			if myMetric == nil {
				log.Errorln("Unable to find metric for", series.Id.CounterId)
				continue
			}

			myMetric.WithLabelValues(datacenterStr, hostStr).Set(float64(series.Value[0]))
		}
	}

	log.Debugln("GetVSphereEsxStats Succeeded")
	log.Debugln("GetVSphereEsxStats LEAVE")

	return nil
}
