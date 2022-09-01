


package v20210115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceEnvironmentPrivateEndpointConnection(ctx *pulumi.Context, args *LookupAppServiceEnvironmentPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceEnvironmentPrivateEndpointConnectionResult, error) {
	var rv LookupAppServiceEnvironmentPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:web/v20210115:getAppServiceEnvironmentPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceEnvironmentPrivateEndpointConnectionArgs struct {
	Name                          string `pulumi:"name"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupAppServiceEnvironmentPrivateEndpointConnectionResult struct {
	Id                                string                              `pulumi:"id"`
	IpAddresses                       []string                            `pulumi:"ipAddresses"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}

func LookupAppServiceEnvironmentPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupAppServiceEnvironmentPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServiceEnvironmentPrivateEndpointConnectionResult, error) {
			args := v.(LookupAppServiceEnvironmentPrivateEndpointConnectionArgs)
			r, err := LookupAppServiceEnvironmentPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupAppServiceEnvironmentPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput)
}

type LookupAppServiceEnvironmentPrivateEndpointConnectionOutputArgs struct {
	Name                          pulumi.StringInput `pulumi:"name"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServiceEnvironmentPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceEnvironmentPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceEnvironmentPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) ToLookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput() LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) ToLookupAppServiceEnvironmentPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentPrivateEndpointConnectionResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentPrivateEndpointConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentPrivateEndpointConnectionResult) *ArmIdWrapperResponse {
		return v.PrivateEndpoint
	}).(ArmIdWrapperResponsePtrOutput)
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentPrivateEndpointConnectionResult) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServiceEnvironmentPrivateEndpointConnectionResultOutput{})
}
