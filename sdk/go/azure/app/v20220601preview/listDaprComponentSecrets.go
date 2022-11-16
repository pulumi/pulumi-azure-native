


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDaprComponentSecrets(ctx *pulumi.Context, args *ListDaprComponentSecretsArgs, opts ...pulumi.InvokeOption) (*ListDaprComponentSecretsResult, error) {
	var rv ListDaprComponentSecretsResult
	err := ctx.Invoke("azure-native:app/v20220601preview:listDaprComponentSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDaprComponentSecretsArgs struct {
	ComponentName     string `pulumi:"componentName"`
	EnvironmentName   string `pulumi:"environmentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDaprComponentSecretsResult struct {
	Value []DaprSecretResponse `pulumi:"value"`
}

func ListDaprComponentSecretsOutput(ctx *pulumi.Context, args ListDaprComponentSecretsOutputArgs, opts ...pulumi.InvokeOption) ListDaprComponentSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDaprComponentSecretsResult, error) {
			args := v.(ListDaprComponentSecretsArgs)
			r, err := ListDaprComponentSecrets(ctx, &args, opts...)
			var s ListDaprComponentSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDaprComponentSecretsResultOutput)
}

type ListDaprComponentSecretsOutputArgs struct {
	ComponentName     pulumi.StringInput `pulumi:"componentName"`
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDaprComponentSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDaprComponentSecretsArgs)(nil)).Elem()
}


type ListDaprComponentSecretsResultOutput struct{ *pulumi.OutputState }

func (ListDaprComponentSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDaprComponentSecretsResult)(nil)).Elem()
}

func (o ListDaprComponentSecretsResultOutput) ToListDaprComponentSecretsResultOutput() ListDaprComponentSecretsResultOutput {
	return o
}

func (o ListDaprComponentSecretsResultOutput) ToListDaprComponentSecretsResultOutputWithContext(ctx context.Context) ListDaprComponentSecretsResultOutput {
	return o
}

func (o ListDaprComponentSecretsResultOutput) Value() DaprSecretResponseArrayOutput {
	return o.ApplyT(func(v ListDaprComponentSecretsResult) []DaprSecretResponse { return v.Value }).(DaprSecretResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDaprComponentSecretsResultOutput{})
}
