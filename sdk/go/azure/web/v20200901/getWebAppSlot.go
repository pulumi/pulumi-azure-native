


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSlot(ctx *pulumi.Context, args *LookupWebAppSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSlotResult, error) {
	var rv LookupWebAppSlotResult
	err := ctx.Invoke("azure-native:web/v20200901:getWebAppSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebAppSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppSlotResult struct {
	AvailabilityState           string                             `pulumi:"availabilityState"`
	ClientAffinityEnabled       *bool                              `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled           *bool                              `pulumi:"clientCertEnabled"`
	ClientCertExclusionPaths    *string                            `pulumi:"clientCertExclusionPaths"`
	ClientCertMode              *string                            `pulumi:"clientCertMode"`
	ContainerSize               *int                               `pulumi:"containerSize"`
	CustomDomainVerificationId  *string                            `pulumi:"customDomainVerificationId"`
	DailyMemoryTimeQuota        *int                               `pulumi:"dailyMemoryTimeQuota"`
	DefaultHostName             string                             `pulumi:"defaultHostName"`
	Enabled                     *bool                              `pulumi:"enabled"`
	EnabledHostNames            []string                           `pulumi:"enabledHostNames"`
	HostNameSslStates           []HostNameSslStateResponse         `pulumi:"hostNameSslStates"`
	HostNames                   []string                           `pulumi:"hostNames"`
	HostNamesDisabled           *bool                              `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile   *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	HttpsOnly                   *bool                              `pulumi:"httpsOnly"`
	HyperV                      *bool                              `pulumi:"hyperV"`
	Id                          string                             `pulumi:"id"`
	Identity                    *ManagedServiceIdentityResponse    `pulumi:"identity"`
	InProgressOperationId       string                             `pulumi:"inProgressOperationId"`
	IsDefaultContainer          bool                               `pulumi:"isDefaultContainer"`
	IsXenon                     *bool                              `pulumi:"isXenon"`
	Kind                        *string                            `pulumi:"kind"`
	LastModifiedTimeUtc         string                             `pulumi:"lastModifiedTimeUtc"`
	Location                    string                             `pulumi:"location"`
	MaxNumberOfWorkers          int                                `pulumi:"maxNumberOfWorkers"`
	Name                        string                             `pulumi:"name"`
	OutboundIpAddresses         string                             `pulumi:"outboundIpAddresses"`
	PossibleOutboundIpAddresses string                             `pulumi:"possibleOutboundIpAddresses"`
	RedundancyMode              *string                            `pulumi:"redundancyMode"`
	RepositorySiteName          string                             `pulumi:"repositorySiteName"`
	Reserved                    *bool                              `pulumi:"reserved"`
	ResourceGroup               string                             `pulumi:"resourceGroup"`
	ScmSiteAlsoStopped          *bool                              `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId                *string                            `pulumi:"serverFarmId"`
	SiteConfig                  *SiteConfigResponse                `pulumi:"siteConfig"`
	SlotSwapStatus              SlotSwapStatusResponse             `pulumi:"slotSwapStatus"`
	State                       string                             `pulumi:"state"`
	SuspendedTill               string                             `pulumi:"suspendedTill"`
	SystemData                  SystemDataResponse                 `pulumi:"systemData"`
	Tags                        map[string]string                  `pulumi:"tags"`
	TargetSwapSlot              string                             `pulumi:"targetSwapSlot"`
	TrafficManagerHostNames     []string                           `pulumi:"trafficManagerHostNames"`
	Type                        string                             `pulumi:"type"`
	UsageState                  string                             `pulumi:"usageState"`
}


func (val *LookupWebAppSlotResult) Defaults() *LookupWebAppSlotResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HyperV) {
		hyperV_ := false
		tmp.HyperV = &hyperV_
	}
	if isZero(tmp.IsXenon) {
		isXenon_ := false
		tmp.IsXenon = &isXenon_
	}
	if isZero(tmp.Reserved) {
		reserved_ := false
		tmp.Reserved = &reserved_
	}
	if isZero(tmp.ScmSiteAlsoStopped) {
		scmSiteAlsoStopped_ := false
		tmp.ScmSiteAlsoStopped = &scmSiteAlsoStopped_
	}
	tmp.SiteConfig = tmp.SiteConfig.Defaults()

	return &tmp
}
