


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLabelByWorkspace(ctx *pulumi.Context, args *LookupLabelByWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupLabelByWorkspaceResult, error) {
	var rv LookupLabelByWorkspaceResult
	err := ctx.Invoke("azure-native:easm/v20220401preview:getLabelByWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabelByWorkspaceArgs struct {
	LabelName         string `pulumi:"labelName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLabelByWorkspaceResult struct {
	Color             *string            `pulumi:"color"`
	DisplayName       *string            `pulumi:"displayName"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupLabelByWorkspaceOutput(ctx *pulumi.Context, args LookupLabelByWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupLabelByWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabelByWorkspaceResult, error) {
			args := v.(LookupLabelByWorkspaceArgs)
			r, err := LookupLabelByWorkspace(ctx, &args, opts...)
			var s LookupLabelByWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabelByWorkspaceResultOutput)
}

type LookupLabelByWorkspaceOutputArgs struct {
	LabelName         pulumi.StringInput `pulumi:"labelName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupLabelByWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabelByWorkspaceArgs)(nil)).Elem()
}


type LookupLabelByWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupLabelByWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabelByWorkspaceResult)(nil)).Elem()
}

func (o LookupLabelByWorkspaceResultOutput) ToLookupLabelByWorkspaceResultOutput() LookupLabelByWorkspaceResultOutput {
	return o
}

func (o LookupLabelByWorkspaceResultOutput) ToLookupLabelByWorkspaceResultOutputWithContext(ctx context.Context) LookupLabelByWorkspaceResultOutput {
	return o
}

func (o LookupLabelByWorkspaceResultOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabelByWorkspaceResult) *string { return v.Color }).(pulumi.StringPtrOutput)
}

func (o LookupLabelByWorkspaceResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabelByWorkspaceResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupLabelByWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelByWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabelByWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelByWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabelByWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelByWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLabelByWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLabelByWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLabelByWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelByWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabelByWorkspaceResultOutput{})
}
