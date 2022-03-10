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

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpecNodes struct {
	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
}

type ClusterSpecServerSideEncryption struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`
	// +optional
	ClusterAddress *string `json:"clusterAddress,omitempty" tf:"cluster_address"`
	// +optional
	ClusterEndpointEncryptionType *string `json:"clusterEndpointEncryptionType,omitempty" tf:"cluster_endpoint_encryption_type"`
	ClusterName                   *string `json:"clusterName" tf:"cluster_name"`
	// +optional
	ConfigurationEndpoint *string `json:"configurationEndpoint,omitempty" tf:"configuration_endpoint"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	IamRoleArn  *string `json:"iamRoleArn" tf:"iam_role_arn"`
	// +optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`
	NodeType          *string `json:"nodeType" tf:"node_type"`
	// +optional
	Nodes []ClusterSpecNodes `json:"nodes,omitempty" tf:"nodes"`
	// +optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn"`
	// +optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name"`
	// +optional
	Port              *int64 `json:"port,omitempty" tf:"port"`
	ReplicationFactor *int64 `json:"replicationFactor" tf:"replication_factor"`
	// +optional
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids"`
	// +optional
	ServerSideEncryption *ClusterSpecServerSideEncryption `json:"serverSideEncryption,omitempty" tf:"server_side_encryption"`
	// +optional
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ClusterStatus struct {
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

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
