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

type ServiceAccountIamPolicyObservation struct {
}

type ServiceAccountIamPolicyParameters struct {

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +kubebuilder:validation:Required
	ServiceAccountID *string `json:"serviceAccountId" tf:"service_account_id,omitempty"`
}

// ServiceAccountIamPolicySpec defines the desired state of ServiceAccountIamPolicy
type ServiceAccountIamPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountIamPolicyParameters `json:"forProvider"`
}

// ServiceAccountIamPolicyStatus defines the observed state of ServiceAccountIamPolicy.
type ServiceAccountIamPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountIamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountIamPolicy is the Schema for the ServiceAccountIamPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type ServiceAccountIamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountIamPolicySpec   `json:"spec"`
	Status            ServiceAccountIamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountIamPolicyList contains a list of ServiceAccountIamPolicys
type ServiceAccountIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountIamPolicy `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountIamPolicy_Kind             = "ServiceAccountIamPolicy"
	ServiceAccountIamPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountIamPolicy_Kind}.String()
	ServiceAccountIamPolicy_KindAPIVersion   = ServiceAccountIamPolicy_Kind + "." + CRDGroupVersion.String()
	ServiceAccountIamPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountIamPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountIamPolicy{}, &ServiceAccountIamPolicyList{})
}