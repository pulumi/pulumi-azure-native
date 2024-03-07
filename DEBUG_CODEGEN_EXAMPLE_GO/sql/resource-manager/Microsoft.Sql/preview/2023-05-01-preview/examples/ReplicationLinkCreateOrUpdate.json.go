package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewReplicationLink(ctx, "replicationLink", &sql.ReplicationLinkArgs{
			DatabaseName:      pulumi.String("gamma-db"),
			LinkId:            pulumi.String("00000000-1111-2222-3333-666666666666"),
			LinkType:          pulumi.String("STANDBY"),
			ResourceGroupName: pulumi.String("Default"),
			ServerName:        pulumi.String("sourcesvr"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
