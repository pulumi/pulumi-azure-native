


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchema(ctx *pulumi.Context, args *LookupSchemaArgs, opts ...pulumi.InvokeOption) (*LookupSchemaResult, error) {
	var rv LookupSchemaResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:getSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SchemaId          string `pulumi:"schemaId"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupSchemaResult struct {
	Description *string `pulumi:"description"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SchemaType  string  `pulumi:"schemaType"`
	Type        string  `pulumi:"type"`
	Value       *string `pulumi:"value"`
}

func LookupSchemaOutput(ctx *pulumi.Context, args LookupSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSchemaResult, error) {
			args := v.(LookupSchemaArgs)
			r, err := LookupSchema(ctx, &args, opts...)
			var s LookupSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSchemaResultOutput)
}

type LookupSchemaOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaId          pulumi.StringInput `pulumi:"schemaId"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaArgs)(nil)).Elem()
}


type LookupSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaResult)(nil)).Elem()
}

func (o LookupSchemaResultOutput) ToLookupSchemaResultOutput() LookupSchemaResultOutput {
	return o
}

func (o LookupSchemaResultOutput) ToLookupSchemaResultOutputWithContext(ctx context.Context) LookupSchemaResultOutput {
	return o
}

func (o LookupSchemaResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.SchemaType }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchemaResultOutput{})
}
