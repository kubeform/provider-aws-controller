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

type Template struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TemplateSpec   `json:"spec,omitempty"`
	Status            TemplateStatus `json:"status,omitempty"`
}

type TemplateSpecBlockDeviceMappingsEbs struct {
	// +optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`
	// +optional
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
	// +optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`
	// +optional
	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`
	// +optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type TemplateSpecBlockDeviceMappings struct {
	// +optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`
	// +optional
	Ebs *TemplateSpecBlockDeviceMappingsEbs `json:"ebs,omitempty" tf:"ebs"`
	// +optional
	NoDevice *string `json:"noDevice,omitempty" tf:"no_device"`
	// +optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name"`
}

type TemplateSpecCapacityReservationSpecificationCapacityReservationTarget struct {
	// +optional
	CapacityReservationID *string `json:"capacityReservationID,omitempty" tf:"capacity_reservation_id"`
}

type TemplateSpecCapacityReservationSpecification struct {
	// +optional
	CapacityReservationPreference *string `json:"capacityReservationPreference,omitempty" tf:"capacity_reservation_preference"`
	// +optional
	CapacityReservationTarget *TemplateSpecCapacityReservationSpecificationCapacityReservationTarget `json:"capacityReservationTarget,omitempty" tf:"capacity_reservation_target"`
}

type TemplateSpecCpuOptions struct {
	// +optional
	CoreCount *int64 `json:"coreCount,omitempty" tf:"core_count"`
	// +optional
	ThreadsPerCore *int64 `json:"threadsPerCore,omitempty" tf:"threads_per_core"`
}

type TemplateSpecCreditSpecification struct {
	// +optional
	CpuCredits *string `json:"cpuCredits,omitempty" tf:"cpu_credits"`
}

type TemplateSpecElasticGpuSpecifications struct {
	Type *string `json:"type" tf:"type"`
}

type TemplateSpecElasticInferenceAccelerator struct {
	Type *string `json:"type" tf:"type"`
}

type TemplateSpecEnclaveOptions struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type TemplateSpecHibernationOptions struct {
	Configured *bool `json:"configured" tf:"configured"`
}

