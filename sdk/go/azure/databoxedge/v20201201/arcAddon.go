


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArcAddon struct {
	pulumi.CustomResourceState

	HostPlatform      pulumi.StringOutput      `pulumi:"hostPlatform"`
	HostPlatformType  pulumi.StringOutput      `pulumi:"hostPlatformType"`
	Kind              pulumi.StringOutput      `pulumi:"kind"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	ResourceGroupName pulumi.StringOutput      `pulumi:"resourceGroupName"`
	ResourceLocation  pulumi.StringOutput      `pulumi:"resourceLocation"`
	ResourceName      pulumi.StringOutput      `pulumi:"resourceName"`
	SubscriptionId    pulumi.StringOutput      `pulumi:"subscriptionId"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
	Version           pulumi.StringOutput      `pulumi:"version"`
}


func NewArcAddon(ctx *pulumi.Context,
	name string, args *ArcAddonArgs, opts ...pulumi.ResourceOption) (*ArcAddon, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceLocation == nil {
		return nil, errors.New("invalid value for required argument 'ResourceLocation'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("ArcForKubernetes")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:ArcAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:ArcAddon"),
		},
	})
	opts = append(opts, aliases)
	var resource ArcAddon
	err := ctx.RegisterResource("azure-native:databoxedge/v20201201:ArcAddon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetArcAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArcAddonState, opts ...pulumi.ResourceOption) (*ArcAddon, error) {
	var resource ArcAddon
	err := ctx.ReadResource("azure-native:databoxedge/v20201201:ArcAddon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type arcAddonState struct {
}

type ArcAddonState struct {
}

func (ArcAddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*arcAddonState)(nil)).Elem()
}

type arcAddonArgs struct {
	AddonName         *string `pulumi:"addonName"`
	DeviceName        string  `pulumi:"deviceName"`
	Kind              string  `pulumi:"kind"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceLocation  string  `pulumi:"resourceLocation"`
	ResourceName      string  `pulumi:"resourceName"`
	RoleName          string  `pulumi:"roleName"`
	SubscriptionId    string  `pulumi:"subscriptionId"`
}


type ArcAddonArgs struct {
	AddonName         pulumi.StringPtrInput
	DeviceName        pulumi.StringInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ResourceLocation  pulumi.StringInput
	ResourceName      pulumi.StringInput
	RoleName          pulumi.StringInput
	SubscriptionId    pulumi.StringInput
}

func (ArcAddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*arcAddonArgs)(nil)).Elem()
}

type ArcAddonInput interface {
	pulumi.Input

	ToArcAddonOutput() ArcAddonOutput
	ToArcAddonOutputWithContext(ctx context.Context) ArcAddonOutput
}

func (*ArcAddon) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcAddon)(nil)).Elem()
}

func (i *ArcAddon) ToArcAddonOutput() ArcAddonOutput {
	return i.ToArcAddonOutputWithContext(context.Background())
}

func (i *ArcAddon) ToArcAddonOutputWithContext(ctx context.Context) ArcAddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcAddonOutput)
}

type ArcAddonOutput struct{ *pulumi.OutputState }

func (ArcAddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcAddon)(nil)).Elem()
}

func (o ArcAddonOutput) ToArcAddonOutput() ArcAddonOutput {
	return o
}

func (o ArcAddonOutput) ToArcAddonOutputWithContext(ctx context.Context) ArcAddonOutput {
	return o
}

func (o ArcAddonOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.HostPlatform }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.HostPlatformType }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) ResourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.ResourceLocation }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.ResourceName }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ArcAddon) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ArcAddonOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ArcAddonOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcAddon) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ArcAddonOutput{})
}
