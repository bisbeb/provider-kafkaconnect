/*
Copyright 2025 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ConnectorPluginParameters are the configurable fields of a ConnectorPlugin.
type ConnectorPluginParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// ConnectorPluginObservation are the observable fields of a ConnectorPlugin.
type ConnectorPluginObservation struct {
	ConfigurableField string `json:"configurableField"`
	ObservableField   string `json:"observableField,omitempty"`
}

// A ConnectorPluginSpec defines the desired state of a ConnectorPlugin.
type ConnectorPluginSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConnectorPluginParameters `json:"forProvider"`
}

// A ConnectorPluginStatus represents the observed state of a ConnectorPlugin.
type ConnectorPluginStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConnectorPluginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A ConnectorPlugin is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,kafkaconnect}
type ConnectorPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectorPluginSpec   `json:"spec"`
	Status ConnectorPluginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectorPluginList contains a list of ConnectorPlugin
type ConnectorPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectorPlugin `json:"items"`
}

// ConnectorPlugin type metadata.
var (
	ConnectorPluginKind             = reflect.TypeOf(ConnectorPlugin{}).Name()
	ConnectorPluginGroupKind        = schema.GroupKind{Group: Group, Kind: ConnectorPluginKind}.String()
	ConnectorPluginKindAPIVersion   = ConnectorPluginKind + "." + SchemeGroupVersion.String()
	ConnectorPluginGroupVersionKind = SchemeGroupVersion.WithKind(ConnectorPluginKind)
)

func init() {
	SchemeBuilder.Register(&ConnectorPlugin{}, &ConnectorPluginList{})
}