type TemplateSpecIamInstanceProfile struct {
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type TemplateSpecInstanceMarketOptionsSpotOptions struct {
	// +optional
	BlockDurationMinutes *int64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`
	// +optional
	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior"`
	// +optional
	MaxPrice *string `json:"maxPrice,omitempty" tf:"max_price"`
	// +optional
	SpotInstanceType *string `json:"spotInstanceType,omitempty" tf:"spot_instance_type"`
	// +optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until"`
}

type TemplateSpecInstanceMarketOptions struct {
	// +optional
	MarketType *string `json:"marketType,omitempty" tf:"market_type"`
	// +optional
	SpotOptions *TemplateSpecInstanceMarketOptionsSpotOptions `json:"spotOptions,omitempty" tf:"spot_options"`
}

type TemplateSpecLicenseSpecification struct {
	LicenseConfigurationArn *string `json:"licenseConfigurationArn" tf:"license_configuration_arn"`
}

type TemplateSpecMetadataOptions struct {
	// +optional
	HttpEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint"`
	// +optional
	HttpPutResponseHopLimit *int64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit"`
	// +optional
	HttpTokens *string `json:"httpTokens,omitempty" tf:"http_tokens"`
}

type TemplateSpecMonitoring struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type TemplateSpecNetworkInterfaces struct {
	// +optional
	AssociateCarrierIPAddress *string `json:"associateCarrierIPAddress,omitempty" tf:"associate_carrier_ip_address"`
	// +optional
	AssociatePublicIPAddress *string `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address"`
	// +optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DeviceIndex *int64 `json:"deviceIndex,omitempty" tf:"device_index"`
	// +optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type"`
	// +optional
	Ipv4AddressCount *int64 `json:"ipv4AddressCount,omitempty" tf:"ipv4_address_count"`
	// +optional
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty" tf:"ipv4_addresses"`
	// +optional
	Ipv6AddressCount *int64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count"`
	// +optional
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses"`
	// +optional
	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty" tf:"network_interface_id"`
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
}

type TemplateSpecPlacement struct {
	// +optional
	Affinity *string `json:"affinity,omitempty" tf:"affinity"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name"`
	// +optional
	HostID *string `json:"hostID,omitempty" tf:"host_id"`
	// +optional
	HostResourceGroupArn *string `json:"hostResourceGroupArn,omitempty" tf:"host_resource_group_arn"`
	// +optional
	PartitionNumber *int64 `json:"partitionNumber,omitempty" tf:"partition_number"`
	// +optional
	SpreadDomain *string `json:"spreadDomain,omitempty" tf:"spread_domain"`
	// +optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy"`
}

type TemplateSpecTagSpecifications struct {
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type TemplateSpec struct {
	State *TemplateSpecResource `json:"state,omitempty" tf:"-"`

	Resource TemplateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TemplateSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	BlockDeviceMappings []TemplateSpecBlockDeviceMappings `json:"blockDeviceMappings,omitempty" tf:"block_device_mappings"`
	// +optional
	CapacityReservationSpecification *TemplateSpecCapacityReservationSpecification `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification"`
	// +optional
	CpuOptions *TemplateSpecCpuOptions `json:"cpuOptions,omitempty" tf:"cpu_options"`
	// +optional
	CreditSpecification *TemplateSpecCreditSpecification `json:"creditSpecification,omitempty" tf:"credit_specification"`
	// +optional
	DefaultVersion *int64 `json:"defaultVersion,omitempty" tf:"default_version"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisableAPITermination *bool `json:"disableAPITermination,omitempty" tf:"disable_api_termination"`
	// +optional
	EbsOptimized *string `json:"ebsOptimized,omitempty" tf:"ebs_optimized"`
	// +optional
	ElasticGpuSpecifications []TemplateSpecElasticGpuSpecifications `json:"elasticGpuSpecifications,omitempty" tf:"elastic_gpu_specifications"`
	// +optional
	ElasticInferenceAccelerator *TemplateSpecElasticInferenceAccelerator `json:"elasticInferenceAccelerator,omitempty" tf:"elastic_inference_accelerator"`
	// +optional
	EnclaveOptions *TemplateSpecEnclaveOptions `json:"enclaveOptions,omitempty" tf:"enclave_options"`
	// +optional
	HibernationOptions *TemplateSpecHibernationOptions `json:"hibernationOptions,omitempty" tf:"hibernation_options"`
	// +optional
	IamInstanceProfile *TemplateSpecIamInstanceProfile `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile"`
	// +optional
	ImageID *string `json:"imageID,omitempty" tf:"image_id"`
	// +optional
	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior"`
	// +optional
	InstanceMarketOptions *TemplateSpecInstanceMarketOptions `json:"instanceMarketOptions,omitempty" tf:"instance_market_options"`
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	KernelID *string `json:"kernelID,omitempty" tf:"kernel_id"`
	// +optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name"`
	// +optional
	LatestVersion *int64 `json:"latestVersion,omitempty" tf:"latest_version"`
	// +optional
	LicenseSpecification []TemplateSpecLicenseSpecification `json:"licenseSpecification,omitempty" tf:"license_specification"`
	// +optional
	MetadataOptions *TemplateSpecMetadataOptions `json:"metadataOptions,omitempty" tf:"metadata_options"`
	// +optional
	Monitoring *TemplateSpecMonitoring `json:"monitoring,omitempty" tf:"monitoring"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	NetworkInterfaces []TemplateSpecNetworkInterfaces `json:"networkInterfaces,omitempty" tf:"network_interfaces"`
	// +optional
	Placement *TemplateSpecPlacement `json:"placement,omitempty" tf:"placement"`
	// +optional
	RamDiskID *string `json:"ramDiskID,omitempty" tf:"ram_disk_id"`
	// +optional
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names"`
	// +optional
	TagSpecifications []TemplateSpecTagSpecifications `json:"tagSpecifications,omitempty" tf:"tag_specifications"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	UpdateDefaultVersion *bool `json:"updateDefaultVersion,omitempty" tf:"update_default_version"`
	// +optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
	// +optional
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids"`
}

type TemplateStatus struct {
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

// TemplateList is a list of Templates
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Template CRD objects
	Items []Template `json:"items,omitempty"`
}
