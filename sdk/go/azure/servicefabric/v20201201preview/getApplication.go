


package v20201201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:servicefabric/v20201201preview:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	Etag                      string                                    `pulumi:"etag"`
	Id                        string                                    `pulumi:"id"`
	Identity                  *ManagedIdentityResponse                  `pulumi:"identity"`
	Location                  *string                                   `pulumi:"location"`
	ManagedIdentities         []ApplicationUserAssignedIdentityResponse `pulumi:"managedIdentities"`
	MaximumNodes              *float64                                  `pulumi:"maximumNodes"`
	Metrics                   []ApplicationMetricDescriptionResponse    `pulumi:"metrics"`
	MinimumNodes              *float64                                  `pulumi:"minimumNodes"`
	Name                      string                                    `pulumi:"name"`
	Parameters                map[string]string                         `pulumi:"parameters"`
	ProvisioningState         string                                    `pulumi:"provisioningState"`
	RemoveApplicationCapacity *bool                                     `pulumi:"removeApplicationCapacity"`
	SystemData                SystemDataResponse                        `pulumi:"systemData"`
	Tags                      map[string]string                         `pulumi:"tags"`
	Type                      string                                    `pulumi:"type"`
	TypeName                  *string                                   `pulumi:"typeName"`
	TypeVersion               *string                                   `pulumi:"typeVersion"`
	UpgradePolicy             *ApplicationUpgradePolicyResponse         `pulumi:"upgradePolicy"`
}


func (val *LookupApplicationResult) Defaults() *LookupApplicationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaximumNodes) {
		maximumNodes_ := 0.0
		tmp.MaximumNodes = &maximumNodes_
	}
	tmp.UpgradePolicy = tmp.UpgradePolicy.Defaults()

	return &tmp
}
