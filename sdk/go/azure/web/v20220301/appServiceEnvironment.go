


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServiceEnvironment struct {
	pulumi.CustomResourceState

	ClusterSettings              NameValuePairResponseArrayOutput              `pulumi:"clusterSettings"`
	CustomDnsSuffixConfiguration CustomDnsSuffixConfigurationResponsePtrOutput `pulumi:"customDnsSuffixConfiguration"`
	DedicatedHostCount           pulumi.IntPtrOutput                           `pulumi:"dedicatedHostCount"`
	DnsSuffix                    pulumi.StringPtrOutput                        `pulumi:"dnsSuffix"`
	FrontEndScaleFactor          pulumi.IntPtrOutput                           `pulumi:"frontEndScaleFactor"`
	HasLinuxWorkers              pulumi.BoolOutput                             `pulumi:"hasLinuxWorkers"`
	InternalLoadBalancingMode    pulumi.StringPtrOutput                        `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount            pulumi.IntPtrOutput                           `pulumi:"ipsslAddressCount"`
	Kind                         pulumi.StringPtrOutput                        `pulumi:"kind"`
	Location                     pulumi.StringOutput                           `pulumi:"location"`
	MaximumNumberOfMachines      pulumi.IntOutput                              `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount               pulumi.IntOutput                              `pulumi:"multiRoleCount"`
	MultiSize                    pulumi.StringPtrOutput                        `pulumi:"multiSize"`
	Name                         pulumi.StringOutput                           `pulumi:"name"`
	NetworkingConfiguration      AseV3NetworkingConfigurationResponsePtrOutput `pulumi:"networkingConfiguration"`
	ProvisioningState            pulumi.StringOutput                           `pulumi:"provisioningState"`
	Status                       pulumi.StringOutput                           `pulumi:"status"`
	Suspended                    pulumi.BoolOutput                             `pulumi:"suspended"`
	Tags                         pulumi.StringMapOutput                        `pulumi:"tags"`
	Type                         pulumi.StringOutput                           `pulumi:"type"`
	UpgradeAvailability          pulumi.StringOutput                           `pulumi:"upgradeAvailability"`
	UpgradePreference            pulumi.StringPtrOutput                        `pulumi:"upgradePreference"`
	UserWhitelistedIpRanges      pulumi.StringArrayOutput                      `pulumi:"userWhitelistedIpRanges"`
	VirtualNetwork               VirtualNetworkProfileResponseOutput           `pulumi:"virtualNetwork"`
	ZoneRedundant                pulumi.BoolPtrOutput                          `pulumi:"zoneRedundant"`
}


func NewAppServiceEnvironment(ctx *pulumi.Context,
	name string, args *AppServiceEnvironmentArgs, opts ...pulumi.ResourceOption) (*AppServiceEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetwork == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetwork'")
	}
	if isZero(args.UpgradePreference) {
		args.UpgradePreference = pulumi.StringPtr("None")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:AppServiceEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceEnvironment
	err := ctx.RegisterResource("azure-native:web/v20220301:AppServiceEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServiceEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceEnvironmentState, opts ...pulumi.ResourceOption) (*AppServiceEnvironment, error) {
	var resource AppServiceEnvironment
	err := ctx.ReadResource("azure-native:web/v20220301:AppServiceEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServiceEnvironmentState struct {
}

type AppServiceEnvironmentState struct {
}

func (AppServiceEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentState)(nil)).Elem()
}

