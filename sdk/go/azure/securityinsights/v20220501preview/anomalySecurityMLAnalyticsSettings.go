


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AnomalySecurityMLAnalyticsSettings struct {
	pulumi.CustomResourceState

	AnomalySettingsVersion   pulumi.IntPtrOutput                                      `pulumi:"anomalySettingsVersion"`
	AnomalyVersion           pulumi.StringOutput                                      `pulumi:"anomalyVersion"`
	CustomizableObservations pulumi.AnyOutput                                         `pulumi:"customizableObservations"`
	Description              pulumi.StringPtrOutput                                   `pulumi:"description"`
	DisplayName              pulumi.StringOutput                                      `pulumi:"displayName"`
	Enabled                  pulumi.BoolOutput                                        `pulumi:"enabled"`
	Etag                     pulumi.StringPtrOutput                                   `pulumi:"etag"`
	Frequency                pulumi.StringOutput                                      `pulumi:"frequency"`
	IsDefaultSettings        pulumi.BoolOutput                                        `pulumi:"isDefaultSettings"`
	Kind                     pulumi.StringOutput                                      `pulumi:"kind"`
	LastModifiedUtc          pulumi.StringOutput                                      `pulumi:"lastModifiedUtc"`
	Name                     pulumi.StringOutput                                      `pulumi:"name"`
	RequiredDataConnectors   SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput `pulumi:"requiredDataConnectors"`
	SettingsDefinitionId     pulumi.StringPtrOutput                                   `pulumi:"settingsDefinitionId"`
	SettingsStatus           pulumi.StringOutput                                      `pulumi:"settingsStatus"`
	SystemData               SystemDataResponseOutput                                 `pulumi:"systemData"`
	Tactics                  pulumi.StringArrayOutput                                 `pulumi:"tactics"`
	Techniques               pulumi.StringArrayOutput                                 `pulumi:"techniques"`
	Type                     pulumi.StringOutput                                      `pulumi:"type"`
}


func NewAnomalySecurityMLAnalyticsSettings(ctx *pulumi.Context,
	name string, args *AnomalySecurityMLAnalyticsSettingsArgs, opts ...pulumi.ResourceOption) (*AnomalySecurityMLAnalyticsSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnomalyVersion == nil {
		return nil, errors.New("invalid value for required argument 'AnomalyVersion'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.IsDefaultSettings == nil {
		return nil, errors.New("invalid value for required argument 'IsDefaultSettings'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SettingsStatus == nil {
		return nil, errors.New("invalid value for required argument 'SettingsStatus'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Anomaly")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AnomalySecurityMLAnalyticsSettings"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AnomalySecurityMLAnalyticsSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource AnomalySecurityMLAnalyticsSettings
	err := ctx.RegisterResource("azure-native:securityinsights/v20220501preview:AnomalySecurityMLAnalyticsSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAnomalySecurityMLAnalyticsSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnomalySecurityMLAnalyticsSettingsState, opts ...pulumi.ResourceOption) (*AnomalySecurityMLAnalyticsSettings, error) {
	var resource AnomalySecurityMLAnalyticsSettings
	err := ctx.ReadResource("azure-native:securityinsights/v20220501preview:AnomalySecurityMLAnalyticsSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type anomalySecurityMLAnalyticsSettingsState struct {
}

type AnomalySecurityMLAnalyticsSettingsState struct {
}

func (AnomalySecurityMLAnalyticsSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalySecurityMLAnalyticsSettingsState)(nil)).Elem()
}

type anomalySecurityMLAnalyticsSettingsArgs struct {
	AnomalySettingsVersion   *int                                    `pulumi:"anomalySettingsVersion"`
	AnomalyVersion           string                                  `pulumi:"anomalyVersion"`
	CustomizableObservations interface{}                             `pulumi:"customizableObservations"`
	Description              *string                                 `pulumi:"description"`
	DisplayName              string                                  `pulumi:"displayName"`
	Enabled                  bool                                    `pulumi:"enabled"`
	Frequency                string                                  `pulumi:"frequency"`
	IsDefaultSettings        bool                                    `pulumi:"isDefaultSettings"`
	Kind                     string                                  `pulumi:"kind"`
	RequiredDataConnectors   []SecurityMLAnalyticsSettingsDataSource `pulumi:"requiredDataConnectors"`
	ResourceGroupName        string                                  `pulumi:"resourceGroupName"`
	SettingsDefinitionId     *string                                 `pulumi:"settingsDefinitionId"`
	SettingsResourceName     *string                                 `pulumi:"settingsResourceName"`
	SettingsStatus           string                                  `pulumi:"settingsStatus"`
	Tactics                  []string                                `pulumi:"tactics"`
	Techniques               []string                                `pulumi:"techniques"`
	WorkspaceName            string                                  `pulumi:"workspaceName"`
}


type AnomalySecurityMLAnalyticsSettingsArgs struct {
	AnomalySettingsVersion   pulumi.IntPtrInput
	AnomalyVersion           pulumi.StringInput
	CustomizableObservations pulumi.Input
	Description              pulumi.StringPtrInput
	DisplayName              pulumi.StringInput
	Enabled                  pulumi.BoolInput
	Frequency                pulumi.StringInput
	IsDefaultSettings        pulumi.BoolInput
	Kind                     pulumi.StringInput
	RequiredDataConnectors   SecurityMLAnalyticsSettingsDataSourceArrayInput
	ResourceGroupName        pulumi.StringInput
	SettingsDefinitionId     pulumi.StringPtrInput
	SettingsResourceName     pulumi.StringPtrInput
	SettingsStatus           pulumi.StringInput
	Tactics                  pulumi.StringArrayInput
	Techniques               pulumi.StringArrayInput
	WorkspaceName            pulumi.StringInput
}

func (AnomalySecurityMLAnalyticsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalySecurityMLAnalyticsSettingsArgs)(nil)).Elem()
}

type AnomalySecurityMLAnalyticsSettingsInput interface {
	pulumi.Input

	ToAnomalySecurityMLAnalyticsSettingsOutput() AnomalySecurityMLAnalyticsSettingsOutput
	ToAnomalySecurityMLAnalyticsSettingsOutputWithContext(ctx context.Context) AnomalySecurityMLAnalyticsSettingsOutput
}

func (*AnomalySecurityMLAnalyticsSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalySecurityMLAnalyticsSettings)(nil)).Elem()
}

