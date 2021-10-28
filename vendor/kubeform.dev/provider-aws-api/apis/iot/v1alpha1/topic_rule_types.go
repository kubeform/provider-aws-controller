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

type TopicRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicRuleSpec   `json:"spec,omitempty"`
	Status            TopicRuleStatus `json:"status,omitempty"`
}

type TopicRuleSpecCloudwatchAlarm struct {
	AlarmName   *string `json:"alarmName" tf:"alarm_name"`
	RoleArn     *string `json:"roleArn" tf:"role_arn"`
	StateReason *string `json:"stateReason" tf:"state_reason"`
	StateValue  *string `json:"stateValue" tf:"state_value"`
}

type TopicRuleSpecCloudwatchMetric struct {
	MetricName      *string `json:"metricName" tf:"metric_name"`
	MetricNamespace *string `json:"metricNamespace" tf:"metric_namespace"`
	// +optional
	MetricTimestamp *string `json:"metricTimestamp,omitempty" tf:"metric_timestamp"`
	MetricUnit      *string `json:"metricUnit" tf:"metric_unit"`
	MetricValue     *string `json:"metricValue" tf:"metric_value"`
	RoleArn         *string `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecDynamodb struct {
	HashKeyField *string `json:"hashKeyField" tf:"hash_key_field"`
	// +optional
	HashKeyType  *string `json:"hashKeyType,omitempty" tf:"hash_key_type"`
	HashKeyValue *string `json:"hashKeyValue" tf:"hash_key_value"`
	// +optional
	Operation *string `json:"operation,omitempty" tf:"operation"`
	// +optional
	PayloadField *string `json:"payloadField,omitempty" tf:"payload_field"`
	// +optional
	RangeKeyField *string `json:"rangeKeyField,omitempty" tf:"range_key_field"`
	// +optional
	RangeKeyType *string `json:"rangeKeyType,omitempty" tf:"range_key_type"`
	// +optional
	RangeKeyValue *string `json:"rangeKeyValue,omitempty" tf:"range_key_value"`
	RoleArn       *string `json:"roleArn" tf:"role_arn"`
	TableName     *string `json:"tableName" tf:"table_name"`
}

type TopicRuleSpecDynamodbv2PutItem struct {
	TableName *string `json:"tableName" tf:"table_name"`
}

type TopicRuleSpecDynamodbv2 struct {
	// +optional
	PutItem *TopicRuleSpecDynamodbv2PutItem `json:"putItem,omitempty" tf:"put_item"`
	RoleArn *string                         `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecElasticsearch struct {
	Endpoint *string `json:"endpoint" tf:"endpoint"`
	ID       *string `json:"ID" tf:"id"`
	Index    *string `json:"index" tf:"index"`
	RoleArn  *string `json:"roleArn" tf:"role_arn"`
	Type     *string `json:"type" tf:"type"`
}

type TopicRuleSpecErrorActionCloudwatchAlarm struct {
	AlarmName   *string `json:"alarmName" tf:"alarm_name"`
	RoleArn     *string `json:"roleArn" tf:"role_arn"`
	StateReason *string `json:"stateReason" tf:"state_reason"`
	StateValue  *string `json:"stateValue" tf:"state_value"`
}

type TopicRuleSpecErrorActionCloudwatchMetric struct {
	MetricName      *string `json:"metricName" tf:"metric_name"`
	MetricNamespace *string `json:"metricNamespace" tf:"metric_namespace"`
	// +optional
	MetricTimestamp *string `json:"metricTimestamp,omitempty" tf:"metric_timestamp"`
	MetricUnit      *string `json:"metricUnit" tf:"metric_unit"`
	MetricValue     *string `json:"metricValue" tf:"metric_value"`
	RoleArn         *string `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecErrorActionDynamodb struct {
	HashKeyField *string `json:"hashKeyField" tf:"hash_key_field"`
	// +optional
	HashKeyType  *string `json:"hashKeyType,omitempty" tf:"hash_key_type"`
	HashKeyValue *string `json:"hashKeyValue" tf:"hash_key_value"`
	// +optional
	Operation *string `json:"operation,omitempty" tf:"operation"`
	// +optional
	PayloadField *string `json:"payloadField,omitempty" tf:"payload_field"`
	// +optional
	RangeKeyField *string `json:"rangeKeyField,omitempty" tf:"range_key_field"`
	// +optional
	RangeKeyType *string `json:"rangeKeyType,omitempty" tf:"range_key_type"`
	// +optional
	RangeKeyValue *string `json:"rangeKeyValue,omitempty" tf:"range_key_value"`
	RoleArn       *string `json:"roleArn" tf:"role_arn"`
	TableName     *string `json:"tableName" tf:"table_name"`
}

type TopicRuleSpecErrorActionDynamodbv2PutItem struct {
	TableName *string `json:"tableName" tf:"table_name"`
}

type TopicRuleSpecErrorActionDynamodbv2 struct {
	// +optional
	PutItem *TopicRuleSpecErrorActionDynamodbv2PutItem `json:"putItem,omitempty" tf:"put_item"`
	RoleArn *string                                    `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecErrorActionElasticsearch struct {
	Endpoint *string `json:"endpoint" tf:"endpoint"`
	ID       *string `json:"ID" tf:"id"`
	Index    *string `json:"index" tf:"index"`
	RoleArn  *string `json:"roleArn" tf:"role_arn"`
	Type     *string `json:"type" tf:"type"`
}

type TopicRuleSpecErrorActionFirehose struct {
	DeliveryStreamName *string `json:"deliveryStreamName" tf:"delivery_stream_name"`
	RoleArn            *string `json:"roleArn" tf:"role_arn"`
	// +optional
	Separator *string `json:"separator,omitempty" tf:"separator"`
}

type TopicRuleSpecErrorActionIotAnalytics struct {
	ChannelName *string `json:"channelName" tf:"channel_name"`
	RoleArn     *string `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecErrorActionIotEvents struct {
	InputName *string `json:"inputName" tf:"input_name"`
	// +optional
	MessageID *string `json:"messageID,omitempty" tf:"message_id"`
	RoleArn   *string `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecErrorActionKinesis struct {
	// +optional
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key"`
	RoleArn      *string `json:"roleArn" tf:"role_arn"`
	StreamName   *string `json:"streamName" tf:"stream_name"`
}

type TopicRuleSpecErrorActionLambda struct {
	FunctionArn *string `json:"functionArn" tf:"function_arn"`
}

type TopicRuleSpecErrorActionRepublish struct {
	// +optional
	Qos     *int64  `json:"qos,omitempty" tf:"qos"`
	RoleArn *string `json:"roleArn" tf:"role_arn"`
	Topic   *string `json:"topic" tf:"topic"`
}

type TopicRuleSpecErrorActionS3 struct {
	BucketName *string `json:"bucketName" tf:"bucket_name"`
	Key        *string `json:"key" tf:"key"`
	RoleArn    *string `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecErrorActionSns struct {
	// +optional
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format"`
	RoleArn       *string `json:"roleArn" tf:"role_arn"`
	TargetArn     *string `json:"targetArn" tf:"target_arn"`
}

type TopicRuleSpecErrorActionSqs struct {
	QueueURL  *string `json:"queueURL" tf:"queue_url"`
	RoleArn   *string `json:"roleArn" tf:"role_arn"`
	UseBase64 *bool   `json:"useBase64" tf:"use_base64"`
}

type TopicRuleSpecErrorActionStepFunctions struct {
	// +optional
	ExecutionNamePrefix *string `json:"executionNamePrefix,omitempty" tf:"execution_name_prefix"`
	RoleArn             *string `json:"roleArn" tf:"role_arn"`
	StateMachineName    *string `json:"stateMachineName" tf:"state_machine_name"`
}

type TopicRuleSpecErrorAction struct {
	// +optional
	CloudwatchAlarm *TopicRuleSpecErrorActionCloudwatchAlarm `json:"cloudwatchAlarm,omitempty" tf:"cloudwatch_alarm"`
	// +optional
	CloudwatchMetric *TopicRuleSpecErrorActionCloudwatchMetric `json:"cloudwatchMetric,omitempty" tf:"cloudwatch_metric"`
	// +optional
	Dynamodb *TopicRuleSpecErrorActionDynamodb `json:"dynamodb,omitempty" tf:"dynamodb"`
	// +optional
	Dynamodbv2 *TopicRuleSpecErrorActionDynamodbv2 `json:"dynamodbv2,omitempty" tf:"dynamodbv2"`
	// +optional
	Elasticsearch *TopicRuleSpecErrorActionElasticsearch `json:"elasticsearch,omitempty" tf:"elasticsearch"`
	// +optional
	Firehose *TopicRuleSpecErrorActionFirehose `json:"firehose,omitempty" tf:"firehose"`
	// +optional
	IotAnalytics *TopicRuleSpecErrorActionIotAnalytics `json:"iotAnalytics,omitempty" tf:"iot_analytics"`
	// +optional
	IotEvents *TopicRuleSpecErrorActionIotEvents `json:"iotEvents,omitempty" tf:"iot_events"`
	// +optional
	Kinesis *TopicRuleSpecErrorActionKinesis `json:"kinesis,omitempty" tf:"kinesis"`
	// +optional
	Lambda *TopicRuleSpecErrorActionLambda `json:"lambda,omitempty" tf:"lambda"`
	// +optional
	Republish *TopicRuleSpecErrorActionRepublish `json:"republish,omitempty" tf:"republish"`
	// +optional
	S3 *TopicRuleSpecErrorActionS3 `json:"s3,omitempty" tf:"s3"`
	// +optional
	Sns *TopicRuleSpecErrorActionSns `json:"sns,omitempty" tf:"sns"`
	// +optional
	Sqs *TopicRuleSpecErrorActionSqs `json:"sqs,omitempty" tf:"sqs"`
	// +optional
	StepFunctions *TopicRuleSpecErrorActionStepFunctions `json:"stepFunctions,omitempty" tf:"step_functions"`
}

type TopicRuleSpecFirehose struct {
	DeliveryStreamName *string `json:"deliveryStreamName" tf:"delivery_stream_name"`
	RoleArn            *string `json:"roleArn" tf:"role_arn"`
	// +optional
	Separator *string `json:"separator,omitempty" tf:"separator"`
}

type TopicRuleSpecIotAnalytics struct {
	ChannelName *string `json:"channelName" tf:"channel_name"`
	RoleArn     *string `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecIotEvents struct {
	InputName *string `json:"inputName" tf:"input_name"`
	// +optional
	MessageID *string `json:"messageID,omitempty" tf:"message_id"`
	RoleArn   *string `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecKinesis struct {
	// +optional
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key"`
	RoleArn      *string `json:"roleArn" tf:"role_arn"`
	StreamName   *string `json:"streamName" tf:"stream_name"`
}

type TopicRuleSpecLambda struct {
	FunctionArn *string `json:"functionArn" tf:"function_arn"`
}

type TopicRuleSpecRepublish struct {
	// +optional
	Qos     *int64  `json:"qos,omitempty" tf:"qos"`
	RoleArn *string `json:"roleArn" tf:"role_arn"`
	Topic   *string `json:"topic" tf:"topic"`
}

type TopicRuleSpecS3 struct {
	BucketName *string `json:"bucketName" tf:"bucket_name"`
	Key        *string `json:"key" tf:"key"`
	RoleArn    *string `json:"roleArn" tf:"role_arn"`
}

type TopicRuleSpecSns struct {
	// +optional
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format"`
	RoleArn       *string `json:"roleArn" tf:"role_arn"`
	TargetArn     *string `json:"targetArn" tf:"target_arn"`
}

type TopicRuleSpecSqs struct {
	QueueURL  *string `json:"queueURL" tf:"queue_url"`
	RoleArn   *string `json:"roleArn" tf:"role_arn"`
	UseBase64 *bool   `json:"useBase64" tf:"use_base64"`
}

type TopicRuleSpecStepFunctions struct {
	// +optional
	ExecutionNamePrefix *string `json:"executionNamePrefix,omitempty" tf:"execution_name_prefix"`
	RoleArn             *string `json:"roleArn" tf:"role_arn"`
	StateMachineName    *string `json:"stateMachineName" tf:"state_machine_name"`
}

type TopicRuleSpec struct {
	State *TopicRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource TopicRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TopicRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CloudwatchAlarm []TopicRuleSpecCloudwatchAlarm `json:"cloudwatchAlarm,omitempty" tf:"cloudwatch_alarm"`
	// +optional
	CloudwatchMetric []TopicRuleSpecCloudwatchMetric `json:"cloudwatchMetric,omitempty" tf:"cloudwatch_metric"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Dynamodb []TopicRuleSpecDynamodb `json:"dynamodb,omitempty" tf:"dynamodb"`
	// +optional
	Dynamodbv2 []TopicRuleSpecDynamodbv2 `json:"dynamodbv2,omitempty" tf:"dynamodbv2"`
	// +optional
	Elasticsearch []TopicRuleSpecElasticsearch `json:"elasticsearch,omitempty" tf:"elasticsearch"`
	Enabled       *bool                        `json:"enabled" tf:"enabled"`
	// +optional
	ErrorAction *TopicRuleSpecErrorAction `json:"errorAction,omitempty" tf:"error_action"`
	// +optional
	Firehose []TopicRuleSpecFirehose `json:"firehose,omitempty" tf:"firehose"`
	// +optional
	IotAnalytics []TopicRuleSpecIotAnalytics `json:"iotAnalytics,omitempty" tf:"iot_analytics"`
	// +optional
	IotEvents []TopicRuleSpecIotEvents `json:"iotEvents,omitempty" tf:"iot_events"`
	// +optional
	Kinesis []TopicRuleSpecKinesis `json:"kinesis,omitempty" tf:"kinesis"`
	// +optional
	Lambda []TopicRuleSpecLambda `json:"lambda,omitempty" tf:"lambda"`
	Name   *string               `json:"name" tf:"name"`
	// +optional
	Republish []TopicRuleSpecRepublish `json:"republish,omitempty" tf:"republish"`
	// +optional
	S3 []TopicRuleSpecS3 `json:"s3,omitempty" tf:"s3"`
	// +optional
	Sns        []TopicRuleSpecSns `json:"sns,omitempty" tf:"sns"`
	Sql        *string            `json:"sql" tf:"sql"`
	SqlVersion *string            `json:"sqlVersion" tf:"sql_version"`
	// +optional
	Sqs []TopicRuleSpecSqs `json:"sqs,omitempty" tf:"sqs"`
	// +optional
	StepFunctions []TopicRuleSpecStepFunctions `json:"stepFunctions,omitempty" tf:"step_functions"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type TopicRuleStatus struct {
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

// TopicRuleList is a list of TopicRules
type TopicRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TopicRule CRD objects
	Items []TopicRule `json:"items,omitempty"`
}
