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

type VirtualGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualGatewaySpec   `json:"spec,omitempty"`
	Status            VirtualGatewayStatus `json:"status,omitempty"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsCertificateFile struct {
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain"`
	PrivateKey       *string `json:"privateKey" tf:"private_key"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsCertificateSds struct {
	SecretName *string `json:"secretName" tf:"secret_name"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsCertificate struct {
	// +optional
	File *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsCertificateFile `json:"file,omitempty" tf:"file"`
	// +optional
	Sds *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsCertificateSds `json:"sds,omitempty" tf:"sds"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch struct {
	Exact []string `json:"exact" tf:"exact"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames struct {
	Match *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch `json:"match" tf:"match"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationTrustAcm struct {
	CertificateAuthorityArns []string `json:"certificateAuthorityArns" tf:"certificate_authority_arns"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationTrustFile struct {
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationTrustSds struct {
	SecretName *string `json:"secretName" tf:"secret_name"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationTrust struct {
	// +optional
	Acm *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationTrustAcm `json:"acm,omitempty" tf:"acm"`
	// +optional
	File *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationTrustFile `json:"file,omitempty" tf:"file"`
	// +optional
	Sds *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationTrustSds `json:"sds,omitempty" tf:"sds"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidation struct {
	// +optional
	SubjectAlternativeNames *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names"`
	Trust                   *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidationTrust                   `json:"trust" tf:"trust"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicyTls struct {
	// +optional
	Certificate *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsCertificate `json:"certificate,omitempty" tf:"certificate"`
	// +optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce"`
	// +optional
	Ports      []int64                                                         `json:"ports,omitempty" tf:"ports"`
	Validation *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTlsValidation `json:"validation" tf:"validation"`
}

type VirtualGatewaySpecSpecBackendDefaultsClientPolicy struct {
	// +optional
	Tls *VirtualGatewaySpecSpecBackendDefaultsClientPolicyTls `json:"tls,omitempty" tf:"tls"`
}

type VirtualGatewaySpecSpecBackendDefaults struct {
	// +optional
	ClientPolicy *VirtualGatewaySpecSpecBackendDefaultsClientPolicy `json:"clientPolicy,omitempty" tf:"client_policy"`
}

type VirtualGatewaySpecSpecListenerConnectionPoolGrpc struct {
	MaxRequests *int64 `json:"maxRequests" tf:"max_requests"`
}

type VirtualGatewaySpecSpecListenerConnectionPoolHttp struct {
	MaxConnections *int64 `json:"maxConnections" tf:"max_connections"`
	// +optional
	MaxPendingRequests *int64 `json:"maxPendingRequests,omitempty" tf:"max_pending_requests"`
}

type VirtualGatewaySpecSpecListenerConnectionPoolHttp2 struct {
	MaxRequests *int64 `json:"maxRequests" tf:"max_requests"`
}

type VirtualGatewaySpecSpecListenerConnectionPool struct {
	// +optional
	Grpc *VirtualGatewaySpecSpecListenerConnectionPoolGrpc `json:"grpc,omitempty" tf:"grpc"`
	// +optional
	Http *VirtualGatewaySpecSpecListenerConnectionPoolHttp `json:"http,omitempty" tf:"http"`
	// +optional
	Http2 *VirtualGatewaySpecSpecListenerConnectionPoolHttp2 `json:"http2,omitempty" tf:"http2"`
}

type VirtualGatewaySpecSpecListenerHealthCheck struct {
	HealthyThreshold *int64 `json:"healthyThreshold" tf:"healthy_threshold"`
	IntervalMillis   *int64 `json:"intervalMillis" tf:"interval_millis"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Port               *int64  `json:"port,omitempty" tf:"port"`
	Protocol           *string `json:"protocol" tf:"protocol"`
	TimeoutMillis      *int64  `json:"timeoutMillis" tf:"timeout_millis"`
	UnhealthyThreshold *int64  `json:"unhealthyThreshold" tf:"unhealthy_threshold"`
}

type VirtualGatewaySpecSpecListenerPortMapping struct {
	Port     *int64  `json:"port" tf:"port"`
	Protocol *string `json:"protocol" tf:"protocol"`
}

type VirtualGatewaySpecSpecListenerTlsCertificateAcm struct {
	CertificateArn *string `json:"certificateArn" tf:"certificate_arn"`
}

type VirtualGatewaySpecSpecListenerTlsCertificateFile struct {
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain"`
	PrivateKey       *string `json:"privateKey" tf:"private_key"`
}

type VirtualGatewaySpecSpecListenerTlsCertificateSds struct {
	SecretName *string `json:"secretName" tf:"secret_name"`
}

type VirtualGatewaySpecSpecListenerTlsCertificate struct {
	// +optional
	Acm *VirtualGatewaySpecSpecListenerTlsCertificateAcm `json:"acm,omitempty" tf:"acm"`
	// +optional
	File *VirtualGatewaySpecSpecListenerTlsCertificateFile `json:"file,omitempty" tf:"file"`
	// +optional
	Sds *VirtualGatewaySpecSpecListenerTlsCertificateSds `json:"sds,omitempty" tf:"sds"`
}

type VirtualGatewaySpecSpecListenerTlsValidationSubjectAlternativeNamesMatch struct {
	Exact []string `json:"exact" tf:"exact"`
}

type VirtualGatewaySpecSpecListenerTlsValidationSubjectAlternativeNames struct {
	Match *VirtualGatewaySpecSpecListenerTlsValidationSubjectAlternativeNamesMatch `json:"match" tf:"match"`
}

type VirtualGatewaySpecSpecListenerTlsValidationTrustFile struct {
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain"`
}

type VirtualGatewaySpecSpecListenerTlsValidationTrustSds struct {
	SecretName *string `json:"secretName" tf:"secret_name"`
}

type VirtualGatewaySpecSpecListenerTlsValidationTrust struct {
	// +optional
	File *VirtualGatewaySpecSpecListenerTlsValidationTrustFile `json:"file,omitempty" tf:"file"`
	// +optional
	Sds *VirtualGatewaySpecSpecListenerTlsValidationTrustSds `json:"sds,omitempty" tf:"sds"`
}

type VirtualGatewaySpecSpecListenerTlsValidation struct {
	// +optional
	SubjectAlternativeNames *VirtualGatewaySpecSpecListenerTlsValidationSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names"`
	Trust                   *VirtualGatewaySpecSpecListenerTlsValidationTrust                   `json:"trust" tf:"trust"`
}

type VirtualGatewaySpecSpecListenerTls struct {
	Certificate *VirtualGatewaySpecSpecListenerTlsCertificate `json:"certificate" tf:"certificate"`
	Mode        *string                                       `json:"mode" tf:"mode"`
	// +optional
	Validation *VirtualGatewaySpecSpecListenerTlsValidation `json:"validation,omitempty" tf:"validation"`
}

type VirtualGatewaySpecSpecListener struct {
	// +optional
	ConnectionPool *VirtualGatewaySpecSpecListenerConnectionPool `json:"connectionPool,omitempty" tf:"connection_pool"`
	// +optional
	HealthCheck *VirtualGatewaySpecSpecListenerHealthCheck `json:"healthCheck,omitempty" tf:"health_check"`
	PortMapping *VirtualGatewaySpecSpecListenerPortMapping `json:"portMapping" tf:"port_mapping"`
	// +optional
	Tls *VirtualGatewaySpecSpecListenerTls `json:"tls,omitempty" tf:"tls"`
}

type VirtualGatewaySpecSpecLoggingAccessLogFile struct {
	Path *string `json:"path" tf:"path"`
}

type VirtualGatewaySpecSpecLoggingAccessLog struct {
	// +optional
	File *VirtualGatewaySpecSpecLoggingAccessLogFile `json:"file,omitempty" tf:"file"`
}

type VirtualGatewaySpecSpecLogging struct {
	// +optional
	AccessLog *VirtualGatewaySpecSpecLoggingAccessLog `json:"accessLog,omitempty" tf:"access_log"`
}

type VirtualGatewaySpecSpec struct {
	// +optional
	BackendDefaults *VirtualGatewaySpecSpecBackendDefaults `json:"backendDefaults,omitempty" tf:"backend_defaults"`
	Listener        *VirtualGatewaySpecSpecListener        `json:"listener" tf:"listener"`
	// +optional
	Logging *VirtualGatewaySpecSpecLogging `json:"logging,omitempty" tf:"logging"`
}

type VirtualGatewaySpec struct {
	State *VirtualGatewaySpecResource `json:"state,omitempty" tf:"-"`

	Resource VirtualGatewaySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VirtualGatewaySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date"`
	// +optional
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date"`
	MeshName        *string `json:"meshName" tf:"mesh_name"`
	// +optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner"`
	Name      *string `json:"name" tf:"name"`
	// +optional
	ResourceOwner *string                 `json:"resourceOwner,omitempty" tf:"resource_owner"`
	Spec          *VirtualGatewaySpecSpec `json:"spec" tf:"spec"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type VirtualGatewayStatus struct {
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

// VirtualGatewayList is a list of VirtualGateways
type VirtualGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualGateway CRD objects
	Items []VirtualGateway `json:"items,omitempty"`
}
