


package v20210325preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsForMIPPolicySync(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsForMIPPolicySyncArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsForMIPPolicySyncResult, error) {
	var rv LookupPrivateEndpointConnectionsForMIPPolicySyncResult
	err := ctx.Invoke("azure-native:m365securityandcompliance/v20210325preview:getPrivateEndpointConnectionsForMIPPolicySync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsForMIPPolicySyncArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsForMIPPolicySyncResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionsForMIPPolicySyncOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsForMIPPolicySyncOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsForMIPPolicySyncResult, error) {
			args := v.(LookupPrivateEndpointConnectionsForMIPPolicySyncArgs)
			r, err := LookupPrivateEndpointConnectionsForMIPPolicySync(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsForMIPPolicySyncResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput)
}

type LookupPrivateEndpointConnectionsForMIPPolicySyncOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsForMIPPolicySyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForMIPPolicySyncArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForMIPPolicySyncResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ToLookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput() LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ToLookupPrivateEndpointConnectionsForMIPPolicySyncResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForMIPPolicySyncResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsForMIPPolicySyncResultOutput{})
}
