package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewProject(ctx, "project", &devcenter.ProjectArgs{
			Description:       pulumi.String("This is my first project."),
			DevCenterId:       pulumi.String("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso"),
			Location:          pulumi.String("centralus"),
			ProjectName:       pulumi.String("DevProject"),
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"CostCenter": pulumi.String("R&D"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
