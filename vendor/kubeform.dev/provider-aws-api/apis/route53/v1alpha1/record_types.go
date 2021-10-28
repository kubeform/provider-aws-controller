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

type Record struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordSpec   `json:"spec,omitempty"`
	Status            RecordStatus `json:"status,omitempty"`
}

type RecordSpecAlias struct {
	EvaluateTargetHealth *bool   `json:"evaluateTargetHealth" tf:"evaluate_target_health"`
	Name                 *string `json:"name" tf:"name"`
	ZoneID               *string `json:"zoneID" tf:"zone_id"`
}

type RecordSpecFailoverRoutingPolicy struct {
	Type *string `json:"type" tf:"type"`
}

type RecordSpecGeolocationRoutingPolicy struct {
	// +optional
	Continent *string `json:"continent,omitempty" tf:"continent"`
	// +optional
	Country *string `json:"country,omitempty" tf:"country"`
	// +optional
	Subdivision *string `json:"subdivision,omitempty" tf:"subdivision"`
}

type RecordSpecLatencyRoutingPolicy struct {
	Region *string `json:"region" tf:"region"`
}

type RecordSpecWeightedRoutingPolicy struct {
	Weight *int64 `json:"weight" tf:"weight"`
}

type RecordSpec struct {
	State *RecordSpecResource `json:"state,omitempty" tf:"-"`

	Resource RecordSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RecordSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Alias []RecordSpecAlias `json:"alias,omitempty" tf:"alias"`
	// +optional
	AllowOverwrite *bool `json:"allowOverwrite,omitempty" tf:"allow_overwrite"`
	// +optional
	FailoverRoutingPolicy []RecordSpecFailoverRoutingPolicy `json:"failoverRoutingPolicy,omitempty" tf:"failover_routing_policy"`
	// +optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn"`
	// +optional
	GeolocationRoutingPolicy []RecordSpecGeolocationRoutingPolicy `json:"geolocationRoutingPolicy,omitempty" tf:"geolocation_routing_policy"`
	// +optional
	HealthCheckID *string `json:"healthCheckID,omitempty" tf:"health_check_id"`
	// +optional
	LatencyRoutingPolicy []RecordSpecLatencyRoutingPolicy `json:"latencyRoutingPolicy,omitempty" tf:"latency_routing_policy"`
	// +optional
	MultivalueAnswerRoutingPolicy *bool   `json:"multivalueAnswerRoutingPolicy,omitempty" tf:"multivalue_answer_routing_policy"`
	Name                          *string `json:"name" tf:"name"`
	// +optional
	Records []string `json:"records,omitempty" tf:"records"`
	// +optional
	SetIdentifier *string `json:"setIdentifier,omitempty" tf:"set_identifier"`
	// +optional
	Ttl  *int64  `json:"ttl,omitempty" tf:"ttl"`
	Type *string `json:"type" tf:"type"`
	// +optional
	WeightedRoutingPolicy []RecordSpecWeightedRoutingPolicy `json:"weightedRoutingPolicy,omitempty" tf:"weighted_routing_policy"`
	ZoneID                *string                           `json:"zoneID" tf:"zone_id"`
}

type RecordStatus struct {
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

// RecordList is a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Record CRD objects
	Items []Record `json:"items,omitempty"`
}
