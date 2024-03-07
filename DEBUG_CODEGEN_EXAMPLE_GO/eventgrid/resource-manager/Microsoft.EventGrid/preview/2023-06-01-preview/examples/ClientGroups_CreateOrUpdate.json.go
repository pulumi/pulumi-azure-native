package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewClientGroup(ctx, "clientGroup", &eventgrid.ClientGroupArgs{
			ClientGroupName:   pulumi.String("exampleClientGroupName1"),
			Description:       pulumi.String("This is a test client group"),
			NamespaceName:     pulumi.String("exampleNamespaceName1"),
			Query:             pulumi.String("attributes.b IN ['a', 'b', 'c']"),
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
