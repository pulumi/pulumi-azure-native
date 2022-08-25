


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Site struct {
	pulumi.CustomResourceState

	AvailabilityState         pulumi.StringOutput                        `pulumi:"availabilityState"`
	ClientAffinityEnabled     pulumi.BoolPtrOutput                       `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled         pulumi.BoolPtrOutput                       `pulumi:"clientCertEnabled"`
	CloningInfo               CloningInfoResponsePtrOutput               `pulumi:"cloningInfo"`
	ContainerSize             pulumi.IntPtrOutput                        `pulumi:"containerSize"`
	DefaultHostName           pulumi.StringOutput                        `pulumi:"defaultHostName"`
	Enabled                   pulumi.BoolPtrOutput                       `pulumi:"enabled"`
	EnabledHostNames          pulumi.StringArrayOutput                   `pulumi:"enabledHostNames"`
	GatewaySiteName           pulumi.StringPtrOutput                     `pulumi:"gatewaySiteName"`
	HostNameSslStates         HostNameSslStateResponseArrayOutput        `pulumi:"hostNameSslStates"`
	HostNames                 pulumi.StringArrayOutput                   `pulumi:"hostNames"`
	HostNamesDisabled         pulumi.BoolPtrOutput                       `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrOutput `pulumi:"hostingEnvironmentProfile"`
	IsDefaultContainer        pulumi.BoolOutput                          `pulumi:"isDefaultContainer"`
	Kind                      pulumi.StringPtrOutput                     `pulumi:"kind"`
	LastModifiedTimeUtc       pulumi.StringOutput                        `pulumi:"lastModifiedTimeUtc"`
	Location                  pulumi.StringOutput                        `pulumi:"location"`
	MaxNumberOfWorkers        pulumi.IntPtrOutput                        `pulumi:"maxNumberOfWorkers"`
	MicroService              pulumi.StringPtrOutput                     `pulumi:"microService"`
	Name                      pulumi.StringPtrOutput                     `pulumi:"name"`
	OutboundIpAddresses       pulumi.StringOutput                        `pulumi:"outboundIpAddresses"`
	PremiumAppDeployed        pulumi.BoolOutput                          `pulumi:"premiumAppDeployed"`
	RepositorySiteName        pulumi.StringOutput                        `pulumi:"repositorySiteName"`
	ResourceGroup             pulumi.StringOutput                        `pulumi:"resourceGroup"`
	ScmSiteAlsoStopped        pulumi.BoolPtrOutput                       `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId              pulumi.StringPtrOutput                     `pulumi:"serverFarmId"`
	SiteConfig                SiteConfigResponsePtrOutput                `pulumi:"siteConfig"`
	State                     pulumi.StringOutput                        `pulumi:"state"`
	Tags                      pulumi.StringMapOutput                     `pulumi:"tags"`
	TargetSwapSlot            pulumi.StringOutput                        `pulumi:"targetSwapSlot"`
	TrafficManagerHostNames   pulumi.StringArrayOutput                   `pulumi:"trafficManagerHostNames"`
	Type                      pulumi.StringPtrOutput                     `pulumi:"type"`
	UsageState                pulumi.StringOutput                        `pulumi:"usageState"`
}


func NewSite(ctx *pulumi.Context,
	name string, args *SiteArgs, opts ...pulumi.ResourceOption) (*Site, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:Site"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:Site"),
		},
	})
	opts = append(opts, aliases)
	var resource Site
	err := ctx.RegisterResource("azure-native:web/v20150801:Site", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteState, opts ...pulumi.ResourceOption) (*Site, error) {
	var resource Site
	err := ctx.ReadResource("azure-native:web/v20150801:Site", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteState struct {
}

type SiteState struct {
}

func (SiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteState)(nil)).Elem()
}

type siteArgs struct {
	ClientAffinityEnabled        *bool                      `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled            *bool                      `pulumi:"clientCertEnabled"`
	CloningInfo                  *CloningInfo               `pulumi:"cloningInfo"`
	ContainerSize                *int                       `pulumi:"containerSize"`
	Enabled                      *bool                      `pulumi:"enabled"`
	ForceDnsRegistration         *string                    `pulumi:"forceDnsRegistration"`
	GatewaySiteName              *string                    `pulumi:"gatewaySiteName"`
	HostNameSslStates            []HostNameSslState         `pulumi:"hostNameSslStates"`
	HostNamesDisabled            *bool                      `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile    *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	Id                           *string                    `pulumi:"id"`
	Kind                         *string                    `pulumi:"kind"`
	Location                     *string                    `pulumi:"location"`
	MaxNumberOfWorkers           *int                       `pulumi:"maxNumberOfWorkers"`
	MicroService                 *string                    `pulumi:"microService"`
	Name                         *string                    `pulumi:"name"`
	ResourceGroupName            string                     `pulumi:"resourceGroupName"`
	ScmSiteAlsoStopped           *bool                      `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId                 *string                    `pulumi:"serverFarmId"`
	SiteConfig                   *SiteConfig                `pulumi:"siteConfig"`
	SkipCustomDomainVerification *string                    `pulumi:"skipCustomDomainVerification"`
	SkipDnsRegistration          *string                    `pulumi:"skipDnsRegistration"`
	Tags                         map[string]string          `pulumi:"tags"`
	TtlInSeconds                 *string                    `pulumi:"ttlInSeconds"`
	Type                         *string                    `pulumi:"type"`
}


