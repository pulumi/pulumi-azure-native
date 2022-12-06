


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSite(ctx *pulumi.Context, args *LookupSiteArgs, opts ...pulumi.InvokeOption) (*LookupSiteResult, error) {
	var rv LookupSiteResult
	err := ctx.Invoke("azure-native:web/v20150801:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteArgs struct {
	Name                string  `pulumi:"name"`
	PropertiesToInclude *string `pulumi:"propertiesToInclude"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupSiteResult struct {
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

func LookupSiteOutput(ctx *pulumi.Context, args LookupSiteOutputArgs, opts ...pulumi.InvokeOption) LookupSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteResult, error) {
			args := v.(LookupSiteArgs)
			r, err := LookupSite(ctx, &args, opts...)
			var s LookupSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteResultOutput)
}

type LookupSiteOutputArgs struct {
	Name                pulumi.StringInput    `pulumi:"name"`
	PropertiesToInclude pulumi.StringPtrInput `pulumi:"propertiesToInclude"`
	ResourceGroupName   pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteArgs)(nil)).Elem()
}


type LookupSiteResultOutput struct{ *pulumi.OutputState }

func (LookupSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteResult)(nil)).Elem()
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutput() LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutputWithContext(ctx context.Context) LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.AvailabilityState }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *bool { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteResultOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteResultOutput) CloningInfo() CloningInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *CloningInfoResponse { return v.CloningInfo }).(CloningInfoResponsePtrOutput)
}

func (o LookupSiteResultOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *int { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

func (o LookupSiteResultOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.DefaultHostName }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteResultOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSiteResult) []string { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupSiteResultOutput) GatewaySiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.GatewaySiteName }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v LookupSiteResult) []HostNameSslStateResponse { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

func (o LookupSiteResultOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSiteResult) []string { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o LookupSiteResultOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *bool { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o LookupSiteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSiteResult) bool { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

func (o LookupSiteResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) MaxNumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *int { return v.MaxNumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o LookupSiteResultOutput) MicroService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.MicroService }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) PremiumAppDeployed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSiteResult) bool { return v.PremiumAppDeployed }).(pulumi.BoolOutput)
}

func (o LookupSiteResultOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.RepositorySiteName }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *bool { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteResultOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *SiteConfigResponse { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

func (o LookupSiteResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteResultOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSiteResult) []string { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupSiteResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.UsageState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteResultOutput{})
}
