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
func (in *LicenseAssociation) DeepCopyInto(out *LicenseAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociation.
func (in *LicenseAssociation) DeepCopy() *LicenseAssociation {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationList) DeepCopyInto(out *LicenseAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationList.
func (in *LicenseAssociationList) DeepCopy() *LicenseAssociationList {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationSpec) DeepCopyInto(out *LicenseAssociationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LicenseAssociationSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationSpec.
func (in *LicenseAssociationSpec) DeepCopy() *LicenseAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationSpecResource) DeepCopyInto(out *LicenseAssociationSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.FreeTrialExpiration != nil {
		in, out := &in.FreeTrialExpiration, &out.FreeTrialExpiration
		*out = new(string)
		**out = **in
	}
	if in.LicenseExpiration != nil {
		in, out := &in.LicenseExpiration, &out.LicenseExpiration
		*out = new(string)
		**out = **in
	}
	if in.LicenseType != nil {
		in, out := &in.LicenseType, &out.LicenseType
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceID != nil {
		in, out := &in.WorkspaceID, &out.WorkspaceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationSpecResource.
func (in *LicenseAssociationSpecResource) DeepCopy() *LicenseAssociationSpecResource {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationStatus) DeepCopyInto(out *LicenseAssociationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationStatus.
func (in *LicenseAssociationStatus) DeepCopy() *LicenseAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpec) DeepCopyInto(out *WorkspaceSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(WorkspaceSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpec.
func (in *WorkspaceSpec) DeepCopy() *WorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpecResource) DeepCopyInto(out *WorkspaceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountAccessType != nil {
		in, out := &in.AccountAccessType, &out.AccountAccessType
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationProviders != nil {
		in, out := &in.AuthenticationProviders, &out.AuthenticationProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DataSources != nil {
		in, out := &in.DataSources, &out.DataSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.GrafanaVersion != nil {
		in, out := &in.GrafanaVersion, &out.GrafanaVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotificationDestinations != nil {
		in, out := &in.NotificationDestinations, &out.NotificationDestinations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OrganizationRoleName != nil {
		in, out := &in.OrganizationRoleName, &out.OrganizationRoleName
		*out = new(string)
		**out = **in
	}
	if in.OrganizationalUnits != nil {
		in, out := &in.OrganizationalUnits, &out.OrganizationalUnits
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PermissionType != nil {
		in, out := &in.PermissionType, &out.PermissionType
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.SamlConfigurationStatus != nil {
		in, out := &in.SamlConfigurationStatus, &out.SamlConfigurationStatus
		*out = new(string)
		**out = **in
	}
	if in.StackSetName != nil {
		in, out := &in.StackSetName, &out.StackSetName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpecResource.
func (in *WorkspaceSpecResource) DeepCopy() *WorkspaceSpecResource {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceStatus) DeepCopyInto(out *WorkspaceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceStatus.
func (in *WorkspaceStatus) DeepCopy() *WorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}
