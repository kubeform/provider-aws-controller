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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cognito/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type CognitoV1alpha1Interface interface {
	RESTClient() rest.Interface
	IdentityPoolsGetter
	IdentityPoolProviderPrincipalTagsGetter
	IdentityPoolRolesAttachmentsGetter
	IdentityProvidersGetter
	ResourceServersGetter
	UsersGetter
	UserGroupsGetter
	UserPoolsGetter
	UserPoolClientsGetter
	UserPoolDomainsGetter
	UserPoolUiCustomizationsGetter
}

// CognitoV1alpha1Client is used to interact with features provided by the cognito.aws.kubeform.com group.
type CognitoV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CognitoV1alpha1Client) IdentityPools(namespace string) IdentityPoolInterface {
	return newIdentityPools(c, namespace)
}

func (c *CognitoV1alpha1Client) IdentityPoolProviderPrincipalTags(namespace string) IdentityPoolProviderPrincipalTagInterface {
	return newIdentityPoolProviderPrincipalTags(c, namespace)
}

func (c *CognitoV1alpha1Client) IdentityPoolRolesAttachments(namespace string) IdentityPoolRolesAttachmentInterface {
	return newIdentityPoolRolesAttachments(c, namespace)
}

func (c *CognitoV1alpha1Client) IdentityProviders(namespace string) IdentityProviderInterface {
	return newIdentityProviders(c, namespace)
}

func (c *CognitoV1alpha1Client) ResourceServers(namespace string) ResourceServerInterface {
	return newResourceServers(c, namespace)
}

func (c *CognitoV1alpha1Client) Users(namespace string) UserInterface {
	return newUsers(c, namespace)
}

func (c *CognitoV1alpha1Client) UserGroups(namespace string) UserGroupInterface {
	return newUserGroups(c, namespace)
}

func (c *CognitoV1alpha1Client) UserPools(namespace string) UserPoolInterface {
	return newUserPools(c, namespace)
}

func (c *CognitoV1alpha1Client) UserPoolClients(namespace string) UserPoolClientInterface {
	return newUserPoolClients(c, namespace)
}

func (c *CognitoV1alpha1Client) UserPoolDomains(namespace string) UserPoolDomainInterface {
	return newUserPoolDomains(c, namespace)
}

func (c *CognitoV1alpha1Client) UserPoolUiCustomizations(namespace string) UserPoolUiCustomizationInterface {
	return newUserPoolUiCustomizations(c, namespace)
}

// NewForConfig creates a new CognitoV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CognitoV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CognitoV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CognitoV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CognitoV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CognitoV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CognitoV1alpha1Client {
	return &CognitoV1alpha1Client{c}
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
func (c *CognitoV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
