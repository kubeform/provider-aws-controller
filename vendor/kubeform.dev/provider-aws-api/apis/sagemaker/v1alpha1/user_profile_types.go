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

type UserProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserProfileSpec   `json:"spec,omitempty"`
	Status            UserProfileStatus `json:"status,omitempty"`
}

type UserProfileSpecUserSettingsJupyterServerAppSettingsDefaultResourceSpec struct {
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn"`
	// +optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn"`
	// +optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn"`
}

type UserProfileSpecUserSettingsJupyterServerAppSettings struct {
	DefaultResourceSpec *UserProfileSpecUserSettingsJupyterServerAppSettingsDefaultResourceSpec `json:"defaultResourceSpec" tf:"default_resource_spec"`
	// +optional
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns"`
}

type UserProfileSpecUserSettingsKernelGatewayAppSettingsCustomImage struct {
	AppImageConfigName *string `json:"appImageConfigName" tf:"app_image_config_name"`
	ImageName          *string `json:"imageName" tf:"image_name"`
	// +optional
	ImageVersionNumber *int64 `json:"imageVersionNumber,omitempty" tf:"image_version_number"`
}

type UserProfileSpecUserSettingsKernelGatewayAppSettingsDefaultResourceSpec struct {
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn"`
	// +optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn"`
	// +optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn"`
}

type UserProfileSpecUserSettingsKernelGatewayAppSettings struct {
	// +optional
	// +kubebuilder:validation:MaxItems=30
	CustomImage         []UserProfileSpecUserSettingsKernelGatewayAppSettingsCustomImage        `json:"customImage,omitempty" tf:"custom_image"`
	DefaultResourceSpec *UserProfileSpecUserSettingsKernelGatewayAppSettingsDefaultResourceSpec `json:"defaultResourceSpec" tf:"default_resource_spec"`
	// +optional
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns"`
}

type UserProfileSpecUserSettingsSharingSettings struct {
	// +optional
	NotebookOutputOption *string `json:"notebookOutputOption,omitempty" tf:"notebook_output_option"`
	// +optional
	S3KmsKeyID *string `json:"s3KmsKeyID,omitempty" tf:"s3_kms_key_id"`
	// +optional
	S3OutputPath *string `json:"s3OutputPath,omitempty" tf:"s3_output_path"`
}

type UserProfileSpecUserSettingsTensorBoardAppSettingsDefaultResourceSpec struct {
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn"`
	// +optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn"`
	// +optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn"`
}

type UserProfileSpecUserSettingsTensorBoardAppSettings struct {
	DefaultResourceSpec *UserProfileSpecUserSettingsTensorBoardAppSettingsDefaultResourceSpec `json:"defaultResourceSpec" tf:"default_resource_spec"`
}

type UserProfileSpecUserSettings struct {
	ExecutionRole *string `json:"executionRole" tf:"execution_role"`
	// +optional
	JupyterServerAppSettings *UserProfileSpecUserSettingsJupyterServerAppSettings `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings"`
	// +optional
	KernelGatewayAppSettings *UserProfileSpecUserSettingsKernelGatewayAppSettings `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	// +optional
	SharingSettings *UserProfileSpecUserSettingsSharingSettings `json:"sharingSettings,omitempty" tf:"sharing_settings"`
	// +optional
	TensorBoardAppSettings *UserProfileSpecUserSettingsTensorBoardAppSettings `json:"tensorBoardAppSettings,omitempty" tf:"tensor_board_app_settings"`
}

type UserProfileSpec struct {
	State *UserProfileSpecResource `json:"state,omitempty" tf:"-"`

	Resource UserProfileSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type UserProfileSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn      *string `json:"arn,omitempty" tf:"arn"`
	DomainID *string `json:"domainID" tf:"domain_id"`
	// +optional
	HomeEfsFileSystemUid *string `json:"homeEfsFileSystemUid,omitempty" tf:"home_efs_file_system_uid"`
	// +optional
	SingleSignOnUserIdentifier *string `json:"singleSignOnUserIdentifier,omitempty" tf:"single_sign_on_user_identifier"`
	// +optional
	SingleSignOnUserValue *string `json:"singleSignOnUserValue,omitempty" tf:"single_sign_on_user_value"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll         *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	UserProfileName *string            `json:"userProfileName" tf:"user_profile_name"`
	// +optional
	UserSettings *UserProfileSpecUserSettings `json:"userSettings,omitempty" tf:"user_settings"`
}

type UserProfileStatus struct {
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

// UserProfileList is a list of UserProfiles
type UserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of UserProfile CRD objects
	Items []UserProfile `json:"items,omitempty"`
}
