


package v20160201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCognitiveServicesAccount(ctx *pulumi.Context, args *LookupCognitiveServicesAccountArgs, opts ...pulumi.InvokeOption) (*LookupCognitiveServicesAccountResult, error) {
	var rv LookupCognitiveServicesAccountResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20160201preview:getCognitiveServicesAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCognitiveServicesAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCognitiveServicesAccountResult struct {
	Endpoint          *string           `pulumi:"endpoint"`
	Etag              *string           `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Sku               *SkuResponse      `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}

func LookupCognitiveServicesAccountOutput(ctx *pulumi.Context, args LookupCognitiveServicesAccountOutputArgs, opts ...pulumi.InvokeOption) LookupCognitiveServicesAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCognitiveServicesAccountResult, error) {
			args := v.(LookupCognitiveServicesAccountArgs)
			r, err := LookupCognitiveServicesAccount(ctx, &args, opts...)
			var s LookupCognitiveServicesAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCognitiveServicesAccountResultOutput)
}

type LookupCognitiveServicesAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCognitiveServicesAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCognitiveServicesAccountArgs)(nil)).Elem()
}


type LookupCognitiveServicesAccountResultOutput struct{ *pulumi.OutputState }

func (LookupCognitiveServicesAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCognitiveServicesAccountResult)(nil)).Elem()
}

func (o LookupCognitiveServicesAccountResultOutput) ToLookupCognitiveServicesAccountResultOutput() LookupCognitiveServicesAccountResultOutput {
	return o
}

func (o LookupCognitiveServicesAccountResultOutput) ToLookupCognitiveServicesAccountResultOutputWithContext(ctx context.Context) LookupCognitiveServicesAccountResultOutput {
	return o
}

func (o LookupCognitiveServicesAccountResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCognitiveServicesAccountResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCognitiveServicesAccountResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCognitiveServicesAccountResultOutput{})
}
