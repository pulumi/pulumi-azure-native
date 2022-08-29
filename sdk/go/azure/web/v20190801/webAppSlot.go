


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSlot struct {
	pulumi.CustomResourceState

	AvailabilityState           pulumi.StringOutput                        `pulumi:"availabilityState"`
	ClientAffinityEnabled       pulumi.BoolPtrOutput                       `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled           pulumi.BoolPtrOutput                       `pulumi:"clientCertEnabled"`
	ClientCertExclusionPaths    pulumi.StringPtrOutput                     `pulumi:"clientCertExclusionPaths"`
	ContainerSize               pulumi.IntPtrOutput                        `pulumi:"containerSize"`
	DailyMemoryTimeQuota        pulumi.IntPtrOutput                        `pulumi:"dailyMemoryTimeQuota"`
	DefaultHostName             pulumi.StringOutput                        `pulumi:"defaultHostName"`
	Enabled                     pulumi.BoolPtrOutput                       `pulumi:"enabled"`
	EnabledHostNames            pulumi.StringArrayOutput                   `pulumi:"enabledHostNames"`
	HostNameSslStates           HostNameSslStateResponseArrayOutput        `pulumi:"hostNameSslStates"`
	HostNames                   pulumi.StringArrayOutput                   `pulumi:"hostNames"`
	HostNamesDisabled           pulumi.BoolPtrOutput                       `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile   HostingEnvironmentProfileResponsePtrOutput `pulumi:"hostingEnvironmentProfile"`
	HttpsOnly                   pulumi.BoolPtrOutput                       `pulumi:"httpsOnly"`
	HyperV                      pulumi.BoolPtrOutput                       `pulumi:"hyperV"`
	Identity                    ManagedServiceIdentityResponsePtrOutput    `pulumi:"identity"`
	InProgressOperationId       pulumi.StringOutput                        `pulumi:"inProgressOperationId"`
	IsDefaultContainer          pulumi.BoolOutput                          `pulumi:"isDefaultContainer"`
	IsXenon                     pulumi.BoolPtrOutput                       `pulumi:"isXenon"`
	Kind                        pulumi.StringPtrOutput                     `pulumi:"kind"`
	LastModifiedTimeUtc         pulumi.StringOutput                        `pulumi:"lastModifiedTimeUtc"`
	Location                    pulumi.StringOutput                        `pulumi:"location"`
	MaxNumberOfWorkers          pulumi.IntOutput                           `pulumi:"maxNumberOfWorkers"`
	Name                        pulumi.StringOutput                        `pulumi:"name"`
	OutboundIpAddresses         pulumi.StringOutput                        `pulumi:"outboundIpAddresses"`
	PossibleOutboundIpAddresses pulumi.StringOutput                        `pulumi:"possibleOutboundIpAddresses"`
	RedundancyMode              pulumi.StringPtrOutput                     `pulumi:"redundancyMode"`
	RepositorySiteName          pulumi.StringOutput                        `pulumi:"repositorySiteName"`
	Reserved                    pulumi.BoolPtrOutput                       `pulumi:"reserved"`
	ResourceGroup               pulumi.StringOutput                        `pulumi:"resourceGroup"`
	ScmSiteAlsoStopped          pulumi.BoolPtrOutput                       `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId                pulumi.StringPtrOutput                     `pulumi:"serverFarmId"`
	SiteConfig                  SiteConfigResponsePtrOutput                `pulumi:"siteConfig"`
	SlotSwapStatus              SlotSwapStatusResponseOutput               `pulumi:"slotSwapStatus"`
	State                       pulumi.StringOutput                        `pulumi:"state"`
	SuspendedTill               pulumi.StringOutput                        `pulumi:"suspendedTill"`
	Tags                        pulumi.StringMapOutput                     `pulumi:"tags"`
	TargetSwapSlot              pulumi.StringOutput                        `pulumi:"targetSwapSlot"`
	TrafficManagerHostNames     pulumi.StringArrayOutput                   `pulumi:"trafficManagerHostNames"`
	Type                        pulumi.StringOutput                        `pulumi:"type"`
	UsageState                  pulumi.StringOutput                        `pulumi:"usageState"`
}


func NewWebAppSlot(ctx *pulumi.Context,
	name string, args *WebAppSlotArgs, opts ...pulumi.ResourceOption) (*WebAppSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.HyperV) {
		args.HyperV = pulumi.BoolPtr(false)
	}
	if isZero(args.IsXenon) {
		args.IsXenon = pulumi.BoolPtr(false)
	}
	if isZero(args.Reserved) {
		args.Reserved = pulumi.BoolPtr(false)
	}
	if isZero(args.ScmSiteAlsoStopped) {
		args.ScmSiteAlsoStopped = pulumi.BoolPtr(false)
	}
	if args.SiteConfig != nil {
		args.SiteConfig = args.SiteConfig.ToSiteConfigPtrOutput().ApplyT(func(v *SiteConfig) *SiteConfig { return v.Defaults() }).(SiteConfigPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSlot
	err := ctx.RegisterResource("azure-native:web/v20190801:WebAppSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSlotState, opts ...pulumi.ResourceOption) (*WebAppSlot, error) {
	var resource WebAppSlot
	err := ctx.ReadResource("azure-native:web/v20190801:WebAppSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSlotState struct {
}

type WebAppSlotState struct {
}

func (WebAppSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotState)(nil)).Elem()
}

type webAppSlotArgs struct {
	ClientAffinityEnabled     *bool                      `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled         *bool                      `pulumi:"clientCertEnabled"`
	ClientCertExclusionPaths  *string                    `pulumi:"clientCertExclusionPaths"`
	CloningInfo               *CloningInfo               `pulumi:"cloningInfo"`
	ContainerSize             *int                       `pulumi:"containerSize"`
	DailyMemoryTimeQuota      *int                       `pulumi:"dailyMemoryTimeQuota"`
	Enabled                   *bool                      `pulumi:"enabled"`
	HostNameSslStates         []HostNameSslState         `pulumi:"hostNameSslStates"`
	HostNamesDisabled         *bool                      `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	HttpsOnly                 *bool                      `pulumi:"httpsOnly"`
	HyperV                    *bool                      `pulumi:"hyperV"`
	Identity                  *ManagedServiceIdentity    `pulumi:"identity"`
	IsXenon                   *bool                      `pulumi:"isXenon"`
	Kind                      *string                    `pulumi:"kind"`
	Location                  *string                    `pulumi:"location"`
	Name                      string                     `pulumi:"name"`
	RedundancyMode            *RedundancyMode            `pulumi:"redundancyMode"`
	Reserved                  *bool                      `pulumi:"reserved"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
	ScmSiteAlsoStopped        *bool                      `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId              *string                    `pulumi:"serverFarmId"`
	SiteConfig                *SiteConfig                `pulumi:"siteConfig"`
	Slot                      *string                    `pulumi:"slot"`
	Tags                      map[string]string          `pulumi:"tags"`
}


type WebAppSlotArgs struct {
	ClientAffinityEnabled     pulumi.BoolPtrInput
	ClientCertEnabled         pulumi.BoolPtrInput
	ClientCertExclusionPaths  pulumi.StringPtrInput
	CloningInfo               CloningInfoPtrInput
	ContainerSize             pulumi.IntPtrInput
	DailyMemoryTimeQuota      pulumi.IntPtrInput
	Enabled                   pulumi.BoolPtrInput
	HostNameSslStates         HostNameSslStateArrayInput
	HostNamesDisabled         pulumi.BoolPtrInput
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	HttpsOnly                 pulumi.BoolPtrInput
	HyperV                    pulumi.BoolPtrInput
	Identity                  ManagedServiceIdentityPtrInput
	IsXenon                   pulumi.BoolPtrInput
	Kind                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	Name                      pulumi.StringInput
	RedundancyMode            RedundancyModePtrInput
	Reserved                  pulumi.BoolPtrInput
	ResourceGroupName         pulumi.StringInput
	ScmSiteAlsoStopped        pulumi.BoolPtrInput
	ServerFarmId              pulumi.StringPtrInput
	SiteConfig                SiteConfigPtrInput
	Slot                      pulumi.StringPtrInput
	Tags                      pulumi.StringMapInput
}

func (WebAppSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotArgs)(nil)).Elem()
}

type WebAppSlotInput interface {
	pulumi.Input

	ToWebAppSlotOutput() WebAppSlotOutput
	ToWebAppSlotOutputWithContext(ctx context.Context) WebAppSlotOutput
}

func (*WebAppSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSlot)(nil)).Elem()
}

func (i *WebAppSlot) ToWebAppSlotOutput() WebAppSlotOutput {
	return i.ToWebAppSlotOutputWithContext(context.Background())
}

func (i *WebAppSlot) ToWebAppSlotOutputWithContext(ctx context.Context) WebAppSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSlotOutput)
}

type WebAppSlotOutput struct{ *pulumi.OutputState }

func (WebAppSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSlot)(nil)).Elem()
}

func (o WebAppSlotOutput) ToWebAppSlotOutput() WebAppSlotOutput {
	return o
}

func (o WebAppSlotOutput) ToWebAppSlotOutputWithContext(ctx context.Context) WebAppSlotOutput {
	return o
}

func (o WebAppSlotOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.AvailabilityState }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) ClientCertExclusionPaths() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringPtrOutput { return v.ClientCertExclusionPaths }).(pulumi.StringPtrOutput)
}

func (o WebAppSlotOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.IntPtrOutput { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

func (o WebAppSlotOutput) DailyMemoryTimeQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.IntPtrOutput { return v.DailyMemoryTimeQuota }).(pulumi.IntPtrOutput)
}

func (o WebAppSlotOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.DefaultHostName }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringArrayOutput { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

func (o WebAppSlotOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v *WebAppSlot) HostNameSslStateResponseArrayOutput { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

func (o WebAppSlotOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringArrayOutput { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o WebAppSlotOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v *WebAppSlot) HostingEnvironmentProfileResponsePtrOutput { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o WebAppSlotOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) HyperV() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.HyperV }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *WebAppSlot) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o WebAppSlotOutput) InProgressOperationId() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.InProgressOperationId }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolOutput { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

func (o WebAppSlotOutput) IsXenon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.IsXenon }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppSlotOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) MaxNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.IntOutput { return v.MaxNumberOfWorkers }).(pulumi.IntOutput)
}

func (o WebAppSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) PossibleOutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.PossibleOutboundIpAddresses }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) RedundancyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringPtrOutput { return v.RedundancyMode }).(pulumi.StringPtrOutput)
}

func (o WebAppSlotOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.RepositorySiteName }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.Reserved }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.BoolPtrOutput { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

func (o WebAppSlotOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringPtrOutput { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

func (o WebAppSlotOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v *WebAppSlot) SiteConfigResponsePtrOutput { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

func (o WebAppSlotOutput) SlotSwapStatus() SlotSwapStatusResponseOutput {
	return o.ApplyT(func(v *WebAppSlot) SlotSwapStatusResponseOutput { return v.SlotSwapStatus }).(SlotSwapStatusResponseOutput)
}

func (o WebAppSlotOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) SuspendedTill() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.SuspendedTill }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WebAppSlotOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringArrayOutput { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

func (o WebAppSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WebAppSlotOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlot) pulumi.StringOutput { return v.UsageState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSlotOutput{})
}
