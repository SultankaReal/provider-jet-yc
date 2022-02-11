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

type AllocationPolicyLocationObservation struct {
}

type AllocationPolicyLocationParameters struct {

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

type AllocationPolicyObservation struct {
	Location []AllocationPolicyLocationObservation `json:"location,omitempty" tf:"location,omitempty"`
}

type AllocationPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Location []AllocationPolicyLocationParameters `json:"location,omitempty" tf:"location,omitempty"`
}

type AutoScaleObservation struct {
}

type AutoScaleParameters struct {

	// +kubebuilder:validation:Required
	Initial *int64 `json:"initial" tf:"initial,omitempty"`

	// +kubebuilder:validation:Required
	Max *int64 `json:"max" tf:"max,omitempty"`

	// +kubebuilder:validation:Required
	Min *int64 `json:"min" tf:"min,omitempty"`
}

type BootDiskObservation struct {
}

type BootDiskParameters struct {

	// +kubebuilder:validation:Optional
	Size *int64 `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ContainerRuntimeObservation struct {
}

type ContainerRuntimeParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type DeployPolicyObservation struct {
}

type DeployPolicyParameters struct {

	// +kubebuilder:validation:Required
	MaxExpansion *int64 `json:"maxExpansion" tf:"max_expansion,omitempty"`

	// +kubebuilder:validation:Required
	MaxUnavailable *int64 `json:"maxUnavailable" tf:"max_unavailable,omitempty"`
}

type FixedScaleObservation struct {
}

type FixedScaleParameters struct {

	// +kubebuilder:validation:Optional
	Size *int64 `json:"size,omitempty" tf:"size,omitempty"`
}

type InstanceTemplateObservation struct {
	BootDisk []BootDiskObservation `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	ContainerRuntime []ContainerRuntimeObservation `json:"containerRuntime,omitempty" tf:"container_runtime,omitempty"`

	NetworkInterface []NetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	PlacementPolicy []PlacementPolicyObservation `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	SchedulingPolicy []SchedulingPolicyObservation `json:"schedulingPolicy,omitempty" tf:"scheduling_policy,omitempty"`
}

type InstanceTemplateParameters struct {

	// +kubebuilder:validation:Optional
	BootDisk []BootDiskParameters `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerRuntime []ContainerRuntimeParameters `json:"containerRuntime,omitempty" tf:"container_runtime,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	NAT *bool `json:"nat,omitempty" tf:"nat,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkAccelerationType *string `json:"networkAccelerationType,omitempty" tf:"network_acceleration_type,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementPolicy []PlacementPolicyParameters `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	// +kubebuilder:validation:Optional
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// +kubebuilder:validation:Optional
	Resources []ResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	SchedulingPolicy []SchedulingPolicyParameters `json:"schedulingPolicy,omitempty" tf:"scheduling_policy,omitempty"`
}

type MaintenancePolicyMaintenanceWindowObservation struct {
}

type MaintenancePolicyMaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// +kubebuilder:validation:Required
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type NetworkInterfaceObservation struct {
}

type NetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	IPv4 *bool `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// +kubebuilder:validation:Optional
	NAT *bool `json:"nat,omitempty" tf:"nat,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`
}

type NodeGroupMaintenancePolicyObservation struct {
	MaintenanceWindow []MaintenancePolicyMaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
}

type NodeGroupMaintenancePolicyParameters struct {

	// +kubebuilder:validation:Required
	AutoRepair *bool `json:"autoRepair" tf:"auto_repair,omitempty"`

	// +kubebuilder:validation:Required
	AutoUpgrade *bool `json:"autoUpgrade" tf:"auto_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenancePolicyMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
}

type NodeGroupObservation struct {
	AllocationPolicy []AllocationPolicyObservation `json:"allocationPolicy,omitempty" tf:"allocation_policy,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	DeployPolicy []DeployPolicyObservation `json:"deployPolicy,omitempty" tf:"deploy_policy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceGroupID *string `json:"instanceGroupId,omitempty" tf:"instance_group_id,omitempty"`

	InstanceTemplate []InstanceTemplateObservation `json:"instanceTemplate,omitempty" tf:"instance_template,omitempty"`

	MaintenancePolicy []NodeGroupMaintenancePolicyObservation `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	ScalePolicy []ScalePolicyObservation `json:"scalePolicy,omitempty" tf:"scale_policy,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	VersionInfo []NodeGroupVersionInfoObservation `json:"versionInfo,omitempty" tf:"version_info,omitempty"`
}

