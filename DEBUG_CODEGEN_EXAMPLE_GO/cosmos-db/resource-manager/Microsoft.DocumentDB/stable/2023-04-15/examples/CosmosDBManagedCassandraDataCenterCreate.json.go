package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewCassandraDataCenter(ctx, "cassandraDataCenter", &documentdb.CassandraDataCenterArgs{
			ClusterName:    pulumi.String("cassandra-prod"),
			DataCenterName: pulumi.String("dc1"),
			Properties: &documentdb.DataCenterResourcePropertiesArgs{
				Base64EncodedCassandraYamlFragment: pulumi.String("Y29tcGFjdGlvbl90aHJvdWdocHV0X21iX3Blcl9zZWM6IDMyCmNvbXBhY3Rpb25fbGFyZ2VfcGFydGl0aW9uX3dhcm5pbmdfdGhyZXNob2xkX21iOiAxMDA="),
				DataCenterLocation:                 pulumi.String("West US 2"),
				DelegatedSubnetId:                  pulumi.String("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1-subnet"),
				NodeCount:                          pulumi.Int(9),
			},
			ResourceGroupName: pulumi.String("cassandra-prod-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
