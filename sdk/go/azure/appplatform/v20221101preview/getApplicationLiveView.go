


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationLiveView(ctx *pulumi.Context, args *LookupApplicationLiveViewArgs, opts ...pulumi.InvokeOption) (*LookupApplicationLiveViewResult, error) {
	var rv LookupApplicationLiveViewResult
	err := ctx.Invoke("azure-native:appplatform/v20221101preview:getApplicationLiveView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationLiveViewArgs struct {
	ApplicationLiveViewName string `pulumi:"applicationLiveViewName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ServiceName             string `pulumi:"serviceName"`
}


type LookupApplicationLiveViewResult struct {
	Id         string                                `pulumi:"id"`
	Name       string                                `pulumi:"name"`
	Properties ApplicationLiveViewPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                    `pulumi:"systemData"`
	Type       string                                `pulumi:"type"`
}

func LookupApplicationLiveViewOutput(ctx *pulumi.Context, args LookupApplicationLiveViewOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationLiveViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationLiveViewResult, error) {
			args := v.(LookupApplicationLiveViewArgs)
			r, err := LookupApplicationLiveView(ctx, &args, opts...)
			var s LookupApplicationLiveViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationLiveViewResultOutput)
}

type LookupApplicationLiveViewOutputArgs struct {
	ApplicationLiveViewName pulumi.StringInput `pulumi:"applicationLiveViewName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName             pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApplicationLiveViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationLiveViewArgs)(nil)).Elem()
}


type LookupApplicationLiveViewResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationLiveViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationLiveViewResult)(nil)).Elem()
}

func (o LookupApplicationLiveViewResultOutput) ToLookupApplicationLiveViewResultOutput() LookupApplicationLiveViewResultOutput {
	return o
}

func (o LookupApplicationLiveViewResultOutput) ToLookupApplicationLiveViewResultOutputWithContext(ctx context.Context) LookupApplicationLiveViewResultOutput {
	return o
}

func (o LookupApplicationLiveViewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationLiveViewResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationLiveViewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationLiveViewResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationLiveViewResultOutput) Properties() ApplicationLiveViewPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApplicationLiveViewResult) ApplicationLiveViewPropertiesResponse { return v.Properties }).(ApplicationLiveViewPropertiesResponseOutput)
}

func (o LookupApplicationLiveViewResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationLiveViewResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApplicationLiveViewResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationLiveViewResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationLiveViewResultOutput{})
}
