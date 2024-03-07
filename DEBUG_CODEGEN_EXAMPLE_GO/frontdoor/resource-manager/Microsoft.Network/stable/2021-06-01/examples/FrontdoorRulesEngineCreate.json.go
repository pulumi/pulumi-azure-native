package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRulesEngine(ctx, "rulesEngine", &network.RulesEngineArgs{
			FrontDoorName:     pulumi.String("frontDoor1"),
			ResourceGroupName: pulumi.String("rg1"),
			Rules: []network.RulesEngineRuleArgs{
				{
					Action: {
						RouteConfigurationOverride: {
							CustomFragment:    "fragment",
							CustomHost:        "www.bing.com",
							CustomPath:        "/api",
							CustomQueryString: "a=b",
							OdataType:         "#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration",
							RedirectProtocol:  "HttpsOnly",
							RedirectType:      "Moved",
						},
					},
					MatchConditions: network.RulesEngineMatchConditionArray{
						{
							RulesEngineMatchValue: pulumi.StringArray{
								pulumi.String("CH"),
							},
							RulesEngineMatchVariable: pulumi.String("RemoteAddr"),
							RulesEngineOperator:      pulumi.String("GeoMatch"),
						},
					},
					MatchProcessingBehavior: pulumi.String("Stop"),
					Name:                    pulumi.String("Rule1"),
					Priority:                pulumi.Int(1),
				},
				{
					Action: {
						ResponseHeaderActions: network.HeaderActionArray{
							{
								HeaderActionType: pulumi.String("Overwrite"),
								HeaderName:       pulumi.String("Cache-Control"),
								Value:            pulumi.String("public, max-age=31536000"),
							},
						},
					},
					MatchConditions: network.RulesEngineMatchConditionArray{
						{
							RulesEngineMatchValue: pulumi.StringArray{
								pulumi.String("jpg"),
							},
							RulesEngineMatchVariable: pulumi.String("RequestFilenameExtension"),
							RulesEngineOperator:      pulumi.String("Equal"),
							Transforms: pulumi.StringArray{
								pulumi.String("Lowercase"),
							},
						},
					},
					Name:     pulumi.String("Rule2"),
					Priority: pulumi.Int(2),
				},
				{
					Action: {
						RouteConfigurationOverride: {
							BackendPool: {
								Id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1",
							},
							CacheConfiguration: {
								CacheDuration:                "P1DT12H20M30S",
								DynamicCompression:           "Disabled",
								QueryParameterStripDirective: "StripOnly",
								QueryParameters:              "a=b,p=q",
							},
							ForwardingProtocol: "HttpsOnly",
							OdataType:          "#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration",
						},
					},
					MatchConditions: network.RulesEngineMatchConditionArray{
						{
							NegateCondition: pulumi.Bool(false),
							RulesEngineMatchValue: pulumi.StringArray{
								pulumi.String("allowoverride"),
							},
							RulesEngineMatchVariable: pulumi.String("RequestHeader"),
							RulesEngineOperator:      pulumi.String("Equal"),
							Selector:                 pulumi.String("Rules-Engine-Route-Forward"),
							Transforms: pulumi.StringArray{
								pulumi.String("Lowercase"),
							},
						},
					},
					Name:     pulumi.String("Rule3"),
					Priority: pulumi.Int(3),
				},
			},
			RulesEngineName: pulumi.String("rulesEngine1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
