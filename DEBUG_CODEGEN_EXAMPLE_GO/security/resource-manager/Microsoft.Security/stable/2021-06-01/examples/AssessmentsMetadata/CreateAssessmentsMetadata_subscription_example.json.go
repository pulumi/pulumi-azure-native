package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewAssessmentMetadataInSubscription(ctx, "assessmentMetadataInSubscription", &security.AssessmentMetadataInSubscriptionArgs{
			AssessmentMetadataName: pulumi.String("ca039e75-a276-4175-aebc-bcd41e4b14b7"),
			AssessmentType:         pulumi.String("CustomerManaged"),
			Categories: pulumi.StringArray{
				pulumi.String("Compute"),
			},
			Description:            pulumi.String("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
			DisplayName:            pulumi.String("Install endpoint protection solution on virtual machine scale sets"),
			ImplementationEffort:   pulumi.String("Low"),
			RemediationDescription: pulumi.String("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
			Severity:               pulumi.String("Medium"),
			Threats: pulumi.StringArray{
				pulumi.String("dataExfiltration"),
				pulumi.String("dataSpillage"),
				pulumi.String("maliciousInsider"),
			},
			UserImpact: pulumi.String("Low"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
