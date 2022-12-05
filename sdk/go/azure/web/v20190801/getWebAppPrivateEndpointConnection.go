


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPrivateEndpointConnection(ctx *pulumi.Context, args *LookupWebAppPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPrivateEndpointConnectionResult, error) {
	var rv LookupWebAppPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:web/v20190801:getWebAppPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPrivateEndpointConnectionArgs struct {
	Name                          string `pulumi:"name"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupWebAppPrivateEndpointConnectionResult struct {
	Id                                string                              `pulumi:"id"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}

func LookupWebAppPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupWebAppPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPrivateEndpointConnectionResult, error) {
			args := v.(LookupWebAppPrivateEndpointConnectionArgs)
			r, err := LookupWebAppPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupWebAppPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPrivateEndpointConnectionResultOutput)
}

type LookupWebAppPrivateEndpointConnectionOutputArgs struct {
	Name                          pulumi.StringInput `pulumi:"name"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupWebAppPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) ToLookupWebAppPrivateEndpointConnectionResultOutput() LookupWebAppPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) ToLookupWebAppPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupWebAppPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionResult) *ArmIdWrapperResponse { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionResult) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebAppPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPrivateEndpointConnectionResultOutput{})
}
