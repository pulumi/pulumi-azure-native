


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTenantAccessSecrets(ctx *pulumi.Context, args *ListTenantAccessSecretsArgs, opts ...pulumi.InvokeOption) (*ListTenantAccessSecretsResult, error) {
	var rv ListTenantAccessSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:listTenantAccessSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTenantAccessSecretsArgs struct {
	AccessName        string `pulumi:"accessName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListTenantAccessSecretsResult struct {
	Enabled      *bool   `pulumi:"enabled"`
	Id           *string `pulumi:"id"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	PrincipalId  *string `pulumi:"principalId"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListTenantAccessSecretsOutput(ctx *pulumi.Context, args ListTenantAccessSecretsOutputArgs, opts ...pulumi.InvokeOption) ListTenantAccessSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTenantAccessSecretsResult, error) {
			args := v.(ListTenantAccessSecretsArgs)
			r, err := ListTenantAccessSecrets(ctx, &args, opts...)
			var s ListTenantAccessSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTenantAccessSecretsResultOutput)
}

type ListTenantAccessSecretsOutputArgs struct {
	AccessName        pulumi.StringInput `pulumi:"accessName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListTenantAccessSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTenantAccessSecretsArgs)(nil)).Elem()
}


type ListTenantAccessSecretsResultOutput struct{ *pulumi.OutputState }

func (ListTenantAccessSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTenantAccessSecretsResult)(nil)).Elem()
}

func (o ListTenantAccessSecretsResultOutput) ToListTenantAccessSecretsResultOutput() ListTenantAccessSecretsResultOutput {
	return o
}

func (o ListTenantAccessSecretsResultOutput) ToListTenantAccessSecretsResultOutputWithContext(ctx context.Context) ListTenantAccessSecretsResultOutput {
	return o
}

func (o ListTenantAccessSecretsResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListTenantAccessSecretsResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListTenantAccessSecretsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTenantAccessSecretsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListTenantAccessSecretsResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTenantAccessSecretsResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListTenantAccessSecretsResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTenantAccessSecretsResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ListTenantAccessSecretsResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTenantAccessSecretsResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTenantAccessSecretsResultOutput{})
}
