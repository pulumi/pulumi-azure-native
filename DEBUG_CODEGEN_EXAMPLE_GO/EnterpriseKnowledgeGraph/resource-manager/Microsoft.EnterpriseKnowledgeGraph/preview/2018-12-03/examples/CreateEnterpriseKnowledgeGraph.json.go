package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/enterpriseknowledgegraph/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := enterpriseknowledgegraph.NewEnterpriseKnowledgeGraph(ctx, "enterpriseKnowledgeGraph", &enterpriseknowledgegraph.EnterpriseKnowledgeGraphArgs{
			Location:          pulumi.String("West US"),
			Properties:        nil,
			ResourceGroupName: pulumi.String("OneResourceGroupName"),
			ResourceName:      pulumi.String("sampleekgname"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
