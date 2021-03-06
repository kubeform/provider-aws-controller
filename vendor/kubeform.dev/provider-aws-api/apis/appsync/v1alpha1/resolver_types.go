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

type Resolver struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverSpec   `json:"spec,omitempty"`
	Status            ResolverStatus `json:"status,omitempty"`
}

type ResolverSpecCachingConfig struct {
	// +optional
	CachingKeys []string `json:"cachingKeys,omitempty" tf:"caching_keys"`
	// +optional
	Ttl *int64 `json:"ttl,omitempty" tf:"ttl"`
}

type ResolverSpecPipelineConfig struct {
	// +optional
	Functions []string `json:"functions,omitempty" tf:"functions"`
}

type ResolverSpecSyncConfigLambdaConflictHandlerConfig struct {
	// +optional
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn,omitempty" tf:"lambda_conflict_handler_arn"`
}

type ResolverSpecSyncConfig struct {
	// +optional
	ConflictDetection *string `json:"conflictDetection,omitempty" tf:"conflict_detection"`
	// +optional
	ConflictHandler *string `json:"conflictHandler,omitempty" tf:"conflict_handler"`
	// +optional
	LambdaConflictHandlerConfig *ResolverSpecSyncConfigLambdaConflictHandlerConfig `json:"lambdaConflictHandlerConfig,omitempty" tf:"lambda_conflict_handler_config"`
}

type ResolverSpec struct {
	State *ResolverSpecResource `json:"state,omitempty" tf:"-"`

	Resource ResolverSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ResolverSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiID *string `json:"apiID" tf:"api_id"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CachingConfig *ResolverSpecCachingConfig `json:"cachingConfig,omitempty" tf:"caching_config"`
	// +optional
	DataSource *string `json:"dataSource,omitempty" tf:"data_source"`
	Field      *string `json:"field" tf:"field"`
	// +optional
	Kind *string `json:"kind,omitempty" tf:"kind"`
	// +optional
	MaxBatchSize *int64 `json:"maxBatchSize,omitempty" tf:"max_batch_size"`
	// +optional
	PipelineConfig *ResolverSpecPipelineConfig `json:"pipelineConfig,omitempty" tf:"pipeline_config"`
	// +optional
	RequestTemplate *string `json:"requestTemplate,omitempty" tf:"request_template"`
	// +optional
	ResponseTemplate *string `json:"responseTemplate,omitempty" tf:"response_template"`
	// +optional
	SyncConfig *ResolverSpecSyncConfig `json:"syncConfig,omitempty" tf:"sync_config"`
	Type       *string                 `json:"type" tf:"type"`
}

type ResolverStatus struct {
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

// ResolverList is a list of Resolvers
type ResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Resolver CRD objects
	Items []Resolver `json:"items,omitempty"`
}
