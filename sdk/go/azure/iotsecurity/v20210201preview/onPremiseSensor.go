


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OnPremiseSensor struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewOnPremiseSensor(ctx *pulumi.Context,
	name string, args *OnPremiseSensorArgs, opts ...pulumi.ResourceOption) (*OnPremiseSensor, error) {
	if args == nil {
		args = &OnPremiseSensorArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:iotsecurity/v20210201preview:OnPremiseSensor"),
		},
		{
			Type: pulumi.String("azure-native:iotsecurity:OnPremiseSensor"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotsecurity:OnPremiseSensor"),
		},
	})
	opts = append(opts, aliases)
	var resource OnPremiseSensor
	err := ctx.RegisterResource("azure-native:iotsecurity/v20210201preview:OnPremiseSensor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOnPremiseSensor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnPremiseSensorState, opts ...pulumi.ResourceOption) (*OnPremiseSensor, error) {
	var resource OnPremiseSensor
	err := ctx.ReadResource("azure-native:iotsecurity/v20210201preview:OnPremiseSensor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type onPremiseSensorState struct {
}

type OnPremiseSensorState struct {
}

func (OnPremiseSensorState) ElementType() reflect.Type {
	return reflect.TypeOf((*onPremiseSensorState)(nil)).Elem()
}

type onPremiseSensorArgs struct {
	OnPremiseSensorName *string `pulumi:"onPremiseSensorName"`
}


type OnPremiseSensorArgs struct {
	OnPremiseSensorName pulumi.StringPtrInput
}

func (OnPremiseSensorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*onPremiseSensorArgs)(nil)).Elem()
}

type OnPremiseSensorInput interface {
	pulumi.Input

	ToOnPremiseSensorOutput() OnPremiseSensorOutput
	ToOnPremiseSensorOutputWithContext(ctx context.Context) OnPremiseSensorOutput
}

func (*OnPremiseSensor) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSensor)(nil))
}

func (i *OnPremiseSensor) ToOnPremiseSensorOutput() OnPremiseSensorOutput {
	return i.ToOnPremiseSensorOutputWithContext(context.Background())
}

func (i *OnPremiseSensor) ToOnPremiseSensorOutputWithContext(ctx context.Context) OnPremiseSensorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseSensorOutput)
}

type OnPremiseSensorOutput struct{ *pulumi.OutputState }

func (OnPremiseSensorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSensor)(nil))
}

func (o OnPremiseSensorOutput) ToOnPremiseSensorOutput() OnPremiseSensorOutput {
	return o
}

func (o OnPremiseSensorOutput) ToOnPremiseSensorOutputWithContext(ctx context.Context) OnPremiseSensorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OnPremiseSensorOutput{})
}
