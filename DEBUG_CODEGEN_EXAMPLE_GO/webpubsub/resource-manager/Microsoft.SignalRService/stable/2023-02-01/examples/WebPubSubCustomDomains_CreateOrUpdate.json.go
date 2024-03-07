package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/webpubsub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := webpubsub.NewWebPubSubCustomDomain(ctx, "webPubSubCustomDomain", &webpubsub.WebPubSubCustomDomainArgs{
			CustomCertificate: &webpubsub.ResourceReferenceArgs{
				Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/customCertificates/myCert"),
			},
			DomainName:        pulumi.String("example.com"),
			Name:              pulumi.String("myDomain"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("myWebPubSubService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
