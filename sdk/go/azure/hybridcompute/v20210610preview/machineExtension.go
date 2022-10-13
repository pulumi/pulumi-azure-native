


package v20210610preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachineExtension struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                      `pulumi:"location"`
	Name       pulumi.StringOutput                      `pulumi:"name"`
	Properties MachineExtensionPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                 `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                   `pulumi:"tags"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewMachineExtension(ctx *pulumi.Context,
	name string, args *MachineExtensionArgs, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190802preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20191212:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200730preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200802:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20211210preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220310:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220510preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220811preview:MachineExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineExtension
	err := ctx.RegisterResource("azure-native:hybridcompute/v20210610preview:MachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineExtensionState, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	var resource MachineExtension
	err := ctx.ReadResource("azure-native:hybridcompute/v20210610preview:MachineExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineExtensionState struct {
}

type MachineExtensionState struct {
}

func (MachineExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionState)(nil)).Elem()
}

type machineExtensionArgs struct {
	ExtensionName     *string                     `pulumi:"extensionName"`
	Location          *string                     `pulumi:"location"`
	MachineName       string                      `pulumi:"machineName"`
	Properties        *MachineExtensionProperties `pulumi:"properties"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	Tags              map[string]string           `pulumi:"tags"`
}


type MachineExtensionArgs struct {
	ExtensionName     pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MachineName       pulumi.StringInput
	Properties        MachineExtensionPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (MachineExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionArgs)(nil)).Elem()
}

type MachineExtensionInput interface {
	pulumi.Input

	ToMachineExtensionOutput() MachineExtensionOutput
	ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput
}

func (*MachineExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtension)(nil)).Elem()
}

func (i *MachineExtension) ToMachineExtensionOutput() MachineExtensionOutput {
	return i.ToMachineExtensionOutputWithContext(context.Background())
}

func (i *MachineExtension) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionOutput)
}

type MachineExtensionOutput struct{ *pulumi.OutputState }

func (MachineExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtension)(nil)).Elem()
}

func (o MachineExtensionOutput) ToMachineExtensionOutput() MachineExtensionOutput {
	return o
}

func (o MachineExtensionOutput) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return o
}

func (o MachineExtensionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MachineExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionOutput) Properties() MachineExtensionPropertiesResponseOutput {
	return o.ApplyT(func(v *MachineExtension) MachineExtensionPropertiesResponseOutput { return v.Properties }).(MachineExtensionPropertiesResponseOutput)
}

func (o MachineExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MachineExtension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MachineExtensionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MachineExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineExtensionOutput{})
}
