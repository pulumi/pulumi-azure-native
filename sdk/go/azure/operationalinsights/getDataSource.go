


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataSource(ctx *pulumi.Context, args *LookupDataSourceArgs, opts ...pulumi.InvokeOption) (*LookupDataSourceResult, error) {
	var rv LookupDataSourceResult
	err := ctx.Invoke("azure-native:operationalinsights:getDataSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataSourceArgs struct {
	DataSourceName    string `pulumi:"dataSourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDataSourceResult struct {
	Etag       *string           `pulumi:"etag"`
	Id         string            `pulumi:"id"`
	Kind       string            `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupDataSourceOutput(ctx *pulumi.Context, args LookupDataSourceOutputArgs, opts ...pulumi.InvokeOption) LookupDataSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataSourceResult, error) {
			args := v.(LookupDataSourceArgs)
			r, err := LookupDataSource(ctx, &args, opts...)
			var s LookupDataSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataSourceResultOutput)
}

type LookupDataSourceOutputArgs struct {
	DataSourceName    pulumi.StringInput `pulumi:"dataSourceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDataSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSourceArgs)(nil)).Elem()
}


type LookupDataSourceResultOutput struct{ *pulumi.OutputState }

func (LookupDataSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSourceResult)(nil)).Elem()
}

func (o LookupDataSourceResultOutput) ToLookupDataSourceResultOutput() LookupDataSourceResultOutput {
	return o
}

func (o LookupDataSourceResultOutput) ToLookupDataSourceResultOutputWithContext(ctx context.Context) LookupDataSourceResultOutput {
	return o
}

func (o LookupDataSourceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDataSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataSourceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDataSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataSourceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDataSourceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupDataSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDataSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataSourceResultOutput{})
}
