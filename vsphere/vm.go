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
	vmBalloonedMemory = (iota + 2048)
	vmCompressedMemory
	vmConsumedOverheadMemory
	vmDistributedCpuEntitlement
	vmDistributedMemoryEntitlement
	//vmFtLatencyStatus
	vmFtLogBandwidth
	vmFtSecondaryLatency
	//vmGuestHeartbeatStatus
	vmGuestMemoryUsage
	vmHostMemoryUsage
	vmOverallCpuDemand
	vmOverallCpuUsage
	vmPrivateMemory
	vmSharedMemory
	vmSsdSwappedMemory
	vmStaticCpuEntitlement
	vmStaticMemoryEntitlement
	vmSwappedMemory
	vmUptimeSeconds
)

var (
	metricsMapVM = make(map[int]*prometheus.GaugeVec)
	//metricsMapVM = make(map[int]*prometheus.Desc)
)

func (c *Client) registerVMMetrics() error {
	log.Debugln("registerVMMetrics ENTER")

	metricName := fmt.Sprintf("%d_ballooned_memory", vmBalloonedMemory)
	log.Debugln("Key:", metricName)

	myMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmBalloonedMemory)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_compressed_memory", vmCompressedMemory)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmCompressedMemory)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_consumed_overhead_memory", vmConsumedOverheadMemory)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmConsumedOverheadMemory)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_distributed_cpu_entitlement", vmDistributedCpuEntitlement)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmDistributedCpuEntitlement)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_distributed_memory_entitlement", vmDistributedMemoryEntitlement)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmDistributedMemoryEntitlement)] = myMetric
	prometheus.MustRegister(myMetric)

	// metricName = fmt.Sprintf("%d_ft_latency_status", vmFtLatencyStatus)
	// log.Debugln("Key:", metricName)

	// myMetric = prometheus.NewGaugeVec(
	// 	prometheus.GaugeOpts{
	// 		Namespace: "vsphere",
	// 		Subsystem: "vm",
	// 		Name:      metricName,
	// 		Help:      metricName,
	// 	},
	// 	[]string{"datacenter"},
	// )
	// metricsMapVM[int(vmFtLatencyStatus)] = myMetric
	// prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_ft_log_bandwidth", vmFtLogBandwidth)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmFtLogBandwidth)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_ft_secondary_latency", vmFtSecondaryLatency)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmFtSecondaryLatency)] = myMetric
	prometheus.MustRegister(myMetric)

	// metricName = fmt.Sprintf("%d_guest_heartbeat_status", vmGuestHeartbeatStatus)
	// log.Debugln("Key:", metricName)

	// myMetric = prometheus.NewGaugeVec(
	// 	prometheus.GaugeOpts{
	// 		Namespace: "vsphere",
	// 		Subsystem: "vm",
	// 		Name:      metricName,
	// 		Help:      metricName,
	// 	},
	// 	[]string{"datacenter"},
	// )
	// metricsMapVM[int(vmGuestHeartbeatStatus)] = myMetric
	// prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_guest_memoruy_usage", vmGuestMemoryUsage)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmGuestMemoryUsage)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_host_memory_usage", vmHostMemoryUsage)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmHostMemoryUsage)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_overall_cpu_demand", vmOverallCpuDemand)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmOverallCpuDemand)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_overall_cpu_usage", vmOverallCpuUsage)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmOverallCpuUsage)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_private_memory", vmPrivateMemory)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmPrivateMemory)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_shared_memory", vmSharedMemory)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmSharedMemory)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_ssd_swapped_memory", vmSsdSwappedMemory)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmSsdSwappedMemory)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_static_cpu_entitlement", vmStaticCpuEntitlement)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmStaticCpuEntitlement)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_static_memory_entitlement", vmStaticMemoryEntitlement)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmStaticMemoryEntitlement)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_swapped_memory", vmSwappedMemory)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmSwappedMemory)] = myMetric
	prometheus.MustRegister(myMetric)

	metricName = fmt.Sprintf("%d_uptime_seconds", vmUptimeSeconds)
	log.Debugln("Key:", metricName)

	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vsphere",
			Subsystem: "vm",
			Name:      metricName,
			Help:      metricName,
		},
		[]string{"datacenter"},
	)
	metricsMapVM[int(vmUptimeSeconds)] = myMetric
	prometheus.MustRegister(myMetric)

	/*
		labels := []string{"datacenter", "vm"}

		metricName := fmt.Sprintf("%d_ballooned_memory", vmBalloonedMemory)
		log.Debugln("Key:", metricName)
		myMetric := prometheus.NewDesc(metricName, "ballooned memory", labels, nil)
		metricsMapVM[int(vmBalloonedMemory)] = myMetric

		metricName = fmt.Sprintf("%d_compressed_memory", vmCompressedMemory)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "compressed memory", labels, nil)
		metricsMapVM[int(vmCompressedMemory)] = myMetric

		metricName = fmt.Sprintf("%d_consumed_overhead_memory", vmConsumedOverheadMemory)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "consumed overhead memory", labels, nil)
		metricsMapVM[int(vmConsumedOverheadMemory)] = myMetric

		metricName = fmt.Sprintf("%d_distributed_cpu_entitlement", vmDistributedCpuEntitlement)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "distributed cpu entitlement", labels, nil)
		metricsMapVM[int(vmDistributedCpuEntitlement)] = myMetric

		metricName = fmt.Sprintf("%d_distributed_memory_entitlement", vmDistributedMemoryEntitlement)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "distributed memory entitlement", labels, nil)
		metricsMapVM[int(vmDistributedMemoryEntitlement)] = myMetric

		// metricName = fmt.Sprintf("%d_ft_latency_status", vmFtLatencyStatus)
		// log.Debugln("Key:", metricName)
		//myMetric = prometheus.NewDesc(metricName, "ft latency status", labels, nil)
		//metricsMapVM[int(vmFtLatencyStatus)] = myMetric

		metricName = fmt.Sprintf("%d_ft_log_bandwidth", vmFtLogBandwidth)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "ft log bandwidth", labels, nil)
		metricsMapVM[int(vmFtLogBandwidth)] = myMetric

		metricName = fmt.Sprintf("%d_ft_secondary_latency", vmFtSecondaryLatency)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "ft secondary latency", labels, nil)
		metricsMapVM[int(vmFtSecondaryLatency)] = myMetric

		// metricName = fmt.Sprintf("%d_guest_heartbeat_status", vmGuestHeartbeatStatus)
		// log.Debugln("Key:", metricName)
		//myMetric = prometheus.NewDesc(metricName, "guest heartbeat status", labels, nil)
		//metricsMapVM[int(vmGuestHeartbeatStatus)] = myMetric

		metricName = fmt.Sprintf("%d_guest_memoruy_usage", vmGuestMemoryUsage)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "guest memoruy usage", labels, nil)
		metricsMapVM[int(vmGuestMemoryUsage)] = myMetric

		metricName = fmt.Sprintf("%d_host_memory_usage", vmHostMemoryUsage)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "host memory usage", labels, nil)
		metricsMapVM[int(vmHostMemoryUsage)] = myMetric

		metricName = fmt.Sprintf("%d_overall_cpu_demand", vmOverallCpuDemand)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "overall cpu demand", labels, nil)
		metricsMapVM[int(vmOverallCpuDemand)] = myMetric

		metricName = fmt.Sprintf("%d_overall_cpu_usage", vmOverallCpuUsage)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "overall cpu usage", labels, nil)
		metricsMapVM[int(vmOverallCpuUsage)] = myMetric

		metricName = fmt.Sprintf("%d_private_memory", vmPrivateMemory)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "private memory", labels, nil)
		metricsMapVM[int(vmPrivateMemory)] = myMetric

		metricName = fmt.Sprintf("%d_shared_memory", vmSharedMemory)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "shared memory", labels, nil)
		metricsMapVM[int(vmSharedMemory)] = myMetric

		metricName = fmt.Sprintf("%d_ssd_swapped_memory", vmSsdSwappedMemory)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "ssd swapped memory", labels, nil)
		metricsMapVM[int(vmSsdSwappedMemory)] = myMetric

		metricName = fmt.Sprintf("%d_static_cpu_entitlement", vmStaticCpuEntitlement)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "static cpu entitlement", labels, nil)
		metricsMapVM[int(vmStaticCpuEntitlement)] = myMetric

		metricName = fmt.Sprintf("%d_static_memory_entitlement", vmStaticMemoryEntitlement)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "static memory entitlement", labels, nil)
		metricsMapVM[int(vmStaticMemoryEntitlement)] = myMetric

		metricName = fmt.Sprintf("%d_swapped_memory", vmSwappedMemory)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "swapped memory", labels, nil)
		metricsMapVM[int(vmSwappedMemory)] = myMetric

		metricName = fmt.Sprintf("%d_uptime_seconds", vmUptimeSeconds)
		log.Debugln("Key:", metricName)
		myMetric = prometheus.NewDesc(metricName, "uptime seconds", labels, nil)
		metricsMapVM[int(vmUptimeSeconds)] = myMetric
	*/

	log.Debugln("registerVMMetrics Succeeded")
	log.Debugln("registerVMMetrics LEAVE")

	return nil
}

