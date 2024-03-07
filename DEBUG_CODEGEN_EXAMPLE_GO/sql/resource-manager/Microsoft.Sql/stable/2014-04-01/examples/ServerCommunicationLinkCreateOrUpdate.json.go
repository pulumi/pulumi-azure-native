package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServerCommunicationLink(ctx, "serverCommunicationLink", &sql.ServerCommunicationLinkArgs{
			CommunicationLinkName: pulumi.String("link1"),
			PartnerServer:         pulumi.String("sqldcrudtest-test"),
			ResourceGroupName:     pulumi.String("sqlcrudtest-7398"),
			ServerName:            pulumi.String("sqlcrudtest-4645"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
