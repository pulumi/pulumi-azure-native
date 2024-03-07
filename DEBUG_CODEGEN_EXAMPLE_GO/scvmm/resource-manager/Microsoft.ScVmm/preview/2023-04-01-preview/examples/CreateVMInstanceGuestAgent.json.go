package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/scvmm/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scvmm.NewVMInstanceGuestAgent(ctx, "vmInstanceGuestAgent", &scvmm.VMInstanceGuestAgentArgs{
			Credentials: &scvmm.GuestCredentialArgs{
				Password: pulumi.String("<password>"),
				Username: pulumi.String("tempuser"),
			},
			HttpProxyConfig: &scvmm.HttpProxyConfigurationArgs{
				HttpsProxy: pulumi.String("http://192.1.2.3:8080"),
			},
			ProvisioningAction: pulumi.String("install"),
			ResourceUri:        pulumi.String("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
