


package v20160627preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRegistryCredentials(ctx *pulumi.Context, args *GetRegistryCredentialsArgs, opts ...pulumi.InvokeOption) (*GetRegistryCredentialsResult, error) {
	var rv GetRegistryCredentialsResult
	err := ctx.Invoke("azure-native:containerregistry/v20160627preview:getRegistryCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRegistryCredentialsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetRegistryCredentialsResult struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}

func GetRegistryCredentialsOutput(ctx *pulumi.Context, args GetRegistryCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetRegistryCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegistryCredentialsResult, error) {
			args := v.(GetRegistryCredentialsArgs)
			r, err := GetRegistryCredentials(ctx, &args, opts...)
			var s GetRegistryCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegistryCredentialsResultOutput)
}

type GetRegistryCredentialsOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetRegistryCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryCredentialsArgs)(nil)).Elem()
}


type GetRegistryCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetRegistryCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryCredentialsResult)(nil)).Elem()
}

func (o GetRegistryCredentialsResultOutput) ToGetRegistryCredentialsResultOutput() GetRegistryCredentialsResultOutput {
	return o
}

func (o GetRegistryCredentialsResultOutput) ToGetRegistryCredentialsResultOutputWithContext(ctx context.Context) GetRegistryCredentialsResultOutput {
	return o
}

func (o GetRegistryCredentialsResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryCredentialsResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GetRegistryCredentialsResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryCredentialsResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegistryCredentialsResultOutput{})
}
