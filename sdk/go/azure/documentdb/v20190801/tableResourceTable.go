


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type TableResourceTable struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                      `pulumi:"location"`
	Name     pulumi.StringOutput                         `pulumi:"name"`
	Resource TableGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                      `pulumi:"tags"`
	Type     pulumi.StringOutput                         `pulumi:"type"`
}


func NewTableResourceTable(ctx *pulumi.Context,
	name string, args *TableResourceTableArgs, opts ...pulumi.ResourceOption) (*TableResourceTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:TableResourceTable"),
		},
	})
	opts = append(opts, aliases)
	var resource TableResourceTable
	err := ctx.RegisterResource("azure-native:documentdb/v20190801:TableResourceTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTableResourceTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableResourceTableState, opts ...pulumi.ResourceOption) (*TableResourceTable, error) {
	var resource TableResourceTable
	err := ctx.ReadResource("azure-native:documentdb/v20190801:TableResourceTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tableResourceTableState struct {
}

type TableResourceTableState struct {
}

func (TableResourceTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableResourceTableState)(nil)).Elem()
}

type tableResourceTableArgs struct {
	AccountName       string            `pulumi:"accountName"`
	Location          *string           `pulumi:"location"`
	Options           map[string]string `pulumi:"options"`
	Resource          TableResource     `pulumi:"resource"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	TableName         *string           `pulumi:"tableName"`
	Tags              map[string]string `pulumi:"tags"`
}


type TableResourceTableArgs struct {
	AccountName       pulumi.StringInput
	Location          pulumi.StringPtrInput
	Options           pulumi.StringMapInput
	Resource          TableResourceInput
	ResourceGroupName pulumi.StringInput
	TableName         pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (TableResourceTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableResourceTableArgs)(nil)).Elem()
}

type TableResourceTableInput interface {
	pulumi.Input

	ToTableResourceTableOutput() TableResourceTableOutput
	ToTableResourceTableOutputWithContext(ctx context.Context) TableResourceTableOutput
}

func (*TableResourceTable) ElementType() reflect.Type {
	return reflect.TypeOf((**TableResourceTable)(nil)).Elem()
}

func (i *TableResourceTable) ToTableResourceTableOutput() TableResourceTableOutput {
	return i.ToTableResourceTableOutputWithContext(context.Background())
}

func (i *TableResourceTable) ToTableResourceTableOutputWithContext(ctx context.Context) TableResourceTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableResourceTableOutput)
}

type TableResourceTableOutput struct{ *pulumi.OutputState }

func (TableResourceTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableResourceTable)(nil)).Elem()
}

func (o TableResourceTableOutput) ToTableResourceTableOutput() TableResourceTableOutput {
	return o
}

func (o TableResourceTableOutput) ToTableResourceTableOutputWithContext(ctx context.Context) TableResourceTableOutput {
	return o
}

func (o TableResourceTableOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableResourceTable) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o TableResourceTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TableResourceTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TableResourceTableOutput) Resource() TableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *TableResourceTable) TableGetPropertiesResponseResourcePtrOutput { return v.Resource }).(TableGetPropertiesResponseResourcePtrOutput)
}

func (o TableResourceTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TableResourceTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o TableResourceTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TableResourceTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TableResourceTableOutput{})
}
