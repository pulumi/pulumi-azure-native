


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateCloud(ctx *pulumi.Context, args *LookupPrivateCloudArgs, opts ...pulumi.InvokeOption) (*LookupPrivateCloudResult, error) {
	var rv LookupPrivateCloudResult
	err := ctx.Invoke("azure-native:avs/v20220501:getPrivateCloud", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateCloudArgs struct {
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPrivateCloudResult struct {
	Availability                 *AvailabilityPropertiesResponse `pulumi:"availability"`
	Circuit                      *CircuitResponse                `pulumi:"circuit"`
	Encryption                   *EncryptionResponse             `pulumi:"encryption"`
	Endpoints                    EndpointsResponse               `pulumi:"endpoints"`
	ExternalCloudLinks           []string                        `pulumi:"externalCloudLinks"`
	Id                           string                          `pulumi:"id"`
	Identity                     *PrivateCloudIdentityResponse   `pulumi:"identity"`
	IdentitySources              []IdentitySourceResponse        `pulumi:"identitySources"`
	Internet                     *string                         `pulumi:"internet"`
	Location                     string                          `pulumi:"location"`
	ManagementCluster            ManagementClusterResponse       `pulumi:"managementCluster"`
	ManagementNetwork            string                          `pulumi:"managementNetwork"`
	Name                         string                          `pulumi:"name"`
	NetworkBlock                 string                          `pulumi:"networkBlock"`
	NsxPublicIpQuotaRaised       string                          `pulumi:"nsxPublicIpQuotaRaised"`
	NsxtCertificateThumbprint    string                          `pulumi:"nsxtCertificateThumbprint"`
	NsxtPassword                 *string                         `pulumi:"nsxtPassword"`
	ProvisioningNetwork          string                          `pulumi:"provisioningNetwork"`
	ProvisioningState            string                          `pulumi:"provisioningState"`
	SecondaryCircuit             *CircuitResponse                `pulumi:"secondaryCircuit"`
	Sku                          SkuResponse                     `pulumi:"sku"`
	Tags                         map[string]string               `pulumi:"tags"`
	Type                         string                          `pulumi:"type"`
	VcenterCertificateThumbprint string                          `pulumi:"vcenterCertificateThumbprint"`
	VcenterPassword              *string                         `pulumi:"vcenterPassword"`
	VmotionNetwork               string                          `pulumi:"vmotionNetwork"`
}


func (val *LookupPrivateCloudResult) Defaults() *LookupPrivateCloudResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Internet) {
		internet_ := "Disabled"
		tmp.Internet = &internet_
	}
	return &tmp
}

func LookupPrivateCloudOutput(ctx *pulumi.Context, args LookupPrivateCloudOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateCloudResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateCloudResult, error) {
			args := v.(LookupPrivateCloudArgs)
			r, err := LookupPrivateCloud(ctx, &args, opts...)
			var s LookupPrivateCloudResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateCloudResultOutput)
}

type LookupPrivateCloudOutputArgs struct {
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateCloudOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateCloudArgs)(nil)).Elem()
}


type LookupPrivateCloudResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateCloudResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateCloudResult)(nil)).Elem()
}

func (o LookupPrivateCloudResultOutput) ToLookupPrivateCloudResultOutput() LookupPrivateCloudResultOutput {
	return o
}

func (o LookupPrivateCloudResultOutput) ToLookupPrivateCloudResultOutputWithContext(ctx context.Context) LookupPrivateCloudResultOutput {
	return o
}

func (o LookupPrivateCloudResultOutput) Availability() AvailabilityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) *AvailabilityPropertiesResponse { return v.Availability }).(AvailabilityPropertiesResponsePtrOutput)
}

func (o LookupPrivateCloudResultOutput) Circuit() CircuitResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) *CircuitResponse { return v.Circuit }).(CircuitResponsePtrOutput)
}

func (o LookupPrivateCloudResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

func (o LookupPrivateCloudResultOutput) Endpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) EndpointsResponse { return v.Endpoints }).(EndpointsResponseOutput)
}

func (o LookupPrivateCloudResultOutput) ExternalCloudLinks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) []string { return v.ExternalCloudLinks }).(pulumi.StringArrayOutput)
}

func (o LookupPrivateCloudResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) Identity() PrivateCloudIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) *PrivateCloudIdentityResponse { return v.Identity }).(PrivateCloudIdentityResponsePtrOutput)
}

func (o LookupPrivateCloudResultOutput) IdentitySources() IdentitySourceResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) []IdentitySourceResponse { return v.IdentitySources }).(IdentitySourceResponseArrayOutput)
}

func (o LookupPrivateCloudResultOutput) Internet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) *string { return v.Internet }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateCloudResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) ManagementCluster() ManagementClusterResponseOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) ManagementClusterResponse { return v.ManagementCluster }).(ManagementClusterResponseOutput)
}

func (o LookupPrivateCloudResultOutput) ManagementNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.ManagementNetwork }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) NetworkBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.NetworkBlock }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) NsxPublicIpQuotaRaised() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.NsxPublicIpQuotaRaised }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) NsxtCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.NsxtCertificateThumbprint }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) NsxtPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) *string { return v.NsxtPassword }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateCloudResultOutput) ProvisioningNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.ProvisioningNetwork }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) SecondaryCircuit() CircuitResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) *CircuitResponse { return v.SecondaryCircuit }).(CircuitResponsePtrOutput)
}

func (o LookupPrivateCloudResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupPrivateCloudResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateCloudResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) VcenterCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.VcenterCertificateThumbprint }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) VcenterPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) *string { return v.VcenterPassword }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateCloudResultOutput) VmotionNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.VmotionNetwork }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateCloudResultOutput{})
}
