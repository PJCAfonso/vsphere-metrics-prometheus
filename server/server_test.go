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
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	assert "github.com/stretchr/testify/assert"

	config "github.com/dvonthenen/vsphere-metrics-prometheus/config"
	"github.com/dvonthenen/vsphere-metrics-prometheus/types"
)

var server *RestServer

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	//create config object
	cfg := config.NewConfig()

	server = NewRestServer(cfg)

	//Run is a blocking call for Negroni... so go routine it
	go func() {
		server.Server.Run(":" + strconv.Itoa(cfg.RestPort))
	}()

	//wait 5 seconds for server to come up
	time.Sleep(5 * time.Second)

	//log.Infoln("Start tests")
	m.Run()
}

func TestVersion(t *testing.T) {
	url := "http://127.0.0.1:" +
		strconv.Itoa(server.Config.RestPort) + "/version"

	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	assert.NotNil(t, body)
	assert.NoError(t, err)

	var ver types.Version
	err = json.Unmarshal(body, &ver)
	assert.NotNil(t, ver)
	assert.NoError(t, err)

	assert.Equal(t, ver.VersionInt, config.VersionInt)
	assert.Equal(t, ver.VersionStr, config.VersionStr)
}
