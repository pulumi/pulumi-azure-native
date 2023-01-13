


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIngestionSettingConnectionStrings(ctx *pulumi.Context, args *ListIngestionSettingConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListIngestionSettingConnectionStringsResult, error) {
	var rv ListIngestionSettingConnectionStringsResult
	err := ctx.Invoke("azure-native:security:listIngestionSettingConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIngestionSettingConnectionStringsArgs struct {
	IngestionSettingName string `pulumi:"ingestionSettingName"`
}


type ListIngestionSettingConnectionStringsResult struct {
	Value []IngestionConnectionStringResponse `pulumi:"value"`
}

func ListIngestionSettingConnectionStringsOutput(ctx *pulumi.Context, args ListIngestionSettingConnectionStringsOutputArgs, opts ...pulumi.InvokeOption) ListIngestionSettingConnectionStringsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIngestionSettingConnectionStringsResult, error) {
			args := v.(ListIngestionSettingConnectionStringsArgs)
			r, err := ListIngestionSettingConnectionStrings(ctx, &args, opts...)
			var s ListIngestionSettingConnectionStringsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIngestionSettingConnectionStringsResultOutput)
}

type ListIngestionSettingConnectionStringsOutputArgs struct {
	IngestionSettingName pulumi.StringInput `pulumi:"ingestionSettingName"`
}

func (ListIngestionSettingConnectionStringsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIngestionSettingConnectionStringsArgs)(nil)).Elem()
}


type ListIngestionSettingConnectionStringsResultOutput struct{ *pulumi.OutputState }

func (ListIngestionSettingConnectionStringsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIngestionSettingConnectionStringsResult)(nil)).Elem()
}

func (o ListIngestionSettingConnectionStringsResultOutput) ToListIngestionSettingConnectionStringsResultOutput() ListIngestionSettingConnectionStringsResultOutput {
	return o
}

func (o ListIngestionSettingConnectionStringsResultOutput) ToListIngestionSettingConnectionStringsResultOutputWithContext(ctx context.Context) ListIngestionSettingConnectionStringsResultOutput {
	return o
}

func (o ListIngestionSettingConnectionStringsResultOutput) Value() IngestionConnectionStringResponseArrayOutput {
	return o.ApplyT(func(v ListIngestionSettingConnectionStringsResult) []IngestionConnectionStringResponse {
		return v.Value
	}).(IngestionConnectionStringResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIngestionSettingConnectionStringsResultOutput{})
}
