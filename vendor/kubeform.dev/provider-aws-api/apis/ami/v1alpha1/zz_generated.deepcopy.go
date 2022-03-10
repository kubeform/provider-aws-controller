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
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ami) DeepCopyInto(out *Ami) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ami.
func (in *Ami) DeepCopy() *Ami {
	if in == nil {
		return nil
	}
	out := new(Ami)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ami) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiList) DeepCopyInto(out *AmiList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ami, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiList.
func (in *AmiList) DeepCopy() *AmiList {
	if in == nil {
		return nil
	}
	out := new(AmiList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AmiList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiSpec) DeepCopyInto(out *AmiSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AmiSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiSpec.
func (in *AmiSpec) DeepCopy() *AmiSpec {
	if in == nil {
		return nil
	}
	out := new(AmiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiSpecEbsBlockDevice) DeepCopyInto(out *AmiSpecEbsBlockDevice) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.OutpostArn != nil {
		in, out := &in.OutpostArn, &out.OutpostArn
		*out = new(string)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiSpecEbsBlockDevice.
func (in *AmiSpecEbsBlockDevice) DeepCopy() *AmiSpecEbsBlockDevice {
	if in == nil {
		return nil
	}
	out := new(AmiSpecEbsBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiSpecEphemeralBlockDevice) DeepCopyInto(out *AmiSpecEphemeralBlockDevice) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiSpecEphemeralBlockDevice.
func (in *AmiSpecEphemeralBlockDevice) DeepCopy() *AmiSpecEphemeralBlockDevice {
	if in == nil {
		return nil
	}
	out := new(AmiSpecEphemeralBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiSpecResource) DeepCopyInto(out *AmiSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.BootMode != nil {
		in, out := &in.BootMode, &out.BootMode
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EbsBlockDevice != nil {
		in, out := &in.EbsBlockDevice, &out.EbsBlockDevice
		*out = make([]AmiSpecEbsBlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnaSupport != nil {
		in, out := &in.EnaSupport, &out.EnaSupport
		*out = new(bool)
		**out = **in
	}
	if in.EphemeralBlockDevice != nil {
		in, out := &in.EphemeralBlockDevice, &out.EphemeralBlockDevice
		*out = make([]AmiSpecEphemeralBlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hypervisor != nil {
		in, out := &in.Hypervisor, &out.Hypervisor
		*out = new(string)
		**out = **in
	}
	if in.ImageLocation != nil {
		in, out := &in.ImageLocation, &out.ImageLocation
		*out = new(string)
		**out = **in
	}
	if in.ImageOwnerAlias != nil {
		in, out := &in.ImageOwnerAlias, &out.ImageOwnerAlias
		*out = new(string)
		**out = **in
	}
	if in.ImageType != nil {
		in, out := &in.ImageType, &out.ImageType
		*out = new(string)
		**out = **in
	}
	if in.KernelID != nil {
		in, out := &in.KernelID, &out.KernelID
		*out = new(string)
		**out = **in
	}
	if in.ManageEbsSnapshots != nil {
		in, out := &in.ManageEbsSnapshots, &out.ManageEbsSnapshots
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(string)
		**out = **in
	}
	if in.PlatformDetails != nil {
		in, out := &in.PlatformDetails, &out.PlatformDetails
		*out = new(string)
		**out = **in
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.RamdiskID != nil {
		in, out := &in.RamdiskID, &out.RamdiskID
		*out = new(string)
		**out = **in
	}
	if in.RootDeviceName != nil {
		in, out := &in.RootDeviceName, &out.RootDeviceName
		*out = new(string)
		**out = **in
	}
	if in.RootSnapshotID != nil {
		in, out := &in.RootSnapshotID, &out.RootSnapshotID
		*out = new(string)
		**out = **in
	}
	if in.SriovNetSupport != nil {
		in, out := &in.SriovNetSupport, &out.SriovNetSupport
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
	if in.UsageOperation != nil {
		in, out := &in.UsageOperation, &out.UsageOperation
		*out = new(string)
		**out = **in
	}
	if in.VirtualizationType != nil {
		in, out := &in.VirtualizationType, &out.VirtualizationType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiSpecResource.
func (in *AmiSpecResource) DeepCopy() *AmiSpecResource {
	if in == nil {
		return nil
	}
	out := new(AmiSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiStatus) DeepCopyInto(out *AmiStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiStatus.
func (in *AmiStatus) DeepCopy() *AmiStatus {
	if in == nil {
		return nil
	}
	out := new(AmiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Copy) DeepCopyInto(out *Copy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Copy.
func (in *Copy) DeepCopy() *Copy {
	if in == nil {
		return nil
	}
	out := new(Copy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Copy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopyList) DeepCopyInto(out *CopyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Copy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyList.
func (in *CopyList) DeepCopy() *CopyList {
	if in == nil {
		return nil
	}
	out := new(CopyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CopyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopySpec) DeepCopyInto(out *CopySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CopySpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopySpec.
func (in *CopySpec) DeepCopy() *CopySpec {
	if in == nil {
		return nil
	}
	out := new(CopySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopySpecEbsBlockDevice) DeepCopyInto(out *CopySpecEbsBlockDevice) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.OutpostArn != nil {
		in, out := &in.OutpostArn, &out.OutpostArn
		*out = new(string)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopySpecEbsBlockDevice.
func (in *CopySpecEbsBlockDevice) DeepCopy() *CopySpecEbsBlockDevice {
	if in == nil {
		return nil
	}
	out := new(CopySpecEbsBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopySpecEphemeralBlockDevice) DeepCopyInto(out *CopySpecEphemeralBlockDevice) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopySpecEphemeralBlockDevice.
func (in *CopySpecEphemeralBlockDevice) DeepCopy() *CopySpecEphemeralBlockDevice {
	if in == nil {
		return nil
	}
	out := new(CopySpecEphemeralBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopySpecResource) DeepCopyInto(out *CopySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.BootMode != nil {
		in, out := &in.BootMode, &out.BootMode
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DestinationOutpostArn != nil {
		in, out := &in.DestinationOutpostArn, &out.DestinationOutpostArn
		*out = new(string)
		**out = **in
	}
	if in.EbsBlockDevice != nil {
		in, out := &in.EbsBlockDevice, &out.EbsBlockDevice
		*out = make([]CopySpecEbsBlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnaSupport != nil {
		in, out := &in.EnaSupport, &out.EnaSupport
		*out = new(bool)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.EphemeralBlockDevice != nil {
		in, out := &in.EphemeralBlockDevice, &out.EphemeralBlockDevice
		*out = make([]CopySpecEphemeralBlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hypervisor != nil {
		in, out := &in.Hypervisor, &out.Hypervisor
		*out = new(string)
		**out = **in
	}
	if in.ImageLocation != nil {
		in, out := &in.ImageLocation, &out.ImageLocation
		*out = new(string)
		**out = **in
	}
	if in.ImageOwnerAlias != nil {
		in, out := &in.ImageOwnerAlias, &out.ImageOwnerAlias
		*out = new(string)
		**out = **in
	}
	if in.ImageType != nil {
		in, out := &in.ImageType, &out.ImageType
		*out = new(string)
		**out = **in
	}
	if in.KernelID != nil {
		in, out := &in.KernelID, &out.KernelID
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.ManageEbsSnapshots != nil {
		in, out := &in.ManageEbsSnapshots, &out.ManageEbsSnapshots
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(string)
		**out = **in
	}
	if in.PlatformDetails != nil {
		in, out := &in.PlatformDetails, &out.PlatformDetails
		*out = new(string)
		**out = **in
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.RamdiskID != nil {
		in, out := &in.RamdiskID, &out.RamdiskID
		*out = new(string)
		**out = **in
	}
	if in.RootDeviceName != nil {
		in, out := &in.RootDeviceName, &out.RootDeviceName
		*out = new(string)
		**out = **in
	}
	if in.RootSnapshotID != nil {
		in, out := &in.RootSnapshotID, &out.RootSnapshotID
		*out = new(string)
		**out = **in
	}
	if in.SourceAmiID != nil {
		in, out := &in.SourceAmiID, &out.SourceAmiID
		*out = new(string)
		**out = **in
	}
	if in.SourceAmiRegion != nil {
		in, out := &in.SourceAmiRegion, &out.SourceAmiRegion
		*out = new(string)
		**out = **in
	}
	if in.SriovNetSupport != nil {
		in, out := &in.SriovNetSupport, &out.SriovNetSupport
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
	if in.UsageOperation != nil {
		in, out := &in.UsageOperation, &out.UsageOperation
		*out = new(string)
		**out = **in
	}
	if in.VirtualizationType != nil {
		in, out := &in.VirtualizationType, &out.VirtualizationType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopySpecResource.
func (in *CopySpecResource) DeepCopy() *CopySpecResource {
	if in == nil {
		return nil
	}
	out := new(CopySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopyStatus) DeepCopyInto(out *CopyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyStatus.
func (in *CopyStatus) DeepCopy() *CopyStatus {
	if in == nil {
		return nil
	}
	out := new(CopyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromInstance) DeepCopyInto(out *FromInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromInstance.
func (in *FromInstance) DeepCopy() *FromInstance {
	if in == nil {
		return nil
	}
	out := new(FromInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FromInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromInstanceList) DeepCopyInto(out *FromInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FromInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromInstanceList.
func (in *FromInstanceList) DeepCopy() *FromInstanceList {
	if in == nil {
		return nil
	}
	out := new(FromInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FromInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromInstanceSpec) DeepCopyInto(out *FromInstanceSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(FromInstanceSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromInstanceSpec.
func (in *FromInstanceSpec) DeepCopy() *FromInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(FromInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromInstanceSpecEbsBlockDevice) DeepCopyInto(out *FromInstanceSpecEbsBlockDevice) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.OutpostArn != nil {
		in, out := &in.OutpostArn, &out.OutpostArn
		*out = new(string)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromInstanceSpecEbsBlockDevice.
func (in *FromInstanceSpecEbsBlockDevice) DeepCopy() *FromInstanceSpecEbsBlockDevice {
	if in == nil {
		return nil
	}
	out := new(FromInstanceSpecEbsBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromInstanceSpecEphemeralBlockDevice) DeepCopyInto(out *FromInstanceSpecEphemeralBlockDevice) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromInstanceSpecEphemeralBlockDevice.
func (in *FromInstanceSpecEphemeralBlockDevice) DeepCopy() *FromInstanceSpecEphemeralBlockDevice {
	if in == nil {
		return nil
	}
	out := new(FromInstanceSpecEphemeralBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromInstanceSpecResource) DeepCopyInto(out *FromInstanceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.BootMode != nil {
		in, out := &in.BootMode, &out.BootMode
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EbsBlockDevice != nil {
		in, out := &in.EbsBlockDevice, &out.EbsBlockDevice
		*out = make([]FromInstanceSpecEbsBlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnaSupport != nil {
		in, out := &in.EnaSupport, &out.EnaSupport
		*out = new(bool)
		**out = **in
	}
	if in.EphemeralBlockDevice != nil {
		in, out := &in.EphemeralBlockDevice, &out.EphemeralBlockDevice
		*out = make([]FromInstanceSpecEphemeralBlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hypervisor != nil {
		in, out := &in.Hypervisor, &out.Hypervisor
		*out = new(string)
		**out = **in
	}
	if in.ImageLocation != nil {
		in, out := &in.ImageLocation, &out.ImageLocation
		*out = new(string)
		**out = **in
	}
	if in.ImageOwnerAlias != nil {
		in, out := &in.ImageOwnerAlias, &out.ImageOwnerAlias
		*out = new(string)
		**out = **in
	}
	if in.ImageType != nil {
		in, out := &in.ImageType, &out.ImageType
		*out = new(string)
		**out = **in
	}
	if in.KernelID != nil {
		in, out := &in.KernelID, &out.KernelID
		*out = new(string)
		**out = **in
	}
	if in.ManageEbsSnapshots != nil {
		in, out := &in.ManageEbsSnapshots, &out.ManageEbsSnapshots
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(string)
		**out = **in
	}
	if in.PlatformDetails != nil {
		in, out := &in.PlatformDetails, &out.PlatformDetails
		*out = new(string)
		**out = **in
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.RamdiskID != nil {
		in, out := &in.RamdiskID, &out.RamdiskID
		*out = new(string)
		**out = **in
	}
	if in.RootDeviceName != nil {
		in, out := &in.RootDeviceName, &out.RootDeviceName
		*out = new(string)
		**out = **in
	}
	if in.RootSnapshotID != nil {
		in, out := &in.RootSnapshotID, &out.RootSnapshotID
		*out = new(string)
		**out = **in
	}
	if in.SnapshotWithoutReboot != nil {
		in, out := &in.SnapshotWithoutReboot, &out.SnapshotWithoutReboot
		*out = new(bool)
		**out = **in
	}
	if in.SourceInstanceID != nil {
		in, out := &in.SourceInstanceID, &out.SourceInstanceID
		*out = new(string)
		**out = **in
	}
	if in.SriovNetSupport != nil {
		in, out := &in.SriovNetSupport, &out.SriovNetSupport
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
	if in.UsageOperation != nil {
		in, out := &in.UsageOperation, &out.UsageOperation
		*out = new(string)
		**out = **in
	}
	if in.VirtualizationType != nil {
		in, out := &in.VirtualizationType, &out.VirtualizationType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromInstanceSpecResource.
func (in *FromInstanceSpecResource) DeepCopy() *FromInstanceSpecResource {
	if in == nil {
		return nil
	}
	out := new(FromInstanceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromInstanceStatus) DeepCopyInto(out *FromInstanceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromInstanceStatus.
func (in *FromInstanceStatus) DeepCopy() *FromInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(FromInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchPermission) DeepCopyInto(out *LaunchPermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchPermission.
func (in *LaunchPermission) DeepCopy() *LaunchPermission {
	if in == nil {
		return nil
	}
	out := new(LaunchPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchPermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchPermissionList) DeepCopyInto(out *LaunchPermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LaunchPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchPermissionList.
func (in *LaunchPermissionList) DeepCopy() *LaunchPermissionList {
	if in == nil {
		return nil
	}
	out := new(LaunchPermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchPermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchPermissionSpec) DeepCopyInto(out *LaunchPermissionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LaunchPermissionSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchPermissionSpec.
func (in *LaunchPermissionSpec) DeepCopy() *LaunchPermissionSpec {
	if in == nil {
		return nil
	}
	out := new(LaunchPermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchPermissionSpecResource) DeepCopyInto(out *LaunchPermissionSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchPermissionSpecResource.
func (in *LaunchPermissionSpecResource) DeepCopy() *LaunchPermissionSpecResource {
	if in == nil {
		return nil
	}
	out := new(LaunchPermissionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchPermissionStatus) DeepCopyInto(out *LaunchPermissionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchPermissionStatus.
func (in *LaunchPermissionStatus) DeepCopy() *LaunchPermissionStatus {
	if in == nil {
		return nil
	}
	out := new(LaunchPermissionStatus)
	in.DeepCopyInto(out)
	return out
}
