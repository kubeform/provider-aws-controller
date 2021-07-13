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

type Webhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebhookSpec   `json:"spec,omitempty"`
	Status            WebhookStatus `json:"status,omitempty"`
}

type WebhookSpecAuthenticationConfiguration struct {
	// +optional
	AllowedIPRange *string `json:"allowedIPRange,omitempty" tf:"allowed_ip_range"`
	// +optional
	SecretToken *string `json:"-" sensitive:"true" tf:"secret_token"`
}

type WebhookSpecFilter struct {
	JsonPath    *string `json:"jsonPath" tf:"json_path"`
	MatchEquals *string `json:"matchEquals" tf:"match_equals"`
}

type WebhookSpec struct {
	State *WebhookSpecResource `json:"state,omitempty" tf:"-"`

	Resource WebhookSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type WebhookSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Authentication *string `json:"authentication" tf:"authentication"`
	// +optional
	AuthenticationConfiguration *WebhookSpecAuthenticationConfiguration `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration"`
	// +kubebuilder:validation:MinItems=1
	Filter []WebhookSpecFilter `json:"filter" tf:"filter"`
	Name   *string             `json:"name" tf:"name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll        *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	TargetAction   *string            `json:"targetAction" tf:"target_action"`
	TargetPipeline *string            `json:"targetPipeline" tf:"target_pipeline"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type WebhookStatus struct {
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

// WebhookList is a list of Webhooks
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Webhook CRD objects
	Items []Webhook `json:"items,omitempty"`
}
