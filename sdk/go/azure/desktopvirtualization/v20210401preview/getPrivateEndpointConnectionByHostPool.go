


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionByHostPool(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionByHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionByHostPoolResult, error) {
	var rv LookupPrivateEndpointConnectionByHostPoolResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20210401preview:getPrivateEndpointConnectionByHostPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionByHostPoolArgs struct {
	HostPoolName                  string `pulumi:"hostPoolName"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointConnectionByHostPoolResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionByHostPoolOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionByHostPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionByHostPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionByHostPoolResult, error) {
			args := v.(LookupPrivateEndpointConnectionByHostPoolArgs)
			r, err := LookupPrivateEndpointConnectionByHostPool(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionByHostPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionByHostPoolResultOutput)
}

type LookupPrivateEndpointConnectionByHostPoolOutputArgs struct {
	HostPoolName                  pulumi.StringInput `pulumi:"hostPoolName"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointConnectionByHostPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByHostPoolArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionByHostPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionByHostPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByHostPoolResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) ToLookupPrivateEndpointConnectionByHostPoolResultOutput() LookupPrivateEndpointConnectionByHostPoolResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) ToLookupPrivateEndpointConnectionByHostPoolResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionByHostPoolResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionByHostPoolResultOutput{})
}
