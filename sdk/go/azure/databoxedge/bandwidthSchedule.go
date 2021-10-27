


package databoxedge

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
			Type: pulumi.String("azure-nextgen:databoxedge:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190301:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190701:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190801:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200501preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20201201:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210601:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210601preview:BandwidthSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource BandwidthSchedule
	err := ctx.RegisterResource("azure-native:databoxedge:BandwidthSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBandwidthSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthScheduleState, opts ...pulumi.ResourceOption) (*BandwidthSchedule, error) {
	var resource BandwidthSchedule
	err := ctx.ReadResource("azure-native:databoxedge:BandwidthSchedule", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*BandwidthSchedule)(nil))
}

func (i *BandwidthSchedule) ToBandwidthScheduleOutput() BandwidthScheduleOutput {
	return i.ToBandwidthScheduleOutputWithContext(context.Background())
}

func (i *BandwidthSchedule) ToBandwidthScheduleOutputWithContext(ctx context.Context) BandwidthScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthScheduleOutput)
}

type BandwidthScheduleOutput struct{ *pulumi.OutputState }

func (BandwidthScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthSchedule)(nil))
}

func (o BandwidthScheduleOutput) ToBandwidthScheduleOutput() BandwidthScheduleOutput {
	return o
}

func (o BandwidthScheduleOutput) ToBandwidthScheduleOutputWithContext(ctx context.Context) BandwidthScheduleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthScheduleInput)(nil)).Elem(), &BandwidthSchedule{})
	pulumi.RegisterOutputType(BandwidthScheduleOutput{})
}
