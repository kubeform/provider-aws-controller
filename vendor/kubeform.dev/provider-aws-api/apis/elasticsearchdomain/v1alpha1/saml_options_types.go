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

type SamlOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SamlOptionsSpec   `json:"spec,omitempty"`
	Status            SamlOptionsStatus `json:"status,omitempty"`
}

type SamlOptionsSpecSamlOptionsIdp struct {
	EntityID        *string `json:"entityID" tf:"entity_id"`
	MetadataContent *string `json:"metadataContent" tf:"metadata_content"`
}

type SamlOptionsSpecSamlOptions struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	Idp *SamlOptionsSpecSamlOptionsIdp `json:"idp,omitempty" tf:"idp"`
	// +optional
	MasterBackendRole *string `json:"masterBackendRole,omitempty" tf:"master_backend_role"`
	// +optional
	MasterUserName *string `json:"-" sensitive:"true" tf:"master_user_name"`
	// +optional
	RolesKey *string `json:"rolesKey,omitempty" tf:"roles_key"`
	// +optional
	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty" tf:"session_timeout_minutes"`
	// +optional
	SubjectKey *string `json:"subjectKey,omitempty" tf:"subject_key"`
}

type SamlOptionsSpec struct {
	State *SamlOptionsSpecResource `json:"state,omitempty" tf:"-"`

	Resource SamlOptionsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SamlOptionsSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DomainName *string `json:"domainName" tf:"domain_name"`
	// +optional
	SamlOptions *SamlOptionsSpecSamlOptions `json:"samlOptions,omitempty" tf:"saml_options"`
}

type SamlOptionsStatus struct {
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

// SamlOptionsList is a list of SamlOptionss
type SamlOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SamlOptions CRD objects
	Items []SamlOptions `json:"items,omitempty"`
}
