package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/communication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := communication.NewSuppressionList(ctx, "suppressionList", &communication.SuppressionListArgs{
			DomainName:          pulumi.String("contoso.com"),
			EmailServiceName:    pulumi.String("contosoEmailService"),
			ListName:            pulumi.String("contosoNewsAlerts"),
			ResourceGroupName:   pulumi.String("contosoResourceGroup"),
			SuppressionListName: pulumi.String("aaaa1111-bbbb-2222-3333-aaaa11112222"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
