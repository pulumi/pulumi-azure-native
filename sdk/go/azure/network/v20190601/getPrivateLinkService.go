


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkService(ctx *pulumi.Context, args *LookupPrivateLinkServiceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServiceResult, error) {
	var rv LookupPrivateLinkServiceResult
	err := ctx.Invoke("azure-native:network/v20190601:getPrivateLinkService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkServiceArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupPrivateLinkServiceResult struct {
	Alias                                string                                            `pulumi:"alias"`
	AutoApproval                         *PrivateLinkServicePropertiesResponseAutoApproval `pulumi:"autoApproval"`
	Etag                                 *string                                           `pulumi:"etag"`
	Fqdns                                []string                                          `pulumi:"fqdns"`
	Id                                   *string                                           `pulumi:"id"`
	IpConfigurations                     []PrivateLinkServiceIpConfigurationResponse       `pulumi:"ipConfigurations"`
	LoadBalancerFrontendIpConfigurations []FrontendIPConfigurationResponse                 `pulumi:"loadBalancerFrontendIpConfigurations"`
	Location                             *string                                           `pulumi:"location"`
	Name                                 string                                            `pulumi:"name"`
	NetworkInterfaces                    []NetworkInterfaceResponse                        `pulumi:"networkInterfaces"`
	PrivateEndpointConnections           []PrivateEndpointConnectionResponse               `pulumi:"privateEndpointConnections"`
	ProvisioningState                    string                                            `pulumi:"provisioningState"`
	Tags                                 map[string]string                                 `pulumi:"tags"`
	Type                                 string                                            `pulumi:"type"`
	Visibility                           *PrivateLinkServicePropertiesResponseVisibility   `pulumi:"visibility"`
}

func LookupPrivateLinkServiceOutput(ctx *pulumi.Context, args LookupPrivateLinkServiceOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkServiceResult, error) {
			args := v.(LookupPrivateLinkServiceArgs)
			r, err := LookupPrivateLinkService(ctx, &args, opts...)
			var s LookupPrivateLinkServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkServiceResultOutput)
}

type LookupPrivateLinkServiceOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput    `pulumi:"serviceName"`
}

func (LookupPrivateLinkServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServiceArgs)(nil)).Elem()
}


type LookupPrivateLinkServiceResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServiceResult)(nil)).Elem()
}

func (o LookupPrivateLinkServiceResultOutput) ToLookupPrivateLinkServiceResultOutput() LookupPrivateLinkServiceResultOutput {
	return o
}

func (o LookupPrivateLinkServiceResultOutput) ToLookupPrivateLinkServiceResultOutputWithContext(ctx context.Context) LookupPrivateLinkServiceResultOutput {
	return o
}

func (o LookupPrivateLinkServiceResultOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.Alias }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkServiceResultOutput) AutoApproval() PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *PrivateLinkServicePropertiesResponseAutoApproval {
		return v.AutoApproval
	}).(PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput)
}

func (o LookupPrivateLinkServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateLinkServiceResultOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []string { return v.Fqdns }).(pulumi.StringArrayOutput)
}

func (o LookupPrivateLinkServiceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateLinkServiceResultOutput) IpConfigurations() PrivateLinkServiceIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []PrivateLinkServiceIpConfigurationResponse {
		return v.IpConfigurations
	}).(PrivateLinkServiceIpConfigurationResponseArrayOutput)
}

func (o LookupPrivateLinkServiceResultOutput) LoadBalancerFrontendIpConfigurations() FrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []FrontendIPConfigurationResponse {
		return v.LoadBalancerFrontendIpConfigurations
	}).(FrontendIPConfigurationResponseArrayOutput)
}

func (o LookupPrivateLinkServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateLinkServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkServiceResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o LookupPrivateLinkServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupPrivateLinkServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateLinkServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkServiceResultOutput) Visibility() PrivateLinkServicePropertiesResponseVisibilityPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *PrivateLinkServicePropertiesResponseVisibility {
		return v.Visibility
	}).(PrivateLinkServicePropertiesResponseVisibilityPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkServiceResultOutput{})
}
