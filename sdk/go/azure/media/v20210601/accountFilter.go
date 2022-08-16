


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountFilter struct {
	pulumi.CustomResourceState

	FirstQuality          FirstQualityResponsePtrOutput           `pulumi:"firstQuality"`
	Name                  pulumi.StringOutput                     `pulumi:"name"`
	PresentationTimeRange PresentationTimeRangeResponsePtrOutput  `pulumi:"presentationTimeRange"`
	SystemData            SystemDataResponseOutput                `pulumi:"systemData"`
	Tracks                FilterTrackSelectionResponseArrayOutput `pulumi:"tracks"`
	Type                  pulumi.StringOutput                     `pulumi:"type"`
}


func NewAccountFilter(ctx *pulumi.Context,
	name string, args *AccountFilterArgs, opts ...pulumi.ResourceOption) (*AccountFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:AccountFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:AccountFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:AccountFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:AccountFilter"),
		},
	})
	opts = append(opts, aliases)
	var resource AccountFilter
	err := ctx.RegisterResource("azure-native:media/v20210601:AccountFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccountFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountFilterState, opts ...pulumi.ResourceOption) (*AccountFilter, error) {
	var resource AccountFilter
	err := ctx.ReadResource("azure-native:media/v20210601:AccountFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accountFilterState struct {
}

type AccountFilterState struct {
}

func (AccountFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountFilterState)(nil)).Elem()
}

type accountFilterArgs struct {
	AccountName           string                 `pulumi:"accountName"`
	FilterName            *string                `pulumi:"filterName"`
	FirstQuality          *FirstQuality          `pulumi:"firstQuality"`
	PresentationTimeRange *PresentationTimeRange `pulumi:"presentationTimeRange"`
	ResourceGroupName     string                 `pulumi:"resourceGroupName"`
	Tracks                []FilterTrackSelection `pulumi:"tracks"`
}


type AccountFilterArgs struct {
	AccountName           pulumi.StringInput
	FilterName            pulumi.StringPtrInput
	FirstQuality          FirstQualityPtrInput
	PresentationTimeRange PresentationTimeRangePtrInput
	ResourceGroupName     pulumi.StringInput
	Tracks                FilterTrackSelectionArrayInput
}

func (AccountFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountFilterArgs)(nil)).Elem()
}

type AccountFilterInput interface {
	pulumi.Input

	ToAccountFilterOutput() AccountFilterOutput
	ToAccountFilterOutputWithContext(ctx context.Context) AccountFilterOutput
}

func (*AccountFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountFilter)(nil)).Elem()
}

func (i *AccountFilter) ToAccountFilterOutput() AccountFilterOutput {
	return i.ToAccountFilterOutputWithContext(context.Background())
}

func (i *AccountFilter) ToAccountFilterOutputWithContext(ctx context.Context) AccountFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountFilterOutput)
}

type AccountFilterOutput struct{ *pulumi.OutputState }

func (AccountFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountFilter)(nil)).Elem()
}

func (o AccountFilterOutput) ToAccountFilterOutput() AccountFilterOutput {
	return o
}

func (o AccountFilterOutput) ToAccountFilterOutputWithContext(ctx context.Context) AccountFilterOutput {
	return o
}

func (o AccountFilterOutput) FirstQuality() FirstQualityResponsePtrOutput {
	return o.ApplyT(func(v *AccountFilter) FirstQualityResponsePtrOutput { return v.FirstQuality }).(FirstQualityResponsePtrOutput)
}

func (o AccountFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountFilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccountFilterOutput) PresentationTimeRange() PresentationTimeRangeResponsePtrOutput {
	return o.ApplyT(func(v *AccountFilter) PresentationTimeRangeResponsePtrOutput { return v.PresentationTimeRange }).(PresentationTimeRangeResponsePtrOutput)
}

func (o AccountFilterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AccountFilter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AccountFilterOutput) Tracks() FilterTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *AccountFilter) FilterTrackSelectionResponseArrayOutput { return v.Tracks }).(FilterTrackSelectionResponseArrayOutput)
}

func (o AccountFilterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountFilter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountFilterOutput{})
}
