


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Table struct {
	pulumi.CustomResourceState

	Name      pulumi.StringOutput `pulumi:"name"`
	TableName pulumi.StringOutput `pulumi:"tableName"`
	Type      pulumi.StringOutput `pulumi:"type"`
}


func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
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
			Type: pulumi.String("azure-native:storage:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:Table"),
		},
	})
	opts = append(opts, aliases)
	var resource Table
	err := ctx.RegisterResource("azure-native:storage/v20210201:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("azure-native:storage/v20210201:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tableState struct {
}

type TableState struct {
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	AccountName       string  `pulumi:"accountName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	TableName         *string `pulumi:"tableName"`
}


type TableArgs struct {
	AccountName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TableName         pulumi.StringPtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TableOutput{})
}
