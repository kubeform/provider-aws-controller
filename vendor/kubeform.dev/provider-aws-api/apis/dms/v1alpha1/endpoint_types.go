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

type Endpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec,omitempty"`
	Status            EndpointStatus `json:"status,omitempty"`
}

type EndpointSpecElasticsearchSettings struct {
	EndpointURI *string `json:"endpointURI" tf:"endpoint_uri"`
	// +optional
	ErrorRetryDuration *int64 `json:"errorRetryDuration,omitempty" tf:"error_retry_duration"`
	// +optional
	FullLoadErrorPercentage *int64  `json:"fullLoadErrorPercentage,omitempty" tf:"full_load_error_percentage"`
	ServiceAccessRoleArn    *string `json:"serviceAccessRoleArn" tf:"service_access_role_arn"`
}

type EndpointSpecKafkaSettings struct {
	Broker *string `json:"broker" tf:"broker"`
	// +optional
	IncludeControlDetails *bool `json:"includeControlDetails,omitempty" tf:"include_control_details"`
	// +optional
	IncludeNullAndEmpty *bool `json:"includeNullAndEmpty,omitempty" tf:"include_null_and_empty"`
	// +optional
	IncludePartitionValue *bool `json:"includePartitionValue,omitempty" tf:"include_partition_value"`
	// +optional
	IncludeTableAlterOperations *bool `json:"includeTableAlterOperations,omitempty" tf:"include_table_alter_operations"`
	// +optional
	IncludeTransactionDetails *bool `json:"includeTransactionDetails,omitempty" tf:"include_transaction_details"`
	// +optional
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format"`
	// +optional
	MessageMaxBytes *int64 `json:"messageMaxBytes,omitempty" tf:"message_max_bytes"`
	// +optional
	NoHexPrefix *bool `json:"noHexPrefix,omitempty" tf:"no_hex_prefix"`
	// +optional
	PartitionIncludeSchemaTable *bool `json:"partitionIncludeSchemaTable,omitempty" tf:"partition_include_schema_table"`
	// +optional
	SaslPassword *string `json:"-" sensitive:"true" tf:"sasl_password"`
	// +optional
	SaslUsername *string `json:"saslUsername,omitempty" tf:"sasl_username"`
	// +optional
	SecurityProtocol *string `json:"securityProtocol,omitempty" tf:"security_protocol"`
	// +optional
	SslCaCertificateArn *string `json:"sslCaCertificateArn,omitempty" tf:"ssl_ca_certificate_arn"`
	// +optional
	SslClientCertificateArn *string `json:"sslClientCertificateArn,omitempty" tf:"ssl_client_certificate_arn"`
	// +optional
	SslClientKeyArn *string `json:"sslClientKeyArn,omitempty" tf:"ssl_client_key_arn"`
	// +optional
	SslClientKeyPassword *string `json:"-" sensitive:"true" tf:"ssl_client_key_password"`
	// +optional
	Topic *string `json:"topic,omitempty" tf:"topic"`
}

type EndpointSpecKinesisSettings struct {
	// +optional
	IncludeControlDetails *bool `json:"includeControlDetails,omitempty" tf:"include_control_details"`
	// +optional
	IncludeNullAndEmpty *bool `json:"includeNullAndEmpty,omitempty" tf:"include_null_and_empty"`
	// +optional
	IncludePartitionValue *bool `json:"includePartitionValue,omitempty" tf:"include_partition_value"`
	// +optional
	IncludeTableAlterOperations *bool `json:"includeTableAlterOperations,omitempty" tf:"include_table_alter_operations"`
	// +optional
	IncludeTransactionDetails *bool `json:"includeTransactionDetails,omitempty" tf:"include_transaction_details"`
	// +optional
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format"`
	// +optional
	PartitionIncludeSchemaTable *bool `json:"partitionIncludeSchemaTable,omitempty" tf:"partition_include_schema_table"`
	// +optional
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn,omitempty" tf:"service_access_role_arn"`
	// +optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn"`
}

type EndpointSpecMongodbSettings struct {
	// +optional
	AuthMechanism *string `json:"authMechanism,omitempty" tf:"auth_mechanism"`
	// +optional
	AuthSource *string `json:"authSource,omitempty" tf:"auth_source"`
	// +optional
	AuthType *string `json:"authType,omitempty" tf:"auth_type"`
	// +optional
	DocsToInvestigate *string `json:"docsToInvestigate,omitempty" tf:"docs_to_investigate"`
	// +optional
	ExtractDocID *string `json:"extractDocID,omitempty" tf:"extract_doc_id"`
	// +optional
	NestingLevel *string `json:"nestingLevel,omitempty" tf:"nesting_level"`
}

