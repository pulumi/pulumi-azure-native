


package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationRuntimeAuthKeys(ctx *pulumi.Context, args *ListIntegrationRuntimeAuthKeysArgs, opts ...pulumi.InvokeOption) (*ListIntegrationRuntimeAuthKeysResult, error) {
	var rv ListIntegrationRuntimeAuthKeysResult
	err := ctx.Invoke("azure-native:datafactory:listIntegrationRuntimeAuthKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationRuntimeAuthKeysArgs struct {
	FactoryName            string `pulumi:"factoryName"`
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type ListIntegrationRuntimeAuthKeysResult struct {
	AuthKey1 *string `pulumi:"authKey1"`
	AuthKey2 *string `pulumi:"authKey2"`
}

func ListIntegrationRuntimeAuthKeysOutput(ctx *pulumi.Context, args ListIntegrationRuntimeAuthKeysOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationRuntimeAuthKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationRuntimeAuthKeysResult, error) {
			args := v.(ListIntegrationRuntimeAuthKeysArgs)
			r, err := ListIntegrationRuntimeAuthKeys(ctx, &args, opts...)
			var s ListIntegrationRuntimeAuthKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationRuntimeAuthKeysResultOutput)
}

type ListIntegrationRuntimeAuthKeysOutputArgs struct {
	FactoryName            pulumi.StringInput `pulumi:"factoryName"`
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListIntegrationRuntimeAuthKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationRuntimeAuthKeysArgs)(nil)).Elem()
}


type ListIntegrationRuntimeAuthKeysResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationRuntimeAuthKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationRuntimeAuthKeysResult)(nil)).Elem()
}

func (o ListIntegrationRuntimeAuthKeysResultOutput) ToListIntegrationRuntimeAuthKeysResultOutput() ListIntegrationRuntimeAuthKeysResultOutput {
	return o
}

func (o ListIntegrationRuntimeAuthKeysResultOutput) ToListIntegrationRuntimeAuthKeysResultOutputWithContext(ctx context.Context) ListIntegrationRuntimeAuthKeysResultOutput {
	return o
}

func (o ListIntegrationRuntimeAuthKeysResultOutput) AuthKey1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIntegrationRuntimeAuthKeysResult) *string { return v.AuthKey1 }).(pulumi.StringPtrOutput)
}

func (o ListIntegrationRuntimeAuthKeysResultOutput) AuthKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIntegrationRuntimeAuthKeysResult) *string { return v.AuthKey2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationRuntimeAuthKeysResultOutput{})
}
