package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewImportSitesController(ctx, "importSitesController", &offazure.ImportSitesControllerArgs{
			DiscoverySolutionId: pulumi.String("lfaswhiwdttdpkrvnrpriauexdjs"),
			Location:            pulumi.String("woctgvdufvkzvjcirjpf"),
			ProvisioningState:   pulumi.String("Created"),
			ResourceGroupName:   pulumi.String("rgmigrate"),
			SiteName:            pulumi.String("O-K1FS"),
			Tags: pulumi.StringMap{
				"key2067": pulumi.String("taqzca"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
