


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteSlot(ctx *pulumi.Context, args *LookupSiteSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteSlotResult, error) {
	var rv LookupSiteSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSlotArgs struct {
	Name                string  `pulumi:"name"`
	PropertiesToInclude *string `pulumi:"propertiesToInclude"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	Slot                string  `pulumi:"slot"`
}


type LookupSiteSlotResult struct {
	AvailabilityState         string                             `pulumi:"availabilityState"`
	ClientAffinityEnabled     *bool                              `pulumi:"clientAffinityEnabled"`
	ClientCertEnabled         *bool                              `pulumi:"clientCertEnabled"`
	CloningInfo               *CloningInfoResponse               `pulumi:"cloningInfo"`
	ContainerSize             *int                               `pulumi:"containerSize"`
	DefaultHostName           string                             `pulumi:"defaultHostName"`
	Enabled                   *bool                              `pulumi:"enabled"`
	EnabledHostNames          []string                           `pulumi:"enabledHostNames"`
	GatewaySiteName           *string                            `pulumi:"gatewaySiteName"`
	HostNameSslStates         []HostNameSslStateResponse         `pulumi:"hostNameSslStates"`
	HostNames                 []string                           `pulumi:"hostNames"`
	HostNamesDisabled         *bool                              `pulumi:"hostNamesDisabled"`
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	Id                        *string                            `pulumi:"id"`
	IsDefaultContainer        bool                               `pulumi:"isDefaultContainer"`
	Kind                      *string                            `pulumi:"kind"`
	LastModifiedTimeUtc       string                             `pulumi:"lastModifiedTimeUtc"`
	Location                  string                             `pulumi:"location"`
	MaxNumberOfWorkers        *int                               `pulumi:"maxNumberOfWorkers"`
	MicroService              *string                            `pulumi:"microService"`
	Name                      *string                            `pulumi:"name"`
	OutboundIpAddresses       string                             `pulumi:"outboundIpAddresses"`
	PremiumAppDeployed        bool                               `pulumi:"premiumAppDeployed"`
	RepositorySiteName        string                             `pulumi:"repositorySiteName"`
	ResourceGroup             string                             `pulumi:"resourceGroup"`
	ScmSiteAlsoStopped        *bool                              `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId              *string                            `pulumi:"serverFarmId"`
	SiteConfig                *SiteConfigResponse                `pulumi:"siteConfig"`
	State                     string                             `pulumi:"state"`
	Tags                      map[string]string                  `pulumi:"tags"`
	TargetSwapSlot            string                             `pulumi:"targetSwapSlot"`
	TrafficManagerHostNames   []string                           `pulumi:"trafficManagerHostNames"`
	Type                      *string                            `pulumi:"type"`
	UsageState                string                             `pulumi:"usageState"`
}

func LookupSiteSlotOutput(ctx *pulumi.Context, args LookupSiteSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteSlotResult, error) {
			args := v.(LookupSiteSlotArgs)
			r, err := LookupSiteSlot(ctx, &args, opts...)
			var s LookupSiteSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteSlotResultOutput)
}

type LookupSiteSlotOutputArgs struct {
	Name                pulumi.StringInput    `pulumi:"name"`
	PropertiesToInclude pulumi.StringPtrInput `pulumi:"propertiesToInclude"`
	ResourceGroupName   pulumi.StringInput    `pulumi:"resourceGroupName"`
	Slot                pulumi.StringInput    `pulumi:"slot"`
}

func (LookupSiteSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSlotArgs)(nil)).Elem()
}


type LookupSiteSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSlotResult)(nil)).Elem()
}

func (o LookupSiteSlotResultOutput) ToLookupSiteSlotResultOutput() LookupSiteSlotResultOutput {
	return o
}

func (o LookupSiteSlotResultOutput) ToLookupSiteSlotResultOutputWithContext(ctx context.Context) LookupSiteSlotResultOutput {
	return o
}

func (o LookupSiteSlotResultOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.AvailabilityState }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *bool { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSlotResultOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSlotResultOutput) CloningInfo() CloningInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *CloningInfoResponse { return v.CloningInfo }).(CloningInfoResponsePtrOutput)
}

func (o LookupSiteSlotResultOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *int { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

func (o LookupSiteSlotResultOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.DefaultHostName }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSlotResultOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) []string { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupSiteSlotResultOutput) GatewaySiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *string { return v.GatewaySiteName }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotResultOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) []HostNameSslStateResponse { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

func (o LookupSiteSlotResultOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) []string { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o LookupSiteSlotResultOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *bool { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSlotResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o LookupSiteSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotResultOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) bool { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

func (o LookupSiteSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) MaxNumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *int { return v.MaxNumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o LookupSiteSlotResultOutput) MicroService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *string { return v.MicroService }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotResultOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) PremiumAppDeployed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) bool { return v.PremiumAppDeployed }).(pulumi.BoolOutput)
}

func (o LookupSiteSlotResultOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.RepositorySiteName }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *bool { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSlotResultOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *string { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotResultOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *SiteConfigResponse { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

func (o LookupSiteSlotResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteSlotResultOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

func (o LookupSiteSlotResultOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) []string { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupSiteSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotResultOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotResult) string { return v.UsageState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteSlotResultOutput{})
}
