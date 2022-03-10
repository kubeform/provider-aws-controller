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

type NetworkProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkProfileSpec   `json:"spec,omitempty"`
	Status            NetworkProfileStatus `json:"status,omitempty"`
}

type NetworkProfileSpec struct {
	State *NetworkProfileSpecResource `json:"state,omitempty" tf:"-"`

	Resource NetworkProfileSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type NetworkProfileSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DownlinkBandwidthBits *int64 `json:"downlinkBandwidthBits,omitempty" tf:"downlink_bandwidth_bits"`
	// +optional
	DownlinkDelayMs *int64 `json:"downlinkDelayMs,omitempty" tf:"downlink_delay_ms"`
	// +optional
	DownlinkJitterMs *int64 `json:"downlinkJitterMs,omitempty" tf:"downlink_jitter_ms"`
	// +optional
	DownlinkLossPercent *int64  `json:"downlinkLossPercent,omitempty" tf:"downlink_loss_percent"`
	Name                *string `json:"name" tf:"name"`
	ProjectArn          *string `json:"projectArn" tf:"project_arn"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	UplinkBandwidthBits *int64 `json:"uplinkBandwidthBits,omitempty" tf:"uplink_bandwidth_bits"`
	// +optional
	UplinkDelayMs *int64 `json:"uplinkDelayMs,omitempty" tf:"uplink_delay_ms"`
	// +optional
	UplinkJitterMs *int64 `json:"uplinkJitterMs,omitempty" tf:"uplink_jitter_ms"`
	// +optional
	UplinkLossPercent *int64 `json:"uplinkLossPercent,omitempty" tf:"uplink_loss_percent"`
}

type NetworkProfileStatus struct {
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

// NetworkProfileList is a list of NetworkProfiles
type NetworkProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkProfile CRD objects
	Items []NetworkProfile `json:"items,omitempty"`
}