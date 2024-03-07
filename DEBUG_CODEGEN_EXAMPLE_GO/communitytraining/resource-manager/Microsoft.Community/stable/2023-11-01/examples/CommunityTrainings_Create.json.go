package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/community/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := community.NewCommunityTraining(ctx, "communityTraining", &community.CommunityTrainingArgs{
			CommunityTrainingName:   pulumi.String("ctApplication"),
			DisasterRecoveryEnabled: pulumi.Bool(true),
			IdentityConfiguration: &community.IdentityConfigurationPropertiesArgs{
				B2cAuthenticationPolicy: pulumi.String("B2C_1_signup_signin"),
				B2cPasswordResetPolicy:  pulumi.String("B2C_1_pwd_reset"),
				ClientId:                pulumi.String("8c92390f-2f30-493d-bd13-d3c3eba3709d"),
				ClientSecret:            pulumi.String("idenityConfigurationClientSecret"),
				CustomLoginParameters:   pulumi.String("custom_hint"),
				DomainName:              pulumi.String("cttenant"),
				IdentityType:            pulumi.String("ADB2C"),
				TeamsEnabled:            pulumi.Bool(false),
				TenantId:                pulumi.String("c1ffbb60-88cf-4b83-b54f-c47ae6220c19"),
			},
			Location:                    pulumi.String("southeastasia"),
			PortalAdminEmailAddress:     pulumi.String("ctadmin@ct.com"),
			PortalName:                  pulumi.String("ctwebsite"),
			PortalOwnerEmailAddress:     pulumi.String("ctcontact@ct.com"),
			PortalOwnerOrganizationName: pulumi.String("CT Portal Owner Organization"),
			ResourceGroupName:           pulumi.String("rgCommunityTaining"),
			Sku: &community.SkuArgs{
				Name: pulumi.String("Commercial"),
				Tier: community.SkuTierStandard,
			},
			ZoneRedundancyEnabled: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
