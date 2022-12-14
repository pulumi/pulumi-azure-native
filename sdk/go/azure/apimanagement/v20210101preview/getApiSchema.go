


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiSchema(ctx *pulumi.Context, args *LookupApiSchemaArgs, opts ...pulumi.InvokeOption) (*LookupApiSchemaResult, error) {
	var rv LookupApiSchemaResult
	err := ctx.Invoke("azure-native:apimanagement/v20210101preview:getApiSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiSchemaArgs struct {
	ApiId             string `pulumi:"apiId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SchemaId          string `pulumi:"schemaId"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiSchemaResult struct {
	ContentType string `pulumi:"contentType"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
}

func LookupApiSchemaOutput(ctx *pulumi.Context, args LookupApiSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupApiSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiSchemaResult, error) {
			args := v.(LookupApiSchemaArgs)
			r, err := LookupApiSchema(ctx, &args, opts...)
			var s LookupApiSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiSchemaResultOutput)
}

type LookupApiSchemaOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaId          pulumi.StringInput `pulumi:"schemaId"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiSchemaArgs)(nil)).Elem()
}


type LookupApiSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupApiSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiSchemaResult)(nil)).Elem()
}

func (o LookupApiSchemaResultOutput) ToLookupApiSchemaResultOutput() LookupApiSchemaResultOutput {
	return o
}

func (o LookupApiSchemaResultOutput) ToLookupApiSchemaResultOutputWithContext(ctx context.Context) LookupApiSchemaResultOutput {
	return o
}

func (o LookupApiSchemaResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiSchemaResult) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o LookupApiSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiSchemaResultOutput{})
}
