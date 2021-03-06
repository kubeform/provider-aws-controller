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

type MountTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MountTargetSpec   `json:"spec,omitempty"`
	Status            MountTargetStatus `json:"status,omitempty"`
}

type MountTargetSpec struct {
	State *MountTargetSpecResource `json:"state,omitempty" tf:"-"`

	Resource MountTargetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MountTargetSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AvailabilityZoneID *string `json:"availabilityZoneID,omitempty" tf:"availability_zone_id"`
	// +optional
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name"`
	// +optional
	DnsName *string `json:"dnsName,omitempty" tf:"dns_name"`
	// +optional
	FileSystemArn *string `json:"fileSystemArn,omitempty" tf:"file_system_arn"`
	FileSystemID  *string `json:"fileSystemID" tf:"file_system_id"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	MountTargetDNSName *string `json:"mountTargetDNSName,omitempty" tf:"mount_target_dns_name"`
	// +optional
	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty" tf:"network_interface_id"`
	// +optional
	OwnerID *string `json:"ownerID,omitempty" tf:"owner_id"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	SubnetID       *string  `json:"subnetID" tf:"subnet_id"`
}

type MountTargetStatus struct {
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

// MountTargetList is a list of MountTargets
type MountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MountTarget CRD objects
	Items []MountTarget `json:"items,omitempty"`
}
