


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateStoreOffer(ctx *pulumi.Context, args *LookupPrivateStoreOfferArgs, opts ...pulumi.InvokeOption) (*LookupPrivateStoreOfferResult, error) {
	var rv LookupPrivateStoreOfferResult
	err := ctx.Invoke("azure-native:marketplace/v20200101:getPrivateStoreOffer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateStoreOfferArgs struct {
	OfferId        string `pulumi:"offerId"`
	PrivateStoreId string `pulumi:"privateStoreId"`
}


type LookupPrivateStoreOfferResult struct {
	CreatedAt                      string            `pulumi:"createdAt"`
	ETag                           *string           `pulumi:"eTag"`
	IconFileUris                   map[string]string `pulumi:"iconFileUris"`
	Id                             string            `pulumi:"id"`
	ModifiedAt                     string            `pulumi:"modifiedAt"`
	Name                           string            `pulumi:"name"`
	OfferDisplayName               string            `pulumi:"offerDisplayName"`
	Plans                          []PlanResponse    `pulumi:"plans"`
	PrivateStoreId                 string            `pulumi:"privateStoreId"`
	PublisherDisplayName           string            `pulumi:"publisherDisplayName"`
	SpecificPlanIdsLimitation      []string          `pulumi:"specificPlanIdsLimitation"`
	Type                           string            `pulumi:"type"`
	UniqueOfferId                  string            `pulumi:"uniqueOfferId"`
	UpdateSuppressedDueIdempotence *bool             `pulumi:"updateSuppressedDueIdempotence"`
}

func LookupPrivateStoreOfferOutput(ctx *pulumi.Context, args LookupPrivateStoreOfferOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateStoreOfferResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateStoreOfferResult, error) {
			args := v.(LookupPrivateStoreOfferArgs)
			r, err := LookupPrivateStoreOffer(ctx, &args, opts...)
			var s LookupPrivateStoreOfferResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateStoreOfferResultOutput)
}

type LookupPrivateStoreOfferOutputArgs struct {
	OfferId        pulumi.StringInput `pulumi:"offerId"`
	PrivateStoreId pulumi.StringInput `pulumi:"privateStoreId"`
}

func (LookupPrivateStoreOfferOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateStoreOfferArgs)(nil)).Elem()
}


type LookupPrivateStoreOfferResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateStoreOfferResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateStoreOfferResult)(nil)).Elem()
}

func (o LookupPrivateStoreOfferResultOutput) ToLookupPrivateStoreOfferResultOutput() LookupPrivateStoreOfferResultOutput {
	return o
}

func (o LookupPrivateStoreOfferResultOutput) ToLookupPrivateStoreOfferResultOutputWithContext(ctx context.Context) LookupPrivateStoreOfferResultOutput {
	return o
}

func (o LookupPrivateStoreOfferResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateStoreOfferResultOutput) IconFileUris() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) map[string]string { return v.IconFileUris }).(pulumi.StringMapOutput)
}

func (o LookupPrivateStoreOfferResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.ModifiedAt }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) OfferDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.OfferDisplayName }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) Plans() PlanResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) []PlanResponse { return v.Plans }).(PlanResponseArrayOutput)
}

func (o LookupPrivateStoreOfferResultOutput) PrivateStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.PrivateStoreId }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) PublisherDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.PublisherDisplayName }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) SpecificPlanIdsLimitation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) []string { return v.SpecificPlanIdsLimitation }).(pulumi.StringArrayOutput)
}

func (o LookupPrivateStoreOfferResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) UniqueOfferId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) string { return v.UniqueOfferId }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreOfferResultOutput) UpdateSuppressedDueIdempotence() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPrivateStoreOfferResult) *bool { return v.UpdateSuppressedDueIdempotence }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateStoreOfferResultOutput{})
}
