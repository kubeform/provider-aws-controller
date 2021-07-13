// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomain) DeepCopyInto(out *ElasticsearchDomain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomain.
func (in *ElasticsearchDomain) DeepCopy() *ElasticsearchDomain {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchDomain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainList) DeepCopyInto(out *ElasticsearchDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticsearchDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainList.
func (in *ElasticsearchDomainList) DeepCopy() *ElasticsearchDomainList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpec) DeepCopyInto(out *ElasticsearchDomainSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ElasticsearchDomainSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpec.
func (in *ElasticsearchDomainSpec) DeepCopy() *ElasticsearchDomainSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecAdvancedSecurityOptions) DeepCopyInto(out *ElasticsearchDomainSpecAdvancedSecurityOptions) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.InternalUserDatabaseEnabled != nil {
		in, out := &in.InternalUserDatabaseEnabled, &out.InternalUserDatabaseEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MasterUserOptions != nil {
		in, out := &in.MasterUserOptions, &out.MasterUserOptions
		*out = new(ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecAdvancedSecurityOptions.
func (in *ElasticsearchDomainSpecAdvancedSecurityOptions) DeepCopy() *ElasticsearchDomainSpecAdvancedSecurityOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecAdvancedSecurityOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions) DeepCopyInto(out *ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions) {
	*out = *in
	if in.MasterUserArn != nil {
		in, out := &in.MasterUserArn, &out.MasterUserArn
		*out = new(string)
		**out = **in
	}
	if in.MasterUserName != nil {
		in, out := &in.MasterUserName, &out.MasterUserName
		*out = new(string)
		**out = **in
	}
	if in.MasterUserPassword != nil {
		in, out := &in.MasterUserPassword, &out.MasterUserPassword
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions.
func (in *ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions) DeepCopy() *ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecClusterConfig) DeepCopyInto(out *ElasticsearchDomainSpecClusterConfig) {
	*out = *in
	if in.DedicatedMasterCount != nil {
		in, out := &in.DedicatedMasterCount, &out.DedicatedMasterCount
		*out = new(int64)
		**out = **in
	}
	if in.DedicatedMasterEnabled != nil {
		in, out := &in.DedicatedMasterEnabled, &out.DedicatedMasterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DedicatedMasterType != nil {
		in, out := &in.DedicatedMasterType, &out.DedicatedMasterType
		*out = new(string)
		**out = **in
	}
	if in.InstanceCount != nil {
		in, out := &in.InstanceCount, &out.InstanceCount
		*out = new(int64)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.WarmCount != nil {
		in, out := &in.WarmCount, &out.WarmCount
		*out = new(int64)
		**out = **in
	}
	if in.WarmEnabled != nil {
		in, out := &in.WarmEnabled, &out.WarmEnabled
		*out = new(bool)
		**out = **in
	}
	if in.WarmType != nil {
		in, out := &in.WarmType, &out.WarmType
		*out = new(string)
		**out = **in
	}
	if in.ZoneAwarenessConfig != nil {
		in, out := &in.ZoneAwarenessConfig, &out.ZoneAwarenessConfig
		*out = new(ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneAwarenessEnabled != nil {
		in, out := &in.ZoneAwarenessEnabled, &out.ZoneAwarenessEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecClusterConfig.
func (in *ElasticsearchDomainSpecClusterConfig) DeepCopy() *ElasticsearchDomainSpecClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig) DeepCopyInto(out *ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig) {
	*out = *in
	if in.AvailabilityZoneCount != nil {
		in, out := &in.AvailabilityZoneCount, &out.AvailabilityZoneCount
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig.
func (in *ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig) DeepCopy() *ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecCognitoOptions) DeepCopyInto(out *ElasticsearchDomainSpecCognitoOptions) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IdentityPoolID != nil {
		in, out := &in.IdentityPoolID, &out.IdentityPoolID
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecCognitoOptions.
func (in *ElasticsearchDomainSpecCognitoOptions) DeepCopy() *ElasticsearchDomainSpecCognitoOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecCognitoOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecDomainEndpointOptions) DeepCopyInto(out *ElasticsearchDomainSpecDomainEndpointOptions) {
	*out = *in
	if in.CustomEndpoint != nil {
		in, out := &in.CustomEndpoint, &out.CustomEndpoint
		*out = new(string)
		**out = **in
	}
	if in.CustomEndpointCertificateArn != nil {
		in, out := &in.CustomEndpointCertificateArn, &out.CustomEndpointCertificateArn
		*out = new(string)
		**out = **in
	}
	if in.CustomEndpointEnabled != nil {
		in, out := &in.CustomEndpointEnabled, &out.CustomEndpointEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EnforceHTTPS != nil {
		in, out := &in.EnforceHTTPS, &out.EnforceHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.TlsSecurityPolicy != nil {
		in, out := &in.TlsSecurityPolicy, &out.TlsSecurityPolicy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecDomainEndpointOptions.
func (in *ElasticsearchDomainSpecDomainEndpointOptions) DeepCopy() *ElasticsearchDomainSpecDomainEndpointOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecDomainEndpointOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecEbsOptions) DeepCopyInto(out *ElasticsearchDomainSpecEbsOptions) {
	*out = *in
	if in.EbsEnabled != nil {
		in, out := &in.EbsEnabled, &out.EbsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecEbsOptions.
func (in *ElasticsearchDomainSpecEbsOptions) DeepCopy() *ElasticsearchDomainSpecEbsOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecEbsOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecEncryptAtRest) DeepCopyInto(out *ElasticsearchDomainSpecEncryptAtRest) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecEncryptAtRest.
func (in *ElasticsearchDomainSpecEncryptAtRest) DeepCopy() *ElasticsearchDomainSpecEncryptAtRest {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecEncryptAtRest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecLogPublishingOptions) DeepCopyInto(out *ElasticsearchDomainSpecLogPublishingOptions) {
	*out = *in
	if in.CloudwatchLogGroupArn != nil {
		in, out := &in.CloudwatchLogGroupArn, &out.CloudwatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogType != nil {
		in, out := &in.LogType, &out.LogType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecLogPublishingOptions.
func (in *ElasticsearchDomainSpecLogPublishingOptions) DeepCopy() *ElasticsearchDomainSpecLogPublishingOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecLogPublishingOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecNodeToNodeEncryption) DeepCopyInto(out *ElasticsearchDomainSpecNodeToNodeEncryption) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecNodeToNodeEncryption.
func (in *ElasticsearchDomainSpecNodeToNodeEncryption) DeepCopy() *ElasticsearchDomainSpecNodeToNodeEncryption {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecNodeToNodeEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecResource) DeepCopyInto(out *ElasticsearchDomainSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AccessPolicies != nil {
		in, out := &in.AccessPolicies, &out.AccessPolicies
		*out = new(string)
		**out = **in
	}
	if in.AdvancedOptions != nil {
		in, out := &in.AdvancedOptions, &out.AdvancedOptions
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.AdvancedSecurityOptions != nil {
		in, out := &in.AdvancedSecurityOptions, &out.AdvancedSecurityOptions
		*out = new(ElasticsearchDomainSpecAdvancedSecurityOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ClusterConfig != nil {
		in, out := &in.ClusterConfig, &out.ClusterConfig
		*out = new(ElasticsearchDomainSpecClusterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CognitoOptions != nil {
		in, out := &in.CognitoOptions, &out.CognitoOptions
		*out = new(ElasticsearchDomainSpecCognitoOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.DomainEndpointOptions != nil {
		in, out := &in.DomainEndpointOptions, &out.DomainEndpointOptions
		*out = new(ElasticsearchDomainSpecDomainEndpointOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.DomainID != nil {
		in, out := &in.DomainID, &out.DomainID
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.EbsOptions != nil {
		in, out := &in.EbsOptions, &out.EbsOptions
		*out = new(ElasticsearchDomainSpecEbsOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ElasticsearchVersion != nil {
		in, out := &in.ElasticsearchVersion, &out.ElasticsearchVersion
		*out = new(string)
		**out = **in
	}
	if in.EncryptAtRest != nil {
		in, out := &in.EncryptAtRest, &out.EncryptAtRest
		*out = new(ElasticsearchDomainSpecEncryptAtRest)
		(*in).DeepCopyInto(*out)
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.KibanaEndpoint != nil {
		in, out := &in.KibanaEndpoint, &out.KibanaEndpoint
		*out = new(string)
		**out = **in
	}
	if in.LogPublishingOptions != nil {
		in, out := &in.LogPublishingOptions, &out.LogPublishingOptions
		*out = make([]ElasticsearchDomainSpecLogPublishingOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeToNodeEncryption != nil {
		in, out := &in.NodeToNodeEncryption, &out.NodeToNodeEncryption
		*out = new(ElasticsearchDomainSpecNodeToNodeEncryption)
		(*in).DeepCopyInto(*out)
	}
	if in.SnapshotOptions != nil {
		in, out := &in.SnapshotOptions, &out.SnapshotOptions
		*out = new(ElasticsearchDomainSpecSnapshotOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.VpcOptions != nil {
		in, out := &in.VpcOptions, &out.VpcOptions
		*out = new(ElasticsearchDomainSpecVpcOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecResource.
func (in *ElasticsearchDomainSpecResource) DeepCopy() *ElasticsearchDomainSpecResource {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecSnapshotOptions) DeepCopyInto(out *ElasticsearchDomainSpecSnapshotOptions) {
	*out = *in
	if in.AutomatedSnapshotStartHour != nil {
		in, out := &in.AutomatedSnapshotStartHour, &out.AutomatedSnapshotStartHour
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecSnapshotOptions.
func (in *ElasticsearchDomainSpecSnapshotOptions) DeepCopy() *ElasticsearchDomainSpecSnapshotOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecSnapshotOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainSpecVpcOptions) DeepCopyInto(out *ElasticsearchDomainSpecVpcOptions) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDS != nil {
		in, out := &in.SecurityGroupIDS, &out.SecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainSpecVpcOptions.
func (in *ElasticsearchDomainSpecVpcOptions) DeepCopy() *ElasticsearchDomainSpecVpcOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainSpecVpcOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainStatus) DeepCopyInto(out *ElasticsearchDomainStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainStatus.
func (in *ElasticsearchDomainStatus) DeepCopy() *ElasticsearchDomainStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PolicySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecResource) DeepCopyInto(out *PolicySpecResource) {
	*out = *in
	if in.AccessPolicies != nil {
		in, out := &in.AccessPolicies, &out.AccessPolicies
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecResource.
func (in *PolicySpecResource) DeepCopy() *PolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(PolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlOptions) DeepCopyInto(out *SamlOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlOptions.
func (in *SamlOptions) DeepCopy() *SamlOptions {
	if in == nil {
		return nil
	}
	out := new(SamlOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SamlOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlOptionsList) DeepCopyInto(out *SamlOptionsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SamlOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlOptionsList.
func (in *SamlOptionsList) DeepCopy() *SamlOptionsList {
	if in == nil {
		return nil
	}
	out := new(SamlOptionsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SamlOptionsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlOptionsSpec) DeepCopyInto(out *SamlOptionsSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SamlOptionsSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlOptionsSpec.
func (in *SamlOptionsSpec) DeepCopy() *SamlOptionsSpec {
	if in == nil {
		return nil
	}
	out := new(SamlOptionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlOptionsSpecResource) DeepCopyInto(out *SamlOptionsSpecResource) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.SamlOptions != nil {
		in, out := &in.SamlOptions, &out.SamlOptions
		*out = new(SamlOptionsSpecSamlOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlOptionsSpecResource.
func (in *SamlOptionsSpecResource) DeepCopy() *SamlOptionsSpecResource {
	if in == nil {
		return nil
	}
	out := new(SamlOptionsSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlOptionsSpecSamlOptions) DeepCopyInto(out *SamlOptionsSpecSamlOptions) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Idp != nil {
		in, out := &in.Idp, &out.Idp
		*out = new(SamlOptionsSpecSamlOptionsIdp)
		(*in).DeepCopyInto(*out)
	}
	if in.MasterBackendRole != nil {
		in, out := &in.MasterBackendRole, &out.MasterBackendRole
		*out = new(string)
		**out = **in
	}
	if in.MasterUserName != nil {
		in, out := &in.MasterUserName, &out.MasterUserName
		*out = new(string)
		**out = **in
	}
	if in.RolesKey != nil {
		in, out := &in.RolesKey, &out.RolesKey
		*out = new(string)
		**out = **in
	}
	if in.SessionTimeoutMinutes != nil {
		in, out := &in.SessionTimeoutMinutes, &out.SessionTimeoutMinutes
		*out = new(int64)
		**out = **in
	}
	if in.SubjectKey != nil {
		in, out := &in.SubjectKey, &out.SubjectKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlOptionsSpecSamlOptions.
func (in *SamlOptionsSpecSamlOptions) DeepCopy() *SamlOptionsSpecSamlOptions {
	if in == nil {
		return nil
	}
	out := new(SamlOptionsSpecSamlOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlOptionsSpecSamlOptionsIdp) DeepCopyInto(out *SamlOptionsSpecSamlOptionsIdp) {
	*out = *in
	if in.EntityID != nil {
		in, out := &in.EntityID, &out.EntityID
		*out = new(string)
		**out = **in
	}
	if in.MetadataContent != nil {
		in, out := &in.MetadataContent, &out.MetadataContent
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlOptionsSpecSamlOptionsIdp.
func (in *SamlOptionsSpecSamlOptionsIdp) DeepCopy() *SamlOptionsSpecSamlOptionsIdp {
	if in == nil {
		return nil
	}
	out := new(SamlOptionsSpecSamlOptionsIdp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlOptionsStatus) DeepCopyInto(out *SamlOptionsStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlOptionsStatus.
func (in *SamlOptionsStatus) DeepCopy() *SamlOptionsStatus {
	if in == nil {
		return nil
	}
	out := new(SamlOptionsStatus)
	in.DeepCopyInto(out)
	return out
}
