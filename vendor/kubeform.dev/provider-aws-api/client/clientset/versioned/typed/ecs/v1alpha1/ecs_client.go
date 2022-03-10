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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecs/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type EcsV1alpha1Interface interface {
	RESTClient() rest.Interface
	AccountSettingDefaultsGetter
	CapacityProvidersGetter
	ClustersGetter
	ClusterCapacityProvidersesGetter
	ServicesGetter
	TagsGetter
	TaskDefinitionsGetter
	TaskSetsGetter
}

// EcsV1alpha1Client is used to interact with features provided by the ecs.aws.kubeform.com group.
type EcsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *EcsV1alpha1Client) AccountSettingDefaults(namespace string) AccountSettingDefaultInterface {
	return newAccountSettingDefaults(c, namespace)
}

func (c *EcsV1alpha1Client) CapacityProviders(namespace string) CapacityProviderInterface {
	return newCapacityProviders(c, namespace)
}

func (c *EcsV1alpha1Client) Clusters(namespace string) ClusterInterface {
	return newClusters(c, namespace)
}

func (c *EcsV1alpha1Client) ClusterCapacityProviderses(namespace string) ClusterCapacityProvidersInterface {
	return newClusterCapacityProviderses(c, namespace)
}

func (c *EcsV1alpha1Client) Services(namespace string) ServiceInterface {
	return newServices(c, namespace)
}

func (c *EcsV1alpha1Client) Tags(namespace string) TagInterface {
	return newTags(c, namespace)
}

func (c *EcsV1alpha1Client) TaskDefinitions(namespace string) TaskDefinitionInterface {
	return newTaskDefinitions(c, namespace)
}

func (c *EcsV1alpha1Client) TaskSets(namespace string) TaskSetInterface {
	return newTaskSets(c, namespace)
}

// NewForConfig creates a new EcsV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*EcsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &EcsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new EcsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *EcsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new EcsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *EcsV1alpha1Client {
	return &EcsV1alpha1Client{c}
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
func (c *EcsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
