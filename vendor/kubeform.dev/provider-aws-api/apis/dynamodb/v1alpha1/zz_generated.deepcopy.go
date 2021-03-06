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
func (in *GlobalTable) DeepCopyInto(out *GlobalTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTable.
func (in *GlobalTable) DeepCopy() *GlobalTable {
	if in == nil {
		return nil
	}
	out := new(GlobalTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTableList) DeepCopyInto(out *GlobalTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTableList.
func (in *GlobalTableList) DeepCopy() *GlobalTableList {
	if in == nil {
		return nil
	}
	out := new(GlobalTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTableSpec) DeepCopyInto(out *GlobalTableSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(GlobalTableSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTableSpec.
func (in *GlobalTableSpec) DeepCopy() *GlobalTableSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTableSpecReplica) DeepCopyInto(out *GlobalTableSpecReplica) {
	*out = *in
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTableSpecReplica.
func (in *GlobalTableSpecReplica) DeepCopy() *GlobalTableSpecReplica {
	if in == nil {
		return nil
	}
	out := new(GlobalTableSpecReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTableSpecResource) DeepCopyInto(out *GlobalTableSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Replica != nil {
		in, out := &in.Replica, &out.Replica
		*out = make([]GlobalTableSpecReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTableSpecResource.
func (in *GlobalTableSpecResource) DeepCopy() *GlobalTableSpecResource {
	if in == nil {
		return nil
	}
	out := new(GlobalTableSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTableStatus) DeepCopyInto(out *GlobalTableStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTableStatus.
func (in *GlobalTableStatus) DeepCopy() *GlobalTableStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamingDestination) DeepCopyInto(out *KinesisStreamingDestination) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamingDestination.
func (in *KinesisStreamingDestination) DeepCopy() *KinesisStreamingDestination {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamingDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KinesisStreamingDestination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamingDestinationList) DeepCopyInto(out *KinesisStreamingDestinationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KinesisStreamingDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamingDestinationList.
func (in *KinesisStreamingDestinationList) DeepCopy() *KinesisStreamingDestinationList {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamingDestinationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KinesisStreamingDestinationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamingDestinationSpec) DeepCopyInto(out *KinesisStreamingDestinationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(KinesisStreamingDestinationSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamingDestinationSpec.
func (in *KinesisStreamingDestinationSpec) DeepCopy() *KinesisStreamingDestinationSpec {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamingDestinationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamingDestinationSpecResource) DeepCopyInto(out *KinesisStreamingDestinationSpecResource) {
	*out = *in
	if in.StreamArn != nil {
		in, out := &in.StreamArn, &out.StreamArn
		*out = new(string)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamingDestinationSpecResource.
func (in *KinesisStreamingDestinationSpecResource) DeepCopy() *KinesisStreamingDestinationSpecResource {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamingDestinationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamingDestinationStatus) DeepCopyInto(out *KinesisStreamingDestinationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamingDestinationStatus.
func (in *KinesisStreamingDestinationStatus) DeepCopy() *KinesisStreamingDestinationStatus {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamingDestinationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Table) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableItem) DeepCopyInto(out *TableItem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableItem.
func (in *TableItem) DeepCopy() *TableItem {
	if in == nil {
		return nil
	}
	out := new(TableItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableItem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableItemList) DeepCopyInto(out *TableItemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TableItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableItemList.
func (in *TableItemList) DeepCopy() *TableItemList {
	if in == nil {
		return nil
	}
	out := new(TableItemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableItemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableItemSpec) DeepCopyInto(out *TableItemSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(TableItemSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableItemSpec.
func (in *TableItemSpec) DeepCopy() *TableItemSpec {
	if in == nil {
		return nil
	}
	out := new(TableItemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableItemSpecResource) DeepCopyInto(out *TableItemSpecResource) {
	*out = *in
	if in.HashKey != nil {
		in, out := &in.HashKey, &out.HashKey
		*out = new(string)
		**out = **in
	}
	if in.Item != nil {
		in, out := &in.Item, &out.Item
		*out = new(string)
		**out = **in
	}
	if in.RangeKey != nil {
		in, out := &in.RangeKey, &out.RangeKey
		*out = new(string)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableItemSpecResource.
func (in *TableItemSpecResource) DeepCopy() *TableItemSpecResource {
	if in == nil {
		return nil
	}
	out := new(TableItemSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableItemStatus) DeepCopyInto(out *TableItemStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableItemStatus.
func (in *TableItemStatus) DeepCopy() *TableItemStatus {
	if in == nil {
		return nil
	}
	out := new(TableItemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableList) DeepCopyInto(out *TableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Table, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableList.
func (in *TableList) DeepCopy() *TableList {
	if in == nil {
		return nil
	}
	out := new(TableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpec) DeepCopyInto(out *TableSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(TableSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpec.
func (in *TableSpec) DeepCopy() *TableSpec {
	if in == nil {
		return nil
	}
	out := new(TableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecAttribute) DeepCopyInto(out *TableSpecAttribute) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecAttribute.
func (in *TableSpecAttribute) DeepCopy() *TableSpecAttribute {
	if in == nil {
		return nil
	}
	out := new(TableSpecAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecGlobalSecondaryIndex) DeepCopyInto(out *TableSpecGlobalSecondaryIndex) {
	*out = *in
	if in.HashKey != nil {
		in, out := &in.HashKey, &out.HashKey
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NonKeyAttributes != nil {
		in, out := &in.NonKeyAttributes, &out.NonKeyAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProjectionType != nil {
		in, out := &in.ProjectionType, &out.ProjectionType
		*out = new(string)
		**out = **in
	}
	if in.RangeKey != nil {
		in, out := &in.RangeKey, &out.RangeKey
		*out = new(string)
		**out = **in
	}
	if in.ReadCapacity != nil {
		in, out := &in.ReadCapacity, &out.ReadCapacity
		*out = new(int64)
		**out = **in
	}
	if in.WriteCapacity != nil {
		in, out := &in.WriteCapacity, &out.WriteCapacity
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecGlobalSecondaryIndex.
func (in *TableSpecGlobalSecondaryIndex) DeepCopy() *TableSpecGlobalSecondaryIndex {
	if in == nil {
		return nil
	}
	out := new(TableSpecGlobalSecondaryIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecLocalSecondaryIndex) DeepCopyInto(out *TableSpecLocalSecondaryIndex) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NonKeyAttributes != nil {
		in, out := &in.NonKeyAttributes, &out.NonKeyAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProjectionType != nil {
		in, out := &in.ProjectionType, &out.ProjectionType
		*out = new(string)
		**out = **in
	}
	if in.RangeKey != nil {
		in, out := &in.RangeKey, &out.RangeKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecLocalSecondaryIndex.
func (in *TableSpecLocalSecondaryIndex) DeepCopy() *TableSpecLocalSecondaryIndex {
	if in == nil {
		return nil
	}
	out := new(TableSpecLocalSecondaryIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecPointInTimeRecovery) DeepCopyInto(out *TableSpecPointInTimeRecovery) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecPointInTimeRecovery.
func (in *TableSpecPointInTimeRecovery) DeepCopy() *TableSpecPointInTimeRecovery {
	if in == nil {
		return nil
	}
	out := new(TableSpecPointInTimeRecovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecReplica) DeepCopyInto(out *TableSpecReplica) {
	*out = *in
	if in.KmsKeyArn != nil {
		in, out := &in.KmsKeyArn, &out.KmsKeyArn
		*out = new(string)
		**out = **in
	}
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecReplica.
func (in *TableSpecReplica) DeepCopy() *TableSpecReplica {
	if in == nil {
		return nil
	}
	out := new(TableSpecReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecResource) DeepCopyInto(out *TableSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Attribute != nil {
		in, out := &in.Attribute, &out.Attribute
		*out = make([]TableSpecAttribute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BillingMode != nil {
		in, out := &in.BillingMode, &out.BillingMode
		*out = new(string)
		**out = **in
	}
	if in.GlobalSecondaryIndex != nil {
		in, out := &in.GlobalSecondaryIndex, &out.GlobalSecondaryIndex
		*out = make([]TableSpecGlobalSecondaryIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HashKey != nil {
		in, out := &in.HashKey, &out.HashKey
		*out = new(string)
		**out = **in
	}
	if in.LocalSecondaryIndex != nil {
		in, out := &in.LocalSecondaryIndex, &out.LocalSecondaryIndex
		*out = make([]TableSpecLocalSecondaryIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PointInTimeRecovery != nil {
		in, out := &in.PointInTimeRecovery, &out.PointInTimeRecovery
		*out = new(TableSpecPointInTimeRecovery)
		(*in).DeepCopyInto(*out)
	}
	if in.RangeKey != nil {
		in, out := &in.RangeKey, &out.RangeKey
		*out = new(string)
		**out = **in
	}
	if in.ReadCapacity != nil {
		in, out := &in.ReadCapacity, &out.ReadCapacity
		*out = new(int64)
		**out = **in
	}
	if in.Replica != nil {
		in, out := &in.Replica, &out.Replica
		*out = make([]TableSpecReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestoreDateTime != nil {
		in, out := &in.RestoreDateTime, &out.RestoreDateTime
		*out = new(string)
		**out = **in
	}
	if in.RestoreSourceName != nil {
		in, out := &in.RestoreSourceName, &out.RestoreSourceName
		*out = new(string)
		**out = **in
	}
	if in.RestoreToLatestTime != nil {
		in, out := &in.RestoreToLatestTime, &out.RestoreToLatestTime
		*out = new(bool)
		**out = **in
	}
	if in.ServerSideEncryption != nil {
		in, out := &in.ServerSideEncryption, &out.ServerSideEncryption
		*out = new(TableSpecServerSideEncryption)
		(*in).DeepCopyInto(*out)
	}
	if in.StreamArn != nil {
		in, out := &in.StreamArn, &out.StreamArn
		*out = new(string)
		**out = **in
	}
	if in.StreamEnabled != nil {
		in, out := &in.StreamEnabled, &out.StreamEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StreamLabel != nil {
		in, out := &in.StreamLabel, &out.StreamLabel
		*out = new(string)
		**out = **in
	}
	if in.StreamViewType != nil {
		in, out := &in.StreamViewType, &out.StreamViewType
		*out = new(string)
		**out = **in
	}
	if in.TableClass != nil {
		in, out := &in.TableClass, &out.TableClass
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
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(TableSpecTtl)
		(*in).DeepCopyInto(*out)
	}
	if in.WriteCapacity != nil {
		in, out := &in.WriteCapacity, &out.WriteCapacity
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecResource.
func (in *TableSpecResource) DeepCopy() *TableSpecResource {
	if in == nil {
		return nil
	}
	out := new(TableSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecServerSideEncryption) DeepCopyInto(out *TableSpecServerSideEncryption) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KmsKeyArn != nil {
		in, out := &in.KmsKeyArn, &out.KmsKeyArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecServerSideEncryption.
func (in *TableSpecServerSideEncryption) DeepCopy() *TableSpecServerSideEncryption {
	if in == nil {
		return nil
	}
	out := new(TableSpecServerSideEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecTtl) DeepCopyInto(out *TableSpecTtl) {
	*out = *in
	if in.AttributeName != nil {
		in, out := &in.AttributeName, &out.AttributeName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecTtl.
func (in *TableSpecTtl) DeepCopy() *TableSpecTtl {
	if in == nil {
		return nil
	}
	out := new(TableSpecTtl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableStatus) DeepCopyInto(out *TableStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableStatus.
func (in *TableStatus) DeepCopy() *TableStatus {
	if in == nil {
		return nil
	}
	out := new(TableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagList) DeepCopyInto(out *TagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagList.
func (in *TagList) DeepCopy() *TagList {
	if in == nil {
		return nil
	}
	out := new(TagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSpec) DeepCopyInto(out *TagSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(TagSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSpec.
func (in *TagSpec) DeepCopy() *TagSpec {
	if in == nil {
		return nil
	}
	out := new(TagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSpecResource) DeepCopyInto(out *TagSpecResource) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.ResourceArn != nil {
		in, out := &in.ResourceArn, &out.ResourceArn
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSpecResource.
func (in *TagSpecResource) DeepCopy() *TagSpecResource {
	if in == nil {
		return nil
	}
	out := new(TagSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagStatus) DeepCopyInto(out *TagStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagStatus.
func (in *TagStatus) DeepCopy() *TagStatus {
	if in == nil {
		return nil
	}
	out := new(TagStatus)
	in.DeepCopyInto(out)
	return out
}
