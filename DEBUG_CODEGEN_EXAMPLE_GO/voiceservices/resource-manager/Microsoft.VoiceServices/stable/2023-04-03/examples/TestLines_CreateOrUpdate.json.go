package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/voiceservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := voiceservices.NewTestLine(ctx, "testLine", &voiceservices.TestLineArgs{
			CommunicationsGatewayName: pulumi.String("myname"),
			Location:                  pulumi.String("useast"),
			PhoneNumber:               pulumi.String("+1-555-1234"),
			Purpose:                   pulumi.String("Automated"),
			ResourceGroupName:         pulumi.String("testrg"),
			TestLineName:              pulumi.String("myline"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
