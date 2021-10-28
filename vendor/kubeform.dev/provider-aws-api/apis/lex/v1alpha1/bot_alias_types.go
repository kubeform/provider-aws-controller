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

type BotAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotAliasSpec   `json:"spec,omitempty"`
	Status            BotAliasStatus `json:"status,omitempty"`
}

type BotAliasSpecConversationLogsLogSettings struct {
	Destination *string `json:"destination" tf:"destination"`
	// +optional
	KmsKeyArn   *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
	LogType     *string `json:"logType" tf:"log_type"`
	ResourceArn *string `json:"resourceArn" tf:"resource_arn"`
	// +optional
	ResourcePrefix *string `json:"resourcePrefix,omitempty" tf:"resource_prefix"`
}

type BotAliasSpecConversationLogs struct {
	IamRoleArn *string `json:"iamRoleArn" tf:"iam_role_arn"`
	// +optional
	LogSettings []BotAliasSpecConversationLogsLogSettings `json:"logSettings,omitempty" tf:"log_settings"`
}

type BotAliasSpec struct {
	State *BotAliasSpecResource `json:"state,omitempty" tf:"-"`

	Resource BotAliasSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BotAliasSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn        *string `json:"arn,omitempty" tf:"arn"`
	BotName    *string `json:"botName" tf:"bot_name"`
	BotVersion *string `json:"botVersion" tf:"bot_version"`
	// +optional
	Checksum *string `json:"checksum,omitempty" tf:"checksum"`
	// +optional
	ConversationLogs *BotAliasSpecConversationLogs `json:"conversationLogs,omitempty" tf:"conversation_logs"`
	// +optional
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date"`
	Name            *string `json:"name" tf:"name"`
}

type BotAliasStatus struct {
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

// BotAliasList is a list of BotAliass
type BotAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BotAlias CRD objects
	Items []BotAlias `json:"items,omitempty"`
}
