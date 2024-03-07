package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationAccountSession(ctx, "integrationAccountSession", &logic.IntegrationAccountSessionArgs{
			Content: pulumi.Any{
				ControlNumber:            "1234",
				ControlNumberChangedTime: "2017-02-21T22:30:11.9923759Z",
			},
			IntegrationAccountName: pulumi.String("testia123"),
			ResourceGroupName:      pulumi.String("testrg123"),
			SessionName:            pulumi.String("testsession123-ICN"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
