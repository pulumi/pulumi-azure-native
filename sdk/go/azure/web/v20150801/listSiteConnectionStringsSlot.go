


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteConnectionStringsSlot(ctx *pulumi.Context, args *ListSiteConnectionStringsSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteConnectionStringsSlotResult, error) {
	var rv ListSiteConnectionStringsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteConnectionStringsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteConnectionStringsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSiteConnectionStringsSlotResult struct {
	Id         *string                                    `pulumi:"id"`
	Kind       *string                                    `pulumi:"kind"`
	Location   string                                     `pulumi:"location"`
	Name       *string                                    `pulumi:"name"`
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	Tags       map[string]string                          `pulumi:"tags"`
	Type       *string                                    `pulumi:"type"`
}

func ListSiteConnectionStringsSlotOutput(ctx *pulumi.Context, args ListSiteConnectionStringsSlotOutputArgs, opts ...pulumi.InvokeOption) ListSiteConnectionStringsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteConnectionStringsSlotResult, error) {
			args := v.(ListSiteConnectionStringsSlotArgs)
			r, err := ListSiteConnectionStringsSlot(ctx, &args, opts...)
			var s ListSiteConnectionStringsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteConnectionStringsSlotResultOutput)
}

type ListSiteConnectionStringsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListSiteConnectionStringsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteConnectionStringsSlotArgs)(nil)).Elem()
}


type ListSiteConnectionStringsSlotResultOutput struct{ *pulumi.OutputState }

func (ListSiteConnectionStringsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteConnectionStringsSlotResult)(nil)).Elem()
}

func (o ListSiteConnectionStringsSlotResultOutput) ToListSiteConnectionStringsSlotResultOutput() ListSiteConnectionStringsSlotResultOutput {
	return o
}

func (o ListSiteConnectionStringsSlotResultOutput) ToListSiteConnectionStringsSlotResultOutputWithContext(ctx context.Context) ListSiteConnectionStringsSlotResultOutput {
	return o
}

func (o ListSiteConnectionStringsSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteConnectionStringsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteConnectionStringsSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteConnectionStringsSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteConnectionStringsSlotResultOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) map[string]ConnStringValueTypePairResponse {
		return v.Properties
	}).(ConnStringValueTypePairResponseMapOutput)
}

func (o ListSiteConnectionStringsSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteConnectionStringsSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteConnectionStringsSlotResultOutput{})
}
