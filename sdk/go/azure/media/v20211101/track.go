


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Track struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	Track             pulumi.AnyOutput    `pulumi:"track"`
	Type              pulumi.StringOutput `pulumi:"type"`
}


func NewTrack(ctx *pulumi.Context,
	name string, args *TrackArgs, opts ...pulumi.ResourceOption) (*Track, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.AssetName == nil {
		return nil, errors.New("invalid value for required argument 'AssetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:Track"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:Track"),
		},
	})
	opts = append(opts, aliases)
	var resource Track
	err := ctx.RegisterResource("azure-native:media/v20211101:Track", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTrack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrackState, opts ...pulumi.ResourceOption) (*Track, error) {
	var resource Track
	err := ctx.ReadResource("azure-native:media/v20211101:Track", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type trackState struct {
}

type TrackState struct {
}

func (TrackState) ElementType() reflect.Type {
	return reflect.TypeOf((*trackState)(nil)).Elem()
}

type trackArgs struct {
	AccountName       string      `pulumi:"accountName"`
	AssetName         string      `pulumi:"assetName"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	Track             interface{} `pulumi:"track"`
	TrackName         *string     `pulumi:"trackName"`
}


type TrackArgs struct {
	AccountName       pulumi.StringInput
	AssetName         pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Track             pulumi.Input
	TrackName         pulumi.StringPtrInput
}

func (TrackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trackArgs)(nil)).Elem()
}

type TrackInput interface {
	pulumi.Input

	ToTrackOutput() TrackOutput
	ToTrackOutputWithContext(ctx context.Context) TrackOutput
}

func (*Track) ElementType() reflect.Type {
	return reflect.TypeOf((**Track)(nil)).Elem()
}

func (i *Track) ToTrackOutput() TrackOutput {
	return i.ToTrackOutputWithContext(context.Background())
}

func (i *Track) ToTrackOutputWithContext(ctx context.Context) TrackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackOutput)
}

type TrackOutput struct{ *pulumi.OutputState }

func (TrackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Track)(nil)).Elem()
}

func (o TrackOutput) ToTrackOutput() TrackOutput {
	return o
}

func (o TrackOutput) ToTrackOutputWithContext(ctx context.Context) TrackOutput {
	return o
}

func (o TrackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Track) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TrackOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Track) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TrackOutput) Track() pulumi.AnyOutput {
	return o.ApplyT(func(v *Track) pulumi.AnyOutput { return v.Track }).(pulumi.AnyOutput)
}

func (o TrackOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Track) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrackOutput{})
}
