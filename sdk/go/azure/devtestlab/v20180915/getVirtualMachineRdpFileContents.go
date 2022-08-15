


package v20180915

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualMachineRdpFileContents(ctx *pulumi.Context, args *GetVirtualMachineRdpFileContentsArgs, opts ...pulumi.InvokeOption) (*GetVirtualMachineRdpFileContentsResult, error) {
	var rv GetVirtualMachineRdpFileContentsResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getVirtualMachineRdpFileContents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualMachineRdpFileContentsArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetVirtualMachineRdpFileContentsResult struct {
	Contents *string `pulumi:"contents"`
}

func GetVirtualMachineRdpFileContentsOutput(ctx *pulumi.Context, args GetVirtualMachineRdpFileContentsOutputArgs, opts ...pulumi.InvokeOption) GetVirtualMachineRdpFileContentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualMachineRdpFileContentsResult, error) {
			args := v.(GetVirtualMachineRdpFileContentsArgs)
			r, err := GetVirtualMachineRdpFileContents(ctx, &args, opts...)
			var s GetVirtualMachineRdpFileContentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualMachineRdpFileContentsResultOutput)
}

type GetVirtualMachineRdpFileContentsOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetVirtualMachineRdpFileContentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualMachineRdpFileContentsArgs)(nil)).Elem()
}


type GetVirtualMachineRdpFileContentsResultOutput struct{ *pulumi.OutputState }

func (GetVirtualMachineRdpFileContentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualMachineRdpFileContentsResult)(nil)).Elem()
}

func (o GetVirtualMachineRdpFileContentsResultOutput) ToGetVirtualMachineRdpFileContentsResultOutput() GetVirtualMachineRdpFileContentsResultOutput {
	return o
}

func (o GetVirtualMachineRdpFileContentsResultOutput) ToGetVirtualMachineRdpFileContentsResultOutputWithContext(ctx context.Context) GetVirtualMachineRdpFileContentsResultOutput {
	return o
}

func (o GetVirtualMachineRdpFileContentsResultOutput) Contents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineRdpFileContentsResult) *string { return v.Contents }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualMachineRdpFileContentsResultOutput{})
}
