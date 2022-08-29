


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebApp(ctx *pulumi.Context, args *LookupWebAppArgs, opts ...pulumi.InvokeOption) (*LookupWebAppResult, error) {
	var rv LookupWebAppResult
	err := ctx.Invoke("azure-native:web/v20210201:getWebApp", args, &rv, opts...)
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
	ClientCertExclusionPaths    *string                            `pulumi:"clientCertExclusionPaths"`
	ClientCertMode              *string                            `pulumi:"clientCertMode"`
	ContainerSize               *int                               `pulumi:"containerSize"`
	CustomDomainVerificationId  *string                            `pulumi:"customDomainVerificationId"`
	DailyMemoryTimeQuota        *int                               `pulumi:"dailyMemoryTimeQuota"`
	DefaultHostName             string                             `pulumi:"defaultHostName"`
	Enabled                     *bool                              `pulumi:"enabled"`
	EnabledHostNames            []string                           `pulumi:"enabledHostNames"`
	ExtendedLocation            *ExtendedLocationResponse          `pulumi:"extendedLocation"`
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
	KeyVaultReferenceIdentity   *string                            `pulumi:"keyVaultReferenceIdentity"`
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
	StorageAccountRequired      *bool                              `pulumi:"storageAccountRequired"`
	SuspendedTill               string                             `pulumi:"suspendedTill"`
	Tags                        map[string]string                  `pulumi:"tags"`
	TargetSwapSlot              string                             `pulumi:"targetSwapSlot"`
	TrafficManagerHostNames     []string                           `pulumi:"trafficManagerHostNames"`
	Type                        string                             `pulumi:"type"`
	UsageState                  string                             `pulumi:"usageState"`
	VirtualNetworkSubnetId      *string                            `pulumi:"virtualNetworkSubnetId"`
}


func (val *LookupWebAppResult) Defaults() *LookupWebAppResult {
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

func LookupWebAppOutput(ctx *pulumi.Context, args LookupWebAppOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppResult, error) {
			args := v.(LookupWebAppArgs)
			r, err := LookupWebApp(ctx, &args, opts...)
			var s LookupWebAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppResultOutput)
}

type LookupWebAppOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppArgs)(nil)).Elem()
}


type LookupWebAppResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppResult)(nil)).Elem()
}

func (o LookupWebAppResultOutput) ToLookupWebAppResultOutput() LookupWebAppResultOutput {
	return o
}

func (o LookupWebAppResultOutput) ToLookupWebAppResultOutputWithContext(ctx context.Context) LookupWebAppResultOutput {
	return o
}

func (o LookupWebAppResultOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.AvailabilityState }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) ClientCertExclusionPaths() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.ClientCertExclusionPaths }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppResultOutput) ClientCertMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.ClientCertMode }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppResultOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *int { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppResultOutput) CustomDomainVerificationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.CustomDomainVerificationId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppResultOutput) DailyMemoryTimeQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *int { return v.DailyMemoryTimeQuota }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppResultOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.DefaultHostName }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []string { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupWebAppResultOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []HostNameSslStateResponse { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

func (o LookupWebAppResultOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []string { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppResultOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o LookupWebAppResultOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) HyperV() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.HyperV }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupWebAppResultOutput) InProgressOperationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.InProgressOperationId }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppResult) bool { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

func (o LookupWebAppResultOutput) IsXenon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.IsXenon }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) MaxNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebAppResult) int { return v.MaxNumberOfWorkers }).(pulumi.IntOutput)
}

func (o LookupWebAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) PossibleOutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.PossibleOutboundIpAddresses }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) RedundancyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.RedundancyMode }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppResultOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.RepositorySiteName }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.Reserved }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppResultOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *SiteConfigResponse { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

func (o LookupWebAppResultOutput) SlotSwapStatus() SlotSwapStatusResponseOutput {
	return o.ApplyT(func(v LookupWebAppResult) SlotSwapStatusResponse { return v.SlotSwapStatus }).(SlotSwapStatusResponseOutput)
}

func (o LookupWebAppResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) StorageAccountRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *bool { return v.StorageAccountRequired }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppResultOutput) SuspendedTill() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.SuspendedTill }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebAppResultOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []string { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.UsageState }).(pulumi.StringOutput)
}

func (o LookupWebAppResultOutput) VirtualNetworkSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppResult) *string { return v.VirtualNetworkSubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppResultOutput{})
}
