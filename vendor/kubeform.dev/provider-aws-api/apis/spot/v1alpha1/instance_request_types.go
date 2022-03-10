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

type InstanceRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceRequestSpec   `json:"spec,omitempty"`
	Status            InstanceRequestStatus `json:"status,omitempty"`
}

type InstanceRequestSpecCapacityReservationSpecificationCapacityReservationTarget struct {
	// +optional
	CapacityReservationID *string `json:"capacityReservationID,omitempty" tf:"capacity_reservation_id"`
}

type InstanceRequestSpecCapacityReservationSpecification struct {
	// +optional
	CapacityReservationPreference *string `json:"capacityReservationPreference,omitempty" tf:"capacity_reservation_preference"`
	// +optional
	CapacityReservationTarget *InstanceRequestSpecCapacityReservationSpecificationCapacityReservationTarget `json:"capacityReservationTarget,omitempty" tf:"capacity_reservation_target"`
}

type InstanceRequestSpecCreditSpecification struct {
	// +optional
	CpuCredits *string `json:"cpuCredits,omitempty" tf:"cpu_credits"`
}

type InstanceRequestSpecEbsBlockDevice struct {
	// +optional
	DeleteOnTermination *bool   `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`
	DeviceName          *string `json:"deviceName" tf:"device_name"`
	// +optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`
	// +optional
	VolumeID *string `json:"volumeID,omitempty" tf:"volume_id"`
	// +optional
	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`
	// +optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type InstanceRequestSpecEnclaveOptions struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type InstanceRequestSpecEphemeralBlockDevice struct {
	DeviceName *string `json:"deviceName" tf:"device_name"`
	// +optional
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device"`
	// +optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name"`
}

type InstanceRequestSpecLaunchTemplate struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type InstanceRequestSpecMetadataOptions struct {
	// +optional
	HttpEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint"`
	// +optional
	HttpPutResponseHopLimit *int64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit"`
	// +optional
	HttpTokens *string `json:"httpTokens,omitempty" tf:"http_tokens"`
	// +optional
	InstanceMetadataTags *string `json:"instanceMetadataTags,omitempty" tf:"instance_metadata_tags"`
}

type InstanceRequestSpecNetworkInterface struct {
	// +optional
	DeleteOnTermination *bool   `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`
	DeviceIndex         *int64  `json:"deviceIndex" tf:"device_index"`
	NetworkInterfaceID  *string `json:"networkInterfaceID" tf:"network_interface_id"`
}

type InstanceRequestSpecRootBlockDevice struct {
	// +optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`
	// +optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`
	// +optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`
	// +optional
	VolumeID *string `json:"volumeID,omitempty" tf:"volume_id"`
	// +optional
	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`
	// +optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type InstanceRequestSpec struct {
	State *InstanceRequestSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceRequestSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceRequestSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Ami *string `json:"ami,omitempty" tf:"ami"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AssociatePublicIPAddress *bool `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	BlockDurationMinutes *int64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`
	// +optional
	CapacityReservationSpecification *InstanceRequestSpecCapacityReservationSpecification `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification"`
	// +optional
	CpuCoreCount *int64 `json:"cpuCoreCount,omitempty" tf:"cpu_core_count"`
	// +optional
	CpuThreadsPerCore *int64 `json:"cpuThreadsPerCore,omitempty" tf:"cpu_threads_per_core"`
	// +optional
	CreditSpecification *InstanceRequestSpecCreditSpecification `json:"creditSpecification,omitempty" tf:"credit_specification"`
	// +optional
	DisableAPITermination *bool `json:"disableAPITermination,omitempty" tf:"disable_api_termination"`
	// +optional
	EbsBlockDevice []InstanceRequestSpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device"`
	// +optional
	EbsOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized"`
	// +optional
	EnclaveOptions *InstanceRequestSpecEnclaveOptions `json:"enclaveOptions,omitempty" tf:"enclave_options"`
	// +optional
	EphemeralBlockDevice []InstanceRequestSpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device"`
	// +optional
	GetPasswordData *bool `json:"getPasswordData,omitempty" tf:"get_password_data"`
	// +optional
	Hibernation *bool `json:"hibernation,omitempty" tf:"hibernation"`
	// +optional
	HostID *string `json:"hostID,omitempty" tf:"host_id"`
	// +optional
	IamInstanceProfile *string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile"`
	// +optional
	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior"`
	// +optional
	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior"`
	// +optional
	InstanceState *string `json:"instanceState,omitempty" tf:"instance_state"`
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	Ipv6AddressCount *int64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count"`
	// +optional
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses"`
	// +optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name"`
	// +optional
	LaunchGroup *string `json:"launchGroup,omitempty" tf:"launch_group"`
	// +optional
	LaunchTemplate *InstanceRequestSpecLaunchTemplate `json:"launchTemplate,omitempty" tf:"launch_template"`
	// +optional
	MetadataOptions *InstanceRequestSpecMetadataOptions `json:"metadataOptions,omitempty" tf:"metadata_options"`
	// +optional
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring"`
	// +optional
	NetworkInterface []InstanceRequestSpecNetworkInterface `json:"networkInterface,omitempty" tf:"network_interface"`
	// +optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn"`
	// +optional
	PasswordData *string `json:"passwordData,omitempty" tf:"password_data"`
	// +optional
	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group"`
	// +optional
	PlacementPartitionNumber *int64 `json:"placementPartitionNumber,omitempty" tf:"placement_partition_number"`
	// +optional
	PrimaryNetworkInterfaceID *string `json:"primaryNetworkInterfaceID,omitempty" tf:"primary_network_interface_id"`
	// +optional
	PrivateDNS *string `json:"privateDNS,omitempty" tf:"private_dns"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	PublicDNS *string `json:"publicDNS,omitempty" tf:"public_dns"`
	// +optional
	PublicIP *string `json:"publicIP,omitempty" tf:"public_ip"`
	// +optional
	RootBlockDevice *InstanceRequestSpecRootBlockDevice `json:"rootBlockDevice,omitempty" tf:"root_block_device"`
	// +optional
	SecondaryPrivateIPS []string `json:"secondaryPrivateIPS,omitempty" tf:"secondary_private_ips"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	// +optional
	SourceDestCheck *bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check"`
	// +optional
	SpotBidStatus *string `json:"spotBidStatus,omitempty" tf:"spot_bid_status"`
	// +optional
	SpotInstanceID *string `json:"spotInstanceID,omitempty" tf:"spot_instance_id"`
	// +optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price"`
	// +optional
	SpotRequestState *string `json:"spotRequestState,omitempty" tf:"spot_request_state"`
	// +optional
	SpotType *string `json:"spotType,omitempty" tf:"spot_type"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy"`
	// +optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
	// +optional
	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64"`
	// +optional
	ValidFrom *string `json:"validFrom,omitempty" tf:"valid_from"`
	// +optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until"`
	// +optional
	VolumeTags *map[string]string `json:"volumeTags,omitempty" tf:"volume_tags"`
	// +optional
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids"`
	// +optional
	WaitForFulfillment *bool `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment"`
}

type InstanceRequestStatus struct {
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

// InstanceRequestList is a list of InstanceRequests
type InstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InstanceRequest CRD objects
	Items []InstanceRequest `json:"items,omitempty"`
}
