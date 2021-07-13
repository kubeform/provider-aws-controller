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

type OptionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OptionGroupSpec   `json:"spec,omitempty"`
	Status            OptionGroupStatus `json:"status,omitempty"`
}

type OptionGroupSpecOptionOptionSettings struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type OptionGroupSpecOption struct {
	// +optional
	DbSecurityGroupMemberships []string `json:"dbSecurityGroupMemberships,omitempty" tf:"db_security_group_memberships"`
	OptionName                 *string  `json:"optionName" tf:"option_name"`
	// +optional
	OptionSettings []OptionGroupSpecOptionOptionSettings `json:"optionSettings,omitempty" tf:"option_settings"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VpcSecurityGroupMemberships []string `json:"vpcSecurityGroupMemberships,omitempty" tf:"vpc_security_group_memberships"`
}

type OptionGroupSpec struct {
	State *OptionGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource OptionGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type OptionGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                *string `json:"arn,omitempty" tf:"arn"`
	EngineName         *string `json:"engineName" tf:"engine_name"`
	MajorEngineVersion *string `json:"majorEngineVersion" tf:"major_engine_version"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	Option []OptionGroupSpecOption `json:"option,omitempty" tf:"option"`
	// +optional
	OptionGroupDescription *string `json:"optionGroupDescription,omitempty" tf:"option_group_description"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type OptionGroupStatus struct {
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

// OptionGroupList is a list of OptionGroups
type OptionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OptionGroup CRD objects
	Items []OptionGroup `json:"items,omitempty"`
}
