


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateStoreCollection(ctx *pulumi.Context, args *LookupPrivateStoreCollectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateStoreCollectionResult, error) {
	var rv LookupPrivateStoreCollectionResult
	err := ctx.Invoke("azure-native:marketplace/v20220301:getPrivateStoreCollection", args, &rv, opts...)
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
	AllSubscriptions          *bool              `pulumi:"allSubscriptions"`
	ApproveAllItems           bool               `pulumi:"approveAllItems"`
	ApproveAllItemsModifiedAt string             `pulumi:"approveAllItemsModifiedAt"`
	Claim                     *string            `pulumi:"claim"`
	CollectionId              string             `pulumi:"collectionId"`
	CollectionName            *string            `pulumi:"collectionName"`
	Enabled                   *bool              `pulumi:"enabled"`
	Id                        string             `pulumi:"id"`
	Name                      string             `pulumi:"name"`
	NumberOfOffers            float64            `pulumi:"numberOfOffers"`
	SubscriptionsList         []string           `pulumi:"subscriptionsList"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	Type                      string             `pulumi:"type"`
}

func LookupPrivateStoreCollectionOutput(ctx *pulumi.Context, args LookupPrivateStoreCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateStoreCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateStoreCollectionResult, error) {
			args := v.(LookupPrivateStoreCollectionArgs)
			r, err := LookupPrivateStoreCollection(ctx, &args, opts...)
			var s LookupPrivateStoreCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateStoreCollectionResultOutput)
}

type LookupPrivateStoreCollectionOutputArgs struct {
	CollectionId   pulumi.StringInput `pulumi:"collectionId"`
	PrivateStoreId pulumi.StringInput `pulumi:"privateStoreId"`
}

func (LookupPrivateStoreCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateStoreCollectionArgs)(nil)).Elem()
}


type LookupPrivateStoreCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateStoreCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateStoreCollectionResult)(nil)).Elem()
}

func (o LookupPrivateStoreCollectionResultOutput) ToLookupPrivateStoreCollectionResultOutput() LookupPrivateStoreCollectionResultOutput {
	return o
}

func (o LookupPrivateStoreCollectionResultOutput) ToLookupPrivateStoreCollectionResultOutputWithContext(ctx context.Context) LookupPrivateStoreCollectionResultOutput {
	return o
}

func (o LookupPrivateStoreCollectionResultOutput) AllSubscriptions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) *bool { return v.AllSubscriptions }).(pulumi.BoolPtrOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) ApproveAllItems() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) bool { return v.ApproveAllItems }).(pulumi.BoolOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) ApproveAllItemsModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) string { return v.ApproveAllItemsModifiedAt }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) Claim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) *string { return v.Claim }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) CollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) string { return v.CollectionId }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) CollectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) *string { return v.CollectionName }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) NumberOfOffers() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) float64 { return v.NumberOfOffers }).(pulumi.Float64Output)
}

func (o LookupPrivateStoreCollectionResultOutput) SubscriptionsList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) []string { return v.SubscriptionsList }).(pulumi.StringArrayOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateStoreCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateStoreCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateStoreCollectionResultOutput{})
}
