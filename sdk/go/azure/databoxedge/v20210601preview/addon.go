


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Addon struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewAddon(ctx *pulumi.Context,
	name string, args *AddonArgs, opts ...pulumi.ResourceOption) (*Addon, error) {
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
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:Addon"),
		},
	})
	opts = append(opts, aliases)
	var resource Addon
	err := ctx.RegisterResource("azure-native:databoxedge/v20210601preview:Addon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddonState, opts ...pulumi.ResourceOption) (*Addon, error) {
	var resource Addon
	err := ctx.ReadResource("azure-native:databoxedge/v20210601preview:Addon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type addonState struct {
}

type AddonState struct {
}

func (AddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*addonState)(nil)).Elem()
}

type addonArgs struct {
	AddonName         *string `pulumi:"addonName"`
	DeviceName        string  `pulumi:"deviceName"`
	Kind              string  `pulumi:"kind"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RoleName          string  `pulumi:"roleName"`
}


type AddonArgs struct {
	AddonName         pulumi.StringPtrInput
	DeviceName        pulumi.StringInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RoleName          pulumi.StringInput
}

func (AddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addonArgs)(nil)).Elem()
}

type AddonInput interface {
	pulumi.Input

	ToAddonOutput() AddonOutput
	ToAddonOutputWithContext(ctx context.Context) AddonOutput
}

func (*Addon) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil)).Elem()
}

func (i *Addon) ToAddonOutput() AddonOutput {
	return i.ToAddonOutputWithContext(context.Background())
}

func (i *Addon) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonOutput)
}

type AddonOutput struct{ *pulumi.OutputState }

func (AddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil)).Elem()
}

func (o AddonOutput) ToAddonOutput() AddonOutput {
	return o
}

func (o AddonOutput) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return o
}

func (o AddonOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AddonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AddonOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Addon) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AddonOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AddonOutput{})
}
