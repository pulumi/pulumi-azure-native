


package v20220808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBotSecrets(ctx *pulumi.Context, args *ListBotSecretsArgs, opts ...pulumi.InvokeOption) (*ListBotSecretsResult, error) {
	var rv ListBotSecretsResult
	err := ctx.Invoke("azure-native:healthbot/v20220808:listBotSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBotSecretsArgs struct {
	BotName           string `pulumi:"botName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListBotSecretsResult struct {
	Secrets []HealthBotKeyResponse `pulumi:"secrets"`
}

func ListBotSecretsOutput(ctx *pulumi.Context, args ListBotSecretsOutputArgs, opts ...pulumi.InvokeOption) ListBotSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBotSecretsResult, error) {
			args := v.(ListBotSecretsArgs)
			r, err := ListBotSecrets(ctx, &args, opts...)
			var s ListBotSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBotSecretsResultOutput)
}

type ListBotSecretsOutputArgs struct {
	BotName           pulumi.StringInput `pulumi:"botName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListBotSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBotSecretsArgs)(nil)).Elem()
}


type ListBotSecretsResultOutput struct{ *pulumi.OutputState }

func (ListBotSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBotSecretsResult)(nil)).Elem()
}

func (o ListBotSecretsResultOutput) ToListBotSecretsResultOutput() ListBotSecretsResultOutput {
	return o
}

func (o ListBotSecretsResultOutput) ToListBotSecretsResultOutputWithContext(ctx context.Context) ListBotSecretsResultOutput {
	return o
}

func (o ListBotSecretsResultOutput) Secrets() HealthBotKeyResponseArrayOutput {
	return o.ApplyT(func(v ListBotSecretsResult) []HealthBotKeyResponse { return v.Secrets }).(HealthBotKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBotSecretsResultOutput{})
}
