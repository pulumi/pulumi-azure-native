


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionByWorkspace(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionByWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionByWorkspaceResult, error) {
	var rv LookupPrivateEndpointConnectionByWorkspaceResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20220401preview:getPrivateEndpointConnectionByWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionByWorkspaceArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	WorkspaceName                 string `pulumi:"workspaceName"`
}


type LookupPrivateEndpointConnectionByWorkspaceResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionByWorkspaceOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionByWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionByWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionByWorkspaceResult, error) {
			args := v.(LookupPrivateEndpointConnectionByWorkspaceArgs)
			r, err := LookupPrivateEndpointConnectionByWorkspace(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionByWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionByWorkspaceResultOutput)
}

type LookupPrivateEndpointConnectionByWorkspaceOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName                 pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupPrivateEndpointConnectionByWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByWorkspaceArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionByWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionByWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByWorkspaceResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) ToLookupPrivateEndpointConnectionByWorkspaceResultOutput() LookupPrivateEndpointConnectionByWorkspaceResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) ToLookupPrivateEndpointConnectionByWorkspaceResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionByWorkspaceResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByWorkspaceResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByWorkspaceResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionByWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionByWorkspaceResultOutput{})
}
