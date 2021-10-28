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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ClientVPNEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClientVPNEndpointSpec   `json:"spec,omitempty"`
	Status            ClientVPNEndpointStatus `json:"status,omitempty"`
}

type ClientVPNEndpointSpecAuthenticationOptions struct {
	// +optional
	ActiveDirectoryID *string `json:"activeDirectoryID,omitempty" tf:"active_directory_id"`
	// +optional
	RootCertificateChainArn *string `json:"rootCertificateChainArn,omitempty" tf:"root_certificate_chain_arn"`
	// +optional
	SamlProviderArn *string `json:"samlProviderArn,omitempty" tf:"saml_provider_arn"`
	Type            *string `json:"type" tf:"type"`
}

type ClientVPNEndpointSpecConnectionLogOptions struct {
	// +optional
	CloudwatchLogGroup *string `json:"cloudwatchLogGroup,omitempty" tf:"cloudwatch_log_group"`
	// +optional
	CloudwatchLogStream *string `json:"cloudwatchLogStream,omitempty" tf:"cloudwatch_log_stream"`
	Enabled             *bool   `json:"enabled" tf:"enabled"`
}

type ClientVPNEndpointSpec struct {
	State *ClientVPNEndpointSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClientVPNEndpointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClientVPNEndpointSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +kubebuilder:validation:MaxItems=2
	AuthenticationOptions []ClientVPNEndpointSpecAuthenticationOptions `json:"authenticationOptions" tf:"authentication_options"`
	ClientCIDRBlock       *string                                      `json:"clientCIDRBlock" tf:"client_cidr_block"`
	ConnectionLogOptions  *ClientVPNEndpointSpecConnectionLogOptions   `json:"connectionLogOptions" tf:"connection_log_options"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DnsName *string `json:"dnsName,omitempty" tf:"dns_name"`
	// +optional
	DnsServers           []string `json:"dnsServers,omitempty" tf:"dns_servers"`
	ServerCertificateArn *string  `json:"serverCertificateArn" tf:"server_certificate_arn"`
	// +optional
	SplitTunnel *bool `json:"splitTunnel,omitempty" tf:"split_tunnel"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	TransportProtocol *string `json:"transportProtocol,omitempty" tf:"transport_protocol"`
}

type ClientVPNEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ClientVPNEndpointList is a list of ClientVPNEndpoints
type ClientVPNEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClientVPNEndpoint CRD objects
	Items []ClientVPNEndpoint `json:"items,omitempty"`
}
