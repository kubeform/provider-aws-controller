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

type Datasource struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasourceSpec   `json:"spec,omitempty"`
	Status            DatasourceStatus `json:"status,omitempty"`
}

type DatasourceSpecDynamodbConfig struct {
	// +optional
	Region    *string `json:"region,omitempty" tf:"region"`
	TableName *string `json:"tableName" tf:"table_name"`
	// +optional
	UseCallerCredentials *bool `json:"useCallerCredentials,omitempty" tf:"use_caller_credentials"`
}

type DatasourceSpecElasticsearchConfig struct {
	Endpoint *string `json:"endpoint" tf:"endpoint"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
}

type DatasourceSpecHttpConfig struct {
	Endpoint *string `json:"endpoint" tf:"endpoint"`
}

type DatasourceSpecLambdaConfig struct {
	FunctionArn *string `json:"functionArn" tf:"function_arn"`
}

type DatasourceSpec struct {
	State *DatasourceSpecResource `json:"state,omitempty" tf:"-"`

	Resource DatasourceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DatasourceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiID *string `json:"apiID" tf:"api_id"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DynamodbConfig *DatasourceSpecDynamodbConfig `json:"dynamodbConfig,omitempty" tf:"dynamodb_config"`
	// +optional
	ElasticsearchConfig *DatasourceSpecElasticsearchConfig `json:"elasticsearchConfig,omitempty" tf:"elasticsearch_config"`
	// +optional
	HttpConfig *DatasourceSpecHttpConfig `json:"httpConfig,omitempty" tf:"http_config"`
	// +optional
	LambdaConfig *DatasourceSpecLambdaConfig `json:"lambdaConfig,omitempty" tf:"lambda_config"`
	Name         *string                     `json:"name" tf:"name"`
	// +optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn"`
	Type           *string `json:"type" tf:"type"`
}

type DatasourceStatus struct {
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

// DatasourceList is a list of Datasources
type DatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Datasource CRD objects
	Items []Datasource `json:"items,omitempty"`
}
