// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package shim

import (
	"context"

	"github.com/hashicorp/terraform-provider-azurerm/internal/services/storage/tableserviceproperties"
	"github.com/jackofallops/giovanni/storage/2023-11-03/table/tables"
)

type StorageTableWrapper interface {
	Create(ctx context.Context, tableName string) error
	Delete(ctx context.Context, tableName string) error
	Exists(ctx context.Context, tableName string) (*bool, error)
	GetServiceProperties(ctx context.Context) (*tableserviceproperties.StorageServiceProperties, error)
	GetACLs(ctx context.Context, tableName string) (*[]tables.SignedIdentifier, error)
	UpdateServiceProperties(ctx context.Context, properties tableserviceproperties.StorageServiceProperties) error
	UpdateACLs(ctx context.Context, tableName string, acls []tables.SignedIdentifier) error
}
