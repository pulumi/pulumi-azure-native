


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouteMap(ctx *pulumi.Context, args *LookupRouteMapArgs, opts ...pulumi.InvokeOption) (*LookupRouteMapResult, error) {
	var rv LookupRouteMapResult
	err := ctx.Invoke("azure-native:network/v20220501:getRouteMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteMapArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteMapName      string `pulumi:"routeMapName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupRouteMapResult struct {
	AssociatedInboundConnections  []string               `pulumi:"associatedInboundConnections"`
	AssociatedOutboundConnections []string               `pulumi:"associatedOutboundConnections"`
	Etag                          string                 `pulumi:"etag"`
	Id                            string                 `pulumi:"id"`
	Name                          string                 `pulumi:"name"`
	ProvisioningState             string                 `pulumi:"provisioningState"`
	Rules                         []RouteMapRuleResponse `pulumi:"rules"`
	Type                          string                 `pulumi:"type"`
}

func LookupRouteMapOutput(ctx *pulumi.Context, args LookupRouteMapOutputArgs, opts ...pulumi.InvokeOption) LookupRouteMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteMapResult, error) {
			args := v.(LookupRouteMapArgs)
			r, err := LookupRouteMap(ctx, &args, opts...)
			var s LookupRouteMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteMapResultOutput)
}

type LookupRouteMapOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RouteMapName      pulumi.StringInput `pulumi:"routeMapName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupRouteMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteMapArgs)(nil)).Elem()
}


type LookupRouteMapResultOutput struct{ *pulumi.OutputState }

func (LookupRouteMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteMapResult)(nil)).Elem()
}

func (o LookupRouteMapResultOutput) ToLookupRouteMapResultOutput() LookupRouteMapResultOutput {
	return o
}

func (o LookupRouteMapResultOutput) ToLookupRouteMapResultOutputWithContext(ctx context.Context) LookupRouteMapResultOutput {
	return o
}

func (o LookupRouteMapResultOutput) AssociatedInboundConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteMapResult) []string { return v.AssociatedInboundConnections }).(pulumi.StringArrayOutput)
}

func (o LookupRouteMapResultOutput) AssociatedOutboundConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteMapResult) []string { return v.AssociatedOutboundConnections }).(pulumi.StringArrayOutput)
}

func (o LookupRouteMapResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupRouteMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouteMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouteMapResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRouteMapResultOutput) Rules() RouteMapRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteMapResult) []RouteMapRuleResponse { return v.Rules }).(RouteMapRuleResponseArrayOutput)
}

func (o LookupRouteMapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteMapResultOutput{})
}
