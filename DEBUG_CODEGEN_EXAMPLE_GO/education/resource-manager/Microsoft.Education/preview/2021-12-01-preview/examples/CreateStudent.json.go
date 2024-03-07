package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/education/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := education.NewStudent(ctx, "student", &education.StudentArgs{
			BillingAccountName: pulumi.String("{billingAccountName}"),
			BillingProfileName: pulumi.String("{billingProfileName}"),
			Budget: &education.AmountArgs{
				Currency: pulumi.String("USD"),
				Value:    pulumi.Float64(100),
			},
			Email:              pulumi.String("test@contoso.com"),
			ExpirationDate:     pulumi.String("2021-11-09T22:13:21.795Z"),
			FirstName:          pulumi.String("test"),
			InvoiceSectionName: pulumi.String("{invoiceSectionName}"),
			LastName:           pulumi.String("user"),
			Role:               pulumi.String("Student"),
			StudentAlias:       pulumi.String("{studentAlias}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
