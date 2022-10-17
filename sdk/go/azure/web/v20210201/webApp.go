


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebApp struct {
	pulumi.CustomResourceState

	AvailabilityState           pulumi.StringOutput                        `pulumi:"availabilityState"`
	ClientAffinityEnabled       pulumi.BoolPtrOutput                       `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled           pulumi.BoolPtrOutput                       `pulumi:"clientCertEnabled"`
	ClientCertExclusionPaths    pulumi.StringPtrOutput                     `pulumi:"clientCertExclusionPaths"`
	ClientCertMode              pulumi.StringPtrOutput                     `pulumi:"clientCertMode"`
	ContainerSize               pulumi.IntPtrOutput                        `pulumi:"containerSize"`
	CustomDomainVerificationId  pulumi.StringPtrOutput                     `pulumi:"customDomainVerificationId"`
	DailyMemoryTimeQuota        pulumi.IntPtrOutput                        `pulumi:"dailyMemoryTimeQuota"`
	DefaultHostName             pulumi.StringOutput                        `pulumi:"defaultHostName"`
	Enabled                     pulumi.BoolPtrOutput                       `pulumi:"enabled"`
	EnabledHostNames            pulumi.StringArrayOutput                   `pulumi:"enabledHostNames"`
	ExtendedLocation            ExtendedLocationResponsePtrOutput          `pulumi:"extendedLocation"`
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
	KeyVaultReferenceIdentity   pulumi.StringPtrOutput                     `pulumi:"keyVaultReferenceIdentity"`
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
	StorageAccountRequired      pulumi.BoolPtrOutput                       `pulumi:"storageAccountRequired"`
	SuspendedTill               pulumi.StringOutput                        `pulumi:"suspendedTill"`
	Tags                        pulumi.StringMapOutput                     `pulumi:"tags"`
	TargetSwapSlot              pulumi.StringOutput                        `pulumi:"targetSwapSlot"`
	TrafficManagerHostNames     pulumi.StringArrayOutput                   `pulumi:"trafficManagerHostNames"`
	Type                        pulumi.StringOutput                        `pulumi:"type"`
	UsageState                  pulumi.StringOutput                        `pulumi:"usageState"`
	VirtualNetworkSubnetId      pulumi.StringPtrOutput                     `pulumi:"virtualNetworkSubnetId"`
}


func NewWebApp(ctx *pulumi.Context,
	name string, args *WebAppArgs, opts ...pulumi.ResourceOption) (*WebApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:web:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebApp"),
		},
	})
	opts = append(opts, aliases)
	var resource WebApp
	err := ctx.RegisterResource("azure-native:web/v20210201:WebApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppState, opts ...pulumi.ResourceOption) (*WebApp, error) {
	var resource WebApp
	err := ctx.ReadResource("azure-native:web/v20210201:WebApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppState struct {
}

type WebAppState struct {
}

func (WebAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppState)(nil)).Elem()
}

type webAppArgs struct {
	ClientAffinityEnabled      *bool                      `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled          *bool                      `pulumi:"clientCertEnabled"`
	ClientCertExclusionPaths   *string                    `pulumi:"clientCertExclusionPaths"`
	ClientCertMode             *ClientCertMode            `pulumi:"clientCertMode"`
	CloningInfo                *CloningInfo               `pulumi:"cloningInfo"`
	ContainerSize              *int                       `pulumi:"containerSize"`
	CustomDomainVerificationId *string                    `pulumi:"customDomainVerificationId"`
	DailyMemoryTimeQuota       *int                       `pulumi:"dailyMemoryTimeQuota"`
	Enabled                    *bool                      `pulumi:"enabled"`
	ExtendedLocation           *ExtendedLocation          `pulumi:"extendedLocation"`
	HostNameSslStates          []HostNameSslState         `pulumi:"hostNameSslStates"`
	HostNamesDisabled          *bool                      `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile  *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	HttpsOnly                  *bool                      `pulumi:"httpsOnly"`
	HyperV                     *bool                      `pulumi:"hyperV"`
	Identity                   *ManagedServiceIdentity    `pulumi:"identity"`
	IsXenon                    *bool                      `pulumi:"isXenon"`
	KeyVaultReferenceIdentity  *string                    `pulumi:"keyVaultReferenceIdentity"`
	Kind                       *string                    `pulumi:"kind"`
	Location                   *string                    `pulumi:"location"`
	Name                       *string                    `pulumi:"name"`
	RedundancyMode             *RedundancyMode            `pulumi:"redundancyMode"`
	Reserved                   *bool                      `pulumi:"reserved"`
	ResourceGroupName          string                     `pulumi:"resourceGroupName"`
	ScmSiteAlsoStopped         *bool                      `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId               *string                    `pulumi:"serverFarmId"`
	SiteConfig                 *SiteConfig                `pulumi:"siteConfig"`
	StorageAccountRequired     *bool                      `pulumi:"storageAccountRequired"`
	Tags                       map[string]string          `pulumi:"tags"`
	VirtualNetworkSubnetId     *string                    `pulumi:"virtualNetworkSubnetId"`
}


