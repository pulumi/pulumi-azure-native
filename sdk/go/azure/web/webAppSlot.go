


package web

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
	if args.HyperV == nil {
		args.HyperV = pulumi.BoolPtr(false)
	}
	if args.IsXenon == nil {
		args.IsXenon = pulumi.BoolPtr(false)
	}
	if args.Reserved == nil {
		args.Reserved = pulumi.BoolPtr(false)
	}
	if args.ScmSiteAlsoStopped == nil {
		args.ScmSiteAlsoStopped = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
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
			Type: pulumi.String("azure-native:web/v20190801:WebAppSlot"),
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
	})
	opts = append(opts, aliases)
	var resource WebAppSlot
	err := ctx.RegisterResource("azure-native:web:WebAppSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSlotState, opts ...pulumi.ResourceOption) (*WebAppSlot, error) {
	var resource WebAppSlot
	err := ctx.ReadResource("azure-native:web:WebAppSlot", name, id, state, &resource, opts...)
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
	Name                       string                     `pulumi:"name"`
	RedundancyMode             *RedundancyMode            `pulumi:"redundancyMode"`
	Reserved                   *bool                      `pulumi:"reserved"`
	ResourceGroupName          string                     `pulumi:"resourceGroupName"`
	ScmSiteAlsoStopped         *bool                      `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId               *string                    `pulumi:"serverFarmId"`
	SiteConfig                 *SiteConfig                `pulumi:"siteConfig"`
	Slot                       *string                    `pulumi:"slot"`
	StorageAccountRequired     *bool                      `pulumi:"storageAccountRequired"`
	Tags                       map[string]string          `pulumi:"tags"`
	VirtualNetworkSubnetId     *string                    `pulumi:"virtualNetworkSubnetId"`
}


type WebAppSlotArgs struct {
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
	Name                       pulumi.StringInput
	RedundancyMode             RedundancyModePtrInput
	Reserved                   pulumi.BoolPtrInput
	ResourceGroupName          pulumi.StringInput
	ScmSiteAlsoStopped         pulumi.BoolPtrInput
	ServerFarmId               pulumi.StringPtrInput
	SiteConfig                 SiteConfigPtrInput
	Slot                       pulumi.StringPtrInput
	StorageAccountRequired     pulumi.BoolPtrInput
	Tags                       pulumi.StringMapInput
	VirtualNetworkSubnetId     pulumi.StringPtrInput
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
	return reflect.TypeOf((*WebAppSlot)(nil))
}

func (i *WebAppSlot) ToWebAppSlotOutput() WebAppSlotOutput {
	return i.ToWebAppSlotOutputWithContext(context.Background())
}

func (i *WebAppSlot) ToWebAppSlotOutputWithContext(ctx context.Context) WebAppSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSlotOutput)
}

type WebAppSlotOutput struct{ *pulumi.OutputState }

func (WebAppSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSlot)(nil))
}

func (o WebAppSlotOutput) ToWebAppSlotOutput() WebAppSlotOutput {
	return o
}

func (o WebAppSlotOutput) ToWebAppSlotOutputWithContext(ctx context.Context) WebAppSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppSlotOutput{})
}
