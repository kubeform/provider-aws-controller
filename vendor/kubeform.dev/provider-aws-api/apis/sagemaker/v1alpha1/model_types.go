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

type Model struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelSpec   `json:"spec,omitempty"`
	Status            ModelStatus `json:"status,omitempty"`
}

type ModelSpecContainerImageConfig struct {
	RepositoryAccessMode *string `json:"repositoryAccessMode" tf:"repository_access_mode"`
}

type ModelSpecContainer struct {
	// +optional
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname"`
	// +optional
	Environment *map[string]string `json:"environment,omitempty" tf:"environment"`
	Image       *string            `json:"image" tf:"image"`
	// +optional
	ImageConfig *ModelSpecContainerImageConfig `json:"imageConfig,omitempty" tf:"image_config"`
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// +optional
	ModelDataURL *string `json:"modelDataURL,omitempty" tf:"model_data_url"`
}

type ModelSpecInferenceExecutionConfig struct {
	Mode *string `json:"mode" tf:"mode"`
}

type ModelSpecPrimaryContainerImageConfig struct {
	RepositoryAccessMode *string `json:"repositoryAccessMode" tf:"repository_access_mode"`
}

type ModelSpecPrimaryContainer struct {
	// +optional
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname"`
	// +optional
	Environment *map[string]string `json:"environment,omitempty" tf:"environment"`
	Image       *string            `json:"image" tf:"image"`
	// +optional
	ImageConfig *ModelSpecPrimaryContainerImageConfig `json:"imageConfig,omitempty" tf:"image_config"`
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// +optional
	ModelDataURL *string `json:"modelDataURL,omitempty" tf:"model_data_url"`
}

type ModelSpecVpcConfig struct {
	// +kubebuilder:validation:MaxItems=5
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +kubebuilder:validation:MaxItems=16
	Subnets []string `json:"subnets" tf:"subnets"`
}

type ModelSpec struct {
	State *ModelSpecResource `json:"state,omitempty" tf:"-"`

	Resource ModelSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ModelSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Container []ModelSpecContainer `json:"container,omitempty" tf:"container"`
	// +optional
	EnableNetworkIsolation *bool   `json:"enableNetworkIsolation,omitempty" tf:"enable_network_isolation"`
	ExecutionRoleArn       *string `json:"executionRoleArn" tf:"execution_role_arn"`
	// +optional
	InferenceExecutionConfig *ModelSpecInferenceExecutionConfig `json:"inferenceExecutionConfig,omitempty" tf:"inference_execution_config"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PrimaryContainer *ModelSpecPrimaryContainer `json:"primaryContainer,omitempty" tf:"primary_container"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	VpcConfig *ModelSpecVpcConfig `json:"vpcConfig,omitempty" tf:"vpc_config"`
}

type ModelStatus struct {
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

// ModelList is a list of Models
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Model CRD objects
	Items []Model `json:"items,omitempty"`
}
