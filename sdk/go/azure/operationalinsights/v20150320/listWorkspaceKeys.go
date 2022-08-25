


package v20150320

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceKeys(ctx *pulumi.Context, args *ListWorkspaceKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceKeysResult, error) {
	var rv ListWorkspaceKeysResult
	err := ctx.Invoke("azure-native:operationalinsights/v20150320:listWorkspaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListWorkspaceKeysResult struct {
	PrimarySharedKey   *string `pulumi:"primarySharedKey"`
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
}

func ListWorkspaceKeysOutput(ctx *pulumi.Context, args ListWorkspaceKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceKeysResult, error) {
			args := v.(ListWorkspaceKeysArgs)
			r, err := ListWorkspaceKeys(ctx, &args, opts...)
			var s ListWorkspaceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceKeysResultOutput)
}

type ListWorkspaceKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceKeysArgs)(nil)).Elem()
}


type ListWorkspaceKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceKeysResult)(nil)).Elem()
}

func (o ListWorkspaceKeysResultOutput) ToListWorkspaceKeysResultOutput() ListWorkspaceKeysResultOutput {
	return o
}

func (o ListWorkspaceKeysResultOutput) ToListWorkspaceKeysResultOutputWithContext(ctx context.Context) ListWorkspaceKeysResultOutput {
	return o
}

func (o ListWorkspaceKeysResultOutput) PrimarySharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) *string { return v.PrimarySharedKey }).(pulumi.StringPtrOutput)
}

func (o ListWorkspaceKeysResultOutput) SecondarySharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) *string { return v.SecondarySharedKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceKeysResultOutput{})
}
