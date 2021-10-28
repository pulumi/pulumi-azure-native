


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataStore struct {
	pulumi.CustomResourceState

	CustomerSecrets    CustomerSecretResponseArrayOutput `pulumi:"customerSecrets"`
	DataStoreTypeId    pulumi.StringOutput               `pulumi:"dataStoreTypeId"`
	ExtendedProperties pulumi.AnyOutput                  `pulumi:"extendedProperties"`
	Name               pulumi.StringOutput               `pulumi:"name"`
	RepositoryId       pulumi.StringPtrOutput            `pulumi:"repositoryId"`
	State              pulumi.StringOutput               `pulumi:"state"`
	Type               pulumi.StringOutput               `pulumi:"type"`
}


func NewDataStore(ctx *pulumi.Context,
	name string, args *DataStoreArgs, opts ...pulumi.ResourceOption) (*DataStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataManagerName == nil {
		return nil, errors.New("invalid value for required argument 'DataManagerName'")
	}
	if args.DataStoreTypeId == nil {
		return nil, errors.New("invalid value for required argument 'DataStoreTypeId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:hybriddata/v20190601:DataStore"),
		},
		{
			Type: pulumi.String("azure-native:hybriddata:DataStore"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybriddata:DataStore"),
		},
		{
			Type: pulumi.String("azure-native:hybriddata/v20160601:DataStore"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybriddata/v20160601:DataStore"),
		},
	})
	opts = append(opts, aliases)
	var resource DataStore
	err := ctx.RegisterResource("azure-native:hybriddata/v20190601:DataStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataStoreState, opts ...pulumi.ResourceOption) (*DataStore, error) {
	var resource DataStore
	err := ctx.ReadResource("azure-native:hybriddata/v20190601:DataStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataStoreState struct {
}

type DataStoreState struct {
}

func (DataStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataStoreState)(nil)).Elem()
}

type dataStoreArgs struct {
	CustomerSecrets    []CustomerSecret `pulumi:"customerSecrets"`
	DataManagerName    string           `pulumi:"dataManagerName"`
	DataStoreName      *string          `pulumi:"dataStoreName"`
	DataStoreTypeId    string           `pulumi:"dataStoreTypeId"`
	ExtendedProperties interface{}      `pulumi:"extendedProperties"`
	RepositoryId       *string          `pulumi:"repositoryId"`
	ResourceGroupName  string           `pulumi:"resourceGroupName"`
	State              State            `pulumi:"state"`
}


type DataStoreArgs struct {
	CustomerSecrets    CustomerSecretArrayInput
	DataManagerName    pulumi.StringInput
	DataStoreName      pulumi.StringPtrInput
	DataStoreTypeId    pulumi.StringInput
	ExtendedProperties pulumi.Input
	RepositoryId       pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	State              StateInput
}

func (DataStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataStoreArgs)(nil)).Elem()
}

type DataStoreInput interface {
	pulumi.Input

	ToDataStoreOutput() DataStoreOutput
	ToDataStoreOutputWithContext(ctx context.Context) DataStoreOutput
}

func (*DataStore) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStore)(nil))
}

func (i *DataStore) ToDataStoreOutput() DataStoreOutput {
	return i.ToDataStoreOutputWithContext(context.Background())
}

func (i *DataStore) ToDataStoreOutputWithContext(ctx context.Context) DataStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreOutput)
}

type DataStoreOutput struct{ *pulumi.OutputState }

func (DataStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStore)(nil))
}

func (o DataStoreOutput) ToDataStoreOutput() DataStoreOutput {
	return o
}

func (o DataStoreOutput) ToDataStoreOutputWithContext(ctx context.Context) DataStoreOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataStoreOutput{})
}
