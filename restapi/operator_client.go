// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package restapi

import (
	"context"

	v1 "github.com/minio/operator/pkg/apis/minio.min.io/v1"
	operatorClientset "github.com/minio/operator/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
)

// OperatorClientI interface with all functions to be implemented
// by mock when testing, it should include all OperatorClientI respective api calls
// that are used within this project.
type OperatorClientI interface {
	TenantDelete(ctx context.Context, namespace string, instanceName string, options metav1.DeleteOptions) error
	TenantGet(ctx context.Context, namespace string, instanceName string, options metav1.GetOptions) (*v1.Tenant, error)
	TenantPatch(ctx context.Context, namespace string, instanceName string, pt types.PatchType, data []byte, options metav1.PatchOptions) (*v1.Tenant, error)
	TenantList(ctx context.Context, namespace string, opts metav1.ListOptions) (*v1.TenantList, error)
}

// Interface implementation
//
// Define the structure of a operator client and define the functions that are actually used
// from the minio operator.
type operatorClient struct {
	client *operatorClientset.Clientset
}

// TenantDelete implements the minio instance delete action from minio operator
func (c *operatorClient) TenantDelete(ctx context.Context, namespace string, instanceName string, options metav1.DeleteOptions) error {
	return c.client.MinioV1().Tenants(namespace).Delete(ctx, instanceName, options)
}

// TenantGet implements the minio instance get action from minio operator
func (c *operatorClient) TenantGet(ctx context.Context, namespace string, instanceName string, options metav1.GetOptions) (*v1.Tenant, error) {
	return c.client.MinioV1().Tenants(namespace).Get(ctx, instanceName, options)
}

// TenantPatch implements the minio instance patch action from minio operator
func (c *operatorClient) TenantPatch(ctx context.Context, namespace string, instanceName string, pt types.PatchType, data []byte, options metav1.PatchOptions) (*v1.Tenant, error) {
	return c.client.MinioV1().Tenants(namespace).Patch(ctx, instanceName, pt, data, options)
}

// TenantList implements the minio instance list action from minio operator
func (c *operatorClient) TenantList(ctx context.Context, namespace string, opts metav1.ListOptions) (*v1.TenantList, error) {
	return c.client.MinioV1().Tenants(namespace).List(ctx, opts)
}
