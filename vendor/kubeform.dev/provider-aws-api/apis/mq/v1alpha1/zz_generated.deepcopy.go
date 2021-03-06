//go:build !ignore_autogenerated
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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Broker) DeepCopyInto(out *Broker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Broker.
func (in *Broker) DeepCopy() *Broker {
	if in == nil {
		return nil
	}
	out := new(Broker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Broker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerList) DeepCopyInto(out *BrokerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Broker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerList.
func (in *BrokerList) DeepCopy() *BrokerList {
	if in == nil {
		return nil
	}
	out := new(BrokerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BrokerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpec) DeepCopyInto(out *BrokerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BrokerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpec.
func (in *BrokerSpec) DeepCopy() *BrokerSpec {
	if in == nil {
		return nil
	}
	out := new(BrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpecConfiguration) DeepCopyInto(out *BrokerSpecConfiguration) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpecConfiguration.
func (in *BrokerSpecConfiguration) DeepCopy() *BrokerSpecConfiguration {
	if in == nil {
		return nil
	}
	out := new(BrokerSpecConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpecEncryptionOptions) DeepCopyInto(out *BrokerSpecEncryptionOptions) {
	*out = *in
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.UseAwsOwnedKey != nil {
		in, out := &in.UseAwsOwnedKey, &out.UseAwsOwnedKey
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpecEncryptionOptions.
func (in *BrokerSpecEncryptionOptions) DeepCopy() *BrokerSpecEncryptionOptions {
	if in == nil {
		return nil
	}
	out := new(BrokerSpecEncryptionOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpecInstances) DeepCopyInto(out *BrokerSpecInstances) {
	*out = *in
	if in.ConsoleURL != nil {
		in, out := &in.ConsoleURL, &out.ConsoleURL
		*out = new(string)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpecInstances.
func (in *BrokerSpecInstances) DeepCopy() *BrokerSpecInstances {
	if in == nil {
		return nil
	}
	out := new(BrokerSpecInstances)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpecLdapServerMetadata) DeepCopyInto(out *BrokerSpecLdapServerMetadata) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RoleBase != nil {
		in, out := &in.RoleBase, &out.RoleBase
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchMatching != nil {
		in, out := &in.RoleSearchMatching, &out.RoleSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchSubtree != nil {
		in, out := &in.RoleSearchSubtree, &out.RoleSearchSubtree
		*out = new(bool)
		**out = **in
	}
	if in.ServiceAccountPassword != nil {
		in, out := &in.ServiceAccountPassword, &out.ServiceAccountPassword
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountUsername != nil {
		in, out := &in.ServiceAccountUsername, &out.ServiceAccountUsername
		*out = new(string)
		**out = **in
	}
	if in.UserBase != nil {
		in, out := &in.UserBase, &out.UserBase
		*out = new(string)
		**out = **in
	}
	if in.UserRoleName != nil {
		in, out := &in.UserRoleName, &out.UserRoleName
		*out = new(string)
		**out = **in
	}
	if in.UserSearchMatching != nil {
		in, out := &in.UserSearchMatching, &out.UserSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.UserSearchSubtree != nil {
		in, out := &in.UserSearchSubtree, &out.UserSearchSubtree
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpecLdapServerMetadata.
func (in *BrokerSpecLdapServerMetadata) DeepCopy() *BrokerSpecLdapServerMetadata {
	if in == nil {
		return nil
	}
	out := new(BrokerSpecLdapServerMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpecLogs) DeepCopyInto(out *BrokerSpecLogs) {
	*out = *in
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(string)
		**out = **in
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpecLogs.
func (in *BrokerSpecLogs) DeepCopy() *BrokerSpecLogs {
	if in == nil {
		return nil
	}
	out := new(BrokerSpecLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpecMaintenanceWindowStartTime) DeepCopyInto(out *BrokerSpecMaintenanceWindowStartTime) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.TimeOfDay != nil {
		in, out := &in.TimeOfDay, &out.TimeOfDay
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpecMaintenanceWindowStartTime.
func (in *BrokerSpecMaintenanceWindowStartTime) DeepCopy() *BrokerSpecMaintenanceWindowStartTime {
	if in == nil {
		return nil
	}
	out := new(BrokerSpecMaintenanceWindowStartTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpecResource) DeepCopyInto(out *BrokerSpecResource) {
	*out = *in
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationStrategy != nil {
		in, out := &in.AuthenticationStrategy, &out.AuthenticationStrategy
		*out = new(string)
		**out = **in
	}
	if in.AutoMinorVersionUpgrade != nil {
		in, out := &in.AutoMinorVersionUpgrade, &out.AutoMinorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.BrokerName != nil {
		in, out := &in.BrokerName, &out.BrokerName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(BrokerSpecConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentMode != nil {
		in, out := &in.DeploymentMode, &out.DeploymentMode
		*out = new(string)
		**out = **in
	}
	if in.EncryptionOptions != nil {
		in, out := &in.EncryptionOptions, &out.EncryptionOptions
		*out = new(BrokerSpecEncryptionOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.HostInstanceType != nil {
		in, out := &in.HostInstanceType, &out.HostInstanceType
		*out = new(string)
		**out = **in
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]BrokerSpecInstances, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LdapServerMetadata != nil {
		in, out := &in.LdapServerMetadata, &out.LdapServerMetadata
		*out = new(BrokerSpecLdapServerMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = new(BrokerSpecLogs)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceWindowStartTime != nil {
		in, out := &in.MaintenanceWindowStartTime, &out.MaintenanceWindowStartTime
		*out = new(BrokerSpecMaintenanceWindowStartTime)
		(*in).DeepCopyInto(*out)
	}
	if in.PubliclyAccessible != nil {
		in, out := &in.PubliclyAccessible, &out.PubliclyAccessible
		*out = new(bool)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = make([]BrokerSpecUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpecResource.
func (in *BrokerSpecResource) DeepCopy() *BrokerSpecResource {
	if in == nil {
		return nil
	}
	out := new(BrokerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpecUser) DeepCopyInto(out *BrokerSpecUser) {
	*out = *in
	if in.ConsoleAccess != nil {
		in, out := &in.ConsoleAccess, &out.ConsoleAccess
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpecUser.
func (in *BrokerSpecUser) DeepCopy() *BrokerSpecUser {
	if in == nil {
		return nil
	}
	out := new(BrokerSpecUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerStatus) DeepCopyInto(out *BrokerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerStatus.
func (in *BrokerStatus) DeepCopy() *BrokerStatus {
	if in == nil {
		return nil
	}
	out := new(BrokerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationList) DeepCopyInto(out *ConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Configuration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationList.
func (in *ConfigurationList) DeepCopy() *ConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ConfigurationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecResource) DeepCopyInto(out *ConfigurationSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationStrategy != nil {
		in, out := &in.AuthenticationStrategy, &out.AuthenticationStrategy
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.LatestRevision != nil {
		in, out := &in.LatestRevision, &out.LatestRevision
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecResource.
func (in *ConfigurationSpecResource) DeepCopy() *ConfigurationSpecResource {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStatus) DeepCopyInto(out *ConfigurationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStatus.
func (in *ConfigurationStatus) DeepCopy() *ConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}
