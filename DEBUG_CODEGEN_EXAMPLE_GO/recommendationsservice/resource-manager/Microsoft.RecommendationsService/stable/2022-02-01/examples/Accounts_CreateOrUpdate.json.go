package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recommendationsservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recommendationsservice.NewAccount(ctx, "account", &recommendationsservice.AccountArgs{
			AccountName: pulumi.String("sampleAccount"),
			Location:    pulumi.String("West US"),
			Properties: recommendationsservice.AccountResourceResponseProperties{
				Configuration: pulumi.String("Capacity"),
				EndpointAuthentications: recommendationsservice.EndpointAuthenticationArray{
					&recommendationsservice.EndpointAuthenticationArgs{
						AadTenantID:   pulumi.String("tenant"),
						PrincipalID:   pulumi.String("oid"),
						PrincipalType: pulumi.String("User"),
					},
				},
			},
			ResourceGroupName: pulumi.String("rg"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Prod"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
