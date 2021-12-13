


package logic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountSchema struct {
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


func NewIntegrationAccountSchema(ctx *pulumi.Context,
	name string, args *IntegrationAccountSchemaArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountSchema, error) {
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
			Type: pulumi.String("azure-native:logic/v20150801preview:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountSchema"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountSchema
	err := ctx.RegisterResource("azure-native:logic:IntegrationAccountSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountSchemaState, opts ...pulumi.ResourceOption) (*IntegrationAccountSchema, error) {
	var resource IntegrationAccountSchema
	err := ctx.ReadResource("azure-native:logic:IntegrationAccountSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountSchemaState struct {
}

type IntegrationAccountSchemaState struct {
}

func (IntegrationAccountSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSchemaState)(nil)).Elem()
}

type integrationAccountSchemaArgs struct {
	Content                *string           `pulumi:"content"`
	ContentType            *string           `pulumi:"contentType"`
	DocumentName           *string           `pulumi:"documentName"`
	FileName               *string           `pulumi:"fileName"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	Metadata               interface{}       `pulumi:"metadata"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SchemaName             *string           `pulumi:"schemaName"`
	SchemaType             string            `pulumi:"schemaType"`
	Tags                   map[string]string `pulumi:"tags"`
	TargetNamespace        *string           `pulumi:"targetNamespace"`
}


type IntegrationAccountSchemaArgs struct {
	Content                pulumi.StringPtrInput
	ContentType            pulumi.StringPtrInput
	DocumentName           pulumi.StringPtrInput
	FileName               pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	ResourceGroupName      pulumi.StringInput
	SchemaName             pulumi.StringPtrInput
	SchemaType             pulumi.StringInput
	Tags                   pulumi.StringMapInput
	TargetNamespace        pulumi.StringPtrInput
}

func (IntegrationAccountSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSchemaArgs)(nil)).Elem()
}

type IntegrationAccountSchemaInput interface {
	pulumi.Input

	ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput
	ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput
}

func (*IntegrationAccountSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSchema)(nil)).Elem()
}

func (i *IntegrationAccountSchema) ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput {
	return i.ToIntegrationAccountSchemaOutputWithContext(context.Background())
}

func (i *IntegrationAccountSchema) ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSchemaOutput)
}

type IntegrationAccountSchemaOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSchema)(nil)).Elem()
}

func (o IntegrationAccountSchemaOutput) ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput {
	return o
}

func (o IntegrationAccountSchemaOutput) ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountSchemaOutput{})
}
