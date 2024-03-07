package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewGuestAgent(ctx, "guestAgent", &azurestackhci.GuestAgentArgs{
			Credentials: &azurestackhci.GuestCredentialArgs{
				Password: pulumi.String("<password>"),
				Username: pulumi.String("tempuser"),
			},
			Name:               pulumi.String("default"),
			ProvisioningAction: pulumi.String("install"),
			ResourceGroupName:  pulumi.String("testrg"),
			VirtualMachineName: pulumi.String("ContosoVm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
