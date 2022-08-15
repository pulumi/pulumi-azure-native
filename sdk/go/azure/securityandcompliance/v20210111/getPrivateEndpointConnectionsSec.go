


package v20210111

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsSec(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsSecArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsSecResult, error) {
	var rv LookupPrivateEndpointConnectionsSecResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210111:getPrivateEndpointConnectionsSec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsSecArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsSecResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionsSecOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsSecOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsSecResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsSecResult, error) {
			args := v.(LookupPrivateEndpointConnectionsSecArgs)
			r, err := LookupPrivateEndpointConnectionsSec(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsSecResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsSecResultOutput)
}

type LookupPrivateEndpointConnectionsSecOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsSecOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsSecArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionsSecResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsSecResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsSecResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) ToLookupPrivateEndpointConnectionsSecResultOutput() LookupPrivateEndpointConnectionsSecResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) ToLookupPrivateEndpointConnectionsSecResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsSecResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsSecResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsSecResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsSecResult) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsSecResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsSecResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsSecResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionsSecResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsSecResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsSecResultOutput{})
}
