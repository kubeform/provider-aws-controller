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

type LogMetricFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogMetricFilterSpec   `json:"spec,omitempty"`
	Status            LogMetricFilterStatus `json:"status,omitempty"`
}

type LogMetricFilterSpecMetricTransformation struct {
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value"`
	// +optional
	Dimensions *map[string]string `json:"dimensions,omitempty" tf:"dimensions"`
	Name       *string            `json:"name" tf:"name"`
	Namespace  *string            `json:"namespace" tf:"namespace"`
	// +optional
	Unit  *string `json:"unit,omitempty" tf:"unit"`
	Value *string `json:"value" tf:"value"`
}

type LogMetricFilterSpec struct {
	State *LogMetricFilterSpecResource `json:"state,omitempty" tf:"-"`

	Resource LogMetricFilterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type LogMetricFilterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	LogGroupName         *string                                  `json:"logGroupName" tf:"log_group_name"`
	MetricTransformation *LogMetricFilterSpecMetricTransformation `json:"metricTransformation" tf:"metric_transformation"`
	Name                 *string                                  `json:"name" tf:"name"`
	Pattern              *string                                  `json:"pattern" tf:"pattern"`
}

type LogMetricFilterStatus struct {
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

// LogMetricFilterList is a list of LogMetricFilters
type LogMetricFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogMetricFilter CRD objects
	Items []LogMetricFilter `json:"items,omitempty"`
}
