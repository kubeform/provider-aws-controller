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

type Application struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec,omitempty"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

type ApplicationSpecApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation struct {
	BucketArn *string `json:"bucketArn" tf:"bucket_arn"`
	FileKey   *string `json:"fileKey" tf:"file_key"`
	// +optional
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version"`
}

type ApplicationSpecApplicationConfigurationApplicationCodeConfigurationCodeContent struct {
	// +optional
	S3ContentLocation *ApplicationSpecApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation `json:"s3ContentLocation,omitempty" tf:"s3_content_location"`
	// +optional
	TextContent *string `json:"textContent,omitempty" tf:"text_content"`
}

type ApplicationSpecApplicationConfigurationApplicationCodeConfiguration struct {
	// +optional
	CodeContent     *ApplicationSpecApplicationConfigurationApplicationCodeConfigurationCodeContent `json:"codeContent,omitempty" tf:"code_content"`
	CodeContentType *string                                                                         `json:"codeContentType" tf:"code_content_type"`
}

type ApplicationSpecApplicationConfigurationApplicationSnapshotConfiguration struct {
	SnapshotsEnabled *bool `json:"snapshotsEnabled" tf:"snapshots_enabled"`
}

type ApplicationSpecApplicationConfigurationEnvironmentPropertiesPropertyGroup struct {
	PropertyGroupID *string            `json:"propertyGroupID" tf:"property_group_id"`
	PropertyMap     *map[string]string `json:"propertyMap" tf:"property_map"`
}

type ApplicationSpecApplicationConfigurationEnvironmentProperties struct {
	// +kubebuilder:validation:MaxItems=50
	PropertyGroup []ApplicationSpecApplicationConfigurationEnvironmentPropertiesPropertyGroup `json:"propertyGroup" tf:"property_group"`
}

type ApplicationSpecApplicationConfigurationFlinkApplicationConfigurationCheckpointConfiguration struct {
	// +optional
	CheckpointInterval *int64 `json:"checkpointInterval,omitempty" tf:"checkpoint_interval"`
	// +optional
	CheckpointingEnabled *bool   `json:"checkpointingEnabled,omitempty" tf:"checkpointing_enabled"`
	ConfigurationType    *string `json:"configurationType" tf:"configuration_type"`
	// +optional
	MinPauseBetweenCheckpoints *int64 `json:"minPauseBetweenCheckpoints,omitempty" tf:"min_pause_between_checkpoints"`
}

type ApplicationSpecApplicationConfigurationFlinkApplicationConfigurationMonitoringConfiguration struct {
	ConfigurationType *string `json:"configurationType" tf:"configuration_type"`
	// +optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
	// +optional
	MetricsLevel *string `json:"metricsLevel,omitempty" tf:"metrics_level"`
}

type ApplicationSpecApplicationConfigurationFlinkApplicationConfigurationParallelismConfiguration struct {
	// +optional
	AutoScalingEnabled *bool   `json:"autoScalingEnabled,omitempty" tf:"auto_scaling_enabled"`
	ConfigurationType  *string `json:"configurationType" tf:"configuration_type"`
	// +optional
	Parallelism *int64 `json:"parallelism,omitempty" tf:"parallelism"`
	// +optional
	ParallelismPerKpu *int64 `json:"parallelismPerKpu,omitempty" tf:"parallelism_per_kpu"`
}

type ApplicationSpecApplicationConfigurationFlinkApplicationConfiguration struct {
	// +optional
	CheckpointConfiguration *ApplicationSpecApplicationConfigurationFlinkApplicationConfigurationCheckpointConfiguration `json:"checkpointConfiguration,omitempty" tf:"checkpoint_configuration"`
	// +optional
	MonitoringConfiguration *ApplicationSpecApplicationConfigurationFlinkApplicationConfigurationMonitoringConfiguration `json:"monitoringConfiguration,omitempty" tf:"monitoring_configuration"`
	// +optional
	ParallelismConfiguration *ApplicationSpecApplicationConfigurationFlinkApplicationConfigurationParallelismConfiguration `json:"parallelismConfiguration,omitempty" tf:"parallelism_configuration"`
}

type ApplicationSpecApplicationConfigurationRunConfigurationApplicationRestoreConfiguration struct {
	// +optional
	ApplicationRestoreType *string `json:"applicationRestoreType,omitempty" tf:"application_restore_type"`
	// +optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name"`
}

type ApplicationSpecApplicationConfigurationRunConfigurationFlinkRunConfiguration struct {
	// +optional
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty" tf:"allow_non_restored_state"`
}

