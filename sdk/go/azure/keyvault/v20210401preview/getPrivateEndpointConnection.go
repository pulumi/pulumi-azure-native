


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:keyvault/v20210401preview:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	VaultName                     string `pulumi:"vaultName"`
}


type LookupPrivateEndpointConnectionResult struct {
	Etag                              *string                                    `pulumi:"etag"`
	Id                                string                                     `pulumi:"id"`
	Location                          string                                     `pulumi:"location"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	Tags                              map[string]string                          `pulumi:"tags"`
	Type                              string                                     `pulumi:"type"`
}

func LookupPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionResult, error) {
			args := v.(LookupPrivateEndpointConnectionArgs)
			r, err := LookupPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionResultOutput)
}

type LookupPrivateEndpointConnectionOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName                     pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionResultOutput() LookupPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionResultOutput{})
}
