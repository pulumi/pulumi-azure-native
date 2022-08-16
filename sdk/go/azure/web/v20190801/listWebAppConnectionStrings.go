


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppConnectionStrings(ctx *pulumi.Context, args *ListWebAppConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListWebAppConnectionStringsResult, error) {
	var rv ListWebAppConnectionStringsResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppConnectionStringsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppConnectionStringsResult struct {
	Id         string                                     `pulumi:"id"`
	Kind       *string                                    `pulumi:"kind"`
	Name       string                                     `pulumi:"name"`
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}

func ListWebAppConnectionStringsOutput(ctx *pulumi.Context, args ListWebAppConnectionStringsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppConnectionStringsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppConnectionStringsResult, error) {
			args := v.(ListWebAppConnectionStringsArgs)
			r, err := ListWebAppConnectionStrings(ctx, &args, opts...)
			var s ListWebAppConnectionStringsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppConnectionStringsResultOutput)
}

type ListWebAppConnectionStringsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppConnectionStringsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppConnectionStringsArgs)(nil)).Elem()
}


type ListWebAppConnectionStringsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppConnectionStringsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppConnectionStringsResult)(nil)).Elem()
}

func (o ListWebAppConnectionStringsResultOutput) ToListWebAppConnectionStringsResultOutput() ListWebAppConnectionStringsResultOutput {
	return o
}

func (o ListWebAppConnectionStringsResultOutput) ToListWebAppConnectionStringsResultOutputWithContext(ctx context.Context) ListWebAppConnectionStringsResultOutput {
	return o
}

func (o ListWebAppConnectionStringsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppConnectionStringsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppConnectionStringsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppConnectionStringsResultOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) map[string]ConnStringValueTypePairResponse {
		return v.Properties
	}).(ConnStringValueTypePairResponseMapOutput)
}

func (o ListWebAppConnectionStringsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppConnectionStringsResultOutput{})
}
