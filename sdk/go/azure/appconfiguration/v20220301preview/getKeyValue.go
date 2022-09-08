


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKeyValue(ctx *pulumi.Context, args *LookupKeyValueArgs, opts ...pulumi.InvokeOption) (*LookupKeyValueResult, error) {
	var rv LookupKeyValueResult
	err := ctx.Invoke("azure-native:appconfiguration/v20220301preview:getKeyValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyValueArgs struct {
	ConfigStoreName   string `pulumi:"configStoreName"`
	KeyValueName      string `pulumi:"keyValueName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupKeyValueResult struct {
	ContentType  *string           `pulumi:"contentType"`
	ETag         string            `pulumi:"eTag"`
	Id           string            `pulumi:"id"`
	Key          string            `pulumi:"key"`
	Label        string            `pulumi:"label"`
	LastModified string            `pulumi:"lastModified"`
	Locked       bool              `pulumi:"locked"`
	Name         string            `pulumi:"name"`
	Tags         map[string]string `pulumi:"tags"`
	Type         string            `pulumi:"type"`
	Value        *string           `pulumi:"value"`
}

func LookupKeyValueOutput(ctx *pulumi.Context, args LookupKeyValueOutputArgs, opts ...pulumi.InvokeOption) LookupKeyValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKeyValueResult, error) {
			args := v.(LookupKeyValueArgs)
			r, err := LookupKeyValue(ctx, &args, opts...)
			var s LookupKeyValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKeyValueResultOutput)
}

type LookupKeyValueOutputArgs struct {
	ConfigStoreName   pulumi.StringInput `pulumi:"configStoreName"`
	KeyValueName      pulumi.StringInput `pulumi:"keyValueName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKeyValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyValueArgs)(nil)).Elem()
}


type LookupKeyValueResultOutput struct{ *pulumi.OutputState }

func (LookupKeyValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyValueResult)(nil)).Elem()
}

func (o LookupKeyValueResultOutput) ToLookupKeyValueResultOutput() LookupKeyValueResultOutput {
	return o
}

func (o LookupKeyValueResultOutput) ToLookupKeyValueResultOutputWithContext(ctx context.Context) LookupKeyValueResultOutput {
	return o
}

func (o LookupKeyValueResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyValueResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o LookupKeyValueResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyValueResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupKeyValueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyValueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKeyValueResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyValueResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupKeyValueResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyValueResult) string { return v.Label }).(pulumi.StringOutput)
}

func (o LookupKeyValueResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyValueResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupKeyValueResultOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupKeyValueResult) bool { return v.Locked }).(pulumi.BoolOutput)
}

func (o LookupKeyValueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyValueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKeyValueResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKeyValueResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupKeyValueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyValueResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupKeyValueResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyValueResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeyValueResultOutput{})
}
