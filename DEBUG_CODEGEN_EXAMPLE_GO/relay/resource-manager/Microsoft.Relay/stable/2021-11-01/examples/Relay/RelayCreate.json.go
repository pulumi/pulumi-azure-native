package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/relay/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := relay.NewWCFRelay(ctx, "wcfRelay", &relay.WCFRelayArgs{
			NamespaceName:               pulumi.String("example-RelayNamespace-9953"),
			RelayName:                   pulumi.String("example-Relay-Wcf-1194"),
			RelayType:                   relay.RelaytypeNetTcp,
			RequiresClientAuthorization: pulumi.Bool(true),
			RequiresTransportSecurity:   pulumi.Bool(true),
			ResourceGroupName:           pulumi.String("resourcegroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
