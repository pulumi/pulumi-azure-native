


package mixedreality

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupObjectAnchorsAccount(ctx *pulumi.Context, args *LookupObjectAnchorsAccountArgs, opts ...pulumi.InvokeOption) (*LookupObjectAnchorsAccountResult, error) {
	var rv LookupObjectAnchorsAccountResult
	err := ctx.Invoke("azure-native:mixedreality:getObjectAnchorsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupObjectAnchorsAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupObjectAnchorsAccountResult struct {
	AccountDomain      string                                `pulumi:"accountDomain"`
	AccountId          string                                `pulumi:"accountId"`
	Id                 string                                `pulumi:"id"`
	Identity           *ObjectAnchorsAccountResponseIdentity `pulumi:"identity"`
	Kind               *SkuResponse                          `pulumi:"kind"`
	Location           string                                `pulumi:"location"`
	Name               string                                `pulumi:"name"`
	Plan               *IdentityResponse                     `pulumi:"plan"`
	Sku                *SkuResponse                          `pulumi:"sku"`
	StorageAccountName *string                               `pulumi:"storageAccountName"`
	SystemData         SystemDataResponse                    `pulumi:"systemData"`
	Tags               map[string]string                     `pulumi:"tags"`
	Type               string                                `pulumi:"type"`
}

func LookupObjectAnchorsAccountOutput(ctx *pulumi.Context, args LookupObjectAnchorsAccountOutputArgs, opts ...pulumi.InvokeOption) LookupObjectAnchorsAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObjectAnchorsAccountResult, error) {
			args := v.(LookupObjectAnchorsAccountArgs)
			r, err := LookupObjectAnchorsAccount(ctx, &args, opts...)
			var s LookupObjectAnchorsAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupObjectAnchorsAccountResultOutput)
}

type LookupObjectAnchorsAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupObjectAnchorsAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectAnchorsAccountArgs)(nil)).Elem()
}


type LookupObjectAnchorsAccountResultOutput struct{ *pulumi.OutputState }

func (LookupObjectAnchorsAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectAnchorsAccountResult)(nil)).Elem()
}

func (o LookupObjectAnchorsAccountResultOutput) ToLookupObjectAnchorsAccountResultOutput() LookupObjectAnchorsAccountResultOutput {
	return o
}

func (o LookupObjectAnchorsAccountResultOutput) ToLookupObjectAnchorsAccountResultOutputWithContext(ctx context.Context) LookupObjectAnchorsAccountResultOutput {
	return o
}

func (o LookupObjectAnchorsAccountResultOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.AccountDomain }).(pulumi.StringOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Identity() ObjectAnchorsAccountResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *ObjectAnchorsAccountResponseIdentity { return v.Identity }).(ObjectAnchorsAccountResponseIdentityPtrOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Kind() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *SkuResponse { return v.Kind }).(SkuResponsePtrOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Plan() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *IdentityResponse { return v.Plan }).(IdentityResponsePtrOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupObjectAnchorsAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectAnchorsAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObjectAnchorsAccountResultOutput{})
}
