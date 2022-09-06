


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTenantAccessGitSecrets(ctx *pulumi.Context, args *ListTenantAccessGitSecretsArgs, opts ...pulumi.InvokeOption) (*ListTenantAccessGitSecretsResult, error) {
	var rv ListTenantAccessGitSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:listTenantAccessGitSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTenantAccessGitSecretsArgs struct {
	AccessName        string `pulumi:"accessName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListTenantAccessGitSecretsResult struct {
	Enabled      *bool   `pulumi:"enabled"`
	Id           *string `pulumi:"id"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListTenantAccessGitSecretsOutput(ctx *pulumi.Context, args ListTenantAccessGitSecretsOutputArgs, opts ...pulumi.InvokeOption) ListTenantAccessGitSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTenantAccessGitSecretsResult, error) {
			args := v.(ListTenantAccessGitSecretsArgs)
			r, err := ListTenantAccessGitSecrets(ctx, &args, opts...)
			var s ListTenantAccessGitSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTenantAccessGitSecretsResultOutput)
}

type ListTenantAccessGitSecretsOutputArgs struct {
	AccessName        pulumi.StringInput `pulumi:"accessName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListTenantAccessGitSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTenantAccessGitSecretsArgs)(nil)).Elem()
}


type ListTenantAccessGitSecretsResultOutput struct{ *pulumi.OutputState }

func (ListTenantAccessGitSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTenantAccessGitSecretsResult)(nil)).Elem()
}

func (o ListTenantAccessGitSecretsResultOutput) ToListTenantAccessGitSecretsResultOutput() ListTenantAccessGitSecretsResultOutput {
	return o
}

func (o ListTenantAccessGitSecretsResultOutput) ToListTenantAccessGitSecretsResultOutputWithContext(ctx context.Context) ListTenantAccessGitSecretsResultOutput {
	return o
}

func (o ListTenantAccessGitSecretsResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListTenantAccessGitSecretsResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListTenantAccessGitSecretsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTenantAccessGitSecretsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListTenantAccessGitSecretsResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTenantAccessGitSecretsResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListTenantAccessGitSecretsResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTenantAccessGitSecretsResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTenantAccessGitSecretsResultOutput{})
}
