package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resourcegraph/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resourcegraph.NewGraphQuery(ctx, "graphQuery", &resourcegraph.GraphQueryArgs{
			Description:       pulumi.String("Docker VMs in PROD"),
			Query:             pulumi.String("where isnotnull(tags['Prod']) and properties.extensions[0].Name == 'docker'"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ResourceName:      pulumi.String("MyDockerVMs"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
