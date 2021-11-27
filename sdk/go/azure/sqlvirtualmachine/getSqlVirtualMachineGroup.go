


package sqlvirtualmachine

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlVirtualMachineGroup(ctx *pulumi.Context, args *LookupSqlVirtualMachineGroupArgs, opts ...pulumi.InvokeOption) (*LookupSqlVirtualMachineGroupResult, error) {
	var rv LookupSqlVirtualMachineGroupResult
	err := ctx.Invoke("azure-native:sqlvirtualmachine:getSqlVirtualMachineGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlVirtualMachineGroupArgs struct {
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	SqlVirtualMachineGroupName string `pulumi:"sqlVirtualMachineGroupName"`
}


type LookupSqlVirtualMachineGroupResult struct {
	ClusterConfiguration string                     `pulumi:"clusterConfiguration"`
	ClusterManagerType   string                     `pulumi:"clusterManagerType"`
	Id                   string                     `pulumi:"id"`
	Location             string                     `pulumi:"location"`
	Name                 string                     `pulumi:"name"`
	ProvisioningState    string                     `pulumi:"provisioningState"`
	ScaleType            string                     `pulumi:"scaleType"`
	SqlImageOffer        *string                    `pulumi:"sqlImageOffer"`
	SqlImageSku          *string                    `pulumi:"sqlImageSku"`
	Tags                 map[string]string          `pulumi:"tags"`
	Type                 string                     `pulumi:"type"`
	WsfcDomainProfile    *WsfcDomainProfileResponse `pulumi:"wsfcDomainProfile"`
}
