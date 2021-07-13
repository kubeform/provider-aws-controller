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

type Subnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec,omitempty"`
	Status            SubnetStatus `json:"status,omitempty"`
}

type SubnetSpec struct {
	State *SubnetSpecResource `json:"state,omitempty" tf:"-"`

	Resource SubnetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SubnetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AssignIpv6AddressOnCreation *bool   `json:"assignIpv6AddressOnCreation,omitempty" tf:"assign_ipv6_address_on_creation"`
	AvailabilityZone            *string `json:"availabilityZone" tf:"availability_zone"`
	// +optional
	AvailabilityZoneID *string `json:"availabilityZoneID,omitempty" tf:"availability_zone_id"`
	// +optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`
	// +optional
	CustomerOwnedIpv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool"`
	// +optional
	Ipv6CIDRBlock *string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block"`
	// +optional
	Ipv6CIDRBlockAssociationID *string `json:"ipv6CIDRBlockAssociationID,omitempty" tf:"ipv6_cidr_block_association_id"`
	// +optional
	MapCustomerOwnedIPOnLaunch *bool `json:"mapCustomerOwnedIPOnLaunch,omitempty" tf:"map_customer_owned_ip_on_launch"`
	// +optional
	MapPublicIPOnLaunch *bool `json:"mapPublicIPOnLaunch,omitempty" tf:"map_public_ip_on_launch"`
	// +optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn"`
	// +optional
	OwnerID *string `json:"ownerID,omitempty" tf:"owner_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type SubnetStatus struct {
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

// SubnetList is a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Subnet CRD objects
	Items []Subnet `json:"items,omitempty"`
}
