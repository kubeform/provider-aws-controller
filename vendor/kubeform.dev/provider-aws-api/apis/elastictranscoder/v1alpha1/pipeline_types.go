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

type Pipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PipelineSpec   `json:"spec,omitempty"`
	Status            PipelineStatus `json:"status,omitempty"`
}

type PipelineSpecContentConfig struct {
	// +optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket"`
	// +optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`
}

type PipelineSpecContentConfigPermissions struct {
	// +optional
	Access []string `json:"access,omitempty" tf:"access"`
	// +optional
	Grantee *string `json:"grantee,omitempty" tf:"grantee"`
	// +optional
	GranteeType *string `json:"granteeType,omitempty" tf:"grantee_type"`
}

type PipelineSpecNotifications struct {
	// +optional
	Completed *string `json:"completed,omitempty" tf:"completed"`
	// +optional
	Error *string `json:"error,omitempty" tf:"error"`
	// +optional
	Progressing *string `json:"progressing,omitempty" tf:"progressing"`
	// +optional
	Warning *string `json:"warning,omitempty" tf:"warning"`
}

type PipelineSpecThumbnailConfig struct {
	// +optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket"`
	// +optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`
}

type PipelineSpecThumbnailConfigPermissions struct {
	// +optional
	Access []string `json:"access,omitempty" tf:"access"`
	// +optional
	Grantee *string `json:"grantee,omitempty" tf:"grantee"`
	// +optional
	GranteeType *string `json:"granteeType,omitempty" tf:"grantee_type"`
}

type PipelineSpec struct {
	State *PipelineSpecResource `json:"state,omitempty" tf:"-"`

	Resource PipelineSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PipelineSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AwsKmsKeyArn *string `json:"awsKmsKeyArn,omitempty" tf:"aws_kms_key_arn"`
	// +optional
	ContentConfig *PipelineSpecContentConfig `json:"contentConfig,omitempty" tf:"content_config"`
	// +optional
	ContentConfigPermissions []PipelineSpecContentConfigPermissions `json:"contentConfigPermissions,omitempty" tf:"content_config_permissions"`
	InputBucket              *string                                `json:"inputBucket" tf:"input_bucket"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Notifications *PipelineSpecNotifications `json:"notifications,omitempty" tf:"notifications"`
	// +optional
	OutputBucket *string `json:"outputBucket,omitempty" tf:"output_bucket"`
	Role         *string `json:"role" tf:"role"`
	// +optional
	ThumbnailConfig *PipelineSpecThumbnailConfig `json:"thumbnailConfig,omitempty" tf:"thumbnail_config"`
	// +optional
	ThumbnailConfigPermissions []PipelineSpecThumbnailConfigPermissions `json:"thumbnailConfigPermissions,omitempty" tf:"thumbnail_config_permissions"`
}

type PipelineStatus struct {
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

// PipelineList is a list of Pipelines
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Pipeline CRD objects
	Items []Pipeline `json:"items,omitempty"`
}
