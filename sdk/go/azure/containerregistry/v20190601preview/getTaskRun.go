


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTaskRun(ctx *pulumi.Context, args *LookupTaskRunArgs, opts ...pulumi.InvokeOption) (*LookupTaskRunResult, error) {
	var rv LookupTaskRunResult
	err := ctx.Invoke("azure-native:containerregistry/v20190601preview:getTaskRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTaskRunArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TaskRunName       string `pulumi:"taskRunName"`
}



type LookupTaskRunResult struct {
	ForceUpdateTag    *string                     `pulumi:"forceUpdateTag"`
	Id                string                      `pulumi:"id"`
	Identity          *IdentityPropertiesResponse `pulumi:"identity"`
	Location          *string                     `pulumi:"location"`
	Name              string                      `pulumi:"name"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	RunRequest        interface{}                 `pulumi:"runRequest"`
	RunResult         RunResponse                 `pulumi:"runResult"`
	SystemData        SystemDataResponse          `pulumi:"systemData"`
	Type              string                      `pulumi:"type"`
}


func (val *LookupTaskRunResult) Defaults() *LookupTaskRunResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.RunResult = *tmp.RunResult.Defaults()

	return &tmp
}
