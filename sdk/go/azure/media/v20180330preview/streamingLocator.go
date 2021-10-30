


package v20180330preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StreamingLocator struct {
	pulumi.CustomResourceState

	AssetName                   pulumi.StringOutput                                      `pulumi:"assetName"`
	ContentKeys                 StreamingLocatorUserDefinedContentKeyResponseArrayOutput `pulumi:"contentKeys"`
	Created                     pulumi.StringOutput                                      `pulumi:"created"`
	DefaultContentKeyPolicyName pulumi.StringPtrOutput                                   `pulumi:"defaultContentKeyPolicyName"`
	EndTime                     pulumi.StringPtrOutput                                   `pulumi:"endTime"`
	Name                        pulumi.StringOutput                                      `pulumi:"name"`
	StartTime                   pulumi.StringPtrOutput                                   `pulumi:"startTime"`
	StreamingLocatorId          pulumi.StringPtrOutput                                   `pulumi:"streamingLocatorId"`
	StreamingPolicyName         pulumi.StringOutput                                      `pulumi:"streamingPolicyName"`
	Type                        pulumi.StringOutput                                      `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:media/v20180330preview:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-nextgen:media:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180701:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20210601:StreamingLocator"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingLocator
	err := ctx.RegisterResource("azure-native:media/v20180330preview:StreamingLocator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStreamingLocator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingLocatorState, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
	var resource StreamingLocator
	err := ctx.ReadResource("azure-native:media/v20180330preview:StreamingLocator", name, id, state, &resource, opts...)
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
	AccountName                 string                                  `pulumi:"accountName"`
	AssetName                   string                                  `pulumi:"assetName"`
	ContentKeys                 []StreamingLocatorUserDefinedContentKey `pulumi:"contentKeys"`
	DefaultContentKeyPolicyName *string                                 `pulumi:"defaultContentKeyPolicyName"`
	EndTime                     *string                                 `pulumi:"endTime"`
	ResourceGroupName           string                                  `pulumi:"resourceGroupName"`
	StartTime                   *string                                 `pulumi:"startTime"`
	StreamingLocatorId          *string                                 `pulumi:"streamingLocatorId"`
	StreamingLocatorName        *string                                 `pulumi:"streamingLocatorName"`
	StreamingPolicyName         string                                  `pulumi:"streamingPolicyName"`
}


type StreamingLocatorArgs struct {
	AccountName                 pulumi.StringInput
	AssetName                   pulumi.StringInput
	ContentKeys                 StreamingLocatorUserDefinedContentKeyArrayInput
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	EndTime                     pulumi.StringPtrInput
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
	return reflect.TypeOf((*StreamingLocator)(nil))
}

func (i *StreamingLocator) ToStreamingLocatorOutput() StreamingLocatorOutput {
	return i.ToStreamingLocatorOutputWithContext(context.Background())
}

func (i *StreamingLocator) ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorOutput)
}

type StreamingLocatorOutput struct{ *pulumi.OutputState }

func (StreamingLocatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingLocator)(nil))
}

func (o StreamingLocatorOutput) ToStreamingLocatorOutput() StreamingLocatorOutput {
	return o
}

func (o StreamingLocatorOutput) ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StreamingLocatorOutput{})
}
