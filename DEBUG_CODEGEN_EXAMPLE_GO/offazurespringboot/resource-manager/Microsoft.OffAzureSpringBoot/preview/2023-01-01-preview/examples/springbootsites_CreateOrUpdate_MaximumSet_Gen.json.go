package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazurespringboot/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazurespringboot.NewSpringbootsite(ctx, "springbootsite", &offazurespringboot.SpringbootsiteArgs{
			ExtendedLocation: &offazurespringboot.SpringbootsitesModelExtendedLocationArgs{
				Name: pulumi.String("rywvpbfsqovhlfirtwisugsdsfsgf"),
				Type: pulumi.String("lvsb"),
			},
			Location: pulumi.String("tgobtvxktootwhhvjtsmpddvlqlrq"),
			Properties: &offazurespringboot.SpringbootsitesPropertiesArgs{
				MasterSiteId:     pulumi.String("xsoimrgshsactearljwuljmi"),
				MigrateProjectId: pulumi.String("wwuattybgco"),
			},
			ResourceGroupName:   pulumi.String("rgspringbootsites"),
			SpringbootsitesName: pulumi.String("xrmzlavpewxtfeitghdrj"),
			Tags: pulumi.StringMap{
				"key3558": pulumi.String("xeuhtglamqzj"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
