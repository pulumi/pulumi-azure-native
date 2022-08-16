


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteMetadataSlot(ctx *pulumi.Context, args *ListSiteMetadataSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteMetadataSlotResult, error) {
	var rv ListSiteMetadataSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteMetadataSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteMetadataSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSiteMetadataSlotResult struct {
	Id         *string           `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   string            `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}

func ListSiteMetadataSlotOutput(ctx *pulumi.Context, args ListSiteMetadataSlotOutputArgs, opts ...pulumi.InvokeOption) ListSiteMetadataSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteMetadataSlotResult, error) {
			args := v.(ListSiteMetadataSlotArgs)
			r, err := ListSiteMetadataSlot(ctx, &args, opts...)
			var s ListSiteMetadataSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteMetadataSlotResultOutput)
}

type ListSiteMetadataSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListSiteMetadataSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteMetadataSlotArgs)(nil)).Elem()
}


type ListSiteMetadataSlotResultOutput struct{ *pulumi.OutputState }

func (ListSiteMetadataSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteMetadataSlotResult)(nil)).Elem()
}

func (o ListSiteMetadataSlotResultOutput) ToListSiteMetadataSlotResultOutput() ListSiteMetadataSlotResultOutput {
	return o
}

func (o ListSiteMetadataSlotResultOutput) ToListSiteMetadataSlotResultOutputWithContext(ctx context.Context) ListSiteMetadataSlotResultOutput {
	return o
}

func (o ListSiteMetadataSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteMetadataSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteMetadataSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteMetadataSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteMetadataSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteMetadataSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteMetadataSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteMetadataSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteMetadataSlotResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteMetadataSlotResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListSiteMetadataSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteMetadataSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteMetadataSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteMetadataSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteMetadataSlotResultOutput{})
}
