package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/connectedvmwarevsphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connectedvmwarevsphere.NewGuestAgent(ctx, "guestAgent", &connectedvmwarevsphere.GuestAgentArgs{
			Credentials: &connectedvmwarevsphere.GuestCredentialArgs{
				Password: pulumi.String("<password>"),
				Username: pulumi.String("tempuser"),
			},
			HttpProxyConfig: &connectedvmwarevsphere.HttpProxyConfigurationArgs{
				HttpsProxy: pulumi.String("http://192.1.2.3:8080"),
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
