


package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiSchema struct {
	pulumi.CustomResourceState

	ContentType pulumi.StringOutput    `pulumi:"contentType"`
	Definitions pulumi.AnyOutput       `pulumi:"definitions"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	Value       pulumi.StringPtrOutput `pulumi:"value"`
}


func NewApiSchema(ctx *pulumi.Context,
	name string, args *ApiSchemaArgs, opts ...pulumi.ResourceOption) (*ApiSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ContentType == nil {
		return nil, errors.New("invalid value for required argument 'ContentType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:ApiSchema"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiSchema
	err := ctx.RegisterResource("azure-native:apimanagement:ApiSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiSchemaState, opts ...pulumi.ResourceOption) (*ApiSchema, error) {
	var resource ApiSchema
	err := ctx.ReadResource("azure-native:apimanagement:ApiSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiSchemaState struct {
}

type ApiSchemaState struct {
}

func (ApiSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiSchemaState)(nil)).Elem()
}

type apiSchemaArgs struct {
	ApiId             string      `pulumi:"apiId"`
	ContentType       string      `pulumi:"contentType"`
	Definitions       interface{} `pulumi:"definitions"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	SchemaId          *string     `pulumi:"schemaId"`
	ServiceName       string      `pulumi:"serviceName"`
	Value             *string     `pulumi:"value"`
}


type ApiSchemaArgs struct {
	ApiId             pulumi.StringInput
	ContentType       pulumi.StringInput
	Definitions       pulumi.Input
	ResourceGroupName pulumi.StringInput
	SchemaId          pulumi.StringPtrInput
	ServiceName       pulumi.StringInput
	Value             pulumi.StringPtrInput
}

func (ApiSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiSchemaArgs)(nil)).Elem()
}

type ApiSchemaInput interface {
	pulumi.Input

	ToApiSchemaOutput() ApiSchemaOutput
	ToApiSchemaOutputWithContext(ctx context.Context) ApiSchemaOutput
}

func (*ApiSchema) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiSchema)(nil))
}

func (i *ApiSchema) ToApiSchemaOutput() ApiSchemaOutput {
	return i.ToApiSchemaOutputWithContext(context.Background())
}

func (i *ApiSchema) ToApiSchemaOutputWithContext(ctx context.Context) ApiSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiSchemaOutput)
}

type ApiSchemaOutput struct{ *pulumi.OutputState }

func (ApiSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiSchema)(nil))
}

func (o ApiSchemaOutput) ToApiSchemaOutput() ApiSchemaOutput {
	return o
}

func (o ApiSchemaOutput) ToApiSchemaOutputWithContext(ctx context.Context) ApiSchemaOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiSchemaOutput{})
}
