package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewThreatIntelligenceIndicator(ctx, "threatIntelligenceIndicator", &securityinsights.ThreatIntelligenceIndicatorArgs{
			Confidence:         pulumi.Int(78),
			CreatedByRef:       pulumi.String("contoso@contoso.com"),
			Description:        pulumi.String("debugging indicators"),
			DisplayName:        pulumi.String("new schema"),
			ExternalReferences: securityinsights.ThreatIntelligenceExternalReferenceArray{},
			GranularMarkings:   securityinsights.ThreatIntelligenceGranularMarkingModelArray{},
			KillChainPhases:    securityinsights.ThreatIntelligenceKillChainPhaseArray{},
			Kind:               pulumi.String("indicator"),
			Labels:             pulumi.StringArray{},
			Modified:           pulumi.String(""),
			Name:               pulumi.String("d9cd6f0b-96b9-3984-17cd-a779d1e15a93"),
			Pattern:            pulumi.String("[url:value = 'https://www.contoso.com']"),
			PatternType:        pulumi.String("url"),
			ResourceGroupName:  pulumi.String("myRg"),
			Revoked:            pulumi.Bool(false),
			Source:             pulumi.String("Azure Sentinel"),
			ThreatIntelligenceTags: pulumi.StringArray{
				pulumi.String("new schema"),
			},
			ThreatTypes: pulumi.StringArray{
				pulumi.String("compromised"),
			},
			ValidFrom:     pulumi.String("2020-04-15T17:44:00.114052Z"),
			ValidUntil:    pulumi.String(""),
			WorkspaceName: pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
