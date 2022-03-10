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

type Workforce struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkforceSpec   `json:"spec,omitempty"`
	Status            WorkforceStatus `json:"status,omitempty"`
}

type WorkforceSpecCognitoConfig struct {
	ClientID *string `json:"clientID" tf:"client_id"`
	UserPool *string `json:"userPool" tf:"user_pool"`
}

type WorkforceSpecOidcConfig struct {
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint"`
	ClientID              *string `json:"clientID" tf:"client_id"`
	ClientSecret          *string `json:"-" sensitive:"true" tf:"client_secret"`
	Issuer                *string `json:"issuer" tf:"issuer"`
	JwksURI               *string `json:"jwksURI" tf:"jwks_uri"`
	LogoutEndpoint        *string `json:"logoutEndpoint" tf:"logout_endpoint"`
	TokenEndpoint         *string `json:"tokenEndpoint" tf:"token_endpoint"`
	UserInfoEndpoint      *string `json:"userInfoEndpoint" tf:"user_info_endpoint"`
}

type WorkforceSpecSourceIPConfig struct {
	// +kubebuilder:validation:MaxItems=10
	Cidrs []string `json:"cidrs" tf:"cidrs"`
}

type WorkforceSpec struct {
	State *WorkforceSpecResource `json:"state,omitempty" tf:"-"`

	Resource WorkforceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type WorkforceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CognitoConfig *WorkforceSpecCognitoConfig `json:"cognitoConfig,omitempty" tf:"cognito_config"`
	// +optional
	OidcConfig *WorkforceSpecOidcConfig `json:"oidcConfig,omitempty" tf:"oidc_config"`
	// +optional
	SourceIPConfig *WorkforceSpecSourceIPConfig `json:"sourceIPConfig,omitempty" tf:"source_ip_config"`
	// +optional
	Subdomain     *string `json:"subdomain,omitempty" tf:"subdomain"`
	WorkforceName *string `json:"workforceName" tf:"workforce_name"`
}

type WorkforceStatus struct {
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

// WorkforceList is a list of Workforces
type WorkforceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Workforce CRD objects
	Items []Workforce `json:"items,omitempty"`
}
