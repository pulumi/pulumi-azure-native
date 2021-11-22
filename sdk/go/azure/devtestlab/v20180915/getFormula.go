


package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFormula(ctx *pulumi.Context, args *LookupFormulaArgs, opts ...pulumi.InvokeOption) (*LookupFormulaResult, error) {
	var rv LookupFormulaResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getFormula", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFormulaArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupFormulaResult struct {
	Author            string                                      `pulumi:"author"`
	CreationDate      string                                      `pulumi:"creationDate"`
	Description       *string                                     `pulumi:"description"`
	FormulaContent    *LabVirtualMachineCreationParameterResponse `pulumi:"formulaContent"`
	Id                string                                      `pulumi:"id"`
	Location          *string                                     `pulumi:"location"`
	Name              string                                      `pulumi:"name"`
	OsType            *string                                     `pulumi:"osType"`
	ProvisioningState string                                      `pulumi:"provisioningState"`
	Tags              map[string]string                           `pulumi:"tags"`
	Type              string                                      `pulumi:"type"`
	UniqueIdentifier  string                                      `pulumi:"uniqueIdentifier"`
	Vm                *FormulaPropertiesFromVmResponse            `pulumi:"vm"`
}
