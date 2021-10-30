


package v20211101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:servicefabric/v20211101preview:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	Id                string                                    `pulumi:"id"`
	Identity          *ManagedIdentityResponse                  `pulumi:"identity"`
	Location          *string                                   `pulumi:"location"`
	ManagedIdentities []ApplicationUserAssignedIdentityResponse `pulumi:"managedIdentities"`
	Name              string                                    `pulumi:"name"`
	Parameters        map[string]string                         `pulumi:"parameters"`
	ProvisioningState string                                    `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                        `pulumi:"systemData"`
	Tags              map[string]string                         `pulumi:"tags"`
	Type              string                                    `pulumi:"type"`
	UpgradePolicy     *ApplicationUpgradePolicyResponse         `pulumi:"upgradePolicy"`
	Version           *string                                   `pulumi:"version"`
}
