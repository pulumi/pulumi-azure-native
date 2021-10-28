


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScheduledSynchronizationSetting struct {
	pulumi.CustomResourceState

	CreatedAt           pulumi.StringOutput      `pulumi:"createdAt"`
	Kind                pulumi.StringOutput      `pulumi:"kind"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput      `pulumi:"provisioningState"`
	RecurrenceInterval  pulumi.StringOutput      `pulumi:"recurrenceInterval"`
	SynchronizationTime pulumi.StringOutput      `pulumi:"synchronizationTime"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	Type                pulumi.StringOutput      `pulumi:"type"`
	UserName            pulumi.StringOutput      `pulumi:"userName"`
}


func NewScheduledSynchronizationSetting(ctx *pulumi.Context,
	name string, args *ScheduledSynchronizationSettingArgs, opts ...pulumi.ResourceOption) (*ScheduledSynchronizationSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.RecurrenceInterval == nil {
		return nil, errors.New("invalid value for required argument 'RecurrenceInterval'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.SynchronizationTime == nil {
		return nil, errors.New("invalid value for required argument 'SynchronizationTime'")
	}
	args.Kind = pulumi.String("ScheduleBased")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/v20201001preview:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20200901:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ScheduledSynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20210801:ScheduledSynchronizationSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledSynchronizationSetting
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:ScheduledSynchronizationSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduledSynchronizationSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledSynchronizationSettingState, opts ...pulumi.ResourceOption) (*ScheduledSynchronizationSetting, error) {
	var resource ScheduledSynchronizationSetting
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:ScheduledSynchronizationSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scheduledSynchronizationSettingState struct {
}

type ScheduledSynchronizationSettingState struct {
}

func (ScheduledSynchronizationSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledSynchronizationSettingState)(nil)).Elem()
}

type scheduledSynchronizationSettingArgs struct {
	AccountName                string  `pulumi:"accountName"`
	Kind                       string  `pulumi:"kind"`
	RecurrenceInterval         string  `pulumi:"recurrenceInterval"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	ShareName                  string  `pulumi:"shareName"`
	SynchronizationSettingName *string `pulumi:"synchronizationSettingName"`
	SynchronizationTime        string  `pulumi:"synchronizationTime"`
}


type ScheduledSynchronizationSettingArgs struct {
	AccountName                pulumi.StringInput
	Kind                       pulumi.StringInput
	RecurrenceInterval         pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	ShareName                  pulumi.StringInput
	SynchronizationSettingName pulumi.StringPtrInput
	SynchronizationTime        pulumi.StringInput
}

func (ScheduledSynchronizationSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledSynchronizationSettingArgs)(nil)).Elem()
}

type ScheduledSynchronizationSettingInput interface {
	pulumi.Input

	ToScheduledSynchronizationSettingOutput() ScheduledSynchronizationSettingOutput
	ToScheduledSynchronizationSettingOutputWithContext(ctx context.Context) ScheduledSynchronizationSettingOutput
}

func (*ScheduledSynchronizationSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledSynchronizationSetting)(nil))
}

func (i *ScheduledSynchronizationSetting) ToScheduledSynchronizationSettingOutput() ScheduledSynchronizationSettingOutput {
	return i.ToScheduledSynchronizationSettingOutputWithContext(context.Background())
}

func (i *ScheduledSynchronizationSetting) ToScheduledSynchronizationSettingOutputWithContext(ctx context.Context) ScheduledSynchronizationSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSynchronizationSettingOutput)
}

type ScheduledSynchronizationSettingOutput struct{ *pulumi.OutputState }

func (ScheduledSynchronizationSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledSynchronizationSetting)(nil))
}

func (o ScheduledSynchronizationSettingOutput) ToScheduledSynchronizationSettingOutput() ScheduledSynchronizationSettingOutput {
	return o
}

func (o ScheduledSynchronizationSettingOutput) ToScheduledSynchronizationSettingOutputWithContext(ctx context.Context) ScheduledSynchronizationSettingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSynchronizationSettingInput)(nil)).Elem(), &ScheduledSynchronizationSetting{})
	pulumi.RegisterOutputType(ScheduledSynchronizationSettingOutput{})
}
