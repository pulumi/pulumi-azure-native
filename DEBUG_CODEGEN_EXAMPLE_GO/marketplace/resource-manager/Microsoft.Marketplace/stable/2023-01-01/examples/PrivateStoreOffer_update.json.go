package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/marketplace/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := marketplace.NewPrivateStoreCollectionOffer(ctx, "privateStoreCollectionOffer", &marketplace.PrivateStoreCollectionOfferArgs{
			CollectionId:   pulumi.String("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
			ETag:           pulumi.String("\"9301f4fd-0000-0100-0000-5e248b350666\""),
			OfferId:        pulumi.String("marketplacetestthirdparty.md-test-third-party-2"),
			PrivateStoreId: pulumi.String("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
			SpecificPlanIdsLimitation: pulumi.StringArray{
				pulumi.String("0001"),
				pulumi.String("0002"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
