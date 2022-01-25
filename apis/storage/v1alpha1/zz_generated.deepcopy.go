//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyServerSideEncryptionByDefaultObservation) DeepCopyInto(out *ApplyServerSideEncryptionByDefaultObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyServerSideEncryptionByDefaultObservation.
func (in *ApplyServerSideEncryptionByDefaultObservation) DeepCopy() *ApplyServerSideEncryptionByDefaultObservation {
	if in == nil {
		return nil
	}
	out := new(ApplyServerSideEncryptionByDefaultObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyServerSideEncryptionByDefaultParameters) DeepCopyInto(out *ApplyServerSideEncryptionByDefaultParameters) {
	*out = *in
	if in.KMSMasterKeyID != nil {
		in, out := &in.KMSMasterKeyID, &out.KMSMasterKeyID
		*out = new(string)
		**out = **in
	}
	if in.SseAlgorithm != nil {
		in, out := &in.SseAlgorithm, &out.SseAlgorithm
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyServerSideEncryptionByDefaultParameters.
func (in *ApplyServerSideEncryptionByDefaultParameters) DeepCopy() *ApplyServerSideEncryptionByDefaultParameters {
	if in == nil {
		return nil
	}
	out := new(ApplyServerSideEncryptionByDefaultParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObservation) DeepCopyInto(out *BucketObservation) {
	*out = *in
	if in.BucketDomainName != nil {
		in, out := &in.BucketDomainName, &out.BucketDomainName
		*out = new(string)
		**out = **in
	}
	if in.CorsRule != nil {
		in, out := &in.CorsRule, &out.CorsRule
		*out = make([]CorsRuleObservation, len(*in))
		copy(*out, *in)
	}
	if in.Grant != nil {
		in, out := &in.Grant, &out.Grant
		*out = make([]GrantObservation, len(*in))
		copy(*out, *in)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LifecycleRule != nil {
		in, out := &in.LifecycleRule, &out.LifecycleRule
		*out = make([]LifecycleRuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = make([]LoggingObservation, len(*in))
		copy(*out, *in)
	}
	if in.ServerSideEncryptionConfiguration != nil {
		in, out := &in.ServerSideEncryptionConfiguration, &out.ServerSideEncryptionConfiguration
		*out = make([]ServerSideEncryptionConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = make([]VersioningObservation, len(*in))
		copy(*out, *in)
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = make([]WebsiteObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObservation.
func (in *BucketObservation) DeepCopy() *BucketObservation {
	if in == nil {
		return nil
	}
	out := new(BucketObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketParameters) DeepCopyInto(out *BucketParameters) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(string)
		**out = **in
	}
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.AccessKeyRef != nil {
		in, out := &in.AccessKeyRef, &out.AccessKeyRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.AccessKeySelector != nil {
		in, out := &in.AccessKeySelector, &out.AccessKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.BucketPrefix != nil {
		in, out := &in.BucketPrefix, &out.BucketPrefix
		*out = new(string)
		**out = **in
	}
	if in.CorsRule != nil {
		in, out := &in.CorsRule, &out.CorsRule
		*out = make([]CorsRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.Grant != nil {
		in, out := &in.Grant, &out.Grant
		*out = make([]GrantParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LifecycleRule != nil {
		in, out := &in.LifecycleRule, &out.LifecycleRule
		*out = make([]LifecycleRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = make([]LoggingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.SecretKeySecretRef != nil {
		in, out := &in.SecretKeySecretRef, &out.SecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServerSideEncryptionConfiguration != nil {
		in, out := &in.ServerSideEncryptionConfiguration, &out.ServerSideEncryptionConfiguration
		*out = make([]ServerSideEncryptionConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = make([]VersioningParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = make([]WebsiteParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebsiteDomain != nil {
		in, out := &in.WebsiteDomain, &out.WebsiteDomain
		*out = new(string)
		**out = **in
	}
	if in.WebsiteEndpoint != nil {
		in, out := &in.WebsiteEndpoint, &out.WebsiteEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketParameters.
func (in *BucketParameters) DeepCopy() *BucketParameters {
	if in == nil {
		return nil
	}
	out := new(BucketParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsRuleObservation) DeepCopyInto(out *CorsRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsRuleObservation.
func (in *CorsRuleObservation) DeepCopy() *CorsRuleObservation {
	if in == nil {
		return nil
	}
	out := new(CorsRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsRuleParameters) DeepCopyInto(out *CorsRuleParameters) {
	*out = *in
	if in.AllowedHeaders != nil {
		in, out := &in.AllowedHeaders, &out.AllowedHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedMethods != nil {
		in, out := &in.AllowedMethods, &out.AllowedMethods
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExposeHeaders != nil {
		in, out := &in.ExposeHeaders, &out.ExposeHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MaxAgeSeconds != nil {
		in, out := &in.MaxAgeSeconds, &out.MaxAgeSeconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsRuleParameters.
func (in *CorsRuleParameters) DeepCopy() *CorsRuleParameters {
	if in == nil {
		return nil
	}
	out := new(CorsRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpirationObservation) DeepCopyInto(out *ExpirationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpirationObservation.
func (in *ExpirationObservation) DeepCopy() *ExpirationObservation {
	if in == nil {
		return nil
	}
	out := new(ExpirationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpirationParameters) DeepCopyInto(out *ExpirationParameters) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.ExpiredObjectDeleteMarker != nil {
		in, out := &in.ExpiredObjectDeleteMarker, &out.ExpiredObjectDeleteMarker
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpirationParameters.
func (in *ExpirationParameters) DeepCopy() *ExpirationParameters {
	if in == nil {
		return nil
	}
	out := new(ExpirationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantObservation) DeepCopyInto(out *GrantObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantObservation.
func (in *GrantObservation) DeepCopy() *GrantObservation {
	if in == nil {
		return nil
	}
	out := new(GrantObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantParameters) DeepCopyInto(out *GrantParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantParameters.
func (in *GrantParameters) DeepCopy() *GrantParameters {
	if in == nil {
		return nil
	}
	out := new(GrantParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleRuleObservation) DeepCopyInto(out *LifecycleRuleObservation) {
	*out = *in
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = make([]ExpirationObservation, len(*in))
		copy(*out, *in)
	}
	if in.NoncurrentVersionExpiration != nil {
		in, out := &in.NoncurrentVersionExpiration, &out.NoncurrentVersionExpiration
		*out = make([]NoncurrentVersionExpirationObservation, len(*in))
		copy(*out, *in)
	}
	if in.NoncurrentVersionTransition != nil {
		in, out := &in.NoncurrentVersionTransition, &out.NoncurrentVersionTransition
		*out = make([]NoncurrentVersionTransitionObservation, len(*in))
		copy(*out, *in)
	}
	if in.Transition != nil {
		in, out := &in.Transition, &out.Transition
		*out = make([]TransitionObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleRuleObservation.
func (in *LifecycleRuleObservation) DeepCopy() *LifecycleRuleObservation {
	if in == nil {
		return nil
	}
	out := new(LifecycleRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleRuleParameters) DeepCopyInto(out *LifecycleRuleParameters) {
	*out = *in
	if in.AbortIncompleteMultipartUploadDays != nil {
		in, out := &in.AbortIncompleteMultipartUploadDays, &out.AbortIncompleteMultipartUploadDays
		*out = new(int64)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = make([]ExpirationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NoncurrentVersionExpiration != nil {
		in, out := &in.NoncurrentVersionExpiration, &out.NoncurrentVersionExpiration
		*out = make([]NoncurrentVersionExpirationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NoncurrentVersionTransition != nil {
		in, out := &in.NoncurrentVersionTransition, &out.NoncurrentVersionTransition
		*out = make([]NoncurrentVersionTransitionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Transition != nil {
		in, out := &in.Transition, &out.Transition
		*out = make([]TransitionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleRuleParameters.
func (in *LifecycleRuleParameters) DeepCopy() *LifecycleRuleParameters {
	if in == nil {
		return nil
	}
	out := new(LifecycleRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingObservation) DeepCopyInto(out *LoggingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingObservation.
func (in *LoggingObservation) DeepCopy() *LoggingObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingParameters) DeepCopyInto(out *LoggingParameters) {
	*out = *in
	if in.TargetBucket != nil {
		in, out := &in.TargetBucket, &out.TargetBucket
		*out = new(string)
		**out = **in
	}
	if in.TargetPrefix != nil {
		in, out := &in.TargetPrefix, &out.TargetPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingParameters.
func (in *LoggingParameters) DeepCopy() *LoggingParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionExpirationObservation) DeepCopyInto(out *NoncurrentVersionExpirationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionExpirationObservation.
func (in *NoncurrentVersionExpirationObservation) DeepCopy() *NoncurrentVersionExpirationObservation {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionExpirationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionExpirationParameters) DeepCopyInto(out *NoncurrentVersionExpirationParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionExpirationParameters.
func (in *NoncurrentVersionExpirationParameters) DeepCopy() *NoncurrentVersionExpirationParameters {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionExpirationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionTransitionObservation) DeepCopyInto(out *NoncurrentVersionTransitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionTransitionObservation.
func (in *NoncurrentVersionTransitionObservation) DeepCopy() *NoncurrentVersionTransitionObservation {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionTransitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionTransitionParameters) DeepCopyInto(out *NoncurrentVersionTransitionParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionTransitionParameters.
func (in *NoncurrentVersionTransitionParameters) DeepCopy() *NoncurrentVersionTransitionParameters {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionTransitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Object) DeepCopyInto(out *Object) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Object.
func (in *Object) DeepCopy() *Object {
	if in == nil {
		return nil
	}
	out := new(Object)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Object) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectList) DeepCopyInto(out *ObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Object, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectList.
func (in *ObjectList) DeepCopy() *ObjectList {
	if in == nil {
		return nil
	}
	out := new(ObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectObservation) DeepCopyInto(out *ObjectObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectObservation.
func (in *ObjectObservation) DeepCopy() *ObjectObservation {
	if in == nil {
		return nil
	}
	out := new(ObjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectParameters) DeepCopyInto(out *ObjectParameters) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(string)
		**out = **in
	}
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.AccessKeyRef != nil {
		in, out := &in.AccessKeyRef, &out.AccessKeyRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.AccessKeySelector != nil {
		in, out := &in.AccessKeySelector, &out.AccessKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.BucketRef != nil {
		in, out := &in.BucketRef, &out.BucketRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.BucketSelector != nil {
		in, out := &in.BucketSelector, &out.BucketSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.SecretKeySecretRef != nil {
		in, out := &in.SecretKeySecretRef, &out.SecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectParameters.
func (in *ObjectParameters) DeepCopy() *ObjectParameters {
	if in == nil {
		return nil
	}
	out := new(ObjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSpec) DeepCopyInto(out *ObjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSpec.
func (in *ObjectSpec) DeepCopy() *ObjectSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStatus) DeepCopyInto(out *ObjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStatus.
func (in *ObjectStatus) DeepCopy() *ObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleObservation) DeepCopyInto(out *RuleObservation) {
	*out = *in
	if in.ApplyServerSideEncryptionByDefault != nil {
		in, out := &in.ApplyServerSideEncryptionByDefault, &out.ApplyServerSideEncryptionByDefault
		*out = make([]ApplyServerSideEncryptionByDefaultObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleObservation.
func (in *RuleObservation) DeepCopy() *RuleObservation {
	if in == nil {
		return nil
	}
	out := new(RuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleParameters) DeepCopyInto(out *RuleParameters) {
	*out = *in
	if in.ApplyServerSideEncryptionByDefault != nil {
		in, out := &in.ApplyServerSideEncryptionByDefault, &out.ApplyServerSideEncryptionByDefault
		*out = make([]ApplyServerSideEncryptionByDefaultParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleParameters.
func (in *RuleParameters) DeepCopy() *RuleParameters {
	if in == nil {
		return nil
	}
	out := new(RuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSideEncryptionConfigurationObservation) DeepCopyInto(out *ServerSideEncryptionConfigurationObservation) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSideEncryptionConfigurationObservation.
func (in *ServerSideEncryptionConfigurationObservation) DeepCopy() *ServerSideEncryptionConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ServerSideEncryptionConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSideEncryptionConfigurationParameters) DeepCopyInto(out *ServerSideEncryptionConfigurationParameters) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSideEncryptionConfigurationParameters.
func (in *ServerSideEncryptionConfigurationParameters) DeepCopy() *ServerSideEncryptionConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ServerSideEncryptionConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitionObservation) DeepCopyInto(out *TransitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitionObservation.
func (in *TransitionObservation) DeepCopy() *TransitionObservation {
	if in == nil {
		return nil
	}
	out := new(TransitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitionParameters) DeepCopyInto(out *TransitionParameters) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitionParameters.
func (in *TransitionParameters) DeepCopy() *TransitionParameters {
	if in == nil {
		return nil
	}
	out := new(TransitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersioningObservation) DeepCopyInto(out *VersioningObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersioningObservation.
func (in *VersioningObservation) DeepCopy() *VersioningObservation {
	if in == nil {
		return nil
	}
	out := new(VersioningObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersioningParameters) DeepCopyInto(out *VersioningParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersioningParameters.
func (in *VersioningParameters) DeepCopy() *VersioningParameters {
	if in == nil {
		return nil
	}
	out := new(VersioningParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteObservation) DeepCopyInto(out *WebsiteObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteObservation.
func (in *WebsiteObservation) DeepCopy() *WebsiteObservation {
	if in == nil {
		return nil
	}
	out := new(WebsiteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteParameters) DeepCopyInto(out *WebsiteParameters) {
	*out = *in
	if in.ErrorDocument != nil {
		in, out := &in.ErrorDocument, &out.ErrorDocument
		*out = new(string)
		**out = **in
	}
	if in.IndexDocument != nil {
		in, out := &in.IndexDocument, &out.IndexDocument
		*out = new(string)
		**out = **in
	}
	if in.RedirectAllRequestsTo != nil {
		in, out := &in.RedirectAllRequestsTo, &out.RedirectAllRequestsTo
		*out = new(string)
		**out = **in
	}
	if in.RoutingRules != nil {
		in, out := &in.RoutingRules, &out.RoutingRules
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteParameters.
func (in *WebsiteParameters) DeepCopy() *WebsiteParameters {
	if in == nil {
		return nil
	}
	out := new(WebsiteParameters)
	in.DeepCopyInto(out)
	return out
}
