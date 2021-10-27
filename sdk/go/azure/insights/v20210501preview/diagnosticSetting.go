


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiagnosticSetting struct {
	pulumi.CustomResourceState

	EventHubAuthorizationRuleId pulumi.StringPtrOutput            `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                pulumi.StringPtrOutput            `pulumi:"eventHubName"`
	LogAnalyticsDestinationType pulumi.StringPtrOutput            `pulumi:"logAnalyticsDestinationType"`
	Logs                        LogSettingsResponseArrayOutput    `pulumi:"logs"`
	MarketplacePartnerId        pulumi.StringPtrOutput            `pulumi:"marketplacePartnerId"`
	Metrics                     MetricSettingsResponseArrayOutput `pulumi:"metrics"`
	Name                        pulumi.StringOutput               `pulumi:"name"`
	ServiceBusRuleId            pulumi.StringPtrOutput            `pulumi:"serviceBusRuleId"`
	StorageAccountId            pulumi.StringPtrOutput            `pulumi:"storageAccountId"`
	SystemData                  SystemDataResponseOutput          `pulumi:"systemData"`
	Type                        pulumi.StringOutput               `pulumi:"type"`
	WorkspaceId                 pulumi.StringPtrOutput            `pulumi:"workspaceId"`
}


func NewDiagnosticSetting(ctx *pulumi.Context,
	name string, args *DiagnosticSettingArgs, opts ...pulumi.ResourceOption) (*DiagnosticSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20210501preview:DiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights:DiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:DiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20170501preview:DiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20170501preview:DiagnosticSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource DiagnosticSetting
	err := ctx.RegisterResource("azure-native:insights/v20210501preview:DiagnosticSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiagnosticSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticSettingState, opts ...pulumi.ResourceOption) (*DiagnosticSetting, error) {
	var resource DiagnosticSetting
	err := ctx.ReadResource("azure-native:insights/v20210501preview:DiagnosticSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diagnosticSettingState struct {
}

type DiagnosticSettingState struct {
}

func (DiagnosticSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticSettingState)(nil)).Elem()
}

type diagnosticSettingArgs struct {
	EventHubAuthorizationRuleId *string          `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string          `pulumi:"eventHubName"`
	LogAnalyticsDestinationType *string          `pulumi:"logAnalyticsDestinationType"`
	Logs                        []LogSettings    `pulumi:"logs"`
	MarketplacePartnerId        *string          `pulumi:"marketplacePartnerId"`
	Metrics                     []MetricSettings `pulumi:"metrics"`
	Name                        *string          `pulumi:"name"`
	ResourceUri                 string           `pulumi:"resourceUri"`
	ServiceBusRuleId            *string          `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string          `pulumi:"storageAccountId"`
	WorkspaceId                 *string          `pulumi:"workspaceId"`
}


type DiagnosticSettingArgs struct {
	EventHubAuthorizationRuleId pulumi.StringPtrInput
	EventHubName                pulumi.StringPtrInput
	LogAnalyticsDestinationType pulumi.StringPtrInput
	Logs                        LogSettingsArrayInput
	MarketplacePartnerId        pulumi.StringPtrInput
	Metrics                     MetricSettingsArrayInput
	Name                        pulumi.StringPtrInput
	ResourceUri                 pulumi.StringInput
	ServiceBusRuleId            pulumi.StringPtrInput
	StorageAccountId            pulumi.StringPtrInput
	WorkspaceId                 pulumi.StringPtrInput
}

func (DiagnosticSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticSettingArgs)(nil)).Elem()
}

type DiagnosticSettingInput interface {
	pulumi.Input

	ToDiagnosticSettingOutput() DiagnosticSettingOutput
	ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput
}

func (*DiagnosticSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticSetting)(nil))
}

func (i *DiagnosticSetting) ToDiagnosticSettingOutput() DiagnosticSettingOutput {
	return i.ToDiagnosticSettingOutputWithContext(context.Background())
}

func (i *DiagnosticSetting) ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticSettingOutput)
}

type DiagnosticSettingOutput struct{ *pulumi.OutputState }

func (DiagnosticSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticSetting)(nil))
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingOutput() DiagnosticSettingOutput {
	return o
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiagnosticSettingInput)(nil)).Elem(), &DiagnosticSetting{})
	pulumi.RegisterOutputType(DiagnosticSettingOutput{})
}
