


package v20180201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Profile struct {
	pulumi.CustomResourceState

	DnsConfig                   DnsConfigResponsePtrOutput     `pulumi:"dnsConfig"`
	Endpoints                   EndpointResponseArrayOutput    `pulumi:"endpoints"`
	Location                    pulumi.StringPtrOutput         `pulumi:"location"`
	MonitorConfig               MonitorConfigResponsePtrOutput `pulumi:"monitorConfig"`
	Name                        pulumi.StringOutput            `pulumi:"name"`
	ProfileStatus               pulumi.StringPtrOutput         `pulumi:"profileStatus"`
	Tags                        pulumi.StringMapOutput         `pulumi:"tags"`
	TrafficRoutingMethod        pulumi.StringPtrOutput         `pulumi:"trafficRoutingMethod"`
	TrafficViewEnrollmentStatus pulumi.StringPtrOutput         `pulumi:"trafficViewEnrollmentStatus"`
	Type                        pulumi.StringOutput            `pulumi:"type"`
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
			Type: pulumi.String("azure-native:network/v20180301:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:Profile"),
		},
	})
	opts = append(opts, aliases)
	var resource Profile
	err := ctx.RegisterResource("azure-native:network/v20180201:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure-native:network/v20180201:Profile", name, id, state, &resource, opts...)
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
	DnsConfig                   *DnsConfig        `pulumi:"dnsConfig"`
	Endpoints                   []EndpointType    `pulumi:"endpoints"`
	Location                    *string           `pulumi:"location"`
	MonitorConfig               *MonitorConfig    `pulumi:"monitorConfig"`
	ProfileName                 *string           `pulumi:"profileName"`
	ProfileStatus               *string           `pulumi:"profileStatus"`
	ResourceGroupName           string            `pulumi:"resourceGroupName"`
	Tags                        map[string]string `pulumi:"tags"`
	TrafficRoutingMethod        *string           `pulumi:"trafficRoutingMethod"`
	TrafficViewEnrollmentStatus *string           `pulumi:"trafficViewEnrollmentStatus"`
}


type ProfileArgs struct {
	DnsConfig                   DnsConfigPtrInput
	Endpoints                   EndpointTypeArrayInput
	Location                    pulumi.StringPtrInput
	MonitorConfig               MonitorConfigPtrInput
	ProfileName                 pulumi.StringPtrInput
	ProfileStatus               pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
	TrafficRoutingMethod        pulumi.StringPtrInput
	TrafficViewEnrollmentStatus pulumi.StringPtrInput
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

func (o ProfileOutput) DnsConfig() DnsConfigResponsePtrOutput {
	return o.ApplyT(func(v *Profile) DnsConfigResponsePtrOutput { return v.DnsConfig }).(DnsConfigResponsePtrOutput)
}

func (o ProfileOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v *Profile) EndpointResponseArrayOutput { return v.Endpoints }).(EndpointResponseArrayOutput)
}

func (o ProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ProfileOutput) MonitorConfig() MonitorConfigResponsePtrOutput {
	return o.ApplyT(func(v *Profile) MonitorConfigResponsePtrOutput { return v.MonitorConfig }).(MonitorConfigResponsePtrOutput)
}

func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
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

func (o ProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfileOutput{})
}
