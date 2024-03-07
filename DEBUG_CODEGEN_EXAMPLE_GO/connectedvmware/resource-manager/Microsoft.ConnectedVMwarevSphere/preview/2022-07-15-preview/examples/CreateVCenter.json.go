package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/connectedvmwarevsphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connectedvmwarevsphere.NewVCenter(ctx, "vCenter", &connectedvmwarevsphere.VCenterArgs{
			Credentials: &connectedvmwarevsphere.VICredentialArgs{
				Password: pulumi.String("<password>"),
				Username: pulumi.String("tempuser"),
			},
			ExtendedLocation: &connectedvmwarevsphere.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
				Type: pulumi.String("customLocation"),
			},
			Fqdn:              pulumi.String("ContosoVMware.contoso.com"),
			Location:          pulumi.String("East US"),
			Port:              pulumi.Int(1234),
			ResourceGroupName: pulumi.String("testrg"),
			VcenterName:       pulumi.String("ContosoVCenter"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
