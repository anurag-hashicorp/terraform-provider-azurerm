// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package tableserviceproperties

import (
	"fmt"

	storageClient "github.com/hashicorp/go-azure-sdk/sdk/client/dataplane/storage"
)

// Client is the base client for Table Storage service properties.
type Client struct {
	Client *storageClient.Client
}

func NewWithBaseUri(baseUri string) (*Client, error) {
	baseClient, err := storageClient.NewStorageClient(baseUri, componentName, apiVersion)
	if err != nil {
		return nil, fmt.Errorf("building base client: %+v", err)
	}

	return &Client{
		Client: baseClient,
	}, nil
}
