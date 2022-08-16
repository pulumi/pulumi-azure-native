


package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProjectKeys(ctx *pulumi.Context, args *GetProjectKeysArgs, opts ...pulumi.InvokeOption) (*GetProjectKeysResult, error) {
	var rv GetProjectKeysResult
	err := ctx.Invoke("azure-native:migrate:getProjectKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProjectKeysArgs struct {
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetProjectKeysResult struct {
	WorkspaceId  string `pulumi:"workspaceId"`
	WorkspaceKey string `pulumi:"workspaceKey"`
}

func GetProjectKeysOutput(ctx *pulumi.Context, args GetProjectKeysOutputArgs, opts ...pulumi.InvokeOption) GetProjectKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectKeysResult, error) {
			args := v.(GetProjectKeysArgs)
			r, err := GetProjectKeys(ctx, &args, opts...)
			var s GetProjectKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectKeysResultOutput)
}

type GetProjectKeysOutputArgs struct {
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetProjectKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectKeysArgs)(nil)).Elem()
}


type GetProjectKeysResultOutput struct{ *pulumi.OutputState }

func (GetProjectKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectKeysResult)(nil)).Elem()
}

func (o GetProjectKeysResultOutput) ToGetProjectKeysResultOutput() GetProjectKeysResultOutput {
	return o
}

func (o GetProjectKeysResultOutput) ToGetProjectKeysResultOutputWithContext(ctx context.Context) GetProjectKeysResultOutput {
	return o
}

func (o GetProjectKeysResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectKeysResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o GetProjectKeysResultOutput) WorkspaceKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectKeysResult) string { return v.WorkspaceKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectKeysResultOutput{})
}
