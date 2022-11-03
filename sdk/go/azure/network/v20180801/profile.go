


package v20180801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Profile struct {
	pulumi.CustomResourceState

	AllowedEndpointRecordTypes  pulumi.StringArrayOutput       `pulumi:"allowedEndpointRecordTypes"`
	DnsConfig                   DnsConfigResponsePtrOutput     `pulumi:"dnsConfig"`
	Endpoints                   EndpointResponseArrayOutput    `pulumi:"endpoints"`
	Location                    pulumi.StringPtrOutput         `pulumi:"location"`
	MaxReturn                   pulumi.Float64PtrOutput        `pulumi:"maxReturn"`
	MonitorConfig               MonitorConfigResponsePtrOutput `pulumi:"monitorConfig"`
	Name                        pulumi.StringPtrOutput         `pulumi:"name"`
	ProfileStatus               pulumi.StringPtrOutput         `pulumi:"profileStatus"`
	Tags                        pulumi.StringMapOutput         `pulumi:"tags"`
	TrafficRoutingMethod        pulumi.StringPtrOutput         `pulumi:"trafficRoutingMethod"`
	TrafficViewEnrollmentStatus pulumi.StringPtrOutput         `pulumi:"trafficViewEnrollmentStatus"`
	Type                        pulumi.StringPtrOutput         `pulumi:"type"`
}


func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20151101:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170501:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:Profile"),
		},
	})
	opts = append(opts, aliases)
	var resource Profile
	err := ctx.RegisterResource("azure-native:network/v20180801:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure-native:network/v20180801:Profile", name, id, state, &resource, opts...)
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
	AllowedEndpointRecordTypes  []string          `pulumi:"allowedEndpointRecordTypes"`
	DnsConfig                   *DnsConfig        `pulumi:"dnsConfig"`
	Endpoints                   []EndpointType    `pulumi:"endpoints"`
	Id                          *string           `pulumi:"id"`
	Location                    *string           `pulumi:"location"`
	MaxReturn                   *float64          `pulumi:"maxReturn"`
	MonitorConfig               *MonitorConfig    `pulumi:"monitorConfig"`
	Name                        *string           `pulumi:"name"`
	ProfileName                 *string           `pulumi:"profileName"`
	ProfileStatus               *string           `pulumi:"profileStatus"`
	ResourceGroupName           string            `pulumi:"resourceGroupName"`
	Tags                        map[string]string `pulumi:"tags"`
	TrafficRoutingMethod        *string           `pulumi:"trafficRoutingMethod"`
	TrafficViewEnrollmentStatus *string           `pulumi:"trafficViewEnrollmentStatus"`
	Type                        *string           `pulumi:"type"`
}


type ProfileArgs struct {
	AllowedEndpointRecordTypes  pulumi.StringArrayInput
	DnsConfig                   DnsConfigPtrInput
	Endpoints                   EndpointTypeArrayInput
	Id                          pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	MaxReturn                   pulumi.Float64PtrInput
	MonitorConfig               MonitorConfigPtrInput
	Name                        pulumi.StringPtrInput
	ProfileName                 pulumi.StringPtrInput
	ProfileStatus               pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
	TrafficRoutingMethod        pulumi.StringPtrInput
	TrafficViewEnrollmentStatus pulumi.StringPtrInput
	Type                        pulumi.StringPtrInput
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

func (o ProfileOutput) AllowedEndpointRecordTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringArrayOutput { return v.AllowedEndpointRecordTypes }).(pulumi.StringArrayOutput)
}

func (o ProfileOutput) DnsConfig() DnsConfigResponsePtrOutput {
	return o.ApplyT(func(v *Profile) DnsConfigResponsePtrOutput { return v.DnsConfig }).(DnsConfigResponsePtrOutput)
}

func (o ProfileOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v *Profile) EndpointResponseArrayOutput { return v.Endpoints }).(EndpointResponseArrayOutput)
}

func (o ProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) MaxReturn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.Float64PtrOutput { return v.MaxReturn }).(pulumi.Float64PtrOutput)
}

func (o ProfileOutput) MonitorConfig() MonitorConfigResponsePtrOutput {
	return o.ApplyT(func(v *Profile) MonitorConfigResponsePtrOutput { return v.MonitorConfig }).(MonitorConfigResponsePtrOutput)
}

func (o ProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) ProfileStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.ProfileStatus }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProfileOutput) TrafficRoutingMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.TrafficRoutingMethod }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) TrafficViewEnrollmentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.TrafficViewEnrollmentStatus }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfileOutput{})
}
