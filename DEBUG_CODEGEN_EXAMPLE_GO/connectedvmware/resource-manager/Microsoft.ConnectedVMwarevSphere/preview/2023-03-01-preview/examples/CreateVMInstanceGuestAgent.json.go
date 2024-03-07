package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/connectedvmwarevsphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connectedvmwarevsphere.NewVMInstanceGuestAgent(ctx, "vmInstanceGuestAgent", &connectedvmwarevsphere.VMInstanceGuestAgentArgs{
			Credentials: &connectedvmwarevsphere.GuestCredentialArgs{
				Password: pulumi.String("<password>"),
				Username: pulumi.String("tempuser"),
			},
			HttpProxyConfig: &connectedvmwarevsphere.HttpProxyConfigurationArgs{
				HttpsProxy: pulumi.String("http://192.1.2.3:8080"),
			},
			PrivateLinkScopeResourceId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
			ProvisioningAction:         pulumi.String("install"),
			ResourceUri:                pulumi.String("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
