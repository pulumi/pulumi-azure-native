


package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebApp(ctx *pulumi.Context, args *LookupWebAppArgs, opts ...pulumi.InvokeOption) (*LookupWebAppResult, error) {
	var rv LookupWebAppResult
	err := ctx.Invoke("azure-native:web/v20160801:getWebApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebAppArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppResult struct {
	AvailabilityState           string                             `pulumi:"availabilityState"`
	ClientAffinityEnabled       *bool                              `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled           *bool                              `pulumi:"clientCertEnabled"`
	ContainerSize               *int                               `pulumi:"containerSize"`
	DailyMemoryTimeQuota        *int                               `pulumi:"dailyMemoryTimeQuota"`
	DefaultHostName             string                             `pulumi:"defaultHostName"`
	Enabled                     *bool                              `pulumi:"enabled"`
	EnabledHostNames            []string                           `pulumi:"enabledHostNames"`
	HostNameSslStates           []HostNameSslStateResponse         `pulumi:"hostNameSslStates"`
	HostNames                   []string                           `pulumi:"hostNames"`
	HostNamesDisabled           *bool                              `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile   *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	HttpsOnly                   *bool                              `pulumi:"httpsOnly"`
	Id                          string                             `pulumi:"id"`
	Identity                    *ManagedServiceIdentityResponse    `pulumi:"identity"`
	IsDefaultContainer          bool                               `pulumi:"isDefaultContainer"`
	Kind                        *string                            `pulumi:"kind"`
	LastModifiedTimeUtc         string                             `pulumi:"lastModifiedTimeUtc"`
	Location                    string                             `pulumi:"location"`
	MaxNumberOfWorkers          int                                `pulumi:"maxNumberOfWorkers"`
	Name                        string                             `pulumi:"name"`
	OutboundIpAddresses         string                             `pulumi:"outboundIpAddresses"`
	PossibleOutboundIpAddresses string                             `pulumi:"possibleOutboundIpAddresses"`
	RepositorySiteName          string                             `pulumi:"repositorySiteName"`
	Reserved                    *bool                              `pulumi:"reserved"`
	ResourceGroup               string                             `pulumi:"resourceGroup"`
	ScmSiteAlsoStopped          *bool                              `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId                *string                            `pulumi:"serverFarmId"`
	SiteConfig                  *SiteConfigResponse                `pulumi:"siteConfig"`
	SlotSwapStatus              SlotSwapStatusResponse             `pulumi:"slotSwapStatus"`
	State                       string                             `pulumi:"state"`
	SuspendedTill               string                             `pulumi:"suspendedTill"`
	Tags                        map[string]string                  `pulumi:"tags"`
	TargetSwapSlot              string                             `pulumi:"targetSwapSlot"`
	TrafficManagerHostNames     []string                           `pulumi:"trafficManagerHostNames"`
	Type                        string                             `pulumi:"type"`
	UsageState                  string                             `pulumi:"usageState"`
}


func (val *LookupWebAppResult) Defaults() *LookupWebAppResult {
	if val == nil {
		return nil
	}
	tmp := *val
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
