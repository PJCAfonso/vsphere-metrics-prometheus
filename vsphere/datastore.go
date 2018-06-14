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

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/mo"
)

const (
	datastoreFreespace = (iota + 1024)
	datastoreUncommitted
	datastoreUsedSpace
	datastoreCapacity
	datastoreProvisioned
)

var (
	metricsMapDatastore = make(map[int]*prometheus.GaugeVec)
)

func (c *Client) registerDatastoreMetrics() error {
	log.Debugln("registerDatastoreMetrics ENTER")

	//freespace
	metricName := fmt.Sprintf("%d_freespace_size", datastoreFreespace)
	log.Debugln("Key:", metricName)

	myMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "datastore",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter", "datastore"},
	)
	metricsMapDatastore[datastoreFreespace] = myMetric
	prometheus.MustRegister(myMetric)

	//uncommitted
	metricName = fmt.Sprintf("%d_uncommitted_size", datastoreUncommitted)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "datastore",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter", "datastore"},
	)
	metricsMapDatastore[datastoreUncommitted] = myMetric
	prometheus.MustRegister(myMetric)

	//usedspace
	metricName = fmt.Sprintf("%d_usedspace_size", datastoreUsedSpace)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "datastore",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter", "datastore"},
	)
	metricsMapDatastore[datastoreUsedSpace] = myMetric
	prometheus.MustRegister(myMetric)

	//capacity
	metricName = fmt.Sprintf("%d_capacity_size", datastoreCapacity)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "datastore",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter", "datastore"},
	)
	metricsMapDatastore[datastoreCapacity] = myMetric
	prometheus.MustRegister(myMetric)

	//provisioned
	metricName = fmt.Sprintf("%d_provisioned_size", datastoreProvisioned)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "datastore",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter", "datastore"},
	)
	metricsMapDatastore[datastoreProvisioned] = myMetric
	prometheus.MustRegister(myMetric)

	log.Debugln("registerDatastoreMetrics Succeeded")
	log.Debugln("registerDatastoreMetrics LEAVE")

	return nil
}

//GetVSphereDatastoreStats gets stats for an individual VM
func (c *Client) GetVSphereDatastoreStats(w http.ResponseWriter, r *http.Request) error {
	log.Debugln("GetVSphereDatastoreStats ENTER")

	vars := mux.Vars(r)
	datacenterStr := vars["datacenter"]
	log.Infoln("Datacenter:", datacenterStr)
	datastoreStr := vars["datastore"]
	log.Infoln("VM:", datastoreStr)

	// Create client
	err := c.getClient()
	if err != nil {
		http.Error(w, "Unable connect to the vCenter Server", http.StatusGone)
		log.Errorln("getClient failed:", err)
		log.Debugln("GetVSphereDatastoreStats LEAVE")

		return err
	}

	//find our objects
	finder := find.NewFinder(c.vClient.Client, false)

	dc, err := finder.Datacenter(*c.ctx, datacenterStr)
	if err != nil {
		http.Error(w, "Unable find the Datacener", http.StatusGone)
		log.Errorln("finder.Datacenter(", datacenterStr, "):", err)
		log.Debugln("GetVSphereDatastoreStats LEAVE")
		return err
	}
	finder.SetDatacenter(dc)

	datastore, err := finder.Datastore(*c.ctx, datastoreStr)
	if err != nil {
		http.Error(w, "Unable find the Datastore", http.StatusGone)
		log.Errorln("finder.Datastore(", datastoreStr, "):", err)
		log.Debugln("GetVSphereDatastoreStats LEAVE")
		return err
	}

	log.Infoln("Datastore:", datastore.Name())
	log.Infoln("Datastore:", datastore.InventoryPath)

	var oDatastore mo.Datastore
	err = datastore.Properties(*c.ctx, datastore.Reference(), []string{"summary"}, &oDatastore)
	if err != nil {
		http.Error(w, "Unable get the Datastore properties", http.StatusBadRequest)
		log.Errorln("datastore.Properties(", datastoreStr, "):", err)
		log.Debugln("GetVSphereDatastoreStats LEAVE")
		return err
	}

	log.Infoln(oDatastore.Self.Value)
	log.Infoln(oDatastore.Summary.Name)
	log.Infoln(oDatastore.Summary.Type)

	myMetric := metricsMapDatastore[datastoreFreespace]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr, datastoreStr).Set(float64(oDatastore.Summary.FreeSpace))
	}
	myMetric = metricsMapDatastore[datastoreUncommitted]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr, datastoreStr).Set(float64(oDatastore.Summary.Uncommitted))
	}
	myMetric = metricsMapDatastore[datastoreUsedSpace]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr, datastoreStr).Set(float64((oDatastore.Summary.Capacity - oDatastore.Summary.FreeSpace)))
	}
	myMetric = metricsMapDatastore[datastoreCapacity]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr, datastoreStr).Set(float64(oDatastore.Summary.Capacity))
	}
	myMetric = metricsMapDatastore[datastoreProvisioned]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr, datastoreStr).Set(float64((oDatastore.Summary.Capacity - oDatastore.Summary.FreeSpace + oDatastore.Summary.Uncommitted)))
	}

	log.Debugln("GetVSphereDatastoreStats Succeeded")
	log.Debugln("GetVSphereDatastoreStats LEAVE")

	return nil
}
