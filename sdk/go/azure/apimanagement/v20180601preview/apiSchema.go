


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiSchema struct {
	pulumi.CustomResourceState

	ContentType pulumi.StringOutput    `pulumi:"contentType"`
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
			Type: pulumi.String("azure-native:apimanagement:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiSchema"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiSchema
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:ApiSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiSchemaState, opts ...pulumi.ResourceOption) (*ApiSchema, error) {
	var resource ApiSchema
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:ApiSchema", name, id, state, &resource, opts...)
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
	ApiId             string  `pulumi:"apiId"`
	ContentType       string  `pulumi:"contentType"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SchemaId          *string `pulumi:"schemaId"`
	ServiceName       string  `pulumi:"serviceName"`
	Value             *string `pulumi:"value"`
}


type ApiSchemaArgs struct {
	ApiId             pulumi.StringInput
	ContentType       pulumi.StringInput
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
	return reflect.TypeOf((**ApiSchema)(nil)).Elem()
}

func (i *ApiSchema) ToApiSchemaOutput() ApiSchemaOutput {
	return i.ToApiSchemaOutputWithContext(context.Background())
}

func (i *ApiSchema) ToApiSchemaOutputWithContext(ctx context.Context) ApiSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiSchemaOutput)
}

type ApiSchemaOutput struct{ *pulumi.OutputState }

func (ApiSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiSchema)(nil)).Elem()
}

func (o ApiSchemaOutput) ToApiSchemaOutput() ApiSchemaOutput {
	return o
}

func (o ApiSchemaOutput) ToApiSchemaOutputWithContext(ctx context.Context) ApiSchemaOutput {
	return o
}

func (o ApiSchemaOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiSchema) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

func (o ApiSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiSchema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiSchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiSchema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApiSchemaOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiSchema) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiSchemaOutput{})
}
