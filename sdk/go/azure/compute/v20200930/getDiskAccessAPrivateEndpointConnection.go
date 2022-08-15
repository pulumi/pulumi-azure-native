


package v20200930

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context, args *LookupDiskAccessAPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupDiskAccessAPrivateEndpointConnectionResult, error) {
	var rv LookupDiskAccessAPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:compute/v20200930:getDiskAccessAPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskAccessAPrivateEndpointConnectionArgs struct {
	DiskAccessName                string `pulumi:"diskAccessName"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupDiskAccessAPrivateEndpointConnectionResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}

func LookupDiskAccessAPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupDiskAccessAPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupDiskAccessAPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskAccessAPrivateEndpointConnectionResult, error) {
			args := v.(LookupDiskAccessAPrivateEndpointConnectionArgs)
			r, err := LookupDiskAccessAPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupDiskAccessAPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskAccessAPrivateEndpointConnectionResultOutput)
}

type LookupDiskAccessAPrivateEndpointConnectionOutputArgs struct {
	DiskAccessName                pulumi.StringInput `pulumi:"diskAccessName"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskAccessAPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskAccessAPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupDiskAccessAPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupDiskAccessAPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskAccessAPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) ToLookupDiskAccessAPrivateEndpointConnectionResultOutput() LookupDiskAccessAPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) ToLookupDiskAccessAPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupDiskAccessAPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskAccessAPrivateEndpointConnectionResultOutput{})
}
