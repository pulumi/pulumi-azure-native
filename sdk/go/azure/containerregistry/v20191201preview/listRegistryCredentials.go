


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRegistryCredentials(ctx *pulumi.Context, args *ListRegistryCredentialsArgs, opts ...pulumi.InvokeOption) (*ListRegistryCredentialsResult, error) {
	var rv ListRegistryCredentialsResult
	err := ctx.Invoke("azure-native:containerregistry/v20191201preview:listRegistryCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryCredentialsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListRegistryCredentialsResult struct {
	Passwords []RegistryPasswordResponse `pulumi:"passwords"`
	Username  *string                    `pulumi:"username"`
}

func ListRegistryCredentialsOutput(ctx *pulumi.Context, args ListRegistryCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListRegistryCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRegistryCredentialsResult, error) {
			args := v.(ListRegistryCredentialsArgs)
			r, err := ListRegistryCredentials(ctx, &args, opts...)
			var s ListRegistryCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRegistryCredentialsResultOutput)
}

type ListRegistryCredentialsOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListRegistryCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryCredentialsArgs)(nil)).Elem()
}


type ListRegistryCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListRegistryCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryCredentialsResult)(nil)).Elem()
}

func (o ListRegistryCredentialsResultOutput) ToListRegistryCredentialsResultOutput() ListRegistryCredentialsResultOutput {
	return o
}

func (o ListRegistryCredentialsResultOutput) ToListRegistryCredentialsResultOutputWithContext(ctx context.Context) ListRegistryCredentialsResultOutput {
	return o
}

func (o ListRegistryCredentialsResultOutput) Passwords() RegistryPasswordResponseArrayOutput {
	return o.ApplyT(func(v ListRegistryCredentialsResult) []RegistryPasswordResponse { return v.Passwords }).(RegistryPasswordResponseArrayOutput)
}

func (o ListRegistryCredentialsResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListRegistryCredentialsResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRegistryCredentialsResultOutput{})
}
