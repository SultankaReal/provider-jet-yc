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

type KafkaTopicObservation_2 struct {
}

type KafkaTopicParameters_2 struct {

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Required
	Partitions *int64 `json:"partitions" tf:"partitions,omitempty"`

	// +kubebuilder:validation:Required
	ReplicationFactor *int64 `json:"replicationFactor" tf:"replication_factor,omitempty"`

	// +kubebuilder:validation:Optional
	TopicConfig []KafkaTopicTopicConfigParameters `json:"topicConfig,omitempty" tf:"topic_config,omitempty"`
}

type KafkaTopicTopicConfigObservation struct {
}

type KafkaTopicTopicConfigParameters struct {

	// +kubebuilder:validation:Optional
	CleanupPolicy *string `json:"cleanupPolicy,omitempty" tf:"cleanup_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CompressionType *string `json:"compressionType,omitempty" tf:"compression_type,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteRetentionMs *string `json:"deleteRetentionMs,omitempty" tf:"delete_retention_ms,omitempty"`

	// +kubebuilder:validation:Optional
	FileDeleteDelayMs *string `json:"fileDeleteDelayMs,omitempty" tf:"file_delete_delay_ms,omitempty"`

	// +kubebuilder:validation:Optional
	FlushMessages *string `json:"flushMessages,omitempty" tf:"flush_messages,omitempty"`

	// +kubebuilder:validation:Optional
	FlushMs *string `json:"flushMs,omitempty" tf:"flush_ms,omitempty"`

	// +kubebuilder:validation:Optional
	MaxMessageBytes *string `json:"maxMessageBytes,omitempty" tf:"max_message_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	MinCompactionLagMs *string `json:"minCompactionLagMs,omitempty" tf:"min_compaction_lag_ms,omitempty"`

	// +kubebuilder:validation:Optional
	MinInsyncReplicas *string `json:"minInsyncReplicas,omitempty" tf:"min_insync_replicas,omitempty"`

	// +kubebuilder:validation:Optional
	Preallocate *bool `json:"preallocate,omitempty" tf:"preallocate,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionBytes *string `json:"retentionBytes,omitempty" tf:"retention_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionMs *string `json:"retentionMs,omitempty" tf:"retention_ms,omitempty"`

	// +kubebuilder:validation:Optional
	SegmentBytes *string `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`
}

// KafkaTopicSpec defines the desired state of KafkaTopic
type KafkaTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KafkaTopicParameters_2 `json:"forProvider"`
}

// KafkaTopicStatus defines the observed state of KafkaTopic.
type KafkaTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KafkaTopicObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaTopic is the Schema for the KafkaTopics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type KafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KafkaTopicSpec   `json:"spec"`
	Status            KafkaTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaTopicList contains a list of KafkaTopics
type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaTopic `json:"items"`
}

// Repository type metadata.
var (
	KafkaTopic_Kind             = "KafkaTopic"
	KafkaTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KafkaTopic_Kind}.String()
	KafkaTopic_KindAPIVersion   = KafkaTopic_Kind + "." + CRDGroupVersion.String()
	KafkaTopic_GroupVersionKind = CRDGroupVersion.WithKind(KafkaTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&KafkaTopic{}, &KafkaTopicList{})
}