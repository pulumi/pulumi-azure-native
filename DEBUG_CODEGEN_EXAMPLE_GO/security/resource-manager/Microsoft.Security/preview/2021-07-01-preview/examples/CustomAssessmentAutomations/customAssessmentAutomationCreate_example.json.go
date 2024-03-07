package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewCustomAssessmentAutomation(ctx, "customAssessmentAutomation", &security.CustomAssessmentAutomationArgs{
			CompressedQuery:                pulumi.String("DQAKAEkAYQBtAF8ARwByAG8AdQBwAA0ACgB8ACAAZQB4AHQAZQBuAGQAIABIAGUAYQBsAHQAaABTAHQAYQB0AHUAcwAgAD0AIABpAGYAZgAoAHQAbwBzAHQAcgBpAG4AZwAoAFIAZQBjAG8AcgBkAC4AVQBzAGUAcgBOAGEAbQBlACkAIABjAG8AbgB0AGEAaQBuAHMAIAAnAHUAcwBlAHIAJwAsACAAJwBVAE4ASABFAEEATABUAEgAWQAnACwAIAAnAEgARQBBAEwAVABIAFkAJwApAA0ACgA="),
			CustomAssessmentAutomationName: pulumi.String("MyCustomAssessmentAutomation"),
			Description:                    pulumi.String("Data should be encrypted"),
			DisplayName:                    pulumi.String("Password Policy"),
			RemediationDescription:         pulumi.String("Encrypt store by..."),
			ResourceGroupName:              pulumi.String("TestResourceGroup"),
			Severity:                       pulumi.String("Medium"),
			SupportedCloud:                 pulumi.String("AWS"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
