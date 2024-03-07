package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewBuildServiceBuilder(ctx, "buildServiceBuilder", &appplatform.BuildServiceBuilderArgs{
			BuildServiceName: pulumi.String("default"),
			BuilderName:      pulumi.String("mybuilder"),
			Properties: &appplatform.BuilderPropertiesArgs{
				BuildpackGroups: appplatform.BuildpacksGroupPropertiesArray{
					&appplatform.BuildpacksGroupPropertiesArgs{
						Buildpacks: appplatform.BuildpackPropertiesArray{
							&appplatform.BuildpackPropertiesArgs{
								Id: pulumi.String("tanzu-buildpacks/java-azure"),
							},
						},
						Name: pulumi.String("mix"),
					},
				},
				Stack: &appplatform.StackPropertiesArgs{
					Id:      pulumi.String("io.buildpacks.stacks.bionic"),
					Version: pulumi.String("base"),
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
