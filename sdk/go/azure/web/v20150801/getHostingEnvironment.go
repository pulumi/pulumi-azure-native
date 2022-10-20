


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupHostingEnvironment(ctx *pulumi.Context, args *LookupHostingEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupHostingEnvironmentResult, error) {
	var rv LookupHostingEnvironmentResult
	err := ctx.Invoke("azure-native:web/v20150801:getHostingEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostingEnvironmentArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupHostingEnvironmentResult struct {
	AllowedMultiSizes         *string                             `pulumi:"allowedMultiSizes"`
	AllowedWorkerSizes        *string                             `pulumi:"allowedWorkerSizes"`
	ApiManagementAccountId    *string                             `pulumi:"apiManagementAccountId"`
	ClusterSettings           []NameValuePairResponse             `pulumi:"clusterSettings"`
	DatabaseEdition           *string                             `pulumi:"databaseEdition"`
	DatabaseServiceObjective  *string                             `pulumi:"databaseServiceObjective"`
	DnsSuffix                 *string                             `pulumi:"dnsSuffix"`
	EnvironmentCapacities     []StampCapacityResponse             `pulumi:"environmentCapacities"`
	EnvironmentIsHealthy      *bool                               `pulumi:"environmentIsHealthy"`
	EnvironmentStatus         *string                             `pulumi:"environmentStatus"`
	Id                        *string                             `pulumi:"id"`
	InternalLoadBalancingMode *string                             `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount         *int                                `pulumi:"ipsslAddressCount"`
	Kind                      *string                             `pulumi:"kind"`
	LastAction                *string                             `pulumi:"lastAction"`
	LastActionResult          *string                             `pulumi:"lastActionResult"`
	Location                  string                              `pulumi:"location"`
	MaximumNumberOfMachines   *int                                `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount            *int                                `pulumi:"multiRoleCount"`
	MultiSize                 *string                             `pulumi:"multiSize"`
	Name                      *string                             `pulumi:"name"`
	NetworkAccessControlList  []NetworkAccessControlEntryResponse `pulumi:"networkAccessControlList"`
	ProvisioningState         *string                             `pulumi:"provisioningState"`
	ResourceGroup             *string                             `pulumi:"resourceGroup"`
	Status                    string                              `pulumi:"status"`
	SubscriptionId            *string                             `pulumi:"subscriptionId"`
	Suspended                 *bool                               `pulumi:"suspended"`
	Tags                      map[string]string                   `pulumi:"tags"`
	Type                      *string                             `pulumi:"type"`
	UpgradeDomains            *int                                `pulumi:"upgradeDomains"`
	VipMappings               []VirtualIPMappingResponse          `pulumi:"vipMappings"`
	VirtualNetwork            *VirtualNetworkProfileResponse      `pulumi:"virtualNetwork"`
	VnetName                  *string                             `pulumi:"vnetName"`
	VnetResourceGroupName     *string                             `pulumi:"vnetResourceGroupName"`
	VnetSubnetName            *string                             `pulumi:"vnetSubnetName"`
	WorkerPools               []WorkerPoolResponse                `pulumi:"workerPools"`
}

func LookupHostingEnvironmentOutput(ctx *pulumi.Context, args LookupHostingEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupHostingEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostingEnvironmentResult, error) {
			args := v.(LookupHostingEnvironmentArgs)
			r, err := LookupHostingEnvironment(ctx, &args, opts...)
			var s LookupHostingEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostingEnvironmentResultOutput)
}

type LookupHostingEnvironmentOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHostingEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostingEnvironmentArgs)(nil)).Elem()
}


type LookupHostingEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupHostingEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostingEnvironmentResult)(nil)).Elem()
}

func (o LookupHostingEnvironmentResultOutput) ToLookupHostingEnvironmentResultOutput() LookupHostingEnvironmentResultOutput {
	return o
}

func (o LookupHostingEnvironmentResultOutput) ToLookupHostingEnvironmentResultOutputWithContext(ctx context.Context) LookupHostingEnvironmentResultOutput {
	return o
}

func (o LookupHostingEnvironmentResultOutput) AllowedMultiSizes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.AllowedMultiSizes }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) AllowedWorkerSizes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.AllowedWorkerSizes }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) ApiManagementAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.ApiManagementAccountId }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) ClusterSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) []NameValuePairResponse { return v.ClusterSettings }).(NameValuePairResponseArrayOutput)
}

func (o LookupHostingEnvironmentResultOutput) DatabaseEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.DatabaseEdition }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) DatabaseServiceObjective() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.DatabaseServiceObjective }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) EnvironmentCapacities() StampCapacityResponseArrayOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) []StampCapacityResponse { return v.EnvironmentCapacities }).(StampCapacityResponseArrayOutput)
}

func (o LookupHostingEnvironmentResultOutput) EnvironmentIsHealthy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *bool { return v.EnvironmentIsHealthy }).(pulumi.BoolPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) EnvironmentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.EnvironmentStatus }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) InternalLoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.InternalLoadBalancingMode }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) IpsslAddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *int { return v.IpsslAddressCount }).(pulumi.IntPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) LastAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.LastAction }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) LastActionResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.LastActionResult }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupHostingEnvironmentResultOutput) MaximumNumberOfMachines() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *int { return v.MaximumNumberOfMachines }).(pulumi.IntPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) MultiRoleCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *int { return v.MultiRoleCount }).(pulumi.IntPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) MultiSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.MultiSize }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) NetworkAccessControlList() NetworkAccessControlEntryResponseArrayOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) []NetworkAccessControlEntryResponse {
		return v.NetworkAccessControlList
	}).(NetworkAccessControlEntryResponseArrayOutput)
}

func (o LookupHostingEnvironmentResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupHostingEnvironmentResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) Suspended() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *bool { return v.Suspended }).(pulumi.BoolPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupHostingEnvironmentResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) UpgradeDomains() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *int { return v.UpgradeDomains }).(pulumi.IntPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) VipMappings() VirtualIPMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) []VirtualIPMappingResponse { return v.VipMappings }).(VirtualIPMappingResponseArrayOutput)
}

func (o LookupHostingEnvironmentResultOutput) VirtualNetwork() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *VirtualNetworkProfileResponse { return v.VirtualNetwork }).(VirtualNetworkProfileResponsePtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) VnetResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.VnetResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) VnetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) *string { return v.VnetSubnetName }).(pulumi.StringPtrOutput)
}

func (o LookupHostingEnvironmentResultOutput) WorkerPools() WorkerPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupHostingEnvironmentResult) []WorkerPoolResponse { return v.WorkerPools }).(WorkerPoolResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostingEnvironmentResultOutput{})
}
