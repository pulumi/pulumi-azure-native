


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SentinelOnboardingState struct {
	pulumi.CustomResourceState

	CustomerManagedKey pulumi.BoolPtrOutput     `pulumi:"customerManagedKey"`
	Etag               pulumi.StringPtrOutput   `pulumi:"etag"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewSentinelOnboardingState(ctx *pulumi.Context,
	name string, args *SentinelOnboardingStateArgs, opts ...pulumi.ResourceOption) (*SentinelOnboardingState, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:SentinelOnboardingState"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:SentinelOnboardingState"),
		},
	})
	opts = append(opts, aliases)
	var resource SentinelOnboardingState
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:SentinelOnboardingState", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSentinelOnboardingState(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SentinelOnboardingStateState, opts ...pulumi.ResourceOption) (*SentinelOnboardingState, error) {
	var resource SentinelOnboardingState
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:SentinelOnboardingState", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sentinelOnboardingStateState struct {
}

type SentinelOnboardingStateState struct {
}

func (SentinelOnboardingStateState) ElementType() reflect.Type {
	return reflect.TypeOf((*sentinelOnboardingStateState)(nil)).Elem()
}

type sentinelOnboardingStateArgs struct {
	CustomerManagedKey          *bool   `pulumi:"customerManagedKey"`
	Etag                        *string `pulumi:"etag"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
	SentinelOnboardingStateName *string `pulumi:"sentinelOnboardingStateName"`
	WorkspaceName               string  `pulumi:"workspaceName"`
}


type SentinelOnboardingStateArgs struct {
	CustomerManagedKey          pulumi.BoolPtrInput
	Etag                        pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	SentinelOnboardingStateName pulumi.StringPtrInput
	WorkspaceName               pulumi.StringInput
}

func (SentinelOnboardingStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sentinelOnboardingStateArgs)(nil)).Elem()
}

type SentinelOnboardingStateInput interface {
	pulumi.Input

	ToSentinelOnboardingStateOutput() SentinelOnboardingStateOutput
	ToSentinelOnboardingStateOutputWithContext(ctx context.Context) SentinelOnboardingStateOutput
}

func (*SentinelOnboardingState) ElementType() reflect.Type {
	return reflect.TypeOf((**SentinelOnboardingState)(nil)).Elem()
}

func (i *SentinelOnboardingState) ToSentinelOnboardingStateOutput() SentinelOnboardingStateOutput {
	return i.ToSentinelOnboardingStateOutputWithContext(context.Background())
}

func (i *SentinelOnboardingState) ToSentinelOnboardingStateOutputWithContext(ctx context.Context) SentinelOnboardingStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentinelOnboardingStateOutput)
}

type SentinelOnboardingStateOutput struct{ *pulumi.OutputState }

func (SentinelOnboardingStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SentinelOnboardingState)(nil)).Elem()
}

func (o SentinelOnboardingStateOutput) ToSentinelOnboardingStateOutput() SentinelOnboardingStateOutput {
	return o
}

func (o SentinelOnboardingStateOutput) ToSentinelOnboardingStateOutputWithContext(ctx context.Context) SentinelOnboardingStateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SentinelOnboardingStateOutput{})
}
