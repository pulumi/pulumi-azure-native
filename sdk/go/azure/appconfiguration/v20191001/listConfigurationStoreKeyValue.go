


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConfigurationStoreKeyValue(ctx *pulumi.Context, args *ListConfigurationStoreKeyValueArgs, opts ...pulumi.InvokeOption) (*ListConfigurationStoreKeyValueResult, error) {
	var rv ListConfigurationStoreKeyValueResult
	err := ctx.Invoke("azure-native:appconfiguration/v20191001:listConfigurationStoreKeyValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigurationStoreKeyValueArgs struct {
	ConfigStoreName   string  `pulumi:"configStoreName"`
	Key               string  `pulumi:"key"`
	Label             *string `pulumi:"label"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ListConfigurationStoreKeyValueResult struct {
	ContentType  string            `pulumi:"contentType"`
	ETag         string            `pulumi:"eTag"`
	Key          string            `pulumi:"key"`
	Label        string            `pulumi:"label"`
	LastModified string            `pulumi:"lastModified"`
	Locked       bool              `pulumi:"locked"`
	Tags         map[string]string `pulumi:"tags"`
	Value        string            `pulumi:"value"`
}

func ListConfigurationStoreKeyValueOutput(ctx *pulumi.Context, args ListConfigurationStoreKeyValueOutputArgs, opts ...pulumi.InvokeOption) ListConfigurationStoreKeyValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConfigurationStoreKeyValueResult, error) {
			args := v.(ListConfigurationStoreKeyValueArgs)
			r, err := ListConfigurationStoreKeyValue(ctx, &args, opts...)
			var s ListConfigurationStoreKeyValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConfigurationStoreKeyValueResultOutput)
}

type ListConfigurationStoreKeyValueOutputArgs struct {
	ConfigStoreName   pulumi.StringInput    `pulumi:"configStoreName"`
	Key               pulumi.StringInput    `pulumi:"key"`
	Label             pulumi.StringPtrInput `pulumi:"label"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListConfigurationStoreKeyValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationStoreKeyValueArgs)(nil)).Elem()
}


type ListConfigurationStoreKeyValueResultOutput struct{ *pulumi.OutputState }

func (ListConfigurationStoreKeyValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationStoreKeyValueResult)(nil)).Elem()
}

func (o ListConfigurationStoreKeyValueResultOutput) ToListConfigurationStoreKeyValueResultOutput() ListConfigurationStoreKeyValueResultOutput {
	return o
}

func (o ListConfigurationStoreKeyValueResultOutput) ToListConfigurationStoreKeyValueResultOutputWithContext(ctx context.Context) ListConfigurationStoreKeyValueResultOutput {
	return o
}

func (o ListConfigurationStoreKeyValueResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeyValueResult) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o ListConfigurationStoreKeyValueResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeyValueResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o ListConfigurationStoreKeyValueResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeyValueResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o ListConfigurationStoreKeyValueResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeyValueResult) string { return v.Label }).(pulumi.StringOutput)
}

func (o ListConfigurationStoreKeyValueResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeyValueResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o ListConfigurationStoreKeyValueResultOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeyValueResult) bool { return v.Locked }).(pulumi.BoolOutput)
}

func (o ListConfigurationStoreKeyValueResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeyValueResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListConfigurationStoreKeyValueResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeyValueResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConfigurationStoreKeyValueResultOutput{})
}
