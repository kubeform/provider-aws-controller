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

type Permission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionSpec   `json:"spec,omitempty"`
	Status            PermissionStatus `json:"status,omitempty"`
}

type PermissionSpec struct {
	State *PermissionSpecResource `json:"state,omitempty" tf:"-"`

	Resource PermissionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PermissionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action *string `json:"action" tf:"action"`
	// +optional
	EventSourceToken *string `json:"eventSourceToken,omitempty" tf:"event_source_token"`
	FunctionName     *string `json:"functionName" tf:"function_name"`
	Principal        *string `json:"principal" tf:"principal"`
	// +optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier"`
	// +optional
	SourceAccount *string `json:"sourceAccount,omitempty" tf:"source_account"`
	// +optional
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn"`
	// +optional
	StatementID *string `json:"statementID,omitempty" tf:"statement_id"`
	// +optional
	StatementIDPrefix *string `json:"statementIDPrefix,omitempty" tf:"statement_id_prefix"`
}

type PermissionStatus struct {
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

// PermissionList is a list of Permissions
type PermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Permission CRD objects
	Items []Permission `json:"items,omitempty"`
}
