


package marketplace

import (
	"context"
	"reflect"

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

func LookupPrivateStoreCollectionOfferOutput(ctx *pulumi.Context, args LookupPrivateStoreCollectionOfferOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateStoreCollectionOfferResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateStoreCollectionOfferResult, error) {
			args := v.(LookupPrivateStoreCollectionOfferArgs)
			r, err := LookupPrivateStoreCollectionOffer(ctx, &args, opts...)
			var s LookupPrivateStoreCollectionOfferResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateStoreCollectionOfferResultOutput)
}

type LookupPrivateStoreCollectionOfferOutputArgs struct {
	CollectionId   pulumi.StringInput `pulumi:"collectionId"`
	OfferId        pulumi.StringInput `pulumi:"offerId"`
	PrivateStoreId pulumi.StringInput `pulumi:"privateStoreId"`
}

func (LookupPrivateStoreCollectionOfferOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateStoreCollectionOfferArgs)(nil)).Elem()
}


type LookupPrivateStoreCollectionOfferResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateStoreCollectionOfferResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateStoreCollectionOfferResult)(nil)).Elem()
}

func (o LookupPrivateStoreCollectionOfferResultOutput) ToLookupPrivateStoreCollectionOfferResultOutput() LookupPrivateStoreCollectionOfferResultOutput {
	return o
}

func (o LookupPrivateStoreCollectionOfferResultOutput) ToLookupPrivateStoreCollectionOfferResultOutputWithContext(ctx context.Context) LookupPrivateStoreCollectionOfferResultOutput {
	return o
}

func (o LookupPrivateStoreCollectionOfferResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) IconFileUris() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) map[string]string { return v.IconFileUris }).(pulumi.StringMapOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.ModifiedAt }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) OfferDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.OfferDisplayName }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) Plans() PlanResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) []PlanResponse { return v.Plans }).(PlanResponseArrayOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) PrivateStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.PrivateStoreId }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) PublisherDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.PublisherDisplayName }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) SpecificPlanIdsLimitation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) []string { return v.SpecificPlanIdsLimitation }).(pulumi.StringArrayOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) UniqueOfferId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) string { return v.UniqueOfferId }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionOfferResultOutput) UpdateSuppressedDueIdempotence() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionOfferResult) *bool { return v.UpdateSuppressedDueIdempotence }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateStoreCollectionOfferResultOutput{})
}
