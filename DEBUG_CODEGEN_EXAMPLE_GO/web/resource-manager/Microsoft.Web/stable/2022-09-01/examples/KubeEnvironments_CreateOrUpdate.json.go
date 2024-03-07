package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewKubeEnvironment(ctx, "kubeEnvironment", &web.KubeEnvironmentArgs{
			Location:          pulumi.String("East US"),
			Name:              pulumi.String("testkubeenv"),
			ResourceGroupName: pulumi.String("examplerg"),
			StaticIp:          pulumi.String("1.2.3.4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
