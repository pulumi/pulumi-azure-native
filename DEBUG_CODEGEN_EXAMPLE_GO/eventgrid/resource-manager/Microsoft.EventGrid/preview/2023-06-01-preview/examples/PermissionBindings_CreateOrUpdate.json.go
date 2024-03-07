package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewPermissionBinding(ctx, "permissionBinding", &eventgrid.PermissionBindingArgs{
			ClientGroupName:       pulumi.String("exampleClientGroupName1"),
			NamespaceName:         pulumi.String("exampleNamespaceName1"),
			Permission:            pulumi.String("Publisher"),
			PermissionBindingName: pulumi.String("examplePermissionBindingName1"),
			ResourceGroupName:     pulumi.String("examplerg"),
			TopicSpaceName:        pulumi.String("exampleTopicSpaceName1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
