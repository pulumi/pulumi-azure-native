package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/scvmm/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scvmm.NewVmmServer(ctx, "vmmServer", &scvmm.VmmServerArgs{
			Credentials: &scvmm.VMMServerPropertiesCredentialsArgs{
				Password: pulumi.String("password"),
				Username: pulumi.String("testuser"),
			},
			ExtendedLocation: &scvmm.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
				Type: pulumi.String("customLocation"),
			},
			Fqdn:              pulumi.String("VMM.contoso.com"),
			Location:          pulumi.String("East US"),
			Port:              pulumi.Int(1234),
			ResourceGroupName: pulumi.String("testrg"),
			VmmServerName:     pulumi.String("ContosoVMMServer"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
