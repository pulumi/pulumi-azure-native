package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/confidentialledger/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := confidentialledger.NewLedger(ctx, "ledger", &confidentialledger.LedgerArgs{
			LedgerName: pulumi.String("DummyLedgerName"),
			Location:   pulumi.String("EastUS"),
			Properties: confidentialledger.LedgerPropertiesResponse{
				AadBasedSecurityPrincipals: confidentialledger.AADBasedSecurityPrincipalArray{
					&confidentialledger.AADBasedSecurityPrincipalArgs{
						LedgerRoleName: pulumi.String("Administrator"),
						PrincipalId:    pulumi.String("34621747-6fc8-4771-a2eb-72f31c461f2e"),
						TenantId:       pulumi.String("bce123b9-2b7b-4975-8360-5ca0b9b1cd08"),
					},
				},
				CertBasedSecurityPrincipals: confidentialledger.CertBasedSecurityPrincipalArray{
					&confidentialledger.CertBasedSecurityPrincipalArgs{
						Cert:           pulumi.String("-----BEGIN CERTIFICATE-----MIIBsjCCATigAwIBAgIUZWIbyG79TniQLd2UxJuU74tqrKcwCgYIKoZIzj0EAwMwEDEOMAwGA1UEAwwFdXNlcjAwHhcNMjEwMzE2MTgwNjExWhcNMjIwMzE2MTgwNjExWjAQMQ4wDAYDVQQDDAV1c2VyMDB2MBAGByqGSM49AgEGBSuBBAAiA2IABBiWSo/j8EFit7aUMm5lF+lUmCu+IgfnpFD+7QMgLKtxRJ3aGSqgS/GpqcYVGddnODtSarNE/HyGKUFUolLPQ5ybHcouUk0kyfA7XMeSoUA4lBz63Wha8wmXo+NdBRo39qNTMFEwHQYDVR0OBBYEFPtuhrwgGjDFHeUUT4nGsXaZn69KMB8GA1UdIwQYMBaAFPtuhrwgGjDFHeUUT4nGsXaZn69KMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwMDaAAwZQIxAOnozm2CyqRwSSQLls5r+mUHRGRyXHXwYtM4Dcst/VEZdmS9fqvHRCHbjUlO/+HNfgIwMWZ4FmsjD3wnPxONOm9YdVn/PRD7SsPRPbOjwBiE4EBGaHDsLjYAGDSGi7NJnSkA-----END CERTIFICATE-----"),
						LedgerRoleName: pulumi.String("Reader"),
					},
				},
				LedgerType: pulumi.String("Public"),
			},
			ResourceGroupName: pulumi.String("DummyResourceGroupName"),
			Tags: pulumi.StringMap{
				"additionalProps1": pulumi.String("additional properties"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
