


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StreamingLocator struct {
	pulumi.CustomResourceState

	AlternativeMediaId          pulumi.StringPtrOutput                        `pulumi:"alternativeMediaId"`
	AssetName                   pulumi.StringOutput                           `pulumi:"assetName"`
	ContentKeys                 StreamingLocatorContentKeyResponseArrayOutput `pulumi:"contentKeys"`
	Created                     pulumi.StringOutput                           `pulumi:"created"`
	DefaultContentKeyPolicyName pulumi.StringPtrOutput                        `pulumi:"defaultContentKeyPolicyName"`
	EndTime                     pulumi.StringPtrOutput                        `pulumi:"endTime"`
	Filters                     pulumi.StringArrayOutput                      `pulumi:"filters"`
	Name                        pulumi.StringOutput                           `pulumi:"name"`
	StartTime                   pulumi.StringPtrOutput                        `pulumi:"startTime"`
	StreamingLocatorId          pulumi.StringPtrOutput                        `pulumi:"streamingLocatorId"`
	StreamingPolicyName         pulumi.StringOutput                           `pulumi:"streamingPolicyName"`
	SystemData                  SystemDataResponseOutput                      `pulumi:"systemData"`
	Type                        pulumi.StringOutput                           `pulumi:"type"`
}


func NewStreamingLocator(ctx *pulumi.Context,
	name string, args *StreamingLocatorArgs, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
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
	if args.StreamingPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'StreamingPolicyName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:StreamingLocator"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingLocator
	err := ctx.RegisterResource("azure-native:media/v20210601:StreamingLocator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStreamingLocator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingLocatorState, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
	var resource StreamingLocator
	err := ctx.ReadResource("azure-native:media/v20210601:StreamingLocator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type streamingLocatorState struct {
}

type StreamingLocatorState struct {
}

func (StreamingLocatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingLocatorState)(nil)).Elem()
}

type streamingLocatorArgs struct {
	AccountName                 string                       `pulumi:"accountName"`
	AlternativeMediaId          *string                      `pulumi:"alternativeMediaId"`
	AssetName                   string                       `pulumi:"assetName"`
	ContentKeys                 []StreamingLocatorContentKey `pulumi:"contentKeys"`
	DefaultContentKeyPolicyName *string                      `pulumi:"defaultContentKeyPolicyName"`
	EndTime                     *string                      `pulumi:"endTime"`
	Filters                     []string                     `pulumi:"filters"`
	ResourceGroupName           string                       `pulumi:"resourceGroupName"`
	StartTime                   *string                      `pulumi:"startTime"`
	StreamingLocatorId          *string                      `pulumi:"streamingLocatorId"`
	StreamingLocatorName        *string                      `pulumi:"streamingLocatorName"`
	StreamingPolicyName         string                       `pulumi:"streamingPolicyName"`
}


type StreamingLocatorArgs struct {
	AccountName                 pulumi.StringInput
	AlternativeMediaId          pulumi.StringPtrInput
	AssetName                   pulumi.StringInput
	ContentKeys                 StreamingLocatorContentKeyArrayInput
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	EndTime                     pulumi.StringPtrInput
	Filters                     pulumi.StringArrayInput
	ResourceGroupName           pulumi.StringInput
	StartTime                   pulumi.StringPtrInput
	StreamingLocatorId          pulumi.StringPtrInput
	StreamingLocatorName        pulumi.StringPtrInput
	StreamingPolicyName         pulumi.StringInput
}

func (StreamingLocatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingLocatorArgs)(nil)).Elem()
}

type StreamingLocatorInput interface {
	pulumi.Input

	ToStreamingLocatorOutput() StreamingLocatorOutput
	ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput
}

func (*StreamingLocator) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingLocator)(nil)).Elem()
}

func (i *StreamingLocator) ToStreamingLocatorOutput() StreamingLocatorOutput {
	return i.ToStreamingLocatorOutputWithContext(context.Background())
}

func (i *StreamingLocator) ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorOutput)
}

type StreamingLocatorOutput struct{ *pulumi.OutputState }

func (StreamingLocatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingLocator)(nil)).Elem()
}

func (o StreamingLocatorOutput) ToStreamingLocatorOutput() StreamingLocatorOutput {
	return o
}

func (o StreamingLocatorOutput) ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput {
	return o
}

func (o StreamingLocatorOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.AssetName }).(pulumi.StringOutput)
}

func (o StreamingLocatorOutput) ContentKeys() StreamingLocatorContentKeyResponseArrayOutput {
	return o.ApplyT(func(v *StreamingLocator) StreamingLocatorContentKeyResponseArrayOutput { return v.ContentKeys }).(StreamingLocatorContentKeyResponseArrayOutput)
}

func (o StreamingLocatorOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o StreamingLocatorOutput) DefaultContentKeyPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.DefaultContentKeyPolicyName }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorOutput) Filters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringArrayOutput { return v.Filters }).(pulumi.StringArrayOutput)
}

func (o StreamingLocatorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StreamingLocatorOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorOutput) StreamingLocatorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.StreamingLocatorId }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorOutput) StreamingPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.StreamingPolicyName }).(pulumi.StringOutput)
}

func (o StreamingLocatorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StreamingLocator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StreamingLocatorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingLocatorOutput{})
}
