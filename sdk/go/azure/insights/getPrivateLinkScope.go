


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkScope(ctx *pulumi.Context, args *LookupPrivateLinkScopeArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkScopeResult, error) {
	var rv LookupPrivateLinkScopeResult
	err := ctx.Invoke("azure-native:insights:getPrivateLinkScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkScopeArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScopeName         string `pulumi:"scopeName"`
}


type LookupPrivateLinkScopeResult struct {
	Id                         string                              `pulumi:"id"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}

func LookupPrivateLinkScopeOutput(ctx *pulumi.Context, args LookupPrivateLinkScopeOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkScopeResult, error) {
			args := v.(LookupPrivateLinkScopeArgs)
			r, err := LookupPrivateLinkScope(ctx, &args, opts...)
			var s LookupPrivateLinkScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkScopeResultOutput)
}

type LookupPrivateLinkScopeOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScopeName         pulumi.StringInput `pulumi:"scopeName"`
}

func (LookupPrivateLinkScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkScopeArgs)(nil)).Elem()
}


type LookupPrivateLinkScopeResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkScopeResult)(nil)).Elem()
}

func (o LookupPrivateLinkScopeResultOutput) ToLookupPrivateLinkScopeResultOutput() LookupPrivateLinkScopeResultOutput {
	return o
}

func (o LookupPrivateLinkScopeResultOutput) ToLookupPrivateLinkScopeResultOutputWithContext(ctx context.Context) LookupPrivateLinkScopeResultOutput {
	return o
}

func (o LookupPrivateLinkScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkScopeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkScopeResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupPrivateLinkScopeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkScopeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateLinkScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkScopeResultOutput{})
}
