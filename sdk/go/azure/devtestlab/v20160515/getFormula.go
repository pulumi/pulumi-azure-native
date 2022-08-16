


package v20160515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupFormula(ctx *pulumi.Context, args *LookupFormulaArgs, opts ...pulumi.InvokeOption) (*LookupFormulaResult, error) {
	var rv LookupFormulaResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getFormula", args, &rv, opts...)
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
	Author            *string                                     `pulumi:"author"`
	CreationDate      string                                      `pulumi:"creationDate"`
	Description       *string                                     `pulumi:"description"`
	FormulaContent    *LabVirtualMachineCreationParameterResponse `pulumi:"formulaContent"`
	Id                string                                      `pulumi:"id"`
	Location          *string                                     `pulumi:"location"`
	Name              string                                      `pulumi:"name"`
	OsType            *string                                     `pulumi:"osType"`
	ProvisioningState *string                                     `pulumi:"provisioningState"`
	Tags              map[string]string                           `pulumi:"tags"`
	Type              string                                      `pulumi:"type"`
	UniqueIdentifier  *string                                     `pulumi:"uniqueIdentifier"`
	Vm                *FormulaPropertiesFromVmResponse            `pulumi:"vm"`
}

func LookupFormulaOutput(ctx *pulumi.Context, args LookupFormulaOutputArgs, opts ...pulumi.InvokeOption) LookupFormulaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFormulaResult, error) {
			args := v.(LookupFormulaArgs)
			r, err := LookupFormula(ctx, &args, opts...)
			var s LookupFormulaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFormulaResultOutput)
}

type LookupFormulaOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupFormulaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFormulaArgs)(nil)).Elem()
}


type LookupFormulaResultOutput struct{ *pulumi.OutputState }

func (LookupFormulaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFormulaResult)(nil)).Elem()
}

func (o LookupFormulaResultOutput) ToLookupFormulaResultOutput() LookupFormulaResultOutput {
	return o
}

func (o LookupFormulaResultOutput) ToLookupFormulaResultOutputWithContext(ctx context.Context) LookupFormulaResultOutput {
	return o
}

func (o LookupFormulaResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupFormulaResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResultOutput) FormulaContent() LabVirtualMachineCreationParameterResponsePtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *LabVirtualMachineCreationParameterResponse { return v.FormulaContent }).(LabVirtualMachineCreationParameterResponsePtrOutput)
}

func (o LookupFormulaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFormulaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFormulaResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFormulaResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFormulaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFormulaResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupFormulaResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResultOutput) Vm() FormulaPropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v LookupFormulaResult) *FormulaPropertiesFromVmResponse { return v.Vm }).(FormulaPropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFormulaResultOutput{})
}
