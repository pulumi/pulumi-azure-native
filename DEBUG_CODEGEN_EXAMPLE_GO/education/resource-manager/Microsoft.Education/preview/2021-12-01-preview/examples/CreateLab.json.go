package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/education/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := education.NewLab(ctx, "lab", &education.LabArgs{
			BillingAccountName: pulumi.String("{billingAccountName}"),
			BillingProfileName: pulumi.String("{billingProfileName}"),
			BudgetPerStudent: &education.AmountArgs{
				Currency: pulumi.String("USD"),
				Value:    pulumi.Float64(100),
			},
			Description:        pulumi.String("example lab description"),
			DisplayName:        pulumi.String("example lab"),
			ExpirationDate:     pulumi.String("2021-12-09T22:11:29.422Z"),
			InvoiceSectionName: pulumi.String("{invoiceSectionName}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
