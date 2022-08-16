


package healthcareapis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspacePrivateEndpointConnection(ctx *pulumi.Context, args *LookupWorkspacePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspacePrivateEndpointConnectionResult, error) {
	var rv LookupWorkspacePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:healthcareapis:getWorkspacePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspacePrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	WorkspaceName                 string `pulumi:"workspaceName"`
}


type LookupWorkspacePrivateEndpointConnectionResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupWorkspacePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupWorkspacePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspacePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspacePrivateEndpointConnectionResult, error) {
			args := v.(LookupWorkspacePrivateEndpointConnectionArgs)
			r, err := LookupWorkspacePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupWorkspacePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspacePrivateEndpointConnectionResultOutput)
}

type LookupWorkspacePrivateEndpointConnectionOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName                 pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspacePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspacePrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupWorkspacePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspacePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspacePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) ToLookupWorkspacePrivateEndpointConnectionResultOutput() LookupWorkspacePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) ToLookupWorkspacePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupWorkspacePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWorkspacePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspacePrivateEndpointConnectionResultOutput{})
}
