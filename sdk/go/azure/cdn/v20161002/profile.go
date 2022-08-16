


package v20161002

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Profile struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput    `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	ResourceState     pulumi.StringOutput    `pulumi:"resourceState"`
	Sku               SkuResponseOutput      `pulumi:"sku"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20170402:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20171012:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190415:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Profile"),
		},
	})
	opts = append(opts, aliases)
	var resource Profile
	err := ctx.RegisterResource("azure-native:cdn/v20161002:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure-native:cdn/v20161002:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type profileState struct {
}

type ProfileState struct {
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	Location          *string           `pulumi:"location"`
	ProfileName       *string           `pulumi:"profileName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               Sku               `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type ProfileArgs struct {
	Location          pulumi.StringPtrInput
	ProfileName       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
	Tags              pulumi.StringMapInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

func (o ProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProfileOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ProfileOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o ProfileOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Profile) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o ProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfileOutput{})
}
