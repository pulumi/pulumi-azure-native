


package signalrservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSignalRPrivateEndpointConnection(ctx *pulumi.Context, args *LookupSignalRPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupSignalRPrivateEndpointConnectionResult, error) {
	var rv LookupSignalRPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:signalrservice:getSignalRPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRPrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupSignalRPrivateEndpointConnectionResult struct {
	Id                                string                                     `pulumi:"id"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	Type                              string                                     `pulumi:"type"`
}

func LookupSignalRPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupSignalRPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRPrivateEndpointConnectionResult, error) {
			args := v.(LookupSignalRPrivateEndpointConnectionArgs)
			r, err := LookupSignalRPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupSignalRPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRPrivateEndpointConnectionResultOutput)
}

type LookupSignalRPrivateEndpointConnectionOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalRPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupSignalRPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) ToLookupSignalRPrivateEndpointConnectionResultOutput() LookupSignalRPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) ToLookupSignalRPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupSignalRPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSignalRPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRPrivateEndpointConnectionResultOutput{})
}
