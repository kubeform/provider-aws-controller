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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecRestoreToPointInTime struct {
	// +optional
	RestoreTime *string `json:"restoreTime,omitempty" tf:"restore_time"`
	// +optional
	SourceDbInstanceIdentifier *string `json:"sourceDbInstanceIdentifier,omitempty" tf:"source_db_instance_identifier"`
	// +optional
	SourceDbiResourceID *string `json:"sourceDbiResourceID,omitempty" tf:"source_dbi_resource_id"`
	// +optional
	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time"`
}

type InstanceSpecS3Import struct {
	BucketName *string `json:"bucketName" tf:"bucket_name"`
	// +optional
	BucketPrefix        *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix"`
	IngestionRole       *string `json:"ingestionRole" tf:"ingestion_role"`
	SourceEngine        *string `json:"sourceEngine" tf:"source_engine"`
	SourceEngineVersion *string `json:"sourceEngineVersion" tf:"source_engine_version"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// +optional
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty" tf:"allocated_storage"`
	// +optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade"`
	// +optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period"`
	// +optional
	BackupWindow *string `json:"backupWindow,omitempty" tf:"backup_window"`
	// +optional
	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier"`
	// +optional
	CharacterSetName *string `json:"characterSetName,omitempty" tf:"character_set_name"`
	// +optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot"`
	// +optional
	CustomerOwnedIPEnabled *bool `json:"customerOwnedIPEnabled,omitempty" tf:"customer_owned_ip_enabled"`
	// +optional
	DbName *string `json:"dbName,omitempty" tf:"db_name"`
	// +optional
	DbSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name"`
	// +optional
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups,omitempty" tf:"delete_automated_backups"`
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// +optional
	DomainIamRoleName *string `json:"domainIamRoleName,omitempty" tf:"domain_iam_role_name"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports"`
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// +optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`
	// +optional
	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual"`
	// +optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier"`
	// +optional
	HostedZoneID *string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id"`
	// +optional
	IamDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled"`
	// +optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier"`
	// +optional
	IdentifierPrefix *string `json:"identifierPrefix,omitempty" tf:"identifier_prefix"`
	InstanceClass    *string `json:"instanceClass" tf:"instance_class"`
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	LatestRestorableTime *string `json:"latestRestorableTime,omitempty" tf:"latest_restorable_time"`
	// +optional
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model"`
	// +optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`
	// +optional
	MaxAllocatedStorage *int64 `json:"maxAllocatedStorage,omitempty" tf:"max_allocated_storage"`
	// +optional
	MonitoringInterval *int64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval"`
	// +optional
	MonitoringRoleArn *string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn"`
	// +optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az"`
	// +optional
	// Deprecated
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NcharCharacterSetName *string `json:"ncharCharacterSetName,omitempty" tf:"nchar_character_set_name"`
	// +optional
	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name"`
	// +optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled"`
	// +optional
	PerformanceInsightsKmsKeyID *string `json:"performanceInsightsKmsKeyID,omitempty" tf:"performance_insights_kms_key_id"`
	// +optional
	PerformanceInsightsRetentionPeriod *int64 `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible"`
	// +optional
	ReplicaMode *string `json:"replicaMode,omitempty" tf:"replica_mode"`
	// +optional
	Replicas []string `json:"replicas,omitempty" tf:"replicas"`
	// +optional
	ReplicateSourceDb *string `json:"replicateSourceDb,omitempty" tf:"replicate_source_db"`
	// +optional
	ResourceID *string `json:"resourceID,omitempty" tf:"resource_id"`
	// +optional
	RestoreToPointInTime *InstanceSpecRestoreToPointInTime `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time"`
	// +optional
	S3Import *InstanceSpecS3Import `json:"s3Import,omitempty" tf:"s3_import"`
	// +optional
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names"`
	// +optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot"`
	// +optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted"`
	// +optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
	// +optional
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
