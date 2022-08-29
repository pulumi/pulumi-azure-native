


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountSchema(ctx *pulumi.Context, args *LookupIntegrationAccountSchemaArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountSchemaResult, error) {
	var rv LookupIntegrationAccountSchemaResult
	err := ctx.Invoke("azure-native:logic/v20190501:getIntegrationAccountSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountSchemaArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SchemaName             string `pulumi:"schemaName"`
}


type LookupIntegrationAccountSchemaResult struct {
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

func LookupIntegrationAccountSchemaOutput(ctx *pulumi.Context, args LookupIntegrationAccountSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountSchemaResult, error) {
			args := v.(LookupIntegrationAccountSchemaArgs)
			r, err := LookupIntegrationAccountSchema(ctx, &args, opts...)
			var s LookupIntegrationAccountSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountSchemaResultOutput)
}

type LookupIntegrationAccountSchemaOutputArgs struct {
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaName             pulumi.StringInput `pulumi:"schemaName"`
}

func (LookupIntegrationAccountSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountSchemaArgs)(nil)).Elem()
}


type LookupIntegrationAccountSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountSchemaResult)(nil)).Elem()
}

func (o LookupIntegrationAccountSchemaResultOutput) ToLookupIntegrationAccountSchemaResultOutput() LookupIntegrationAccountSchemaResultOutput {
	return o
}

func (o LookupIntegrationAccountSchemaResultOutput) ToLookupIntegrationAccountSchemaResultOutputWithContext(ctx context.Context) LookupIntegrationAccountSchemaResultOutput {
	return o
}

func (o LookupIntegrationAccountSchemaResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) ContentLinkResponse { return v.ContentLink }).(ContentLinkResponseOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) DocumentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) *string { return v.DocumentName }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) FileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) *string { return v.FileName }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) string { return v.SchemaType }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountSchemaResultOutput{})
}
