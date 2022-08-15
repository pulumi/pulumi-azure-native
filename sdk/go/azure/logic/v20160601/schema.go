


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schema struct {
	pulumi.CustomResourceState

	ChangedTime     pulumi.StringOutput       `pulumi:"changedTime"`
	Content         pulumi.StringPtrOutput    `pulumi:"content"`
	ContentLink     ContentLinkResponseOutput `pulumi:"contentLink"`
	ContentType     pulumi.StringPtrOutput    `pulumi:"contentType"`
	CreatedTime     pulumi.StringOutput       `pulumi:"createdTime"`
	DocumentName    pulumi.StringPtrOutput    `pulumi:"documentName"`
	FileName        pulumi.StringPtrOutput    `pulumi:"fileName"`
	Location        pulumi.StringPtrOutput    `pulumi:"location"`
	Metadata        pulumi.AnyOutput          `pulumi:"metadata"`
	Name            pulumi.StringOutput       `pulumi:"name"`
	SchemaType      pulumi.StringOutput       `pulumi:"schemaType"`
	Tags            pulumi.StringMapOutput    `pulumi:"tags"`
	TargetNamespace pulumi.StringPtrOutput    `pulumi:"targetNamespace"`
	Type            pulumi.StringOutput       `pulumi:"type"`
}


func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaType == nil {
		return nil, errors.New("invalid value for required argument 'SchemaType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:Schema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:Schema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Schema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Schema"),
		},
	})
	opts = append(opts, aliases)
	var resource Schema
	err := ctx.RegisterResource("azure-native:logic/v20160601:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("azure-native:logic/v20160601:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type schemaState struct {
}

type SchemaState struct {
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	Content                *string           `pulumi:"content"`
	ContentType            *string           `pulumi:"contentType"`
	DocumentName           *string           `pulumi:"documentName"`
	FileName               *string           `pulumi:"fileName"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	Metadata               interface{}       `pulumi:"metadata"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SchemaName             *string           `pulumi:"schemaName"`
	SchemaType             SchemaType        `pulumi:"schemaType"`
	Tags                   map[string]string `pulumi:"tags"`
	TargetNamespace        *string           `pulumi:"targetNamespace"`
}


type SchemaArgs struct {
	Content                pulumi.StringPtrInput
	ContentType            pulumi.StringPtrInput
	DocumentName           pulumi.StringPtrInput
	FileName               pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	ResourceGroupName      pulumi.StringInput
	SchemaName             pulumi.StringPtrInput
	SchemaType             SchemaTypeInput
	Tags                   pulumi.StringMapInput
	TargetNamespace        pulumi.StringPtrInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

func (o SchemaOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o SchemaOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v *Schema) ContentLinkResponseOutput { return v.ContentLink }).(ContentLinkResponseOutput)
}

func (o SchemaOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o SchemaOutput) DocumentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.DocumentName }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) FileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.FileName }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *Schema) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o SchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SchemaOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.SchemaType }).(pulumi.StringOutput)
}

func (o SchemaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SchemaOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SchemaOutput{})
}
