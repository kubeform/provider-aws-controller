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
func (in *Attachment) DeepCopyInto(out *Attachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attachment.
func (in *Attachment) DeepCopy() *Attachment {
	if in == nil {
		return nil
	}
	out := new(Attachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Attachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachmentList) DeepCopyInto(out *AttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Attachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentList.
func (in *AttachmentList) DeepCopy() *AttachmentList {
	if in == nil {
		return nil
	}
	out := new(AttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachmentSpec) DeepCopyInto(out *AttachmentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AttachmentSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentSpec.
func (in *AttachmentSpec) DeepCopy() *AttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(AttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachmentSpecResource) DeepCopyInto(out *AttachmentSpecResource) {
	*out = *in
	if in.Elb != nil {
		in, out := &in.Elb, &out.Elb
		*out = new(string)
		**out = **in
	}
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentSpecResource.
func (in *AttachmentSpecResource) DeepCopy() *AttachmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(AttachmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachmentStatus) DeepCopyInto(out *AttachmentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentStatus.
func (in *AttachmentStatus) DeepCopy() *AttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(AttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elb) DeepCopyInto(out *Elb) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elb.
func (in *Elb) DeepCopy() *Elb {
	if in == nil {
		return nil
	}
	out := new(Elb)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Elb) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElbList) DeepCopyInto(out *ElbList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Elb, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElbList.
func (in *ElbList) DeepCopy() *ElbList {
	if in == nil {
		return nil
	}
	out := new(ElbList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElbList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElbSpec) DeepCopyInto(out *ElbSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ElbSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElbSpec.
func (in *ElbSpec) DeepCopy() *ElbSpec {
	if in == nil {
		return nil
	}
	out := new(ElbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElbSpecAccessLogs) DeepCopyInto(out *ElbSpecAccessLogs) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.BucketPrefix != nil {
		in, out := &in.BucketPrefix, &out.BucketPrefix
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElbSpecAccessLogs.
func (in *ElbSpecAccessLogs) DeepCopy() *ElbSpecAccessLogs {
	if in == nil {
		return nil
	}
	out := new(ElbSpecAccessLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElbSpecHealthCheck) DeepCopyInto(out *ElbSpecHealthCheck) {
	*out = *in
	if in.HealthyThreshold != nil {
		in, out := &in.HealthyThreshold, &out.HealthyThreshold
		*out = new(int64)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	if in.UnhealthyThreshold != nil {
		in, out := &in.UnhealthyThreshold, &out.UnhealthyThreshold
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElbSpecHealthCheck.
func (in *ElbSpecHealthCheck) DeepCopy() *ElbSpecHealthCheck {
	if in == nil {
		return nil
	}
	out := new(ElbSpecHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElbSpecListener) DeepCopyInto(out *ElbSpecListener) {
	*out = *in
	if in.InstancePort != nil {
		in, out := &in.InstancePort, &out.InstancePort
		*out = new(int64)
		**out = **in
	}
	if in.InstanceProtocol != nil {
		in, out := &in.InstanceProtocol, &out.InstanceProtocol
		*out = new(string)
		**out = **in
	}
	if in.LbPort != nil {
		in, out := &in.LbPort, &out.LbPort
		*out = new(int64)
		**out = **in
	}
	if in.LbProtocol != nil {
		in, out := &in.LbProtocol, &out.LbProtocol
		*out = new(string)
		**out = **in
	}
	if in.SslCertificateID != nil {
		in, out := &in.SslCertificateID, &out.SslCertificateID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElbSpecListener.
func (in *ElbSpecListener) DeepCopy() *ElbSpecListener {
	if in == nil {
		return nil
	}
	out := new(ElbSpecListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElbSpecResource) DeepCopyInto(out *ElbSpecResource) {
	*out = *in
	if in.AccessLogs != nil {
		in, out := &in.AccessLogs, &out.AccessLogs
		*out = new(ElbSpecAccessLogs)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConnectionDraining != nil {
		in, out := &in.ConnectionDraining, &out.ConnectionDraining
		*out = new(bool)
		**out = **in
	}
	if in.ConnectionDrainingTimeout != nil {
		in, out := &in.ConnectionDrainingTimeout, &out.ConnectionDrainingTimeout
		*out = new(int64)
		**out = **in
	}
	if in.CrossZoneLoadBalancing != nil {
		in, out := &in.CrossZoneLoadBalancing, &out.CrossZoneLoadBalancing
		*out = new(bool)
		**out = **in
	}
	if in.DnsName != nil {
		in, out := &in.DnsName, &out.DnsName
		*out = new(string)
		**out = **in
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(ElbSpecHealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.IdleTimeout != nil {
		in, out := &in.IdleTimeout, &out.IdleTimeout
		*out = new(int64)
		**out = **in
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Internal != nil {
		in, out := &in.Internal, &out.Internal
		*out = new(bool)
		**out = **in
	}
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = make([]ElbSpecListener, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceSecurityGroup != nil {
		in, out := &in.SourceSecurityGroup, &out.SourceSecurityGroup
		*out = new(string)
		**out = **in
	}
	if in.SourceSecurityGroupID != nil {
		in, out := &in.SourceSecurityGroupID, &out.SourceSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
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
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElbSpecResource.
func (in *ElbSpecResource) DeepCopy() *ElbSpecResource {
	if in == nil {
		return nil
	}
	out := new(ElbSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElbStatus) DeepCopyInto(out *ElbStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElbStatus.
func (in *ElbStatus) DeepCopy() *ElbStatus {
	if in == nil {
		return nil
	}
	out := new(ElbStatus)
	in.DeepCopyInto(out)
	return out
}
