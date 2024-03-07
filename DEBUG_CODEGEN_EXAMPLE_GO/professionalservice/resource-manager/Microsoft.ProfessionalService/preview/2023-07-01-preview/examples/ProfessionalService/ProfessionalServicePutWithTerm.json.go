package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/professionalservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := professionalservice.NewProfessionalServiceSubscriptionLevel(ctx, "professionalServiceSubscriptionLevel", &professionalservice.ProfessionalServiceSubscriptionLevelArgs{
			Location: pulumi.String("global"),
			Name:     pulumi.String("MyContosoPS"),
			Properties: &professionalservice.ProfessionalServiceCreationPropertiesArgs{
				BillingPeriod: pulumi.String("P1Y"),
				OfferId:       pulumi.String("testprofservice"),
				PublisherId:   pulumi.String("microsoft-contoso"),
				QuoteId:       pulumi.String("quoteabc"),
				SkuId:         pulumi.String("ff051f4f-a6d9-4cbc-8d9a-2a41bd468abc"),
				TermUnit:      pulumi.String("P3Y"),
			},
			ResourceGroupName: pulumi.String("my-ps-rg"),
			ResourceName:      pulumi.String("MyContosoPS"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
