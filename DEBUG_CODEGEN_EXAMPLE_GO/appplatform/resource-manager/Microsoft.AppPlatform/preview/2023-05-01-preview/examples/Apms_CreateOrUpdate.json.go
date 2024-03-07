package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewApm(ctx, "apm", &appplatform.ApmArgs{
			ApmName: pulumi.String("myappinsights"),
			Properties: &appplatform.ApmPropertiesArgs{
				Properties: pulumi.StringMap{
					"any-string":    pulumi.String("any-string"),
					"sampling-rate": pulumi.String("12.0"),
				},
				Secrets: pulumi.StringMap{
					"connection-string": pulumi.String("XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXX;XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXXXXXXXX"),
				},
				Type: pulumi.String("ApplicationInsights"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
