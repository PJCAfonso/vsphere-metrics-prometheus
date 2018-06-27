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

package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	negroni "github.com/urfave/negroni"

	config "github.com/dvonthenen/vsphere-metrics-prometheus/config"
	"github.com/dvonthenen/vsphere-metrics-prometheus/vsphere"
)

//RestServer representation for a REST API server
type RestServer struct {
	Config  *config.Config
	vClient *vsphere.Client
	Server  *negroni.Negroni
}

//NewRestServer generates a new REST API server
func NewRestServer(cfg *config.Config) *RestServer {
	restServer := &RestServer{
		Config:  cfg,
		vClient: vsphere.NewClient(cfg),
	}

	err := restServer.vClient.RegisterMetrics()
	if err != nil {
		log.Errorln("registerMetrics Failed:", err)
		return nil
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		err := restServer.getVersion(w, r)
		if err != nil {
			log.Errorln("getVersion Failed:", err)
		}
	}).Methods("GET")
	if cfg.VSphereType == string(config.VSphereRoleEsx) {
		mux.HandleFunc("/datacenter/{datacenter}/host/{host}/metrics", func(w http.ResponseWriter, r *http.Request) {
			err := restServer.vClient.GetVSphereEsxStats(w, r)
			if err != nil {
				log.Errorln("getVSphereStats Failed:", err)
			}
			promhttp.Handler().ServeHTTP(w, r)
		}).Methods("GET")
	} else if cfg.VSphereType == string(config.VSphereRoleDatastore) {
		mux.HandleFunc("/datacenter/{datacenter}/datastore/{datastore}/metrics", func(w http.ResponseWriter, r *http.Request) {
			err := restServer.vClient.GetVSphereDatastoreStats(w, r)
			if err != nil {
				log.Errorln("GetVSphereDatastoreStats Failed:", err)
			}
			promhttp.Handler().ServeHTTP(w, r)
		}).Methods("GET")
	} else if cfg.VSphereType == string(config.VSphereRoleVirtualMachine) {
		mux.HandleFunc("/datacenter/{datacenter}/vm/{vm}/metrics", func(w http.ResponseWriter, r *http.Request) {
			err := restServer.vClient.GetVSphereVMStats(w, r)
			if err != nil {
				log.Errorln("GetVSphereVMStats Failed:", err)
			}
			promhttp.Handler().ServeHTTP(w, r)
		}).Methods("GET")
	}

	server := negroni.Classic()
	server.UseHandler(mux)

	restServer.Server = server

	return restServer
}
