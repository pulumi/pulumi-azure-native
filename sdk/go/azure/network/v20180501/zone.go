


package v20180501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Zone struct {
	pulumi.CustomResourceState

	Etag                           pulumi.StringPtrOutput         `pulumi:"etag"`
	Location                       pulumi.StringOutput            `pulumi:"location"`
	MaxNumberOfRecordSets          pulumi.Float64Output           `pulumi:"maxNumberOfRecordSets"`
	MaxNumberOfRecordsPerRecordSet pulumi.Float64Output           `pulumi:"maxNumberOfRecordsPerRecordSet"`
	Name                           pulumi.StringOutput            `pulumi:"name"`
	NameServers                    pulumi.StringArrayOutput       `pulumi:"nameServers"`
	NumberOfRecordSets             pulumi.Float64Output           `pulumi:"numberOfRecordSets"`
	RegistrationVirtualNetworks    SubResourceResponseArrayOutput `pulumi:"registrationVirtualNetworks"`
	ResolutionVirtualNetworks      SubResourceResponseArrayOutput `pulumi:"resolutionVirtualNetworks"`
	Tags                           pulumi.StringMapOutput         `pulumi:"tags"`
	Type                           pulumi.StringOutput            `pulumi:"type"`
	ZoneType                       pulumi.StringPtrOutput         `pulumi:"zoneType"`
}


func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.ZoneType) {
		args.ZoneType = ZoneType("Public")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150504preview:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160401:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301preview:Zone"),
		},
	})
	opts = append(opts, aliases)
	var resource Zone
	err := ctx.RegisterResource("azure-native:network/v20180501:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("azure-native:network/v20180501:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type zoneState struct {
}

type ZoneState struct {
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	Location                    *string           `pulumi:"location"`
	RegistrationVirtualNetworks []SubResource     `pulumi:"registrationVirtualNetworks"`
	ResolutionVirtualNetworks   []SubResource     `pulumi:"resolutionVirtualNetworks"`
	ResourceGroupName           string            `pulumi:"resourceGroupName"`
	Tags                        map[string]string `pulumi:"tags"`
	ZoneName                    *string           `pulumi:"zoneName"`
	ZoneType                    *ZoneType         `pulumi:"zoneType"`
}


type ZoneArgs struct {
	Location                    pulumi.StringPtrInput
	RegistrationVirtualNetworks SubResourceArrayInput
	ResolutionVirtualNetworks   SubResourceArrayInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
	ZoneName                    pulumi.StringPtrInput
	ZoneType                    ZoneTypePtrInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

func (o ZoneOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ZoneOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ZoneOutput) MaxNumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v *Zone) pulumi.Float64Output { return v.MaxNumberOfRecordSets }).(pulumi.Float64Output)
}

func (o ZoneOutput) MaxNumberOfRecordsPerRecordSet() pulumi.Float64Output {
	return o.ApplyT(func(v *Zone) pulumi.Float64Output { return v.MaxNumberOfRecordsPerRecordSet }).(pulumi.Float64Output)
}

func (o ZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ZoneOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringArrayOutput { return v.NameServers }).(pulumi.StringArrayOutput)
}

func (o ZoneOutput) NumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v *Zone) pulumi.Float64Output { return v.NumberOfRecordSets }).(pulumi.Float64Output)
}

func (o ZoneOutput) RegistrationVirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *Zone) SubResourceResponseArrayOutput { return v.RegistrationVirtualNetworks }).(SubResourceResponseArrayOutput)
}

func (o ZoneOutput) ResolutionVirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *Zone) SubResourceResponseArrayOutput { return v.ResolutionVirtualNetworks }).(SubResourceResponseArrayOutput)
}

func (o ZoneOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ZoneOutput) ZoneType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.ZoneType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ZoneOutput{})
}
