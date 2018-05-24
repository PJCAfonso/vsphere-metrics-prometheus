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

package main

import (
	"flag"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dvonthenen/vsphere-metrics-prometheus/config"
	"github.com/dvonthenen/vsphere-metrics-prometheus/server"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.Infoln("Initializing the Prometheus metric collection for vSphere...")
}

func main() {
	cfg := config.NewConfig()
	fs := flag.NewFlagSet("vsphere-proxy", flag.ExitOnError)
	cfg.AddFlags(fs)
	fs.Parse(os.Args[1:])

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Warnln("Invalid log level. Defaulting to info.")
		level = log.InfoLevel
	} else {
		log.Infoln("Set logging to", cfg.LogLevel)
	}
	log.SetLevel(level)

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		log.Debugln(pair[0], "=", pair[1])
	}

	restServer := server.NewRestServer(cfg)
	restServer.Server.Run(":" + strconv.Itoa(cfg.RestPort))
}
