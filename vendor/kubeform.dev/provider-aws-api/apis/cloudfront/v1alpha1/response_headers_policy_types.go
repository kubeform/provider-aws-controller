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

type ResponseHeadersPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResponseHeadersPolicySpec   `json:"spec,omitempty"`
	Status            ResponseHeadersPolicyStatus `json:"status,omitempty"`
}

type ResponseHeadersPolicySpecCorsConfigAccessControlAllowHeaders struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type ResponseHeadersPolicySpecCorsConfigAccessControlAllowMethods struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type ResponseHeadersPolicySpecCorsConfigAccessControlAllowOrigins struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type ResponseHeadersPolicySpecCorsConfigAccessControlExposeHeaders struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type ResponseHeadersPolicySpecCorsConfig struct {
	AccessControlAllowCredentials *bool                                                         `json:"accessControlAllowCredentials" tf:"access_control_allow_credentials"`
	AccessControlAllowHeaders     *ResponseHeadersPolicySpecCorsConfigAccessControlAllowHeaders `json:"accessControlAllowHeaders" tf:"access_control_allow_headers"`
	AccessControlAllowMethods     *ResponseHeadersPolicySpecCorsConfigAccessControlAllowMethods `json:"accessControlAllowMethods" tf:"access_control_allow_methods"`
	AccessControlAllowOrigins     *ResponseHeadersPolicySpecCorsConfigAccessControlAllowOrigins `json:"accessControlAllowOrigins" tf:"access_control_allow_origins"`
	// +optional
	AccessControlExposeHeaders *ResponseHeadersPolicySpecCorsConfigAccessControlExposeHeaders `json:"accessControlExposeHeaders,omitempty" tf:"access_control_expose_headers"`
	// +optional
	AccessControlMaxAgeSec *int64 `json:"accessControlMaxAgeSec,omitempty" tf:"access_control_max_age_sec"`
	OriginOverride         *bool  `json:"originOverride" tf:"origin_override"`
}

type ResponseHeadersPolicySpecCustomHeadersConfigItems struct {
	Header   *string `json:"header" tf:"header"`
	Override *bool   `json:"override" tf:"override"`
	Value    *string `json:"value" tf:"value"`
}

type ResponseHeadersPolicySpecCustomHeadersConfig struct {
	// +optional
	Items []ResponseHeadersPolicySpecCustomHeadersConfigItems `json:"items,omitempty" tf:"items"`
}

type ResponseHeadersPolicySpecSecurityHeadersConfigContentSecurityPolicy struct {
	ContentSecurityPolicy *string `json:"contentSecurityPolicy" tf:"content_security_policy"`
	Override              *bool   `json:"override" tf:"override"`
}

type ResponseHeadersPolicySpecSecurityHeadersConfigContentTypeOptions struct {
	Override *bool `json:"override" tf:"override"`
}

type ResponseHeadersPolicySpecSecurityHeadersConfigFrameOptions struct {
	FrameOption *string `json:"frameOption" tf:"frame_option"`
	Override    *bool   `json:"override" tf:"override"`
}

type ResponseHeadersPolicySpecSecurityHeadersConfigReferrerPolicy struct {
	Override       *bool   `json:"override" tf:"override"`
	ReferrerPolicy *string `json:"referrerPolicy" tf:"referrer_policy"`
}

type ResponseHeadersPolicySpecSecurityHeadersConfigStrictTransportSecurity struct {
	AccessControlMaxAgeSec *int64 `json:"accessControlMaxAgeSec" tf:"access_control_max_age_sec"`
	// +optional
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty" tf:"include_subdomains"`
	Override          *bool `json:"override" tf:"override"`
	// +optional
	Preload *bool `json:"preload,omitempty" tf:"preload"`
}

type ResponseHeadersPolicySpecSecurityHeadersConfigXssProtection struct {
	// +optional
	ModeBlock  *bool `json:"modeBlock,omitempty" tf:"mode_block"`
	Override   *bool `json:"override" tf:"override"`
	Protection *bool `json:"protection" tf:"protection"`
	// +optional
	ReportURI *string `json:"reportURI,omitempty" tf:"report_uri"`
}

type ResponseHeadersPolicySpecSecurityHeadersConfig struct {
	// +optional
	ContentSecurityPolicy *ResponseHeadersPolicySpecSecurityHeadersConfigContentSecurityPolicy `json:"contentSecurityPolicy,omitempty" tf:"content_security_policy"`
	// +optional
	ContentTypeOptions *ResponseHeadersPolicySpecSecurityHeadersConfigContentTypeOptions `json:"contentTypeOptions,omitempty" tf:"content_type_options"`
	// +optional
	FrameOptions *ResponseHeadersPolicySpecSecurityHeadersConfigFrameOptions `json:"frameOptions,omitempty" tf:"frame_options"`
	// +optional
	ReferrerPolicy *ResponseHeadersPolicySpecSecurityHeadersConfigReferrerPolicy `json:"referrerPolicy,omitempty" tf:"referrer_policy"`
	// +optional
	StrictTransportSecurity *ResponseHeadersPolicySpecSecurityHeadersConfigStrictTransportSecurity `json:"strictTransportSecurity,omitempty" tf:"strict_transport_security"`
	// +optional
	XssProtection *ResponseHeadersPolicySpecSecurityHeadersConfigXssProtection `json:"xssProtection,omitempty" tf:"xss_protection"`
}

type ResponseHeadersPolicySpec struct {
	State *ResponseHeadersPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource ResponseHeadersPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ResponseHeadersPolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Comment *string `json:"comment,omitempty" tf:"comment"`
	// +optional
	CorsConfig *ResponseHeadersPolicySpecCorsConfig `json:"corsConfig,omitempty" tf:"cors_config"`
	// +optional
	CustomHeadersConfig *ResponseHeadersPolicySpecCustomHeadersConfig `json:"customHeadersConfig,omitempty" tf:"custom_headers_config"`
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	Name *string `json:"name" tf:"name"`
	// +optional
	SecurityHeadersConfig *ResponseHeadersPolicySpecSecurityHeadersConfig `json:"securityHeadersConfig,omitempty" tf:"security_headers_config"`
}

type ResponseHeadersPolicyStatus struct {
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

// ResponseHeadersPolicyList is a list of ResponseHeadersPolicys
type ResponseHeadersPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ResponseHeadersPolicy CRD objects
	Items []ResponseHeadersPolicy `json:"items,omitempty"`
}