type SiteArgs struct {
	ClientAffinityEnabled        pulumi.BoolPtrInput
	ClientCertEnabled            pulumi.BoolPtrInput
	CloningInfo                  CloningInfoPtrInput
	ContainerSize                pulumi.IntPtrInput
	Enabled                      pulumi.BoolPtrInput
	ForceDnsRegistration         pulumi.StringPtrInput
	GatewaySiteName              pulumi.StringPtrInput
	HostNameSslStates            HostNameSslStateArrayInput
	HostNamesDisabled            pulumi.BoolPtrInput
	HostingEnvironmentProfile    HostingEnvironmentProfilePtrInput
	Id                           pulumi.StringPtrInput
	Kind                         pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	MaxNumberOfWorkers           pulumi.IntPtrInput
	MicroService                 pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	ScmSiteAlsoStopped           pulumi.BoolPtrInput
	ServerFarmId                 pulumi.StringPtrInput
	SiteConfig                   SiteConfigPtrInput
	SkipCustomDomainVerification pulumi.StringPtrInput
	SkipDnsRegistration          pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
	TtlInSeconds                 pulumi.StringPtrInput
	Type                         pulumi.StringPtrInput
}

func (SiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteArgs)(nil)).Elem()
}

type SiteInput interface {
	pulumi.Input

	ToSiteOutput() SiteOutput
	ToSiteOutputWithContext(ctx context.Context) SiteOutput
}

func (*Site) ElementType() reflect.Type {
	return reflect.TypeOf((**Site)(nil)).Elem()
}

func (i *Site) ToSiteOutput() SiteOutput {
	return i.ToSiteOutputWithContext(context.Background())
}

func (i *Site) ToSiteOutputWithContext(ctx context.Context) SiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteOutput)
}

type SiteOutput struct{ *pulumi.OutputState }

func (SiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Site)(nil)).Elem()
}

func (o SiteOutput) ToSiteOutput() SiteOutput {
	return o
}

func (o SiteOutput) ToSiteOutputWithContext(ctx context.Context) SiteOutput {
	return o
}

func (o SiteOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.AvailabilityState }).(pulumi.StringOutput)
}

func (o SiteOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.BoolPtrOutput { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.BoolPtrOutput { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteOutput) CloningInfo() CloningInfoResponsePtrOutput {
	return o.ApplyT(func(v *Site) CloningInfoResponsePtrOutput { return v.CloningInfo }).(CloningInfoResponsePtrOutput)
}

func (o SiteOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.IntPtrOutput { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

func (o SiteOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.DefaultHostName }).(pulumi.StringOutput)
}

func (o SiteOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Site) pulumi.StringArrayOutput { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

func (o SiteOutput) GatewaySiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.GatewaySiteName }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v *Site) HostNameSslStateResponseArrayOutput { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

func (o SiteOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Site) pulumi.StringArrayOutput { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o SiteOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.BoolPtrOutput { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

func (o SiteOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v *Site) HostingEnvironmentProfileResponsePtrOutput { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o SiteOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v *Site) pulumi.BoolOutput { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

func (o SiteOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o SiteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteOutput) MaxNumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.IntPtrOutput { return v.MaxNumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o SiteOutput) MicroService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.MicroService }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

func (o SiteOutput) PremiumAppDeployed() pulumi.BoolOutput {
	return o.ApplyT(func(v *Site) pulumi.BoolOutput { return v.PremiumAppDeployed }).(pulumi.BoolOutput)
}

func (o SiteOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.RepositorySiteName }).(pulumi.StringOutput)
}

func (o SiteOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o SiteOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.BoolPtrOutput { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

func (o SiteOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v *Site) SiteConfigResponsePtrOutput { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

func (o SiteOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o SiteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Site) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

func (o SiteOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Site) pulumi.StringArrayOutput { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

func (o SiteOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o SiteOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.UsageState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteOutput{})
}
