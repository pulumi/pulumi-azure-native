


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGlobalSchema(ctx *pulumi.Context, args *LookupGlobalSchemaArgs, opts ...pulumi.InvokeOption) (*LookupGlobalSchemaResult, error) {
	var rv LookupGlobalSchemaResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getGlobalSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalSchemaArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SchemaId          string `pulumi:"schemaId"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGlobalSchemaResult struct {
	Description *string     `pulumi:"description"`
	Id          string      `pulumi:"id"`
	Name        string      `pulumi:"name"`
	SchemaType  string      `pulumi:"schemaType"`
	Type        string      `pulumi:"type"`
	Value       interface{} `pulumi:"value"`
}

func LookupGlobalSchemaOutput(ctx *pulumi.Context, args LookupGlobalSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalSchemaResult, error) {
			args := v.(LookupGlobalSchemaArgs)
			r, err := LookupGlobalSchema(ctx, &args, opts...)
			var s LookupGlobalSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalSchemaResultOutput)
}

type LookupGlobalSchemaOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaId          pulumi.StringInput `pulumi:"schemaId"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGlobalSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalSchemaArgs)(nil)).Elem()
}


type LookupGlobalSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalSchemaResult)(nil)).Elem()
}

func (o LookupGlobalSchemaResultOutput) ToLookupGlobalSchemaResultOutput() LookupGlobalSchemaResultOutput {
	return o
}

func (o LookupGlobalSchemaResultOutput) ToLookupGlobalSchemaResultOutputWithContext(ctx context.Context) LookupGlobalSchemaResultOutput {
	return o
}

func (o LookupGlobalSchemaResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGlobalSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGlobalSchemaResultOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) string { return v.SchemaType }).(pulumi.StringOutput)
}

func (o LookupGlobalSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupGlobalSchemaResultOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) interface{} { return v.Value }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalSchemaResultOutput{})
}
