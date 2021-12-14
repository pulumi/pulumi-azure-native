


package v20191101

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

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:Trigger"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:Trigger"),
		},
	})
	opts = append(opts, aliases)
	var resource Trigger
	err := ctx.RegisterResource("azure-native:datashare/v20191101:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("azure-native:datashare/v20191101:Trigger", name, id, state, &resource, opts...)
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
	AccountName           string  `pulumi:"accountName"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	TriggerName           *string `pulumi:"triggerName"`
}


type TriggerArgs struct {
	AccountName           pulumi.StringInput
	Kind                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	TriggerName           pulumi.StringPtrInput
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
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
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
