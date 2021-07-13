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

type Budget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetSpec   `json:"spec,omitempty"`
	Status            BudgetStatus `json:"status,omitempty"`
}

type BudgetSpecCostTypes struct {
	// +optional
	IncludeCredit *bool `json:"includeCredit,omitempty" tf:"include_credit"`
	// +optional
	IncludeDiscount *bool `json:"includeDiscount,omitempty" tf:"include_discount"`
	// +optional
	IncludeOtherSubscription *bool `json:"includeOtherSubscription,omitempty" tf:"include_other_subscription"`
	// +optional
	IncludeRecurring *bool `json:"includeRecurring,omitempty" tf:"include_recurring"`
	// +optional
	IncludeRefund *bool `json:"includeRefund,omitempty" tf:"include_refund"`
	// +optional
	IncludeSubscription *bool `json:"includeSubscription,omitempty" tf:"include_subscription"`
	// +optional
	IncludeSupport *bool `json:"includeSupport,omitempty" tf:"include_support"`
	// +optional
	IncludeTax *bool `json:"includeTax,omitempty" tf:"include_tax"`
	// +optional
	IncludeUpfront *bool `json:"includeUpfront,omitempty" tf:"include_upfront"`
	// +optional
	UseAmortized *bool `json:"useAmortized,omitempty" tf:"use_amortized"`
	// +optional
	UseBlended *bool `json:"useBlended,omitempty" tf:"use_blended"`
}

type BudgetSpecNotification struct {
	ComparisonOperator *string `json:"comparisonOperator" tf:"comparison_operator"`
	NotificationType   *string `json:"notificationType" tf:"notification_type"`
	// +optional
	SubscriberEmailAddresses []string `json:"subscriberEmailAddresses,omitempty" tf:"subscriber_email_addresses"`
	// +optional
	SubscriberSnsTopicArns []string `json:"subscriberSnsTopicArns,omitempty" tf:"subscriber_sns_topic_arns"`
	Threshold              *float64 `json:"threshold" tf:"threshold"`
	ThresholdType          *string  `json:"thresholdType" tf:"threshold_type"`
}

type BudgetSpec struct {
	State *BudgetSpecResource `json:"state,omitempty" tf:"-"`

	Resource BudgetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BudgetSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// +optional
	Arn        *string `json:"arn,omitempty" tf:"arn"`
	BudgetType *string `json:"budgetType" tf:"budget_type"`
	// +optional
	CostFilters *map[string]string `json:"costFilters,omitempty" tf:"cost_filters"`
	// +optional
	CostTypes   *BudgetSpecCostTypes `json:"costTypes,omitempty" tf:"cost_types"`
	LimitAmount *string              `json:"limitAmount" tf:"limit_amount"`
	LimitUnit   *string              `json:"limitUnit" tf:"limit_unit"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	Notification []BudgetSpecNotification `json:"notification,omitempty" tf:"notification"`
	// +optional
	TimePeriodEnd   *string `json:"timePeriodEnd,omitempty" tf:"time_period_end"`
	TimePeriodStart *string `json:"timePeriodStart" tf:"time_period_start"`
	TimeUnit        *string `json:"timeUnit" tf:"time_unit"`
}

type BudgetStatus struct {
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

// BudgetList is a list of Budgets
type BudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Budget CRD objects
	Items []Budget `json:"items,omitempty"`
}
