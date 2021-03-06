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

type Function struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec,omitempty"`
	Status            FunctionStatus `json:"status,omitempty"`
}

type FunctionSpecDeadLetterConfig struct {
	TargetArn *string `json:"targetArn" tf:"target_arn"`
}

type FunctionSpecEnvironment struct {
	// +optional
	Variables *map[string]string `json:"variables,omitempty" tf:"variables"`
}

type FunctionSpecFileSystemConfig struct {
	Arn            *string `json:"arn" tf:"arn"`
	LocalMountPath *string `json:"localMountPath" tf:"local_mount_path"`
}

type FunctionSpecImageConfig struct {
	// +optional
	Command []string `json:"command,omitempty" tf:"command"`
	// +optional
	EntryPoint []string `json:"entryPoint,omitempty" tf:"entry_point"`
	// +optional
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory"`
}

type FunctionSpecTracingConfig struct {
	Mode *string `json:"mode" tf:"mode"`
}

type FunctionSpecVpcConfig struct {
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	SubnetIDS        []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type FunctionSpec struct {
	State *FunctionSpecResource `json:"state,omitempty" tf:"-"`

	Resource FunctionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FunctionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Architectures []string `json:"architectures,omitempty" tf:"architectures"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CodeSigningConfigArn *string `json:"codeSigningConfigArn,omitempty" tf:"code_signing_config_arn"`
	// +optional
	DeadLetterConfig *FunctionSpecDeadLetterConfig `json:"deadLetterConfig,omitempty" tf:"dead_letter_config"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Environment *FunctionSpecEnvironment `json:"environment,omitempty" tf:"environment"`
	// +optional
	FileSystemConfig *FunctionSpecFileSystemConfig `json:"fileSystemConfig,omitempty" tf:"file_system_config"`
	// +optional
	Filename     *string `json:"filename,omitempty" tf:"filename"`
	FunctionName *string `json:"functionName" tf:"function_name"`
	// +optional
	Handler *string `json:"handler,omitempty" tf:"handler"`
	// +optional
	ImageConfig *FunctionSpecImageConfig `json:"imageConfig,omitempty" tf:"image_config"`
	// +optional
	ImageURI *string `json:"imageURI,omitempty" tf:"image_uri"`
	// +optional
	InvokeArn *string `json:"invokeArn,omitempty" tf:"invoke_arn"`
	// +optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
	// +optional
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	Layers []string `json:"layers,omitempty" tf:"layers"`
	// +optional
	MemorySize *int64 `json:"memorySize,omitempty" tf:"memory_size"`
	// +optional
	PackageType *string `json:"packageType,omitempty" tf:"package_type"`
	// +optional
	Publish *bool `json:"publish,omitempty" tf:"publish"`
	// +optional
	QualifiedArn *string `json:"qualifiedArn,omitempty" tf:"qualified_arn"`
	// +optional
	ReservedConcurrentExecutions *int64  `json:"reservedConcurrentExecutions,omitempty" tf:"reserved_concurrent_executions"`
	Role                         *string `json:"role" tf:"role"`
	// +optional
	Runtime *string `json:"runtime,omitempty" tf:"runtime"`
	// +optional
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket"`
	// +optional
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key"`
	// +optional
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version"`
	// +optional
	SigningJobArn *string `json:"signingJobArn,omitempty" tf:"signing_job_arn"`
	// +optional
	SigningProfileVersionArn *string `json:"signingProfileVersionArn,omitempty" tf:"signing_profile_version_arn"`
	// +optional
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash"`
	// +optional
	SourceCodeSize *int64 `json:"sourceCodeSize,omitempty" tf:"source_code_size"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// +optional
	TracingConfig *FunctionSpecTracingConfig `json:"tracingConfig,omitempty" tf:"tracing_config"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VpcConfig *FunctionSpecVpcConfig `json:"vpcConfig,omitempty" tf:"vpc_config"`
}

type FunctionStatus struct {
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

// FunctionList is a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Function CRD objects
	Items []Function `json:"items,omitempty"`
}
