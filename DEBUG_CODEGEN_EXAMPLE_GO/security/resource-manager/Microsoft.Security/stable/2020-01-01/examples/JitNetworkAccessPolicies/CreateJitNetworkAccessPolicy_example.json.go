package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewJitNetworkAccessPolicy(ctx, "jitNetworkAccessPolicy", &security.JitNetworkAccessPolicyArgs{
			AscLocation:                pulumi.String("westeurope"),
			JitNetworkAccessPolicyName: pulumi.String("default"),
			Kind:                       pulumi.String("Basic"),
			Requests: security.JitNetworkAccessRequestArray{
				&security.JitNetworkAccessRequestArgs{
					Requestor:    pulumi.String("barbara@contoso.com"),
					StartTimeUtc: pulumi.String("2018-05-17T08:06:45.5691611Z"),
					VirtualMachines: security.JitNetworkAccessRequestVirtualMachineArray{
						&security.JitNetworkAccessRequestVirtualMachineArgs{
							Id: pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
							Ports: security.JitNetworkAccessRequestPortArray{
								&security.JitNetworkAccessRequestPortArgs{
									AllowedSourceAddressPrefix: pulumi.String("192.127.0.2"),
									EndTimeUtc:                 pulumi.String("2018-05-17T09:06:45.5691611Z"),
									Number:                     pulumi.Int(3389),
									Status:                     pulumi.String("Initiated"),
									StatusReason:               pulumi.String("UserRequested"),
								},
							},
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("myRg1"),
			VirtualMachines: security.JitNetworkAccessPolicyVirtualMachineArray{
				&security.JitNetworkAccessPolicyVirtualMachineArgs{
					Id: pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
					Ports: security.JitNetworkAccessPortRuleArray{
						&security.JitNetworkAccessPortRuleArgs{
							AllowedSourceAddressPrefix: pulumi.String("*"),
							MaxRequestAccessDuration:   pulumi.String("PT3H"),
							Number:                     pulumi.Int(22),
							Protocol:                   pulumi.String("*"),
						},
						&security.JitNetworkAccessPortRuleArgs{
							AllowedSourceAddressPrefix: pulumi.String("*"),
							MaxRequestAccessDuration:   pulumi.String("PT3H"),
							Number:                     pulumi.Int(3389),
							Protocol:                   pulumi.String("*"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
