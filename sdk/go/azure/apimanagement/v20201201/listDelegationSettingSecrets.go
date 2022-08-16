


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDelegationSettingSecrets(ctx *pulumi.Context, args *ListDelegationSettingSecretsArgs, opts ...pulumi.InvokeOption) (*ListDelegationSettingSecretsResult, error) {
	var rv ListDelegationSettingSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:listDelegationSettingSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDelegationSettingSecretsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListDelegationSettingSecretsResult struct {
	ValidationKey *string `pulumi:"validationKey"`
}

func ListDelegationSettingSecretsOutput(ctx *pulumi.Context, args ListDelegationSettingSecretsOutputArgs, opts ...pulumi.InvokeOption) ListDelegationSettingSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDelegationSettingSecretsResult, error) {
			args := v.(ListDelegationSettingSecretsArgs)
			r, err := ListDelegationSettingSecrets(ctx, &args, opts...)
			var s ListDelegationSettingSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDelegationSettingSecretsResultOutput)
}

type ListDelegationSettingSecretsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListDelegationSettingSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDelegationSettingSecretsArgs)(nil)).Elem()
}


type ListDelegationSettingSecretsResultOutput struct{ *pulumi.OutputState }

func (ListDelegationSettingSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDelegationSettingSecretsResult)(nil)).Elem()
}

func (o ListDelegationSettingSecretsResultOutput) ToListDelegationSettingSecretsResultOutput() ListDelegationSettingSecretsResultOutput {
	return o
}

func (o ListDelegationSettingSecretsResultOutput) ToListDelegationSettingSecretsResultOutputWithContext(ctx context.Context) ListDelegationSettingSecretsResultOutput {
	return o
}

func (o ListDelegationSettingSecretsResultOutput) ValidationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListDelegationSettingSecretsResult) *string { return v.ValidationKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDelegationSettingSecretsResultOutput{})
}
