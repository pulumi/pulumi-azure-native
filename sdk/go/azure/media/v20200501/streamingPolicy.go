


package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StreamingPolicy struct {
	pulumi.CustomResourceState

	CommonEncryptionCbcs        CommonEncryptionCbcsResponsePtrOutput `pulumi:"commonEncryptionCbcs"`
	CommonEncryptionCenc        CommonEncryptionCencResponsePtrOutput `pulumi:"commonEncryptionCenc"`
	Created                     pulumi.StringOutput                   `pulumi:"created"`
	DefaultContentKeyPolicyName pulumi.StringPtrOutput                `pulumi:"defaultContentKeyPolicyName"`
	EnvelopeEncryption          EnvelopeEncryptionResponsePtrOutput   `pulumi:"envelopeEncryption"`
	Name                        pulumi.StringOutput                   `pulumi:"name"`
	NoEncryption                NoEncryptionResponsePtrOutput         `pulumi:"noEncryption"`
	SystemData                  SystemDataResponseOutput              `pulumi:"systemData"`
	Type                        pulumi.StringOutput                   `pulumi:"type"`
}


func NewStreamingPolicy(ctx *pulumi.Context,
	name string, args *StreamingPolicyArgs, opts ...pulumi.ResourceOption) (*StreamingPolicy, error) {
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
			Type: pulumi.String("azure-native:media:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:StreamingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingPolicy
	err := ctx.RegisterResource("azure-native:media/v20200501:StreamingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStreamingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingPolicyState, opts ...pulumi.ResourceOption) (*StreamingPolicy, error) {
	var resource StreamingPolicy
	err := ctx.ReadResource("azure-native:media/v20200501:StreamingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type streamingPolicyState struct {
}

type StreamingPolicyState struct {
}

func (StreamingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingPolicyState)(nil)).Elem()
}

type streamingPolicyArgs struct {
	AccountName                 string                `pulumi:"accountName"`
	CommonEncryptionCbcs        *CommonEncryptionCbcs `pulumi:"commonEncryptionCbcs"`
	CommonEncryptionCenc        *CommonEncryptionCenc `pulumi:"commonEncryptionCenc"`
	DefaultContentKeyPolicyName *string               `pulumi:"defaultContentKeyPolicyName"`
	EnvelopeEncryption          *EnvelopeEncryption   `pulumi:"envelopeEncryption"`
	NoEncryption                *NoEncryption         `pulumi:"noEncryption"`
	ResourceGroupName           string                `pulumi:"resourceGroupName"`
	StreamingPolicyName         *string               `pulumi:"streamingPolicyName"`
}


type StreamingPolicyArgs struct {
	AccountName                 pulumi.StringInput
	CommonEncryptionCbcs        CommonEncryptionCbcsPtrInput
	CommonEncryptionCenc        CommonEncryptionCencPtrInput
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	EnvelopeEncryption          EnvelopeEncryptionPtrInput
	NoEncryption                NoEncryptionPtrInput
	ResourceGroupName           pulumi.StringInput
	StreamingPolicyName         pulumi.StringPtrInput
}

func (StreamingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingPolicyArgs)(nil)).Elem()
}

type StreamingPolicyInput interface {
	pulumi.Input

	ToStreamingPolicyOutput() StreamingPolicyOutput
	ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput
}

func (*StreamingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicy)(nil)).Elem()
}

func (i *StreamingPolicy) ToStreamingPolicyOutput() StreamingPolicyOutput {
	return i.ToStreamingPolicyOutputWithContext(context.Background())
}

func (i *StreamingPolicy) ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyOutput)
}

type StreamingPolicyOutput struct{ *pulumi.OutputState }

func (StreamingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicy)(nil)).Elem()
}

func (o StreamingPolicyOutput) ToStreamingPolicyOutput() StreamingPolicyOutput {
	return o
}

func (o StreamingPolicyOutput) ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StreamingPolicyOutput{})
}
