


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchema(ctx *pulumi.Context, args *LookupSchemaArgs, opts ...pulumi.InvokeOption) (*LookupSchemaResult, error) {
	var rv LookupSchemaResult
	err := ctx.Invoke("azure-native:logic/v20160601:getSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SchemaName             string `pulumi:"schemaName"`
}


type LookupSchemaResult struct {
	ChangedTime     string              `pulumi:"changedTime"`
	Content         *string             `pulumi:"content"`
	ContentLink     ContentLinkResponse `pulumi:"contentLink"`
	ContentType     *string             `pulumi:"contentType"`
	CreatedTime     string              `pulumi:"createdTime"`
	DocumentName    *string             `pulumi:"documentName"`
	FileName        *string             `pulumi:"fileName"`
	Id              string              `pulumi:"id"`
	Location        *string             `pulumi:"location"`
	Metadata        interface{}         `pulumi:"metadata"`
	Name            string              `pulumi:"name"`
	SchemaType      string              `pulumi:"schemaType"`
	Tags            map[string]string   `pulumi:"tags"`
	TargetNamespace *string             `pulumi:"targetNamespace"`
	Type            string              `pulumi:"type"`
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
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaName             pulumi.StringInput `pulumi:"schemaName"`
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

func (o LookupSchemaResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v LookupSchemaResult) ContentLinkResponse { return v.ContentLink }).(ContentLinkResponseOutput)
}

func (o LookupSchemaResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) DocumentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.DocumentName }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) FileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.FileName }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSchemaResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.SchemaType }).(pulumi.StringOutput)
}

func (o LookupSchemaResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSchemaResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSchemaResultOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchemaResultOutput{})
}
