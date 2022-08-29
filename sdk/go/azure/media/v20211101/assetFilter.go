


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssetFilter struct {
	pulumi.CustomResourceState

	FirstQuality          FirstQualityResponsePtrOutput           `pulumi:"firstQuality"`
	Name                  pulumi.StringOutput                     `pulumi:"name"`
	PresentationTimeRange PresentationTimeRangeResponsePtrOutput  `pulumi:"presentationTimeRange"`
	SystemData            SystemDataResponseOutput                `pulumi:"systemData"`
	Tracks                FilterTrackSelectionResponseArrayOutput `pulumi:"tracks"`
	Type                  pulumi.StringOutput                     `pulumi:"type"`
}


func NewAssetFilter(ctx *pulumi.Context,
	name string, args *AssetFilterArgs, opts ...pulumi.ResourceOption) (*AssetFilter, error) {
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
			Type: pulumi.String("azure-native:media:AssetFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:AssetFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:AssetFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:AssetFilter"),
		},
	})
	opts = append(opts, aliases)
	var resource AssetFilter
	err := ctx.RegisterResource("azure-native:media/v20211101:AssetFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetFilterState, opts ...pulumi.ResourceOption) (*AssetFilter, error) {
	var resource AssetFilter
	err := ctx.ReadResource("azure-native:media/v20211101:AssetFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type assetFilterState struct {
}

type AssetFilterState struct {
}

func (AssetFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetFilterState)(nil)).Elem()
}

type assetFilterArgs struct {
	AccountName           string                 `pulumi:"accountName"`
	AssetName             string                 `pulumi:"assetName"`
	FilterName            *string                `pulumi:"filterName"`
	FirstQuality          *FirstQuality          `pulumi:"firstQuality"`
	PresentationTimeRange *PresentationTimeRange `pulumi:"presentationTimeRange"`
	ResourceGroupName     string                 `pulumi:"resourceGroupName"`
	Tracks                []FilterTrackSelection `pulumi:"tracks"`
}


type AssetFilterArgs struct {
	AccountName           pulumi.StringInput
	AssetName             pulumi.StringInput
	FilterName            pulumi.StringPtrInput
	FirstQuality          FirstQualityPtrInput
	PresentationTimeRange PresentationTimeRangePtrInput
	ResourceGroupName     pulumi.StringInput
	Tracks                FilterTrackSelectionArrayInput
}

func (AssetFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetFilterArgs)(nil)).Elem()
}

type AssetFilterInput interface {
	pulumi.Input

	ToAssetFilterOutput() AssetFilterOutput
	ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput
}

func (*AssetFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetFilter)(nil)).Elem()
}

func (i *AssetFilter) ToAssetFilterOutput() AssetFilterOutput {
	return i.ToAssetFilterOutputWithContext(context.Background())
}

func (i *AssetFilter) ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFilterOutput)
}

type AssetFilterOutput struct{ *pulumi.OutputState }

func (AssetFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetFilter)(nil)).Elem()
}

func (o AssetFilterOutput) ToAssetFilterOutput() AssetFilterOutput {
	return o
}

func (o AssetFilterOutput) ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput {
	return o
}

func (o AssetFilterOutput) FirstQuality() FirstQualityResponsePtrOutput {
	return o.ApplyT(func(v *AssetFilter) FirstQualityResponsePtrOutput { return v.FirstQuality }).(FirstQualityResponsePtrOutput)
}

func (o AssetFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssetFilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AssetFilterOutput) PresentationTimeRange() PresentationTimeRangeResponsePtrOutput {
	return o.ApplyT(func(v *AssetFilter) PresentationTimeRangeResponsePtrOutput { return v.PresentationTimeRange }).(PresentationTimeRangeResponsePtrOutput)
}

func (o AssetFilterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AssetFilter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AssetFilterOutput) Tracks() FilterTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *AssetFilter) FilterTrackSelectionResponseArrayOutput { return v.Tracks }).(FilterTrackSelectionResponseArrayOutput)
}

func (o AssetFilterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AssetFilter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetFilterOutput{})
}
