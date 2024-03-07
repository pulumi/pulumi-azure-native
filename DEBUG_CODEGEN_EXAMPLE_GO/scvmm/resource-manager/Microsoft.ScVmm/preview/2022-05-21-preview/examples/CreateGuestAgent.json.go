package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/scvmm/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scvmm.NewGuestAgent(ctx, "guestAgent", &scvmm.GuestAgentArgs{
			Credentials: &scvmm.GuestCredentialArgs{
				Password: pulumi.String("<password>"),
				Username: pulumi.String("tempuser"),
			},
			GuestAgentName: pulumi.String("default"),
			HttpProxyConfig: &scvmm.HttpProxyConfigurationArgs{
				HttpsProxy: pulumi.String("http://192.1.2.3:8080"),
			},
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
