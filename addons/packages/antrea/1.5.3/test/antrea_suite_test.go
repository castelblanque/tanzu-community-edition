// Copyright 2021 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package antrea_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type AntreaConfig struct {
	InfraProvider string `yaml:"infraProvider"`
	Antrea        struct {
		Config struct {
			ServiceCIDR           string `yaml:"serviceCIDR"`
			ServiceCIDRv6         string `yaml:"serviceCIDRv6"`
			TunnelType            string `yaml:"tunnelType"`
			TrafficEncryptionMode string `yaml:"trafficEncryptionMode"`
			WireGuard             struct {
				Port int `yaml:"port,omitempty"`
			} `yaml:"wireGuard"`
			TrafficEncapMode        string `yaml:"trafficEncapMode"`
			NoSNAT                  bool   `yaml:"noSNAT"`
			DisableUDPTunnelOffload bool   `yaml:"disableUdpTunnelOffload"`
			DefaultMTU              string `yaml:"defaultMTU"`
			TLSCipherSuites         string `yaml:"tlsCipherSuites"`
			FeatureGates            struct {
				AntreaProxy        bool `yaml:"AntreaProxy"`
				EndpointSlice      bool `yaml:"EndpointSlice"`
				AntreaTraceflow    bool `yaml:"AntreaTraceflow"`
				NodePortLocal      bool `yaml:"NodePortLocal"`
				AntreaPolicy       bool `yaml:"AntreaPolicy"`
				FlowExporter       bool `yaml:"FlowExporter"`
				NetworkPolicyStats bool `yaml:"NetworkPolicyStats"`
				Egress             bool `yaml:"Egress"`
				AntreaIPAM         bool `yaml:"AntreaIPAM"`
				ServiceExternalIP  bool `yaml:"ServiceExternalIP"`
				Multicast          bool `yaml:"Multicast"`
			} `yaml:"featureGates"`
			NodePortLocal struct {
				Enabled   bool   `yaml:"enabled"`
				PortRange string `yaml:"portRange"`
			} `yaml:"nodePortLocal"`
			FlowExporter struct {
				Address           string `yaml:"address"`
				PollInterval      string `yaml:"pollInterval"`
				ExportTimeout     string `yaml:"exportTimeout"`
				IdleExportTimeout string `yaml:"idleExportTimeout"`
			} `yaml:"flowExporter"`
			KubeAPIServerOverride   string   `yaml:"kubeAPIServerOverride,omitempty"`
			TransportInterface      string   `yaml:"transportInterface,omitempty"`
			TransportInterfaceCIDRs []string `yaml:"transportInterfaceCIDRs,omitempty"`
			MulticastInterfaces     []string `yaml:"multicastInterfaces,omitempty"`
			EnableUsageReporting    bool     `yaml:"enableUsageReporting"`
			AntreaProxy             struct {
				ProxyAll             bool     `yaml:"proxyAll"`
				NodePortAddresses    []string `yaml:"nodePortAddresses"`
				SkipServices         []string `yaml:"skipServices"`
				ProxyLoadBalancerIPS bool     `yaml:"proxyLoadBalancerIPs"`
			} `yaml:"antreaProxy"`
			Egress struct {
				ExceptCIDRs []string `yaml:"exceptCIDRs"`
			} `yaml:"egress"`
		} `yaml:"config"`
	} `yaml:"antrea"`
}

func TestAntrea(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Antrea Addons Templates Suite")
}
