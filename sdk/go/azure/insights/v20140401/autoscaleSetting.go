


package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type AutoscaleSetting struct {
	pulumi.CustomResourceState

	Enabled                pulumi.BoolPtrOutput                     `pulumi:"enabled"`
	Location               pulumi.StringOutput                      `pulumi:"location"`
	Name                   pulumi.StringOutput                      `pulumi:"name"`
	Notifications          AutoscaleNotificationResponseArrayOutput `pulumi:"notifications"`
	Profiles               AutoscaleProfileResponseArrayOutput      `pulumi:"profiles"`
	Tags                   pulumi.StringMapOutput                   `pulumi:"tags"`
	TargetResourceLocation pulumi.StringPtrOutput                   `pulumi:"targetResourceLocation"`
	TargetResourceUri      pulumi.StringPtrOutput                   `pulumi:"targetResourceUri"`
	Type                   pulumi.StringOutput                      `pulumi:"type"`
}


func NewAutoscaleSetting(ctx *pulumi.Context,
	name string, args *AutoscaleSettingArgs, opts ...pulumi.ResourceOption) (*AutoscaleSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Profiles == nil {
		return nil, errors.New("invalid value for required argument 'Profiles'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.Enabled) {
		args.Enabled = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20150401:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210501preview:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20221001:AutoscaleSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource AutoscaleSetting
	err := ctx.RegisterResource("azure-native:insights/v20140401:AutoscaleSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAutoscaleSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscaleSettingState, opts ...pulumi.ResourceOption) (*AutoscaleSetting, error) {
	var resource AutoscaleSetting
	err := ctx.ReadResource("azure-native:insights/v20140401:AutoscaleSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type autoscaleSettingState struct {
}

type AutoscaleSettingState struct {
}

func (AutoscaleSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscaleSettingState)(nil)).Elem()
}

type autoscaleSettingArgs struct {
	AutoscaleSettingName   *string                 `pulumi:"autoscaleSettingName"`
	Enabled                *bool                   `pulumi:"enabled"`
	Location               *string                 `pulumi:"location"`
	Name                   *string                 `pulumi:"name"`
	Notifications          []AutoscaleNotification `pulumi:"notifications"`
	Profiles               []AutoscaleProfile      `pulumi:"profiles"`
	ResourceGroupName      string                  `pulumi:"resourceGroupName"`
	Tags                   map[string]string       `pulumi:"tags"`
	TargetResourceLocation *string                 `pulumi:"targetResourceLocation"`
	TargetResourceUri      *string                 `pulumi:"targetResourceUri"`
}


type AutoscaleSettingArgs struct {
	AutoscaleSettingName   pulumi.StringPtrInput
	Enabled                pulumi.BoolPtrInput
	Location               pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	Notifications          AutoscaleNotificationArrayInput
	Profiles               AutoscaleProfileArrayInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	TargetResourceLocation pulumi.StringPtrInput
	TargetResourceUri      pulumi.StringPtrInput
}

func (AutoscaleSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscaleSettingArgs)(nil)).Elem()
}

type AutoscaleSettingInput interface {
	pulumi.Input

	ToAutoscaleSettingOutput() AutoscaleSettingOutput
	ToAutoscaleSettingOutputWithContext(ctx context.Context) AutoscaleSettingOutput
}

func (*AutoscaleSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSetting)(nil)).Elem()
}

func (i *AutoscaleSetting) ToAutoscaleSettingOutput() AutoscaleSettingOutput {
	return i.ToAutoscaleSettingOutputWithContext(context.Background())
}

func (i *AutoscaleSetting) ToAutoscaleSettingOutputWithContext(ctx context.Context) AutoscaleSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingOutput)
}

type AutoscaleSettingOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSetting)(nil)).Elem()
}

func (o AutoscaleSettingOutput) ToAutoscaleSettingOutput() AutoscaleSettingOutput {
	return o
}

func (o AutoscaleSettingOutput) ToAutoscaleSettingOutputWithContext(ctx context.Context) AutoscaleSettingOutput {
	return o
}

func (o AutoscaleSettingOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoscaleSetting) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AutoscaleSettingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoscaleSetting) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AutoscaleSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoscaleSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AutoscaleSettingOutput) Notifications() AutoscaleNotificationResponseArrayOutput {
	return o.ApplyT(func(v *AutoscaleSetting) AutoscaleNotificationResponseArrayOutput { return v.Notifications }).(AutoscaleNotificationResponseArrayOutput)
}

func (o AutoscaleSettingOutput) Profiles() AutoscaleProfileResponseArrayOutput {
	return o.ApplyT(func(v *AutoscaleSetting) AutoscaleProfileResponseArrayOutput { return v.Profiles }).(AutoscaleProfileResponseArrayOutput)
}

func (o AutoscaleSettingOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AutoscaleSetting) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AutoscaleSettingOutput) TargetResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscaleSetting) pulumi.StringPtrOutput { return v.TargetResourceLocation }).(pulumi.StringPtrOutput)
}

func (o AutoscaleSettingOutput) TargetResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscaleSetting) pulumi.StringPtrOutput { return v.TargetResourceUri }).(pulumi.StringPtrOutput)
}

func (o AutoscaleSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoscaleSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoscaleSettingOutput{})
}
