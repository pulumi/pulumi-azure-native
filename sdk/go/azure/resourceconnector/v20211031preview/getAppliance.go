


package v20211031preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppliance(ctx *pulumi.Context, args *LookupApplianceArgs, opts ...pulumi.InvokeOption) (*LookupApplianceResult, error) {
	var rv LookupApplianceResult
	err := ctx.Invoke("azure-native:resourceconnector/v20211031preview:getAppliance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApplianceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupApplianceResult struct {
	Distro               *string                                          `pulumi:"distro"`
	Id                   string                                           `pulumi:"id"`
	Identity             *IdentityResponse                                `pulumi:"identity"`
	InfrastructureConfig *AppliancePropertiesResponseInfrastructureConfig `pulumi:"infrastructureConfig"`
	Location             string                                           `pulumi:"location"`
	Name                 string                                           `pulumi:"name"`
	ProvisioningState    string                                           `pulumi:"provisioningState"`
	PublicKey            *string                                          `pulumi:"publicKey"`
	Status               string                                           `pulumi:"status"`
	SystemData           SystemDataResponse                               `pulumi:"systemData"`
	Tags                 map[string]string                                `pulumi:"tags"`
	Type                 string                                           `pulumi:"type"`
	Version              string                                           `pulumi:"version"`
}


func (val *LookupApplianceResult) Defaults() *LookupApplianceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Distro) {
		distro_ := "AKSEdge"
		tmp.Distro = &distro_
	}
	return &tmp
}
