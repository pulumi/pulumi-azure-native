package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityandcompliance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityandcompliance.NewPrivateLinkServicesForEDMUpload(ctx, "privateLinkServicesForEDMUpload", &securityandcompliance.PrivateLinkServicesForEDMUploadArgs{
			Kind:     securityandcompliance.Kind_Fhir_R4,
			Location: pulumi.String("westus2"),
			Properties: &securityandcompliance.ServicesPropertiesArgs{
				AccessPolicies: securityandcompliance.ServiceAccessPolicyEntryArray{
					&securityandcompliance.ServiceAccessPolicyEntryArgs{
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
