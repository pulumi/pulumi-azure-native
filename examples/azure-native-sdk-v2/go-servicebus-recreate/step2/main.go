// Copyright 2021, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v3"
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}

		namespace, err := servicebus.NewNamespace(ctx, "ns", &servicebus.NamespaceArgs{
			ResourceGroupName: resourceGroup.Name,
		})
		if err != nil {
			return err
		}

		topic, err := servicebus.NewTopic(ctx, "topic", &servicebus.TopicArgs{
			ResourceGroupName:          resourceGroup.Name,
			NamespaceName:              namespace.Name,
			RequiresDuplicateDetection: pulumi.Bool(true), // Changing this should trigger recreation otherwise will be rejected.
		})
		if err != nil {
			return err
		}

		// Export the topic name
		ctx.Export("topicName", topic.Name)

		return nil
	})
}
