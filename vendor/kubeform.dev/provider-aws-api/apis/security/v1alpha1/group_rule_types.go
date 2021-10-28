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

type GroupRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupRuleSpec   `json:"spec,omitempty"`
	Status            GroupRuleStatus `json:"status,omitempty"`
}

type GroupRuleSpec struct {
	State *GroupRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource GroupRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GroupRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	FromPort    *int64  `json:"fromPort" tf:"from_port"`
	// +optional
	Ipv6CIDRBlocks []string `json:"ipv6CIDRBlocks,omitempty" tf:"ipv6_cidr_blocks"`
	// +optional
	PrefixListIDS   []string `json:"prefixListIDS,omitempty" tf:"prefix_list_ids"`
	Protocol        *string  `json:"protocol" tf:"protocol"`
	SecurityGroupID *string  `json:"securityGroupID" tf:"security_group_id"`
	// +optional
	Self *bool `json:"self,omitempty" tf:"self"`
	// +optional
	SourceSecurityGroupID *string `json:"sourceSecurityGroupID,omitempty" tf:"source_security_group_id"`
	ToPort                *int64  `json:"toPort" tf:"to_port"`
	// Type of rule, ingress (inbound) or egress (outbound).
	Type *string `json:"type" tf:"type"`
}

type GroupRuleStatus struct {
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

// GroupRuleList is a list of GroupRules
type GroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GroupRule CRD objects
	Items []GroupRule `json:"items,omitempty"`
}
