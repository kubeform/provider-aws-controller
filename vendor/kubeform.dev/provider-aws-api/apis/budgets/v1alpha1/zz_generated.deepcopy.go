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
func (in *Budget) DeepCopyInto(out *Budget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Budget.
func (in *Budget) DeepCopy() *Budget {
	if in == nil {
		return nil
	}
	out := new(Budget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Budget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetAction) DeepCopyInto(out *BudgetAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetAction.
func (in *BudgetAction) DeepCopy() *BudgetAction {
	if in == nil {
		return nil
	}
	out := new(BudgetAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionList) DeepCopyInto(out *BudgetActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BudgetAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionList.
func (in *BudgetActionList) DeepCopy() *BudgetActionList {
	if in == nil {
		return nil
	}
	out := new(BudgetActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionSpec) DeepCopyInto(out *BudgetActionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BudgetActionSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionSpec.
func (in *BudgetActionSpec) DeepCopy() *BudgetActionSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionSpecActionThreshold) DeepCopyInto(out *BudgetActionSpecActionThreshold) {
	*out = *in
	if in.ActionThresholdType != nil {
		in, out := &in.ActionThresholdType, &out.ActionThresholdType
		*out = new(string)
		**out = **in
	}
	if in.ActionThresholdValue != nil {
		in, out := &in.ActionThresholdValue, &out.ActionThresholdValue
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionSpecActionThreshold.
func (in *BudgetActionSpecActionThreshold) DeepCopy() *BudgetActionSpecActionThreshold {
	if in == nil {
		return nil
	}
	out := new(BudgetActionSpecActionThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionSpecDefinition) DeepCopyInto(out *BudgetActionSpecDefinition) {
	*out = *in
	if in.IamActionDefinition != nil {
		in, out := &in.IamActionDefinition, &out.IamActionDefinition
		*out = new(BudgetActionSpecDefinitionIamActionDefinition)
		(*in).DeepCopyInto(*out)
	}
	if in.ScpActionDefinition != nil {
		in, out := &in.ScpActionDefinition, &out.ScpActionDefinition
		*out = new(BudgetActionSpecDefinitionScpActionDefinition)
		(*in).DeepCopyInto(*out)
	}
	if in.SsmActionDefinition != nil {
		in, out := &in.SsmActionDefinition, &out.SsmActionDefinition
		*out = new(BudgetActionSpecDefinitionSsmActionDefinition)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionSpecDefinition.
func (in *BudgetActionSpecDefinition) DeepCopy() *BudgetActionSpecDefinition {
	if in == nil {
		return nil
	}
	out := new(BudgetActionSpecDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionSpecDefinitionIamActionDefinition) DeepCopyInto(out *BudgetActionSpecDefinitionIamActionDefinition) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PolicyArn != nil {
		in, out := &in.PolicyArn, &out.PolicyArn
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionSpecDefinitionIamActionDefinition.
func (in *BudgetActionSpecDefinitionIamActionDefinition) DeepCopy() *BudgetActionSpecDefinitionIamActionDefinition {
	if in == nil {
		return nil
	}
	out := new(BudgetActionSpecDefinitionIamActionDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionSpecDefinitionScpActionDefinition) DeepCopyInto(out *BudgetActionSpecDefinitionScpActionDefinition) {
	*out = *in
	if in.PolicyID != nil {
		in, out := &in.PolicyID, &out.PolicyID
		*out = new(string)
		**out = **in
	}
	if in.TargetIDS != nil {
		in, out := &in.TargetIDS, &out.TargetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionSpecDefinitionScpActionDefinition.
func (in *BudgetActionSpecDefinitionScpActionDefinition) DeepCopy() *BudgetActionSpecDefinitionScpActionDefinition {
	if in == nil {
		return nil
	}
	out := new(BudgetActionSpecDefinitionScpActionDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionSpecDefinitionSsmActionDefinition) DeepCopyInto(out *BudgetActionSpecDefinitionSsmActionDefinition) {
	*out = *in
	if in.ActionSubType != nil {
		in, out := &in.ActionSubType, &out.ActionSubType
		*out = new(string)
		**out = **in
	}
	if in.InstanceIDS != nil {
		in, out := &in.InstanceIDS, &out.InstanceIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionSpecDefinitionSsmActionDefinition.
func (in *BudgetActionSpecDefinitionSsmActionDefinition) DeepCopy() *BudgetActionSpecDefinitionSsmActionDefinition {
	if in == nil {
		return nil
	}
	out := new(BudgetActionSpecDefinitionSsmActionDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionSpecResource) DeepCopyInto(out *BudgetActionSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.ActionID != nil {
		in, out := &in.ActionID, &out.ActionID
		*out = new(string)
		**out = **in
	}
	if in.ActionThreshold != nil {
		in, out := &in.ActionThreshold, &out.ActionThreshold
		*out = new(BudgetActionSpecActionThreshold)
		(*in).DeepCopyInto(*out)
	}
	if in.ActionType != nil {
		in, out := &in.ActionType, &out.ActionType
		*out = new(string)
		**out = **in
	}
	if in.ApprovalModel != nil {
		in, out := &in.ApprovalModel, &out.ApprovalModel
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.BudgetName != nil {
		in, out := &in.BudgetName, &out.BudgetName
		*out = new(string)
		**out = **in
	}
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(BudgetActionSpecDefinition)
		(*in).DeepCopyInto(*out)
	}
	if in.ExecutionRoleArn != nil {
		in, out := &in.ExecutionRoleArn, &out.ExecutionRoleArn
		*out = new(string)
		**out = **in
	}
	if in.NotificationType != nil {
		in, out := &in.NotificationType, &out.NotificationType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Subscriber != nil {
		in, out := &in.Subscriber, &out.Subscriber
		*out = make([]BudgetActionSpecSubscriber, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionSpecResource.
func (in *BudgetActionSpecResource) DeepCopy() *BudgetActionSpecResource {
	if in == nil {
		return nil
	}
	out := new(BudgetActionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionSpecSubscriber) DeepCopyInto(out *BudgetActionSpecSubscriber) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionType != nil {
		in, out := &in.SubscriptionType, &out.SubscriptionType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionSpecSubscriber.
func (in *BudgetActionSpecSubscriber) DeepCopy() *BudgetActionSpecSubscriber {
	if in == nil {
		return nil
	}
	out := new(BudgetActionSpecSubscriber)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetActionStatus) DeepCopyInto(out *BudgetActionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetActionStatus.
func (in *BudgetActionStatus) DeepCopy() *BudgetActionStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetList) DeepCopyInto(out *BudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Budget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetList.
func (in *BudgetList) DeepCopy() *BudgetList {
	if in == nil {
		return nil
	}
	out := new(BudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSpec) DeepCopyInto(out *BudgetSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BudgetSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSpec.
func (in *BudgetSpec) DeepCopy() *BudgetSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSpecCostFilter) DeepCopyInto(out *BudgetSpecCostFilter) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSpecCostFilter.
func (in *BudgetSpecCostFilter) DeepCopy() *BudgetSpecCostFilter {
	if in == nil {
		return nil
	}
	out := new(BudgetSpecCostFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSpecCostTypes) DeepCopyInto(out *BudgetSpecCostTypes) {
	*out = *in
	if in.IncludeCredit != nil {
		in, out := &in.IncludeCredit, &out.IncludeCredit
		*out = new(bool)
		**out = **in
	}
	if in.IncludeDiscount != nil {
		in, out := &in.IncludeDiscount, &out.IncludeDiscount
		*out = new(bool)
		**out = **in
	}
	if in.IncludeOtherSubscription != nil {
		in, out := &in.IncludeOtherSubscription, &out.IncludeOtherSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRecurring != nil {
		in, out := &in.IncludeRecurring, &out.IncludeRecurring
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRefund != nil {
		in, out := &in.IncludeRefund, &out.IncludeRefund
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSubscription != nil {
		in, out := &in.IncludeSubscription, &out.IncludeSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSupport != nil {
		in, out := &in.IncludeSupport, &out.IncludeSupport
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTax != nil {
		in, out := &in.IncludeTax, &out.IncludeTax
		*out = new(bool)
		**out = **in
	}
	if in.IncludeUpfront != nil {
		in, out := &in.IncludeUpfront, &out.IncludeUpfront
		*out = new(bool)
		**out = **in
	}
	if in.UseAmortized != nil {
		in, out := &in.UseAmortized, &out.UseAmortized
		*out = new(bool)
		**out = **in
	}
	if in.UseBlended != nil {
		in, out := &in.UseBlended, &out.UseBlended
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSpecCostTypes.
func (in *BudgetSpecCostTypes) DeepCopy() *BudgetSpecCostTypes {
	if in == nil {
		return nil
	}
	out := new(BudgetSpecCostTypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSpecNotification) DeepCopyInto(out *BudgetSpecNotification) {
	*out = *in
	if in.ComparisonOperator != nil {
		in, out := &in.ComparisonOperator, &out.ComparisonOperator
		*out = new(string)
		**out = **in
	}
	if in.NotificationType != nil {
		in, out := &in.NotificationType, &out.NotificationType
		*out = new(string)
		**out = **in
	}
	if in.SubscriberEmailAddresses != nil {
		in, out := &in.SubscriberEmailAddresses, &out.SubscriberEmailAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubscriberSnsTopicArns != nil {
		in, out := &in.SubscriberSnsTopicArns, &out.SubscriberSnsTopicArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdType != nil {
		in, out := &in.ThresholdType, &out.ThresholdType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSpecNotification.
func (in *BudgetSpecNotification) DeepCopy() *BudgetSpecNotification {
	if in == nil {
		return nil
	}
	out := new(BudgetSpecNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSpecResource) DeepCopyInto(out *BudgetSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.BudgetType != nil {
		in, out := &in.BudgetType, &out.BudgetType
		*out = new(string)
		**out = **in
	}
	if in.CostFilter != nil {
		in, out := &in.CostFilter, &out.CostFilter
		*out = make([]BudgetSpecCostFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CostFilters != nil {
		in, out := &in.CostFilters, &out.CostFilters
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.CostTypes != nil {
		in, out := &in.CostTypes, &out.CostTypes
		*out = new(BudgetSpecCostTypes)
		(*in).DeepCopyInto(*out)
	}
	if in.LimitAmount != nil {
		in, out := &in.LimitAmount, &out.LimitAmount
		*out = new(string)
		**out = **in
	}
	if in.LimitUnit != nil {
		in, out := &in.LimitUnit, &out.LimitUnit
		*out = new(string)
		**out = **in
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
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]BudgetSpecNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimePeriodEnd != nil {
		in, out := &in.TimePeriodEnd, &out.TimePeriodEnd
		*out = new(string)
		**out = **in
	}
	if in.TimePeriodStart != nil {
		in, out := &in.TimePeriodStart, &out.TimePeriodStart
		*out = new(string)
		**out = **in
	}
	if in.TimeUnit != nil {
		in, out := &in.TimeUnit, &out.TimeUnit
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSpecResource.
func (in *BudgetSpecResource) DeepCopy() *BudgetSpecResource {
	if in == nil {
		return nil
	}
	out := new(BudgetSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetStatus) DeepCopyInto(out *BudgetStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetStatus.
func (in *BudgetStatus) DeepCopy() *BudgetStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetStatus)
	in.DeepCopyInto(out)
	return out
}
