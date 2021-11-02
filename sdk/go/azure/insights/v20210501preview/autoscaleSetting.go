


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoscaleSetting struct {
	pulumi.CustomResourceState

	Enabled                   pulumi.BoolPtrOutput                       `pulumi:"enabled"`
	Location                  pulumi.StringOutput                        `pulumi:"location"`
	Name                      pulumi.StringOutput                        `pulumi:"name"`
	Notifications             AutoscaleNotificationResponseArrayOutput   `pulumi:"notifications"`
	PredictiveAutoscalePolicy PredictiveAutoscalePolicyResponsePtrOutput `pulumi:"predictiveAutoscalePolicy"`
	Profiles                  AutoscaleProfileResponseArrayOutput        `pulumi:"profiles"`
	SystemData                SystemDataResponseOutput                   `pulumi:"systemData"`
	Tags                      pulumi.StringMapOutput                     `pulumi:"tags"`
	TargetResourceLocation    pulumi.StringPtrOutput                     `pulumi:"targetResourceLocation"`
	TargetResourceUri         pulumi.StringPtrOutput                     `pulumi:"targetResourceUri"`
	Type                      pulumi.StringOutput                        `pulumi:"type"`
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
	if args.Enabled == nil {
		args.Enabled = pulumi.BoolPtr(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20210501preview:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20140401:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20140401:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20150401:AutoscaleSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20150401:AutoscaleSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource AutoscaleSetting
	err := ctx.RegisterResource("azure-native:insights/v20210501preview:AutoscaleSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAutoscaleSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscaleSettingState, opts ...pulumi.ResourceOption) (*AutoscaleSetting, error) {
	var resource AutoscaleSetting
	err := ctx.ReadResource("azure-native:insights/v20210501preview:AutoscaleSetting", name, id, state, &resource, opts...)
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
	AutoscaleSettingName      *string                    `pulumi:"autoscaleSettingName"`
	Enabled                   *bool                      `pulumi:"enabled"`
	Location                  *string                    `pulumi:"location"`
	Name                      *string                    `pulumi:"name"`
	Notifications             []AutoscaleNotification    `pulumi:"notifications"`
	PredictiveAutoscalePolicy *PredictiveAutoscalePolicy `pulumi:"predictiveAutoscalePolicy"`
	Profiles                  []AutoscaleProfile         `pulumi:"profiles"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
	Tags                      map[string]string          `pulumi:"tags"`
	TargetResourceLocation    *string                    `pulumi:"targetResourceLocation"`
	TargetResourceUri         *string                    `pulumi:"targetResourceUri"`
}


type AutoscaleSettingArgs struct {
	AutoscaleSettingName      pulumi.StringPtrInput
	Enabled                   pulumi.BoolPtrInput
	Location                  pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	Notifications             AutoscaleNotificationArrayInput
	PredictiveAutoscalePolicy PredictiveAutoscalePolicyPtrInput
	Profiles                  AutoscaleProfileArrayInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
	TargetResourceLocation    pulumi.StringPtrInput
	TargetResourceUri         pulumi.StringPtrInput
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
	return reflect.TypeOf((*AutoscaleSetting)(nil))
}

func (i *AutoscaleSetting) ToAutoscaleSettingOutput() AutoscaleSettingOutput {
	return i.ToAutoscaleSettingOutputWithContext(context.Background())
}

func (i *AutoscaleSetting) ToAutoscaleSettingOutputWithContext(ctx context.Context) AutoscaleSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingOutput)
}

type AutoscaleSettingOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSetting)(nil))
}

func (o AutoscaleSettingOutput) ToAutoscaleSettingOutput() AutoscaleSettingOutput {
	return o
}

func (o AutoscaleSettingOutput) ToAutoscaleSettingOutputWithContext(ctx context.Context) AutoscaleSettingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AutoscaleSettingOutput{})
}
