


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIngestionSettingTokens(ctx *pulumi.Context, args *ListIngestionSettingTokensArgs, opts ...pulumi.InvokeOption) (*ListIngestionSettingTokensResult, error) {
	var rv ListIngestionSettingTokensResult
	err := ctx.Invoke("azure-native:security:listIngestionSettingTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIngestionSettingTokensArgs struct {
	IngestionSettingName string `pulumi:"ingestionSettingName"`
}


type ListIngestionSettingTokensResult struct {
	Token string `pulumi:"token"`
}

func ListIngestionSettingTokensOutput(ctx *pulumi.Context, args ListIngestionSettingTokensOutputArgs, opts ...pulumi.InvokeOption) ListIngestionSettingTokensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIngestionSettingTokensResult, error) {
			args := v.(ListIngestionSettingTokensArgs)
			r, err := ListIngestionSettingTokens(ctx, &args, opts...)
			var s ListIngestionSettingTokensResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIngestionSettingTokensResultOutput)
}

type ListIngestionSettingTokensOutputArgs struct {
	IngestionSettingName pulumi.StringInput `pulumi:"ingestionSettingName"`
}

func (ListIngestionSettingTokensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIngestionSettingTokensArgs)(nil)).Elem()
}


type ListIngestionSettingTokensResultOutput struct{ *pulumi.OutputState }

func (ListIngestionSettingTokensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIngestionSettingTokensResult)(nil)).Elem()
}

func (o ListIngestionSettingTokensResultOutput) ToListIngestionSettingTokensResultOutput() ListIngestionSettingTokensResultOutput {
	return o
}

func (o ListIngestionSettingTokensResultOutput) ToListIngestionSettingTokensResultOutputWithContext(ctx context.Context) ListIngestionSettingTokensResultOutput {
	return o
}

func (o ListIngestionSettingTokensResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v ListIngestionSettingTokensResult) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIngestionSettingTokensResultOutput{})
}
