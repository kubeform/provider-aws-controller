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

type Secret struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretSpec   `json:"spec,omitempty"`
	Status            SecretStatus `json:"status,omitempty"`
}

type SecretSpecReplica struct {
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	LastAccessedDate *string `json:"lastAccessedDate,omitempty" tf:"last_accessed_date"`
	Region           *string `json:"region" tf:"region"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StatusMessage *string `json:"statusMessage,omitempty" tf:"status_message"`
}

type SecretSpecRotationRules struct {
	AutomaticallyAfterDays *int64 `json:"automaticallyAfterDays" tf:"automatically_after_days"`
}

type SecretSpec struct {
	State *SecretSpecResource `json:"state,omitempty" tf:"-"`

	Resource SecretSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SecretSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ForceOverwriteReplicaSecret *bool `json:"forceOverwriteReplicaSecret,omitempty" tf:"force_overwrite_replica_secret"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	Policy *string `json:"policy,omitempty" tf:"policy"`
	// +optional
	RecoveryWindowInDays *int64 `json:"recoveryWindowInDays,omitempty" tf:"recovery_window_in_days"`
	// +optional
	Replica []SecretSpecReplica `json:"replica,omitempty" tf:"replica"`
	// +optional
	// Deprecated
	RotationEnabled *bool `json:"rotationEnabled,omitempty" tf:"rotation_enabled"`
	// +optional
	// Deprecated
	RotationLambdaArn *string `json:"rotationLambdaArn,omitempty" tf:"rotation_lambda_arn"`
	// +optional
	// Deprecated
	RotationRules *SecretSpecRotationRules `json:"rotationRules,omitempty" tf:"rotation_rules"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type SecretStatus struct {
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

// SecretList is a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Secret CRD objects
	Items []Secret `json:"items,omitempty"`
}
