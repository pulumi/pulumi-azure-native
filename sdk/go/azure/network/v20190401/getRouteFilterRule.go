


package v20190401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouteFilterRule(ctx *pulumi.Context, args *LookupRouteFilterRuleArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterRuleResult, error) {
	var rv LookupRouteFilterRuleResult
	err := ctx.Invoke("azure-native:network/v20190401:getRouteFilterRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteFilterRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteFilterName   string `pulumi:"routeFilterName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupRouteFilterRuleResult struct {
	Access              string   `pulumi:"access"`
	Communities         []string `pulumi:"communities"`
	Etag                string   `pulumi:"etag"`
	Id                  *string  `pulumi:"id"`
	Location            *string  `pulumi:"location"`
	Name                *string  `pulumi:"name"`
	ProvisioningState   string   `pulumi:"provisioningState"`
	RouteFilterRuleType string   `pulumi:"routeFilterRuleType"`
}

func LookupRouteFilterRuleOutput(ctx *pulumi.Context, args LookupRouteFilterRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRouteFilterRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteFilterRuleResult, error) {
			args := v.(LookupRouteFilterRuleArgs)
			r, err := LookupRouteFilterRule(ctx, &args, opts...)
			var s LookupRouteFilterRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteFilterRuleResultOutput)
}

type LookupRouteFilterRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RouteFilterName   pulumi.StringInput `pulumi:"routeFilterName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupRouteFilterRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterRuleArgs)(nil)).Elem()
}


type LookupRouteFilterRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRouteFilterRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterRuleResult)(nil)).Elem()
}

func (o LookupRouteFilterRuleResultOutput) ToLookupRouteFilterRuleResultOutput() LookupRouteFilterRuleResultOutput {
	return o
}

func (o LookupRouteFilterRuleResultOutput) ToLookupRouteFilterRuleResultOutputWithContext(ctx context.Context) LookupRouteFilterRuleResultOutput {
	return o
}

func (o LookupRouteFilterRuleResultOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) string { return v.Access }).(pulumi.StringOutput)
}

func (o LookupRouteFilterRuleResultOutput) Communities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) []string { return v.Communities }).(pulumi.StringArrayOutput)
}

func (o LookupRouteFilterRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupRouteFilterRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRouteFilterRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupRouteFilterRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRouteFilterRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRouteFilterRuleResultOutput) RouteFilterRuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) string { return v.RouteFilterRuleType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteFilterRuleResultOutput{})
}
