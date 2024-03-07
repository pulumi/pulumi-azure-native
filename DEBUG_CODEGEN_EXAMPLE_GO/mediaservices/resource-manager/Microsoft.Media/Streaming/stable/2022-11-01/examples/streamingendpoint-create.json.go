package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewStreamingEndpoint(ctx, "streamingEndpoint", &media.StreamingEndpointArgs{
			AccessControl: &media.StreamingEndpointAccessControlArgs{
				Akamai: &media.AkamaiAccessControlArgs{
					AkamaiSignatureHeaderAuthenticationKeyList: media.AkamaiSignatureHeaderAuthenticationKeyArray{
						&media.AkamaiSignatureHeaderAuthenticationKeyArgs{
							Base64Key:  pulumi.String("dGVzdGlkMQ=="),
							Expiration: pulumi.String("2029-12-31T16:00:00-08:00"),
							Identifier: pulumi.String("id1"),
						},
						&media.AkamaiSignatureHeaderAuthenticationKeyArgs{
							Base64Key:  pulumi.String("dGVzdGlkMQ=="),
							Expiration: pulumi.String("2030-12-31T16:00:00-08:00"),
							Identifier: pulumi.String("id2"),
						},
					},
				},
				Ip: &media.IPAccessControlArgs{
					Allow: media.IPRangeArray{
						&media.IPRangeArgs{
							Address: pulumi.String("192.168.1.1"),
							Name:    pulumi.String("AllowedIp"),
						},
					},
				},
			},
			AccountName:           pulumi.String("slitestmedia10"),
			AvailabilitySetName:   pulumi.String("availableset"),
			CdnEnabled:            pulumi.Bool(false),
			Description:           pulumi.String("test event 1"),
			Location:              pulumi.String("West US"),
			ResourceGroupName:     pulumi.String("mediaresources"),
			ScaleUnits:            pulumi.Int(1),
			StreamingEndpointName: pulumi.String("myStreamingEndpoint1"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
