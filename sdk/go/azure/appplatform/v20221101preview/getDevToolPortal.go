


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDevToolPortal(ctx *pulumi.Context, args *LookupDevToolPortalArgs, opts ...pulumi.InvokeOption) (*LookupDevToolPortalResult, error) {
	var rv LookupDevToolPortalResult
	err := ctx.Invoke("azure-native:appplatform/v20221101preview:getDevToolPortal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDevToolPortalArgs struct {
	DevToolPortalName string `pulumi:"devToolPortalName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupDevToolPortalResult struct {
	Id         string                          `pulumi:"id"`
	Name       string                          `pulumi:"name"`
	Properties DevToolPortalPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse              `pulumi:"systemData"`
	Type       string                          `pulumi:"type"`
}


func (val *LookupDevToolPortalResult) Defaults() *LookupDevToolPortalResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupDevToolPortalOutput(ctx *pulumi.Context, args LookupDevToolPortalOutputArgs, opts ...pulumi.InvokeOption) LookupDevToolPortalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDevToolPortalResult, error) {
			args := v.(LookupDevToolPortalArgs)
			r, err := LookupDevToolPortal(ctx, &args, opts...)
			var s LookupDevToolPortalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDevToolPortalResultOutput)
}

type LookupDevToolPortalOutputArgs struct {
	DevToolPortalName pulumi.StringInput `pulumi:"devToolPortalName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupDevToolPortalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevToolPortalArgs)(nil)).Elem()
}


type LookupDevToolPortalResultOutput struct{ *pulumi.OutputState }

func (LookupDevToolPortalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevToolPortalResult)(nil)).Elem()
}

func (o LookupDevToolPortalResultOutput) ToLookupDevToolPortalResultOutput() LookupDevToolPortalResultOutput {
	return o
}

func (o LookupDevToolPortalResultOutput) ToLookupDevToolPortalResultOutputWithContext(ctx context.Context) LookupDevToolPortalResultOutput {
	return o
}

func (o LookupDevToolPortalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDevToolPortalResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDevToolPortalResultOutput) Properties() DevToolPortalPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) DevToolPortalPropertiesResponse { return v.Properties }).(DevToolPortalPropertiesResponseOutput)
}

func (o LookupDevToolPortalResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDevToolPortalResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevToolPortalResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDevToolPortalResultOutput{})
}
