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

type ClusterSpecAutoTerminationPolicy struct {
	// +optional
	IdleTimeout *int64 `json:"idleTimeout,omitempty" tf:"idle_timeout"`
}

type ClusterSpecBootstrapAction struct {
	// +optional
	Args []string `json:"args,omitempty" tf:"args"`
	Name *string  `json:"name" tf:"name"`
	Path *string  `json:"path" tf:"path"`
}

type ClusterSpecCoreInstanceFleetInstanceTypeConfigsConfigurations struct {
	// +optional
	Classification *string `json:"classification,omitempty" tf:"classification"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
}

type ClusterSpecCoreInstanceFleetInstanceTypeConfigsEbsConfig struct {
	// +optional
	Iops *int64  `json:"iops,omitempty" tf:"iops"`
	Size *int64  `json:"size" tf:"size"`
	Type *string `json:"type" tf:"type"`
	// +optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type ClusterSpecCoreInstanceFleetInstanceTypeConfigs struct {
	// +optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`
	// +optional
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice,omitempty" tf:"bid_price_as_percentage_of_on_demand_price"`
	// +optional
	Configurations []ClusterSpecCoreInstanceFleetInstanceTypeConfigsConfigurations `json:"configurations,omitempty" tf:"configurations"`
	// +optional
	EbsConfig    []ClusterSpecCoreInstanceFleetInstanceTypeConfigsEbsConfig `json:"ebsConfig,omitempty" tf:"ebs_config"`
	InstanceType *string                                                    `json:"instanceType" tf:"instance_type"`
	// +optional
	WeightedCapacity *int64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type ClusterSpecCoreInstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	AllocationStrategy *string `json:"allocationStrategy" tf:"allocation_strategy"`
}

type ClusterSpecCoreInstanceFleetLaunchSpecificationsSpotSpecification struct {
	AllocationStrategy *string `json:"allocationStrategy" tf:"allocation_strategy"`
	// +optional
	BlockDurationMinutes   *int64  `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`
	TimeoutAction          *string `json:"timeoutAction" tf:"timeout_action"`
	TimeoutDurationMinutes *int64  `json:"timeoutDurationMinutes" tf:"timeout_duration_minutes"`
}

type ClusterSpecCoreInstanceFleetLaunchSpecifications struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	OnDemandSpecification []ClusterSpecCoreInstanceFleetLaunchSpecificationsOnDemandSpecification `json:"onDemandSpecification,omitempty" tf:"on_demand_specification"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	SpotSpecification []ClusterSpecCoreInstanceFleetLaunchSpecificationsSpotSpecification `json:"spotSpecification,omitempty" tf:"spot_specification"`
}

type ClusterSpecCoreInstanceFleet struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InstanceTypeConfigs []ClusterSpecCoreInstanceFleetInstanceTypeConfigs `json:"instanceTypeConfigs,omitempty" tf:"instance_type_configs"`
	// +optional
	LaunchSpecifications *ClusterSpecCoreInstanceFleetLaunchSpecifications `json:"launchSpecifications,omitempty" tf:"launch_specifications"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ProvisionedOnDemandCapacity *int64 `json:"provisionedOnDemandCapacity,omitempty" tf:"provisioned_on_demand_capacity"`
	// +optional
	ProvisionedSpotCapacity *int64 `json:"provisionedSpotCapacity,omitempty" tf:"provisioned_spot_capacity"`
	// +optional
	TargetOnDemandCapacity *int64 `json:"targetOnDemandCapacity,omitempty" tf:"target_on_demand_capacity"`
	// +optional
	TargetSpotCapacity *int64 `json:"targetSpotCapacity,omitempty" tf:"target_spot_capacity"`
}

type ClusterSpecCoreInstanceGroupEbsConfig struct {
	// +optional
	Iops *int64  `json:"iops,omitempty" tf:"iops"`
	Size *int64  `json:"size" tf:"size"`
	Type *string `json:"type" tf:"type"`
	// +optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type ClusterSpecCoreInstanceGroup struct {
	// +optional
	AutoscalingPolicy *string `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy"`
	// +optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`
	// +optional
	EbsConfig []ClusterSpecCoreInstanceGroupEbsConfig `json:"ebsConfig,omitempty" tf:"ebs_config"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InstanceCount *int64  `json:"instanceCount,omitempty" tf:"instance_count"`
	InstanceType  *string `json:"instanceType" tf:"instance_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type ClusterSpecEc2Attributes struct {
	// +optional
	AdditionalMasterSecurityGroups *string `json:"additionalMasterSecurityGroups,omitempty" tf:"additional_master_security_groups"`
	// +optional
	AdditionalSlaveSecurityGroups *string `json:"additionalSlaveSecurityGroups,omitempty" tf:"additional_slave_security_groups"`
	// +optional
	EmrManagedMasterSecurityGroup *string `json:"emrManagedMasterSecurityGroup,omitempty" tf:"emr_managed_master_security_group"`
	// +optional
	EmrManagedSlaveSecurityGroup *string `json:"emrManagedSlaveSecurityGroup,omitempty" tf:"emr_managed_slave_security_group"`
	InstanceProfile              *string `json:"instanceProfile" tf:"instance_profile"`
	// +optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name"`
	// +optional
	ServiceAccessSecurityGroup *string `json:"serviceAccessSecurityGroup,omitempty" tf:"service_access_security_group"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids"`
}

