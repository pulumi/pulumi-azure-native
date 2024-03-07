package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blockchain/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blockchain.NewTransactionNode(ctx, "transactionNode", &blockchain.TransactionNodeArgs{
			BlockchainMemberName: pulumi.String("contosemember1"),
			Location:             pulumi.String("southeastasia"),
			Password:             pulumi.String("<password>"),
			ResourceGroupName:    pulumi.String("mygroup"),
			TransactionNodeName:  pulumi.String("txnode2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
