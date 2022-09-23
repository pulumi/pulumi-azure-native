


package v20180501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SerialPort struct {
	pulumi.CustomResourceState

	Name  pulumi.StringOutput    `pulumi:"name"`
	State pulumi.StringPtrOutput `pulumi:"state"`
	Type  pulumi.StringOutput    `pulumi:"type"`
}


func NewSerialPort(ctx *pulumi.Context,
	name string, args *SerialPortArgs, opts ...pulumi.ResourceOption) (*SerialPort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentResource == nil {
		return nil, errors.New("invalid value for required argument 'ParentResource'")
	}
	if args.ParentResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ResourceProviderNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:serialconsole:SerialPort"),
		},
	})
	opts = append(opts, aliases)
	var resource SerialPort
	err := ctx.RegisterResource("azure-native:serialconsole/v20180501:SerialPort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSerialPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SerialPortState, opts ...pulumi.ResourceOption) (*SerialPort, error) {
	var resource SerialPort
	err := ctx.ReadResource("azure-native:serialconsole/v20180501:SerialPort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serialPortState struct {
}

type SerialPortState struct {
}

func (SerialPortState) ElementType() reflect.Type {
	return reflect.TypeOf((*serialPortState)(nil)).Elem()
}

type serialPortArgs struct {
	ParentResource            string               `pulumi:"parentResource"`
	ParentResourceType        string               `pulumi:"parentResourceType"`
	ResourceGroupName         string               `pulumi:"resourceGroupName"`
	ResourceProviderNamespace string               `pulumi:"resourceProviderNamespace"`
	SerialPort                *string              `pulumi:"serialPort"`
	State                     *SerialPortStateEnum `pulumi:"state"`
}


type SerialPortArgs struct {
	ParentResource            pulumi.StringInput
	ParentResourceType        pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	ResourceProviderNamespace pulumi.StringInput
	SerialPort                pulumi.StringPtrInput
	State                     SerialPortStateEnumPtrInput
}

func (SerialPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serialPortArgs)(nil)).Elem()
}

type SerialPortInput interface {
	pulumi.Input

	ToSerialPortOutput() SerialPortOutput
	ToSerialPortOutputWithContext(ctx context.Context) SerialPortOutput
}

func (*SerialPort) ElementType() reflect.Type {
	return reflect.TypeOf((**SerialPort)(nil)).Elem()
}

func (i *SerialPort) ToSerialPortOutput() SerialPortOutput {
	return i.ToSerialPortOutputWithContext(context.Background())
}

func (i *SerialPort) ToSerialPortOutputWithContext(ctx context.Context) SerialPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SerialPortOutput)
}

type SerialPortOutput struct{ *pulumi.OutputState }

func (SerialPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SerialPort)(nil)).Elem()
}

func (o SerialPortOutput) ToSerialPortOutput() SerialPortOutput {
	return o
}

func (o SerialPortOutput) ToSerialPortOutputWithContext(ctx context.Context) SerialPortOutput {
	return o
}

func (o SerialPortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SerialPort) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SerialPortOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SerialPort) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o SerialPortOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SerialPort) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SerialPortOutput{})
}
