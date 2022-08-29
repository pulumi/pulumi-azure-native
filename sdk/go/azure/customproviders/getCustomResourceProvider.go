


package customproviders

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomResourceProvider(ctx *pulumi.Context, args *LookupCustomResourceProviderArgs, opts ...pulumi.InvokeOption) (*LookupCustomResourceProviderResult, error) {
	var rv LookupCustomResourceProviderResult
	err := ctx.Invoke("azure-native:customproviders:getCustomResourceProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomResourceProviderArgs struct {
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	ResourceProviderName string `pulumi:"resourceProviderName"`
}


type LookupCustomResourceProviderResult struct {
	Actions           []CustomRPActionRouteDefinitionResponse       `pulumi:"actions"`
	Id                string                                        `pulumi:"id"`
	Location          string                                        `pulumi:"location"`
	Name              string                                        `pulumi:"name"`
	ProvisioningState string                                        `pulumi:"provisioningState"`
	ResourceTypes     []CustomRPResourceTypeRouteDefinitionResponse `pulumi:"resourceTypes"`
	Tags              map[string]string                             `pulumi:"tags"`
	Type              string                                        `pulumi:"type"`
	Validations       []CustomRPValidationsResponse                 `pulumi:"validations"`
}

func LookupCustomResourceProviderOutput(ctx *pulumi.Context, args LookupCustomResourceProviderOutputArgs, opts ...pulumi.InvokeOption) LookupCustomResourceProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomResourceProviderResult, error) {
			args := v.(LookupCustomResourceProviderArgs)
			r, err := LookupCustomResourceProvider(ctx, &args, opts...)
			var s LookupCustomResourceProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomResourceProviderResultOutput)
}

type LookupCustomResourceProviderOutputArgs struct {
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceProviderName pulumi.StringInput `pulumi:"resourceProviderName"`
}

func (LookupCustomResourceProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomResourceProviderArgs)(nil)).Elem()
}


type LookupCustomResourceProviderResultOutput struct{ *pulumi.OutputState }

func (LookupCustomResourceProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomResourceProviderResult)(nil)).Elem()
}

func (o LookupCustomResourceProviderResultOutput) ToLookupCustomResourceProviderResultOutput() LookupCustomResourceProviderResultOutput {
	return o
}

func (o LookupCustomResourceProviderResultOutput) ToLookupCustomResourceProviderResultOutputWithContext(ctx context.Context) LookupCustomResourceProviderResultOutput {
	return o
}

func (o LookupCustomResourceProviderResultOutput) Actions() CustomRPActionRouteDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) []CustomRPActionRouteDefinitionResponse { return v.Actions }).(CustomRPActionRouteDefinitionResponseArrayOutput)
}

func (o LookupCustomResourceProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomResourceProviderResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCustomResourceProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomResourceProviderResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCustomResourceProviderResultOutput) ResourceTypes() CustomRPResourceTypeRouteDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) []CustomRPResourceTypeRouteDefinitionResponse {
		return v.ResourceTypes
	}).(CustomRPResourceTypeRouteDefinitionResponseArrayOutput)
}

func (o LookupCustomResourceProviderResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCustomResourceProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCustomResourceProviderResultOutput) Validations() CustomRPValidationsResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomResourceProviderResult) []CustomRPValidationsResponse { return v.Validations }).(CustomRPValidationsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomResourceProviderResultOutput{})
}
