// Copyright 2025, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const (
	tagAtScopeTok  = "azure-native:resources:TagAtScope"
	tagAtScopePath = "/{scope}/providers/Microsoft.Resources/tags/default"
)

// tagAtScope overrides Create, Update, and Delete for TagAtScope to use the PATCH endpoint
// (Tags_UpdateAtScope) instead of the default PUT (Tags_CreateOrUpdateAtScope).
//
// This uses operation=Merge for Create/Update so that tags managed by Pulumi are added or
// updated without disturbing tags on the scope that Pulumi does not own.  Delete uses
// operation=Delete to remove only the tags that were declared in the resource, leaving
// any other tags on the scope untouched.
func tagAtScope(
	lookupResource resources.ResourceLookupFunc,
	crudClientFactory crud.ResourceCrudClientFactory,
	azureClient azure.AzureClient,
) (*CustomResource, error) {
	var client crud.ResourceCrudClient
	var asyncStyle string
	if lookupResource != nil && crudClientFactory != nil {
		res, found, err := lookupResource(tagAtScopeTok)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("resource %q not found", tagAtScopeTok)
		}
		client = crudClientFactory(&res)
		// PUT async-style is the same for PATCH
		asyncStyle = res.PutAsyncStyle
	}

	// patch calls the PATCH endpoint with the given operation ("Merge" or "Delete").
	// inputs must be in Pulumi SDK shape; PrepareAzureRESTBody converts them to ARM format.
	patch := func(ctx context.Context, id, operation string, inputs resource.PropertyMap) (map[string]any, error) {
		body, err := client.PrepareAzureRESTBody(id, inputs, nil)
		if err != nil {
			return nil, err
		}
		body["operation"] = operation
		queryParams := map[string]any{"api-version": client.ApiVersion()}
		resp, _, err := azureClient.Patch(ctx, id, body, queryParams, asyncStyle)
		if err != nil {
			return nil, err
		}
		return client.ResponseBodyToSdkOutputs(resp), nil
	}

	return &CustomResource{
		path: tagAtScopePath,
		tok:  tagAtScopeTok,

		CanCreate: func(ctx context.Context, id string) error {
			// we should always be able to `Create` since we are using
			// Merge-based updates on a default resource
			return nil
		},

		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
			return patch(ctx, id, "Merge", inputs)
		},

		Update: func(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]any, error) {
			return patch(ctx, id, "Merge", news)
		},

		// Delete removes only the tags declared in this resource, leaving other tags intact.
		Delete: func(ctx context.Context, id string, previousInputs, state resource.PropertyMap) error {
			_, err := patch(ctx, id, "Delete", previousInputs)
			return err
		},
	}, nil
}
