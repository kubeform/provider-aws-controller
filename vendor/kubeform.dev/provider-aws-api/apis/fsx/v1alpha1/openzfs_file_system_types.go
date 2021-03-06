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

type OpenzfsFileSystem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpenzfsFileSystemSpec   `json:"spec,omitempty"`
	Status            OpenzfsFileSystemStatus `json:"status,omitempty"`
}

type OpenzfsFileSystemSpecDiskIopsConfiguration struct {
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
}

type OpenzfsFileSystemSpecRootVolumeConfigurationNfsExportsClientConfigurations struct {
	Clients *string `json:"clients" tf:"clients"`
	// +kubebuilder:validation:MaxItems=20
	// +kubebuilder:validation:MinItems=1
	Options []string `json:"options" tf:"options"`
}

type OpenzfsFileSystemSpecRootVolumeConfigurationNfsExports struct {
	// +kubebuilder:validation:MaxItems=25
	ClientConfigurations []OpenzfsFileSystemSpecRootVolumeConfigurationNfsExportsClientConfigurations `json:"clientConfigurations" tf:"client_configurations"`
}

type OpenzfsFileSystemSpecRootVolumeConfigurationUserAndGroupQuotas struct {
	ID                      *int64  `json:"ID" tf:"id"`
	StorageCapacityQuotaGib *int64  `json:"storageCapacityQuotaGib" tf:"storage_capacity_quota_gib"`
	Type                    *string `json:"type" tf:"type"`
}

type OpenzfsFileSystemSpecRootVolumeConfiguration struct {
	// +optional
	CopyTagsToSnapshots *bool `json:"copyTagsToSnapshots,omitempty" tf:"copy_tags_to_snapshots"`
	// +optional
	DataCompressionType *string `json:"dataCompressionType,omitempty" tf:"data_compression_type"`
	// +optional
	NfsExports *OpenzfsFileSystemSpecRootVolumeConfigurationNfsExports `json:"nfsExports,omitempty" tf:"nfs_exports"`
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	UserAndGroupQuotas []OpenzfsFileSystemSpecRootVolumeConfigurationUserAndGroupQuotas `json:"userAndGroupQuotas,omitempty" tf:"user_and_group_quotas"`
}

type OpenzfsFileSystemSpec struct {
	State *OpenzfsFileSystemSpecResource `json:"state,omitempty" tf:"-"`

	Resource OpenzfsFileSystemSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type OpenzfsFileSystemSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AutomaticBackupRetentionDays *int64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days"`
	// +optional
	BackupID *string `json:"backupID,omitempty" tf:"backup_id"`
	// +optional
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups"`
	// +optional
	CopyTagsToVolumes *bool `json:"copyTagsToVolumes,omitempty" tf:"copy_tags_to_volumes"`
	// +optional
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time"`
	DeploymentType                *string `json:"deploymentType" tf:"deployment_type"`
	// +optional
	DiskIopsConfiguration *OpenzfsFileSystemSpecDiskIopsConfiguration `json:"diskIopsConfiguration,omitempty" tf:"disk_iops_configuration"`
	// +optional
	DnsName *string `json:"dnsName,omitempty" tf:"dns_name"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	NetworkInterfaceIDS []string `json:"networkInterfaceIDS,omitempty" tf:"network_interface_ids"`
	// +optional
	OwnerID *string `json:"ownerID,omitempty" tf:"owner_id"`
	// +optional
	RootVolumeConfiguration *OpenzfsFileSystemSpecRootVolumeConfiguration `json:"rootVolumeConfiguration,omitempty" tf:"root_volume_configuration"`
	// +optional
	RootVolumeID *string `json:"rootVolumeID,omitempty" tf:"root_volume_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids"`
	// +optional
	StorageCapacity *int64 `json:"storageCapacity,omitempty" tf:"storage_capacity"`
	// +optional
	StorageType *string  `json:"storageType,omitempty" tf:"storage_type"`
	SubnetIDS   []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll            *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	ThroughputCapacity *int64             `json:"throughputCapacity" tf:"throughput_capacity"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
	// +optional
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time"`
}

type OpenzfsFileSystemStatus struct {
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

// OpenzfsFileSystemList is a list of OpenzfsFileSystems
type OpenzfsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpenzfsFileSystem CRD objects
	Items []OpenzfsFileSystem `json:"items,omitempty"`
}
