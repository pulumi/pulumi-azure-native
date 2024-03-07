package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/signalrservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := signalrservice.NewSignalRCustomDomain(ctx, "signalRCustomDomain", &signalrservice.SignalRCustomDomainArgs{
			CustomCertificate: &signalrservice.ResourceReferenceArgs{
				Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/customCertificates/myCert"),
			},
			DomainName:        pulumi.String("example.com"),
			Name:              pulumi.String("myDomain"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("mySignalRService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
