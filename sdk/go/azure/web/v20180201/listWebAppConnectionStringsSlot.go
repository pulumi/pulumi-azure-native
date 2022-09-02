


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppConnectionStringsSlot(ctx *pulumi.Context, args *ListWebAppConnectionStringsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppConnectionStringsSlotResult, error) {
	var rv ListWebAppConnectionStringsSlotResult
	err := ctx.Invoke("azure-native:web/v20180201:listWebAppConnectionStringsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppConnectionStringsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppConnectionStringsSlotResult struct {
	Id         string                                     `pulumi:"id"`
	Kind       *string                                    `pulumi:"kind"`
	Name       string                                     `pulumi:"name"`
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}

func ListWebAppConnectionStringsSlotOutput(ctx *pulumi.Context, args ListWebAppConnectionStringsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppConnectionStringsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppConnectionStringsSlotResult, error) {
			args := v.(ListWebAppConnectionStringsSlotArgs)
			r, err := ListWebAppConnectionStringsSlot(ctx, &args, opts...)
			var s ListWebAppConnectionStringsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppConnectionStringsSlotResultOutput)
}

type ListWebAppConnectionStringsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppConnectionStringsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppConnectionStringsSlotArgs)(nil)).Elem()
}


type ListWebAppConnectionStringsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppConnectionStringsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppConnectionStringsSlotResult)(nil)).Elem()
}

func (o ListWebAppConnectionStringsSlotResultOutput) ToListWebAppConnectionStringsSlotResultOutput() ListWebAppConnectionStringsSlotResultOutput {
	return o
}

func (o ListWebAppConnectionStringsSlotResultOutput) ToListWebAppConnectionStringsSlotResultOutputWithContext(ctx context.Context) ListWebAppConnectionStringsSlotResultOutput {
	return o
}

func (o ListWebAppConnectionStringsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppConnectionStringsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppConnectionStringsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppConnectionStringsSlotResultOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) map[string]ConnStringValueTypePairResponse {
		return v.Properties
	}).(ConnStringValueTypePairResponseMapOutput)
}

func (o ListWebAppConnectionStringsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppConnectionStringsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppConnectionStringsSlotResultOutput{})
}
