


package marketplace

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateStoreCollection(ctx *pulumi.Context, args *LookupPrivateStoreCollectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateStoreCollectionResult, error) {
	var rv LookupPrivateStoreCollectionResult
	err := ctx.Invoke("azure-native:marketplace:getPrivateStoreCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateStoreCollectionArgs struct {
	CollectionId   string `pulumi:"collectionId"`
	PrivateStoreId string `pulumi:"privateStoreId"`
}


type LookupPrivateStoreCollectionResult struct {
	AllSubscriptions  *bool              `pulumi:"allSubscriptions"`
	Claim             *string            `pulumi:"claim"`
	CollectionId      string             `pulumi:"collectionId"`
	CollectionName    *string            `pulumi:"collectionName"`
	Enabled           *bool              `pulumi:"enabled"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	NumberOfOffers    float64            `pulumi:"numberOfOffers"`
	SubscriptionsList []string           `pulumi:"subscriptionsList"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}
