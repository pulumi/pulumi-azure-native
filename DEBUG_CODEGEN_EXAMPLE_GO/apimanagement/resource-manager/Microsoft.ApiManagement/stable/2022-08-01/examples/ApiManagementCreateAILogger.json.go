package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewLogger(ctx, "logger", &apimanagement.LoggerArgs{
			Credentials: pulumi.StringMap{
				"instrumentationKey": pulumi.String("11................a1"),
			},
			Description:       pulumi.String("adding a new logger"),
			LoggerId:          pulumi.String("loggerId"),
			LoggerType:        pulumi.String("applicationInsights"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
