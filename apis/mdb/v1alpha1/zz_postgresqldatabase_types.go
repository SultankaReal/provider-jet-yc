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

type PostgresqlDatabaseExtensionObservation struct {
}

type PostgresqlDatabaseExtensionParameters struct {

	// +kubebuilder:validation:Required
	// (Required) The name of the database.
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Version of the extension.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PostgresqlDatabaseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PostgresqlDatabaseParameters struct {

	// +crossplane:generate:reference:type=PostgresqlCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) Set of database extensions. The structure is documented below
	Extension []PostgresqlDatabaseExtensionParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) POSIX locale for string sorting order. Forbidden to change in an existing database.
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) POSIX locale for character classification. Forbidden to change in an existing database.
	LcType *string `json:"lcType,omitempty" tf:"lc_type,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The name of the database.
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Name of the user assigned as the owner of the database. Forbidden to change in an existing database.
	Owner *string `json:"owner" tf:"owner,omitempty"`
}

// PostgresqlDatabaseSpec defines the desired state of PostgresqlDatabase
type PostgresqlDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlDatabaseParameters `json:"forProvider"`
}

// PostgresqlDatabaseStatus defines the observed state of PostgresqlDatabase.
type PostgresqlDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlDatabase is the Schema for the PostgresqlDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type PostgresqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlDatabaseSpec   `json:"spec"`
	Status            PostgresqlDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlDatabaseList contains a list of PostgresqlDatabases
type PostgresqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlDatabase `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlDatabase_Kind             = "PostgresqlDatabase"
	PostgresqlDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PostgresqlDatabase_Kind}.String()
	PostgresqlDatabase_KindAPIVersion   = PostgresqlDatabase_Kind + "." + CRDGroupVersion.String()
	PostgresqlDatabase_GroupVersionKind = CRDGroupVersion.WithKind(PostgresqlDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlDatabase{}, &PostgresqlDatabaseList{})
}