type WebAppArgs struct {
	ClientAffinityEnabled      pulumi.BoolPtrInput
	ClientCertEnabled          pulumi.BoolPtrInput
	ClientCertExclusionPaths   pulumi.StringPtrInput
	ClientCertMode             ClientCertModePtrInput
	CloningInfo                CloningInfoPtrInput
	ContainerSize              pulumi.IntPtrInput
	CustomDomainVerificationId pulumi.StringPtrInput
	DailyMemoryTimeQuota       pulumi.IntPtrInput
	Enabled                    pulumi.BoolPtrInput
	ExtendedLocation           ExtendedLocationPtrInput
	HostNameSslStates          HostNameSslStateArrayInput
	HostNamesDisabled          pulumi.BoolPtrInput
	HostingEnvironmentProfile  HostingEnvironmentProfilePtrInput
	HttpsOnly                  pulumi.BoolPtrInput
	HyperV                     pulumi.BoolPtrInput
	Identity                   ManagedServiceIdentityPtrInput
	IsXenon                    pulumi.BoolPtrInput
	KeyVaultReferenceIdentity  pulumi.StringPtrInput
	Kind                       pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	RedundancyMode             RedundancyModePtrInput
	Reserved                   pulumi.BoolPtrInput
	ResourceGroupName          pulumi.StringInput
	ScmSiteAlsoStopped         pulumi.BoolPtrInput
	ServerFarmId               pulumi.StringPtrInput
	SiteConfig                 SiteConfigPtrInput
	StorageAccountRequired     pulumi.BoolPtrInput
	Tags                       pulumi.StringMapInput
	VirtualNetworkSubnetId     pulumi.StringPtrInput
}

func (WebAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppArgs)(nil)).Elem()
}

type WebAppInput interface {
	pulumi.Input

	ToWebAppOutput() WebAppOutput
	ToWebAppOutputWithContext(ctx context.Context) WebAppOutput
}

func (*WebApp) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApp)(nil)).Elem()
}

func (i *WebApp) ToWebAppOutput() WebAppOutput {
	return i.ToWebAppOutputWithContext(context.Background())
}

func (i *WebApp) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppOutput)
}

type WebAppOutput struct{ *pulumi.OutputState }

func (WebAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApp)(nil)).Elem()
}

func (o WebAppOutput) ToWebAppOutput() WebAppOutput {
	return o
}

func (o WebAppOutput) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return o
}

func (o WebAppOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.AvailabilityState }).(pulumi.StringOutput)
}

func (o WebAppOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) ClientCertExclusionPaths() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.ClientCertExclusionPaths }).(pulumi.StringPtrOutput)
}

func (o WebAppOutput) ClientCertMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.ClientCertMode }).(pulumi.StringPtrOutput)
}

func (o WebAppOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.IntPtrOutput { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

func (o WebAppOutput) CustomDomainVerificationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.CustomDomainVerificationId }).(pulumi.StringPtrOutput)
}

func (o WebAppOutput) DailyMemoryTimeQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.IntPtrOutput { return v.DailyMemoryTimeQuota }).(pulumi.IntPtrOutput)
}

func (o WebAppOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.DefaultHostName }).(pulumi.StringOutput)
}

func (o WebAppOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringArrayOutput { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

func (o WebAppOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *WebApp) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o WebAppOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v *WebApp) HostNameSslStateResponseArrayOutput { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

func (o WebAppOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringArrayOutput { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o WebAppOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v *WebApp) HostingEnvironmentProfileResponsePtrOutput { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o WebAppOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) HyperV() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.HyperV }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *WebApp) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o WebAppOutput) InProgressOperationId() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.InProgressOperationId }).(pulumi.StringOutput)
}

func (o WebAppOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolOutput { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

func (o WebAppOutput) IsXenon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.IsXenon }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

func (o WebAppOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o WebAppOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WebAppOutput) MaxNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v *WebApp) pulumi.IntOutput { return v.MaxNumberOfWorkers }).(pulumi.IntOutput)
}

func (o WebAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

func (o WebAppOutput) PossibleOutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.PossibleOutboundIpAddresses }).(pulumi.StringOutput)
}

func (o WebAppOutput) RedundancyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.RedundancyMode }).(pulumi.StringPtrOutput)
}

func (o WebAppOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.RepositorySiteName }).(pulumi.StringOutput)
}

func (o WebAppOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.Reserved }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o WebAppOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

func (o WebAppOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v *WebApp) SiteConfigResponsePtrOutput { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

func (o WebAppOutput) SlotSwapStatus() SlotSwapStatusResponseOutput {
	return o.ApplyT(func(v *WebApp) SlotSwapStatusResponseOutput { return v.SlotSwapStatus }).(SlotSwapStatusResponseOutput)
}

func (o WebAppOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o WebAppOutput) StorageAccountRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.StorageAccountRequired }).(pulumi.BoolPtrOutput)
}

func (o WebAppOutput) SuspendedTill() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.SuspendedTill }).(pulumi.StringOutput)
}

func (o WebAppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WebAppOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

func (o WebAppOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringArrayOutput { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

func (o WebAppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WebAppOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.UsageState }).(pulumi.StringOutput)
}

func (o WebAppOutput) VirtualNetworkSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.VirtualNetworkSubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppOutput{})
}
