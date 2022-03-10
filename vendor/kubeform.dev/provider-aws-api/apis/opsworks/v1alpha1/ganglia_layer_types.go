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

type GangliaLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GangliaLayerSpec   `json:"spec,omitempty"`
	Status            GangliaLayerStatus `json:"status,omitempty"`
}

type GangliaLayerSpecCloudwatchConfigurationLogStreams struct {
	// +optional
	BatchCount *int64 `json:"batchCount,omitempty" tf:"batch_count"`
	// +optional
	BatchSize *int64 `json:"batchSize,omitempty" tf:"batch_size"`
	// +optional
	BufferDuration *int64 `json:"bufferDuration,omitempty" tf:"buffer_duration"`
	// +optional
	DatetimeFormat *string `json:"datetimeFormat,omitempty" tf:"datetime_format"`
	// +optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`
	File     *string `json:"file" tf:"file"`
	// +optional
	FileFingerprintLines *string `json:"fileFingerprintLines,omitempty" tf:"file_fingerprint_lines"`
	// +optional
	InitialPosition *string `json:"initialPosition,omitempty" tf:"initial_position"`
	LogGroupName    *string `json:"logGroupName" tf:"log_group_name"`
	// +optional
	MultilineStartPattern *string `json:"multilineStartPattern,omitempty" tf:"multiline_start_pattern"`
	// +optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
}

type GangliaLayerSpecCloudwatchConfiguration struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	LogStreams []GangliaLayerSpecCloudwatchConfigurationLogStreams `json:"logStreams,omitempty" tf:"log_streams"`
}

type GangliaLayerSpecEbsVolume struct {
	// +optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	Iops          *int64  `json:"iops,omitempty" tf:"iops"`
	MountPoint    *string `json:"mountPoint" tf:"mount_point"`
	NumberOfDisks *int64  `json:"numberOfDisks" tf:"number_of_disks"`
	// +optional
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level"`
	Size      *int64  `json:"size" tf:"size"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type GangliaLayerSpec struct {
	State *GangliaLayerSpecResource `json:"state,omitempty" tf:"-"`

	Resource GangliaLayerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GangliaLayerSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AutoAssignElasticIPS *bool `json:"autoAssignElasticIPS,omitempty" tf:"auto_assign_elastic_ips"`
	// +optional
	AutoAssignPublicIPS *bool `json:"autoAssignPublicIPS,omitempty" tf:"auto_assign_public_ips"`
	// +optional
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing"`
	// +optional
	CloudwatchConfiguration *GangliaLayerSpecCloudwatchConfiguration `json:"cloudwatchConfiguration,omitempty" tf:"cloudwatch_configuration"`
	// +optional
	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes"`
	// +optional
	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes"`
	// +optional
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn"`
	// +optional
	CustomJSON *string `json:"customJSON,omitempty" tf:"custom_json"`
	// +optional
	CustomSecurityGroupIDS []string `json:"customSecurityGroupIDS,omitempty" tf:"custom_security_group_ids"`
	// +optional
	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes"`
	// +optional
	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes"`
	// +optional
	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes"`
	// +optional
	DrainElbOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown"`
	// +optional
	EbsVolume []GangliaLayerSpecEbsVolume `json:"ebsVolume,omitempty" tf:"ebs_volume"`
	// +optional
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer"`
	// +optional
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot"`
	// +optional
	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout"`
	// +optional
	Name     *string `json:"name,omitempty" tf:"name"`
	Password *string `json:"password" tf:"password"`
	StackID  *string `json:"stackID" tf:"stack_id"`
	// +optional
	SystemPackages []string `json:"systemPackages,omitempty" tf:"system_packages"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
	// +optional
	UseEbsOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type GangliaLayerStatus struct {
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

// GangliaLayerList is a list of GangliaLayers
type GangliaLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GangliaLayer CRD objects
	Items []GangliaLayer `json:"items,omitempty"`
}
