package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewCassandraCluster(ctx, "cassandraCluster", &documentdb.CassandraClusterArgs{
			ClusterName: pulumi.String("cassandra-prod"),
			Location:    pulumi.String("West US"),
			Properties: &documentdb.ClusterResourcePropertiesArgs{
				AuthenticationMethod: pulumi.String("Cassandra"),
				CassandraVersion:     pulumi.String("3.11"),
				ClientCertificates: documentdb.CertificateArray{
					&documentdb.CertificateArgs{
						Pem: pulumi.String("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
					},
				},
				ClusterNameOverride:         pulumi.String("ClusterNameIllegalForAzureResource"),
				DelegatedManagementSubnetId: pulumi.String("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
				ExternalGossipCertificates: documentdb.CertificateArray{
					&documentdb.CertificateArgs{
						Pem: pulumi.String("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
					},
				},
				ExternalSeedNodes: documentdb.SeedNodeArray{
					&documentdb.SeedNodeArgs{
						IpAddress: pulumi.String("10.52.221.2"),
					},
					&documentdb.SeedNodeArgs{
						IpAddress: pulumi.String("10.52.221.3"),
					},
					&documentdb.SeedNodeArgs{
						IpAddress: pulumi.String("10.52.221.4"),
					},
				},
				HoursBetweenBackups:           pulumi.Int(24),
				InitialCassandraAdminPassword: pulumi.String("mypassword"),
			},
			ResourceGroupName: pulumi.String("cassandra-prod-rg"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
