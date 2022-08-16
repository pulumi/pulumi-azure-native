


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSlot(ctx *pulumi.Context, args *LookupWebAppSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSlotResult, error) {
	var rv LookupWebAppSlotResult
	err := ctx.Invoke("azure-native:web/v20180201:getWebAppSlot", args, &rv, opts...)
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
	ContainerSize               *int                               `pulumi:"containerSize"`
	DailyMemoryTimeQuota        *int                               `pulumi:"dailyMemoryTimeQuota"`
	DefaultHostName             string                             `pulumi:"defaultHostName"`
	Enabled                     *bool                              `pulumi:"enabled"`
	EnabledHostNames            []string                           `pulumi:"enabledHostNames"`
	GeoDistributions            []GeoDistributionResponse          `pulumi:"geoDistributions"`
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

func LookupWebAppSlotOutput(ctx *pulumi.Context, args LookupWebAppSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSlotResult, error) {
			args := v.(LookupWebAppSlotArgs)
			r, err := LookupWebAppSlot(ctx, &args, opts...)
			var s LookupWebAppSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSlotResultOutput)
}

type LookupWebAppSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSlotArgs)(nil)).Elem()
}


type LookupWebAppSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSlotResult)(nil)).Elem()
}

func (o LookupWebAppSlotResultOutput) ToLookupWebAppSlotResultOutput() LookupWebAppSlotResultOutput {
	return o
}

func (o LookupWebAppSlotResultOutput) ToLookupWebAppSlotResultOutputWithContext(ctx context.Context) LookupWebAppSlotResultOutput {
	return o
}

func (o LookupWebAppSlotResultOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.AvailabilityState }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) ClientCertExclusionPaths() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *string { return v.ClientCertExclusionPaths }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSlotResultOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *int { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppSlotResultOutput) DailyMemoryTimeQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *int { return v.DailyMemoryTimeQuota }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppSlotResultOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.DefaultHostName }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) []string { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppSlotResultOutput) GeoDistributions() GeoDistributionResponseArrayOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) []GeoDistributionResponse { return v.GeoDistributions }).(GeoDistributionResponseArrayOutput)
}

func (o LookupWebAppSlotResultOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) []HostNameSslStateResponse { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

func (o LookupWebAppSlotResultOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) []string { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppSlotResultOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o LookupWebAppSlotResultOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) HyperV() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.HyperV }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupWebAppSlotResultOutput) InProgressOperationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.InProgressOperationId }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) bool { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

func (o LookupWebAppSlotResultOutput) IsXenon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.IsXenon }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSlotResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) MaxNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) int { return v.MaxNumberOfWorkers }).(pulumi.IntOutput)
}

func (o LookupWebAppSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) PossibleOutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.PossibleOutboundIpAddresses }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) RedundancyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *string { return v.RedundancyMode }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSlotResultOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.RepositorySiteName }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.Reserved }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *bool { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSlotResultOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *string { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSlotResultOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) *SiteConfigResponse { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

func (o LookupWebAppSlotResultOutput) SlotSwapStatus() SlotSwapStatusResponseOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) SlotSwapStatusResponse { return v.SlotSwapStatus }).(SlotSwapStatusResponseOutput)
}

func (o LookupWebAppSlotResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) SuspendedTill() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.SuspendedTill }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebAppSlotResultOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) []string { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotResultOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotResult) string { return v.UsageState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSlotResultOutput{})
}
