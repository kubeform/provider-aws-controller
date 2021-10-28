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
func (in *CloudformationType) DeepCopyInto(out *CloudformationType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationType.
func (in *CloudformationType) DeepCopy() *CloudformationType {
	if in == nil {
		return nil
	}
	out := new(CloudformationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeList) DeepCopyInto(out *CloudformationTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudformationType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeList.
func (in *CloudformationTypeList) DeepCopy() *CloudformationTypeList {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeSpec) DeepCopyInto(out *CloudformationTypeSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CloudformationTypeSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeSpec.
func (in *CloudformationTypeSpec) DeepCopy() *CloudformationTypeSpec {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeSpecLoggingConfig) DeepCopyInto(out *CloudformationTypeSpecLoggingConfig) {
	*out = *in
	if in.LogGroupName != nil {
		in, out := &in.LogGroupName, &out.LogGroupName
		*out = new(string)
		**out = **in
	}
	if in.LogRoleArn != nil {
		in, out := &in.LogRoleArn, &out.LogRoleArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeSpecLoggingConfig.
func (in *CloudformationTypeSpecLoggingConfig) DeepCopy() *CloudformationTypeSpecLoggingConfig {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeSpecLoggingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeSpecResource) DeepCopyInto(out *CloudformationTypeSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DefaultVersionID != nil {
		in, out := &in.DefaultVersionID, &out.DefaultVersionID
		*out = new(string)
		**out = **in
	}
	if in.DeprecatedStatus != nil {
		in, out := &in.DeprecatedStatus, &out.DeprecatedStatus
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DocumentationURL != nil {
		in, out := &in.DocumentationURL, &out.DocumentationURL
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleArn != nil {
		in, out := &in.ExecutionRoleArn, &out.ExecutionRoleArn
		*out = new(string)
		**out = **in
	}
	if in.IsDefaultVersion != nil {
		in, out := &in.IsDefaultVersion, &out.IsDefaultVersion
		*out = new(bool)
		**out = **in
	}
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = new(CloudformationTypeSpecLoggingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningType != nil {
		in, out := &in.ProvisioningType, &out.ProvisioningType
		*out = new(string)
		**out = **in
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
	}
	if in.SchemaHandlerPackage != nil {
		in, out := &in.SchemaHandlerPackage, &out.SchemaHandlerPackage
		*out = new(string)
		**out = **in
	}
	if in.SourceURL != nil {
		in, out := &in.SourceURL, &out.SourceURL
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.TypeArn != nil {
		in, out := &in.TypeArn, &out.TypeArn
		*out = new(string)
		**out = **in
	}
	if in.TypeName != nil {
		in, out := &in.TypeName, &out.TypeName
		*out = new(string)
		**out = **in
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeSpecResource.
func (in *CloudformationTypeSpecResource) DeepCopy() *CloudformationTypeSpecResource {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeStatus) DeepCopyInto(out *CloudformationTypeStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeStatus.
func (in *CloudformationTypeStatus) DeepCopy() *CloudformationTypeStatus {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeStatus)
	in.DeepCopyInto(out)
	return out
}
