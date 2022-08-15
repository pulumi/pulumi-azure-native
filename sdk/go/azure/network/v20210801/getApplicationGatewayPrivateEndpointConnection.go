


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGatewayPrivateEndpointConnection(ctx *pulumi.Context, args *LookupApplicationGatewayPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGatewayPrivateEndpointConnectionResult, error) {
	var rv LookupApplicationGatewayPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:network/v20210801:getApplicationGatewayPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApplicationGatewayPrivateEndpointConnectionArgs struct {
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	ConnectionName         string `pulumi:"connectionName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupApplicationGatewayPrivateEndpointConnectionResult struct {
	Etag                              string                                     `pulumi:"etag"`
	Id                                *string                                    `pulumi:"id"`
	LinkIdentifier                    string                                     `pulumi:"linkIdentifier"`
	Name                              *string                                    `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponse                    `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	Type                              string                                     `pulumi:"type"`
}


func (val *LookupApplicationGatewayPrivateEndpointConnectionResult) Defaults() *LookupApplicationGatewayPrivateEndpointConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateEndpoint = *tmp.PrivateEndpoint.Defaults()

	return &tmp
}

func LookupApplicationGatewayPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupApplicationGatewayPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationGatewayPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationGatewayPrivateEndpointConnectionResult, error) {
			args := v.(LookupApplicationGatewayPrivateEndpointConnectionArgs)
			r, err := LookupApplicationGatewayPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupApplicationGatewayPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationGatewayPrivateEndpointConnectionResultOutput)
}

type LookupApplicationGatewayPrivateEndpointConnectionOutputArgs struct {
	ApplicationGatewayName pulumi.StringInput `pulumi:"applicationGatewayName"`
	ConnectionName         pulumi.StringInput `pulumi:"connectionName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationGatewayPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGatewayPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupApplicationGatewayPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationGatewayPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGatewayPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) ToLookupApplicationGatewayPrivateEndpointConnectionResultOutput() LookupApplicationGatewayPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) ToLookupApplicationGatewayPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupApplicationGatewayPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) LinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) string { return v.LinkIdentifier }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponseOutput)
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationGatewayPrivateEndpointConnectionResultOutput{})
}
