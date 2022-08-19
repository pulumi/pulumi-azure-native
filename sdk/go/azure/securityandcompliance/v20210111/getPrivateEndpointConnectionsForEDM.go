


package v20210111

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsForEDM(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsForEDMArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsForEDMResult, error) {
	var rv LookupPrivateEndpointConnectionsForEDMResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210111:getPrivateEndpointConnectionsForEDM", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsForEDMArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsForEDMResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionsForEDMOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsForEDMOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsForEDMResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsForEDMResult, error) {
			args := v.(LookupPrivateEndpointConnectionsForEDMArgs)
			r, err := LookupPrivateEndpointConnectionsForEDM(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsForEDMResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsForEDMResultOutput)
}

type LookupPrivateEndpointConnectionsForEDMOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsForEDMOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForEDMArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionsForEDMResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsForEDMResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForEDMResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) ToLookupPrivateEndpointConnectionsForEDMResultOutput() LookupPrivateEndpointConnectionsForEDMResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) ToLookupPrivateEndpointConnectionsForEDMResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsForEDMResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForEDMResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForEDMResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForEDMResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForEDMResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForEDMResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForEDMResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionsForEDMResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForEDMResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsForEDMResultOutput{})
}