//GetVSphereVMStats gets stats for an individual VM
func (c *Client) GetVSphereVMStats(w http.ResponseWriter, r *http.Request) error {
	log.Debugln("GetVSphereVMStats ENTER")

	vars := mux.Vars(r)
	datacenterStr := vars["datacenter"]
	log.Infoln("Datacenter:", datacenterStr)
	vmStr := vars["vm"]
	log.Infoln("VM:", vmStr)

	// Create client
	err := c.getClient()
	if err != nil {
		http.Error(w, "Unable connect to the vCenter Server", http.StatusGone)
		log.Errorln("getClient failed:", err)
		log.Debugln("GetVSphereVMStats LEAVE")

		return err
	}

	//find our objects
	finder := find.NewFinder(c.vClient.Client, false)

	dc, err := finder.Datacenter(*c.ctx, datacenterStr)
	if err != nil {
		http.Error(w, "Unable find the Datacener", http.StatusGone)
		log.Errorln("finder.Datacenter(", datacenterStr, "):", err)
		log.Debugln("GetVSphereVMStats LEAVE")
		return err
	}
	finder.SetDatacenter(dc)

	vm, err := finder.VirtualMachine(*c.ctx, vmStr)
	if err != nil {
		http.Error(w, "Unable find the VirtualMachine", http.StatusGone)
		log.Errorln("finder.VirtualMachine(", vmStr, "):", err)
		log.Debugln("GetVSphereVMStats LEAVE")
		return err
	}

	log.Infoln("VM:", vm.Name())
	log.Infoln("VM:", vm.InventoryPath)

	var oVM mo.VirtualMachine
	err = vm.Properties(*c.ctx, vm.Reference(), []string{"config", "summary"}, &oVM)
	if err != nil {
		http.Error(w, "Unable get the VirtualMachine properties", http.StatusBadRequest)
		log.Errorln("vm.Properties(", vmStr, "):", err)
		log.Debugln("GetVSphereVMStats LEAVE")
		return err
	}

	log.Infoln(oVM.Self.Value)
	log.Infoln(oVM.Summary.Config.Name)
	log.Infoln(oVM.Summary.Config.GuestFullName)
	log.Infoln(oVM.Summary.Config.GuestId)
	log.Infoln(oVM.Summary.Config.InstanceUuid)
	log.Infoln(string(oVM.Summary.OverallStatus))
	log.Infoln(string(oVM.OverallStatus))

	myMetric := metricsMapVM[vmBalloonedMemory]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.BalloonedMemory))
	}
	myMetric = metricsMapVM[vmCompressedMemory]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.CompressedMemory))
	}
	myMetric = metricsMapVM[vmConsumedOverheadMemory]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.ConsumedOverheadMemory))
	}
	myMetric = metricsMapVM[vmDistributedCpuEntitlement]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.DistributedCpuEntitlement))
	}
	myMetric = metricsMapVM[vmDistributedMemoryEntitlement]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.DistributedMemoryEntitlement))
	}
	// myMetric = metricsMapVM[vmFtLatencyStatus]
	// if myMetric != nil {
	// 	myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.FtLatencyStatus))
	// }
	myMetric = metricsMapVM[vmFtLogBandwidth]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.FtLogBandwidth))
	}
	myMetric = metricsMapVM[vmFtSecondaryLatency]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.FtSecondaryLatency))
	}
	// myMetric = metricsMapVM[vmGuestHeartbeatStatus]
	// if myMetric != nil {
	// 	myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.GuestHeartbeatStatus))
	// }
	myMetric = metricsMapVM[vmGuestMemoryUsage]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.GuestMemoryUsage))
	}
	myMetric = metricsMapVM[vmHostMemoryUsage]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.HostMemoryUsage))
	}
	myMetric = metricsMapVM[vmOverallCpuDemand]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.OverallCpuDemand))
	}
	myMetric = metricsMapVM[vmOverallCpuUsage]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.OverallCpuUsage))
	}
	myMetric = metricsMapVM[vmPrivateMemory]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.PrivateMemory))
	}
	myMetric = metricsMapVM[vmSharedMemory]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.SharedMemory))
	}
	myMetric = metricsMapVM[vmSsdSwappedMemory]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.SsdSwappedMemory))
	}
	myMetric = metricsMapVM[vmStaticCpuEntitlement]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.StaticCpuEntitlement))
	}
	myMetric = metricsMapVM[vmStaticMemoryEntitlement]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.StaticMemoryEntitlement))
	}
	myMetric = metricsMapVM[vmSwappedMemory]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.SwappedMemory))
	}
	myMetric = metricsMapVM[vmUptimeSeconds]
	if myMetric != nil {
		myMetric.WithLabelValues(datacenterStr).Set(float64(oVM.Summary.QuickStats.UptimeSeconds))
	}

	/*
		myMetric := metricsMapVM[vmBalloonedMemory]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.BalloonedMemory),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmCompressedMemory]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.CompressedMemory),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmConsumedOverheadMemory]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.ConsumedOverheadMemory),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmDistributedCpuEntitlement]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.DistributedCpuEntitlement),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmDistributedMemoryEntitlement]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.DistributedMemoryEntitlement),
				datacenterStr, vmStr)
		}
		// myMetric = metricsMapVM[vmFtLatencyStatus]
		// if myMetric != nil {
		//	prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.FtLatencyStatus),
		//		datacenterStr, vmStr)
		// }
		myMetric = metricsMapVM[vmFtLogBandwidth]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.FtLogBandwidth),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmFtSecondaryLatency]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.FtSecondaryLatency),
				datacenterStr, vmStr)
		}
		// myMetric = metricsMapVM[vmGuestHeartbeatStatus]
		// if myMetric != nil {
		//	prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.GuestHeartbeatStatus),
		//		datacenterStr, vmStr)
		// }
		myMetric = metricsMapVM[vmGuestMemoryUsage]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.GuestMemoryUsage),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmHostMemoryUsage]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.HostMemoryUsage),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmOverallCpuDemand]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.OverallCpuDemand),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmOverallCpuUsage]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.OverallCpuUsage),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmPrivateMemory]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.PrivateMemory),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmSharedMemory]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.SharedMemory),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmSsdSwappedMemory]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.SsdSwappedMemory),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmStaticCpuEntitlement]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.StaticCpuEntitlement),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmStaticMemoryEntitlement]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.StaticMemoryEntitlement),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmSwappedMemory]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.SwappedMemory),
				datacenterStr, vmStr)
		}
		myMetric = metricsMapVM[vmUptimeSeconds]
		if myMetric != nil {
			prometheus.MustNewConstMetric(myMetric, prometheus.GaugeValue, float64(oVM.Summary.QuickStats.UptimeSeconds),
				datacenterStr, vmStr)
		}
	*/

	log.Debugln("GetVSphereVMStats Succeeded")
	log.Debugln("GetVSphereVMStats LEAVE")

	return nil
}
