// Copyright 2025, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const customDomainWithSslTok = "azure-native:cdn:CustomDomain"

func customDomainWithSsl(crudClientFactory crud.ResourceCrudClientFactory, lookupResource resources.ResourceLookupFunc) (*CustomResource, error) {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}",
		tok:  customDomainWithSslTok,
		Schema: func(resource *ResourceDefinition) (*ResourceDefinition, error) {
			if resource == nil {
				return nil, fmt.Errorf("resource %q not found in the spec", customDomainWithSslTok)
			}
			return &ResourceDefinition{}, nil
		},
		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
			return nil, nil
		},
		Update: func(ctx context.Context, id string, inputs resource.PropertyMap, state resource.PropertyMap) (map[string]any, error) {
			return nil, nil
		},
	}, nil
}
