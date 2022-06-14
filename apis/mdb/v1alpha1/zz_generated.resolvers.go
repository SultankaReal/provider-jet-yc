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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1"
	v1alpha11 "github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this MongodbCluster.
func (mg *MongodbCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this MySQLCluster.
func (mg *MySQLCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this MySQLDatabase.
func (mg *MySQLDatabase) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &MySQLClusterList{},
			Managed: &MySQLCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MySQLUser.
func (mg *MySQLUser) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &MySQLClusterList{},
			Managed: &MySQLCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PostgresqlCluster.
func (mg *PostgresqlCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &PostgresqlClusterList{},
			Managed: &PostgresqlCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PostgresqlUser.
func (mg *PostgresqlUser) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &PostgresqlClusterList{},
			Managed: &PostgresqlCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RedisCluster.
func (mg *RedisCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}
