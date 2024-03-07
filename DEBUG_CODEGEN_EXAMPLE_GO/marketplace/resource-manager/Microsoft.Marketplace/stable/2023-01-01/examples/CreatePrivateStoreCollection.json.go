package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/marketplace/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := marketplace.NewPrivateStoreCollection(ctx, "privateStoreCollection", &marketplace.PrivateStoreCollectionArgs{
			AllSubscriptions: pulumi.Bool(false),
			Claim:            pulumi.String(""),
			CollectionId:     pulumi.String("d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1"),
			CollectionName:   pulumi.String("Test Collection"),
			PrivateStoreId:   pulumi.String("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
			SubscriptionsList: pulumi.StringArray{
				pulumi.String("b340914e-353d-453a-85fb-8f9b65b51f91"),
				pulumi.String("f2baa04d-5bfc-461b-b6d8-61b403c9ec48"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
