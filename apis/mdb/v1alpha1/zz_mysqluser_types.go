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

type MySQLUserConnectionLimitsObservation struct {
}

type MySQLUserConnectionLimitsParameters struct {

	// +kubebuilder:validation:Optional
	MaxConnectionsPerHour *float64 `json:"maxConnectionsPerHour,omitempty" tf:"max_connections_per_hour,omitempty"`

	// +kubebuilder:validation:Optional
	MaxQuestionsPerHour *float64 `json:"maxQuestionsPerHour,omitempty" tf:"max_questions_per_hour,omitempty"`

	// +kubebuilder:validation:Optional
	MaxUpdatesPerHour *float64 `json:"maxUpdatesPerHour,omitempty" tf:"max_updates_per_hour,omitempty"`

	// +kubebuilder:validation:Optional
	MaxUserConnections *float64 `json:"maxUserConnections,omitempty" tf:"max_user_connections,omitempty"`
}

type MySQLUserObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MySQLUserParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationPlugin *string `json:"authenticationPlugin,omitempty" tf:"authentication_plugin,omitempty"`

	// +crossplane:generate:reference:type=MysqlCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnectionLimits []MySQLUserConnectionLimitsParameters `json:"connectionLimits,omitempty" tf:"connection_limits,omitempty"`

	// +kubebuilder:validation:Optional
	GlobalPermissions []*string `json:"globalPermissions,omitempty" tf:"global_permissions,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Permission []MySQLUserPermissionParameters `json:"permission,omitempty" tf:"permission,omitempty"`
}

type MySQLUserPermissionObservation struct {
}

type MySQLUserPermissionParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

// MySQLUserSpec defines the desired state of MySQLUser
type MySQLUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MySQLUserParameters `json:"forProvider"`
}

// MySQLUserStatus defines the observed state of MySQLUser.
type MySQLUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MySQLUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLUser is the Schema for the MySQLUsers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type MySQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MySQLUserSpec   `json:"spec"`
	Status            MySQLUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLUserList contains a list of MySQLUsers
type MySQLUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLUser `json:"items"`
}

// Repository type metadata.
var (
	MySQLUser_Kind             = "MySQLUser"
	MySQLUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MySQLUser_Kind}.String()
	MySQLUser_KindAPIVersion   = MySQLUser_Kind + "." + CRDGroupVersion.String()
	MySQLUser_GroupVersionKind = CRDGroupVersion.WithKind(MySQLUser_Kind)
)

func init() {
	SchemeBuilder.Register(&MySQLUser{}, &MySQLUserList{})
}
