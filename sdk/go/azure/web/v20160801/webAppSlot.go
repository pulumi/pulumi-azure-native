


package v20160801

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
	Identity                    ManagedServiceIdentityResponsePtrOutput    `pulumi:"identity"`
	IsDefaultContainer          pulumi.BoolOutput                          `pulumi:"isDefaultContainer"`
	Kind                        pulumi.StringPtrOutput                     `pulumi:"kind"`
	LastModifiedTimeUtc         pulumi.StringOutput                        `pulumi:"lastModifiedTimeUtc"`
	Location                    pulumi.StringOutput                        `pulumi:"location"`
	MaxNumberOfWorkers          pulumi.IntOutput                           `pulumi:"maxNumberOfWorkers"`
	Name                        pulumi.StringOutput                        `pulumi:"name"`
	OutboundIpAddresses         pulumi.StringOutput                        `pulumi:"outboundIpAddresses"`
	PossibleOutboundIpAddresses pulumi.StringOutput                        `pulumi:"possibleOutboundIpAddresses"`
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
	if args.Reserved == nil {
		args.Reserved = pulumi.BoolPtr(false)
	}
	if args.ScmSiteAlsoStopped == nil {
		args.ScmSiteAlsoStopped = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSlot
	err := ctx.RegisterResource("azure-native:web/v20160801:WebAppSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSlotState, opts ...pulumi.ResourceOption) (*WebAppSlot, error) {
	var resource WebAppSlot
	err := ctx.ReadResource("azure-native:web/v20160801:WebAppSlot", name, id, state, &resource, opts...)
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
	ClientAffinityEnabled        *bool                      `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled            *bool                      `pulumi:"clientCertEnabled"`
	CloningInfo                  *CloningInfo               `pulumi:"cloningInfo"`
	ContainerSize                *int                       `pulumi:"containerSize"`
	DailyMemoryTimeQuota         *int                       `pulumi:"dailyMemoryTimeQuota"`
	Enabled                      *bool                      `pulumi:"enabled"`
	ForceDnsRegistration         *bool                      `pulumi:"forceDnsRegistration"`
	HostNameSslStates            []HostNameSslState         `pulumi:"hostNameSslStates"`
	HostNamesDisabled            *bool                      `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile    *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	HttpsOnly                    *bool                      `pulumi:"httpsOnly"`
	Identity                     *ManagedServiceIdentity    `pulumi:"identity"`
	Kind                         *string                    `pulumi:"kind"`
	Location                     *string                    `pulumi:"location"`
	Name                         string                     `pulumi:"name"`
	Reserved                     *bool                      `pulumi:"reserved"`
	ResourceGroupName            string                     `pulumi:"resourceGroupName"`
	ScmSiteAlsoStopped           *bool                      `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId                 *string                    `pulumi:"serverFarmId"`
	SiteConfig                   *SiteConfig                `pulumi:"siteConfig"`
	SkipCustomDomainVerification *bool                      `pulumi:"skipCustomDomainVerification"`
	SkipDnsRegistration          *bool                      `pulumi:"skipDnsRegistration"`
	Slot                         *string                    `pulumi:"slot"`
	SnapshotInfo                 *SnapshotRecoveryRequest   `pulumi:"snapshotInfo"`
	Tags                         map[string]string          `pulumi:"tags"`
	TtlInSeconds                 *string                    `pulumi:"ttlInSeconds"`
}


type WebAppSlotArgs struct {
	ClientAffinityEnabled        pulumi.BoolPtrInput
	ClientCertEnabled            pulumi.BoolPtrInput
	CloningInfo                  CloningInfoPtrInput
	ContainerSize                pulumi.IntPtrInput
	DailyMemoryTimeQuota         pulumi.IntPtrInput
	Enabled                      pulumi.BoolPtrInput
	ForceDnsRegistration         pulumi.BoolPtrInput
	HostNameSslStates            HostNameSslStateArrayInput
	HostNamesDisabled            pulumi.BoolPtrInput
	HostingEnvironmentProfile    HostingEnvironmentProfilePtrInput
	HttpsOnly                    pulumi.BoolPtrInput
	Identity                     ManagedServiceIdentityPtrInput
	Kind                         pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringInput
	Reserved                     pulumi.BoolPtrInput
	ResourceGroupName            pulumi.StringInput
	ScmSiteAlsoStopped           pulumi.BoolPtrInput
	ServerFarmId                 pulumi.StringPtrInput
	SiteConfig                   SiteConfigPtrInput
	SkipCustomDomainVerification pulumi.BoolPtrInput
	SkipDnsRegistration          pulumi.BoolPtrInput
	Slot                         pulumi.StringPtrInput
	SnapshotInfo                 SnapshotRecoveryRequestPtrInput
	Tags                         pulumi.StringMapInput
	TtlInSeconds                 pulumi.StringPtrInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*WebAppSlotInput)(nil)).Elem(), &WebAppSlot{})
	pulumi.RegisterOutputType(WebAppSlotOutput{})
}
