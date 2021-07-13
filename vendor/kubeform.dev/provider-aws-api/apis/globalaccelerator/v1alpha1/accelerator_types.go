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

type Accelerator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcceleratorSpec   `json:"spec,omitempty"`
	Status            AcceleratorStatus `json:"status,omitempty"`
}

type AcceleratorSpecAttributes struct {
	// +optional
	FlowLogsEnabled *bool `json:"flowLogsEnabled,omitempty" tf:"flow_logs_enabled"`
	// +optional
	FlowLogsS3Bucket *string `json:"flowLogsS3Bucket,omitempty" tf:"flow_logs_s3_bucket"`
	// +optional
	FlowLogsS3Prefix *string `json:"flowLogsS3Prefix,omitempty" tf:"flow_logs_s3_prefix"`
}

type AcceleratorSpecIpSets struct {
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty" tf:"ip_addresses"`
	// +optional
	IpFamily *string `json:"ipFamily,omitempty" tf:"ip_family"`
}

type AcceleratorSpec struct {
	State *AcceleratorSpecResource `json:"state,omitempty" tf:"-"`

	Resource AcceleratorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AcceleratorSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Attributes *AcceleratorSpecAttributes `json:"attributes,omitempty" tf:"attributes"`
	// +optional
	DnsName *string `json:"dnsName,omitempty" tf:"dns_name"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	HostedZoneID *string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id"`
	// +optional
	IpAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type"`
	// +optional
	IpSets []AcceleratorSpecIpSets `json:"ipSets,omitempty" tf:"ip_sets"`
	Name   *string                 `json:"name" tf:"name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type AcceleratorStatus struct {
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

// AcceleratorList is a list of Accelerators
type AcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Accelerator CRD objects
	Items []Accelerator `json:"items,omitempty"`
}
