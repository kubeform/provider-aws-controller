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

type HealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HealthCheckSpec   `json:"spec,omitempty"`
	Status            HealthCheckStatus `json:"status,omitempty"`
}

type HealthCheckSpec struct {
	State *HealthCheckSpecResource `json:"state,omitempty" tf:"-"`

	Resource HealthCheckSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type HealthCheckSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ChildHealthThreshold *int64 `json:"childHealthThreshold,omitempty" tf:"child_health_threshold"`
	// +optional
	// +kubebuilder:validation:MaxItems=256
	ChildHealthchecks []string `json:"childHealthchecks,omitempty" tf:"child_healthchecks"`
	// +optional
	CloudwatchAlarmName *string `json:"cloudwatchAlarmName,omitempty" tf:"cloudwatch_alarm_name"`
	// +optional
	CloudwatchAlarmRegion *string `json:"cloudwatchAlarmRegion,omitempty" tf:"cloudwatch_alarm_region"`
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// +optional
	EnableSni *bool `json:"enableSni,omitempty" tf:"enable_sni"`
	// +optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`
	// +optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn"`
	// +optional
	InsufficientDataHealthStatus *string `json:"insufficientDataHealthStatus,omitempty" tf:"insufficient_data_health_status"`
	// +optional
	InvertHealthcheck *bool `json:"invertHealthcheck,omitempty" tf:"invert_healthcheck"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	MeasureLatency *bool `json:"measureLatency,omitempty" tf:"measure_latency"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=3
	Regions []string `json:"regions,omitempty" tf:"regions"`
	// +optional
	RequestInterval *int64 `json:"requestInterval,omitempty" tf:"request_interval"`
	// +optional
	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path"`
	// +optional
	RoutingControlArn *string `json:"routingControlArn,omitempty" tf:"routing_control_arn"`
	// +optional
	SearchString *string `json:"searchString,omitempty" tf:"search_string"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	Type    *string            `json:"type" tf:"type"`
}

type HealthCheckStatus struct {
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

// HealthCheckList is a list of HealthChecks
type HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HealthCheck CRD objects
	Items []HealthCheck `json:"items,omitempty"`
}
