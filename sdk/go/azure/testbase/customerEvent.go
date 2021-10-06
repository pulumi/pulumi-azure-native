


package testbase

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomerEvent struct {
	pulumi.CustomResourceState

	EventName  pulumi.StringOutput                          `pulumi:"eventName"`
	Name       pulumi.StringOutput                          `pulumi:"name"`
	Receivers  NotificationEventReceiverResponseArrayOutput `pulumi:"receivers"`
	SystemData SystemDataResponseOutput                     `pulumi:"systemData"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewCustomerEvent(ctx *pulumi.Context,
	name string, args *CustomerEventArgs, opts ...pulumi.ResourceOption) (*CustomerEvent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventName == nil {
		return nil, errors.New("invalid value for required argument 'EventName'")
	}
	if args.Receivers == nil {
		return nil, errors.New("invalid value for required argument 'Receivers'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TestBaseAccountName == nil {
		return nil, errors.New("invalid value for required argument 'TestBaseAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:testbase:CustomerEvent"),
		},
		{
			Type: pulumi.String("azure-native:testbase/v20201216preview:CustomerEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:testbase/v20201216preview:CustomerEvent"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomerEvent
	err := ctx.RegisterResource("azure-native:testbase:CustomerEvent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomerEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerEventState, opts ...pulumi.ResourceOption) (*CustomerEvent, error) {
	var resource CustomerEvent
	err := ctx.ReadResource("azure-native:testbase:CustomerEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customerEventState struct {
}

type CustomerEventState struct {
}

func (CustomerEventState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerEventState)(nil)).Elem()
}

type customerEventArgs struct {
	CustomerEventName   *string                     `pulumi:"customerEventName"`
	EventName           string                      `pulumi:"eventName"`
	Receivers           []NotificationEventReceiver `pulumi:"receivers"`
	ResourceGroupName   string                      `pulumi:"resourceGroupName"`
	TestBaseAccountName string                      `pulumi:"testBaseAccountName"`
}


type CustomerEventArgs struct {
	CustomerEventName   pulumi.StringPtrInput
	EventName           pulumi.StringInput
	Receivers           NotificationEventReceiverArrayInput
	ResourceGroupName   pulumi.StringInput
	TestBaseAccountName pulumi.StringInput
}

func (CustomerEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerEventArgs)(nil)).Elem()
}

type CustomerEventInput interface {
	pulumi.Input

	ToCustomerEventOutput() CustomerEventOutput
	ToCustomerEventOutputWithContext(ctx context.Context) CustomerEventOutput
}

func (*CustomerEvent) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerEvent)(nil))
}

func (i *CustomerEvent) ToCustomerEventOutput() CustomerEventOutput {
	return i.ToCustomerEventOutputWithContext(context.Background())
}

func (i *CustomerEvent) ToCustomerEventOutputWithContext(ctx context.Context) CustomerEventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerEventOutput)
}

type CustomerEventOutput struct{ *pulumi.OutputState }

func (CustomerEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerEvent)(nil))
}

func (o CustomerEventOutput) ToCustomerEventOutput() CustomerEventOutput {
	return o
}

func (o CustomerEventOutput) ToCustomerEventOutputWithContext(ctx context.Context) CustomerEventOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomerEventOutput{})
}
