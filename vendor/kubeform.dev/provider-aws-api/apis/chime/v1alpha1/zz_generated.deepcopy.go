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
func (in *VoiceConnector) DeepCopyInto(out *VoiceConnector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnector.
func (in *VoiceConnector) DeepCopy() *VoiceConnector {
	if in == nil {
		return nil
	}
	out := new(VoiceConnector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorGroup) DeepCopyInto(out *VoiceConnectorGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorGroup.
func (in *VoiceConnectorGroup) DeepCopy() *VoiceConnectorGroup {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorGroupList) DeepCopyInto(out *VoiceConnectorGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VoiceConnectorGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorGroupList.
func (in *VoiceConnectorGroupList) DeepCopy() *VoiceConnectorGroupList {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorGroupSpec) DeepCopyInto(out *VoiceConnectorGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VoiceConnectorGroupSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorGroupSpec.
func (in *VoiceConnectorGroupSpec) DeepCopy() *VoiceConnectorGroupSpec {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorGroupSpecConnector) DeepCopyInto(out *VoiceConnectorGroupSpecConnector) {
	*out = *in
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.VoiceConnectorID != nil {
		in, out := &in.VoiceConnectorID, &out.VoiceConnectorID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorGroupSpecConnector.
func (in *VoiceConnectorGroupSpecConnector) DeepCopy() *VoiceConnectorGroupSpecConnector {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorGroupSpecConnector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorGroupSpecResource) DeepCopyInto(out *VoiceConnectorGroupSpecResource) {
	*out = *in
	if in.Connector != nil {
		in, out := &in.Connector, &out.Connector
		*out = make([]VoiceConnectorGroupSpecConnector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorGroupSpecResource.
func (in *VoiceConnectorGroupSpecResource) DeepCopy() *VoiceConnectorGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorGroupStatus) DeepCopyInto(out *VoiceConnectorGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorGroupStatus.
func (in *VoiceConnectorGroupStatus) DeepCopy() *VoiceConnectorGroupStatus {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorList) DeepCopyInto(out *VoiceConnectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VoiceConnector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorList.
func (in *VoiceConnectorList) DeepCopy() *VoiceConnectorList {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorLogging) DeepCopyInto(out *VoiceConnectorLogging) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorLogging.
func (in *VoiceConnectorLogging) DeepCopy() *VoiceConnectorLogging {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorLogging) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorLoggingList) DeepCopyInto(out *VoiceConnectorLoggingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VoiceConnectorLogging, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorLoggingList.
func (in *VoiceConnectorLoggingList) DeepCopy() *VoiceConnectorLoggingList {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorLoggingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorLoggingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorLoggingSpec) DeepCopyInto(out *VoiceConnectorLoggingSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VoiceConnectorLoggingSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorLoggingSpec.
func (in *VoiceConnectorLoggingSpec) DeepCopy() *VoiceConnectorLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorLoggingSpecResource) DeepCopyInto(out *VoiceConnectorLoggingSpecResource) {
	*out = *in
	if in.EnableSipLogs != nil {
		in, out := &in.EnableSipLogs, &out.EnableSipLogs
		*out = new(bool)
		**out = **in
	}
	if in.VoiceConnectorID != nil {
		in, out := &in.VoiceConnectorID, &out.VoiceConnectorID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorLoggingSpecResource.
func (in *VoiceConnectorLoggingSpecResource) DeepCopy() *VoiceConnectorLoggingSpecResource {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorLoggingSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorLoggingStatus) DeepCopyInto(out *VoiceConnectorLoggingStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorLoggingStatus.
func (in *VoiceConnectorLoggingStatus) DeepCopy() *VoiceConnectorLoggingStatus {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorLoggingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorOrigination) DeepCopyInto(out *VoiceConnectorOrigination) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorOrigination.
func (in *VoiceConnectorOrigination) DeepCopy() *VoiceConnectorOrigination {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorOrigination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorOrigination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorOriginationList) DeepCopyInto(out *VoiceConnectorOriginationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VoiceConnectorOrigination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorOriginationList.
func (in *VoiceConnectorOriginationList) DeepCopy() *VoiceConnectorOriginationList {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorOriginationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorOriginationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorOriginationSpec) DeepCopyInto(out *VoiceConnectorOriginationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VoiceConnectorOriginationSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorOriginationSpec.
func (in *VoiceConnectorOriginationSpec) DeepCopy() *VoiceConnectorOriginationSpec {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorOriginationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorOriginationSpecResource) DeepCopyInto(out *VoiceConnectorOriginationSpecResource) {
	*out = *in
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = make([]VoiceConnectorOriginationSpecRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VoiceConnectorID != nil {
		in, out := &in.VoiceConnectorID, &out.VoiceConnectorID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorOriginationSpecResource.
func (in *VoiceConnectorOriginationSpecResource) DeepCopy() *VoiceConnectorOriginationSpecResource {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorOriginationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorOriginationSpecRoute) DeepCopyInto(out *VoiceConnectorOriginationSpecRoute) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorOriginationSpecRoute.
func (in *VoiceConnectorOriginationSpecRoute) DeepCopy() *VoiceConnectorOriginationSpecRoute {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorOriginationSpecRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorOriginationStatus) DeepCopyInto(out *VoiceConnectorOriginationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorOriginationStatus.
func (in *VoiceConnectorOriginationStatus) DeepCopy() *VoiceConnectorOriginationStatus {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorOriginationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorSpec) DeepCopyInto(out *VoiceConnectorSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VoiceConnectorSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorSpec.
func (in *VoiceConnectorSpec) DeepCopy() *VoiceConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorSpecResource) DeepCopyInto(out *VoiceConnectorSpecResource) {
	*out = *in
	if in.AwsRegion != nil {
		in, out := &in.AwsRegion, &out.AwsRegion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OutboundHostName != nil {
		in, out := &in.OutboundHostName, &out.OutboundHostName
		*out = new(string)
		**out = **in
	}
	if in.RequireEncryption != nil {
		in, out := &in.RequireEncryption, &out.RequireEncryption
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorSpecResource.
func (in *VoiceConnectorSpecResource) DeepCopy() *VoiceConnectorSpecResource {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorStatus) DeepCopyInto(out *VoiceConnectorStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorStatus.
func (in *VoiceConnectorStatus) DeepCopy() *VoiceConnectorStatus {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorStreaming) DeepCopyInto(out *VoiceConnectorStreaming) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorStreaming.
func (in *VoiceConnectorStreaming) DeepCopy() *VoiceConnectorStreaming {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorStreaming)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorStreaming) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorStreamingList) DeepCopyInto(out *VoiceConnectorStreamingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VoiceConnectorStreaming, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorStreamingList.
func (in *VoiceConnectorStreamingList) DeepCopy() *VoiceConnectorStreamingList {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorStreamingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorStreamingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorStreamingSpec) DeepCopyInto(out *VoiceConnectorStreamingSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VoiceConnectorStreamingSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorStreamingSpec.
func (in *VoiceConnectorStreamingSpec) DeepCopy() *VoiceConnectorStreamingSpec {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorStreamingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorStreamingSpecResource) DeepCopyInto(out *VoiceConnectorStreamingSpecResource) {
	*out = *in
	if in.DataRetention != nil {
		in, out := &in.DataRetention, &out.DataRetention
		*out = new(int64)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.StreamingNotificationTargets != nil {
		in, out := &in.StreamingNotificationTargets, &out.StreamingNotificationTargets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VoiceConnectorID != nil {
		in, out := &in.VoiceConnectorID, &out.VoiceConnectorID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorStreamingSpecResource.
func (in *VoiceConnectorStreamingSpecResource) DeepCopy() *VoiceConnectorStreamingSpecResource {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorStreamingSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorStreamingStatus) DeepCopyInto(out *VoiceConnectorStreamingStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorStreamingStatus.
func (in *VoiceConnectorStreamingStatus) DeepCopy() *VoiceConnectorStreamingStatus {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorStreamingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTermination) DeepCopyInto(out *VoiceConnectorTermination) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTermination.
func (in *VoiceConnectorTermination) DeepCopy() *VoiceConnectorTermination {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTermination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorTermination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationCredentials) DeepCopyInto(out *VoiceConnectorTerminationCredentials) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationCredentials.
func (in *VoiceConnectorTerminationCredentials) DeepCopy() *VoiceConnectorTerminationCredentials {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorTerminationCredentials) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationCredentialsList) DeepCopyInto(out *VoiceConnectorTerminationCredentialsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VoiceConnectorTerminationCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationCredentialsList.
func (in *VoiceConnectorTerminationCredentialsList) DeepCopy() *VoiceConnectorTerminationCredentialsList {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationCredentialsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorTerminationCredentialsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationCredentialsSpec) DeepCopyInto(out *VoiceConnectorTerminationCredentialsSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VoiceConnectorTerminationCredentialsSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationCredentialsSpec.
func (in *VoiceConnectorTerminationCredentialsSpec) DeepCopy() *VoiceConnectorTerminationCredentialsSpec {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationCredentialsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationCredentialsSpecCredentials) DeepCopyInto(out *VoiceConnectorTerminationCredentialsSpecCredentials) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationCredentialsSpecCredentials.
func (in *VoiceConnectorTerminationCredentialsSpecCredentials) DeepCopy() *VoiceConnectorTerminationCredentialsSpecCredentials {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationCredentialsSpecCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationCredentialsSpecResource) DeepCopyInto(out *VoiceConnectorTerminationCredentialsSpecResource) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]VoiceConnectorTerminationCredentialsSpecCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VoiceConnectorID != nil {
		in, out := &in.VoiceConnectorID, &out.VoiceConnectorID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationCredentialsSpecResource.
func (in *VoiceConnectorTerminationCredentialsSpecResource) DeepCopy() *VoiceConnectorTerminationCredentialsSpecResource {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationCredentialsSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationCredentialsStatus) DeepCopyInto(out *VoiceConnectorTerminationCredentialsStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationCredentialsStatus.
func (in *VoiceConnectorTerminationCredentialsStatus) DeepCopy() *VoiceConnectorTerminationCredentialsStatus {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationCredentialsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationList) DeepCopyInto(out *VoiceConnectorTerminationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VoiceConnectorTermination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationList.
func (in *VoiceConnectorTerminationList) DeepCopy() *VoiceConnectorTerminationList {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VoiceConnectorTerminationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationSpec) DeepCopyInto(out *VoiceConnectorTerminationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VoiceConnectorTerminationSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationSpec.
func (in *VoiceConnectorTerminationSpec) DeepCopy() *VoiceConnectorTerminationSpec {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationSpecResource) DeepCopyInto(out *VoiceConnectorTerminationSpecResource) {
	*out = *in
	if in.CallingRegions != nil {
		in, out := &in.CallingRegions, &out.CallingRegions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CidrAllowList != nil {
		in, out := &in.CidrAllowList, &out.CidrAllowList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CpsLimit != nil {
		in, out := &in.CpsLimit, &out.CpsLimit
		*out = new(int64)
		**out = **in
	}
	if in.DefaultPhoneNumber != nil {
		in, out := &in.DefaultPhoneNumber, &out.DefaultPhoneNumber
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.VoiceConnectorID != nil {
		in, out := &in.VoiceConnectorID, &out.VoiceConnectorID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationSpecResource.
func (in *VoiceConnectorTerminationSpecResource) DeepCopy() *VoiceConnectorTerminationSpecResource {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceConnectorTerminationStatus) DeepCopyInto(out *VoiceConnectorTerminationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceConnectorTerminationStatus.
func (in *VoiceConnectorTerminationStatus) DeepCopy() *VoiceConnectorTerminationStatus {
	if in == nil {
		return nil
	}
	out := new(VoiceConnectorTerminationStatus)
	in.DeepCopyInto(out)
	return out
}