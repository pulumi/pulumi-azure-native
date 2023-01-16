


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRouteFilter(ctx *pulumi.Context, args *LookupRouteFilterArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterResult, error) {
	var rv LookupRouteFilterResult
	err := ctx.Invoke("azure-native:network/v20180201:getRouteFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteFilterArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RouteFilterName   string  `pulumi:"routeFilterName"`
}


type LookupRouteFilterResult struct {
	Etag              string                               `pulumi:"etag"`
	Id                *string                              `pulumi:"id"`
	Location          string                               `pulumi:"location"`
	Name              string                               `pulumi:"name"`
	Peerings          []ExpressRouteCircuitPeeringResponse `pulumi:"peerings"`
	ProvisioningState string                               `pulumi:"provisioningState"`
	Rules             []RouteFilterRuleResponse            `pulumi:"rules"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              string                               `pulumi:"type"`
}

func LookupRouteFilterOutput(ctx *pulumi.Context, args LookupRouteFilterOutputArgs, opts ...pulumi.InvokeOption) LookupRouteFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteFilterResult, error) {
			args := v.(LookupRouteFilterArgs)
			r, err := LookupRouteFilter(ctx, &args, opts...)
			var s LookupRouteFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteFilterResultOutput)
}

type LookupRouteFilterOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	RouteFilterName   pulumi.StringInput    `pulumi:"routeFilterName"`
}

func (LookupRouteFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterArgs)(nil)).Elem()
}


type LookupRouteFilterResultOutput struct{ *pulumi.OutputState }

func (LookupRouteFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterResult)(nil)).Elem()
}

func (o LookupRouteFilterResultOutput) ToLookupRouteFilterResultOutput() LookupRouteFilterResultOutput {
	return o
}

func (o LookupRouteFilterResultOutput) ToLookupRouteFilterResultOutputWithContext(ctx context.Context) LookupRouteFilterResultOutput {
	return o
}

func (o LookupRouteFilterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupRouteFilterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRouteFilterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRouteFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouteFilterResultOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) []ExpressRouteCircuitPeeringResponse { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o LookupRouteFilterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRouteFilterResultOutput) Rules() RouteFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) []RouteFilterRuleResponse { return v.Rules }).(RouteFilterRuleResponseArrayOutput)
}

func (o LookupRouteFilterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRouteFilterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteFilterResultOutput{})
}
