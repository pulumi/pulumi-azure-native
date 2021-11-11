


package v20150401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseAccountCassandraKeyspace struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewDatabaseAccountCassandraKeyspace(ctx *pulumi.Context,
	name string, args *DatabaseAccountCassandraKeyspaceArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraKeyspace, error) {
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
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountCassandraKeyspace"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountCassandraKeyspace
	err := ctx.RegisterResource("azure-native:documentdb/v20150401:DatabaseAccountCassandraKeyspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountCassandraKeyspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountCassandraKeyspaceState, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraKeyspace, error) {
	var resource DatabaseAccountCassandraKeyspace
	err := ctx.ReadResource("azure-native:documentdb/v20150401:DatabaseAccountCassandraKeyspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountCassandraKeyspaceState struct {
}

type DatabaseAccountCassandraKeyspaceState struct {
}

func (DatabaseAccountCassandraKeyspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraKeyspaceState)(nil)).Elem()
}

type databaseAccountCassandraKeyspaceArgs struct {
	AccountName       string                    `pulumi:"accountName"`
	KeyspaceName      *string                   `pulumi:"keyspaceName"`
	Options           map[string]string         `pulumi:"options"`
	Resource          CassandraKeyspaceResource `pulumi:"resource"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
}


type DatabaseAccountCassandraKeyspaceArgs struct {
	AccountName       pulumi.StringInput
	KeyspaceName      pulumi.StringPtrInput
	Options           pulumi.StringMapInput
	Resource          CassandraKeyspaceResourceInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountCassandraKeyspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraKeyspaceArgs)(nil)).Elem()
}

type DatabaseAccountCassandraKeyspaceInput interface {
	pulumi.Input

	ToDatabaseAccountCassandraKeyspaceOutput() DatabaseAccountCassandraKeyspaceOutput
	ToDatabaseAccountCassandraKeyspaceOutputWithContext(ctx context.Context) DatabaseAccountCassandraKeyspaceOutput
}

func (*DatabaseAccountCassandraKeyspace) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountCassandraKeyspace)(nil))
}

func (i *DatabaseAccountCassandraKeyspace) ToDatabaseAccountCassandraKeyspaceOutput() DatabaseAccountCassandraKeyspaceOutput {
	return i.ToDatabaseAccountCassandraKeyspaceOutputWithContext(context.Background())
}

func (i *DatabaseAccountCassandraKeyspace) ToDatabaseAccountCassandraKeyspaceOutputWithContext(ctx context.Context) DatabaseAccountCassandraKeyspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountCassandraKeyspaceOutput)
}

type DatabaseAccountCassandraKeyspaceOutput struct{ *pulumi.OutputState }

func (DatabaseAccountCassandraKeyspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountCassandraKeyspace)(nil))
}

func (o DatabaseAccountCassandraKeyspaceOutput) ToDatabaseAccountCassandraKeyspaceOutput() DatabaseAccountCassandraKeyspaceOutput {
	return o
}

func (o DatabaseAccountCassandraKeyspaceOutput) ToDatabaseAccountCassandraKeyspaceOutputWithContext(ctx context.Context) DatabaseAccountCassandraKeyspaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountCassandraKeyspaceOutput{})
}
