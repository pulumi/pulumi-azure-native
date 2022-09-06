


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiPortal(ctx *pulumi.Context, args *LookupApiPortalArgs, opts ...pulumi.InvokeOption) (*LookupApiPortalResult, error) {
	var rv LookupApiPortalResult
	err := ctx.Invoke("azure-native:appplatform/v20220501preview:getApiPortal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiPortalArgs struct {
	ApiPortalName     string `pulumi:"apiPortalName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiPortalResult struct {
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties ApiPortalPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                `pulumi:"sku"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Type       string                      `pulumi:"type"`
}


func (val *LookupApiPortalResult) Defaults() *LookupApiPortalResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupApiPortalOutput(ctx *pulumi.Context, args LookupApiPortalOutputArgs, opts ...pulumi.InvokeOption) LookupApiPortalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiPortalResult, error) {
			args := v.(LookupApiPortalArgs)
			r, err := LookupApiPortal(ctx, &args, opts...)
			var s LookupApiPortalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiPortalResultOutput)
}

type LookupApiPortalOutputArgs struct {
	ApiPortalName     pulumi.StringInput `pulumi:"apiPortalName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiPortalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPortalArgs)(nil)).Elem()
}


type LookupApiPortalResultOutput struct{ *pulumi.OutputState }

func (LookupApiPortalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPortalResult)(nil)).Elem()
}

func (o LookupApiPortalResultOutput) ToLookupApiPortalResultOutput() LookupApiPortalResultOutput {
	return o
}

func (o LookupApiPortalResultOutput) ToLookupApiPortalResultOutputWithContext(ctx context.Context) LookupApiPortalResultOutput {
	return o
}

func (o LookupApiPortalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiPortalResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiPortalResultOutput) Properties() ApiPortalPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApiPortalResult) ApiPortalPropertiesResponse { return v.Properties }).(ApiPortalPropertiesResponseOutput)
}

func (o LookupApiPortalResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApiPortalResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupApiPortalResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApiPortalResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApiPortalResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiPortalResultOutput{})
}
