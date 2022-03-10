/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-aws-api/apis/network/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type NetworkV1alpha1Interface interface {
	RESTClient() rest.Interface
	AclsGetter
	AclAssociationsGetter
	AclRulesGetter
	InterfacesGetter
	InterfaceAttachmentsGetter
	InterfaceSgAttachmentsGetter
}

// NetworkV1alpha1Client is used to interact with features provided by the network.aws.kubeform.com group.
type NetworkV1alpha1Client struct {
	restClient rest.Interface
}

func (c *NetworkV1alpha1Client) Acls(namespace string) AclInterface {
	return newAcls(c, namespace)
}

func (c *NetworkV1alpha1Client) AclAssociations(namespace string) AclAssociationInterface {
	return newAclAssociations(c, namespace)
}

func (c *NetworkV1alpha1Client) AclRules(namespace string) AclRuleInterface {
	return newAclRules(c, namespace)
}

func (c *NetworkV1alpha1Client) Interfaces(namespace string) InterfaceInterface {
	return newInterfaces(c, namespace)
}

func (c *NetworkV1alpha1Client) InterfaceAttachments(namespace string) InterfaceAttachmentInterface {
	return newInterfaceAttachments(c, namespace)
}

func (c *NetworkV1alpha1Client) InterfaceSgAttachments(namespace string) InterfaceSgAttachmentInterface {
	return newInterfaceSgAttachments(c, namespace)
}

// NewForConfig creates a new NetworkV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*NetworkV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NetworkV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new NetworkV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NetworkV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NetworkV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *NetworkV1alpha1Client {
	return &NetworkV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *NetworkV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
