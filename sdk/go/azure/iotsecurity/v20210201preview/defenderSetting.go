


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DefenderSetting struct {
	pulumi.CustomResourceState

	DeviceQuota                  pulumi.IntOutput                                       `pulumi:"deviceQuota"`
	EvaluationEndTime            pulumi.StringOutput                                    `pulumi:"evaluationEndTime"`
	MdeIntegration               DefenderSettingsPropertiesResponseMdeIntegrationOutput `pulumi:"mdeIntegration"`
	Name                         pulumi.StringOutput                                    `pulumi:"name"`
	OnboardingKind               pulumi.StringOutput                                    `pulumi:"onboardingKind"`
	SentinelWorkspaceResourceIds pulumi.StringArrayOutput                               `pulumi:"sentinelWorkspaceResourceIds"`
	Type                         pulumi.StringOutput                                    `pulumi:"type"`
}


func NewDefenderSetting(ctx *pulumi.Context,
	name string, args *DefenderSettingArgs, opts ...pulumi.ResourceOption) (*DefenderSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceQuota == nil {
		return nil, errors.New("invalid value for required argument 'DeviceQuota'")
	}
	if args.MdeIntegration == nil {
		return nil, errors.New("invalid value for required argument 'MdeIntegration'")
	}
	if args.OnboardingKind == nil {
		return nil, errors.New("invalid value for required argument 'OnboardingKind'")
	}
	if args.SentinelWorkspaceResourceIds == nil {
		return nil, errors.New("invalid value for required argument 'SentinelWorkspaceResourceIds'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:iotsecurity/v20210201preview:DefenderSetting"),
		},
		{
			Type: pulumi.String("azure-native:iotsecurity:DefenderSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotsecurity:DefenderSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource DefenderSetting
	err := ctx.RegisterResource("azure-native:iotsecurity/v20210201preview:DefenderSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDefenderSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefenderSettingState, opts ...pulumi.ResourceOption) (*DefenderSetting, error) {
	var resource DefenderSetting
	err := ctx.ReadResource("azure-native:iotsecurity/v20210201preview:DefenderSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type defenderSettingState struct {
}

type DefenderSettingState struct {
}

func (DefenderSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*defenderSettingState)(nil)).Elem()
}

type defenderSettingArgs struct {
	DeviceQuota                  int                                      `pulumi:"deviceQuota"`
	MdeIntegration               DefenderSettingsPropertiesMdeIntegration `pulumi:"mdeIntegration"`
	OnboardingKind               string                                   `pulumi:"onboardingKind"`
	SentinelWorkspaceResourceIds []string                                 `pulumi:"sentinelWorkspaceResourceIds"`
}


type DefenderSettingArgs struct {
	DeviceQuota                  pulumi.IntInput
	MdeIntegration               DefenderSettingsPropertiesMdeIntegrationInput
	OnboardingKind               pulumi.StringInput
	SentinelWorkspaceResourceIds pulumi.StringArrayInput
}

func (DefenderSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defenderSettingArgs)(nil)).Elem()
}

type DefenderSettingInput interface {
	pulumi.Input

	ToDefenderSettingOutput() DefenderSettingOutput
	ToDefenderSettingOutputWithContext(ctx context.Context) DefenderSettingOutput
}

func (*DefenderSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderSetting)(nil))
}

func (i *DefenderSetting) ToDefenderSettingOutput() DefenderSettingOutput {
	return i.ToDefenderSettingOutputWithContext(context.Background())
}

func (i *DefenderSetting) ToDefenderSettingOutputWithContext(ctx context.Context) DefenderSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderSettingOutput)
}

type DefenderSettingOutput struct{ *pulumi.OutputState }

func (DefenderSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderSetting)(nil))
}

func (o DefenderSettingOutput) ToDefenderSettingOutput() DefenderSettingOutput {
	return o
}

func (o DefenderSettingOutput) ToDefenderSettingOutputWithContext(ctx context.Context) DefenderSettingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderSettingInput)(nil)).Elem(), &DefenderSetting{})
	pulumi.RegisterOutputType(DefenderSettingOutput{})
}
