// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkingV1beta1 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1beta1) DestinationRules(namespace string) v1beta1.DestinationRuleInterface {
	return &FakeDestinationRules{c, namespace}
}

func (c *FakeNetworkingV1beta1) Gateways(namespace string) v1beta1.GatewayInterface {
	return &FakeGateways{c, namespace}
}

func (c *FakeNetworkingV1beta1) ProxyConfigs(namespace string) v1beta1.ProxyConfigInterface {
	return &FakeProxyConfigs{c, namespace}
}

func (c *FakeNetworkingV1beta1) ServiceEntries(namespace string) v1beta1.ServiceEntryInterface {
	return &FakeServiceEntries{c, namespace}
}

func (c *FakeNetworkingV1beta1) ServiceNameMappings(namespace string) v1beta1.ServiceNameMappingInterface {
	return &FakeServiceNameMappings{c, namespace}
}

func (c *FakeNetworkingV1beta1) Sidecars(namespace string) v1beta1.SidecarInterface {
	return &FakeSidecars{c, namespace}
}

func (c *FakeNetworkingV1beta1) VirtualServices(namespace string) v1beta1.VirtualServiceInterface {
	return &FakeVirtualServices{c, namespace}
}

func (c *FakeNetworkingV1beta1) WorkloadEntries(namespace string) v1beta1.WorkloadEntryInterface {
	return &FakeWorkloadEntries{c, namespace}
}

func (c *FakeNetworkingV1beta1) WorkloadGroups(namespace string) v1beta1.WorkloadGroupInterface {
	return &FakeWorkloadGroups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkingV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
