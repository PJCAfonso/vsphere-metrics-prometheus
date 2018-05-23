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
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dvonthenen/vsphere-metrics-prometheus/config"
	"github.com/dvonthenen/vsphere-metrics-prometheus/types"
)

func (s *RestServer) getVersion(w http.ResponseWriter, r *http.Request) error {
	log.Debugln("getVersion ENTER")

	ver := types.Version{
		VersionInt: config.VersionInt,
		VersionStr: config.VersionStr,
	}

	response, err := json.MarshalIndent(ver, "", "  ")
	if err != nil {
		http.Error(w, "Unable to marshall the response", http.StatusBadRequest)
		log.Debugln("getVersion LEAVE")
		return err
	}

	log.Debugln("response:", string(response))
	fmt.Fprintf(w, string(response))

	log.Debugln("getVersion Succeeded")
	log.Debugln("getVersion LEAVE")

	return nil
}