type ApplicationSpecApplicationConfigurationRunConfiguration struct {
	// +optional
	ApplicationRestoreConfiguration *ApplicationSpecApplicationConfigurationRunConfigurationApplicationRestoreConfiguration `json:"applicationRestoreConfiguration,omitempty" tf:"application_restore_configuration"`
	// +optional
	FlinkRunConfiguration *ApplicationSpecApplicationConfigurationRunConfigurationFlinkRunConfiguration `json:"flinkRunConfiguration,omitempty" tf:"flink_run_configuration"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputParallelism struct {
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfigurationInputLambdaProcessor struct {
	ResourceArn *string `json:"resourceArn" tf:"resource_arn"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfiguration struct {
	InputLambdaProcessor *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfigurationInputLambdaProcessor `json:"inputLambdaProcessor" tf:"input_lambda_processor"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumn struct {
	// +optional
	Mapping *string `json:"mapping,omitempty" tf:"mapping"`
	Name    *string `json:"name" tf:"name"`
	SqlType *string `json:"sqlType" tf:"sql_type"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParameters struct {
	RecordColumnDelimiter *string `json:"recordColumnDelimiter" tf:"record_column_delimiter"`
	RecordRowDelimiter    *string `json:"recordRowDelimiter" tf:"record_row_delimiter"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParameters struct {
	RecordRowPath *string `json:"recordRowPath" tf:"record_row_path"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParameters struct {
	// +optional
	CsvMappingParameters *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParameters `json:"csvMappingParameters,omitempty" tf:"csv_mapping_parameters"`
	// +optional
	JsonMappingParameters *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParameters `json:"jsonMappingParameters,omitempty" tf:"json_mapping_parameters"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormat struct {
	MappingParameters *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParameters `json:"mappingParameters" tf:"mapping_parameters"`
	RecordFormatType  *string                                                                                                          `json:"recordFormatType" tf:"record_format_type"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchema struct {
	// +kubebuilder:validation:MaxItems=1000
	RecordColumn []ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumn `json:"recordColumn" tf:"record_column"`
	// +optional
	RecordEncoding *string                                                                                         `json:"recordEncoding,omitempty" tf:"record_encoding"`
	RecordFormat   *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormat `json:"recordFormat" tf:"record_format"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfiguration struct {
	// +optional
	InputStartingPosition *string `json:"inputStartingPosition,omitempty" tf:"input_starting_position"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputKinesisFirehoseInput struct {
	ResourceArn *string `json:"resourceArn" tf:"resource_arn"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInput struct {
	ResourceArn *string `json:"resourceArn" tf:"resource_arn"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInput struct {
	// +optional
	InAppStreamNames []string `json:"inAppStreamNames,omitempty" tf:"in_app_stream_names"`
	// +optional
	InputID *string `json:"inputID,omitempty" tf:"input_id"`
	// +optional
	InputParallelism *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputParallelism `json:"inputParallelism,omitempty" tf:"input_parallelism"`
	// +optional
	InputProcessingConfiguration *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfiguration `json:"inputProcessingConfiguration,omitempty" tf:"input_processing_configuration"`
	InputSchema                  *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputSchema                  `json:"inputSchema" tf:"input_schema"`
	// +optional
	InputStartingPositionConfiguration []ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfiguration `json:"inputStartingPositionConfiguration,omitempty" tf:"input_starting_position_configuration"`
	// +optional
	KinesisFirehoseInput *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputKinesisFirehoseInput `json:"kinesisFirehoseInput,omitempty" tf:"kinesis_firehose_input"`
	// +optional
	KinesisStreamsInput *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInput `json:"kinesisStreamsInput,omitempty" tf:"kinesis_streams_input"`
	NamePrefix          *string                                                                                     `json:"namePrefix" tf:"name_prefix"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchema struct {
	RecordFormatType *string `json:"recordFormatType" tf:"record_format_type"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutput struct {
	ResourceArn *string `json:"resourceArn" tf:"resource_arn"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutputKinesisStreamsOutput struct {
	ResourceArn *string `json:"resourceArn" tf:"resource_arn"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutput struct {
	ResourceArn *string `json:"resourceArn" tf:"resource_arn"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutput struct {
	DestinationSchema *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchema `json:"destinationSchema" tf:"destination_schema"`
	// +optional
	KinesisFirehoseOutput *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutput `json:"kinesisFirehoseOutput,omitempty" tf:"kinesis_firehose_output"`
	// +optional
	KinesisStreamsOutput *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutputKinesisStreamsOutput `json:"kinesisStreamsOutput,omitempty" tf:"kinesis_streams_output"`
	// +optional
	LambdaOutput *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutput `json:"lambdaOutput,omitempty" tf:"lambda_output"`
	Name         *string                                                                               `json:"name" tf:"name"`
	// +optional
	OutputID *string `json:"outputID,omitempty" tf:"output_id"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumn struct {
	// +optional
	Mapping *string `json:"mapping,omitempty" tf:"mapping"`
	Name    *string `json:"name" tf:"name"`
	SqlType *string `json:"sqlType" tf:"sql_type"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersCsvMappingParameters struct {
	RecordColumnDelimiter *string `json:"recordColumnDelimiter" tf:"record_column_delimiter"`
	RecordRowDelimiter    *string `json:"recordRowDelimiter" tf:"record_row_delimiter"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParameters struct {
	RecordRowPath *string `json:"recordRowPath" tf:"record_row_path"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParameters struct {
	// +optional
	CsvMappingParameters *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersCsvMappingParameters `json:"csvMappingParameters,omitempty" tf:"csv_mapping_parameters"`
	// +optional
	JsonMappingParameters *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParameters `json:"jsonMappingParameters,omitempty" tf:"json_mapping_parameters"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormat struct {
	MappingParameters *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParameters `json:"mappingParameters" tf:"mapping_parameters"`
	RecordFormatType  *string                                                                                                                            `json:"recordFormatType" tf:"record_format_type"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchema struct {
	// +kubebuilder:validation:MaxItems=1000
	RecordColumn []ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumn `json:"recordColumn" tf:"record_column"`
	// +optional
	RecordEncoding *string                                                                                                           `json:"recordEncoding,omitempty" tf:"record_encoding"`
	RecordFormat   *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormat `json:"recordFormat" tf:"record_format"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSource struct {
	BucketArn *string `json:"bucketArn" tf:"bucket_arn"`
	FileKey   *string `json:"fileKey" tf:"file_key"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSource struct {
	// +optional
	ReferenceID           *string                                                                                                     `json:"referenceID,omitempty" tf:"reference_id"`
	ReferenceSchema       *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchema       `json:"referenceSchema" tf:"reference_schema"`
	S3ReferenceDataSource *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSource `json:"s3ReferenceDataSource" tf:"s3_reference_data_source"`
	TableName             *string                                                                                                     `json:"tableName" tf:"table_name"`
}

type ApplicationSpecApplicationConfigurationSqlApplicationConfiguration struct {
	// +optional
	Input *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationInput `json:"input,omitempty" tf:"input"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	Output []ApplicationSpecApplicationConfigurationSqlApplicationConfigurationOutput `json:"output,omitempty" tf:"output"`
	// +optional
	ReferenceDataSource *ApplicationSpecApplicationConfigurationSqlApplicationConfigurationReferenceDataSource `json:"referenceDataSource,omitempty" tf:"reference_data_source"`
}

type ApplicationSpecApplicationConfigurationVpcConfiguration struct {
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:MinItems=1
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +kubebuilder:validation:MaxItems=16
	// +kubebuilder:validation:MinItems=1
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	VpcConfigurationID *string `json:"vpcConfigurationID,omitempty" tf:"vpc_configuration_id"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type ApplicationSpecApplicationConfiguration struct {
	ApplicationCodeConfiguration *ApplicationSpecApplicationConfigurationApplicationCodeConfiguration `json:"applicationCodeConfiguration" tf:"application_code_configuration"`
	// +optional
	ApplicationSnapshotConfiguration *ApplicationSpecApplicationConfigurationApplicationSnapshotConfiguration `json:"applicationSnapshotConfiguration,omitempty" tf:"application_snapshot_configuration"`
	// +optional
	EnvironmentProperties *ApplicationSpecApplicationConfigurationEnvironmentProperties `json:"environmentProperties,omitempty" tf:"environment_properties"`
	// +optional
	FlinkApplicationConfiguration *ApplicationSpecApplicationConfigurationFlinkApplicationConfiguration `json:"flinkApplicationConfiguration,omitempty" tf:"flink_application_configuration"`
	// +optional
	RunConfiguration *ApplicationSpecApplicationConfigurationRunConfiguration `json:"runConfiguration,omitempty" tf:"run_configuration"`
	// +optional
	SqlApplicationConfiguration *ApplicationSpecApplicationConfigurationSqlApplicationConfiguration `json:"sqlApplicationConfiguration,omitempty" tf:"sql_application_configuration"`
	// +optional
	VpcConfiguration *ApplicationSpecApplicationConfigurationVpcConfiguration `json:"vpcConfiguration,omitempty" tf:"vpc_configuration"`
}

type ApplicationSpecCloudwatchLoggingOptions struct {
	// +optional
	CloudwatchLoggingOptionID *string `json:"cloudwatchLoggingOptionID,omitempty" tf:"cloudwatch_logging_option_id"`
	LogStreamArn              *string `json:"logStreamArn" tf:"log_stream_arn"`
}

type ApplicationSpec struct {
	State *ApplicationSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApplicationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApplicationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApplicationConfiguration *ApplicationSpecApplicationConfiguration `json:"applicationConfiguration,omitempty" tf:"application_configuration"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CloudwatchLoggingOptions *ApplicationSpecCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options"`
	// +optional
	CreateTimestamp *string `json:"createTimestamp,omitempty" tf:"create_timestamp"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ForceStop *bool `json:"forceStop,omitempty" tf:"force_stop"`
	// +optional
	LastUpdateTimestamp  *string `json:"lastUpdateTimestamp,omitempty" tf:"last_update_timestamp"`
	Name                 *string `json:"name" tf:"name"`
	RuntimeEnvironment   *string `json:"runtimeEnvironment" tf:"runtime_environment"`
	ServiceExecutionRole *string `json:"serviceExecutionRole" tf:"service_execution_role"`
	// +optional
	StartApplication *bool `json:"startApplication,omitempty" tf:"start_application"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	VersionID *int64 `json:"versionID,omitempty" tf:"version_id"`
}

type ApplicationStatus struct {
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

// ApplicationList is a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Application CRD objects
	Items []Application `json:"items,omitempty"`
}
