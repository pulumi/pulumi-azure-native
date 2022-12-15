


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteConnectionStrings(ctx *pulumi.Context, args *ListSiteConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListSiteConnectionStringsResult, error) {
	var rv ListSiteConnectionStringsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteConnectionStringsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteConnectionStringsResult struct {
	Id         *string                                    `pulumi:"id"`
	Kind       *string                                    `pulumi:"kind"`
	Location   string                                     `pulumi:"location"`
	Name       *string                                    `pulumi:"name"`
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	Tags       map[string]string                          `pulumi:"tags"`
	Type       *string                                    `pulumi:"type"`
}

func ListSiteConnectionStringsOutput(ctx *pulumi.Context, args ListSiteConnectionStringsOutputArgs, opts ...pulumi.InvokeOption) ListSiteConnectionStringsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteConnectionStringsResult, error) {
			args := v.(ListSiteConnectionStringsArgs)
			r, err := ListSiteConnectionStrings(ctx, &args, opts...)
			var s ListSiteConnectionStringsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteConnectionStringsResultOutput)
}

type ListSiteConnectionStringsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSiteConnectionStringsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteConnectionStringsArgs)(nil)).Elem()
}


type ListSiteConnectionStringsResultOutput struct{ *pulumi.OutputState }

func (ListSiteConnectionStringsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteConnectionStringsResult)(nil)).Elem()
}

func (o ListSiteConnectionStringsResultOutput) ToListSiteConnectionStringsResultOutput() ListSiteConnectionStringsResultOutput {
	return o
}

func (o ListSiteConnectionStringsResultOutput) ToListSiteConnectionStringsResultOutputWithContext(ctx context.Context) ListSiteConnectionStringsResultOutput {
	return o
}

func (o ListSiteConnectionStringsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteConnectionStringsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteConnectionStringsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteConnectionStringsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteConnectionStringsResultOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsResult) map[string]ConnStringValueTypePairResponse {
		return v.Properties
	}).(ConnStringValueTypePairResponseMapOutput)
}

func (o ListSiteConnectionStringsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteConnectionStringsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteConnectionStringsResultOutput{})
}
