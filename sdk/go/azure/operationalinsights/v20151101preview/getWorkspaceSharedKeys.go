


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWorkspaceSharedKeys(ctx *pulumi.Context, args *GetWorkspaceSharedKeysArgs, opts ...pulumi.InvokeOption) (*GetWorkspaceSharedKeysResult, error) {
	var rv GetWorkspaceSharedKeysResult
	err := ctx.Invoke("azure-native:operationalinsights/v20151101preview:getWorkspaceSharedKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetWorkspaceSharedKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetWorkspaceSharedKeysResult struct {
	PrimarySharedKey   *string `pulumi:"primarySharedKey"`
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
}

func GetWorkspaceSharedKeysOutput(ctx *pulumi.Context, args GetWorkspaceSharedKeysOutputArgs, opts ...pulumi.InvokeOption) GetWorkspaceSharedKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWorkspaceSharedKeysResult, error) {
			args := v.(GetWorkspaceSharedKeysArgs)
			r, err := GetWorkspaceSharedKeys(ctx, &args, opts...)
			var s GetWorkspaceSharedKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWorkspaceSharedKeysResultOutput)
}

type GetWorkspaceSharedKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetWorkspaceSharedKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspaceSharedKeysArgs)(nil)).Elem()
}


type GetWorkspaceSharedKeysResultOutput struct{ *pulumi.OutputState }

func (GetWorkspaceSharedKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspaceSharedKeysResult)(nil)).Elem()
}

func (o GetWorkspaceSharedKeysResultOutput) ToGetWorkspaceSharedKeysResultOutput() GetWorkspaceSharedKeysResultOutput {
	return o
}

func (o GetWorkspaceSharedKeysResultOutput) ToGetWorkspaceSharedKeysResultOutputWithContext(ctx context.Context) GetWorkspaceSharedKeysResultOutput {
	return o
}

func (o GetWorkspaceSharedKeysResultOutput) PrimarySharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWorkspaceSharedKeysResult) *string { return v.PrimarySharedKey }).(pulumi.StringPtrOutput)
}

func (o GetWorkspaceSharedKeysResultOutput) SecondarySharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWorkspaceSharedKeysResult) *string { return v.SecondarySharedKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWorkspaceSharedKeysResultOutput{})
}
