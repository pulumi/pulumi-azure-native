package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewTrustedAccessRoleBinding(ctx, "trustedAccessRoleBinding", &containerservice.TrustedAccessRoleBindingArgs{
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("clustername1"),
			Roles: pulumi.StringArray{
				pulumi.String("Microsoft.MachineLearningServices/workspaces/reader"),
				pulumi.String("Microsoft.MachineLearningServices/workspaces/writer"),
			},
			SourceResourceId:             pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c"),
			TrustedAccessRoleBindingName: pulumi.String("binding1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
