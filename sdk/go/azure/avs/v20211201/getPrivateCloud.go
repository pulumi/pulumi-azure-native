


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateCloud(ctx *pulumi.Context, args *LookupPrivateCloudArgs, opts ...pulumi.InvokeOption) (*LookupPrivateCloudResult, error) {
	var rv LookupPrivateCloudResult
	err := ctx.Invoke("azure-native:avs/v20211201:getPrivateCloud", args, &rv, opts...)
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
