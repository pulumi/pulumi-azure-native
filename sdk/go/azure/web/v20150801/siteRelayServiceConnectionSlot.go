


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteRelayServiceConnectionSlot struct {
	pulumi.CustomResourceState

	BiztalkUri               pulumi.StringPtrOutput `pulumi:"biztalkUri"`
	EntityConnectionString   pulumi.StringPtrOutput `pulumi:"entityConnectionString"`
	EntityName               pulumi.StringPtrOutput `pulumi:"entityName"`
	Hostname                 pulumi.StringPtrOutput `pulumi:"hostname"`
	Kind                     pulumi.StringPtrOutput `pulumi:"kind"`
	Location                 pulumi.StringOutput    `pulumi:"location"`
	Name                     pulumi.StringPtrOutput `pulumi:"name"`
	Port                     pulumi.IntPtrOutput    `pulumi:"port"`
	ResourceConnectionString pulumi.StringPtrOutput `pulumi:"resourceConnectionString"`
	ResourceType             pulumi.StringPtrOutput `pulumi:"resourceType"`
	Tags                     pulumi.StringMapOutput `pulumi:"tags"`
	Type                     pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, args *SiteRelayServiceConnectionSlotArgs, opts ...pulumi.ResourceOption) (*SiteRelayServiceConnectionSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteRelayServiceConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteRelayServiceConnectionSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteRelayServiceConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteRelayServiceConnectionSlotState, opts ...pulumi.ResourceOption) (*SiteRelayServiceConnectionSlot, error) {
	var resource SiteRelayServiceConnectionSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteRelayServiceConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteRelayServiceConnectionSlotState struct {
}

type SiteRelayServiceConnectionSlotState struct {
}

func (SiteRelayServiceConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteRelayServiceConnectionSlotState)(nil)).Elem()
}

type siteRelayServiceConnectionSlotArgs struct {
	BiztalkUri               *string           `pulumi:"biztalkUri"`
	EntityConnectionString   *string           `pulumi:"entityConnectionString"`
	EntityName               *string           `pulumi:"entityName"`
	Hostname                 *string           `pulumi:"hostname"`
	Id                       *string           `pulumi:"id"`
	Kind                     *string           `pulumi:"kind"`
	Location                 *string           `pulumi:"location"`
	Name                     string            `pulumi:"name"`
	Port                     *int              `pulumi:"port"`
	ResourceConnectionString *string           `pulumi:"resourceConnectionString"`
	ResourceGroupName        string            `pulumi:"resourceGroupName"`
	ResourceType             *string           `pulumi:"resourceType"`
	Slot                     string            `pulumi:"slot"`
	Tags                     map[string]string `pulumi:"tags"`
	Type                     *string           `pulumi:"type"`
}


type SiteRelayServiceConnectionSlotArgs struct {
	BiztalkUri               pulumi.StringPtrInput
	EntityConnectionString   pulumi.StringPtrInput
	EntityName               pulumi.StringPtrInput
	Hostname                 pulumi.StringPtrInput
	Id                       pulumi.StringPtrInput
	Kind                     pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Name                     pulumi.StringInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ResourceType             pulumi.StringPtrInput
	Slot                     pulumi.StringInput
	Tags                     pulumi.StringMapInput
	Type                     pulumi.StringPtrInput
}

func (SiteRelayServiceConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteRelayServiceConnectionSlotArgs)(nil)).Elem()
}

type SiteRelayServiceConnectionSlotInput interface {
	pulumi.Input

	ToSiteRelayServiceConnectionSlotOutput() SiteRelayServiceConnectionSlotOutput
	ToSiteRelayServiceConnectionSlotOutputWithContext(ctx context.Context) SiteRelayServiceConnectionSlotOutput
}

func (*SiteRelayServiceConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteRelayServiceConnectionSlot)(nil)).Elem()
}

func (i *SiteRelayServiceConnectionSlot) ToSiteRelayServiceConnectionSlotOutput() SiteRelayServiceConnectionSlotOutput {
	return i.ToSiteRelayServiceConnectionSlotOutputWithContext(context.Background())
}

func (i *SiteRelayServiceConnectionSlot) ToSiteRelayServiceConnectionSlotOutputWithContext(ctx context.Context) SiteRelayServiceConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteRelayServiceConnectionSlotOutput)
}

type SiteRelayServiceConnectionSlotOutput struct{ *pulumi.OutputState }

func (SiteRelayServiceConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteRelayServiceConnectionSlot)(nil)).Elem()
}

func (o SiteRelayServiceConnectionSlotOutput) ToSiteRelayServiceConnectionSlotOutput() SiteRelayServiceConnectionSlotOutput {
	return o
}

func (o SiteRelayServiceConnectionSlotOutput) ToSiteRelayServiceConnectionSlotOutputWithContext(ctx context.Context) SiteRelayServiceConnectionSlotOutput {
	return o
}

func (o SiteRelayServiceConnectionSlotOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteRelayServiceConnectionSlotOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteRelayServiceConnectionSlotOutput{})
}