type ClusterSpecKerberosAttributes struct {
	// +optional
	AdDomainJoinPassword *string `json:"-" sensitive:"true" tf:"ad_domain_join_password"`
	// +optional
	AdDomainJoinUser *string `json:"adDomainJoinUser,omitempty" tf:"ad_domain_join_user"`
	// +optional
	CrossRealmTrustPrincipalPassword *string `json:"-" sensitive:"true" tf:"cross_realm_trust_principal_password"`
	KdcAdminPassword                 *string `json:"-" sensitive:"true" tf:"kdc_admin_password"`
	Realm                            *string `json:"realm" tf:"realm"`
}

type ClusterSpecMasterInstanceFleetInstanceTypeConfigsConfigurations struct {
	// +optional
	Classification *string `json:"classification,omitempty" tf:"classification"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
}

type ClusterSpecMasterInstanceFleetInstanceTypeConfigsEbsConfig struct {
	// +optional
	Iops *int64  `json:"iops,omitempty" tf:"iops"`
	Size *int64  `json:"size" tf:"size"`
	Type *string `json:"type" tf:"type"`
	// +optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type ClusterSpecMasterInstanceFleetInstanceTypeConfigs struct {
	// +optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`
	// +optional
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice,omitempty" tf:"bid_price_as_percentage_of_on_demand_price"`
	// +optional
	Configurations []ClusterSpecMasterInstanceFleetInstanceTypeConfigsConfigurations `json:"configurations,omitempty" tf:"configurations"`
	// +optional
	EbsConfig    []ClusterSpecMasterInstanceFleetInstanceTypeConfigsEbsConfig `json:"ebsConfig,omitempty" tf:"ebs_config"`
	InstanceType *string                                                      `json:"instanceType" tf:"instance_type"`
	// +optional
	WeightedCapacity *int64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type ClusterSpecMasterInstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	AllocationStrategy *string `json:"allocationStrategy" tf:"allocation_strategy"`
}

type ClusterSpecMasterInstanceFleetLaunchSpecificationsSpotSpecification struct {
	AllocationStrategy *string `json:"allocationStrategy" tf:"allocation_strategy"`
	// +optional
	BlockDurationMinutes   *int64  `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`
	TimeoutAction          *string `json:"timeoutAction" tf:"timeout_action"`
	TimeoutDurationMinutes *int64  `json:"timeoutDurationMinutes" tf:"timeout_duration_minutes"`
}

type ClusterSpecMasterInstanceFleetLaunchSpecifications struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	OnDemandSpecification []ClusterSpecMasterInstanceFleetLaunchSpecificationsOnDemandSpecification `json:"onDemandSpecification,omitempty" tf:"on_demand_specification"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	SpotSpecification []ClusterSpecMasterInstanceFleetLaunchSpecificationsSpotSpecification `json:"spotSpecification,omitempty" tf:"spot_specification"`
}

type ClusterSpecMasterInstanceFleet struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InstanceTypeConfigs []ClusterSpecMasterInstanceFleetInstanceTypeConfigs `json:"instanceTypeConfigs,omitempty" tf:"instance_type_configs"`
	// +optional
	LaunchSpecifications *ClusterSpecMasterInstanceFleetLaunchSpecifications `json:"launchSpecifications,omitempty" tf:"launch_specifications"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ProvisionedOnDemandCapacity *int64 `json:"provisionedOnDemandCapacity,omitempty" tf:"provisioned_on_demand_capacity"`
	// +optional
	ProvisionedSpotCapacity *int64 `json:"provisionedSpotCapacity,omitempty" tf:"provisioned_spot_capacity"`
	// +optional
	TargetOnDemandCapacity *int64 `json:"targetOnDemandCapacity,omitempty" tf:"target_on_demand_capacity"`
	// +optional
	TargetSpotCapacity *int64 `json:"targetSpotCapacity,omitempty" tf:"target_spot_capacity"`
}

