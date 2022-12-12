


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteMetadata(ctx *pulumi.Context, args *ListSiteMetadataArgs, opts ...pulumi.InvokeOption) (*ListSiteMetadataResult, error) {
	var rv ListSiteMetadataResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteMetadataArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteMetadataResult struct {
	Id         *string           `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   string            `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}

func ListSiteMetadataOutput(ctx *pulumi.Context, args ListSiteMetadataOutputArgs, opts ...pulumi.InvokeOption) ListSiteMetadataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteMetadataResult, error) {
			args := v.(ListSiteMetadataArgs)
			r, err := ListSiteMetadata(ctx, &args, opts...)
			var s ListSiteMetadataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteMetadataResultOutput)
}

type ListSiteMetadataOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSiteMetadataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteMetadataArgs)(nil)).Elem()
}


type ListSiteMetadataResultOutput struct{ *pulumi.OutputState }

func (ListSiteMetadataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteMetadataResult)(nil)).Elem()
}

func (o ListSiteMetadataResultOutput) ToListSiteMetadataResultOutput() ListSiteMetadataResultOutput {
	return o
}

func (o ListSiteMetadataResultOutput) ToListSiteMetadataResultOutputWithContext(ctx context.Context) ListSiteMetadataResultOutput {
	return o
}

func (o ListSiteMetadataResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteMetadataResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteMetadataResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteMetadataResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteMetadataResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteMetadataResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteMetadataResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteMetadataResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteMetadataResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteMetadataResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListSiteMetadataResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteMetadataResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteMetadataResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteMetadataResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteMetadataResultOutput{})
}
