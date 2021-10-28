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

type ClientVPNRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClientVPNRouteSpec   `json:"spec,omitempty"`
	Status            ClientVPNRouteStatus `json:"status,omitempty"`
}

type ClientVPNRouteSpec struct {
	State *ClientVPNRouteSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClientVPNRouteSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClientVPNRouteSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClientVPNEndpointID *string `json:"clientVPNEndpointID" tf:"client_vpn_endpoint_id"`
	// +optional
	Description          *string `json:"description,omitempty" tf:"description"`
	DestinationCIDRBlock *string `json:"destinationCIDRBlock" tf:"destination_cidr_block"`
	// +optional
	Origin            *string `json:"origin,omitempty" tf:"origin"`
	TargetVpcSubnetID *string `json:"targetVpcSubnetID" tf:"target_vpc_subnet_id"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ClientVPNRouteStatus struct {
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

// ClientVPNRouteList is a list of ClientVPNRoutes
type ClientVPNRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClientVPNRoute CRD objects
	Items []ClientVPNRoute `json:"items,omitempty"`
}