type ClusterSpecMasterInstanceGroupEbsConfig struct {
	// +optional
	Iops *int64  `json:"iops,omitempty" tf:"iops"`
	Size *int64  `json:"size" tf:"size"`
	Type *string `json:"type" tf:"type"`
	// +optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type ClusterSpecMasterInstanceGroup struct {
	// +optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`
	// +optional
	EbsConfig []ClusterSpecMasterInstanceGroupEbsConfig `json:"ebsConfig,omitempty" tf:"ebs_config"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InstanceCount *int64  `json:"instanceCount,omitempty" tf:"instance_count"`
	InstanceType  *string `json:"instanceType" tf:"instance_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type ClusterSpecStepHadoopJarStep struct {
	// +optional
	Args []string `json:"args,omitempty" tf:"args"`
	Jar  *string  `json:"jar" tf:"jar"`
	// +optional
	MainClass *string `json:"mainClass,omitempty" tf:"main_class"`
	// +optional
	Properties *map[string]string `json:"properties,omitempty" tf:"properties"`
}

type ClusterSpecStep struct {
	ActionOnFailure *string                       `json:"actionOnFailure" tf:"action_on_failure"`
	HadoopJarStep   *ClusterSpecStepHadoopJarStep `json:"hadoopJarStep" tf:"hadoop_jar_step"`
	Name            *string                       `json:"name" tf:"name"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalInfo *string `json:"additionalInfo,omitempty" tf:"additional_info"`
	// +optional
	Applications []string `json:"applications,omitempty" tf:"applications"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AutoTerminationPolicy *ClusterSpecAutoTerminationPolicy `json:"autoTerminationPolicy,omitempty" tf:"auto_termination_policy"`
	// +optional
	AutoscalingRole *string `json:"autoscalingRole,omitempty" tf:"autoscaling_role"`
	// +optional
	BootstrapAction []ClusterSpecBootstrapAction `json:"bootstrapAction,omitempty" tf:"bootstrap_action"`
	// +optional
	ClusterState *string `json:"clusterState,omitempty" tf:"cluster_state"`
	// +optional
	Configurations *string `json:"configurations,omitempty" tf:"configurations"`
	// +optional
	ConfigurationsJSON *string `json:"configurationsJSON,omitempty" tf:"configurations_json"`
	// +optional
	CoreInstanceFleet *ClusterSpecCoreInstanceFleet `json:"coreInstanceFleet,omitempty" tf:"core_instance_fleet"`
	// +optional
	CoreInstanceGroup *ClusterSpecCoreInstanceGroup `json:"coreInstanceGroup,omitempty" tf:"core_instance_group"`
	// +optional
	CustomAmiID *string `json:"customAmiID,omitempty" tf:"custom_ami_id"`
	// +optional
	EbsRootVolumeSize *int64 `json:"ebsRootVolumeSize,omitempty" tf:"ebs_root_volume_size"`
	// +optional
	Ec2Attributes *ClusterSpecEc2Attributes `json:"ec2Attributes,omitempty" tf:"ec2_attributes"`
	// +optional
	KeepJobFlowAliveWhenNoSteps *bool `json:"keepJobFlowAliveWhenNoSteps,omitempty" tf:"keep_job_flow_alive_when_no_steps"`
	// +optional
	KerberosAttributes *ClusterSpecKerberosAttributes `json:"kerberosAttributes,omitempty" tf:"kerberos_attributes"`
	// +optional
	LogEncryptionKmsKeyID *string `json:"logEncryptionKmsKeyID,omitempty" tf:"log_encryption_kms_key_id"`
	// +optional
	LogURI *string `json:"logURI,omitempty" tf:"log_uri"`
	// +optional
	MasterInstanceFleet *ClusterSpecMasterInstanceFleet `json:"masterInstanceFleet,omitempty" tf:"master_instance_fleet"`
	// +optional
	MasterInstanceGroup *ClusterSpecMasterInstanceGroup `json:"masterInstanceGroup,omitempty" tf:"master_instance_group"`
	// +optional
	MasterPublicDNS *string `json:"masterPublicDNS,omitempty" tf:"master_public_dns"`
	Name            *string `json:"name" tf:"name"`
	ReleaseLabel    *string `json:"releaseLabel" tf:"release_label"`
	// +optional
	ScaleDownBehavior *string `json:"scaleDownBehavior,omitempty" tf:"scale_down_behavior"`
	// +optional
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration"`
	ServiceRole           *string `json:"serviceRole" tf:"service_role"`
	// +optional
	Step []ClusterSpecStep `json:"step,omitempty" tf:"step"`
	// +optional
	StepConcurrencyLevel *int64 `json:"stepConcurrencyLevel,omitempty" tf:"step_concurrency_level"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	TerminationProtection *bool `json:"terminationProtection,omitempty" tf:"termination_protection"`
	// +optional
	VisibleToAllUsers *bool `json:"visibleToAllUsers,omitempty" tf:"visible_to_all_users"`
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
