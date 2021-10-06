


package v20150521preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFormulaResource(ctx *pulumi.Context, args *LookupFormulaResourceArgs, opts ...pulumi.InvokeOption) (*LookupFormulaResourceResult, error) {
	var rv LookupFormulaResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getFormulaResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFormulaResourceArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFormulaResourceResult struct {
	Author            *string                          `pulumi:"author"`
	CreationDate      *string                          `pulumi:"creationDate"`
	Description       *string                          `pulumi:"description"`
	FormulaContent    *LabVirtualMachineResponse       `pulumi:"formulaContent"`
	Id                *string                          `pulumi:"id"`
	Location          *string                          `pulumi:"location"`
	Name              *string                          `pulumi:"name"`
	OsType            *string                          `pulumi:"osType"`
	ProvisioningState *string                          `pulumi:"provisioningState"`
	Tags              map[string]string                `pulumi:"tags"`
	Type              *string                          `pulumi:"type"`
	Vm                *FormulaPropertiesFromVmResponse `pulumi:"vm"`
}
