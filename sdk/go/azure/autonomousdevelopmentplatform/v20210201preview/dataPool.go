


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataPool struct {
	pulumi.CustomResourceState

	DataPoolId        pulumi.StringOutput                 `pulumi:"dataPoolId"`
	Locations         DataPoolLocationResponseArrayOutput `pulumi:"locations"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                 `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput            `pulumi:"systemData"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewDataPool(ctx *pulumi.Context,
	name string, args *DataPoolArgs, opts ...pulumi.ResourceOption) (*DataPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:autonomousdevelopmentplatform/v20210201preview:DataPool"),
		},
		{
			Type: pulumi.String("azure-native:autonomousdevelopmentplatform:DataPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:autonomousdevelopmentplatform:DataPool"),
		},
		{
			Type: pulumi.String("azure-native:autonomousdevelopmentplatform/v20200701preview:DataPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:autonomousdevelopmentplatform/v20200701preview:DataPool"),
		},
	})
	opts = append(opts, aliases)
	var resource DataPool
	err := ctx.RegisterResource("azure-native:autonomousdevelopmentplatform/v20210201preview:DataPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataPoolState, opts ...pulumi.ResourceOption) (*DataPool, error) {
	var resource DataPool
	err := ctx.ReadResource("azure-native:autonomousdevelopmentplatform/v20210201preview:DataPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataPoolState struct {
}

type DataPoolState struct {
}

func (DataPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataPoolState)(nil)).Elem()
}

type dataPoolArgs struct {
	AccountName       string             `pulumi:"accountName"`
	DataPoolName      *string            `pulumi:"dataPoolName"`
	Locations         []DataPoolLocation `pulumi:"locations"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
}


type DataPoolArgs struct {
	AccountName       pulumi.StringInput
	DataPoolName      pulumi.StringPtrInput
	Locations         DataPoolLocationArrayInput
	ResourceGroupName pulumi.StringInput
}

func (DataPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataPoolArgs)(nil)).Elem()
}

type DataPoolInput interface {
	pulumi.Input

	ToDataPoolOutput() DataPoolOutput
	ToDataPoolOutputWithContext(ctx context.Context) DataPoolOutput
}

func (*DataPool) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPool)(nil))
}

func (i *DataPool) ToDataPoolOutput() DataPoolOutput {
	return i.ToDataPoolOutputWithContext(context.Background())
}

func (i *DataPool) ToDataPoolOutputWithContext(ctx context.Context) DataPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolOutput)
}

type DataPoolOutput struct{ *pulumi.OutputState }

func (DataPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPool)(nil))
}

func (o DataPoolOutput) ToDataPoolOutput() DataPoolOutput {
	return o
}

func (o DataPoolOutput) ToDataPoolOutputWithContext(ctx context.Context) DataPoolOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataPoolInput)(nil)).Elem(), &DataPool{})
	pulumi.RegisterOutputType(DataPoolOutput{})
}
