


package v20170418

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20170418:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	Etag       string                                     `pulumi:"etag"`
	Id         string                                     `pulumi:"id"`
	Identity   *IdentityResponse                          `pulumi:"identity"`
	Kind       *string                                    `pulumi:"kind"`
	Location   *string                                    `pulumi:"location"`
	Name       string                                     `pulumi:"name"`
	Properties CognitiveServicesAccountPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                               `pulumi:"sku"`
	Tags       map[string]string                          `pulumi:"tags"`
	Type       string                                     `pulumi:"type"`
}


func (val *LookupAccountResult) Defaults() *LookupAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}


type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupAccountResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Properties() CognitiveServicesAccountPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAccountResult) CognitiveServicesAccountPropertiesResponse { return v.Properties }).(CognitiveServicesAccountPropertiesResponseOutput)
}

func (o LookupAccountResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
