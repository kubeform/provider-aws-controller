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

type ApnsSandboxChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApnsSandboxChannelSpec   `json:"spec,omitempty"`
	Status            ApnsSandboxChannelStatus `json:"status,omitempty"`
}

type ApnsSandboxChannelSpec struct {
	State *ApnsSandboxChannelSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApnsSandboxChannelSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ApnsSandboxChannelSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApplicationID *string `json:"applicationID" tf:"application_id"`
	// +optional
	BundleID *string `json:"-" sensitive:"true" tf:"bundle_id"`
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod,omitempty" tf:"default_authentication_method"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	PrivateKey *string `json:"-" sensitive:"true" tf:"private_key"`
	// +optional
	TeamID *string `json:"-" sensitive:"true" tf:"team_id"`
	// +optional
	TokenKey *string `json:"-" sensitive:"true" tf:"token_key"`
	// +optional
	TokenKeyID *string `json:"-" sensitive:"true" tf:"token_key_id"`
}

type ApnsSandboxChannelStatus struct {
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

// ApnsSandboxChannelList is a list of ApnsSandboxChannels
type ApnsSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApnsSandboxChannel CRD objects
	Items []ApnsSandboxChannel `json:"items,omitempty"`
}
