package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/m365securityandcompliance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := m365securityandcompliance.NewPrivateLinkServicesForM365SecurityCenter(ctx, "privateLinkServicesForM365SecurityCenter", &m365securityandcompliance.PrivateLinkServicesForM365SecurityCenterArgs{
			Kind:     m365securityandcompliance.Kind_Fhir_R4,
			Location: pulumi.String("westus2"),
			Properties: &m365securityandcompliance.ServicesPropertiesArgs{
				AccessPolicies: m365securityandcompliance.ServiceAccessPolicyEntryArray{
					&m365securityandcompliance.ServiceAccessPolicyEntryArgs{
						ObjectId: pulumi.String("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("service2"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
