


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectionGateway(ctx *pulumi.Context, args *LookupConnectionGatewayArgs, opts ...pulumi.InvokeOption) (*LookupConnectionGatewayResult, error) {
	var rv LookupConnectionGatewayResult
	err := ctx.Invoke("azure-native:web:getConnectionGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionGatewayArgs struct {
	ConnectionGatewayName string  `pulumi:"connectionGatewayName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	SubscriptionId        *string `pulumi:"subscriptionId"`
}


type LookupConnectionGatewayResult struct {
	Etag       *string                                       `pulumi:"etag"`
	Id         string                                        `pulumi:"id"`
	Location   *string                                       `pulumi:"location"`
	Name       string                                        `pulumi:"name"`
	Properties ConnectionGatewayDefinitionResponseProperties `pulumi:"properties"`
	Tags       map[string]string                             `pulumi:"tags"`
	Type       string                                        `pulumi:"type"`
}

func LookupConnectionGatewayOutput(ctx *pulumi.Context, args LookupConnectionGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionGatewayResult, error) {
			args := v.(LookupConnectionGatewayArgs)
			r, err := LookupConnectionGateway(ctx, &args, opts...)
			var s LookupConnectionGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionGatewayResultOutput)
}

type LookupConnectionGatewayOutputArgs struct {
	ConnectionGatewayName pulumi.StringInput    `pulumi:"connectionGatewayName"`
	ResourceGroupName     pulumi.StringInput    `pulumi:"resourceGroupName"`
	SubscriptionId        pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupConnectionGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionGatewayArgs)(nil)).Elem()
}


type LookupConnectionGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionGatewayResult)(nil)).Elem()
}

func (o LookupConnectionGatewayResultOutput) ToLookupConnectionGatewayResultOutput() LookupConnectionGatewayResultOutput {
	return o
}

func (o LookupConnectionGatewayResultOutput) ToLookupConnectionGatewayResultOutputWithContext(ctx context.Context) LookupConnectionGatewayResultOutput {
	return o
}

func (o LookupConnectionGatewayResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectionGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectionGatewayResultOutput) Properties() ConnectionGatewayDefinitionResponsePropertiesOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) ConnectionGatewayDefinitionResponseProperties {
		return v.Properties
	}).(ConnectionGatewayDefinitionResponsePropertiesOutput)
}

func (o LookupConnectionGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConnectionGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionGatewayResultOutput{})
}
