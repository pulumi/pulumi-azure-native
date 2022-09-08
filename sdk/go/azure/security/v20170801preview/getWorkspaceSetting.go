


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceSetting(ctx *pulumi.Context, args *LookupWorkspaceSettingArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceSettingResult, error) {
	var rv LookupWorkspaceSettingResult
	err := ctx.Invoke("azure-native:security/v20170801preview:getWorkspaceSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceSettingArgs struct {
	WorkspaceSettingName string `pulumi:"workspaceSettingName"`
}


type LookupWorkspaceSettingResult struct {
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Scope       string `pulumi:"scope"`
	Type        string `pulumi:"type"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupWorkspaceSettingOutput(ctx *pulumi.Context, args LookupWorkspaceSettingOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceSettingResult, error) {
			args := v.(LookupWorkspaceSettingArgs)
			r, err := LookupWorkspaceSetting(ctx, &args, opts...)
			var s LookupWorkspaceSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceSettingResultOutput)
}

type LookupWorkspaceSettingOutputArgs struct {
	WorkspaceSettingName pulumi.StringInput `pulumi:"workspaceSettingName"`
}

func (LookupWorkspaceSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSettingArgs)(nil)).Elem()
}


type LookupWorkspaceSettingResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSettingResult)(nil)).Elem()
}

func (o LookupWorkspaceSettingResultOutput) ToLookupWorkspaceSettingResultOutput() LookupWorkspaceSettingResultOutput {
	return o
}

func (o LookupWorkspaceSettingResultOutput) ToLookupWorkspaceSettingResultOutputWithContext(ctx context.Context) LookupWorkspaceSettingResultOutput {
	return o
}

func (o LookupWorkspaceSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceSettingResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupWorkspaceSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkspaceSettingResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceSettingResultOutput{})
}
