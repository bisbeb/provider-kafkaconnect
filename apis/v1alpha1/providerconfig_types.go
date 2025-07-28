package v1alpha1

import (
    "reflect"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime/schema"

    xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ProviderConfigSpec defines the desired state of ProviderConfig.
type ProviderConfigSpec struct {
    // Credentials required to authenticate to Kafka Connect.
    // +optional
    Credentials ProviderCredentials `json:"credentials"`
    
    // KafkaConnectURL is the base URL of the Kafka Connect instance
    // +kubebuilder:validation:Required
    KafkaConnectURL string `json:"kafkaConnectUrl"`
    
    // TLS configuration for connecting to Kafka Connect
    // +optional
    TLS *TLSConfig `json:"tls,omitempty"`
}

// TLSConfig contains TLS configuration
type TLSConfig struct {
    // InsecureSkipVerify disables TLS certificate verification
    // +optional
    InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
    
    // CABundle is a PEM encoded CA bundle which will be used to validate the server certificate
    // +optional
    CABundle []byte `json:"caBundle,omitempty"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
    // Source of the provider credentials.
    // +kubebuilder:validation:Enum=None;Secret;InjectedIdentity;Environment;Filesystem
    Source xpv1.CredentialsSource `json:"source"`

    xpv1.CommonCredentialSelectors `json:",inline"`
}

// ProviderConfigStatus defines the observed state of ProviderConfig.
type ProviderConfigStatus struct {
    xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProviderConfig is the Schema for the ProviderConfigs API.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,kafkaconnect}
// +kubebuilder:subresource:status
type ProviderConfig struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   ProviderConfigSpec   `json:"spec"`
    Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []ProviderConfig `json:"items"`
}

// ProviderConfig type metadata.
var (
    ProviderConfigKind             = reflect.TypeOf(ProviderConfig{}).Name()
    ProviderConfigGroupKind        = schema.GroupKind{Group: Group, Kind: ProviderConfigKind}.String()
    ProviderConfigKindAPIVersion   = ProviderConfigKind + "." + SchemeGroupVersion.String()
    ProviderConfigGroupVersionKind = SchemeGroupVersion.WithKind(ProviderConfigKind)
)

func init() {
    SchemeBuilder.Register(&ProviderConfig{}, &ProviderConfigList{})
}
