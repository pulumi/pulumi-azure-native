


package v20211101preview

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
	Tags              pulumi.StringMapOutput              `pulumi:"tags"`
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
			Type: pulumi.String("azure-native:autonomousdevelopmentplatform:DataPool"),
		},
		{
			Type: pulumi.String("azure-native:autonomousdevelopmentplatform/v20200701preview:DataPool"),
		},
		{
			Type: pulumi.String("azure-native:autonomousdevelopmentplatform/v20210201preview:DataPool"),
		},
	})
	opts = append(opts, aliases)
	var resource DataPool
	err := ctx.RegisterResource("azure-native:autonomousdevelopmentplatform/v20211101preview:DataPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataPoolState, opts ...pulumi.ResourceOption) (*DataPool, error) {
	var resource DataPool
	err := ctx.ReadResource("azure-native:autonomousdevelopmentplatform/v20211101preview:DataPool", name, id, state, &resource, opts...)
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
	Tags              map[string]string  `pulumi:"tags"`
}


type DataPoolArgs struct {
	AccountName       pulumi.StringInput
	DataPoolName      pulumi.StringPtrInput
	Locations         DataPoolLocationArrayInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
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
	pulumi.RegisterOutputType(DataPoolOutput{})
}
