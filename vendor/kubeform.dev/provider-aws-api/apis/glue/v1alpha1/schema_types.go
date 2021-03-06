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

type Schema struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchemaSpec   `json:"spec,omitempty"`
	Status            SchemaStatus `json:"status,omitempty"`
}

type SchemaSpec struct {
	State *SchemaSpecResource `json:"state,omitempty" tf:"-"`

	Resource SchemaSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SchemaSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn           *string `json:"arn,omitempty" tf:"arn"`
	Compatibility *string `json:"compatibility" tf:"compatibility"`
	DataFormat    *string `json:"dataFormat" tf:"data_format"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	LatestSchemaVersion *int64 `json:"latestSchemaVersion,omitempty" tf:"latest_schema_version"`
	// +optional
	NextSchemaVersion *int64 `json:"nextSchemaVersion,omitempty" tf:"next_schema_version"`
	// +optional
	RegistryArn *string `json:"registryArn,omitempty" tf:"registry_arn"`
	// +optional
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name"`
	// +optional
	SchemaCheckpoint *int64  `json:"schemaCheckpoint,omitempty" tf:"schema_checkpoint"`
	SchemaDefinition *string `json:"schemaDefinition" tf:"schema_definition"`
	SchemaName       *string `json:"schemaName" tf:"schema_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type SchemaStatus struct {
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

// SchemaList is a list of Schemas
type SchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Schema CRD objects
	Items []Schema `json:"items,omitempty"`
}
