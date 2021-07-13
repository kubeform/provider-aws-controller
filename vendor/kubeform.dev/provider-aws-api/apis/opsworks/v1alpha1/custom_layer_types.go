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

type CustomLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomLayerSpec   `json:"spec,omitempty"`
	Status            CustomLayerStatus `json:"status,omitempty"`
}

type CustomLayerSpecEbsVolume struct {
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

type CustomLayerSpec struct {
	State *CustomLayerSpecResource `json:"state,omitempty" tf:"-"`

	Resource CustomLayerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CustomLayerSpecResource struct {
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
	EbsVolume []CustomLayerSpecEbsVolume `json:"ebsVolume,omitempty" tf:"ebs_volume"`
	// +optional
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer"`
	// +optional
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot"`
	// +optional
	InstanceShutdownTimeout *int64  `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout"`
	Name                    *string `json:"name" tf:"name"`
	ShortName               *string `json:"shortName" tf:"short_name"`
	StackID                 *string `json:"stackID" tf:"stack_id"`
	// +optional
	SystemPackages []string `json:"systemPackages,omitempty" tf:"system_packages"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	UseEbsOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances"`
}

type CustomLayerStatus struct {
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

// CustomLayerList is a list of CustomLayers
type CustomLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CustomLayer CRD objects
	Items []CustomLayer `json:"items,omitempty"`
}
