


package mixedreality

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSpatialAnchorsAccount(ctx *pulumi.Context, args *LookupSpatialAnchorsAccountArgs, opts ...pulumi.InvokeOption) (*LookupSpatialAnchorsAccountResult, error) {
	var rv LookupSpatialAnchorsAccountResult
	err := ctx.Invoke("azure-native:mixedreality:getSpatialAnchorsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSpatialAnchorsAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSpatialAnchorsAccountResult struct {
	AccountDomain      string             `pulumi:"accountDomain"`
	AccountId          string             `pulumi:"accountId"`
	Id                 string             `pulumi:"id"`
	Identity           *IdentityResponse  `pulumi:"identity"`
	Kind               *SkuResponse       `pulumi:"kind"`
	Location           string             `pulumi:"location"`
	Name               string             `pulumi:"name"`
	Plan               *IdentityResponse  `pulumi:"plan"`
	Sku                *SkuResponse       `pulumi:"sku"`
	StorageAccountName *string            `pulumi:"storageAccountName"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Tags               map[string]string  `pulumi:"tags"`
	Type               string             `pulumi:"type"`
}

func LookupSpatialAnchorsAccountOutput(ctx *pulumi.Context, args LookupSpatialAnchorsAccountOutputArgs, opts ...pulumi.InvokeOption) LookupSpatialAnchorsAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSpatialAnchorsAccountResult, error) {
			args := v.(LookupSpatialAnchorsAccountArgs)
			r, err := LookupSpatialAnchorsAccount(ctx, &args, opts...)
			var s LookupSpatialAnchorsAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSpatialAnchorsAccountResultOutput)
}

type LookupSpatialAnchorsAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSpatialAnchorsAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpatialAnchorsAccountArgs)(nil)).Elem()
}


type LookupSpatialAnchorsAccountResultOutput struct{ *pulumi.OutputState }

func (LookupSpatialAnchorsAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpatialAnchorsAccountResult)(nil)).Elem()
}

func (o LookupSpatialAnchorsAccountResultOutput) ToLookupSpatialAnchorsAccountResultOutput() LookupSpatialAnchorsAccountResultOutput {
	return o
}

func (o LookupSpatialAnchorsAccountResultOutput) ToLookupSpatialAnchorsAccountResultOutputWithContext(ctx context.Context) LookupSpatialAnchorsAccountResultOutput {
	return o
}

func (o LookupSpatialAnchorsAccountResultOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.AccountDomain }).(pulumi.StringOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Kind() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) *SkuResponse { return v.Kind }).(SkuResponsePtrOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Plan() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) *IdentityResponse { return v.Plan }).(IdentityResponsePtrOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSpatialAnchorsAccountResultOutput{})
}