func (i *AnomalySecurityMLAnalyticsSettings) ToAnomalySecurityMLAnalyticsSettingsOutput() AnomalySecurityMLAnalyticsSettingsOutput {
	return i.ToAnomalySecurityMLAnalyticsSettingsOutputWithContext(context.Background())
}

func (i *AnomalySecurityMLAnalyticsSettings) ToAnomalySecurityMLAnalyticsSettingsOutputWithContext(ctx context.Context) AnomalySecurityMLAnalyticsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomalySecurityMLAnalyticsSettingsOutput)
}

type AnomalySecurityMLAnalyticsSettingsOutput struct{ *pulumi.OutputState }

func (AnomalySecurityMLAnalyticsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalySecurityMLAnalyticsSettings)(nil)).Elem()
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) ToAnomalySecurityMLAnalyticsSettingsOutput() AnomalySecurityMLAnalyticsSettingsOutput {
	return o
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) ToAnomalySecurityMLAnalyticsSettingsOutputWithContext(ctx context.Context) AnomalySecurityMLAnalyticsSettingsOutput {
	return o
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) AnomalySettingsVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.IntPtrOutput { return v.AnomalySettingsVersion }).(pulumi.IntPtrOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) AnomalyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.AnomalyVersion }).(pulumi.StringOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) CustomizableObservations() pulumi.AnyOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.AnyOutput { return v.CustomizableObservations }).(pulumi.AnyOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.Frequency }).(pulumi.StringOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) IsDefaultSettings() pulumi.BoolOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.BoolOutput { return v.IsDefaultSettings }).(pulumi.BoolOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) RequiredDataConnectors() SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput {
		return v.RequiredDataConnectors
	}).(SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) SettingsDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringPtrOutput { return v.SettingsDefinitionId }).(pulumi.StringPtrOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) SettingsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.SettingsStatus }).(pulumi.StringOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

func (o AnomalySecurityMLAnalyticsSettingsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySecurityMLAnalyticsSettings) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AnomalySecurityMLAnalyticsSettingsOutput{})
}
