package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazurespringboot/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazurespringboot.NewSpringbootserver(ctx, "springbootserver", &offazurespringboot.SpringbootserverArgs{
			Properties: offazurespringboot.SpringbootserversPropertiesResponse{
				Errors:               offazurespringboot.ErrorArray{},
				FqdnAndIpAddressList: pulumi.StringArray{},
				MachineArmId:         pulumi.String("fvfkiapbqsprnbzczdfmuryknrna"),
				Port:                 pulumi.Int(10),
				Server:               pulumi.String("thhuxocfyqpeluqcgnypi"),
				SpringBootApps:       pulumi.Int(17),
				TotalApps:            pulumi.Int(5),
			},
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
