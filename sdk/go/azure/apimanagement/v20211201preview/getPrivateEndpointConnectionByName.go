


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionByName(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionByNameArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionByNameResult, error) {
	var rv LookupPrivateEndpointConnectionByNameResult
	err := ctx.Invoke("azure-native:apimanagement/v20211201preview:getPrivateEndpointConnectionByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionByNameArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ServiceName                   string `pulumi:"serviceName"`
}


type LookupPrivateEndpointConnectionByNameResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionByNameOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionByNameOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionByNameResult, error) {
			args := v.(LookupPrivateEndpointConnectionByNameArgs)
			r, err := LookupPrivateEndpointConnectionByName(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionByNameResultOutput)
}

type LookupPrivateEndpointConnectionByNameOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName                   pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPrivateEndpointConnectionByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByNameArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionByNameResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByNameResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) ToLookupPrivateEndpointConnectionByNameResultOutput() LookupPrivateEndpointConnectionByNameResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) ToLookupPrivateEndpointConnectionByNameResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionByNameResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionByNameResultOutput{})
}
