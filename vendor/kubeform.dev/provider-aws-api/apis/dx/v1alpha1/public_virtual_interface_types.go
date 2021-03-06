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

type PublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            PublicVirtualInterfaceStatus `json:"status,omitempty"`
}

type PublicVirtualInterfaceSpec struct {
	State *PublicVirtualInterfaceSpecResource `json:"state,omitempty" tf:"-"`

	Resource PublicVirtualInterfaceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PublicVirtualInterfaceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddressFamily *string `json:"addressFamily" tf:"address_family"`
	// +optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address"`
	// +optional
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device"`
	BgpAsn    *int64  `json:"bgpAsn" tf:"bgp_asn"`
	// +optional
	BgpAuthKey   *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key"`
	ConnectionID *string `json:"connectionID" tf:"connection_id"`
	// +optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address"`
	Name            *string `json:"name" tf:"name"`
	// +kubebuilder:validation:MinItems=1
	RouteFilterPrefixes []string `json:"routeFilterPrefixes" tf:"route_filter_prefixes"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	Vlan    *int64             `json:"vlan" tf:"vlan"`
}

type PublicVirtualInterfaceStatus struct {
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

// PublicVirtualInterfaceList is a list of PublicVirtualInterfaces
type PublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PublicVirtualInterface CRD objects
	Items []PublicVirtualInterface `json:"items,omitempty"`
}
