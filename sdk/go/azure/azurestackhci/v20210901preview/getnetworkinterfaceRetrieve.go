


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetnetworkinterfaceRetrieve(ctx *pulumi.Context, args *GetnetworkinterfaceRetrieveArgs, opts ...pulumi.InvokeOption) (*GetnetworkinterfaceRetrieveResult, error) {
	var rv GetnetworkinterfaceRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getnetworkinterfaceRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetnetworkinterfaceRetrieveArgs struct {
	NetworkinterfacesName string `pulumi:"networkinterfacesName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type GetnetworkinterfaceRetrieveResult struct {
	DnsSettings       *InterfaceDNSSettingsResponse  `pulumi:"dnsSettings"`
	ExtendedLocation  *ExtendedLocationResponse      `pulumi:"extendedLocation"`
	Id                string                         `pulumi:"id"`
	IpConfigurations  []IpConfigurationResponse      `pulumi:"ipConfigurations"`
	Location          string                         `pulumi:"location"`
	MacAddress        *string                        `pulumi:"macAddress"`
	Name              string                         `pulumi:"name"`
	ProvisioningState string                         `pulumi:"provisioningState"`
	ResourceName      *string                        `pulumi:"resourceName"`
	Status            NetworkInterfaceStatusResponse `pulumi:"status"`
	SystemData        SystemDataResponse             `pulumi:"systemData"`
	Tags              map[string]string              `pulumi:"tags"`
	Type              string                         `pulumi:"type"`
}

func GetnetworkinterfaceRetrieveOutput(ctx *pulumi.Context, args GetnetworkinterfaceRetrieveOutputArgs, opts ...pulumi.InvokeOption) GetnetworkinterfaceRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetnetworkinterfaceRetrieveResult, error) {
			args := v.(GetnetworkinterfaceRetrieveArgs)
			r, err := GetnetworkinterfaceRetrieve(ctx, &args, opts...)
			var s GetnetworkinterfaceRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetnetworkinterfaceRetrieveResultOutput)
}

type GetnetworkinterfaceRetrieveOutputArgs struct {
	NetworkinterfacesName pulumi.StringInput `pulumi:"networkinterfacesName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetnetworkinterfaceRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetnetworkinterfaceRetrieveArgs)(nil)).Elem()
}


type GetnetworkinterfaceRetrieveResultOutput struct{ *pulumi.OutputState }

func (GetnetworkinterfaceRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetnetworkinterfaceRetrieveResult)(nil)).Elem()
}

func (o GetnetworkinterfaceRetrieveResultOutput) ToGetnetworkinterfaceRetrieveResultOutput() GetnetworkinterfaceRetrieveResultOutput {
	return o
}

func (o GetnetworkinterfaceRetrieveResultOutput) ToGetnetworkinterfaceRetrieveResultOutputWithContext(ctx context.Context) GetnetworkinterfaceRetrieveResultOutput {
	return o
}

func (o GetnetworkinterfaceRetrieveResultOutput) DnsSettings() InterfaceDNSSettingsResponsePtrOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) *InterfaceDNSSettingsResponse { return v.DnsSettings }).(InterfaceDNSSettingsResponsePtrOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) IpConfigurations() IpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) []IpConfigurationResponse { return v.IpConfigurations }).(IpConfigurationResponseArrayOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) Status() NetworkInterfaceStatusResponseOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) NetworkInterfaceStatusResponse { return v.Status }).(NetworkInterfaceStatusResponseOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetnetworkinterfaceRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetnetworkinterfaceRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetnetworkinterfaceRetrieveResultOutput{})
}
