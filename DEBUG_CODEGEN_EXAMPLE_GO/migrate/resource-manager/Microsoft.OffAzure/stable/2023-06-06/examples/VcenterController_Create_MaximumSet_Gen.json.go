package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewVcenterController(ctx, "vcenterController", &offazure.VcenterControllerArgs{
			Fqdn:              pulumi.String("mzrejobgzhpxhynsfpumuhk"),
			FriendlyName:      pulumi.String("qqsnhutpvockzhwuesuohnfirvxmw"),
			Port:              pulumi.String("ahagklwzutlumcdhawkrfzmpfypcz"),
			ProvisioningState: pulumi.String("Created"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			RunAsAccountId:    pulumi.String("orrfsvlcuoagqlkkqsd"),
			SiteName:          pulumi.String("Snu-38v6"),
			VcenterName:       pulumi.String("e72lZRSD"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
