package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/policyinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := policyinsights.NewAttestationAtSubscription(ctx, "attestationAtSubscription", &policyinsights.AttestationAtSubscriptionArgs{
			AttestationName:    pulumi.String("790996e6-9871-4b1f-9cd9-ec42cd6ced1e"),
			ComplianceState:    pulumi.String("Compliant"),
			PolicyAssignmentId: pulumi.String("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
