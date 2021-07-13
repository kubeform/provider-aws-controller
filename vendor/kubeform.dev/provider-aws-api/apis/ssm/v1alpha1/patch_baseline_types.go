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

type PatchBaseline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PatchBaselineSpec   `json:"spec,omitempty"`
	Status            PatchBaselineStatus `json:"status,omitempty"`
}

type PatchBaselineSpecApprovalRulePatchFilter struct {
	Key *string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=20
	// +kubebuilder:validation:MinItems=1
	Values []string `json:"values" tf:"values"`
}

type PatchBaselineSpecApprovalRule struct {
	// +optional
	ApproveAfterDays *int64 `json:"approveAfterDays,omitempty" tf:"approve_after_days"`
	// +optional
	ApproveUntilDate *string `json:"approveUntilDate,omitempty" tf:"approve_until_date"`
	// +optional
	ComplianceLevel *string `json:"complianceLevel,omitempty" tf:"compliance_level"`
	// +optional
	EnableNonSecurity *bool `json:"enableNonSecurity,omitempty" tf:"enable_non_security"`
	// +kubebuilder:validation:MaxItems=10
	PatchFilter []PatchBaselineSpecApprovalRulePatchFilter `json:"patchFilter" tf:"patch_filter"`
}

type PatchBaselineSpecGlobalFilter struct {
	Key *string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=20
	// +kubebuilder:validation:MinItems=1
	Values []string `json:"values" tf:"values"`
}

type PatchBaselineSpecSource struct {
	Configuration *string `json:"configuration" tf:"configuration"`
	Name          *string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=20
	Products []string `json:"products" tf:"products"`
}

type PatchBaselineSpec struct {
	State *PatchBaselineSpecResource `json:"state,omitempty" tf:"-"`

	Resource PatchBaselineSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PatchBaselineSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApprovalRule []PatchBaselineSpecApprovalRule `json:"approvalRule,omitempty" tf:"approval_rule"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	ApprovedPatches []string `json:"approvedPatches,omitempty" tf:"approved_patches"`
	// +optional
	ApprovedPatchesComplianceLevel *string `json:"approvedPatchesComplianceLevel,omitempty" tf:"approved_patches_compliance_level"`
	// +optional
	ApprovedPatchesEnableNonSecurity *bool `json:"approvedPatchesEnableNonSecurity,omitempty" tf:"approved_patches_enable_non_security"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	GlobalFilter []PatchBaselineSpecGlobalFilter `json:"globalFilter,omitempty" tf:"global_filter"`
	Name         *string                         `json:"name" tf:"name"`
	// +optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	RejectedPatches []string `json:"rejectedPatches,omitempty" tf:"rejected_patches"`
	// +optional
	RejectedPatchesAction *string `json:"rejectedPatchesAction,omitempty" tf:"rejected_patches_action"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	Source []PatchBaselineSpecSource `json:"source,omitempty" tf:"source"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type PatchBaselineStatus struct {
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

// PatchBaselineList is a list of PatchBaselines
type PatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PatchBaseline CRD objects
	Items []PatchBaseline `json:"items,omitempty"`
}
