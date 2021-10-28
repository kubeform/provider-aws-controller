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

type SqsQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqsQueueSpec   `json:"spec,omitempty"`
	Status            SqsQueueStatus `json:"status,omitempty"`
}

type SqsQueueSpec struct {
	State *SqsQueueSpecResource `json:"state,omitempty" tf:"-"`

	Resource SqsQueueSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SqsQueueSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication"`
	// +optional
	DeduplicationScope *string `json:"deduplicationScope,omitempty" tf:"deduplication_scope"`
	// +optional
	DelaySeconds *int64 `json:"delaySeconds,omitempty" tf:"delay_seconds"`
	// +optional
	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue"`
	// +optional
	FifoThroughputLimit *string `json:"fifoThroughputLimit,omitempty" tf:"fifo_throughput_limit"`
	// +optional
	KmsDataKeyReusePeriodSeconds *int64 `json:"kmsDataKeyReusePeriodSeconds,omitempty" tf:"kms_data_key_reuse_period_seconds"`
	// +optional
	KmsMasterKeyID *string `json:"kmsMasterKeyID,omitempty" tf:"kms_master_key_id"`
	// +optional
	MaxMessageSize *int64 `json:"maxMessageSize,omitempty" tf:"max_message_size"`
	// +optional
	MessageRetentionSeconds *int64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	Policy *string `json:"policy,omitempty" tf:"policy"`
	// +optional
	ReceiveWaitTimeSeconds *int64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds"`
	// +optional
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
	// +optional
	VisibilityTimeoutSeconds *int64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds"`
}

type SqsQueueStatus struct {
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

// SqsQueueList is a list of SqsQueues
type SqsQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqsQueue CRD objects
	Items []SqsQueue `json:"items,omitempty"`
}
