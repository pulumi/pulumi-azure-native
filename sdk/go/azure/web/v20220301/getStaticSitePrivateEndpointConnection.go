


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSitePrivateEndpointConnection(ctx *pulumi.Context, args *LookupStaticSitePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupStaticSitePrivateEndpointConnectionResult, error) {
	var rv LookupStaticSitePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:web/v20220301:getStaticSitePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSitePrivateEndpointConnectionArgs struct {
	Name                          string `pulumi:"name"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupStaticSitePrivateEndpointConnectionResult struct {
	Id                                string                              `pulumi:"id"`
	IpAddresses                       []string                            `pulumi:"ipAddresses"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}

func LookupStaticSitePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupStaticSitePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSitePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSitePrivateEndpointConnectionResult, error) {
			args := v.(LookupStaticSitePrivateEndpointConnectionArgs)
			r, err := LookupStaticSitePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupStaticSitePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSitePrivateEndpointConnectionResultOutput)
}

type LookupStaticSitePrivateEndpointConnectionOutputArgs struct {
	Name                          pulumi.StringInput `pulumi:"name"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSitePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSitePrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupStaticSitePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSitePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSitePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) ToLookupStaticSitePrivateEndpointConnectionResultOutput() LookupStaticSitePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) ToLookupStaticSitePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupStaticSitePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSitePrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStaticSitePrivateEndpointConnectionResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSitePrivateEndpointConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSitePrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSitePrivateEndpointConnectionResult) *ArmIdWrapperResponse {
		return v.PrivateEndpoint
	}).(ArmIdWrapperResponsePtrOutput)
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSitePrivateEndpointConnectionResult) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSitePrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStaticSitePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSitePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSitePrivateEndpointConnectionResultOutput{})
}
