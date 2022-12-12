


package v20210111

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsComp(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsCompArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsCompResult, error) {
	var rv LookupPrivateEndpointConnectionsCompResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210111:getPrivateEndpointConnectionsComp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsCompArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsCompResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionsCompOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsCompOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsCompResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsCompResult, error) {
			args := v.(LookupPrivateEndpointConnectionsCompArgs)
			r, err := LookupPrivateEndpointConnectionsComp(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsCompResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsCompResultOutput)
}

type LookupPrivateEndpointConnectionsCompOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsCompOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsCompArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionsCompResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsCompResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsCompResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) ToLookupPrivateEndpointConnectionsCompResultOutput() LookupPrivateEndpointConnectionsCompResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) ToLookupPrivateEndpointConnectionsCompResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsCompResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsCompResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsCompResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsCompResult) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsCompResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsCompResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsCompResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionsCompResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsCompResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsCompResultOutput{})
}
