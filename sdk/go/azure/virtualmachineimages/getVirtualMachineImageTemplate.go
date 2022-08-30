


package virtualmachineimages

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineImageTemplate(ctx *pulumi.Context, args *LookupVirtualMachineImageTemplateArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineImageTemplateResult, error) {
	var rv LookupVirtualMachineImageTemplateResult
	err := ctx.Invoke("azure-native:virtualmachineimages:getVirtualMachineImageTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineImageTemplateArgs struct {
	ImageTemplateName string `pulumi:"imageTemplateName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVirtualMachineImageTemplateResult struct {
	BuildTimeoutInMinutes *int                               `pulumi:"buildTimeoutInMinutes"`
	Customize             []interface{}                      `pulumi:"customize"`
	Distribute            []interface{}                      `pulumi:"distribute"`
	Id                    string                             `pulumi:"id"`
	Identity              ImageTemplateIdentityResponse      `pulumi:"identity"`
	LastRunStatus         ImageTemplateLastRunStatusResponse `pulumi:"lastRunStatus"`
	Location              string                             `pulumi:"location"`
	Name                  string                             `pulumi:"name"`
	ProvisioningError     ProvisioningErrorResponse          `pulumi:"provisioningError"`
	ProvisioningState     string                             `pulumi:"provisioningState"`
	Source                interface{}                        `pulumi:"source"`
	Tags                  map[string]string                  `pulumi:"tags"`
	Type                  string                             `pulumi:"type"`
	VmProfile             *ImageTemplateVmProfileResponse    `pulumi:"vmProfile"`
}


func (val *LookupVirtualMachineImageTemplateResult) Defaults() *LookupVirtualMachineImageTemplateResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.BuildTimeoutInMinutes) {
		buildTimeoutInMinutes_ := 0
		tmp.BuildTimeoutInMinutes = &buildTimeoutInMinutes_
	}
	tmp.VmProfile = tmp.VmProfile.Defaults()

	return &tmp
}

func LookupVirtualMachineImageTemplateOutput(ctx *pulumi.Context, args LookupVirtualMachineImageTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineImageTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineImageTemplateResult, error) {
			args := v.(LookupVirtualMachineImageTemplateArgs)
			r, err := LookupVirtualMachineImageTemplate(ctx, &args, opts...)
			var s LookupVirtualMachineImageTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineImageTemplateResultOutput)
}

type LookupVirtualMachineImageTemplateOutputArgs struct {
	ImageTemplateName pulumi.StringInput `pulumi:"imageTemplateName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVirtualMachineImageTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineImageTemplateArgs)(nil)).Elem()
}


type LookupVirtualMachineImageTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineImageTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineImageTemplateResult)(nil)).Elem()
}

func (o LookupVirtualMachineImageTemplateResultOutput) ToLookupVirtualMachineImageTemplateResultOutput() LookupVirtualMachineImageTemplateResultOutput {
	return o
}

func (o LookupVirtualMachineImageTemplateResultOutput) ToLookupVirtualMachineImageTemplateResultOutputWithContext(ctx context.Context) LookupVirtualMachineImageTemplateResultOutput {
	return o
}

func (o LookupVirtualMachineImageTemplateResultOutput) BuildTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) *int { return v.BuildTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Customize() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) []interface{} { return v.Customize }).(pulumi.ArrayOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Distribute() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) []interface{} { return v.Distribute }).(pulumi.ArrayOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Identity() ImageTemplateIdentityResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) ImageTemplateIdentityResponse { return v.Identity }).(ImageTemplateIdentityResponseOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) LastRunStatus() ImageTemplateLastRunStatusResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) ImageTemplateLastRunStatusResponse {
		return v.LastRunStatus
	}).(ImageTemplateLastRunStatusResponseOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) ProvisioningError() ProvisioningErrorResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) ProvisioningErrorResponse { return v.ProvisioningError }).(ProvisioningErrorResponseOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Source() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) interface{} { return v.Source }).(pulumi.AnyOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineImageTemplateResultOutput) VmProfile() ImageTemplateVmProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineImageTemplateResult) *ImageTemplateVmProfileResponse { return v.VmProfile }).(ImageTemplateVmProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineImageTemplateResultOutput{})
}
