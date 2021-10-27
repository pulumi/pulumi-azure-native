


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscriptionDiagnosticSetting struct {
	pulumi.CustomResourceState

	EventHubAuthorizationRuleId pulumi.StringPtrOutput                     `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                pulumi.StringPtrOutput                     `pulumi:"eventHubName"`
	Logs                        SubscriptionLogSettingsResponseArrayOutput `pulumi:"logs"`
	MarketplacePartnerId        pulumi.StringPtrOutput                     `pulumi:"marketplacePartnerId"`
	Name                        pulumi.StringOutput                        `pulumi:"name"`
	ServiceBusRuleId            pulumi.StringPtrOutput                     `pulumi:"serviceBusRuleId"`
	StorageAccountId            pulumi.StringPtrOutput                     `pulumi:"storageAccountId"`
	SystemData                  SystemDataResponseOutput                   `pulumi:"systemData"`
	Type                        pulumi.StringOutput                        `pulumi:"type"`
	WorkspaceId                 pulumi.StringPtrOutput                     `pulumi:"workspaceId"`
}


func NewSubscriptionDiagnosticSetting(ctx *pulumi.Context,
	name string, args *SubscriptionDiagnosticSettingArgs, opts ...pulumi.ResourceOption) (*SubscriptionDiagnosticSetting, error) {
	if args == nil {
		args = &SubscriptionDiagnosticSettingArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20210501preview:SubscriptionDiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights:SubscriptionDiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:SubscriptionDiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20170501preview:SubscriptionDiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20170501preview:SubscriptionDiagnosticSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource SubscriptionDiagnosticSetting
	err := ctx.RegisterResource("azure-native:insights/v20210501preview:SubscriptionDiagnosticSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubscriptionDiagnosticSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionDiagnosticSettingState, opts ...pulumi.ResourceOption) (*SubscriptionDiagnosticSetting, error) {
	var resource SubscriptionDiagnosticSetting
	err := ctx.ReadResource("azure-native:insights/v20210501preview:SubscriptionDiagnosticSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subscriptionDiagnosticSettingState struct {
}

type SubscriptionDiagnosticSettingState struct {
}

func (SubscriptionDiagnosticSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionDiagnosticSettingState)(nil)).Elem()
}

type subscriptionDiagnosticSettingArgs struct {
	EventHubAuthorizationRuleId *string                   `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string                   `pulumi:"eventHubName"`
	Logs                        []SubscriptionLogSettings `pulumi:"logs"`
	MarketplacePartnerId        *string                   `pulumi:"marketplacePartnerId"`
	Name                        *string                   `pulumi:"name"`
	ServiceBusRuleId            *string                   `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string                   `pulumi:"storageAccountId"`
	WorkspaceId                 *string                   `pulumi:"workspaceId"`
}


type SubscriptionDiagnosticSettingArgs struct {
	EventHubAuthorizationRuleId pulumi.StringPtrInput
	EventHubName                pulumi.StringPtrInput
	Logs                        SubscriptionLogSettingsArrayInput
	MarketplacePartnerId        pulumi.StringPtrInput
	Name                        pulumi.StringPtrInput
	ServiceBusRuleId            pulumi.StringPtrInput
	StorageAccountId            pulumi.StringPtrInput
	WorkspaceId                 pulumi.StringPtrInput
}

func (SubscriptionDiagnosticSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionDiagnosticSettingArgs)(nil)).Elem()
}

type SubscriptionDiagnosticSettingInput interface {
	pulumi.Input

	ToSubscriptionDiagnosticSettingOutput() SubscriptionDiagnosticSettingOutput
	ToSubscriptionDiagnosticSettingOutputWithContext(ctx context.Context) SubscriptionDiagnosticSettingOutput
}

func (*SubscriptionDiagnosticSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionDiagnosticSetting)(nil))
}

func (i *SubscriptionDiagnosticSetting) ToSubscriptionDiagnosticSettingOutput() SubscriptionDiagnosticSettingOutput {
	return i.ToSubscriptionDiagnosticSettingOutputWithContext(context.Background())
}

func (i *SubscriptionDiagnosticSetting) ToSubscriptionDiagnosticSettingOutputWithContext(ctx context.Context) SubscriptionDiagnosticSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionDiagnosticSettingOutput)
}

type SubscriptionDiagnosticSettingOutput struct{ *pulumi.OutputState }

func (SubscriptionDiagnosticSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionDiagnosticSetting)(nil))
}

func (o SubscriptionDiagnosticSettingOutput) ToSubscriptionDiagnosticSettingOutput() SubscriptionDiagnosticSettingOutput {
	return o
}

func (o SubscriptionDiagnosticSettingOutput) ToSubscriptionDiagnosticSettingOutputWithContext(ctx context.Context) SubscriptionDiagnosticSettingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionDiagnosticSettingInput)(nil)).Elem(), &SubscriptionDiagnosticSetting{})
	pulumi.RegisterOutputType(SubscriptionDiagnosticSettingOutput{})
}
