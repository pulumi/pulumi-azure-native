


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppMetadataSlot(ctx *pulumi.Context, args *ListWebAppMetadataSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppMetadataSlotResult, error) {
	var rv ListWebAppMetadataSlotResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppMetadataSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppMetadataSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppMetadataSlotResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListWebAppMetadataSlotOutput(ctx *pulumi.Context, args ListWebAppMetadataSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppMetadataSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppMetadataSlotResult, error) {
			args := v.(ListWebAppMetadataSlotArgs)
			r, err := ListWebAppMetadataSlot(ctx, &args, opts...)
			var s ListWebAppMetadataSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppMetadataSlotResultOutput)
}

type ListWebAppMetadataSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppMetadataSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppMetadataSlotArgs)(nil)).Elem()
}


type ListWebAppMetadataSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppMetadataSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppMetadataSlotResult)(nil)).Elem()
}

func (o ListWebAppMetadataSlotResultOutput) ToListWebAppMetadataSlotResultOutput() ListWebAppMetadataSlotResultOutput {
	return o
}

func (o ListWebAppMetadataSlotResultOutput) ToListWebAppMetadataSlotResultOutputWithContext(ctx context.Context) ListWebAppMetadataSlotResultOutput {
	return o
}

func (o ListWebAppMetadataSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppMetadataSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppMetadataSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppMetadataSlotResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListWebAppMetadataSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppMetadataSlotResultOutput{})
}
