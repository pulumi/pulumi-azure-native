


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataContainer struct {
	pulumi.CustomResourceState

	DataContainerProperties DataContainerResponseOutput `pulumi:"dataContainerProperties"`
	Name                    pulumi.StringOutput         `pulumi:"name"`
	SystemData              SystemDataResponseOutput    `pulumi:"systemData"`
	Type                    pulumi.StringOutput         `pulumi:"type"`
}


func NewDataContainer(ctx *pulumi.Context,
	name string, args *DataContainerArgs, opts ...pulumi.ResourceOption) (*DataContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'DataContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.DataContainerProperties = args.DataContainerProperties.ToDataContainerTypeOutput().ApplyT(func(v DataContainerType) DataContainerType { return *v.Defaults() }).(DataContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:DataContainer"),
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
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220501:DataContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataContainerState, opts ...pulumi.ResourceOption) (*DataContainer, error) {
	var resource DataContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220501:DataContainer", name, id, state, &resource, opts...)
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
	DataContainerProperties DataContainerType `pulumi:"dataContainerProperties"`
	Name                    *string           `pulumi:"name"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	WorkspaceName           string            `pulumi:"workspaceName"`
}


type DataContainerArgs struct {
	DataContainerProperties DataContainerTypeInput
	Name                    pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	WorkspaceName           pulumi.StringInput
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

func (o DataContainerOutput) DataContainerProperties() DataContainerResponseOutput {
	return o.ApplyT(func(v *DataContainer) DataContainerResponseOutput { return v.DataContainerProperties }).(DataContainerResponseOutput)
}

func (o DataContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
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
