


package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TableServiceProperties struct {
	pulumi.CustomResourceState

	Cors CorsRulesResponsePtrOutput `pulumi:"cors"`
	Name pulumi.StringOutput        `pulumi:"name"`
	Type pulumi.StringOutput        `pulumi:"type"`
}


func NewTableServiceProperties(ctx *pulumi.Context,
	name string, args *TableServicePropertiesArgs, opts ...pulumi.ResourceOption) (*TableServiceProperties, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage/v20190601:TableServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:TableServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:TableServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:TableServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:TableServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:TableServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:TableServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:TableServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:TableServiceProperties"),
		},
	})
	opts = append(opts, aliases)
	var resource TableServiceProperties
	err := ctx.RegisterResource("azure-native:storage:TableServiceProperties", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTableServiceProperties(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableServicePropertiesState, opts ...pulumi.ResourceOption) (*TableServiceProperties, error) {
	var resource TableServiceProperties
	err := ctx.ReadResource("azure-native:storage:TableServiceProperties", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tableServicePropertiesState struct {
}

type TableServicePropertiesState struct {
}

func (TableServicePropertiesState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableServicePropertiesState)(nil)).Elem()
}

type tableServicePropertiesArgs struct {
	AccountName       string     `pulumi:"accountName"`
	Cors              *CorsRules `pulumi:"cors"`
	ResourceGroupName string     `pulumi:"resourceGroupName"`
	TableServiceName  *string    `pulumi:"tableServiceName"`
}


type TableServicePropertiesArgs struct {
	AccountName       pulumi.StringInput
	Cors              CorsRulesPtrInput
	ResourceGroupName pulumi.StringInput
	TableServiceName  pulumi.StringPtrInput
}

func (TableServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableServicePropertiesArgs)(nil)).Elem()
}

type TableServicePropertiesInput interface {
	pulumi.Input

	ToTableServicePropertiesOutput() TableServicePropertiesOutput
	ToTableServicePropertiesOutputWithContext(ctx context.Context) TableServicePropertiesOutput
}

func (*TableServiceProperties) ElementType() reflect.Type {
	return reflect.TypeOf((**TableServiceProperties)(nil)).Elem()
}

func (i *TableServiceProperties) ToTableServicePropertiesOutput() TableServicePropertiesOutput {
	return i.ToTableServicePropertiesOutputWithContext(context.Background())
}

func (i *TableServiceProperties) ToTableServicePropertiesOutputWithContext(ctx context.Context) TableServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableServicePropertiesOutput)
}

type TableServicePropertiesOutput struct{ *pulumi.OutputState }

func (TableServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableServiceProperties)(nil)).Elem()
}

func (o TableServicePropertiesOutput) ToTableServicePropertiesOutput() TableServicePropertiesOutput {
	return o
}

func (o TableServicePropertiesOutput) ToTableServicePropertiesOutputWithContext(ctx context.Context) TableServicePropertiesOutput {
	return o
}

func (o TableServicePropertiesOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v *TableServiceProperties) CorsRulesResponsePtrOutput { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o TableServicePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TableServiceProperties) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TableServicePropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TableServiceProperties) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TableServicePropertiesOutput{})
}
