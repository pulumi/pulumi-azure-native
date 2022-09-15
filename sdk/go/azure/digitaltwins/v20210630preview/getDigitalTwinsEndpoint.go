


package v20210630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDigitalTwinsEndpoint(ctx *pulumi.Context, args *LookupDigitalTwinsEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDigitalTwinsEndpointResult, error) {
	var rv LookupDigitalTwinsEndpointResult
	err := ctx.Invoke("azure-native:digitaltwins/v20210630preview:getDigitalTwinsEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDigitalTwinsEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupDigitalTwinsEndpointResult struct {
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	Properties interface{}        `pulumi:"properties"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupDigitalTwinsEndpointOutput(ctx *pulumi.Context, args LookupDigitalTwinsEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupDigitalTwinsEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDigitalTwinsEndpointResult, error) {
			args := v.(LookupDigitalTwinsEndpointArgs)
			r, err := LookupDigitalTwinsEndpoint(ctx, &args, opts...)
			var s LookupDigitalTwinsEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDigitalTwinsEndpointResultOutput)
}

type LookupDigitalTwinsEndpointOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupDigitalTwinsEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDigitalTwinsEndpointArgs)(nil)).Elem()
}


type LookupDigitalTwinsEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupDigitalTwinsEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDigitalTwinsEndpointResult)(nil)).Elem()
}

func (o LookupDigitalTwinsEndpointResultOutput) ToLookupDigitalTwinsEndpointResultOutput() LookupDigitalTwinsEndpointResultOutput {
	return o
}

func (o LookupDigitalTwinsEndpointResultOutput) ToLookupDigitalTwinsEndpointResultOutputWithContext(ctx context.Context) LookupDigitalTwinsEndpointResultOutput {
	return o
}

func (o LookupDigitalTwinsEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinsEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinsEndpointResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupDigitalTwinsEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDigitalTwinsEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDigitalTwinsEndpointResultOutput{})
}
