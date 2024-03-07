package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewBuildpackBinding(ctx, "buildpackBinding", &appplatform.BuildpackBindingArgs{
			BuildServiceName:     pulumi.String("default"),
			BuilderName:          pulumi.String("default"),
			BuildpackBindingName: pulumi.String("myBuildpackBinding"),
			Properties: &appplatform.BuildpackBindingPropertiesArgs{
				BindingType: pulumi.String("ApplicationInsights"),
				LaunchProperties: &appplatform.BuildpackBindingLaunchPropertiesArgs{
					Properties: pulumi.StringMap{
						"abc":           pulumi.String("def"),
						"any-string":    pulumi.String("any-string"),
						"sampling-rate": pulumi.String("12.0"),
					},
					Secrets: pulumi.StringMap{
						"connection-string": pulumi.String("XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXX;XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXXXXXXXX"),
					},
				},
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
