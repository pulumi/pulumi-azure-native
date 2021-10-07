


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GremlinResourceGremlinDatabase struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                                `pulumi:"location"`
	Name     pulumi.StringOutput                                   `pulumi:"name"`
	Resource GremlinDatabaseGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                                `pulumi:"tags"`
	Type     pulumi.StringOutput                                   `pulumi:"type"`
}


func NewGremlinResourceGremlinDatabase(ctx *pulumi.Context,
	name string, args *GremlinResourceGremlinDatabaseArgs, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinDatabase, error) {
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
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200901:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210115:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210301preview:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210315:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210401preview:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210415:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210515:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210615:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210701preview:GremlinResourceGremlinDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource GremlinResourceGremlinDatabase
	err := ctx.RegisterResource("azure-native:documentdb/v20190801:GremlinResourceGremlinDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGremlinResourceGremlinDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GremlinResourceGremlinDatabaseState, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinDatabase, error) {
	var resource GremlinResourceGremlinDatabase
	err := ctx.ReadResource("azure-native:documentdb/v20190801:GremlinResourceGremlinDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gremlinResourceGremlinDatabaseState struct {
}

type GremlinResourceGremlinDatabaseState struct {
}

func (GremlinResourceGremlinDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinDatabaseState)(nil)).Elem()
}

type gremlinResourceGremlinDatabaseArgs struct {
	AccountName       string                  `pulumi:"accountName"`
	DatabaseName      *string                 `pulumi:"databaseName"`
	Location          *string                 `pulumi:"location"`
	Options           map[string]string       `pulumi:"options"`
	Resource          GremlinDatabaseResource `pulumi:"resource"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type GremlinResourceGremlinDatabaseArgs struct {
	AccountName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Options           pulumi.StringMapInput
	Resource          GremlinDatabaseResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (GremlinResourceGremlinDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinDatabaseArgs)(nil)).Elem()
}

type GremlinResourceGremlinDatabaseInput interface {
	pulumi.Input

	ToGremlinResourceGremlinDatabaseOutput() GremlinResourceGremlinDatabaseOutput
	ToGremlinResourceGremlinDatabaseOutputWithContext(ctx context.Context) GremlinResourceGremlinDatabaseOutput
}

func (*GremlinResourceGremlinDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinResourceGremlinDatabase)(nil))
}

func (i *GremlinResourceGremlinDatabase) ToGremlinResourceGremlinDatabaseOutput() GremlinResourceGremlinDatabaseOutput {
	return i.ToGremlinResourceGremlinDatabaseOutputWithContext(context.Background())
}

func (i *GremlinResourceGremlinDatabase) ToGremlinResourceGremlinDatabaseOutputWithContext(ctx context.Context) GremlinResourceGremlinDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinResourceGremlinDatabaseOutput)
}

type GremlinResourceGremlinDatabaseOutput struct{ *pulumi.OutputState }

func (GremlinResourceGremlinDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinResourceGremlinDatabase)(nil))
}

func (o GremlinResourceGremlinDatabaseOutput) ToGremlinResourceGremlinDatabaseOutput() GremlinResourceGremlinDatabaseOutput {
	return o
}

func (o GremlinResourceGremlinDatabaseOutput) ToGremlinResourceGremlinDatabaseOutputWithContext(ctx context.Context) GremlinResourceGremlinDatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GremlinResourceGremlinDatabaseOutput{})
}
