


package security

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
			Type: pulumi.String("azure-native:security/v20150601preview:AdaptiveApplicationControl"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:AdaptiveApplicationControl"),
		},
	})
	opts = append(opts, aliases)
	var resource AdaptiveApplicationControl
	err := ctx.RegisterResource("azure-native:security:AdaptiveApplicationControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAdaptiveApplicationControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdaptiveApplicationControlState, opts ...pulumi.ResourceOption) (*AdaptiveApplicationControl, error) {
	var resource AdaptiveApplicationControl
	err := ctx.ReadResource("azure-native:security:AdaptiveApplicationControl", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**AdaptiveApplicationControl)(nil)).Elem()
}

func (i *AdaptiveApplicationControl) ToAdaptiveApplicationControlOutput() AdaptiveApplicationControlOutput {
	return i.ToAdaptiveApplicationControlOutputWithContext(context.Background())
}

func (i *AdaptiveApplicationControl) ToAdaptiveApplicationControlOutputWithContext(ctx context.Context) AdaptiveApplicationControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdaptiveApplicationControlOutput)
}

type AdaptiveApplicationControlOutput struct{ *pulumi.OutputState }

func (AdaptiveApplicationControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdaptiveApplicationControl)(nil)).Elem()
}

func (o AdaptiveApplicationControlOutput) ToAdaptiveApplicationControlOutput() AdaptiveApplicationControlOutput {
	return o
}

func (o AdaptiveApplicationControlOutput) ToAdaptiveApplicationControlOutputWithContext(ctx context.Context) AdaptiveApplicationControlOutput {
	return o
}

func (o AdaptiveApplicationControlOutput) ConfigurationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.ConfigurationStatus }).(pulumi.StringOutput)
}

func (o AdaptiveApplicationControlOutput) EnforcementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringPtrOutput { return v.EnforcementMode }).(pulumi.StringPtrOutput)
}

func (o AdaptiveApplicationControlOutput) Issues() AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
		return v.Issues
	}).(AdaptiveApplicationControlIssueSummaryResponseArrayOutput)
}

func (o AdaptiveApplicationControlOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AdaptiveApplicationControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AdaptiveApplicationControlOutput) PathRecommendations() PathRecommendationResponseArrayOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) PathRecommendationResponseArrayOutput {
		return v.PathRecommendations
	}).(PathRecommendationResponseArrayOutput)
}

func (o AdaptiveApplicationControlOutput) ProtectionMode() ProtectionModeResponsePtrOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) ProtectionModeResponsePtrOutput { return v.ProtectionMode }).(ProtectionModeResponsePtrOutput)
}

func (o AdaptiveApplicationControlOutput) RecommendationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.RecommendationStatus }).(pulumi.StringOutput)
}

func (o AdaptiveApplicationControlOutput) SourceSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.SourceSystem }).(pulumi.StringOutput)
}

func (o AdaptiveApplicationControlOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AdaptiveApplicationControlOutput) VmRecommendations() VmRecommendationResponseArrayOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) VmRecommendationResponseArrayOutput { return v.VmRecommendations }).(VmRecommendationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AdaptiveApplicationControlOutput{})
}
