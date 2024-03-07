package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewRaiPolicy(ctx, "raiPolicy", &cognitiveservices.RaiPolicyArgs{
			AccountName: pulumi.String("accountName"),
			Properties: cognitiveservices.RaiPolicyPropertiesResponse{
				BasePolicyName: pulumi.String("112"),
				ContentFilters: cognitiveservices.RaiPolicyContentFilterArray{
					&cognitiveservices.RaiPolicyContentFilterArgs{
						AllowedContentLevel: pulumi.String("Low"),
						Blocking:            pulumi.Bool(true),
						Name:                pulumi.String("hate"),
					},
					&cognitiveservices.RaiPolicyContentFilterArgs{
						AllowedContentLevel: pulumi.String("Low"),
						Name:                pulumi.String("sexual"),
					},
					&cognitiveservices.RaiPolicyContentFilterArgs{
						Enabled: pulumi.Bool(false),
						Name:    pulumi.String("violence"),
					},
					&cognitiveservices.RaiPolicyContentFilterArgs{
						Enabled: pulumi.Bool(false),
						Name:    pulumi.String("DefaultHateSpeechBlockList"),
					},
				},
			},
			RaiPolicyName:     pulumi.String("raiPolicyName"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
