package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewWorkspacePolicyFragment(ctx, "workspacePolicyFragment", &apimanagement.WorkspacePolicyFragmentArgs{
			Description:       pulumi.String("A policy fragment example"),
			Format:            pulumi.String("xml"),
			Id:                pulumi.String("policyFragment1"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Value:             pulumi.String("<fragment><json-to-xml apply=\"always\" consider-accept-header=\"false\" /></fragment>"),
			WorkspaceId:       pulumi.String("wks1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
