


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationRuntimeAuthKey(ctx *pulumi.Context, args *ListIntegrationRuntimeAuthKeyArgs, opts ...pulumi.InvokeOption) (*ListIntegrationRuntimeAuthKeyResult, error) {
	var rv ListIntegrationRuntimeAuthKeyResult
	err := ctx.Invoke("azure-native:synapse/v20210301:listIntegrationRuntimeAuthKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationRuntimeAuthKeyArgs struct {
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type ListIntegrationRuntimeAuthKeyResult struct {
	AuthKey1 *string `pulumi:"authKey1"`
	AuthKey2 *string `pulumi:"authKey2"`
}

func ListIntegrationRuntimeAuthKeyOutput(ctx *pulumi.Context, args ListIntegrationRuntimeAuthKeyOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationRuntimeAuthKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationRuntimeAuthKeyResult, error) {
			args := v.(ListIntegrationRuntimeAuthKeyArgs)
			r, err := ListIntegrationRuntimeAuthKey(ctx, &args, opts...)
			var s ListIntegrationRuntimeAuthKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationRuntimeAuthKeyResultOutput)
}

type ListIntegrationRuntimeAuthKeyOutputArgs struct {
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName          pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListIntegrationRuntimeAuthKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationRuntimeAuthKeyArgs)(nil)).Elem()
}


type ListIntegrationRuntimeAuthKeyResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationRuntimeAuthKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationRuntimeAuthKeyResult)(nil)).Elem()
}

func (o ListIntegrationRuntimeAuthKeyResultOutput) ToListIntegrationRuntimeAuthKeyResultOutput() ListIntegrationRuntimeAuthKeyResultOutput {
	return o
}

func (o ListIntegrationRuntimeAuthKeyResultOutput) ToListIntegrationRuntimeAuthKeyResultOutputWithContext(ctx context.Context) ListIntegrationRuntimeAuthKeyResultOutput {
	return o
}

func (o ListIntegrationRuntimeAuthKeyResultOutput) AuthKey1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIntegrationRuntimeAuthKeyResult) *string { return v.AuthKey1 }).(pulumi.StringPtrOutput)
}

func (o ListIntegrationRuntimeAuthKeyResultOutput) AuthKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIntegrationRuntimeAuthKeyResult) *string { return v.AuthKey2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationRuntimeAuthKeyResultOutput{})
}
