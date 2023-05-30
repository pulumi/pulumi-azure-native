// Copyright 2022, Pulumi Corporation.  All rights reserved.

package main

import (
	"fmt"

	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		config, err := authorization.GetClientConfig(ctx)
		if err != nil {
			return err
		}

		// The managed identity (MSI) authentication method will not have an object id, but it's
		// unlikely our tests run on Azure VMs with this configuration.
		if config.ClientId == "" || config.ObjectId == "" || config.SubscriptionId == "" || config.TenantId == "" {
			return fmt.Errorf("Expected client, object, subscription, and tenant ids in auth config %v", config)
		}

		return nil
	})
}
