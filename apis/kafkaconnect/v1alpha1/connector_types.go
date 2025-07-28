package v1alpha1

import (
    "reflect"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime/schema"

    xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ConnectorParameters are the configurable fields of a Connector.
type ConnectorParameters struct {
    // Name of the connector
    Name string `json:"name"`
    
    // ConnectorClass is the Java class for the connector
    // +kubebuilder:validation:Required
    ConnectorClass string `json:"connectorClass"`
    
    // TasksMax is the maximum number of tasks
    // +kubebuilder:default=1
    TasksMax int `json:"tasksMax,omitempty"`
    
    // Config contains connector-specific configuration
    // +kubebuilder:pruning:PreserveUnknownFields
    Config map[string]string `json:"config"`
    
    // KafkaConnectURL is the URL of the Kafka Connect instance
    // +optional
    KafkaConnectURL string `json:"kafkaConnectUrl,omitempty"`
}

// ConnectorObservation are the observable fields of a Connector.
type ConnectorObservation struct {
    // State of the connector
    State string `json:"state,omitempty"`
    
    // WorkerID that the connector is running on
    WorkerID string `json:"workerId,omitempty"`
    
    // Tasks information
    Tasks []TaskStatus `json:"tasks,omitempty"`
}

// TaskStatus represents the status of a connector task
type TaskStatus struct {
    ID       int    `json:"id"`
    State    string `json:"state"`
    WorkerID string `json:"workerId,omitempty"`
    Trace    string `json:"trace,omitempty"`
}

// A ConnectorSpec defines the desired state of a Connector.
type ConnectorSpec struct {
    xpv1.ResourceSpec `json:",inline"`
    ForProvider       ConnectorParameters `json:"forProvider"`
}

// A ConnectorStatus represents the observed state of a Connector.
type ConnectorStatus struct {
    xpv1.ResourceStatus `json:",inline"`
    AtProvider          ConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,kafkaconnect}

// A Connector is a managed resource representing a Kafka Connect connector
type Connector struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   ConnectorSpec   `json:"spec"`
    Status ConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectorList contains a list of Connector
type ConnectorList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []Connector `json:"items"`
}

// Connector type metadata.
var (
    ConnectorKind             = reflect.TypeOf(Connector{}).Name()
    ConnectorGroupKind        = schema.GroupKind{Group: Group, Kind: ConnectorKind}.String()
    ConnectorKindAPIVersion   = ConnectorKind + "." + SchemeGroupVersion.String()
    ConnectorGroupVersionKind = SchemeGroupVersion.WithKind(ConnectorKind)
)

func init() {
    SchemeBuilder.Register(&Connector{}, &ConnectorList{})
}

// GetCondition of this Connector.
func (mg *Connector) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
    return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Connector.
func (mg *Connector) GetDeletionPolicy() xpv1.DeletionPolicy {
    return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Connector.
func (mg *Connector) GetManagementPolicies() xpv1.ManagementPolicies {
    return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Connector.
func (mg *Connector) GetProviderConfigReference() *xpv1.Reference {
    return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Connector.
func (mg *Connector) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
    return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Connector.
func (mg *Connector) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
    return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Connector.
func (mg *Connector) SetConditions(c ...xpv1.Condition) {
    mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Connector.
func (mg *Connector) SetDeletionPolicy(r xpv1.DeletionPolicy) {
    mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Connector.
func (mg *Connector) SetManagementPolicies(r xpv1.ManagementPolicies) {
    mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Connector.
func (mg *Connector) SetProviderConfigReference(r *xpv1.Reference) {
    mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Connector.
func (mg *Connector) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
    mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Connector.
func (mg *Connector) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
    mg.Spec.WriteConnectionSecretToReference = r
}
