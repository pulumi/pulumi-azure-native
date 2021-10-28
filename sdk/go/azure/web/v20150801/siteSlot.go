


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteSlot struct {
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


func NewSiteSlot(ctx *pulumi.Context,
	name string, args *SiteSlotArgs, opts ...pulumi.ResourceOption) (*SiteSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:SiteSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteSlotState, opts ...pulumi.ResourceOption) (*SiteSlot, error) {
	var resource SiteSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteSlotState struct {
}

type SiteSlotState struct {
}

func (SiteSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSlotState)(nil)).Elem()
}

type siteSlotArgs struct {
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
	Name                         string                     `pulumi:"name"`
	ResourceGroupName            string                     `pulumi:"resourceGroupName"`
	ScmSiteAlsoStopped           *bool                      `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId                 *string                    `pulumi:"serverFarmId"`
	SiteConfig                   *SiteConfig                `pulumi:"siteConfig"`
	SkipCustomDomainVerification *string                    `pulumi:"skipCustomDomainVerification"`
	SkipDnsRegistration          *string                    `pulumi:"skipDnsRegistration"`
	Slot                         *string                    `pulumi:"slot"`
	Tags                         map[string]string          `pulumi:"tags"`
	TtlInSeconds                 *string                    `pulumi:"ttlInSeconds"`
	Type                         *string                    `pulumi:"type"`
}


type SiteSlotArgs struct {
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
	Name                         pulumi.StringInput
	ResourceGroupName            pulumi.StringInput
	ScmSiteAlsoStopped           pulumi.BoolPtrInput
	ServerFarmId                 pulumi.StringPtrInput
	SiteConfig                   SiteConfigPtrInput
	SkipCustomDomainVerification pulumi.StringPtrInput
	SkipDnsRegistration          pulumi.StringPtrInput
	Slot                         pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
	TtlInSeconds                 pulumi.StringPtrInput
	Type                         pulumi.StringPtrInput
}

func (SiteSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSlotArgs)(nil)).Elem()
}

type SiteSlotInput interface {
	pulumi.Input

	ToSiteSlotOutput() SiteSlotOutput
	ToSiteSlotOutputWithContext(ctx context.Context) SiteSlotOutput
}

func (*SiteSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSlot)(nil))
}

func (i *SiteSlot) ToSiteSlotOutput() SiteSlotOutput {
	return i.ToSiteSlotOutputWithContext(context.Background())
}

func (i *SiteSlot) ToSiteSlotOutputWithContext(ctx context.Context) SiteSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteSlotOutput)
}

type SiteSlotOutput struct{ *pulumi.OutputState }

func (SiteSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSlot)(nil))
}

func (o SiteSlotOutput) ToSiteSlotOutput() SiteSlotOutput {
	return o
}

func (o SiteSlotOutput) ToSiteSlotOutputWithContext(ctx context.Context) SiteSlotOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteSlotInput)(nil)).Elem(), &SiteSlot{})
	pulumi.RegisterOutputType(SiteSlotOutput{})
}
