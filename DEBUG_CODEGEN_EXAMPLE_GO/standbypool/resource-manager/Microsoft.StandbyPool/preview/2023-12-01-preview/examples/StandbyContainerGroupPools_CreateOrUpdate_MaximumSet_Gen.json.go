package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/standbypool/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := standbypool.NewStandbyContainerGroupPool(ctx, "standbyContainerGroupPool", &standbypool.StandbyContainerGroupPoolArgs{
			ContainerGroupProperties: standbypool.ContainerGroupPropertiesResponse{
				ContainerGroupProfile: &standbypool.ContainerGroupProfileArgs{
					Id:       pulumi.String("/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile"),
					Revision: pulumi.Float64(1),
				},
				SubnetIds: standbypool.SubnetArray{
					&standbypool.SubnetArgs{
						Id: pulumi.String("/subscriptions/8cf6c1b6-c80f-437c-87ad-45fbaff54f73/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet"),
					},
				},
			},
			ElasticityProfile: &standbypool.StandbyContainerGroupPoolElasticityProfileArgs{
				MaxReadyCapacity: pulumi.Float64(688),
				RefillPolicy:     pulumi.String("always"),
			},
			Location:                      pulumi.String("West US"),
			ResourceGroupName:             pulumi.String("rgstandbypool"),
			StandbyContainerGroupPoolName: pulumi.String("pool"),
			Tags:                          nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
