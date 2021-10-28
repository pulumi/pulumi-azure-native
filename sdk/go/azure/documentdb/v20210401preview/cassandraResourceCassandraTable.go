


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CassandraResourceCassandraTable struct {
	pulumi.CustomResourceState

	Identity ManagedServiceIdentityResponsePtrOutput              `pulumi:"identity"`
	Location pulumi.StringPtrOutput                               `pulumi:"location"`
	Name     pulumi.StringOutput                                  `pulumi:"name"`
	Options  CassandraTableGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource CassandraTableGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                               `pulumi:"tags"`
	Type     pulumi.StringOutput                                  `pulumi:"type"`
}


func NewCassandraResourceCassandraTable(ctx *pulumi.Context,
	name string, args *CassandraResourceCassandraTableArgs, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.KeyspaceName == nil {
		return nil, errors.New("invalid value for required argument 'KeyspaceName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210401preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200901:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210115:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210301preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210315:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210415:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210515:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210615:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210701preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20211015:CassandraResourceCassandraTable"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraResourceCassandraTable
	err := ctx.RegisterResource("azure-native:documentdb/v20210401preview:CassandraResourceCassandraTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCassandraResourceCassandraTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraResourceCassandraTableState, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraTable, error) {
	var resource CassandraResourceCassandraTable
	err := ctx.ReadResource("azure-native:documentdb/v20210401preview:CassandraResourceCassandraTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cassandraResourceCassandraTableState struct {
}

type CassandraResourceCassandraTableState struct {
}

func (CassandraResourceCassandraTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraTableState)(nil)).Elem()
}

type cassandraResourceCassandraTableArgs struct {
	AccountName       string                  `pulumi:"accountName"`
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	KeyspaceName      string                  `pulumi:"keyspaceName"`
	Location          *string                 `pulumi:"location"`
	Options           *CreateUpdateOptions    `pulumi:"options"`
	Resource          CassandraTableResource  `pulumi:"resource"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	TableName         *string                 `pulumi:"tableName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type CassandraResourceCassandraTableArgs struct {
	AccountName       pulumi.StringInput
	Identity          ManagedServiceIdentityPtrInput
	KeyspaceName      pulumi.StringInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsPtrInput
	Resource          CassandraTableResourceInput
	ResourceGroupName pulumi.StringInput
	TableName         pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (CassandraResourceCassandraTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraTableArgs)(nil)).Elem()
}

type CassandraResourceCassandraTableInput interface {
	pulumi.Input

	ToCassandraResourceCassandraTableOutput() CassandraResourceCassandraTableOutput
	ToCassandraResourceCassandraTableOutputWithContext(ctx context.Context) CassandraResourceCassandraTableOutput
}

func (*CassandraResourceCassandraTable) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraResourceCassandraTable)(nil))
}

func (i *CassandraResourceCassandraTable) ToCassandraResourceCassandraTableOutput() CassandraResourceCassandraTableOutput {
	return i.ToCassandraResourceCassandraTableOutputWithContext(context.Background())
}

func (i *CassandraResourceCassandraTable) ToCassandraResourceCassandraTableOutputWithContext(ctx context.Context) CassandraResourceCassandraTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraResourceCassandraTableOutput)
}

type CassandraResourceCassandraTableOutput struct{ *pulumi.OutputState }

func (CassandraResourceCassandraTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraResourceCassandraTable)(nil))
}

func (o CassandraResourceCassandraTableOutput) ToCassandraResourceCassandraTableOutput() CassandraResourceCassandraTableOutput {
	return o
}

func (o CassandraResourceCassandraTableOutput) ToCassandraResourceCassandraTableOutputWithContext(ctx context.Context) CassandraResourceCassandraTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CassandraResourceCassandraTableOutput{})
}
