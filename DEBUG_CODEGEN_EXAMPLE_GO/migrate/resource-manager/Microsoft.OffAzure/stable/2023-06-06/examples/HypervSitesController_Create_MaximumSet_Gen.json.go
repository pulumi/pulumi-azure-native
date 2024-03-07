package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewHypervSitesController(ctx, "hypervSitesController", &offazure.HypervSitesControllerArgs{
			AgentDetails: &offazure.SiteAgentPropertiesArgs{
				KeyVaultId:  pulumi.String("awxurtbjmofxuciewsqfgpkccpzw"),
				KeyVaultUri: pulumi.String("qizphgqwage"),
			},
			ApplianceName:       pulumi.String("jnpsjguxuzuxbhasiqfijf"),
			DiscoverySolutionId: pulumi.String("yxajidsykozchjkuxj"),
			Location:            pulumi.String("sctymxdndonxgejdhi"),
			ProvisioningState:   pulumi.String("Created"),
			ResourceGroupName:   pulumi.String("rgmigrate"),
			ServicePrincipalIdentityDetails: &offazure.SiteSpnPropertiesArgs{
				AadAuthority:  pulumi.String("yanzipdww"),
				ApplicationId: pulumi.String("tspgrujepxyxuprkqvfuqbbjrweeqa"),
				Audience:      pulumi.String("oepwfaozztzvegmzvswafvotj"),
				ObjectId:      pulumi.String("tqrjngpgxnnto"),
				RawCertData:   pulumi.String("dotvgkslkmsgvtekgojnhcdrryk"),
				TenantId:      pulumi.String("vesmyhu"),
			},
			SiteName: pulumi.String("692-u-93F93"),
			Tags: pulumi.StringMap{
				"key9741": pulumi.String("fdbzykkpvhnhsrhiam"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
