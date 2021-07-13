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

type Policy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec,omitempty"`
	Status            PolicyStatus `json:"status,omitempty"`
}

type PolicySpecExcludeMap struct {
	// +optional
	Account []string `json:"account,omitempty" tf:"account"`
	// +optional
	Orgunit []string `json:"orgunit,omitempty" tf:"orgunit"`
}

type PolicySpecIncludeMap struct {
	// +optional
	Account []string `json:"account,omitempty" tf:"account"`
	// +optional
	Orgunit []string `json:"orgunit,omitempty" tf:"orgunit"`
}

type PolicySpecSecurityServicePolicyData struct {
	// +optional
	ManagedServiceData *string `json:"managedServiceData,omitempty" tf:"managed_service_data"`
	Type               *string `json:"type" tf:"type"`
}

type PolicySpec struct {
	State *PolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource PolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	DeleteAllPolicyResources *bool `json:"deleteAllPolicyResources,omitempty" tf:"delete_all_policy_resources"`
	// +optional
	ExcludeMap          *PolicySpecExcludeMap `json:"excludeMap,omitempty" tf:"exclude_map"`
	ExcludeResourceTags *bool                 `json:"excludeResourceTags" tf:"exclude_resource_tags"`
	// +optional
	IncludeMap *PolicySpecIncludeMap `json:"includeMap,omitempty" tf:"include_map"`
	Name       *string               `json:"name" tf:"name"`
	// +optional
	PolicyUpdateToken *string `json:"policyUpdateToken,omitempty" tf:"policy_update_token"`
	// +optional
	RemediationEnabled *bool `json:"remediationEnabled,omitempty" tf:"remediation_enabled"`
	// +optional
	ResourceTags *map[string]string `json:"resourceTags,omitempty" tf:"resource_tags"`
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// +optional
	ResourceTypeList          []string                             `json:"resourceTypeList,omitempty" tf:"resource_type_list"`
	SecurityServicePolicyData *PolicySpecSecurityServicePolicyData `json:"securityServicePolicyData" tf:"security_service_policy_data"`
}

type PolicyStatus struct {
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

// PolicyList is a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Policy CRD objects
	Items []Policy `json:"items,omitempty"`
}
