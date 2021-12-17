


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CassandraResourceCassandraKeyspace struct {
	pulumi.CustomResourceState

	Identity ManagedServiceIdentityResponsePtrOutput                 `pulumi:"identity"`
	Location pulumi.StringPtrOutput                                  `pulumi:"location"`
	Name     pulumi.StringOutput                                     `pulumi:"name"`
	Options  CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource CassandraKeyspaceGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type     pulumi.StringOutput                                     `pulumi:"type"`
}


func NewCassandraResourceCassandraKeyspace(ctx *pulumi.Context,
	name string, args *CassandraResourceCassandraKeyspaceArgs, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraKeyspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:CassandraResourceCassandraKeyspace"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraResourceCassandraKeyspace
	err := ctx.RegisterResource("azure-native:documentdb/v20210401preview:CassandraResourceCassandraKeyspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCassandraResourceCassandraKeyspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraResourceCassandraKeyspaceState, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraKeyspace, error) {
	var resource CassandraResourceCassandraKeyspace
	err := ctx.ReadResource("azure-native:documentdb/v20210401preview:CassandraResourceCassandraKeyspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cassandraResourceCassandraKeyspaceState struct {
}

type CassandraResourceCassandraKeyspaceState struct {
}

func (CassandraResourceCassandraKeyspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraKeyspaceState)(nil)).Elem()
}

type cassandraResourceCassandraKeyspaceArgs struct {
	AccountName       string                    `pulumi:"accountName"`
	Identity          *ManagedServiceIdentity   `pulumi:"identity"`
	KeyspaceName      *string                   `pulumi:"keyspaceName"`
	Location          *string                   `pulumi:"location"`
	Options           *CreateUpdateOptions      `pulumi:"options"`
	Resource          CassandraKeyspaceResource `pulumi:"resource"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type CassandraResourceCassandraKeyspaceArgs struct {
	AccountName       pulumi.StringInput
	Identity          ManagedServiceIdentityPtrInput
	KeyspaceName      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsPtrInput
	Resource          CassandraKeyspaceResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (CassandraResourceCassandraKeyspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraKeyspaceArgs)(nil)).Elem()
}

type CassandraResourceCassandraKeyspaceInput interface {
	pulumi.Input

	ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput
	ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput
}

func (*CassandraResourceCassandraKeyspace) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraResourceCassandraKeyspace)(nil)).Elem()
}

func (i *CassandraResourceCassandraKeyspace) ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput {
	return i.ToCassandraResourceCassandraKeyspaceOutputWithContext(context.Background())
}

func (i *CassandraResourceCassandraKeyspace) ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraResourceCassandraKeyspaceOutput)
}

type CassandraResourceCassandraKeyspaceOutput struct{ *pulumi.OutputState }

func (CassandraResourceCassandraKeyspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraResourceCassandraKeyspace)(nil)).Elem()
}

func (o CassandraResourceCassandraKeyspaceOutput) ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput {
	return o
}

func (o CassandraResourceCassandraKeyspaceOutput) ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CassandraResourceCassandraKeyspaceOutput{})
}
