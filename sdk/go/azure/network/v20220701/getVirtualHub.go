


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHub(ctx *pulumi.Context, args *LookupVirtualHubArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubResult, error) {
	var rv LookupVirtualHubResult
	err := ctx.Invoke("azure-native:network/v20220701:getVirtualHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubResult struct {
	AddressPrefix                       *string                                      `pulumi:"addressPrefix"`
	AllowBranchToBranchTraffic          *bool                                        `pulumi:"allowBranchToBranchTraffic"`
	AzureFirewall                       *SubResourceResponse                         `pulumi:"azureFirewall"`
	BgpConnections                      []SubResourceResponse                        `pulumi:"bgpConnections"`
	Etag                                string                                       `pulumi:"etag"`
	ExpressRouteGateway                 *SubResourceResponse                         `pulumi:"expressRouteGateway"`
	HubRoutingPreference                *string                                      `pulumi:"hubRoutingPreference"`
	Id                                  *string                                      `pulumi:"id"`
	IpConfigurations                    []SubResourceResponse                        `pulumi:"ipConfigurations"`
	Kind                                string                                       `pulumi:"kind"`
	Location                            string                                       `pulumi:"location"`
	Name                                string                                       `pulumi:"name"`
	P2SVpnGateway                       *SubResourceResponse                         `pulumi:"p2SVpnGateway"`
	PreferredRoutingGateway             *string                                      `pulumi:"preferredRoutingGateway"`
	ProvisioningState                   string                                       `pulumi:"provisioningState"`
	RouteMaps                           []SubResourceResponse                        `pulumi:"routeMaps"`
	RouteTable                          *VirtualHubRouteTableResponse                `pulumi:"routeTable"`
	RoutingState                        string                                       `pulumi:"routingState"`
	SecurityPartnerProvider             *SubResourceResponse                         `pulumi:"securityPartnerProvider"`
	SecurityProviderName                *string                                      `pulumi:"securityProviderName"`
	Sku                                 *string                                      `pulumi:"sku"`
	Tags                                map[string]string                            `pulumi:"tags"`
	Type                                string                                       `pulumi:"type"`
	VirtualHubRouteTableV2s             []VirtualHubRouteTableV2Response             `pulumi:"virtualHubRouteTableV2s"`
	VirtualRouterAsn                    *float64                                     `pulumi:"virtualRouterAsn"`
	VirtualRouterAutoScaleConfiguration *VirtualRouterAutoScaleConfigurationResponse `pulumi:"virtualRouterAutoScaleConfiguration"`
	VirtualRouterIps                    []string                                     `pulumi:"virtualRouterIps"`
	VirtualWan                          *SubResourceResponse                         `pulumi:"virtualWan"`
	VpnGateway                          *SubResourceResponse                         `pulumi:"vpnGateway"`
}

func LookupVirtualHubOutput(ctx *pulumi.Context, args LookupVirtualHubOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubResult, error) {
			args := v.(LookupVirtualHubArgs)
			r, err := LookupVirtualHub(ctx, &args, opts...)
			var s LookupVirtualHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubResultOutput)
}

type LookupVirtualHubOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubArgs)(nil)).Elem()
}


type LookupVirtualHubResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubResult)(nil)).Elem()
}

func (o LookupVirtualHubResultOutput) ToLookupVirtualHubResultOutput() LookupVirtualHubResultOutput {
	return o
}

func (o LookupVirtualHubResultOutput) ToLookupVirtualHubResultOutputWithContext(ctx context.Context) LookupVirtualHubResultOutput {
	return o
}

func (o LookupVirtualHubResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubResultOutput) AllowBranchToBranchTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *bool { return v.AllowBranchToBranchTraffic }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualHubResultOutput) AzureFirewall() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.AzureFirewall }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualHubResultOutput) BgpConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []SubResourceResponse { return v.BgpConnections }).(SubResourceResponseArrayOutput)
}

func (o LookupVirtualHubResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) ExpressRouteGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.ExpressRouteGateway }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualHubResultOutput) HubRoutingPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.HubRoutingPreference }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubResultOutput) IpConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []SubResourceResponse { return v.IpConfigurations }).(SubResourceResponseArrayOutput)
}

func (o LookupVirtualHubResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) P2SVpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.P2SVpnGateway }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualHubResultOutput) PreferredRoutingGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.PreferredRoutingGateway }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) RouteMaps() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []SubResourceResponse { return v.RouteMaps }).(SubResourceResponseArrayOutput)
}

func (o LookupVirtualHubResultOutput) RouteTable() VirtualHubRouteTableResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *VirtualHubRouteTableResponse { return v.RouteTable }).(VirtualHubRouteTableResponsePtrOutput)
}

func (o LookupVirtualHubResultOutput) RoutingState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.RoutingState }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) SecurityPartnerProvider() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.SecurityPartnerProvider }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualHubResultOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) VirtualHubRouteTableV2s() VirtualHubRouteTableV2ResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []VirtualHubRouteTableV2Response { return v.VirtualHubRouteTableV2s }).(VirtualHubRouteTableV2ResponseArrayOutput)
}

func (o LookupVirtualHubResultOutput) VirtualRouterAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *float64 { return v.VirtualRouterAsn }).(pulumi.Float64PtrOutput)
}

func (o LookupVirtualHubResultOutput) VirtualRouterAutoScaleConfiguration() VirtualRouterAutoScaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *VirtualRouterAutoScaleConfigurationResponse {
		return v.VirtualRouterAutoScaleConfiguration
	}).(VirtualRouterAutoScaleConfigurationResponsePtrOutput)
}

func (o LookupVirtualHubResultOutput) VirtualRouterIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []string { return v.VirtualRouterIps }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualHubResultOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualHubResultOutput) VpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.VpnGateway }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubResultOutput{})
}
