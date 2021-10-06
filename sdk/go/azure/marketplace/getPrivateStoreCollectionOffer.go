


package marketplace

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateStoreCollectionOffer(ctx *pulumi.Context, args *LookupPrivateStoreCollectionOfferArgs, opts ...pulumi.InvokeOption) (*LookupPrivateStoreCollectionOfferResult, error) {
	var rv LookupPrivateStoreCollectionOfferResult
	err := ctx.Invoke("azure-native:marketplace:getPrivateStoreCollectionOffer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateStoreCollectionOfferArgs struct {
	CollectionId   string `pulumi:"collectionId"`
	OfferId        string `pulumi:"offerId"`
	PrivateStoreId string `pulumi:"privateStoreId"`
}


type LookupPrivateStoreCollectionOfferResult struct {
	CreatedAt                      string             `pulumi:"createdAt"`
	ETag                           *string            `pulumi:"eTag"`
	IconFileUris                   map[string]string  `pulumi:"iconFileUris"`
	Id                             string             `pulumi:"id"`
	ModifiedAt                     string             `pulumi:"modifiedAt"`
	Name                           string             `pulumi:"name"`
	OfferDisplayName               string             `pulumi:"offerDisplayName"`
	Plans                          []PlanResponse     `pulumi:"plans"`
	PrivateStoreId                 string             `pulumi:"privateStoreId"`
	PublisherDisplayName           string             `pulumi:"publisherDisplayName"`
	SpecificPlanIdsLimitation      []string           `pulumi:"specificPlanIdsLimitation"`
	SystemData                     SystemDataResponse `pulumi:"systemData"`
	Type                           string             `pulumi:"type"`
	UniqueOfferId                  string             `pulumi:"uniqueOfferId"`
	UpdateSuppressedDueIdempotence *bool              `pulumi:"updateSuppressedDueIdempotence"`
}
