package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/testbase/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := testbase.NewCustomerEvent(ctx, "customerEvent", &testbase.CustomerEventArgs{
			CustomerEventName: pulumi.String("WeeklySummary"),
			EventName:         pulumi.String("WeeklySummary"),
			Receivers: []testbase.NotificationEventReceiverArgs{
				{
					ReceiverType: pulumi.String("UserObjects"),
					ReceiverValue: {
						UserObjectReceiverValue: {
							UserObjectIds: pulumi.StringArray{
								pulumi.String("245245245245325"),
								pulumi.String("365365365363565"),
							},
						},
					},
				},
				{
					ReceiverType: pulumi.String("DistributionGroup"),
					ReceiverValue: {
						DistributionGroupListReceiverValue: {
							DistributionGroups: pulumi.StringArray{
								pulumi.String("test@microsoft.com"),
							},
						},
					},
				},
			},
			ResourceGroupName:   pulumi.String("contoso-rg1"),
			TestBaseAccountName: pulumi.String("contoso-testBaseAccount1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
