


package v20210404preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Spacecraft struct {
	pulumi.CustomResourceState

	AuthorizationStatus         pulumi.StringOutput               `pulumi:"authorizationStatus"`
	AuthorizationStatusExtended pulumi.StringOutput               `pulumi:"authorizationStatusExtended"`
	Etag                        pulumi.StringOutput               `pulumi:"etag"`
	Links                       SpacecraftLinkResponseArrayOutput `pulumi:"links"`
	Location                    pulumi.StringOutput               `pulumi:"location"`
	Name                        pulumi.StringOutput               `pulumi:"name"`
	NoradId                     pulumi.StringOutput               `pulumi:"noradId"`
	SystemData                  SystemDataResponseOutput          `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput            `pulumi:"tags"`
	TitleLine                   pulumi.StringPtrOutput            `pulumi:"titleLine"`
	TleLine1                    pulumi.StringPtrOutput            `pulumi:"tleLine1"`
	TleLine2                    pulumi.StringPtrOutput            `pulumi:"tleLine2"`
	Type                        pulumi.StringOutput               `pulumi:"type"`
}


func NewSpacecraft(ctx *pulumi.Context,
	name string, args *SpacecraftArgs, opts ...pulumi.ResourceOption) (*Spacecraft, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NoradId == nil {
		return nil, errors.New("invalid value for required argument 'NoradId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:orbital:Spacecraft"),
		},
		{
			Type: pulumi.String("azure-native:orbital/v20220301:Spacecraft"),
		},
	})
	opts = append(opts, aliases)
	var resource Spacecraft
	err := ctx.RegisterResource("azure-native:orbital/v20210404preview:Spacecraft", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSpacecraft(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpacecraftState, opts ...pulumi.ResourceOption) (*Spacecraft, error) {
	var resource Spacecraft
	err := ctx.ReadResource("azure-native:orbital/v20210404preview:Spacecraft", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type spacecraftState struct {
}

type SpacecraftState struct {
}

func (SpacecraftState) ElementType() reflect.Type {
	return reflect.TypeOf((*spacecraftState)(nil)).Elem()
}

type spacecraftArgs struct {
	Links             []SpacecraftLink  `pulumi:"links"`
	Location          *string           `pulumi:"location"`
	NoradId           string            `pulumi:"noradId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SpacecraftName    *string           `pulumi:"spacecraftName"`
	Tags              map[string]string `pulumi:"tags"`
	TitleLine         *string           `pulumi:"titleLine"`
	TleLine1          *string           `pulumi:"tleLine1"`
	TleLine2          *string           `pulumi:"tleLine2"`
}


type SpacecraftArgs struct {
	Links             SpacecraftLinkArrayInput
	Location          pulumi.StringPtrInput
	NoradId           pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SpacecraftName    pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	TitleLine         pulumi.StringPtrInput
	TleLine1          pulumi.StringPtrInput
	TleLine2          pulumi.StringPtrInput
}

func (SpacecraftArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spacecraftArgs)(nil)).Elem()
}

type SpacecraftInput interface {
	pulumi.Input

	ToSpacecraftOutput() SpacecraftOutput
	ToSpacecraftOutputWithContext(ctx context.Context) SpacecraftOutput
}

func (*Spacecraft) ElementType() reflect.Type {
	return reflect.TypeOf((**Spacecraft)(nil)).Elem()
}

func (i *Spacecraft) ToSpacecraftOutput() SpacecraftOutput {
	return i.ToSpacecraftOutputWithContext(context.Background())
}

func (i *Spacecraft) ToSpacecraftOutputWithContext(ctx context.Context) SpacecraftOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacecraftOutput)
}

type SpacecraftOutput struct{ *pulumi.OutputState }

func (SpacecraftOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Spacecraft)(nil)).Elem()
}

func (o SpacecraftOutput) ToSpacecraftOutput() SpacecraftOutput {
	return o
}

func (o SpacecraftOutput) ToSpacecraftOutputWithContext(ctx context.Context) SpacecraftOutput {
	return o
}

func (o SpacecraftOutput) AuthorizationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringOutput { return v.AuthorizationStatus }).(pulumi.StringOutput)
}

func (o SpacecraftOutput) AuthorizationStatusExtended() pulumi.StringOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringOutput { return v.AuthorizationStatusExtended }).(pulumi.StringOutput)
}

func (o SpacecraftOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SpacecraftOutput) Links() SpacecraftLinkResponseArrayOutput {
	return o.ApplyT(func(v *Spacecraft) SpacecraftLinkResponseArrayOutput { return v.Links }).(SpacecraftLinkResponseArrayOutput)
}

func (o SpacecraftOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SpacecraftOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SpacecraftOutput) NoradId() pulumi.StringOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringOutput { return v.NoradId }).(pulumi.StringOutput)
}

func (o SpacecraftOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Spacecraft) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SpacecraftOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SpacecraftOutput) TitleLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringPtrOutput { return v.TitleLine }).(pulumi.StringPtrOutput)
}

func (o SpacecraftOutput) TleLine1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringPtrOutput { return v.TleLine1 }).(pulumi.StringPtrOutput)
}

func (o SpacecraftOutput) TleLine2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringPtrOutput { return v.TleLine2 }).(pulumi.StringPtrOutput)
}

func (o SpacecraftOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Spacecraft) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SpacecraftOutput{})
}
