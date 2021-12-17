


package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProactiveDetectionConfiguration struct {
	pulumi.CustomResourceState

	CustomEmails                   pulumi.StringArrayOutput                                                                    `pulumi:"customEmails"`
	Enabled                        pulumi.BoolPtrOutput                                                                        `pulumi:"enabled"`
	LastUpdatedTime                pulumi.StringPtrOutput                                                                      `pulumi:"lastUpdatedTime"`
	Name                           pulumi.StringPtrOutput                                                                      `pulumi:"name"`
	RuleDefinitions                ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput `pulumi:"ruleDefinitions"`
	SendEmailsToSubscriptionOwners pulumi.BoolPtrOutput                                                                        `pulumi:"sendEmailsToSubscriptionOwners"`
}


func NewProactiveDetectionConfiguration(ctx *pulumi.Context,
	name string, args *ProactiveDetectionConfigurationArgs, opts ...pulumi.ResourceOption) (*ProactiveDetectionConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20150501:ProactiveDetectionConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180501preview:ProactiveDetectionConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource ProactiveDetectionConfiguration
	err := ctx.RegisterResource("azure-native:insights:ProactiveDetectionConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProactiveDetectionConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProactiveDetectionConfigurationState, opts ...pulumi.ResourceOption) (*ProactiveDetectionConfiguration, error) {
	var resource ProactiveDetectionConfiguration
	err := ctx.ReadResource("azure-native:insights:ProactiveDetectionConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type proactiveDetectionConfigurationState struct {
}

type ProactiveDetectionConfigurationState struct {
}

func (ProactiveDetectionConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*proactiveDetectionConfigurationState)(nil)).Elem()
}

type proactiveDetectionConfigurationArgs struct {
	ConfigurationId                *string                                                                     `pulumi:"configurationId"`
	CustomEmails                   []string                                                                    `pulumi:"customEmails"`
	Enabled                        *bool                                                                       `pulumi:"enabled"`
	LastUpdatedTime                *string                                                                     `pulumi:"lastUpdatedTime"`
	Name                           *string                                                                     `pulumi:"name"`
	ResourceGroupName              string                                                                      `pulumi:"resourceGroupName"`
	ResourceName                   string                                                                      `pulumi:"resourceName"`
	RuleDefinitions                *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions `pulumi:"ruleDefinitions"`
	SendEmailsToSubscriptionOwners *bool                                                                       `pulumi:"sendEmailsToSubscriptionOwners"`
}


type ProactiveDetectionConfigurationArgs struct {
	ConfigurationId                pulumi.StringPtrInput
	CustomEmails                   pulumi.StringArrayInput
	Enabled                        pulumi.BoolPtrInput
	LastUpdatedTime                pulumi.StringPtrInput
	Name                           pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	ResourceName                   pulumi.StringInput
	RuleDefinitions                ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrInput
	SendEmailsToSubscriptionOwners pulumi.BoolPtrInput
}

func (ProactiveDetectionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proactiveDetectionConfigurationArgs)(nil)).Elem()
}

type ProactiveDetectionConfigurationInput interface {
	pulumi.Input

	ToProactiveDetectionConfigurationOutput() ProactiveDetectionConfigurationOutput
	ToProactiveDetectionConfigurationOutputWithContext(ctx context.Context) ProactiveDetectionConfigurationOutput
}

func (*ProactiveDetectionConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ProactiveDetectionConfiguration)(nil)).Elem()
}

func (i *ProactiveDetectionConfiguration) ToProactiveDetectionConfigurationOutput() ProactiveDetectionConfigurationOutput {
	return i.ToProactiveDetectionConfigurationOutputWithContext(context.Background())
}

func (i *ProactiveDetectionConfiguration) ToProactiveDetectionConfigurationOutputWithContext(ctx context.Context) ProactiveDetectionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProactiveDetectionConfigurationOutput)
}

type ProactiveDetectionConfigurationOutput struct{ *pulumi.OutputState }

func (ProactiveDetectionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProactiveDetectionConfiguration)(nil)).Elem()
}

func (o ProactiveDetectionConfigurationOutput) ToProactiveDetectionConfigurationOutput() ProactiveDetectionConfigurationOutput {
	return o
}

func (o ProactiveDetectionConfigurationOutput) ToProactiveDetectionConfigurationOutputWithContext(ctx context.Context) ProactiveDetectionConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProactiveDetectionConfigurationOutput{})
}