type NodeGroupParameters struct {

	// +kubebuilder:validation:Optional
	AllocationPolicy []AllocationPolicyParameters `json:"allocationPolicy,omitempty" tf:"allocation_policy,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedUnsafeSysctls []*string `json:"allowedUnsafeSysctls,omitempty" tf:"allowed_unsafe_sysctls,omitempty"`

	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DeployPolicy []DeployPolicyParameters `json:"deployPolicy,omitempty" tf:"deploy_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	InstanceTemplate []InstanceTemplateParameters `json:"instanceTemplate" tf:"instance_template,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenancePolicy []NodeGroupMaintenancePolicyParameters `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NodeLabels map[string]*string `json:"nodeLabels,omitempty" tf:"node_labels,omitempty"`

	// +kubebuilder:validation:Optional
	NodeTaints []*string `json:"nodeTaints,omitempty" tf:"node_taints,omitempty"`

	// +kubebuilder:validation:Required
	ScalePolicy []ScalePolicyParameters `json:"scalePolicy" tf:"scale_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NodeGroupVersionInfoObservation struct {
	CurrentVersion *string `json:"currentVersion,omitempty" tf:"current_version,omitempty"`

	NewRevisionAvailable *bool `json:"newRevisionAvailable,omitempty" tf:"new_revision_available,omitempty"`

	NewRevisionSummary *string `json:"newRevisionSummary,omitempty" tf:"new_revision_summary,omitempty"`

	VersionDeprecated *bool `json:"versionDeprecated,omitempty" tf:"version_deprecated,omitempty"`
}

type NodeGroupVersionInfoParameters struct {
}

type PlacementPolicyObservation struct {
}

type PlacementPolicyParameters struct {

	// +kubebuilder:validation:Required
	PlacementGroupID *string `json:"placementGroupId" tf:"placement_group_id,omitempty"`
}

type ResourcesObservation struct {
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Optional
	CoreFraction *int64 `json:"coreFraction,omitempty" tf:"core_fraction,omitempty"`

	// +kubebuilder:validation:Optional
	Cores *int64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// +kubebuilder:validation:Optional
	Gpus *int64 `json:"gpus,omitempty" tf:"gpus,omitempty"`

	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`
}

type ScalePolicyObservation struct {
	AutoScale []AutoScaleObservation `json:"autoScale,omitempty" tf:"auto_scale,omitempty"`

	FixedScale []FixedScaleObservation `json:"fixedScale,omitempty" tf:"fixed_scale,omitempty"`
}

type ScalePolicyParameters struct {

	// +kubebuilder:validation:Optional
	AutoScale []AutoScaleParameters `json:"autoScale,omitempty" tf:"auto_scale,omitempty"`

	// +kubebuilder:validation:Optional
	FixedScale []FixedScaleParameters `json:"fixedScale,omitempty" tf:"fixed_scale,omitempty"`
}

type SchedulingPolicyObservation struct {
}

type SchedulingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

// NodeGroupSpec defines the desired state of NodeGroup
type NodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeGroupParameters `json:"forProvider"`
}

// NodeGroupStatus defines the observed state of NodeGroup.
type NodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroup is the Schema for the NodeGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type NodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeGroupSpec   `json:"spec"`
	Status            NodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroupList contains a list of NodeGroups
type NodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeGroup `json:"items"`
}

// Repository type metadata.
var (
	NodeGroup_Kind             = "NodeGroup"
	NodeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeGroup_Kind}.String()
	NodeGroup_KindAPIVersion   = NodeGroup_Kind + "." + CRDGroupVersion.String()
	NodeGroup_GroupVersionKind = CRDGroupVersion.WithKind(NodeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeGroup{}, &NodeGroupList{})
}
