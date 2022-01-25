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

type CiliumObservation struct {
}

type CiliumParameters struct {
}

type ClusterObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KMSProvider []KMSProviderObservation `json:"kmsProvider,omitempty" tf:"kms_provider,omitempty"`

	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	Master []MasterObservation `json:"master,omitempty" tf:"master,omitempty"`

	NetworkImplementation []NetworkImplementationObservation `json:"networkImplementation,omitempty" tf:"network_implementation,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	ClusterIPv4Range *string `json:"clusterIpv4Range,omitempty" tf:"cluster_ipv4_range,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIPv6Range *string `json:"clusterIpv6Range,omitempty" tf:"cluster_ipv6_range,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSProvider []KMSProviderParameters `json:"kmsProvider,omitempty" tf:"kms_provider,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	Master []MasterParameters `json:"master" tf:"master,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkImplementation []NetworkImplementationParameters `json:"networkImplementation,omitempty" tf:"network_implementation,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkPolicyProvider *string `json:"networkPolicyProvider,omitempty" tf:"network_policy_provider,omitempty"`

	// +kubebuilder:validation:Optional
	NodeIPv4CidrMaskSize *int64 `json:"nodeIpv4CidrMaskSize,omitempty" tf:"node_ipv4_cidr_mask_size,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/iam/v1alpha1.ServiceAccount
	// +kubebuilder:validation:Optional
	NodeServiceAccountID *string `json:"nodeServiceAccountId,omitempty" tf:"node_service_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	NodeServiceAccountIDRef *v1.Reference `json:"nodeServiceAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NodeServiceAccountIDSelector *v1.Selector `json:"nodeServiceAccountIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ReleaseChannel *string `json:"releaseChannel,omitempty" tf:"release_channel,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/iam/v1alpha1.ServiceAccount
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceIPv4Range *string `json:"serviceIpv4Range,omitempty" tf:"service_ipv4_range,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceIPv6Range *string `json:"serviceIpv6Range,omitempty" tf:"service_ipv6_range,omitempty"`
}

type KMSProviderObservation struct {
}

type KMSProviderParameters struct {

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/kms/v1alpha1.SymmetricKey
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`
}

type LocationObservation struct {
}

type LocationParameters struct {

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type MaintenancePolicyObservation struct {
	MaintenanceWindow []MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
}

type MaintenancePolicyParameters struct {

	// +kubebuilder:validation:Required
	AutoUpgrade *bool `json:"autoUpgrade" tf:"auto_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
}

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// +kubebuilder:validation:Required
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type MasterObservation struct {
	ClusterCACertificate *string `json:"clusterCaCertificate,omitempty" tf:"cluster_ca_certificate,omitempty"`

	ExternalV4Address *string `json:"externalV4Address,omitempty" tf:"external_v4_address,omitempty"`

	ExternalV4Endpoint *string `json:"externalV4Endpoint,omitempty" tf:"external_v4_endpoint,omitempty"`

	InternalV4Address *string `json:"internalV4Address,omitempty" tf:"internal_v4_address,omitempty"`

	InternalV4Endpoint *string `json:"internalV4Endpoint,omitempty" tf:"internal_v4_endpoint,omitempty"`

	MaintenancePolicy []MaintenancePolicyObservation `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	Regional []RegionalObservation `json:"regional,omitempty" tf:"regional,omitempty"`

	VersionInfo []VersionInfoObservation `json:"versionInfo,omitempty" tf:"version_info,omitempty"`

	Zonal []ZonalObservation `json:"zonal,omitempty" tf:"zonal,omitempty"`
}

type MasterParameters struct {

	// +kubebuilder:validation:Optional
	MaintenancePolicy []MaintenancePolicyParameters `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIP *bool `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Regional []RegionalParameters `json:"regional,omitempty" tf:"regional,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// +kubebuilder:validation:Required
	VersionInfo []VersionInfoParameters `json:"versionInfo" tf:"version_info,omitempty"`

	// +kubebuilder:validation:Optional
	Zonal []ZonalParameters `json:"zonal,omitempty" tf:"zonal,omitempty"`
}

type NetworkImplementationObservation struct {
	Cilium []CiliumObservation `json:"cilium,omitempty" tf:"cilium,omitempty"`
}

type NetworkImplementationParameters struct {

	// +kubebuilder:validation:Optional
	Cilium []CiliumParameters `json:"cilium,omitempty" tf:"cilium,omitempty"`
}

type RegionalObservation struct {
	Location []LocationObservation `json:"location,omitempty" tf:"location,omitempty"`
}

type RegionalParameters struct {

	// +kubebuilder:validation:Optional
	Location []LocationParameters `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type VersionInfoObservation struct {
	CurrentVersion *string `json:"currentVersion,omitempty" tf:"current_version,omitempty"`

	NewRevisionAvailable *bool `json:"newRevisionAvailable,omitempty" tf:"new_revision_available,omitempty"`

	NewRevisionSummary *string `json:"newRevisionSummary,omitempty" tf:"new_revision_summary,omitempty"`

	VersionDeprecated *bool `json:"versionDeprecated,omitempty" tf:"version_deprecated,omitempty"`
}

type VersionInfoParameters struct {
}

type ZonalObservation struct {
}

type ZonalParameters struct {

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
