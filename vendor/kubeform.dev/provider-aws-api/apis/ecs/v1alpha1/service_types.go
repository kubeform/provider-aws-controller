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

type Service struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec,omitempty"`
	Status            ServiceStatus `json:"status,omitempty"`
}

type ServiceSpecCapacityProviderStrategy struct {
	// +optional
	Base             *int64  `json:"base,omitempty" tf:"base"`
	CapacityProvider *string `json:"capacityProvider" tf:"capacity_provider"`
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type ServiceSpecDeploymentCircuitBreaker struct {
	Enable   *bool `json:"enable" tf:"enable"`
	Rollback *bool `json:"rollback" tf:"rollback"`
}

type ServiceSpecDeploymentController struct {
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServiceSpecLoadBalancer struct {
	ContainerName *string `json:"containerName" tf:"container_name"`
	ContainerPort *int64  `json:"containerPort" tf:"container_port"`
	// +optional
	ElbName *string `json:"elbName,omitempty" tf:"elb_name"`
	// +optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn"`
}

type ServiceSpecNetworkConfiguration struct {
	// +optional
	AssignPublicIP *bool `json:"assignPublicIP,omitempty" tf:"assign_public_ip"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	Subnets        []string `json:"subnets" tf:"subnets"`
}

type ServiceSpecOrderedPlacementStrategy struct {
	// +optional
	Field *string `json:"field,omitempty" tf:"field"`
	Type  *string `json:"type" tf:"type"`
}

type ServiceSpecPlacementConstraints struct {
	// +optional
	Expression *string `json:"expression,omitempty" tf:"expression"`
	Type       *string `json:"type" tf:"type"`
}

type ServiceSpecServiceRegistries struct {
	// +optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name"`
	// +optional
	ContainerPort *int64 `json:"containerPort,omitempty" tf:"container_port"`
	// +optional
	Port        *int64  `json:"port,omitempty" tf:"port"`
	RegistryArn *string `json:"registryArn" tf:"registry_arn"`
}

type ServiceSpec struct {
	State *ServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ServiceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CapacityProviderStrategy []ServiceSpecCapacityProviderStrategy `json:"capacityProviderStrategy,omitempty" tf:"capacity_provider_strategy"`
	// +optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster"`
	// +optional
	DeploymentCircuitBreaker *ServiceSpecDeploymentCircuitBreaker `json:"deploymentCircuitBreaker,omitempty" tf:"deployment_circuit_breaker"`
	// +optional
	DeploymentController *ServiceSpecDeploymentController `json:"deploymentController,omitempty" tf:"deployment_controller"`
	// +optional
	DeploymentMaximumPercent *int64 `json:"deploymentMaximumPercent,omitempty" tf:"deployment_maximum_percent"`
	// +optional
	DeploymentMinimumHealthyPercent *int64 `json:"deploymentMinimumHealthyPercent,omitempty" tf:"deployment_minimum_healthy_percent"`
	// +optional
	DesiredCount *int64 `json:"desiredCount,omitempty" tf:"desired_count"`
	// +optional
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags"`
	// +optional
	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command"`
	// +optional
	ForceNewDeployment *bool `json:"forceNewDeployment,omitempty" tf:"force_new_deployment"`
	// +optional
	HealthCheckGracePeriodSeconds *int64 `json:"healthCheckGracePeriodSeconds,omitempty" tf:"health_check_grace_period_seconds"`
	// +optional
	IamRole *string `json:"iamRole,omitempty" tf:"iam_role"`
	// +optional
	LaunchType *string `json:"launchType,omitempty" tf:"launch_type"`
	// +optional
	LoadBalancer []ServiceSpecLoadBalancer `json:"loadBalancer,omitempty" tf:"load_balancer"`
	Name         *string                   `json:"name" tf:"name"`
	// +optional
	NetworkConfiguration *ServiceSpecNetworkConfiguration `json:"networkConfiguration,omitempty" tf:"network_configuration"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	OrderedPlacementStrategy []ServiceSpecOrderedPlacementStrategy `json:"orderedPlacementStrategy,omitempty" tf:"ordered_placement_strategy"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	PlacementConstraints []ServiceSpecPlacementConstraints `json:"placementConstraints,omitempty" tf:"placement_constraints"`
	// +optional
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version"`
	// +optional
	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags"`
	// +optional
	SchedulingStrategy *string `json:"schedulingStrategy,omitempty" tf:"scheduling_strategy"`
	// +optional
	ServiceRegistries *ServiceSpecServiceRegistries `json:"serviceRegistries,omitempty" tf:"service_registries"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	TaskDefinition *string `json:"taskDefinition,omitempty" tf:"task_definition"`
	// +optional
	WaitForSteadyState *bool `json:"waitForSteadyState,omitempty" tf:"wait_for_steady_state"`
}

type ServiceStatus struct {
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

// ServiceList is a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Service CRD objects
	Items []Service `json:"items,omitempty"`
}
