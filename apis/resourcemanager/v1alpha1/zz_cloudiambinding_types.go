/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CloudIamBindingObservation struct {
}

type CloudIamBindingParameters struct {

	// +kubebuilder:validation:Required
	CloudID *string `json:"cloudId" tf:"cloud_id,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// CloudIamBindingSpec defines the desired state of CloudIamBinding
type CloudIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudIamBindingParameters `json:"forProvider"`
}

// CloudIamBindingStatus defines the observed state of CloudIamBinding.
type CloudIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudIamBinding is the Schema for the CloudIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type CloudIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudIamBindingSpec   `json:"spec"`
	Status            CloudIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudIamBindingList contains a list of CloudIamBindings
type CloudIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIamBinding `json:"items"`
}

// Repository type metadata.
var (
	CloudIamBinding_Kind             = "CloudIamBinding"
	CloudIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudIamBinding_Kind}.String()
	CloudIamBinding_KindAPIVersion   = CloudIamBinding_Kind + "." + CRDGroupVersion.String()
	CloudIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(CloudIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudIamBinding{}, &CloudIamBindingList{})
}