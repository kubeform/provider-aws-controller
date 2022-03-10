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

type ClusterSpecBrokerNodeGroupInfo struct {
	// +optional
	AzDistribution *string  `json:"azDistribution,omitempty" tf:"az_distribution"`
	ClientSubnets  []string `json:"clientSubnets" tf:"client_subnets"`
	EbsVolumeSize  *int64   `json:"ebsVolumeSize" tf:"ebs_volume_size"`
	InstanceType   *string  `json:"instanceType" tf:"instance_type"`
	SecurityGroups []string `json:"securityGroups" tf:"security_groups"`
}

type ClusterSpecClientAuthenticationSasl struct {
	// +optional
	Iam *bool `json:"iam,omitempty" tf:"iam"`
	// +optional
	Scram *bool `json:"scram,omitempty" tf:"scram"`
}

type ClusterSpecClientAuthenticationTls struct {
	// +optional
	CertificateAuthorityArns []string `json:"certificateAuthorityArns,omitempty" tf:"certificate_authority_arns"`
}

type ClusterSpecClientAuthentication struct {
	// +optional
	Sasl *ClusterSpecClientAuthenticationSasl `json:"sasl,omitempty" tf:"sasl"`
	// +optional
	Tls *ClusterSpecClientAuthenticationTls `json:"tls,omitempty" tf:"tls"`
}

type ClusterSpecConfigurationInfo struct {
	Arn      *string `json:"arn" tf:"arn"`
	Revision *int64  `json:"revision" tf:"revision"`
}

type ClusterSpecEncryptionInfoEncryptionInTransit struct {
	// +optional
	ClientBroker *string `json:"clientBroker,omitempty" tf:"client_broker"`
	// +optional
	InCluster *bool `json:"inCluster,omitempty" tf:"in_cluster"`
}

type ClusterSpecEncryptionInfo struct {
	// +optional
	EncryptionAtRestKmsKeyArn *string `json:"encryptionAtRestKmsKeyArn,omitempty" tf:"encryption_at_rest_kms_key_arn"`
	// +optional
	EncryptionInTransit *ClusterSpecEncryptionInfoEncryptionInTransit `json:"encryptionInTransit,omitempty" tf:"encryption_in_transit"`
}

type ClusterSpecLoggingInfoBrokerLogsCloudwatchLogs struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	LogGroup *string `json:"logGroup,omitempty" tf:"log_group"`
}

type ClusterSpecLoggingInfoBrokerLogsFirehose struct {
	// +optional
	DeliveryStream *string `json:"deliveryStream,omitempty" tf:"delivery_stream"`
	Enabled        *bool   `json:"enabled" tf:"enabled"`
}

type ClusterSpecLoggingInfoBrokerLogsS3 struct {
	// +optional
	Bucket  *string `json:"bucket,omitempty" tf:"bucket"`
	Enabled *bool   `json:"enabled" tf:"enabled"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type ClusterSpecLoggingInfoBrokerLogs struct {
	// +optional
	CloudwatchLogs *ClusterSpecLoggingInfoBrokerLogsCloudwatchLogs `json:"cloudwatchLogs,omitempty" tf:"cloudwatch_logs"`
	// +optional
	Firehose *ClusterSpecLoggingInfoBrokerLogsFirehose `json:"firehose,omitempty" tf:"firehose"`
	// +optional
	S3 *ClusterSpecLoggingInfoBrokerLogsS3 `json:"s3,omitempty" tf:"s3"`
}

type ClusterSpecLoggingInfo struct {
	BrokerLogs *ClusterSpecLoggingInfoBrokerLogs `json:"brokerLogs" tf:"broker_logs"`
}

type ClusterSpecOpenMonitoringPrometheusJmxExporter struct {
	EnabledInBroker *bool `json:"enabledInBroker" tf:"enabled_in_broker"`
}

type ClusterSpecOpenMonitoringPrometheusNodeExporter struct {
	EnabledInBroker *bool `json:"enabledInBroker" tf:"enabled_in_broker"`
}

type ClusterSpecOpenMonitoringPrometheus struct {
	// +optional
	JmxExporter *ClusterSpecOpenMonitoringPrometheusJmxExporter `json:"jmxExporter,omitempty" tf:"jmx_exporter"`
	// +optional
	NodeExporter *ClusterSpecOpenMonitoringPrometheusNodeExporter `json:"nodeExporter,omitempty" tf:"node_exporter"`
}

type ClusterSpecOpenMonitoring struct {
	Prometheus *ClusterSpecOpenMonitoringPrometheus `json:"prometheus" tf:"prometheus"`
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
	BootstrapBrokers *string `json:"bootstrapBrokers,omitempty" tf:"bootstrap_brokers"`
	// +optional
	BootstrapBrokersSaslIam *string `json:"bootstrapBrokersSaslIam,omitempty" tf:"bootstrap_brokers_sasl_iam"`
	// +optional
	BootstrapBrokersSaslScram *string `json:"bootstrapBrokersSaslScram,omitempty" tf:"bootstrap_brokers_sasl_scram"`
	// +optional
	BootstrapBrokersTls *string                         `json:"bootstrapBrokersTls,omitempty" tf:"bootstrap_brokers_tls"`
	BrokerNodeGroupInfo *ClusterSpecBrokerNodeGroupInfo `json:"brokerNodeGroupInfo" tf:"broker_node_group_info"`
	// +optional
	ClientAuthentication *ClusterSpecClientAuthentication `json:"clientAuthentication,omitempty" tf:"client_authentication"`
	ClusterName          *string                          `json:"clusterName" tf:"cluster_name"`
	// +optional
	ConfigurationInfo *ClusterSpecConfigurationInfo `json:"configurationInfo,omitempty" tf:"configuration_info"`
	// +optional
	CurrentVersion *string `json:"currentVersion,omitempty" tf:"current_version"`
	// +optional
	EncryptionInfo *ClusterSpecEncryptionInfo `json:"encryptionInfo,omitempty" tf:"encryption_info"`
	// +optional
	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty" tf:"enhanced_monitoring"`
	KafkaVersion       *string `json:"kafkaVersion" tf:"kafka_version"`
	// +optional
	LoggingInfo         *ClusterSpecLoggingInfo `json:"loggingInfo,omitempty" tf:"logging_info"`
	NumberOfBrokerNodes *int64                  `json:"numberOfBrokerNodes" tf:"number_of_broker_nodes"`
	// +optional
	OpenMonitoring *ClusterSpecOpenMonitoring `json:"openMonitoring,omitempty" tf:"open_monitoring"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	ZookeeperConnectString *string `json:"zookeeperConnectString,omitempty" tf:"zookeeper_connect_string"`
	// +optional
	ZookeeperConnectStringTls *string `json:"zookeeperConnectStringTls,omitempty" tf:"zookeeper_connect_string_tls"`
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
