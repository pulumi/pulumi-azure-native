


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiManagementService(ctx *pulumi.Context, args *LookupApiManagementServiceArgs, opts ...pulumi.InvokeOption) (*LookupApiManagementServiceResult, error) {
	var rv LookupApiManagementServiceResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:getApiManagementService", args, &rv, opts...)
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
	AdditionalLocations         []AdditionalLocationResponse              `pulumi:"additionalLocations"`
	ApiVersionConstraint        *ApiVersionConstraintResponse             `pulumi:"apiVersionConstraint"`
	Certificates                []CertificateConfigurationResponse        `pulumi:"certificates"`
	CreatedAtUtc                string                                    `pulumi:"createdAtUtc"`
	CustomProperties            map[string]string                         `pulumi:"customProperties"`
	DeveloperPortalUrl          string                                    `pulumi:"developerPortalUrl"`
	DisableGateway              *bool                                     `pulumi:"disableGateway"`
	EnableClientCertificate     *bool                                     `pulumi:"enableClientCertificate"`
	Etag                        string                                    `pulumi:"etag"`
	GatewayRegionalUrl          string                                    `pulumi:"gatewayRegionalUrl"`
	GatewayUrl                  string                                    `pulumi:"gatewayUrl"`
	HostnameConfigurations      []HostnameConfigurationResponse           `pulumi:"hostnameConfigurations"`
	Id                          string                                    `pulumi:"id"`
	Identity                    *ApiManagementServiceIdentityResponse     `pulumi:"identity"`
	Location                    string                                    `pulumi:"location"`
	ManagementApiUrl            string                                    `pulumi:"managementApiUrl"`
	Name                        string                                    `pulumi:"name"`
	NotificationSenderEmail     *string                                   `pulumi:"notificationSenderEmail"`
	PortalUrl                   string                                    `pulumi:"portalUrl"`
	PrivateIPAddresses          []string                                  `pulumi:"privateIPAddresses"`
	ProvisioningState           string                                    `pulumi:"provisioningState"`
	PublicIPAddresses           []string                                  `pulumi:"publicIPAddresses"`
	PublisherEmail              string                                    `pulumi:"publisherEmail"`
	PublisherName               string                                    `pulumi:"publisherName"`
	ScmUrl                      string                                    `pulumi:"scmUrl"`
	Sku                         ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	Tags                        map[string]string                         `pulumi:"tags"`
	TargetProvisioningState     string                                    `pulumi:"targetProvisioningState"`
	Type                        string                                    `pulumi:"type"`
	VirtualNetworkConfiguration *VirtualNetworkConfigurationResponse      `pulumi:"virtualNetworkConfiguration"`
	VirtualNetworkType          *string                                   `pulumi:"virtualNetworkType"`
}


func (val *LookupApiManagementServiceResult) Defaults() *LookupApiManagementServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableGateway) {
		disableGateway_ := false
		tmp.DisableGateway = &disableGateway_
	}
	if isZero(tmp.EnableClientCertificate) {
		enableClientCertificate_ := false
		tmp.EnableClientCertificate = &enableClientCertificate_
	}
	if isZero(tmp.VirtualNetworkType) {
		virtualNetworkType_ := "None"
		tmp.VirtualNetworkType = &virtualNetworkType_
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

func (o LookupApiManagementServiceResultOutput) AdditionalLocations() AdditionalLocationResponseArrayOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) []AdditionalLocationResponse { return v.AdditionalLocations }).(AdditionalLocationResponseArrayOutput)
}

func (o LookupApiManagementServiceResultOutput) ApiVersionConstraint() ApiVersionConstraintResponsePtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *ApiVersionConstraintResponse { return v.ApiVersionConstraint }).(ApiVersionConstraintResponsePtrOutput)
}

func (o LookupApiManagementServiceResultOutput) Certificates() CertificateConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) []CertificateConfigurationResponse { return v.Certificates }).(CertificateConfigurationResponseArrayOutput)
}

func (o LookupApiManagementServiceResultOutput) CreatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.CreatedAtUtc }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) CustomProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) map[string]string { return v.CustomProperties }).(pulumi.StringMapOutput)
}

func (o LookupApiManagementServiceResultOutput) DeveloperPortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.DeveloperPortalUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) DisableGateway() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *bool { return v.DisableGateway }).(pulumi.BoolPtrOutput)
}

func (o LookupApiManagementServiceResultOutput) EnableClientCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *bool { return v.EnableClientCertificate }).(pulumi.BoolPtrOutput)
}

func (o LookupApiManagementServiceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) GatewayRegionalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.GatewayRegionalUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) GatewayUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.GatewayUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) HostnameConfigurations() HostnameConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) []HostnameConfigurationResponse {
		return v.HostnameConfigurations
	}).(HostnameConfigurationResponseArrayOutput)
}

func (o LookupApiManagementServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) Identity() ApiManagementServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *ApiManagementServiceIdentityResponse { return v.Identity }).(ApiManagementServiceIdentityResponsePtrOutput)
}

func (o LookupApiManagementServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) ManagementApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.ManagementApiUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) NotificationSenderEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *string { return v.NotificationSenderEmail }).(pulumi.StringPtrOutput)
}

func (o LookupApiManagementServiceResultOutput) PortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.PortalUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) PrivateIPAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) []string { return v.PrivateIPAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupApiManagementServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) PublicIPAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) []string { return v.PublicIPAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupApiManagementServiceResultOutput) PublisherEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.PublisherEmail }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) PublisherName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.PublisherName }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) ScmUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) string { return v.ScmUrl }).(pulumi.StringOutput)
}

func (o LookupApiManagementServiceResultOutput) Sku() ApiManagementServiceSkuPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) ApiManagementServiceSkuPropertiesResponse { return v.Sku }).(ApiManagementServiceSkuPropertiesResponseOutput)
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

func (o LookupApiManagementServiceResultOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *VirtualNetworkConfigurationResponse {
		return v.VirtualNetworkConfiguration
	}).(VirtualNetworkConfigurationResponsePtrOutput)
}

func (o LookupApiManagementServiceResultOutput) VirtualNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiManagementServiceResult) *string { return v.VirtualNetworkType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiManagementServiceResultOutput{})
}
