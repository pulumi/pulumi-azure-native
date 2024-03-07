package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managedidentity/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managedidentity.NewFederatedIdentityCredential(ctx, "federatedIdentityCredential", &managedidentity.FederatedIdentityCredentialArgs{
			Audiences: pulumi.StringArray{
				pulumi.String("api://AzureADTokenExchange"),
			},
			FederatedIdentityCredentialResourceName: pulumi.String("ficResourceName"),
			Issuer:                                  pulumi.String("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID"),
			ResourceGroupName:                       pulumi.String("rgName"),
			ResourceName:                            pulumi.String("resourceName"),
			Subject:                                 pulumi.String("system:serviceaccount:ns:svcaccount"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
