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

type ReplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationConfigurationSpec   `json:"spec,omitempty"`
	Status            ReplicationConfigurationStatus `json:"status,omitempty"`
}

type ReplicationConfigurationSpecReplicationConfigurationRuleDestination struct {
	Region     *string `json:"region" tf:"region"`
	RegistryID *string `json:"registryID" tf:"registry_id"`
}

type ReplicationConfigurationSpecReplicationConfigurationRuleRepositoryFilter struct {
	Filter     *string `json:"filter" tf:"filter"`
	FilterType *string `json:"filterType" tf:"filter_type"`
}

type ReplicationConfigurationSpecReplicationConfigurationRule struct {
	// +kubebuilder:validation:MaxItems=25
	Destination []ReplicationConfigurationSpecReplicationConfigurationRuleDestination `json:"destination" tf:"destination"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:MinItems=1
	RepositoryFilter []ReplicationConfigurationSpecReplicationConfigurationRuleRepositoryFilter `json:"repositoryFilter,omitempty" tf:"repository_filter"`
}

type ReplicationConfigurationSpecReplicationConfiguration struct {
	// +kubebuilder:validation:MaxItems=10
	Rule []ReplicationConfigurationSpecReplicationConfigurationRule `json:"rule" tf:"rule"`
}

type ReplicationConfigurationSpec struct {
	State *ReplicationConfigurationSpecResource `json:"state,omitempty" tf:"-"`

	Resource ReplicationConfigurationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ReplicationConfigurationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	RegistryID *string `json:"registryID,omitempty" tf:"registry_id"`
	// +optional
	ReplicationConfiguration *ReplicationConfigurationSpecReplicationConfiguration `json:"replicationConfiguration,omitempty" tf:"replication_configuration"`
}

type ReplicationConfigurationStatus struct {
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

// ReplicationConfigurationList is a list of ReplicationConfigurations
type ReplicationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ReplicationConfiguration CRD objects
	Items []ReplicationConfiguration `json:"items,omitempty"`
}
