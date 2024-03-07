package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewSitesController(ctx, "sitesController", &offazure.SitesControllerArgs{
			AgentDetails: &offazure.SiteAgentPropertiesArgs{
				KeyVaultId:  pulumi.String("awxurtbjmofxuciewsqfgpkccpzw"),
				KeyVaultUri: pulumi.String("qizphgqwage"),
			},
			ApplianceName:       pulumi.String("zkzibwptff"),
			DiscoverySolutionId: pulumi.String("lvskpghpphpfumbzxroakznqplomiy"),
			Location:            pulumi.String("mh"),
			ResourceGroupName:   pulumi.String("rgmigrate"),
			ServicePrincipalIdentityDetails: &offazure.SiteSpnPropertiesArgs{
				AadAuthority:  pulumi.String("yanzipdww"),
				ApplicationId: pulumi.String("tspgrujepxyxuprkqvfuqbbjrweeqa"),
				Audience:      pulumi.String("oepwfaozztzvegmzvswafvotj"),
				ObjectId:      pulumi.String("tqrjngpgxnnto"),
				RawCertData:   pulumi.String("dotvgkslkmsgvtekgojnhcdrryk"),
				TenantId:      pulumi.String("vesmyhu"),
			},
			SiteName: pulumi.String("3-MmUUL6025"),
			Tags: pulumi.StringMap{
				"key7046": pulumi.String("qywdefracwwzlugbwgqbt"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
