package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsmq.NewBrokerAuthorization(ctx, "brokerAuthorization", &iotoperationsmq.BrokerAuthorizationArgs{
			AuthorizationName: pulumi.String("C15G"),
			AuthorizationPolicies: &iotoperationsmq.AuthorizationConfigArgs{
				EnableCache: pulumi.Bool(true),
				Rules: iotoperationsmq.AuthorizationBasicRuleArray{
					&iotoperationsmq.AuthorizationBasicRuleArgs{
						BrokerResources: iotoperationsmq.ResourceInfoDefinitionArray{
							&iotoperationsmq.ResourceInfoDefinitionArgs{
								Method: pulumi.String("Connect"),
								Topics: pulumi.StringArray{
									pulumi.String("v"),
								},
							},
						},
						Principals: &iotoperationsmq.PrincipalDefinitionArgs{
							Attributes: pulumi.StringMapArray{
								nil,
							},
							Clientids: pulumi.StringArray{
								pulumi.String("smrfzvniq"),
							},
							Usernames: pulumi.StringArray{
								pulumi.String("jtwwmsrzriat"),
							},
						},
					},
				},
			},
			BrokerName: pulumi.String("7E0-tXS-6u1h-Vx396----"),
			ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
				Name: pulumi.String("an"),
				Type: pulumi.String("CustomLocation"),
			},
			ListenerRef: pulumi.StringArray{
				pulumi.String("mxgpbyb"),
			},
			Location:          pulumi.String("bvgohixie"),
			MqName:            pulumi.String("Zz22-b2VC-9"),
			ResourceGroupName: pulumi.String("rgiotoperationsmq"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
