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

type VirtualService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualServiceSpec   `json:"spec,omitempty"`
	Status            VirtualServiceStatus `json:"status,omitempty"`
}

type VirtualServiceSpecSpecProviderVirtualNode struct {
	VirtualNodeName *string `json:"virtualNodeName" tf:"virtual_node_name"`
}

type VirtualServiceSpecSpecProviderVirtualRouter struct {
	VirtualRouterName *string `json:"virtualRouterName" tf:"virtual_router_name"`
}

type VirtualServiceSpecSpecProvider struct {
	// +optional
	VirtualNode *VirtualServiceSpecSpecProviderVirtualNode `json:"virtualNode,omitempty" tf:"virtual_node"`
	// +optional
	VirtualRouter *VirtualServiceSpecSpecProviderVirtualRouter `json:"virtualRouter,omitempty" tf:"virtual_router"`
}

type VirtualServiceSpecSpec struct {
	// +optional
	Provider *VirtualServiceSpecSpecProvider `json:"provider,omitempty" tf:"provider"`
}

type VirtualServiceSpec struct {
	State *VirtualServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource VirtualServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VirtualServiceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date"`
	// +optional
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date"`
	MeshName        *string `json:"meshName" tf:"mesh_name"`
	// +optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner"`
	Name      *string `json:"name" tf:"name"`
	// +optional
	ResourceOwner *string                 `json:"resourceOwner,omitempty" tf:"resource_owner"`
	Spec          *VirtualServiceSpecSpec `json:"spec" tf:"spec"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type VirtualServiceStatus struct {
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

// VirtualServiceList is a list of VirtualServices
type VirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualService CRD objects
	Items []VirtualService `json:"items,omitempty"`
}
