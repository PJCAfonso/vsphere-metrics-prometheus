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
	"context"
	"errors"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"

	config "github.com/dvonthenen/vsphere-metrics-prometheus/config"
)

var (
	//ErrClientParamsNil - The govmomi client parameters are nil. Need to re-init.
	ErrClientParamsNil = errors.New("The govmomi client parameters are nil. Need to re-init")
)

//Client representation for a REST API server
type Client struct {
	config  *config.Config
	ctx     *context.Context
	vClient *govmomi.Client
}

//NewClient generates a new VSphere client
func NewClient(cfg *config.Config) *Client {
	client := &Client{
		config:  cfg,
		vClient: nil,
		ctx:     nil,
	}

	return client
}

//RegisterMetrics performs a Prometheus registration for all metrics
func (c *Client) RegisterMetrics() error {
	log.Debugln("RegisterMetrics ENTER")

	err := c.registerEsxMetrics()
	if err != nil {
		log.Debugln("registerEsxMetrics Failed:", err)
		log.Debugln("RegisterMetrics LEAVE")
		return err
	}

	err = c.registerDatastoreMetrics()
	if err != nil {
		log.Debugln("registerDatastoreMetrics Failed:", err)
		log.Debugln("RegisterMetrics LEAVE")
		return err
	}

	err = c.registerVMMetrics()
	if err != nil {
		log.Debugln("registerVMMetrics Failed:", err)
		log.Debugln("RegisterMetrics LEAVE")
		return err
	}

	log.Debugln("RegisterMetrics Succeeded")
	log.Debugln("RegisterMetrics LEAVE")
	return nil
}

//IsValid returns an error if the connection isnt valid
func (c *Client) IsValid() error {
	log.Debugln("IsSessionValid ENTER")

	if c.vClient == nil || c.ctx == nil {
		log.Debugln("IsSessionValid Failed. client or ctx is nil")
		log.Debugln("IsSessionValid LEAVE")
		return ErrClientParamsNil
	}

	var mgr mo.SessionManager
	err := mo.RetrieveProperties(context.Background(), c.vClient, c.vClient.ServiceContent.PropertyCollector, *c.vClient.ServiceContent.SessionManager, &mgr)
	if err != nil {
		log.Debugln("IsSessionValid Failed:", err)
		log.Debugln("IsSessionValid LEAVE")
		return err
	}

	log.Debugln("IsSessionValid Succeeded")
	log.Debugln("IsSessionValid LEAVE")
	return nil
}

//Logout performs a logout of the govmomi client
func (c *Client) Logout() {
	log.Debugln("Logout ENTER")

	if c.vClient == nil || c.ctx == nil {
		log.Debugln("client or ctx is nil. No need to logout.")
		log.Debugln("Logout LEAVE")
		return
	}

	err := c.vClient.Logout(*c.ctx)
	if err != nil {
		log.Warnln("Logout Failed:", err)
	}

	c.ctx = nil
	c.vClient = nil

	log.Debugln("Logout Succeeded")
	log.Debugln("Logout LEAVE")
}

func (c *Client) getClient() error {
	log.Debugln("getClient ENTER")

	// Does a connection already exist? Then reuse it!
	if c.IsValid() == nil {
		log.Infoln("Reusing vSphere Client")
		log.Debugln("getClient LEAVE")
		return nil
	}

	//Logout out just in case and clear out state in Client
	c.Logout()

	// Default context
	ctx := context.Background()
	c.ctx = &ctx

	// Setup URL object
	urlBase := "https://username:password@host/sdk"
	u, err := soap.ParseURL(urlBase)
	if err != nil {
		log.Infoln("Failed to parse URL object")
		log.Debugln("getClient LEAVE")
		return err
	}

	u.User = url.UserPassword(c.config.VSphereUser, c.config.VSpherePass)
	if c.config.VSpherePort > 0 {
		u.Host = c.config.VSphereHostname + ":" + strconv.Itoa(c.config.VSpherePort)
	} else {
		u.Host = c.config.VSphereHostname
	}

	log.Debugln("ConnectionStr:", u.String())
	log.Debugln("Insecure:", c.config.VSphereInsecure)

	// Connect and login to ESX or vCenter
	c.vClient, err = govmomi.NewClient(ctx, u, c.config.VSphereInsecure)
	if err != nil {
		log.Infoln("govmomi.NewClient failed:", err)
		log.Debugln("getClient LEAVE")
		return err
	}

	log.Debugln("getClient Succeeded")
	log.Debugln("getClient LEAVE")

	return nil
}
