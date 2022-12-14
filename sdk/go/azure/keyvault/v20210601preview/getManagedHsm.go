


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedHsm(ctx *pulumi.Context, args *LookupManagedHsmArgs, opts ...pulumi.InvokeOption) (*LookupManagedHsmResult, error) {
	var rv LookupManagedHsmResult
	err := ctx.Invoke("azure-native:keyvault/v20210601preview:getManagedHsm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedHsmArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedHsmResult struct {
	Id         string                       `pulumi:"id"`
	Location   *string                      `pulumi:"location"`
	Name       string                       `pulumi:"name"`
	Properties ManagedHsmPropertiesResponse `pulumi:"properties"`
	Sku        *ManagedHsmSkuResponse       `pulumi:"sku"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Tags       map[string]string            `pulumi:"tags"`
	Type       string                       `pulumi:"type"`
}


func (val *LookupManagedHsmResult) Defaults() *LookupManagedHsmResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupManagedHsmOutput(ctx *pulumi.Context, args LookupManagedHsmOutputArgs, opts ...pulumi.InvokeOption) LookupManagedHsmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedHsmResult, error) {
			args := v.(LookupManagedHsmArgs)
			r, err := LookupManagedHsm(ctx, &args, opts...)
			var s LookupManagedHsmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedHsmResultOutput)
}

type LookupManagedHsmOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedHsmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedHsmArgs)(nil)).Elem()
}


type LookupManagedHsmResultOutput struct{ *pulumi.OutputState }

func (LookupManagedHsmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedHsmResult)(nil)).Elem()
}

func (o LookupManagedHsmResultOutput) ToLookupManagedHsmResultOutput() LookupManagedHsmResultOutput {
	return o
}

func (o LookupManagedHsmResultOutput) ToLookupManagedHsmResultOutputWithContext(ctx context.Context) LookupManagedHsmResultOutput {
	return o
}

func (o LookupManagedHsmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedHsmResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupManagedHsmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedHsmResultOutput) Properties() ManagedHsmPropertiesResponseOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) ManagedHsmPropertiesResponse { return v.Properties }).(ManagedHsmPropertiesResponseOutput)
}

func (o LookupManagedHsmResultOutput) Sku() ManagedHsmSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) *ManagedHsmSkuResponse { return v.Sku }).(ManagedHsmSkuResponsePtrOutput)
}

func (o LookupManagedHsmResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagedHsmResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedHsmResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedHsmResultOutput{})
}
