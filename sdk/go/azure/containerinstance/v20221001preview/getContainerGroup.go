


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerGroup(ctx *pulumi.Context, args *LookupContainerGroupArgs, opts ...pulumi.InvokeOption) (*LookupContainerGroupResult, error) {
	var rv LookupContainerGroupResult
	err := ctx.Invoke("azure-native:containerinstance/v20221001preview:getContainerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupContainerGroupArgs struct {
	ContainerGroupName string `pulumi:"containerGroupName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupContainerGroupResult struct {
	ConfidentialComputeProperties *ConfidentialComputePropertiesResponse       `pulumi:"confidentialComputeProperties"`
	Containers                    []ContainerResponse                          `pulumi:"containers"`
	Diagnostics                   *ContainerGroupDiagnosticsResponse           `pulumi:"diagnostics"`
	DnsConfig                     *DnsConfigurationResponse                    `pulumi:"dnsConfig"`
	EncryptionProperties          *EncryptionPropertiesResponse                `pulumi:"encryptionProperties"`
	Extensions                    []DeploymentExtensionSpecResponse            `pulumi:"extensions"`
	Id                            string                                       `pulumi:"id"`
	Identity                      *ContainerGroupIdentityResponse              `pulumi:"identity"`
	ImageRegistryCredentials      []ImageRegistryCredentialResponse            `pulumi:"imageRegistryCredentials"`
	InitContainers                []InitContainerDefinitionResponse            `pulumi:"initContainers"`
	InstanceView                  ContainerGroupPropertiesResponseInstanceView `pulumi:"instanceView"`
	IpAddress                     *IpAddressResponse                           `pulumi:"ipAddress"`
	Location                      *string                                      `pulumi:"location"`
	Name                          string                                       `pulumi:"name"`
	OsType                        string                                       `pulumi:"osType"`
	Priority                      *string                                      `pulumi:"priority"`
	ProvisioningState             string                                       `pulumi:"provisioningState"`
	RestartPolicy                 *string                                      `pulumi:"restartPolicy"`
	Sku                           *string                                      `pulumi:"sku"`
	SubnetIds                     []ContainerGroupSubnetIdResponse             `pulumi:"subnetIds"`
	Tags                          map[string]string                            `pulumi:"tags"`
	Type                          string                                       `pulumi:"type"`
	Volumes                       []VolumeResponse                             `pulumi:"volumes"`
	Zones                         []string                                     `pulumi:"zones"`
}


func (val *LookupContainerGroupResult) Defaults() *LookupContainerGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.IpAddress = tmp.IpAddress.Defaults()

	return &tmp
}

func LookupContainerGroupOutput(ctx *pulumi.Context, args LookupContainerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupContainerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerGroupResult, error) {
			args := v.(LookupContainerGroupArgs)
			r, err := LookupContainerGroup(ctx, &args, opts...)
			var s LookupContainerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerGroupResultOutput)
}

type LookupContainerGroupOutputArgs struct {
	ContainerGroupName pulumi.StringInput `pulumi:"containerGroupName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupArgs)(nil)).Elem()
}


type LookupContainerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupContainerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupResult)(nil)).Elem()
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutput() LookupContainerGroupResultOutput {
	return o
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutputWithContext(ctx context.Context) LookupContainerGroupResultOutput {
	return o
}

func (o LookupContainerGroupResultOutput) ConfidentialComputeProperties() ConfidentialComputePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ConfidentialComputePropertiesResponse {
		return v.ConfidentialComputeProperties
	}).(ConfidentialComputePropertiesResponsePtrOutput)
}

func (o LookupContainerGroupResultOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ContainerResponse { return v.Containers }).(ContainerResponseArrayOutput)
}

func (o LookupContainerGroupResultOutput) Diagnostics() ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ContainerGroupDiagnosticsResponse { return v.Diagnostics }).(ContainerGroupDiagnosticsResponsePtrOutput)
}

func (o LookupContainerGroupResultOutput) DnsConfig() DnsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *DnsConfigurationResponse { return v.DnsConfig }).(DnsConfigurationResponsePtrOutput)
}

func (o LookupContainerGroupResultOutput) EncryptionProperties() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *EncryptionPropertiesResponse { return v.EncryptionProperties }).(EncryptionPropertiesResponsePtrOutput)
}

func (o LookupContainerGroupResultOutput) Extensions() DeploymentExtensionSpecResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []DeploymentExtensionSpecResponse { return v.Extensions }).(DeploymentExtensionSpecResponseArrayOutput)
}

func (o LookupContainerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerGroupResultOutput) Identity() ContainerGroupIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ContainerGroupIdentityResponse { return v.Identity }).(ContainerGroupIdentityResponsePtrOutput)
}

func (o LookupContainerGroupResultOutput) ImageRegistryCredentials() ImageRegistryCredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ImageRegistryCredentialResponse {
		return v.ImageRegistryCredentials
	}).(ImageRegistryCredentialResponseArrayOutput)
}

func (o LookupContainerGroupResultOutput) InitContainers() InitContainerDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []InitContainerDefinitionResponse { return v.InitContainers }).(InitContainerDefinitionResponseArrayOutput)
}

func (o LookupContainerGroupResultOutput) InstanceView() ContainerGroupPropertiesResponseInstanceViewOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) ContainerGroupPropertiesResponseInstanceView { return v.InstanceView }).(ContainerGroupPropertiesResponseInstanceViewOutput)
}

func (o LookupContainerGroupResultOutput) IpAddress() IpAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *IpAddressResponse { return v.IpAddress }).(IpAddressResponsePtrOutput)
}

func (o LookupContainerGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupContainerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContainerGroupResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.OsType }).(pulumi.StringOutput)
}

func (o LookupContainerGroupResultOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

func (o LookupContainerGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupContainerGroupResultOutput) RestartPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.RestartPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupContainerGroupResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LookupContainerGroupResultOutput) SubnetIds() ContainerGroupSubnetIdResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ContainerGroupSubnetIdResponse { return v.SubnetIds }).(ContainerGroupSubnetIdResponseArrayOutput)
}

func (o LookupContainerGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupContainerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupContainerGroupResultOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []VolumeResponse { return v.Volumes }).(VolumeResponseArrayOutput)
}

func (o LookupContainerGroupResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerGroupResultOutput{})
}
