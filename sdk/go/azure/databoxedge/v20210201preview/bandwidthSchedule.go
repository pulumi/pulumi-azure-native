


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BandwidthSchedule struct {
	pulumi.CustomResourceState

	Days       pulumi.StringArrayOutput `pulumi:"days"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	RateInMbps pulumi.IntOutput         `pulumi:"rateInMbps"`
	Start      pulumi.StringOutput      `pulumi:"start"`
	Stop       pulumi.StringOutput      `pulumi:"stop"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewBandwidthSchedule(ctx *pulumi.Context,
	name string, args *BandwidthScheduleArgs, opts ...pulumi.ResourceOption) (*BandwidthSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Days == nil {
		return nil, errors.New("invalid value for required argument 'Days'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.RateInMbps == nil {
		return nil, errors.New("invalid value for required argument 'RateInMbps'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	if args.Stop == nil {
		return nil, errors.New("invalid value for required argument 'Stop'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:BandwidthSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource BandwidthSchedule
	err := ctx.RegisterResource("azure-native:databoxedge/v20210201preview:BandwidthSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBandwidthSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthScheduleState, opts ...pulumi.ResourceOption) (*BandwidthSchedule, error) {
	var resource BandwidthSchedule
	err := ctx.ReadResource("azure-native:databoxedge/v20210201preview:BandwidthSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type bandwidthScheduleState struct {
}

type BandwidthScheduleState struct {
}

func (BandwidthScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthScheduleState)(nil)).Elem()
}

type bandwidthScheduleArgs struct {
	Days              []string `pulumi:"days"`
	DeviceName        string   `pulumi:"deviceName"`
	Name              *string  `pulumi:"name"`
	RateInMbps        int      `pulumi:"rateInMbps"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Start             string   `pulumi:"start"`
	Stop              string   `pulumi:"stop"`
}


type BandwidthScheduleArgs struct {
	Days              pulumi.StringArrayInput
	DeviceName        pulumi.StringInput
	Name              pulumi.StringPtrInput
	RateInMbps        pulumi.IntInput
	ResourceGroupName pulumi.StringInput
	Start             pulumi.StringInput
	Stop              pulumi.StringInput
}

func (BandwidthScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthScheduleArgs)(nil)).Elem()
}

type BandwidthScheduleInput interface {
	pulumi.Input

	ToBandwidthScheduleOutput() BandwidthScheduleOutput
	ToBandwidthScheduleOutputWithContext(ctx context.Context) BandwidthScheduleOutput
}

func (*BandwidthSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthSchedule)(nil)).Elem()
}

func (i *BandwidthSchedule) ToBandwidthScheduleOutput() BandwidthScheduleOutput {
	return i.ToBandwidthScheduleOutputWithContext(context.Background())
}

func (i *BandwidthSchedule) ToBandwidthScheduleOutputWithContext(ctx context.Context) BandwidthScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthScheduleOutput)
}

type BandwidthScheduleOutput struct{ *pulumi.OutputState }

func (BandwidthScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthSchedule)(nil)).Elem()
}

func (o BandwidthScheduleOutput) ToBandwidthScheduleOutput() BandwidthScheduleOutput {
	return o
}

func (o BandwidthScheduleOutput) ToBandwidthScheduleOutputWithContext(ctx context.Context) BandwidthScheduleOutput {
	return o
}

func (o BandwidthScheduleOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BandwidthSchedule) pulumi.StringArrayOutput { return v.Days }).(pulumi.StringArrayOutput)
}

func (o BandwidthScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BandwidthScheduleOutput) RateInMbps() pulumi.IntOutput {
	return o.ApplyT(func(v *BandwidthSchedule) pulumi.IntOutput { return v.RateInMbps }).(pulumi.IntOutput)
}

func (o BandwidthScheduleOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthSchedule) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

func (o BandwidthScheduleOutput) Stop() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthSchedule) pulumi.StringOutput { return v.Stop }).(pulumi.StringOutput)
}

func (o BandwidthScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BandwidthSchedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BandwidthScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BandwidthScheduleOutput{})
}
