package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/policyinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := policyinsights.NewAttestationAtResourceGroup(ctx, "attestationAtResourceGroup", &policyinsights.AttestationAtResourceGroupArgs{
			AssessmentDate:  pulumi.String("2021-06-10T00:00:00Z"),
			AttestationName: pulumi.String("790996e6-9871-4b1f-9cd9-ec42cd6ced1e"),
			Comments:        pulumi.String("This subscription has passed a security audit."),
			ComplianceState: pulumi.String("Compliant"),
			Evidence: []policyinsights.AttestationEvidenceArgs{
				{
					Description: pulumi.String("The results of the security audit."),
					SourceUri:   pulumi.String("https://gist.github.com/contoso/9573e238762c60166c090ae16b814011"),
				},
			},
			ExpiresOn: pulumi.String("2021-06-15T00:00:00Z"),
			Metadata: pulumi.Any{
				DepartmentId: "NYC-MARKETING-1",
			},
			Owner:                       pulumi.String("55a32e28-3aa5-4eea-9b5a-4cd85153b966"),
			PolicyAssignmentId:          pulumi.String("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
			PolicyDefinitionReferenceId: pulumi.String("0b158b46-ff42-4799-8e39-08a5c23b4551"),
			ResourceGroupName:           pulumi.String("myRg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
