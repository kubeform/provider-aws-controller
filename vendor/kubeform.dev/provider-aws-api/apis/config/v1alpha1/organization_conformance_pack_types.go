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

type OrganizationConformancePack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationConformancePackSpec   `json:"spec,omitempty"`
	Status            OrganizationConformancePackStatus `json:"status,omitempty"`
}

type OrganizationConformancePackSpecInputParameter struct {
	ParameterName  *string `json:"parameterName" tf:"parameter_name"`
	ParameterValue *string `json:"parameterValue" tf:"parameter_value"`
}

type OrganizationConformancePackSpec struct {
	State *OrganizationConformancePackSpecResource `json:"state,omitempty" tf:"-"`

	Resource OrganizationConformancePackSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type OrganizationConformancePackSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	DeliveryS3Bucket *string `json:"deliveryS3Bucket,omitempty" tf:"delivery_s3_bucket"`
	// +optional
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix,omitempty" tf:"delivery_s3_key_prefix"`
	// +optional
	// +kubebuilder:validation:MaxItems=1000
	ExcludedAccounts []string `json:"excludedAccounts,omitempty" tf:"excluded_accounts"`
	// +optional
	// +kubebuilder:validation:MaxItems=60
	InputParameter []OrganizationConformancePackSpecInputParameter `json:"inputParameter,omitempty" tf:"input_parameter"`
	Name           *string                                         `json:"name" tf:"name"`
	// +optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body"`
	// +optional
	TemplateS3URI *string `json:"templateS3URI,omitempty" tf:"template_s3_uri"`
}

type OrganizationConformancePackStatus struct {
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

// OrganizationConformancePackList is a list of OrganizationConformancePacks
type OrganizationConformancePackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationConformancePack CRD objects
	Items []OrganizationConformancePack `json:"items,omitempty"`
}
