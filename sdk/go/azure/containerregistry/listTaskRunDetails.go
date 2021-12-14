


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTaskRunDetails(ctx *pulumi.Context, args *ListTaskRunDetailsArgs, opts ...pulumi.InvokeOption) (*ListTaskRunDetailsResult, error) {
	var rv ListTaskRunDetailsResult
	err := ctx.Invoke("azure-native:containerregistry:listTaskRunDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListTaskRunDetailsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TaskRunName       string `pulumi:"taskRunName"`
}



type ListTaskRunDetailsResult struct {
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


func (val *ListTaskRunDetailsResult) Defaults() *ListTaskRunDetailsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.RunResult = *tmp.RunResult.Defaults()

	return &tmp
}
