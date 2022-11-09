


package v20201101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkServicePrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateLinkServicePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServicePrivateEndpointConnectionResult, error) {
	var rv LookupPrivateLinkServicePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:network/v20201101:getPrivateLinkServicePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateLinkServicePrivateEndpointConnectionArgs struct {
	Expand            *string `pulumi:"expand"`
	PeConnectionName  string  `pulumi:"peConnectionName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupPrivateLinkServicePrivateEndpointConnectionResult struct {
	Etag                              string                                     `pulumi:"etag"`
	Id                                *string                                    `pulumi:"id"`
	LinkIdentifier                    string                                     `pulumi:"linkIdentifier"`
	Name                              *string                                    `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponse                    `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	Type                              string                                     `pulumi:"type"`
}


func (val *LookupPrivateLinkServicePrivateEndpointConnectionResult) Defaults() *LookupPrivateLinkServicePrivateEndpointConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateEndpoint = *tmp.PrivateEndpoint.Defaults()

	return &tmp
}

func LookupPrivateLinkServicePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupPrivateLinkServicePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkServicePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkServicePrivateEndpointConnectionResult, error) {
			args := v.(LookupPrivateLinkServicePrivateEndpointConnectionArgs)
			r, err := LookupPrivateLinkServicePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupPrivateLinkServicePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkServicePrivateEndpointConnectionResultOutput)
}

type LookupPrivateLinkServicePrivateEndpointConnectionOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	PeConnectionName  pulumi.StringInput    `pulumi:"peConnectionName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput    `pulumi:"serviceName"`
}

func (LookupPrivateLinkServicePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServicePrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupPrivateLinkServicePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServicePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) ToLookupPrivateLinkServicePrivateEndpointConnectionResultOutput() LookupPrivateLinkServicePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) ToLookupPrivateLinkServicePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupPrivateLinkServicePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) LinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string { return v.LinkIdentifier }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponseOutput)
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkServicePrivateEndpointConnectionResultOutput{})
}
