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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/apigateway/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type ApigatewayV1alpha1Interface interface {
	RESTClient() rest.Interface
	AccountsGetter
	ApiKeysGetter
	ApigatewayResourcesGetter
	AuthorizersGetter
	BasePathMappingsGetter
	ClientCertificatesGetter
	DeploymentsGetter
	DocumentationPartsGetter
	DocumentationVersionsGetter
	DomainNamesGetter
	GatewayResponsesGetter
	IntegrationsGetter
	IntegrationResponsesGetter
	MethodsGetter
	MethodResponsesGetter
	MethodSettingsesGetter
	ModelsGetter
	RequestValidatorsGetter
	RestAPIsGetter
	RestAPIPoliciesGetter
	StagesGetter
	UsagePlansGetter
	UsagePlanKeysGetter
	VpcLinksGetter
}

// ApigatewayV1alpha1Client is used to interact with features provided by the apigateway.aws.kubeform.com group.
type ApigatewayV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ApigatewayV1alpha1Client) Accounts(namespace string) AccountInterface {
	return newAccounts(c, namespace)
}

func (c *ApigatewayV1alpha1Client) ApiKeys(namespace string) ApiKeyInterface {
	return newApiKeys(c, namespace)
}

func (c *ApigatewayV1alpha1Client) ApigatewayResources(namespace string) ApigatewayResourceInterface {
	return newApigatewayResources(c, namespace)
}

func (c *ApigatewayV1alpha1Client) Authorizers(namespace string) AuthorizerInterface {
	return newAuthorizers(c, namespace)
}

func (c *ApigatewayV1alpha1Client) BasePathMappings(namespace string) BasePathMappingInterface {
	return newBasePathMappings(c, namespace)
}

func (c *ApigatewayV1alpha1Client) ClientCertificates(namespace string) ClientCertificateInterface {
	return newClientCertificates(c, namespace)
}

func (c *ApigatewayV1alpha1Client) Deployments(namespace string) DeploymentInterface {
	return newDeployments(c, namespace)
}

func (c *ApigatewayV1alpha1Client) DocumentationParts(namespace string) DocumentationPartInterface {
	return newDocumentationParts(c, namespace)
}

func (c *ApigatewayV1alpha1Client) DocumentationVersions(namespace string) DocumentationVersionInterface {
	return newDocumentationVersions(c, namespace)
}

func (c *ApigatewayV1alpha1Client) DomainNames(namespace string) DomainNameInterface {
	return newDomainNames(c, namespace)
}

func (c *ApigatewayV1alpha1Client) GatewayResponses(namespace string) GatewayResponseInterface {
	return newGatewayResponses(c, namespace)
}

func (c *ApigatewayV1alpha1Client) Integrations(namespace string) IntegrationInterface {
	return newIntegrations(c, namespace)
}

func (c *ApigatewayV1alpha1Client) IntegrationResponses(namespace string) IntegrationResponseInterface {
	return newIntegrationResponses(c, namespace)
}

func (c *ApigatewayV1alpha1Client) Methods(namespace string) MethodInterface {
	return newMethods(c, namespace)
}

func (c *ApigatewayV1alpha1Client) MethodResponses(namespace string) MethodResponseInterface {
	return newMethodResponses(c, namespace)
}

func (c *ApigatewayV1alpha1Client) MethodSettingses(namespace string) MethodSettingsInterface {
	return newMethodSettingses(c, namespace)
}

func (c *ApigatewayV1alpha1Client) Models(namespace string) ModelInterface {
	return newModels(c, namespace)
}

func (c *ApigatewayV1alpha1Client) RequestValidators(namespace string) RequestValidatorInterface {
	return newRequestValidators(c, namespace)
}

func (c *ApigatewayV1alpha1Client) RestAPIs(namespace string) RestAPIInterface {
	return newRestAPIs(c, namespace)
}

func (c *ApigatewayV1alpha1Client) RestAPIPolicies(namespace string) RestAPIPolicyInterface {
	return newRestAPIPolicies(c, namespace)
}

func (c *ApigatewayV1alpha1Client) Stages(namespace string) StageInterface {
	return newStages(c, namespace)
}

func (c *ApigatewayV1alpha1Client) UsagePlans(namespace string) UsagePlanInterface {
	return newUsagePlans(c, namespace)
}

func (c *ApigatewayV1alpha1Client) UsagePlanKeys(namespace string) UsagePlanKeyInterface {
	return newUsagePlanKeys(c, namespace)
}

func (c *ApigatewayV1alpha1Client) VpcLinks(namespace string) VpcLinkInterface {
	return newVpcLinks(c, namespace)
}

// NewForConfig creates a new ApigatewayV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ApigatewayV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ApigatewayV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ApigatewayV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ApigatewayV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ApigatewayV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ApigatewayV1alpha1Client {
	return &ApigatewayV1alpha1Client{c}
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
func (c *ApigatewayV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
