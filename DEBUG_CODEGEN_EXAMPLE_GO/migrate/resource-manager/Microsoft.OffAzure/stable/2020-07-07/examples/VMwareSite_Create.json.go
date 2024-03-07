package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewSite(ctx, "site", &offazure.SiteArgs{
			Location: pulumi.String("eastus"),
			Properties: &offazure.SitePropertiesArgs{
				ServicePrincipalIdentityDetails: &offazure.SiteSpnPropertiesArgs{
					AadAuthority:  pulumi.String("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
					ApplicationId: pulumi.String("e9f013df-2a2a-4871-b766-e79867f30348"),
					Audience:      pulumi.String("https://72f988bf-86f1-41af-91ab-2d7cd011db47/MaheshSite17ac9agentauthaadapp"),
					ObjectId:      pulumi.String("2cd492bc-7ef3-4ee0-b301-59a88108b47b"),
					TenantId:      pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				},
			},
			ResourceGroupName: pulumi.String("pajindTest"),
			SiteName:          pulumi.String("appliance1e39site"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
