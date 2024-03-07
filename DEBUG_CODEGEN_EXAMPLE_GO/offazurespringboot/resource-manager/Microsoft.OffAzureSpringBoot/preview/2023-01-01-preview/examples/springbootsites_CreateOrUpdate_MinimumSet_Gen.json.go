package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazurespringboot/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazurespringboot.NewSpringbootsite(ctx, "springbootsite", &offazurespringboot.SpringbootsiteArgs{
			Location:            pulumi.String("tgobtvxktootwhhvjtsmpddvlqlrq"),
			ResourceGroupName:   pulumi.String("rgspringbootsites"),
			SpringbootsitesName: pulumi.String("xrmzlavpewxtfeitghdrj"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
