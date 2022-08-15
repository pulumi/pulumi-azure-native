


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHubVirtualNetworkConnection(ctx *pulumi.Context, args *LookupHubVirtualNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupHubVirtualNetworkConnectionResult, error) {
	var rv LookupHubVirtualNetworkConnectionResult
	err := ctx.Invoke("azure-native:network/v20210301:getHubVirtualNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubVirtualNetworkConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupHubVirtualNetworkConnectionResult struct {
	AllowHubToRemoteVnetTransit         *bool                         `pulumi:"allowHubToRemoteVnetTransit"`
	AllowRemoteVnetToUseHubVnetGateways *bool                         `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	EnableInternetSecurity              *bool                         `pulumi:"enableInternetSecurity"`
	Etag                                string                        `pulumi:"etag"`
	Id                                  *string                       `pulumi:"id"`
	Name                                *string                       `pulumi:"name"`
	ProvisioningState                   string                        `pulumi:"provisioningState"`
	RemoteVirtualNetwork                *SubResourceResponse          `pulumi:"remoteVirtualNetwork"`
	RoutingConfiguration                *RoutingConfigurationResponse `pulumi:"routingConfiguration"`
}

func LookupHubVirtualNetworkConnectionOutput(ctx *pulumi.Context, args LookupHubVirtualNetworkConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupHubVirtualNetworkConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHubVirtualNetworkConnectionResult, error) {
			args := v.(LookupHubVirtualNetworkConnectionArgs)
			r, err := LookupHubVirtualNetworkConnection(ctx, &args, opts...)
			var s LookupHubVirtualNetworkConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHubVirtualNetworkConnectionResultOutput)
}

type LookupHubVirtualNetworkConnectionOutputArgs struct {
	ConnectionName    pulumi.StringInput `pulumi:"connectionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupHubVirtualNetworkConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubVirtualNetworkConnectionArgs)(nil)).Elem()
}


type LookupHubVirtualNetworkConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupHubVirtualNetworkConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubVirtualNetworkConnectionResult)(nil)).Elem()
}

func (o LookupHubVirtualNetworkConnectionResultOutput) ToLookupHubVirtualNetworkConnectionResultOutput() LookupHubVirtualNetworkConnectionResultOutput {
	return o
}

func (o LookupHubVirtualNetworkConnectionResultOutput) ToLookupHubVirtualNetworkConnectionResultOutputWithContext(ctx context.Context) LookupHubVirtualNetworkConnectionResultOutput {
	return o
}

func (o LookupHubVirtualNetworkConnectionResultOutput) AllowHubToRemoteVnetTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *bool { return v.AllowHubToRemoteVnetTransit }).(pulumi.BoolPtrOutput)
}

func (o LookupHubVirtualNetworkConnectionResultOutput) AllowRemoteVnetToUseHubVnetGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *bool { return v.AllowRemoteVnetToUseHubVnetGateways }).(pulumi.BoolPtrOutput)
}

func (o LookupHubVirtualNetworkConnectionResultOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *bool { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

func (o LookupHubVirtualNetworkConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupHubVirtualNetworkConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupHubVirtualNetworkConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupHubVirtualNetworkConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupHubVirtualNetworkConnectionResultOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *SubResourceResponse { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o LookupHubVirtualNetworkConnectionResultOutput) RoutingConfiguration() RoutingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *RoutingConfigurationResponse {
		return v.RoutingConfiguration
	}).(RoutingConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHubVirtualNetworkConnectionResultOutput{})
}
