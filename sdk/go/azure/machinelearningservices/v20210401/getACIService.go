


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupACIService(ctx *pulumi.Context, args *LookupACIServiceArgs, opts ...pulumi.InvokeOption) (*LookupACIServiceResult, error) {
	var rv LookupACIServiceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210401:getACIService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupACIServiceArgs struct {
	Expand            *bool  `pulumi:"expand"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupACIServiceResult struct {
	Id         string             `pulumi:"id"`
	Identity   *IdentityResponse  `pulumi:"identity"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	Properties interface{}        `pulumi:"properties"`
	Sku        *SkuResponse       `pulumi:"sku"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	Type       string             `pulumi:"type"`
}

func LookupACIServiceOutput(ctx *pulumi.Context, args LookupACIServiceOutputArgs, opts ...pulumi.InvokeOption) LookupACIServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupACIServiceResult, error) {
			args := v.(LookupACIServiceArgs)
			r, err := LookupACIService(ctx, &args, opts...)
			var s LookupACIServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupACIServiceResultOutput)
}

type LookupACIServiceOutputArgs struct {
	Expand            pulumi.BoolPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput  `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput  `pulumi:"serviceName"`
	WorkspaceName     pulumi.StringInput  `pulumi:"workspaceName"`
}

func (LookupACIServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupACIServiceArgs)(nil)).Elem()
}


type LookupACIServiceResultOutput struct{ *pulumi.OutputState }

func (LookupACIServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupACIServiceResult)(nil)).Elem()
}

func (o LookupACIServiceResultOutput) ToLookupACIServiceResultOutput() LookupACIServiceResultOutput {
	return o
}

func (o LookupACIServiceResultOutput) ToLookupACIServiceResultOutputWithContext(ctx context.Context) LookupACIServiceResultOutput {
	return o
}

func (o LookupACIServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupACIServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupACIServiceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupACIServiceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupACIServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupACIServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupACIServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupACIServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupACIServiceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupACIServiceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupACIServiceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupACIServiceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupACIServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupACIServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupACIServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupACIServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupACIServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupACIServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupACIServiceResultOutput{})
}
