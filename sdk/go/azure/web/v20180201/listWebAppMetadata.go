


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppMetadata(ctx *pulumi.Context, args *ListWebAppMetadataArgs, opts ...pulumi.InvokeOption) (*ListWebAppMetadataResult, error) {
	var rv ListWebAppMetadataResult
	err := ctx.Invoke("azure-native:web/v20180201:listWebAppMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppMetadataArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppMetadataResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListWebAppMetadataOutput(ctx *pulumi.Context, args ListWebAppMetadataOutputArgs, opts ...pulumi.InvokeOption) ListWebAppMetadataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppMetadataResult, error) {
			args := v.(ListWebAppMetadataArgs)
			r, err := ListWebAppMetadata(ctx, &args, opts...)
			var s ListWebAppMetadataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppMetadataResultOutput)
}

type ListWebAppMetadataOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppMetadataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppMetadataArgs)(nil)).Elem()
}


type ListWebAppMetadataResultOutput struct{ *pulumi.OutputState }

func (ListWebAppMetadataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppMetadataResult)(nil)).Elem()
}

func (o ListWebAppMetadataResultOutput) ToListWebAppMetadataResultOutput() ListWebAppMetadataResultOutput {
	return o
}

func (o ListWebAppMetadataResultOutput) ToListWebAppMetadataResultOutputWithContext(ctx context.Context) ListWebAppMetadataResultOutput {
	return o
}

func (o ListWebAppMetadataResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppMetadataResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppMetadataResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppMetadataResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppMetadataResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppMetadataResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListWebAppMetadataResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppMetadataResultOutput{})
}
