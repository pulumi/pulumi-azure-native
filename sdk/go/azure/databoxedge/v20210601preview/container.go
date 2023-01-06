


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Container struct {
	pulumi.CustomResourceState

	ContainerStatus pulumi.StringOutput          `pulumi:"containerStatus"`
	CreatedDateTime pulumi.StringOutput          `pulumi:"createdDateTime"`
	DataFormat      pulumi.StringOutput          `pulumi:"dataFormat"`
	Name            pulumi.StringOutput          `pulumi:"name"`
	RefreshDetails  RefreshDetailsResponseOutput `pulumi:"refreshDetails"`
	SystemData      SystemDataResponseOutput     `pulumi:"systemData"`
	Type            pulumi.StringOutput          `pulumi:"type"`
}


func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFormat == nil {
		return nil, errors.New("invalid value for required argument 'DataFormat'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:Container"),
		},
	})
	opts = append(opts, aliases)
	var resource Container
	err := ctx.RegisterResource("azure-native:databoxedge/v20210601preview:Container", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerState, opts ...pulumi.ResourceOption) (*Container, error) {
	var resource Container
	err := ctx.ReadResource("azure-native:databoxedge/v20210601preview:Container", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type containerState struct {
}

type ContainerState struct {
}

func (ContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerState)(nil)).Elem()
}

type containerArgs struct {
	ContainerName      *string `pulumi:"containerName"`
	DataFormat         string  `pulumi:"dataFormat"`
	DeviceName         string  `pulumi:"deviceName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	StorageAccountName string  `pulumi:"storageAccountName"`
}


type ContainerArgs struct {
	ContainerName      pulumi.StringPtrInput
	DataFormat         pulumi.StringInput
	DeviceName         pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	StorageAccountName pulumi.StringInput
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerArgs)(nil)).Elem()
}

type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(ctx context.Context) ContainerOutput
}

func (*Container) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (i *Container) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i *Container) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func (o ContainerOutput) ContainerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.ContainerStatus }).(pulumi.StringOutput)
}

func (o ContainerOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o ContainerOutput) DataFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.DataFormat }).(pulumi.StringOutput)
}

func (o ContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerOutput) RefreshDetails() RefreshDetailsResponseOutput {
	return o.ApplyT(func(v *Container) RefreshDetailsResponseOutput { return v.RefreshDetails }).(RefreshDetailsResponseOutput)
}

func (o ContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Container) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerOutput{})
}
