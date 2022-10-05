


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataContainer struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput         `pulumi:"name"`
	Properties DataContainerResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput    `pulumi:"systemData"`
	Type       pulumi.StringOutput         `pulumi:"type"`
}


func NewDataContainer(ctx *pulumi.Context,
	name string, args *DataContainerArgs, opts ...pulumi.ResourceOption) (*DataContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:DataContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource DataContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:DataContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataContainerState, opts ...pulumi.ResourceOption) (*DataContainer, error) {
	var resource DataContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:DataContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataContainerState struct {
}

type DataContainerState struct {
}

func (DataContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataContainerState)(nil)).Elem()
}

type dataContainerArgs struct {
	Name              *string           `pulumi:"name"`
	Properties        DataContainerType `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	WorkspaceName     string            `pulumi:"workspaceName"`
}


type DataContainerArgs struct {
	Name              pulumi.StringPtrInput
	Properties        DataContainerTypeInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (DataContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataContainerArgs)(nil)).Elem()
}

type DataContainerInput interface {
	pulumi.Input

	ToDataContainerOutput() DataContainerOutput
	ToDataContainerOutputWithContext(ctx context.Context) DataContainerOutput
}

func (*DataContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**DataContainer)(nil)).Elem()
}

func (i *DataContainer) ToDataContainerOutput() DataContainerOutput {
	return i.ToDataContainerOutputWithContext(context.Background())
}

func (i *DataContainer) ToDataContainerOutputWithContext(ctx context.Context) DataContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataContainerOutput)
}

type DataContainerOutput struct{ *pulumi.OutputState }

func (DataContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataContainer)(nil)).Elem()
}

func (o DataContainerOutput) ToDataContainerOutput() DataContainerOutput {
	return o
}

func (o DataContainerOutput) ToDataContainerOutputWithContext(ctx context.Context) DataContainerOutput {
	return o
}

func (o DataContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataContainerOutput) Properties() DataContainerResponseOutput {
	return o.ApplyT(func(v *DataContainer) DataContainerResponseOutput { return v.Properties }).(DataContainerResponseOutput)
}

func (o DataContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DataContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataContainerOutput{})
}