type appServiceEnvironmentArgs struct {
	ClusterSettings              []NameValuePair               `pulumi:"clusterSettings"`
	CustomDnsSuffixConfiguration *CustomDnsSuffixConfiguration `pulumi:"customDnsSuffixConfiguration"`
	DedicatedHostCount           *int                          `pulumi:"dedicatedHostCount"`
	DnsSuffix                    *string                       `pulumi:"dnsSuffix"`
	FrontEndScaleFactor          *int                          `pulumi:"frontEndScaleFactor"`
	InternalLoadBalancingMode    *string                       `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount            *int                          `pulumi:"ipsslAddressCount"`
	Kind                         *string                       `pulumi:"kind"`
	Location                     *string                       `pulumi:"location"`
	MultiSize                    *string                       `pulumi:"multiSize"`
	Name                         *string                       `pulumi:"name"`
	NetworkingConfiguration      *AseV3NetworkingConfiguration `pulumi:"networkingConfiguration"`
	ResourceGroupName            string                        `pulumi:"resourceGroupName"`
	Tags                         map[string]string             `pulumi:"tags"`
	UpgradePreference            *string                       `pulumi:"upgradePreference"`
	UserWhitelistedIpRanges      []string                      `pulumi:"userWhitelistedIpRanges"`
	VirtualNetwork               VirtualNetworkProfile         `pulumi:"virtualNetwork"`
	ZoneRedundant                *bool                         `pulumi:"zoneRedundant"`
}


type AppServiceEnvironmentArgs struct {
	ClusterSettings              NameValuePairArrayInput
	CustomDnsSuffixConfiguration CustomDnsSuffixConfigurationPtrInput
	DedicatedHostCount           pulumi.IntPtrInput
	DnsSuffix                    pulumi.StringPtrInput
	FrontEndScaleFactor          pulumi.IntPtrInput
	InternalLoadBalancingMode    pulumi.StringPtrInput
	IpsslAddressCount            pulumi.IntPtrInput
	Kind                         pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	MultiSize                    pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	NetworkingConfiguration      AseV3NetworkingConfigurationPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
	UpgradePreference            pulumi.StringPtrInput
	UserWhitelistedIpRanges      pulumi.StringArrayInput
	VirtualNetwork               VirtualNetworkProfileInput
	ZoneRedundant                pulumi.BoolPtrInput
}

func (AppServiceEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentArgs)(nil)).Elem()
}

type AppServiceEnvironmentInput interface {
	pulumi.Input

	ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput
	ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput
}

func (*AppServiceEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironment)(nil)).Elem()
}

func (i *AppServiceEnvironment) ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput {
	return i.ToAppServiceEnvironmentOutputWithContext(context.Background())
}

func (i *AppServiceEnvironment) ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceEnvironmentOutput)
}

type AppServiceEnvironmentOutput struct{ *pulumi.OutputState }

func (AppServiceEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironment)(nil)).Elem()
}

func (o AppServiceEnvironmentOutput) ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput {
	return o
}

func (o AppServiceEnvironmentOutput) ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput {
	return o
}

func (o AppServiceEnvironmentOutput) ClusterSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) NameValuePairResponseArrayOutput { return v.ClusterSettings }).(NameValuePairResponseArrayOutput)
}

func (o AppServiceEnvironmentOutput) CustomDnsSuffixConfiguration() CustomDnsSuffixConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) CustomDnsSuffixConfigurationResponsePtrOutput {
		return v.CustomDnsSuffixConfiguration
	}).(CustomDnsSuffixConfigurationResponsePtrOutput)
}

func (o AppServiceEnvironmentOutput) DedicatedHostCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntPtrOutput { return v.DedicatedHostCount }).(pulumi.IntPtrOutput)
}

func (o AppServiceEnvironmentOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) FrontEndScaleFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntPtrOutput { return v.FrontEndScaleFactor }).(pulumi.IntPtrOutput)
}

func (o AppServiceEnvironmentOutput) HasLinuxWorkers() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.BoolOutput { return v.HasLinuxWorkers }).(pulumi.BoolOutput)
}

func (o AppServiceEnvironmentOutput) InternalLoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.InternalLoadBalancingMode }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) IpsslAddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntPtrOutput { return v.IpsslAddressCount }).(pulumi.IntPtrOutput)
}

func (o AppServiceEnvironmentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) MaximumNumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntOutput { return v.MaximumNumberOfMachines }).(pulumi.IntOutput)
}

func (o AppServiceEnvironmentOutput) MultiRoleCount() pulumi.IntOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntOutput { return v.MultiRoleCount }).(pulumi.IntOutput)
}

func (o AppServiceEnvironmentOutput) MultiSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.MultiSize }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) NetworkingConfiguration() AseV3NetworkingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) AseV3NetworkingConfigurationResponsePtrOutput {
		return v.NetworkingConfiguration
	}).(AseV3NetworkingConfigurationResponsePtrOutput)
}

func (o AppServiceEnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) Suspended() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.BoolOutput { return v.Suspended }).(pulumi.BoolOutput)
}

func (o AppServiceEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AppServiceEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) UpgradeAvailability() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.UpgradeAvailability }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) UpgradePreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.UpgradePreference }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) UserWhitelistedIpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringArrayOutput { return v.UserWhitelistedIpRanges }).(pulumi.StringArrayOutput)
}

func (o AppServiceEnvironmentOutput) VirtualNetwork() VirtualNetworkProfileResponseOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) VirtualNetworkProfileResponseOutput { return v.VirtualNetwork }).(VirtualNetworkProfileResponseOutput)
}

func (o AppServiceEnvironmentOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServiceEnvironmentOutput{})
}
