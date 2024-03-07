package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewSqlSitesController(ctx, "sqlSitesController", &offazure.SqlSitesControllerArgs{
			DiscoveryScenario: pulumi.String("Migrate"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteAppliancePropertiesCollection: []offazure.SiteAppliancePropertiesArgs{
				{
					AgentDetails: {
						KeyVaultId:  pulumi.String("awxurtbjmofxuciewsqfgpkccpzw"),
						KeyVaultUri: pulumi.String("qizphgqwage"),
					},
					ApplianceName: pulumi.String("zpbkpigahvexsxevwafzgsu"),
					ServicePrincipalIdentityDetails: {
						AadAuthority:  pulumi.String("yanzipdww"),
						ApplicationId: pulumi.String("tspgrujepxyxuprkqvfuqbbjrweeqa"),
						Audience:      pulumi.String("oepwfaozztzvegmzvswafvotj"),
						ObjectId:      pulumi.String("tqrjngpgxnnto"),
						RawCertData:   pulumi.String("dotvgkslkmsgvtekgojnhcdrryk"),
						TenantId:      pulumi.String("vesmyhu"),
					},
				},
			},
			SiteName:    pulumi.String("qo36fi7J8y-b"),
			SqlSiteName: pulumi.String("A-4l5sbX5"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
