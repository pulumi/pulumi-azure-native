


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectedEnvironmentsDaprComponentSecrets(ctx *pulumi.Context, args *ListConnectedEnvironmentsDaprComponentSecretsArgs, opts ...pulumi.InvokeOption) (*ListConnectedEnvironmentsDaprComponentSecretsResult, error) {
	var rv ListConnectedEnvironmentsDaprComponentSecretsResult
	err := ctx.Invoke("azure-native:app/v20220601preview:listConnectedEnvironmentsDaprComponentSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedEnvironmentsDaprComponentSecretsArgs struct {
	ComponentName            string `pulumi:"componentName"`
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type ListConnectedEnvironmentsDaprComponentSecretsResult struct {
	Value []SecretResponse `pulumi:"value"`
}

func ListConnectedEnvironmentsDaprComponentSecretsOutput(ctx *pulumi.Context, args ListConnectedEnvironmentsDaprComponentSecretsOutputArgs, opts ...pulumi.InvokeOption) ListConnectedEnvironmentsDaprComponentSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConnectedEnvironmentsDaprComponentSecretsResult, error) {
			args := v.(ListConnectedEnvironmentsDaprComponentSecretsArgs)
			r, err := ListConnectedEnvironmentsDaprComponentSecrets(ctx, &args, opts...)
			var s ListConnectedEnvironmentsDaprComponentSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConnectedEnvironmentsDaprComponentSecretsResultOutput)
}

type ListConnectedEnvironmentsDaprComponentSecretsOutputArgs struct {
	ComponentName            pulumi.StringInput `pulumi:"componentName"`
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListConnectedEnvironmentsDaprComponentSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedEnvironmentsDaprComponentSecretsArgs)(nil)).Elem()
}


type ListConnectedEnvironmentsDaprComponentSecretsResultOutput struct{ *pulumi.OutputState }

func (ListConnectedEnvironmentsDaprComponentSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedEnvironmentsDaprComponentSecretsResult)(nil)).Elem()
}

func (o ListConnectedEnvironmentsDaprComponentSecretsResultOutput) ToListConnectedEnvironmentsDaprComponentSecretsResultOutput() ListConnectedEnvironmentsDaprComponentSecretsResultOutput {
	return o
}

func (o ListConnectedEnvironmentsDaprComponentSecretsResultOutput) ToListConnectedEnvironmentsDaprComponentSecretsResultOutputWithContext(ctx context.Context) ListConnectedEnvironmentsDaprComponentSecretsResultOutput {
	return o
}

func (o ListConnectedEnvironmentsDaprComponentSecretsResultOutput) Value() SecretResponseArrayOutput {
	return o.ApplyT(func(v ListConnectedEnvironmentsDaprComponentSecretsResult) []SecretResponse { return v.Value }).(SecretResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectedEnvironmentsDaprComponentSecretsResultOutput{})
}
