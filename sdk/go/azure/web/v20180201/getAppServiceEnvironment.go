


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceEnvironment(ctx *pulumi.Context, args *LookupAppServiceEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceEnvironmentResult, error) {
	var rv LookupAppServiceEnvironmentResult
	err := ctx.Invoke("azure-native:web/v20180201:getAppServiceEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceEnvironmentArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAppServiceEnvironmentResult struct {
	AllowedMultiSizes          string                              `pulumi:"allowedMultiSizes"`
	AllowedWorkerSizes         string                              `pulumi:"allowedWorkerSizes"`
	ApiManagementAccountId     *string                             `pulumi:"apiManagementAccountId"`
	ClusterSettings            []NameValuePairResponse             `pulumi:"clusterSettings"`
	DatabaseEdition            string                              `pulumi:"databaseEdition"`
	DatabaseServiceObjective   string                              `pulumi:"databaseServiceObjective"`
	DefaultFrontEndScaleFactor int                                 `pulumi:"defaultFrontEndScaleFactor"`
	DnsSuffix                  *string                             `pulumi:"dnsSuffix"`
	DynamicCacheEnabled        *bool                               `pulumi:"dynamicCacheEnabled"`
	EnvironmentCapacities      []StampCapacityResponse             `pulumi:"environmentCapacities"`
	EnvironmentIsHealthy       bool                                `pulumi:"environmentIsHealthy"`
	EnvironmentStatus          string                              `pulumi:"environmentStatus"`
	FrontEndScaleFactor        *int                                `pulumi:"frontEndScaleFactor"`
	HasLinuxWorkers            *bool                               `pulumi:"hasLinuxWorkers"`
	Id                         string                              `pulumi:"id"`
	InternalLoadBalancingMode  *string                             `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount          *int                                `pulumi:"ipsslAddressCount"`
	Kind                       *string                             `pulumi:"kind"`
	LastAction                 string                              `pulumi:"lastAction"`
	LastActionResult           string                              `pulumi:"lastActionResult"`
	Location                   string                              `pulumi:"location"`
	MaximumNumberOfMachines    int                                 `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount             *int                                `pulumi:"multiRoleCount"`
	MultiSize                  *string                             `pulumi:"multiSize"`
	Name                       string                              `pulumi:"name"`
	NetworkAccessControlList   []NetworkAccessControlEntryResponse `pulumi:"networkAccessControlList"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	ResourceGroup              string                              `pulumi:"resourceGroup"`
	SslCertKeyVaultId          *string                             `pulumi:"sslCertKeyVaultId"`
	SslCertKeyVaultSecretName  *string                             `pulumi:"sslCertKeyVaultSecretName"`
	Status                     string                              `pulumi:"status"`
	SubscriptionId             string                              `pulumi:"subscriptionId"`
	Suspended                  *bool                               `pulumi:"suspended"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	UpgradeDomains             int                                 `pulumi:"upgradeDomains"`
	UserWhitelistedIpRanges    []string                            `pulumi:"userWhitelistedIpRanges"`
	VipMappings                []VirtualIPMappingResponse          `pulumi:"vipMappings"`
	VirtualNetwork             VirtualNetworkProfileResponse       `pulumi:"virtualNetwork"`
	VnetName                   *string                             `pulumi:"vnetName"`
	VnetResourceGroupName      *string                             `pulumi:"vnetResourceGroupName"`
	VnetSubnetName             *string                             `pulumi:"vnetSubnetName"`
	WorkerPools                []WorkerPoolResponse                `pulumi:"workerPools"`
}

func LookupAppServiceEnvironmentOutput(ctx *pulumi.Context, args LookupAppServiceEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupAppServiceEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServiceEnvironmentResult, error) {
			args := v.(LookupAppServiceEnvironmentArgs)
			r, err := LookupAppServiceEnvironment(ctx, &args, opts...)
			var s LookupAppServiceEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServiceEnvironmentResultOutput)
}

type LookupAppServiceEnvironmentOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServiceEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceEnvironmentArgs)(nil)).Elem()
}


type LookupAppServiceEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupAppServiceEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceEnvironmentResult)(nil)).Elem()
}

func (o LookupAppServiceEnvironmentResultOutput) ToLookupAppServiceEnvironmentResultOutput() LookupAppServiceEnvironmentResultOutput {
	return o
}

func (o LookupAppServiceEnvironmentResultOutput) ToLookupAppServiceEnvironmentResultOutputWithContext(ctx context.Context) LookupAppServiceEnvironmentResultOutput {
	return o
}

func (o LookupAppServiceEnvironmentResultOutput) AllowedMultiSizes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.AllowedMultiSizes }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) AllowedWorkerSizes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.AllowedWorkerSizes }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) ApiManagementAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.ApiManagementAccountId }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) ClusterSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) []NameValuePairResponse { return v.ClusterSettings }).(NameValuePairResponseArrayOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) DatabaseEdition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.DatabaseEdition }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) DatabaseServiceObjective() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.DatabaseServiceObjective }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) DefaultFrontEndScaleFactor() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) int { return v.DefaultFrontEndScaleFactor }).(pulumi.IntOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) DynamicCacheEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *bool { return v.DynamicCacheEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) EnvironmentCapacities() StampCapacityResponseArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) []StampCapacityResponse { return v.EnvironmentCapacities }).(StampCapacityResponseArrayOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) EnvironmentIsHealthy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) bool { return v.EnvironmentIsHealthy }).(pulumi.BoolOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) EnvironmentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.EnvironmentStatus }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) FrontEndScaleFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *int { return v.FrontEndScaleFactor }).(pulumi.IntPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) HasLinuxWorkers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *bool { return v.HasLinuxWorkers }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) InternalLoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.InternalLoadBalancingMode }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) IpsslAddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *int { return v.IpsslAddressCount }).(pulumi.IntPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) LastAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.LastAction }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) LastActionResult() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.LastActionResult }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) MaximumNumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) int { return v.MaximumNumberOfMachines }).(pulumi.IntOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) MultiRoleCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *int { return v.MultiRoleCount }).(pulumi.IntPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) MultiSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.MultiSize }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) NetworkAccessControlList() NetworkAccessControlEntryResponseArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) []NetworkAccessControlEntryResponse {
		return v.NetworkAccessControlList
	}).(NetworkAccessControlEntryResponseArrayOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) SslCertKeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.SslCertKeyVaultId }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) SslCertKeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.SslCertKeyVaultSecretName }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Suspended() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *bool { return v.Suspended }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) UpgradeDomains() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) int { return v.UpgradeDomains }).(pulumi.IntOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) UserWhitelistedIpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) []string { return v.UserWhitelistedIpRanges }).(pulumi.StringArrayOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) VipMappings() VirtualIPMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) []VirtualIPMappingResponse { return v.VipMappings }).(VirtualIPMappingResponseArrayOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) VirtualNetwork() VirtualNetworkProfileResponseOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) VirtualNetworkProfileResponse { return v.VirtualNetwork }).(VirtualNetworkProfileResponseOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) VnetResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.VnetResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) VnetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) *string { return v.VnetSubnetName }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceEnvironmentResultOutput) WorkerPools() WorkerPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupAppServiceEnvironmentResult) []WorkerPoolResponse { return v.WorkerPools }).(WorkerPoolResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServiceEnvironmentResultOutput{})
}
