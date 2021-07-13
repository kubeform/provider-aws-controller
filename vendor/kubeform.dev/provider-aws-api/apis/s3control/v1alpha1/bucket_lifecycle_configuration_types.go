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

type BucketLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketLifecycleConfigurationSpec   `json:"spec,omitempty"`
	Status            BucketLifecycleConfigurationStatus `json:"status,omitempty"`
}

type BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload struct {
	DaysAfterInitiation *int64 `json:"daysAfterInitiation" tf:"days_after_initiation"`
}

type BucketLifecycleConfigurationSpecRuleExpiration struct {
	// +optional
	Date *string `json:"date,omitempty" tf:"date"`
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
	// +optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker"`
}

type BucketLifecycleConfigurationSpecRuleFilter struct {
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type BucketLifecycleConfigurationSpecRule struct {
	// +optional
	AbortIncompleteMultipartUpload *BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload"`
	// +optional
	Expiration *BucketLifecycleConfigurationSpecRuleExpiration `json:"expiration,omitempty" tf:"expiration"`
	// +optional
	Filter *BucketLifecycleConfigurationSpecRuleFilter `json:"filter,omitempty" tf:"filter"`
	ID     *string                                     `json:"ID" tf:"id"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type BucketLifecycleConfigurationSpec struct {
	State *BucketLifecycleConfigurationSpecResource `json:"state,omitempty" tf:"-"`

	Resource BucketLifecycleConfigurationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BucketLifecycleConfigurationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bucket *string `json:"bucket" tf:"bucket"`
	// +kubebuilder:validation:MinItems=1
	Rule []BucketLifecycleConfigurationSpecRule `json:"rule" tf:"rule"`
}

type BucketLifecycleConfigurationStatus struct {
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

// BucketLifecycleConfigurationList is a list of BucketLifecycleConfigurations
type BucketLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BucketLifecycleConfiguration CRD objects
	Items []BucketLifecycleConfiguration `json:"items,omitempty"`
}
