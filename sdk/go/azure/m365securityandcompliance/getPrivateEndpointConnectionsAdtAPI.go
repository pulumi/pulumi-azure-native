


package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsAdtAPI(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsAdtAPIArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsAdtAPIResult, error) {
	var rv LookupPrivateEndpointConnectionsAdtAPIResult
	err := ctx.Invoke("azure-native:m365securityandcompliance:getPrivateEndpointConnectionsAdtAPI", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsAdtAPIArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsAdtAPIResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionsAdtAPIOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsAdtAPIOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsAdtAPIResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsAdtAPIResult, error) {
			args := v.(LookupPrivateEndpointConnectionsAdtAPIArgs)
			r, err := LookupPrivateEndpointConnectionsAdtAPI(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsAdtAPIResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsAdtAPIResultOutput)
}

type LookupPrivateEndpointConnectionsAdtAPIOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsAdtAPIOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsAdtAPIArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionsAdtAPIResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsAdtAPIResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsAdtAPIResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) ToLookupPrivateEndpointConnectionsAdtAPIResultOutput() LookupPrivateEndpointConnectionsAdtAPIResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) ToLookupPrivateEndpointConnectionsAdtAPIResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsAdtAPIResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionsAdtAPIResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsAdtAPIResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsAdtAPIResultOutput{})
}
