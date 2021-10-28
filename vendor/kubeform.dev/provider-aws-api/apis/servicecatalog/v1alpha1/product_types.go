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

type Product struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductSpec   `json:"spec,omitempty"`
	Status            ProductStatus `json:"status,omitempty"`
}

type ProductSpecProvisioningArtifactParameters struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	TemplatePhysicalID *string `json:"templatePhysicalID,omitempty" tf:"template_physical_id"`
	// +optional
	TemplateURL *string `json:"templateURL,omitempty" tf:"template_url"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ProductSpec struct {
	State *ProductSpecResource `json:"state,omitempty" tf:"-"`

	Resource ProductSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ProductSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Distributor *string `json:"distributor,omitempty" tf:"distributor"`
	// +optional
	HasDefaultPath                 *bool                                      `json:"hasDefaultPath,omitempty" tf:"has_default_path"`
	Name                           *string                                    `json:"name" tf:"name"`
	Owner                          *string                                    `json:"owner" tf:"owner"`
	ProvisioningArtifactParameters *ProductSpecProvisioningArtifactParameters `json:"provisioningArtifactParameters" tf:"provisioning_artifact_parameters"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	SupportDescription *string `json:"supportDescription,omitempty" tf:"support_description"`
	// +optional
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email"`
	// +optional
	SupportURL *string `json:"supportURL,omitempty" tf:"support_url"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	Type    *string            `json:"type" tf:"type"`
}

type ProductStatus struct {
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

// ProductList is a list of Products
type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Product CRD objects
	Items []Product `json:"items,omitempty"`
}
