


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementGroupDiagnosticSetting struct {
	pulumi.CustomResourceState

	EventHubAuthorizationRuleId pulumi.StringPtrOutput                        `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                pulumi.StringPtrOutput                        `pulumi:"eventHubName"`
	Logs                        ManagementGroupLogSettingsResponseArrayOutput `pulumi:"logs"`
	MarketplacePartnerId        pulumi.StringPtrOutput                        `pulumi:"marketplacePartnerId"`
	Name                        pulumi.StringOutput                           `pulumi:"name"`
	ServiceBusRuleId            pulumi.StringPtrOutput                        `pulumi:"serviceBusRuleId"`
	StorageAccountId            pulumi.StringPtrOutput                        `pulumi:"storageAccountId"`
	SystemData                  SystemDataResponseOutput                      `pulumi:"systemData"`
	Type                        pulumi.StringOutput                           `pulumi:"type"`
	WorkspaceId                 pulumi.StringPtrOutput                        `pulumi:"workspaceId"`
}


func NewManagementGroupDiagnosticSetting(ctx *pulumi.Context,
	name string, args *ManagementGroupDiagnosticSettingArgs, opts ...pulumi.ResourceOption) (*ManagementGroupDiagnosticSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:ManagementGroupDiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200101preview:ManagementGroupDiagnosticSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementGroupDiagnosticSetting
	err := ctx.RegisterResource("azure-native:insights/v20210501preview:ManagementGroupDiagnosticSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementGroupDiagnosticSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupDiagnosticSettingState, opts ...pulumi.ResourceOption) (*ManagementGroupDiagnosticSetting, error) {
	var resource ManagementGroupDiagnosticSetting
	err := ctx.ReadResource("azure-native:insights/v20210501preview:ManagementGroupDiagnosticSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementGroupDiagnosticSettingState struct {
}

type ManagementGroupDiagnosticSettingState struct {
}

func (ManagementGroupDiagnosticSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupDiagnosticSettingState)(nil)).Elem()
}

type managementGroupDiagnosticSettingArgs struct {
	EventHubAuthorizationRuleId *string                      `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string                      `pulumi:"eventHubName"`
	Logs                        []ManagementGroupLogSettings `pulumi:"logs"`
	ManagementGroupId           string                       `pulumi:"managementGroupId"`
	MarketplacePartnerId        *string                      `pulumi:"marketplacePartnerId"`
	Name                        *string                      `pulumi:"name"`
	ServiceBusRuleId            *string                      `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string                      `pulumi:"storageAccountId"`
	WorkspaceId                 *string                      `pulumi:"workspaceId"`
}


type ManagementGroupDiagnosticSettingArgs struct {
	EventHubAuthorizationRuleId pulumi.StringPtrInput
	EventHubName                pulumi.StringPtrInput
	Logs                        ManagementGroupLogSettingsArrayInput
	ManagementGroupId           pulumi.StringInput
	MarketplacePartnerId        pulumi.StringPtrInput
	Name                        pulumi.StringPtrInput
	ServiceBusRuleId            pulumi.StringPtrInput
	StorageAccountId            pulumi.StringPtrInput
	WorkspaceId                 pulumi.StringPtrInput
}

func (ManagementGroupDiagnosticSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupDiagnosticSettingArgs)(nil)).Elem()
}

type ManagementGroupDiagnosticSettingInput interface {
	pulumi.Input

	ToManagementGroupDiagnosticSettingOutput() ManagementGroupDiagnosticSettingOutput
	ToManagementGroupDiagnosticSettingOutputWithContext(ctx context.Context) ManagementGroupDiagnosticSettingOutput
}

func (*ManagementGroupDiagnosticSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupDiagnosticSetting)(nil)).Elem()
}

func (i *ManagementGroupDiagnosticSetting) ToManagementGroupDiagnosticSettingOutput() ManagementGroupDiagnosticSettingOutput {
	return i.ToManagementGroupDiagnosticSettingOutputWithContext(context.Background())
}

func (i *ManagementGroupDiagnosticSetting) ToManagementGroupDiagnosticSettingOutputWithContext(ctx context.Context) ManagementGroupDiagnosticSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupDiagnosticSettingOutput)
}

type ManagementGroupDiagnosticSettingOutput struct{ *pulumi.OutputState }

func (ManagementGroupDiagnosticSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupDiagnosticSetting)(nil)).Elem()
}

func (o ManagementGroupDiagnosticSettingOutput) ToManagementGroupDiagnosticSettingOutput() ManagementGroupDiagnosticSettingOutput {
	return o
}

func (o ManagementGroupDiagnosticSettingOutput) ToManagementGroupDiagnosticSettingOutputWithContext(ctx context.Context) ManagementGroupDiagnosticSettingOutput {
	return o
}

func (o ManagementGroupDiagnosticSettingOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) pulumi.StringPtrOutput { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) pulumi.StringPtrOutput { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) Logs() ManagementGroupLogSettingsResponseArrayOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) ManagementGroupLogSettingsResponseArrayOutput { return v.Logs }).(ManagementGroupLogSettingsResponseArrayOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) MarketplacePartnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) pulumi.StringPtrOutput { return v.MarketplacePartnerId }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) pulumi.StringPtrOutput { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ManagementGroupDiagnosticSettingOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDiagnosticSetting) pulumi.StringPtrOutput { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupDiagnosticSettingOutput{})
}
