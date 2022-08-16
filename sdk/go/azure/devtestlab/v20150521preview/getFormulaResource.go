


package v20150521preview

import (
	"context"
	"reflect"

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

func LookupFormulaResourceOutput(ctx *pulumi.Context, args LookupFormulaResourceOutputArgs, opts ...pulumi.InvokeOption) LookupFormulaResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFormulaResourceResult, error) {
			args := v.(LookupFormulaResourceArgs)
			r, err := LookupFormulaResource(ctx, &args, opts...)
			var s LookupFormulaResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFormulaResourceResultOutput)
}

type LookupFormulaResourceOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFormulaResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFormulaResourceArgs)(nil)).Elem()
}


type LookupFormulaResourceResultOutput struct{ *pulumi.OutputState }

func (LookupFormulaResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFormulaResourceResult)(nil)).Elem()
}

func (o LookupFormulaResourceResultOutput) ToLookupFormulaResourceResultOutput() LookupFormulaResourceResultOutput {
	return o
}

func (o LookupFormulaResourceResultOutput) ToLookupFormulaResourceResultOutputWithContext(ctx context.Context) LookupFormulaResourceResultOutput {
	return o
}

func (o LookupFormulaResourceResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) FormulaContent() LabVirtualMachineResponsePtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *LabVirtualMachineResponse { return v.FormulaContent }).(LabVirtualMachineResponsePtrOutput)
}

func (o LookupFormulaResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFormulaResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupFormulaResourceResultOutput) Vm() FormulaPropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v LookupFormulaResourceResult) *FormulaPropertiesFromVmResponse { return v.Vm }).(FormulaPropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFormulaResourceResultOutput{})
}
