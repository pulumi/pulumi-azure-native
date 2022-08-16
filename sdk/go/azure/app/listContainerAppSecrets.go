


package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListContainerAppSecrets(ctx *pulumi.Context, args *ListContainerAppSecretsArgs, opts ...pulumi.InvokeOption) (*ListContainerAppSecretsResult, error) {
	var rv ListContainerAppSecretsResult
	err := ctx.Invoke("azure-native:app:listContainerAppSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListContainerAppSecretsArgs struct {
	ContainerAppName  string `pulumi:"containerAppName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListContainerAppSecretsResult struct {
	Value []ContainerAppSecretResponse `pulumi:"value"`
}

func ListContainerAppSecretsOutput(ctx *pulumi.Context, args ListContainerAppSecretsOutputArgs, opts ...pulumi.InvokeOption) ListContainerAppSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListContainerAppSecretsResult, error) {
			args := v.(ListContainerAppSecretsArgs)
			r, err := ListContainerAppSecrets(ctx, &args, opts...)
			var s ListContainerAppSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListContainerAppSecretsResultOutput)
}

type ListContainerAppSecretsOutputArgs struct {
	ContainerAppName  pulumi.StringInput `pulumi:"containerAppName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListContainerAppSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppSecretsArgs)(nil)).Elem()
}


type ListContainerAppSecretsResultOutput struct{ *pulumi.OutputState }

func (ListContainerAppSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppSecretsResult)(nil)).Elem()
}

func (o ListContainerAppSecretsResultOutput) ToListContainerAppSecretsResultOutput() ListContainerAppSecretsResultOutput {
	return o
}

func (o ListContainerAppSecretsResultOutput) ToListContainerAppSecretsResultOutputWithContext(ctx context.Context) ListContainerAppSecretsResultOutput {
	return o
}

func (o ListContainerAppSecretsResultOutput) Value() ContainerAppSecretResponseArrayOutput {
	return o.ApplyT(func(v ListContainerAppSecretsResult) []ContainerAppSecretResponse { return v.Value }).(ContainerAppSecretResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListContainerAppSecretsResultOutput{})
}
