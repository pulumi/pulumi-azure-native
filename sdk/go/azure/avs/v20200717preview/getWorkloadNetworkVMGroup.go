


package v20200717preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkVMGroup(ctx *pulumi.Context, args *LookupWorkloadNetworkVMGroupArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkVMGroupResult, error) {
	var rv LookupWorkloadNetworkVMGroupResult
	err := ctx.Invoke("azure-native:avs/v20200717preview:getWorkloadNetworkVMGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkVMGroupArgs struct {
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VmGroupId         string `pulumi:"vmGroupId"`
}


type LookupWorkloadNetworkVMGroupResult struct {
	DisplayName       *string  `pulumi:"displayName"`
	Id                string   `pulumi:"id"`
	Members           []string `pulumi:"members"`
	Name              string   `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Status            string   `pulumi:"status"`
	Type              string   `pulumi:"type"`
}
