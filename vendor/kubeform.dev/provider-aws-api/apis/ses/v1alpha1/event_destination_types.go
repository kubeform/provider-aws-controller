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

type EventDestination struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventDestinationSpec   `json:"spec,omitempty"`
	Status            EventDestinationStatus `json:"status,omitempty"`
}

type EventDestinationSpecCloudwatchDestination struct {
	DefaultValue  *string `json:"defaultValue" tf:"default_value"`
	DimensionName *string `json:"dimensionName" tf:"dimension_name"`
	ValueSource   *string `json:"valueSource" tf:"value_source"`
}

type EventDestinationSpecKinesisDestination struct {
	RoleArn   *string `json:"roleArn" tf:"role_arn"`
	StreamArn *string `json:"streamArn" tf:"stream_arn"`
}

type EventDestinationSpecSnsDestination struct {
	TopicArn *string `json:"topicArn" tf:"topic_arn"`
}

type EventDestinationSpec struct {
	State *EventDestinationSpecResource `json:"state,omitempty" tf:"-"`

	Resource EventDestinationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EventDestinationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CloudwatchDestination []EventDestinationSpecCloudwatchDestination `json:"cloudwatchDestination,omitempty" tf:"cloudwatch_destination"`
	ConfigurationSetName  *string                                     `json:"configurationSetName" tf:"configuration_set_name"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	KinesisDestination *EventDestinationSpecKinesisDestination `json:"kinesisDestination,omitempty" tf:"kinesis_destination"`
	MatchingTypes      []string                                `json:"matchingTypes" tf:"matching_types"`
	Name               *string                                 `json:"name" tf:"name"`
	// +optional
	SnsDestination *EventDestinationSpecSnsDestination `json:"snsDestination,omitempty" tf:"sns_destination"`
}

type EventDestinationStatus struct {
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

// EventDestinationList is a list of EventDestinations
type EventDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventDestination CRD objects
	Items []EventDestination `json:"items,omitempty"`
}
