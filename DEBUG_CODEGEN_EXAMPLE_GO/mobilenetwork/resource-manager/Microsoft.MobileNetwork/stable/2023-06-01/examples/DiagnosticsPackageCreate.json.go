package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewDiagnosticsPackage(ctx, "diagnosticsPackage", &mobilenetwork.DiagnosticsPackageArgs{
			DiagnosticsPackageName:     pulumi.String("dp1"),
			PacketCoreControlPlaneName: pulumi.String("TestPacketCoreCP"),
			ResourceGroupName:          pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
