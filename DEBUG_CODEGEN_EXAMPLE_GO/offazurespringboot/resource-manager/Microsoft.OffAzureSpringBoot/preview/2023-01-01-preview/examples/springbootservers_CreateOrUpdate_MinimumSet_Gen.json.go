package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazurespringboot/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazurespringboot.NewSpringbootserver(ctx, "springbootserver", &offazurespringboot.SpringbootserverArgs{
			ResourceGroupName:     pulumi.String("rgspringbootservers"),
			SiteName:              pulumi.String("hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj"),
			SpringbootserversName: pulumi.String("zkarbqnwnxeozvjrkpdqmgnwedwgtwcmmyqwaijkn"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
