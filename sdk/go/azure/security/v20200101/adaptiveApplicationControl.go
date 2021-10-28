


package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdaptiveApplicationControl struct {
	pulumi.CustomResourceState

	ConfigurationStatus  pulumi.StringOutput                                       `pulumi:"configurationStatus"`
	EnforcementMode      pulumi.StringPtrOutput                                    `pulumi:"enforcementMode"`
	Issues               AdaptiveApplicationControlIssueSummaryResponseArrayOutput `pulumi:"issues"`
	Location             pulumi.StringOutput                                       `pulumi:"location"`
	Name                 pulumi.StringOutput                                       `pulumi:"name"`
	PathRecommendations  PathRecommendationResponseArrayOutput                     `pulumi:"pathRecommendations"`
	ProtectionMode       ProtectionModeResponsePtrOutput                           `pulumi:"protectionMode"`
	RecommendationStatus pulumi.StringOutput                                       `pulumi:"recommendationStatus"`
	SourceSystem         pulumi.StringOutput                                       `pulumi:"sourceSystem"`
	Type                 pulumi.StringOutput                                       `pulumi:"type"`
	VmRecommendations    VmRecommendationResponseArrayOutput                       `pulumi:"vmRecommendations"`
}


func NewAdaptiveApplicationControl(ctx *pulumi.Context,
	name string, args *AdaptiveApplicationControlArgs, opts ...pulumi.ResourceOption) (*AdaptiveApplicationControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AscLocation == nil {
		return nil, errors.New("invalid value for required argument 'AscLocation'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security/v20200101:AdaptiveApplicationControl"),
		},
		{
			Type: pulumi.String("azure-native:security:AdaptiveApplicationControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:security:AdaptiveApplicationControl"),
		},
		{
			Type: pulumi.String("azure-native:security/v20150601preview:AdaptiveApplicationControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20150601preview:AdaptiveApplicationControl"),
		},
	})
	opts = append(opts, aliases)
	var resource AdaptiveApplicationControl
	err := ctx.RegisterResource("azure-native:security/v20200101:AdaptiveApplicationControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAdaptiveApplicationControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdaptiveApplicationControlState, opts ...pulumi.ResourceOption) (*AdaptiveApplicationControl, error) {
	var resource AdaptiveApplicationControl
	err := ctx.ReadResource("azure-native:security/v20200101:AdaptiveApplicationControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adaptiveApplicationControlState struct {
}

type AdaptiveApplicationControlState struct {
}

func (AdaptiveApplicationControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*adaptiveApplicationControlState)(nil)).Elem()
}

type adaptiveApplicationControlArgs struct {
	AscLocation         string               `pulumi:"ascLocation"`
	EnforcementMode     *string              `pulumi:"enforcementMode"`
	GroupName           *string              `pulumi:"groupName"`
	PathRecommendations []PathRecommendation `pulumi:"pathRecommendations"`
	ProtectionMode      *ProtectionMode      `pulumi:"protectionMode"`
	VmRecommendations   []VmRecommendation   `pulumi:"vmRecommendations"`
}


type AdaptiveApplicationControlArgs struct {
	AscLocation         pulumi.StringInput
	EnforcementMode     pulumi.StringPtrInput
	GroupName           pulumi.StringPtrInput
	PathRecommendations PathRecommendationArrayInput
	ProtectionMode      ProtectionModePtrInput
	VmRecommendations   VmRecommendationArrayInput
}

func (AdaptiveApplicationControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adaptiveApplicationControlArgs)(nil)).Elem()
}

type AdaptiveApplicationControlInput interface {
	pulumi.Input

	ToAdaptiveApplicationControlOutput() AdaptiveApplicationControlOutput
	ToAdaptiveApplicationControlOutputWithContext(ctx context.Context) AdaptiveApplicationControlOutput
}

func (*AdaptiveApplicationControl) ElementType() reflect.Type {
	return reflect.TypeOf((*AdaptiveApplicationControl)(nil))
}

func (i *AdaptiveApplicationControl) ToAdaptiveApplicationControlOutput() AdaptiveApplicationControlOutput {
	return i.ToAdaptiveApplicationControlOutputWithContext(context.Background())
}

func (i *AdaptiveApplicationControl) ToAdaptiveApplicationControlOutputWithContext(ctx context.Context) AdaptiveApplicationControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdaptiveApplicationControlOutput)
}

type AdaptiveApplicationControlOutput struct{ *pulumi.OutputState }

func (AdaptiveApplicationControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdaptiveApplicationControl)(nil))
}

func (o AdaptiveApplicationControlOutput) ToAdaptiveApplicationControlOutput() AdaptiveApplicationControlOutput {
	return o
}

func (o AdaptiveApplicationControlOutput) ToAdaptiveApplicationControlOutputWithContext(ctx context.Context) AdaptiveApplicationControlOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdaptiveApplicationControlInput)(nil)).Elem(), &AdaptiveApplicationControl{})
	pulumi.RegisterOutputType(AdaptiveApplicationControlOutput{})
}
