


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteGateway(ctx *pulumi.Context, args *LookupExpressRouteGatewayArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteGatewayResult, error) {
	var rv LookupExpressRouteGatewayResult
	err := ctx.Invoke("azure-native:network/v20190801:getExpressRouteGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteGatewayArgs struct {
	ExpressRouteGatewayName string `pulumi:"expressRouteGatewayName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteGatewayResult struct {
	AutoScaleConfiguration  *ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration `pulumi:"autoScaleConfiguration"`
	Etag                    string                                                       `pulumi:"etag"`
	ExpressRouteConnections []ExpressRouteConnectionResponse                             `pulumi:"expressRouteConnections"`
	Id                      *string                                                      `pulumi:"id"`
	Location                *string                                                      `pulumi:"location"`
	Name                    string                                                       `pulumi:"name"`
	ProvisioningState       string                                                       `pulumi:"provisioningState"`
	Tags                    map[string]string                                            `pulumi:"tags"`
	Type                    string                                                       `pulumi:"type"`
	VirtualHub              VirtualHubIdResponse                                         `pulumi:"virtualHub"`
}

func LookupExpressRouteGatewayOutput(ctx *pulumi.Context, args LookupExpressRouteGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteGatewayResult, error) {
			args := v.(LookupExpressRouteGatewayArgs)
			r, err := LookupExpressRouteGateway(ctx, &args, opts...)
			var s LookupExpressRouteGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteGatewayResultOutput)
}

type LookupExpressRouteGatewayOutputArgs struct {
	ExpressRouteGatewayName pulumi.StringInput `pulumi:"expressRouteGatewayName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteGatewayArgs)(nil)).Elem()
}


type LookupExpressRouteGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteGatewayResult)(nil)).Elem()
}

func (o LookupExpressRouteGatewayResultOutput) ToLookupExpressRouteGatewayResultOutput() LookupExpressRouteGatewayResultOutput {
	return o
}

func (o LookupExpressRouteGatewayResultOutput) ToLookupExpressRouteGatewayResultOutputWithContext(ctx context.Context) LookupExpressRouteGatewayResultOutput {
	return o
}

func (o LookupExpressRouteGatewayResultOutput) AutoScaleConfiguration() ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) *ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration {
		return v.AutoScaleConfiguration
	}).(ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput)
}

func (o LookupExpressRouteGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupExpressRouteGatewayResultOutput) ExpressRouteConnections() ExpressRouteConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) []ExpressRouteConnectionResponse {
		return v.ExpressRouteConnections
	}).(ExpressRouteConnectionResponseArrayOutput)
}

func (o LookupExpressRouteGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExpressRouteGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExpressRouteGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupExpressRouteGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupExpressRouteGatewayResultOutput) VirtualHub() VirtualHubIdResponseOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) VirtualHubIdResponse { return v.VirtualHub }).(VirtualHubIdResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteGatewayResultOutput{})
}
