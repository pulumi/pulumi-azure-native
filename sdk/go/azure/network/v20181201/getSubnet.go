


package v20181201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	var rv LookupSubnetResult
	err := ctx.Invoke("azure-native:network/v20181201:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubnetArgs struct {
	Expand             *string `pulumi:"expand"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SubnetName         string  `pulumi:"subnetName"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type LookupSubnetResult struct {
	AddressPrefix           *string                                   `pulumi:"addressPrefix"`
	AddressPrefixes         []string                                  `pulumi:"addressPrefixes"`
	Delegations             []DelegationResponse                      `pulumi:"delegations"`
	Etag                    *string                                   `pulumi:"etag"`
	Id                      *string                                   `pulumi:"id"`
	InterfaceEndpoints      []InterfaceEndpointResponse               `pulumi:"interfaceEndpoints"`
	IpConfigurationProfiles []IPConfigurationProfileResponse          `pulumi:"ipConfigurationProfiles"`
	IpConfigurations        []IPConfigurationResponse                 `pulumi:"ipConfigurations"`
	Name                    *string                                   `pulumi:"name"`
	NetworkSecurityGroup    *NetworkSecurityGroupResponse             `pulumi:"networkSecurityGroup"`
	ProvisioningState       *string                                   `pulumi:"provisioningState"`
	Purpose                 string                                    `pulumi:"purpose"`
	ResourceNavigationLinks []ResourceNavigationLinkResponse          `pulumi:"resourceNavigationLinks"`
	RouteTable              *RouteTableResponse                       `pulumi:"routeTable"`
	ServiceAssociationLinks []ServiceAssociationLinkResponse          `pulumi:"serviceAssociationLinks"`
	ServiceEndpointPolicies []ServiceEndpointPolicyResponse           `pulumi:"serviceEndpointPolicies"`
	ServiceEndpoints        []ServiceEndpointPropertiesFormatResponse `pulumi:"serviceEndpoints"`
}

func LookupSubnetOutput(ctx *pulumi.Context, args LookupSubnetOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetResult, error) {
			args := v.(LookupSubnetArgs)
			r, err := LookupSubnet(ctx, &args, opts...)
			var s LookupSubnetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetResultOutput)
}

type LookupSubnetOutputArgs struct {
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	SubnetName         pulumi.StringInput    `pulumi:"subnetName"`
	VirtualNetworkName pulumi.StringInput    `pulumi:"virtualNetworkName"`
}

func (LookupSubnetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetArgs)(nil)).Elem()
}


type LookupSubnetResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetResult)(nil)).Elem()
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutput() LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutputWithContext(ctx context.Context) LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o LookupSubnetResultOutput) Delegations() DelegationResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []DelegationResponse { return v.Delegations }).(DelegationResponseArrayOutput)
}

func (o LookupSubnetResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) InterfaceEndpoints() InterfaceEndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []InterfaceEndpointResponse { return v.InterfaceEndpoints }).(InterfaceEndpointResponseArrayOutput)
}

func (o LookupSubnetResultOutput) IpConfigurationProfiles() IPConfigurationProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []IPConfigurationProfileResponse { return v.IpConfigurationProfiles }).(IPConfigurationProfileResponseArrayOutput)
}

func (o LookupSubnetResultOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []IPConfigurationResponse { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

func (o LookupSubnetResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o LookupSubnetResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Purpose }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []ResourceNavigationLinkResponse { return v.ResourceNavigationLinks }).(ResourceNavigationLinkResponseArrayOutput)
}

func (o LookupSubnetResultOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *RouteTableResponse { return v.RouteTable }).(RouteTableResponsePtrOutput)
}

func (o LookupSubnetResultOutput) ServiceAssociationLinks() ServiceAssociationLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []ServiceAssociationLinkResponse { return v.ServiceAssociationLinks }).(ServiceAssociationLinkResponseArrayOutput)
}

func (o LookupSubnetResultOutput) ServiceEndpointPolicies() ServiceEndpointPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []ServiceEndpointPolicyResponse { return v.ServiceEndpointPolicies }).(ServiceEndpointPolicyResponseArrayOutput)
}

func (o LookupSubnetResultOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []ServiceEndpointPropertiesFormatResponse { return v.ServiceEndpoints }).(ServiceEndpointPropertiesFormatResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetResultOutput{})
}
