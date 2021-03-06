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

type Listener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ListenerSpec   `json:"spec,omitempty"`
	Status            ListenerStatus `json:"status,omitempty"`
}

type ListenerSpecDefaultActionAuthenticateCognito struct {
	// +optional
	AuthenticationRequestExtraParams *map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params"`
	// +optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request"`
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// +optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name"`
	// +optional
	SessionTimeout   *int64  `json:"sessionTimeout,omitempty" tf:"session_timeout"`
	UserPoolArn      *string `json:"userPoolArn" tf:"user_pool_arn"`
	UserPoolClientID *string `json:"userPoolClientID" tf:"user_pool_client_id"`
	UserPoolDomain   *string `json:"userPoolDomain" tf:"user_pool_domain"`
}

type ListenerSpecDefaultActionAuthenticateOidc struct {
	// +optional
	AuthenticationRequestExtraParams *map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params"`
	AuthorizationEndpoint            *string            `json:"authorizationEndpoint" tf:"authorization_endpoint"`
	ClientID                         *string            `json:"clientID" tf:"client_id"`
	ClientSecret                     *string            `json:"-" sensitive:"true" tf:"client_secret"`
	Issuer                           *string            `json:"issuer" tf:"issuer"`
	// +optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request"`
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// +optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name"`
	// +optional
	SessionTimeout   *int64  `json:"sessionTimeout,omitempty" tf:"session_timeout"`
	TokenEndpoint    *string `json:"tokenEndpoint" tf:"token_endpoint"`
	UserInfoEndpoint *string `json:"userInfoEndpoint" tf:"user_info_endpoint"`
}

type ListenerSpecDefaultActionFixedResponse struct {
	ContentType *string `json:"contentType" tf:"content_type"`
	// +optional
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body"`
	// +optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code"`
}

type ListenerSpecDefaultActionForwardStickiness struct {
	Duration *int64 `json:"duration" tf:"duration"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type ListenerSpecDefaultActionForwardTargetGroup struct {
	Arn *string `json:"arn" tf:"arn"`
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type ListenerSpecDefaultActionForward struct {
	// +optional
	Stickiness *ListenerSpecDefaultActionForwardStickiness `json:"stickiness,omitempty" tf:"stickiness"`
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:MinItems=1
	TargetGroup []ListenerSpecDefaultActionForwardTargetGroup `json:"targetGroup" tf:"target_group"`
}

type ListenerSpecDefaultActionRedirect struct {
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Port *string `json:"port,omitempty" tf:"port"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	Query      *string `json:"query,omitempty" tf:"query"`
	StatusCode *string `json:"statusCode" tf:"status_code"`
}

type ListenerSpecDefaultAction struct {
	// +optional
	AuthenticateCognito *ListenerSpecDefaultActionAuthenticateCognito `json:"authenticateCognito,omitempty" tf:"authenticate_cognito"`
	// +optional
	AuthenticateOidc *ListenerSpecDefaultActionAuthenticateOidc `json:"authenticateOidc,omitempty" tf:"authenticate_oidc"`
	// +optional
	FixedResponse *ListenerSpecDefaultActionFixedResponse `json:"fixedResponse,omitempty" tf:"fixed_response"`
	// +optional
	Forward *ListenerSpecDefaultActionForward `json:"forward,omitempty" tf:"forward"`
	// +optional
	Order *int64 `json:"order,omitempty" tf:"order"`
	// +optional
	Redirect *ListenerSpecDefaultActionRedirect `json:"redirect,omitempty" tf:"redirect"`
	// +optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn"`
	Type           *string `json:"type" tf:"type"`
}

type ListenerSpec struct {
	State *ListenerSpecResource `json:"state,omitempty" tf:"-"`

	Resource ListenerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ListenerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AlpnPolicy *string `json:"alpnPolicy,omitempty" tf:"alpn_policy"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CertificateArn  *string                     `json:"certificateArn,omitempty" tf:"certificate_arn"`
	DefaultAction   []ListenerSpecDefaultAction `json:"defaultAction" tf:"default_action"`
	LoadBalancerArn *string                     `json:"loadBalancerArn" tf:"load_balancer_arn"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	SslPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ListenerStatus struct {
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

// ListenerList is a list of Listeners
type ListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Listener CRD objects
	Items []Listener `json:"items,omitempty"`
}
