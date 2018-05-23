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

package config

import (
	"flag"
	"strconv"
)

//consts exported out of package
const (
	//VersionInt in INT form
	VersionInt = 1

	//VersionStr in string form
	VersionStr = "0.0.1"

	//DefaultRestPort REST API port
	DefaultRestPort = 9444

	//Default vSphere port
	DefaultVSpherePort = 0
)

//Config is the representation of the config
type Config struct {
	LogLevel string
	Debug    bool

	//RestEndPoint string
	RestPort int

	VSphereHostname string
	VSpherePort     int
	VSphereInsecure bool
	VSphereUser     string
	VSpherePass     string
}

//AddFlags adds flags to the command line parsing
func (cfg *Config) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.LogLevel, "loglevel", cfg.LogLevel, "Set the logging level")
	fs.BoolVar(&cfg.Debug, "debug", cfg.Debug, "Debug mode")

	//fs.StringVar(&cfg.RestEndPoint, "rest.endpoint", cfg.RestEndPoint, "Address for REST endpoint")
	fs.IntVar(&cfg.RestPort, "rest.port", cfg.RestPort, "Port to serve up REST endpoint")

	fs.StringVar(&cfg.VSphereHostname, "vsphere.hostname", cfg.VSphereHostname, "vCenter Server hostname")
	fs.IntVar(&cfg.VSpherePort, "vsphere.port", cfg.VSpherePort, "vCenter Server port")
	fs.BoolVar(&cfg.VSphereInsecure, "vsphere.insecure", cfg.VSphereInsecure, "vCenter Server insecure mode")
	fs.StringVar(&cfg.VSphereUser, "vsphere.username", cfg.VSphereUser, "vCenter Server Username")
	fs.StringVar(&cfg.VSpherePass, "vsphere.password", cfg.VSpherePass, "vCenter Server Password")
}

//NewConfig creates a new Config object
func NewConfig() *Config {
	/*
		ip, err := autoDiscoverIP()
		if err != nil {
			ip = "127.0.0.1"
		}
	*/

	return &Config{
		LogLevel: env("LOG_LEVEL", "info"),
		Debug:    envBool("DEBUG", "false"),
		//RestEndPoint:    env("REST_ENDPOINT", ip),
		RestPort:        envInt("REST_PORT", strconv.Itoa(DefaultRestPort)),
		VSphereHostname: env("VSPHERE_HOSTNAME", ""),
		VSpherePort:     envInt("VSPHERE_PORT", strconv.Itoa(DefaultVSpherePort)),
		VSphereInsecure: envBool("VSPHERE_INSECURE", "false"),
		VSphereUser:     env("VSPHERE_USERNAME", ""),
		VSpherePass:     env("VSPHERE_PASSWORD", ""),
	}
}

//autoDiscoverIP attempt to discover the IP for this host
/*
func autoDiscoverIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Warnln("Failed to get Interfaces", err)
		return "", err
	}

	var ip string
	for _, i := range ifaces {
		if strings.Contains(i.Name, "lo") || strings.Contains(i.Name, "docker") {
			log.Debugln("Skipping interface:", i.Name)
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			log.Infoln("Failed to get IPs on Interface", err)
			continue
		}
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP.String()
				log.Debugln("IPNet:", ip)
			case *net.IPAddr:
				ip = v.IP.String()
				log.Debugln("IPAddr:", ip)
			}
			if len(ip) > 0 {
				break
			}
		}

		log.Infoln("IP Discovered:", ip)
		break
	}

	return ip, nil
}
*/
