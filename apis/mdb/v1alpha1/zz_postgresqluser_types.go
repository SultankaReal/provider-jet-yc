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

type PostgresqlUserObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PostgresqlUserParameters struct {

	// +crossplane:generate:reference:type=PostgresqlCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnLimit *float64 `json:"connLimit,omitempty" tf:"conn_limit,omitempty"`

	// +kubebuilder:validation:Optional
	Grants []*string `json:"grants,omitempty" tf:"grants,omitempty"`

	// +kubebuilder:validation:Optional
	Login *bool `json:"login,omitempty" tf:"login,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Permission []PostgresqlUserPermissionParameters `json:"permission,omitempty" tf:"permission,omitempty"`

	// +kubebuilder:validation:Optional
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type PostgresqlUserPermissionObservation struct {
}

type PostgresqlUserPermissionParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`
}

// PostgresqlUserSpec defines the desired state of PostgresqlUser
type PostgresqlUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlUserParameters `json:"forProvider"`
}

// PostgresqlUserStatus defines the observed state of PostgresqlUser.
type PostgresqlUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlUser is the Schema for the PostgresqlUsers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type PostgresqlUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlUserSpec   `json:"spec"`
	Status            PostgresqlUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlUserList contains a list of PostgresqlUsers
type PostgresqlUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlUser `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlUser_Kind             = "PostgresqlUser"
	PostgresqlUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PostgresqlUser_Kind}.String()
	PostgresqlUser_KindAPIVersion   = PostgresqlUser_Kind + "." + CRDGroupVersion.String()
	PostgresqlUser_GroupVersionKind = CRDGroupVersion.WithKind(PostgresqlUser_Kind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlUser{}, &PostgresqlUserList{})
}
