package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewSecurityOperator(ctx, "securityOperator", &security.SecurityOperatorArgs{
			PricingName:          pulumi.String("CloudPosture"),
			SecurityOperatorName: pulumi.String("DefenderCSPMSecurityOperator"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
