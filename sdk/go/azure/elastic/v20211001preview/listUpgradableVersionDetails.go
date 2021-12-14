


package v20211001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListUpgradableVersionDetails(ctx *pulumi.Context, args *ListUpgradableVersionDetailsArgs, opts ...pulumi.InvokeOption) (*ListUpgradableVersionDetailsResult, error) {
	var rv ListUpgradableVersionDetailsResult
	err := ctx.Invoke("azure-native:elastic/v20211001preview:listUpgradableVersionDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListUpgradableVersionDetailsArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListUpgradableVersionDetailsResult struct {
	CurrentVersion     *string  `pulumi:"currentVersion"`
	UpgradableVersions []string `pulumi:"upgradableVersions"`
}
