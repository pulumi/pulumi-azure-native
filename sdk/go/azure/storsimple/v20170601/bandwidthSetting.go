


package v20170601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BandwidthSetting struct {
	pulumi.CustomResourceState

	Kind        pulumi.StringPtrOutput               `pulumi:"kind"`
	Name        pulumi.StringOutput                  `pulumi:"name"`
	Schedules   BandwidthScheduleResponseArrayOutput `pulumi:"schedules"`
	Type        pulumi.StringOutput                  `pulumi:"type"`
	VolumeCount pulumi.IntOutput                     `pulumi:"volumeCount"`
}


func NewBandwidthSetting(ctx *pulumi.Context,
	name string, args *BandwidthSettingArgs, opts ...pulumi.ResourceOption) (*BandwidthSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Schedules == nil {
		return nil, errors.New("invalid value for required argument 'Schedules'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple:BandwidthSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource BandwidthSetting
	err := ctx.RegisterResource("azure-native:storsimple/v20170601:BandwidthSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBandwidthSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthSettingState, opts ...pulumi.ResourceOption) (*BandwidthSetting, error) {
	var resource BandwidthSetting
	err := ctx.ReadResource("azure-native:storsimple/v20170601:BandwidthSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type bandwidthSettingState struct {
}

type BandwidthSettingState struct {
}

func (BandwidthSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthSettingState)(nil)).Elem()
}

type bandwidthSettingArgs struct {
	BandwidthSettingName *string             `pulumi:"bandwidthSettingName"`
	Kind                 *Kind               `pulumi:"kind"`
	ManagerName          string              `pulumi:"managerName"`
	ResourceGroupName    string              `pulumi:"resourceGroupName"`
	Schedules            []BandwidthSchedule `pulumi:"schedules"`
}


type BandwidthSettingArgs struct {
	BandwidthSettingName pulumi.StringPtrInput
	Kind                 KindPtrInput
	ManagerName          pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	Schedules            BandwidthScheduleArrayInput
}

func (BandwidthSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthSettingArgs)(nil)).Elem()
}

type BandwidthSettingInput interface {
	pulumi.Input

	ToBandwidthSettingOutput() BandwidthSettingOutput
	ToBandwidthSettingOutputWithContext(ctx context.Context) BandwidthSettingOutput
}

func (*BandwidthSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthSetting)(nil))
}

func (i *BandwidthSetting) ToBandwidthSettingOutput() BandwidthSettingOutput {
	return i.ToBandwidthSettingOutputWithContext(context.Background())
}

func (i *BandwidthSetting) ToBandwidthSettingOutputWithContext(ctx context.Context) BandwidthSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthSettingOutput)
}

type BandwidthSettingOutput struct{ *pulumi.OutputState }

func (BandwidthSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthSetting)(nil))
}

func (o BandwidthSettingOutput) ToBandwidthSettingOutput() BandwidthSettingOutput {
	return o
}

func (o BandwidthSettingOutput) ToBandwidthSettingOutputWithContext(ctx context.Context) BandwidthSettingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BandwidthSettingOutput{})
}
