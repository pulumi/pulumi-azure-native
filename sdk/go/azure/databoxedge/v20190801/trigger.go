


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Trigger struct {
	pulumi.CustomResourceState

	Kind pulumi.StringOutput `pulumi:"kind"`
	Name pulumi.StringOutput `pulumi:"name"`
	Type pulumi.StringOutput `pulumi:"type"`
}


func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190801:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190301:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190701:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200501preview:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901preview:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20201201:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Trigger"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201preview:Trigger"),
		},
	})
	opts = append(opts, aliases)
	var resource Trigger
	err := ctx.RegisterResource("azure-native:databoxedge/v20190801:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("azure-native:databoxedge/v20190801:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type triggerState struct {
}

type TriggerState struct {
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	DeviceName        string  `pulumi:"deviceName"`
	Kind              string  `pulumi:"kind"`
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type TriggerArgs struct {
	DeviceName        pulumi.StringInput
	Kind              pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil))
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil))
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TriggerOutput{})
}
