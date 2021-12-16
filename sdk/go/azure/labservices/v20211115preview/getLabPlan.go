


package v20211115preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLabPlan(ctx *pulumi.Context, args *LookupLabPlanArgs, opts ...pulumi.InvokeOption) (*LookupLabPlanResult, error) {
	var rv LookupLabPlanResult
	err := ctx.Invoke("azure-native:labservices/v20211115preview:getLabPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupLabPlanArgs struct {
	LabPlanName       string `pulumi:"labPlanName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLabPlanResult struct {
	AllowedRegions             []string                       `pulumi:"allowedRegions"`
	DefaultAutoShutdownProfile *AutoShutdownProfileResponse   `pulumi:"defaultAutoShutdownProfile"`
	DefaultConnectionProfile   *ConnectionProfileResponse     `pulumi:"defaultConnectionProfile"`
	DefaultNetworkProfile      *LabPlanNetworkProfileResponse `pulumi:"defaultNetworkProfile"`
	Id                         string                         `pulumi:"id"`
	LinkedLmsInstance          *string                        `pulumi:"linkedLmsInstance"`
	Location                   string                         `pulumi:"location"`
	Name                       string                         `pulumi:"name"`
	ProvisioningState          string                         `pulumi:"provisioningState"`
	SharedGalleryId            *string                        `pulumi:"sharedGalleryId"`
	SupportInfo                *SupportInfoResponse           `pulumi:"supportInfo"`
	SystemData                 SystemDataResponse             `pulumi:"systemData"`
	Tags                       map[string]string              `pulumi:"tags"`
	Type                       string                         `pulumi:"type"`
}


func (val *LookupLabPlanResult) Defaults() *LookupLabPlanResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DefaultAutoShutdownProfile = tmp.DefaultAutoShutdownProfile.Defaults()

	tmp.DefaultConnectionProfile = tmp.DefaultConnectionProfile.Defaults()

	return &tmp
}
