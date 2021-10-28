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

type ResolverFirewallRuleGroupAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverFirewallRuleGroupAssociationSpec   `json:"spec,omitempty"`
	Status            ResolverFirewallRuleGroupAssociationStatus `json:"status,omitempty"`
}

type ResolverFirewallRuleGroupAssociationSpec struct {
	State *ResolverFirewallRuleGroupAssociationSpecResource `json:"state,omitempty" tf:"-"`

	Resource ResolverFirewallRuleGroupAssociationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ResolverFirewallRuleGroupAssociationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                 *string `json:"arn,omitempty" tf:"arn"`
	FirewallRuleGroupID *string `json:"firewallRuleGroupID" tf:"firewall_rule_group_id"`
	// +optional
	MutationProtection *string `json:"mutationProtection,omitempty" tf:"mutation_protection"`
	Name               *string `json:"name" tf:"name"`
	Priority           *int64  `json:"priority" tf:"priority"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	VpcID   *string            `json:"vpcID" tf:"vpc_id"`
}

type ResolverFirewallRuleGroupAssociationStatus struct {
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

// ResolverFirewallRuleGroupAssociationList is a list of ResolverFirewallRuleGroupAssociations
type ResolverFirewallRuleGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ResolverFirewallRuleGroupAssociation CRD objects
	Items []ResolverFirewallRuleGroupAssociation `json:"items,omitempty"`
}
