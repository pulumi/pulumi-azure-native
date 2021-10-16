


package automation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridRunbookWorkerGroup(ctx *pulumi.Context, args *LookupHybridRunbookWorkerGroupArgs, opts ...pulumi.InvokeOption) (*LookupHybridRunbookWorkerGroupResult, error) {
	var rv LookupHybridRunbookWorkerGroupResult
	err := ctx.Invoke("azure-native:automation:getHybridRunbookWorkerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridRunbookWorkerGroupArgs struct {
	AutomationAccountName        string `pulumi:"automationAccountName"`
	HybridRunbookWorkerGroupName string `pulumi:"hybridRunbookWorkerGroupName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupHybridRunbookWorkerGroupResult struct {
	Credential           *RunAsCredentialAssociationPropertyResponse `pulumi:"credential"`
	GroupType            *string                                     `pulumi:"groupType"`
	HybridRunbookWorkers []HybridRunbookWorkerLegacyResponse         `pulumi:"hybridRunbookWorkers"`
	Id                   *string                                     `pulumi:"id"`
	Name                 *string                                     `pulumi:"name"`
	SystemData           SystemDataResponse                          `pulumi:"systemData"`
	Type                 string                                      `pulumi:"type"`
}
