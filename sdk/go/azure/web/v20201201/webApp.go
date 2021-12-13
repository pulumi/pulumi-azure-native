


package v20201201

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
			Type: pulumi.String("azure-native:web/v20210101:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebApp"),
		},
	})
	opts = append(opts, aliases)
	var resource WebApp
	err := ctx.RegisterResource("azure-native:web/v20201201:WebApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppState, opts ...pulumi.ResourceOption) (*WebApp, error) {
	var resource WebApp
	err := ctx.ReadResource("azure-native:web/v20201201:WebApp", name, id, state, &resource, opts...)
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

func init() {
	pulumi.RegisterOutputType(WebAppOutput{})
}
