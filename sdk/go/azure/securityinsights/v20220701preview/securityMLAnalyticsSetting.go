


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SecurityMLAnalyticsSetting struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewSecurityMLAnalyticsSetting(ctx *pulumi.Context,
	name string, args *SecurityMLAnalyticsSettingArgs, opts ...pulumi.ResourceOption) (*SecurityMLAnalyticsSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:SecurityMLAnalyticsSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:SecurityMLAnalyticsSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:SecurityMLAnalyticsSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:SecurityMLAnalyticsSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:SecurityMLAnalyticsSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:SecurityMLAnalyticsSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:SecurityMLAnalyticsSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:SecurityMLAnalyticsSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityMLAnalyticsSetting
	err := ctx.RegisterResource("azure-native:securityinsights/v20220701preview:SecurityMLAnalyticsSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityMLAnalyticsSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityMLAnalyticsSettingState, opts ...pulumi.ResourceOption) (*SecurityMLAnalyticsSetting, error) {
	var resource SecurityMLAnalyticsSetting
	err := ctx.ReadResource("azure-native:securityinsights/v20220701preview:SecurityMLAnalyticsSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityMLAnalyticsSettingState struct {
}

type SecurityMLAnalyticsSettingState struct {
}

func (SecurityMLAnalyticsSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityMLAnalyticsSettingState)(nil)).Elem()
}

type securityMLAnalyticsSettingArgs struct {
	Kind                 string  `pulumi:"kind"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
	SettingsResourceName *string `pulumi:"settingsResourceName"`
	WorkspaceName        string  `pulumi:"workspaceName"`
}


type SecurityMLAnalyticsSettingArgs struct {
	Kind                 pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	SettingsResourceName pulumi.StringPtrInput
	WorkspaceName        pulumi.StringInput
}

func (SecurityMLAnalyticsSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityMLAnalyticsSettingArgs)(nil)).Elem()
}

type SecurityMLAnalyticsSettingInput interface {
	pulumi.Input

	ToSecurityMLAnalyticsSettingOutput() SecurityMLAnalyticsSettingOutput
	ToSecurityMLAnalyticsSettingOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingOutput
}

func (*SecurityMLAnalyticsSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityMLAnalyticsSetting)(nil)).Elem()
}

func (i *SecurityMLAnalyticsSetting) ToSecurityMLAnalyticsSettingOutput() SecurityMLAnalyticsSettingOutput {
	return i.ToSecurityMLAnalyticsSettingOutputWithContext(context.Background())
}

func (i *SecurityMLAnalyticsSetting) ToSecurityMLAnalyticsSettingOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityMLAnalyticsSettingOutput)
}

type SecurityMLAnalyticsSettingOutput struct{ *pulumi.OutputState }

func (SecurityMLAnalyticsSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityMLAnalyticsSetting)(nil)).Elem()
}

func (o SecurityMLAnalyticsSettingOutput) ToSecurityMLAnalyticsSettingOutput() SecurityMLAnalyticsSettingOutput {
	return o
}

func (o SecurityMLAnalyticsSettingOutput) ToSecurityMLAnalyticsSettingOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingOutput {
	return o
}

func (o SecurityMLAnalyticsSettingOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityMLAnalyticsSetting) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SecurityMLAnalyticsSettingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityMLAnalyticsSetting) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SecurityMLAnalyticsSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityMLAnalyticsSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecurityMLAnalyticsSettingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SecurityMLAnalyticsSetting) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SecurityMLAnalyticsSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityMLAnalyticsSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingOutput{})
}
