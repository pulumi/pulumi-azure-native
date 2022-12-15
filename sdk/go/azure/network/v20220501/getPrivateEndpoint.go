


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azure-native:network/v20220501:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateEndpointArgs struct {
	Expand              *string `pulumi:"expand"`
	PrivateEndpointName string  `pulumi:"privateEndpointName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointResult struct {
	ApplicationSecurityGroups           []ApplicationSecurityGroupResponse        `pulumi:"applicationSecurityGroups"`
	CustomDnsConfigs                    []CustomDnsConfigPropertiesFormatResponse `pulumi:"customDnsConfigs"`
	CustomNetworkInterfaceName          *string                                   `pulumi:"customNetworkInterfaceName"`
	Etag                                string                                    `pulumi:"etag"`
	ExtendedLocation                    *ExtendedLocationResponse                 `pulumi:"extendedLocation"`
	Id                                  *string                                   `pulumi:"id"`
	IpConfigurations                    []PrivateEndpointIPConfigurationResponse  `pulumi:"ipConfigurations"`
	Location                            *string                                   `pulumi:"location"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse    `pulumi:"manualPrivateLinkServiceConnections"`
	Name                                string                                    `pulumi:"name"`
	NetworkInterfaces                   []NetworkInterfaceResponse                `pulumi:"networkInterfaces"`
	PrivateLinkServiceConnections       []PrivateLinkServiceConnectionResponse    `pulumi:"privateLinkServiceConnections"`
	ProvisioningState                   string                                    `pulumi:"provisioningState"`
	Subnet                              *SubnetResponse                           `pulumi:"subnet"`
	Tags                                map[string]string                         `pulumi:"tags"`
	Type                                string                                    `pulumi:"type"`
}


func (val *LookupPrivateEndpointResult) Defaults() *LookupPrivateEndpointResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Subnet = tmp.Subnet.Defaults()

	return &tmp
}

func LookupPrivateEndpointOutput(ctx *pulumi.Context, args LookupPrivateEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointResult, error) {
			args := v.(LookupPrivateEndpointArgs)
			r, err := LookupPrivateEndpoint(ctx, &args, opts...)
			var s LookupPrivateEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointResultOutput)
}

type LookupPrivateEndpointOutputArgs struct {
	Expand              pulumi.StringPtrInput `pulumi:"expand"`
	PrivateEndpointName pulumi.StringInput    `pulumi:"privateEndpointName"`
	ResourceGroupName   pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointArgs)(nil)).Elem()
}


type LookupPrivateEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointResult)(nil)).Elem()
}

func (o LookupPrivateEndpointResultOutput) ToLookupPrivateEndpointResultOutput() LookupPrivateEndpointResultOutput {
	return o
}

func (o LookupPrivateEndpointResultOutput) ToLookupPrivateEndpointResultOutputWithContext(ctx context.Context) LookupPrivateEndpointResultOutput {
	return o
}

func (o LookupPrivateEndpointResultOutput) ApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []ApplicationSecurityGroupResponse {
		return v.ApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) CustomDnsConfigs() CustomDnsConfigPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []CustomDnsConfigPropertiesFormatResponse {
		return v.CustomDnsConfigs
	}).(CustomDnsConfigPropertiesFormatResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) CustomNetworkInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.CustomNetworkInterfaceName }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupPrivateEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointResultOutput) IpConfigurations() PrivateEndpointIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateEndpointIPConfigurationResponse {
		return v.IpConfigurations
	}).(PrivateEndpointIPConfigurationResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointResultOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateLinkServiceConnectionResponse {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) PrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateLinkServiceConnectionResponse {
		return v.PrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

func (o LookupPrivateEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointResultOutput{})
}
