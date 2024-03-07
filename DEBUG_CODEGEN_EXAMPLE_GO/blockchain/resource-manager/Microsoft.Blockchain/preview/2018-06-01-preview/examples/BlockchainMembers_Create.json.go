package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blockchain/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blockchain.NewBlockchainMember(ctx, "blockchainMember", &blockchain.BlockchainMemberArgs{
			BlockchainMemberName:                pulumi.String("contosemember1"),
			Consortium:                          pulumi.String("ContoseConsortium"),
			ConsortiumManagementAccountPassword: pulumi.String("<consortiumManagementAccountPassword>"),
			Location:                            pulumi.String("southeastasia"),
			Password:                            pulumi.String("<password>"),
			Protocol:                            pulumi.String("Quorum"),
			ResourceGroupName:                   pulumi.String("mygroup"),
			ValidatorNodesSku: &blockchain.BlockchainMemberNodesSkuArgs{
				Capacity: pulumi.Int(2),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
