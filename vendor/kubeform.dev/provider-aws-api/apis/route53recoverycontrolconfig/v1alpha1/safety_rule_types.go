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

type SafetyRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SafetyRuleSpec   `json:"spec,omitempty"`
	Status            SafetyRuleStatus `json:"status,omitempty"`
}

type SafetyRuleSpecRuleConfig struct {
	Inverted  *bool   `json:"inverted" tf:"inverted"`
	Threshold *int64  `json:"threshold" tf:"threshold"`
	Type      *string `json:"type" tf:"type"`
}

type SafetyRuleSpec struct {
	State *SafetyRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource SafetyRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SafetyRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AssertedControls []string `json:"assertedControls,omitempty" tf:"asserted_controls"`
	ControlPanelArn  *string  `json:"controlPanelArn" tf:"control_panel_arn"`
	// +optional
	GatingControls []string                  `json:"gatingControls,omitempty" tf:"gating_controls"`
	Name           *string                   `json:"name" tf:"name"`
	RuleConfig     *SafetyRuleSpecRuleConfig `json:"ruleConfig" tf:"rule_config"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TargetControls []string `json:"targetControls,omitempty" tf:"target_controls"`
	WaitPeriodMs   *int64   `json:"waitPeriodMs" tf:"wait_period_ms"`
}

type SafetyRuleStatus struct {
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

// SafetyRuleList is a list of SafetyRules
type SafetyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SafetyRule CRD objects
	Items []SafetyRule `json:"items,omitempty"`
}
