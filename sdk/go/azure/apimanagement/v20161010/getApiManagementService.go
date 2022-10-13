


package v20161010

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiManagementService(ctx *pulumi.Context, args *LookupApiManagementServiceArgs, opts ...pulumi.InvokeOption) (*LookupApiManagementServiceResult, error) {
	var rv LookupApiManagementServiceResult
	err := ctx.Invoke("azure-native:apimanagement/v20161010:getApiManagementService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiManagementServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiManagementServiceResult struct {
	AdditionalLocations     []AdditionalRegionResponse                `pulumi:"additionalLocations"`
	AddresserEmail          *string                                   `pulumi:"addresserEmail"`
	CreatedAtUtc            string                                    `pulumi:"createdAtUtc"`
	CustomProperties        map[string]string                         `pulumi:"customProperties"`
	Etag                    string                                    `pulumi:"etag"`
	HostnameConfigurations  []HostnameConfigurationResponse           `pulumi:"hostnameConfigurations"`
	Id                      string                                    `pulumi:"id"`
	Location                string                                    `pulumi:"location"`
	ManagementApiUrl        string                                    `pulumi:"managementApiUrl"`
	Name                    *string                                   `pulumi:"name"`
	PortalUrl               string                                    `pulumi:"portalUrl"`
	ProvisioningState       string                                    `pulumi:"provisioningState"`
	PublisherEmail          string                                    `pulumi:"publisherEmail"`
	PublisherName           string                                    `pulumi:"publisherName"`
	RuntimeUrl              string                                    `pulumi:"runtimeUrl"`
	ScmUrl                  string                                    `pulumi:"scmUrl"`
	Sku                     ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	StaticIPs               []string                                  `pulumi:"staticIPs"`
	Tags                    map[string]string                         `pulumi:"tags"`
	TargetProvisioningState string                                    `pulumi:"targetProvisioningState"`
	Type                    string                                    `pulumi:"type"`
	VpnType                 *string                                   `pulumi:"vpnType"`
	Vpnconfiguration        *VirtualNetworkConfigurationResponse      `pulumi:"vpnconfiguration"`
}


func (val *LookupApiManagementServiceResult) Defaults() *LookupApiManagementServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = *tmp.Sku.Defaults()

	if isZero(tmp.VpnType) {
		vpnType_ := "None"
		tmp.VpnType = &vpnType_
	}
	return &tmp
}

func LookupApiManagementServiceOutput(ctx *pulumi.Context, args LookupApiManagementServiceOutputArgs, opts ...pulumi.InvokeOption) LookupApiManagementServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiManagementServiceResult, error) {
			args := v.(LookupApiManagementServiceArgs)
			r, err := LookupApiManagementService(ctx, &args, opts...)
			var s LookupApiManagementServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiManagementServiceResultOutput)
}

type LookupApiManagementServiceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiManagementServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiManagementServiceArgs)(nil)).Elem()
}


type LookupApiManagementServiceResultOutput struct{ *pulumi.OutputState }

func (LookupApiManagementServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiManagementServiceResult)(nil)).Elem()
}

func (o LookupApiManagementServiceResultOutput) ToLookupApiManagementServiceResultOutput() LookupApiManagementServiceResultOutput {
	return o
}

func (o LookupApiManagementServiceResultOutput) ToLookupApiManagementServiceResultOutputWithContext(ctx context.Context) LookupApiManagementServiceResultOutput {
	return o
}

func (o LookupApiManagementServiceResultOutput) AdditionalLocations() AdditionalRegionResponseArrayOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) []AdditionalRegionResponse { return v.AdditionalLocations }).(AdditionalRegionResponseArrayOutput)
}

func (o LookupApiManagementServiceResultOutput) AddresserEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *string { return v.AddresserEmail }).(pulumi.StringPtrOutput)
}

func (o LookupApiManagementServiceResultOutput) CreatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.CreatedAtUtc }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) CustomProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) map[string]string { return v.CustomProperties }).(pulumi.StringMapOutput)
}

func (o LookupApiManagementServiceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) HostnameConfigurations() HostnameConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) []HostnameConfigurationResponse {
		return v.HostnameConfigurations
	}).(HostnameConfigurationResponseArrayOutput)
}

func (o LookupApiManagementServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) ManagementApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.ManagementApiUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupApiManagementServiceResultOutput) PortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.PortalUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) PublisherEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.PublisherEmail }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) PublisherName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.PublisherName }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) RuntimeUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.RuntimeUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) ScmUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.ScmUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) Sku() ApiManagementServiceSkuPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) ApiManagementServiceSkuPropertiesResponse { return v.Sku }).(ApiManagementServiceSkuPropertiesResponseOutput)
}

func (o LookupApiManagementServiceResultOutput) StaticIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) []string { return v.StaticIPs }).(pulumi.StringArrayOutput)
}

func (o LookupApiManagementServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApiManagementServiceResultOutput) TargetProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.TargetProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *string { return v.VpnType }).(pulumi.StringPtrOutput)
}

func (o LookupApiManagementServiceResultOutput) Vpnconfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *VirtualNetworkConfigurationResponse {
		return v.Vpnconfiguration
	}).(VirtualNetworkConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiManagementServiceResultOutput{})
}
