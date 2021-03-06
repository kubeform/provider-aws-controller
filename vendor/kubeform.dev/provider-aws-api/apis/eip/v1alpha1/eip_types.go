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

type Eip struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EipSpec   `json:"spec,omitempty"`
	Status            EipStatus `json:"status,omitempty"`
}

type EipSpec struct {
	State *EipSpecResource `json:"state,omitempty" tf:"-"`

	Resource EipSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EipSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// +optional
	AllocationID *string `json:"allocationID,omitempty" tf:"allocation_id"`
	// +optional
	AssociateWithPrivateIP *string `json:"associateWithPrivateIP,omitempty" tf:"associate_with_private_ip"`
	// +optional
	AssociationID *string `json:"associationID,omitempty" tf:"association_id"`
	// +optional
	CarrierIP *string `json:"carrierIP,omitempty" tf:"carrier_ip"`
	// +optional
	CustomerOwnedIP *string `json:"customerOwnedIP,omitempty" tf:"customer_owned_ip"`
	// +optional
	CustomerOwnedIpv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool"`
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// +optional
	Instance *string `json:"instance,omitempty" tf:"instance"`
	// +optional
	NetworkBorderGroup *string `json:"networkBorderGroup,omitempty" tf:"network_border_group"`
	// +optional
	NetworkInterface *string `json:"networkInterface,omitempty" tf:"network_interface"`
	// +optional
	PrivateDNS *string `json:"privateDNS,omitempty" tf:"private_dns"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	PublicDNS *string `json:"publicDNS,omitempty" tf:"public_dns"`
	// +optional
	PublicIP *string `json:"publicIP,omitempty" tf:"public_ip"`
	// +optional
	PublicIpv4Pool *string `json:"publicIpv4Pool,omitempty" tf:"public_ipv4_pool"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Vpc *bool `json:"vpc,omitempty" tf:"vpc"`
}

type EipStatus struct {
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

// EipList is a list of Eips
type EipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Eip CRD objects
	Items []Eip `json:"items,omitempty"`
}
