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

type GlacierVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlacierVaultSpec   `json:"spec,omitempty"`
	Status            GlacierVaultStatus `json:"status,omitempty"`
}

type GlacierVaultSpecNotification struct {
	Events   []string `json:"events" tf:"events"`
	SnsTopic *string  `json:"snsTopic" tf:"sns_topic"`
}

type GlacierVaultSpec struct {
	State *GlacierVaultSpecResource `json:"state,omitempty" tf:"-"`

	Resource GlacierVaultSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GlacierVaultSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	Notification *GlacierVaultSpecNotification `json:"notification,omitempty" tf:"notification"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type GlacierVaultStatus struct {
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

// GlacierVaultList is a list of GlacierVaults
type GlacierVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlacierVault CRD objects
	Items []GlacierVault `json:"items,omitempty"`
}