type EndpointSpecS3Settings struct {
	// +optional
	AddColumnName *bool `json:"addColumnName,omitempty" tf:"add_column_name"`
	// +optional
	BucketFolder *string `json:"bucketFolder,omitempty" tf:"bucket_folder"`
	// +optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name"`
	// +optional
	CannedACLForObjects *string `json:"cannedACLForObjects,omitempty" tf:"canned_acl_for_objects"`
	// +optional
	CdcInsertsAndUpdates *bool `json:"cdcInsertsAndUpdates,omitempty" tf:"cdc_inserts_and_updates"`
	// +optional
	CdcInsertsOnly *bool `json:"cdcInsertsOnly,omitempty" tf:"cdc_inserts_only"`
	// +optional
	CdcMaxBatchInterval *int64 `json:"cdcMaxBatchInterval,omitempty" tf:"cdc_max_batch_interval"`
	// +optional
	CdcMinFileSize *int64 `json:"cdcMinFileSize,omitempty" tf:"cdc_min_file_size"`
	// +optional
	CdcPath *string `json:"cdcPath,omitempty" tf:"cdc_path"`
	// +optional
	CompressionType *string `json:"compressionType,omitempty" tf:"compression_type"`
	// +optional
	CsvDelimiter *string `json:"csvDelimiter,omitempty" tf:"csv_delimiter"`
	// +optional
	CsvNoSupValue *string `json:"csvNoSupValue,omitempty" tf:"csv_no_sup_value"`
	// +optional
	CsvNullValue *string `json:"csvNullValue,omitempty" tf:"csv_null_value"`
	// +optional
	CsvRowDelimiter *string `json:"csvRowDelimiter,omitempty" tf:"csv_row_delimiter"`
	// +optional
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format"`
	// +optional
	DataPageSize *int64 `json:"dataPageSize,omitempty" tf:"data_page_size"`
	// +optional
	DatePartitionDelimiter *string `json:"datePartitionDelimiter,omitempty" tf:"date_partition_delimiter"`
	// +optional
	DatePartitionEnabled *bool `json:"datePartitionEnabled,omitempty" tf:"date_partition_enabled"`
	// +optional
	DatePartitionSequence *string `json:"datePartitionSequence,omitempty" tf:"date_partition_sequence"`
	// +optional
	DictPageSizeLimit *int64 `json:"dictPageSizeLimit,omitempty" tf:"dict_page_size_limit"`
	// +optional
	EnableStatistics *bool `json:"enableStatistics,omitempty" tf:"enable_statistics"`
	// +optional
	EncodingType *string `json:"encodingType,omitempty" tf:"encoding_type"`
	// +optional
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode"`
	// +optional
	ExternalTableDefinition *string `json:"externalTableDefinition,omitempty" tf:"external_table_definition"`
	// +optional
	IgnoreHeadersRow *int64 `json:"ignoreHeadersRow,omitempty" tf:"ignore_headers_row"`
	// +optional
	IncludeOpForFullLoad *bool `json:"includeOpForFullLoad,omitempty" tf:"include_op_for_full_load"`
	// +optional
	MaxFileSize *int64 `json:"maxFileSize,omitempty" tf:"max_file_size"`
	// +optional
	ParquetTimestampInMillisecond *bool `json:"parquetTimestampInMillisecond,omitempty" tf:"parquet_timestamp_in_millisecond"`
	// +optional
	ParquetVersion *string `json:"parquetVersion,omitempty" tf:"parquet_version"`
	// +optional
	PreserveTransactions *bool `json:"preserveTransactions,omitempty" tf:"preserve_transactions"`
	// +optional
	Rfc4180 *bool `json:"rfc4180,omitempty" tf:"rfc_4180"`
	// +optional
	RowGroupLength *int64 `json:"rowGroupLength,omitempty" tf:"row_group_length"`
	// +optional
	ServerSideEncryptionKmsKeyID *string `json:"serverSideEncryptionKmsKeyID,omitempty" tf:"server_side_encryption_kms_key_id"`
	// +optional
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn,omitempty" tf:"service_access_role_arn"`
	// +optional
	TimestampColumnName *string `json:"timestampColumnName,omitempty" tf:"timestamp_column_name"`
	// +optional
	UseCsvNoSupValue *bool `json:"useCsvNoSupValue,omitempty" tf:"use_csv_no_sup_value"`
}

type EndpointSpec struct {
	State *EndpointSpecResource `json:"state,omitempty" tf:"-"`

	Resource EndpointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EndpointSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn"`
	// +optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name"`
	// +optional
	ElasticsearchSettings *EndpointSpecElasticsearchSettings `json:"elasticsearchSettings,omitempty" tf:"elasticsearch_settings"`
	// +optional
	EndpointArn  *string `json:"endpointArn,omitempty" tf:"endpoint_arn"`
	EndpointID   *string `json:"endpointID" tf:"endpoint_id"`
	EndpointType *string `json:"endpointType" tf:"endpoint_type"`
	EngineName   *string `json:"engineName" tf:"engine_name"`
	// +optional
	ExtraConnectionAttributes *string `json:"extraConnectionAttributes,omitempty" tf:"extra_connection_attributes"`
	// +optional
	KafkaSettings *EndpointSpecKafkaSettings `json:"kafkaSettings,omitempty" tf:"kafka_settings"`
	// +optional
	KinesisSettings *EndpointSpecKinesisSettings `json:"kinesisSettings,omitempty" tf:"kinesis_settings"`
	// +optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
	// +optional
	MongodbSettings *EndpointSpecMongodbSettings `json:"mongodbSettings,omitempty" tf:"mongodb_settings"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	S3Settings *EndpointSpecS3Settings `json:"s3Settings,omitempty" tf:"s3_settings"`
	// +optional
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn,omitempty" tf:"secrets_manager_access_role_arn"`
	// +optional
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn"`
	// +optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name"`
	// +optional
	ServiceAccessRole *string `json:"serviceAccessRole,omitempty" tf:"service_access_role"`
	// +optional
	SslMode *string `json:"sslMode,omitempty" tf:"ssl_mode"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type EndpointStatus struct {
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

// EndpointList is a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Endpoint CRD objects
	Items []Endpoint `json:"items,omitempty"`
}
