


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkScopedResource(ctx *pulumi.Context, args *LookupPrivateLinkScopedResourceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkScopedResourceResult, error) {
	var rv LookupPrivateLinkScopedResourceResult
	err := ctx.Invoke("azure-native:insights/v20210701preview:getPrivateLinkScopedResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkScopedResourceArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScopeName         string `pulumi:"scopeName"`
}


type LookupPrivateLinkScopedResourceResult struct {
	Id                string             `pulumi:"id"`
	LinkedResourceId  *string            `pulumi:"linkedResourceId"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupPrivateLinkScopedResourceOutput(ctx *pulumi.Context, args LookupPrivateLinkScopedResourceOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkScopedResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkScopedResourceResult, error) {
			args := v.(LookupPrivateLinkScopedResourceArgs)
			r, err := LookupPrivateLinkScopedResource(ctx, &args, opts...)
			var s LookupPrivateLinkScopedResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkScopedResourceResultOutput)
}

type LookupPrivateLinkScopedResourceOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScopeName         pulumi.StringInput `pulumi:"scopeName"`
}

func (LookupPrivateLinkScopedResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkScopedResourceArgs)(nil)).Elem()
}


type LookupPrivateLinkScopedResourceResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkScopedResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkScopedResourceResult)(nil)).Elem()
}

func (o LookupPrivateLinkScopedResourceResultOutput) ToLookupPrivateLinkScopedResourceResultOutput() LookupPrivateLinkScopedResourceResultOutput {
	return o
}

func (o LookupPrivateLinkScopedResourceResultOutput) ToLookupPrivateLinkScopedResourceResultOutputWithContext(ctx context.Context) LookupPrivateLinkScopedResourceResultOutput {
	return o
}

func (o LookupPrivateLinkScopedResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopedResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkScopedResourceResultOutput) LinkedResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopedResourceResult) *string { return v.LinkedResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateLinkScopedResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopedResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkScopedResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopedResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkScopedResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopedResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateLinkScopedResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopedResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkScopedResourceResultOutput{})
}
