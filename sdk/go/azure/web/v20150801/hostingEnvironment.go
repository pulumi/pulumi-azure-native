


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type HostingEnvironment struct {
	pulumi.CustomResourceState

	AllowedMultiSizes         pulumi.StringPtrOutput                       `pulumi:"allowedMultiSizes"`
	AllowedWorkerSizes        pulumi.StringPtrOutput                       `pulumi:"allowedWorkerSizes"`
	ApiManagementAccountId    pulumi.StringPtrOutput                       `pulumi:"apiManagementAccountId"`
	ClusterSettings           NameValuePairResponseArrayOutput             `pulumi:"clusterSettings"`
	DatabaseEdition           pulumi.StringPtrOutput                       `pulumi:"databaseEdition"`
	DatabaseServiceObjective  pulumi.StringPtrOutput                       `pulumi:"databaseServiceObjective"`
	DnsSuffix                 pulumi.StringPtrOutput                       `pulumi:"dnsSuffix"`
	EnvironmentCapacities     StampCapacityResponseArrayOutput             `pulumi:"environmentCapacities"`
	EnvironmentIsHealthy      pulumi.BoolPtrOutput                         `pulumi:"environmentIsHealthy"`
	EnvironmentStatus         pulumi.StringPtrOutput                       `pulumi:"environmentStatus"`
	InternalLoadBalancingMode pulumi.StringPtrOutput                       `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount         pulumi.IntPtrOutput                          `pulumi:"ipsslAddressCount"`
	Kind                      pulumi.StringPtrOutput                       `pulumi:"kind"`
	LastAction                pulumi.StringPtrOutput                       `pulumi:"lastAction"`
	LastActionResult          pulumi.StringPtrOutput                       `pulumi:"lastActionResult"`
	Location                  pulumi.StringOutput                          `pulumi:"location"`
	MaximumNumberOfMachines   pulumi.IntPtrOutput                          `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount            pulumi.IntPtrOutput                          `pulumi:"multiRoleCount"`
	MultiSize                 pulumi.StringPtrOutput                       `pulumi:"multiSize"`
	Name                      pulumi.StringPtrOutput                       `pulumi:"name"`
	NetworkAccessControlList  NetworkAccessControlEntryResponseArrayOutput `pulumi:"networkAccessControlList"`
	ProvisioningState         pulumi.StringPtrOutput                       `pulumi:"provisioningState"`
	ResourceGroup             pulumi.StringPtrOutput                       `pulumi:"resourceGroup"`
	Status                    pulumi.StringOutput                          `pulumi:"status"`
	SubscriptionId            pulumi.StringPtrOutput                       `pulumi:"subscriptionId"`
	Suspended                 pulumi.BoolPtrOutput                         `pulumi:"suspended"`
	Tags                      pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                      pulumi.StringPtrOutput                       `pulumi:"type"`
	UpgradeDomains            pulumi.IntPtrOutput                          `pulumi:"upgradeDomains"`
	VipMappings               VirtualIPMappingResponseArrayOutput          `pulumi:"vipMappings"`
	VirtualNetwork            VirtualNetworkProfileResponsePtrOutput       `pulumi:"virtualNetwork"`
	VnetName                  pulumi.StringPtrOutput                       `pulumi:"vnetName"`
	VnetResourceGroupName     pulumi.StringPtrOutput                       `pulumi:"vnetResourceGroupName"`
	VnetSubnetName            pulumi.StringPtrOutput                       `pulumi:"vnetSubnetName"`
	WorkerPools               WorkerPoolResponseArrayOutput                `pulumi:"workerPools"`
}


func NewHostingEnvironment(ctx *pulumi.Context,
	name string, args *HostingEnvironmentArgs, opts ...pulumi.ResourceOption) (*HostingEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:HostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:HostingEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource HostingEnvironment
	err := ctx.RegisterResource("azure-native:web/v20150801:HostingEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHostingEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostingEnvironmentState, opts ...pulumi.ResourceOption) (*HostingEnvironment, error) {
	var resource HostingEnvironment
	err := ctx.ReadResource("azure-native:web/v20150801:HostingEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hostingEnvironmentState struct {
}

type HostingEnvironmentState struct {
}

func (HostingEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostingEnvironmentState)(nil)).Elem()
}

type hostingEnvironmentArgs struct {
	AllowedMultiSizes         *string                     `pulumi:"allowedMultiSizes"`
	AllowedWorkerSizes        *string                     `pulumi:"allowedWorkerSizes"`
	ApiManagementAccountId    *string                     `pulumi:"apiManagementAccountId"`
	ClusterSettings           []NameValuePair             `pulumi:"clusterSettings"`
	DatabaseEdition           *string                     `pulumi:"databaseEdition"`
	DatabaseServiceObjective  *string                     `pulumi:"databaseServiceObjective"`
	DnsSuffix                 *string                     `pulumi:"dnsSuffix"`
	EnvironmentCapacities     []StampCapacity             `pulumi:"environmentCapacities"`
	EnvironmentIsHealthy      *bool                       `pulumi:"environmentIsHealthy"`
	EnvironmentStatus         *string                     `pulumi:"environmentStatus"`
	Id                        *string                     `pulumi:"id"`
	InternalLoadBalancingMode *InternalLoadBalancingMode  `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount         *int                        `pulumi:"ipsslAddressCount"`
	Kind                      *string                     `pulumi:"kind"`
	LastAction                *string                     `pulumi:"lastAction"`
	LastActionResult          *string                     `pulumi:"lastActionResult"`
	Location                  *string                     `pulumi:"location"`
	MaximumNumberOfMachines   *int                        `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount            *int                        `pulumi:"multiRoleCount"`
	MultiSize                 *string                     `pulumi:"multiSize"`
	Name                      *string                     `pulumi:"name"`
	NetworkAccessControlList  []NetworkAccessControlEntry `pulumi:"networkAccessControlList"`
	ProvisioningState         *ProvisioningState          `pulumi:"provisioningState"`
	ResourceGroup             *string                     `pulumi:"resourceGroup"`
	ResourceGroupName         string                      `pulumi:"resourceGroupName"`
	Status                    HostingEnvironmentStatus    `pulumi:"status"`
	SubscriptionId            *string                     `pulumi:"subscriptionId"`
	Suspended                 *bool                       `pulumi:"suspended"`
	Tags                      map[string]string           `pulumi:"tags"`
	Type                      *string                     `pulumi:"type"`
	UpgradeDomains            *int                        `pulumi:"upgradeDomains"`
	VipMappings               []VirtualIPMapping          `pulumi:"vipMappings"`
	VirtualNetwork            *VirtualNetworkProfile      `pulumi:"virtualNetwork"`
	VnetName                  *string                     `pulumi:"vnetName"`
	VnetResourceGroupName     *string                     `pulumi:"vnetResourceGroupName"`
	VnetSubnetName            *string                     `pulumi:"vnetSubnetName"`
	WorkerPools               []WorkerPool                `pulumi:"workerPools"`
}


type HostingEnvironmentArgs struct {
	AllowedMultiSizes         pulumi.StringPtrInput
	AllowedWorkerSizes        pulumi.StringPtrInput
	ApiManagementAccountId    pulumi.StringPtrInput
	ClusterSettings           NameValuePairArrayInput
	DatabaseEdition           pulumi.StringPtrInput
	DatabaseServiceObjective  pulumi.StringPtrInput
	DnsSuffix                 pulumi.StringPtrInput
	EnvironmentCapacities     StampCapacityArrayInput
	EnvironmentIsHealthy      pulumi.BoolPtrInput
	EnvironmentStatus         pulumi.StringPtrInput
	Id                        pulumi.StringPtrInput
	InternalLoadBalancingMode InternalLoadBalancingModePtrInput
	IpsslAddressCount         pulumi.IntPtrInput
	Kind                      pulumi.StringPtrInput
	LastAction                pulumi.StringPtrInput
	LastActionResult          pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	MaximumNumberOfMachines   pulumi.IntPtrInput
	MultiRoleCount            pulumi.IntPtrInput
	MultiSize                 pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	NetworkAccessControlList  NetworkAccessControlEntryArrayInput
	ProvisioningState         ProvisioningStatePtrInput
	ResourceGroup             pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Status                    HostingEnvironmentStatusInput
	SubscriptionId            pulumi.StringPtrInput
	Suspended                 pulumi.BoolPtrInput
	Tags                      pulumi.StringMapInput
	Type                      pulumi.StringPtrInput
	UpgradeDomains            pulumi.IntPtrInput
	VipMappings               VirtualIPMappingArrayInput
	VirtualNetwork            VirtualNetworkProfilePtrInput
	VnetName                  pulumi.StringPtrInput
	VnetResourceGroupName     pulumi.StringPtrInput
	VnetSubnetName            pulumi.StringPtrInput
	WorkerPools               WorkerPoolArrayInput
}

func (HostingEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostingEnvironmentArgs)(nil)).Elem()
}

type HostingEnvironmentInput interface {
	pulumi.Input

	ToHostingEnvironmentOutput() HostingEnvironmentOutput
	ToHostingEnvironmentOutputWithContext(ctx context.Context) HostingEnvironmentOutput
}

func (*HostingEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironment)(nil)).Elem()
}

func (i *HostingEnvironment) ToHostingEnvironmentOutput() HostingEnvironmentOutput {
	return i.ToHostingEnvironmentOutputWithContext(context.Background())
}

func (i *HostingEnvironment) ToHostingEnvironmentOutputWithContext(ctx context.Context) HostingEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentOutput)
}

type HostingEnvironmentOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironment)(nil)).Elem()
}

func (o HostingEnvironmentOutput) ToHostingEnvironmentOutput() HostingEnvironmentOutput {
	return o
}

func (o HostingEnvironmentOutput) ToHostingEnvironmentOutputWithContext(ctx context.Context) HostingEnvironmentOutput {
	return o
}

func (o HostingEnvironmentOutput) AllowedMultiSizes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.AllowedMultiSizes }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) AllowedWorkerSizes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.AllowedWorkerSizes }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) ApiManagementAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.ApiManagementAccountId }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) ClusterSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v *HostingEnvironment) NameValuePairResponseArrayOutput { return v.ClusterSettings }).(NameValuePairResponseArrayOutput)
}

func (o HostingEnvironmentOutput) DatabaseEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.DatabaseEdition }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) DatabaseServiceObjective() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.DatabaseServiceObjective }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) EnvironmentCapacities() StampCapacityResponseArrayOutput {
	return o.ApplyT(func(v *HostingEnvironment) StampCapacityResponseArrayOutput { return v.EnvironmentCapacities }).(StampCapacityResponseArrayOutput)
}

func (o HostingEnvironmentOutput) EnvironmentIsHealthy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.BoolPtrOutput { return v.EnvironmentIsHealthy }).(pulumi.BoolPtrOutput)
}

func (o HostingEnvironmentOutput) EnvironmentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.EnvironmentStatus }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) InternalLoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.InternalLoadBalancingMode }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) IpsslAddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.IntPtrOutput { return v.IpsslAddressCount }).(pulumi.IntPtrOutput)
}

func (o HostingEnvironmentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) LastAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.LastAction }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) LastActionResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.LastActionResult }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o HostingEnvironmentOutput) MaximumNumberOfMachines() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.IntPtrOutput { return v.MaximumNumberOfMachines }).(pulumi.IntPtrOutput)
}

func (o HostingEnvironmentOutput) MultiRoleCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.IntPtrOutput { return v.MultiRoleCount }).(pulumi.IntPtrOutput)
}

func (o HostingEnvironmentOutput) MultiSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.MultiSize }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) NetworkAccessControlList() NetworkAccessControlEntryResponseArrayOutput {
	return o.ApplyT(func(v *HostingEnvironment) NetworkAccessControlEntryResponseArrayOutput {
		return v.NetworkAccessControlList
	}).(NetworkAccessControlEntryResponseArrayOutput)
}

func (o HostingEnvironmentOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o HostingEnvironmentOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) Suspended() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.BoolPtrOutput { return v.Suspended }).(pulumi.BoolPtrOutput)
}

func (o HostingEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o HostingEnvironmentOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) UpgradeDomains() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.IntPtrOutput { return v.UpgradeDomains }).(pulumi.IntPtrOutput)
}

func (o HostingEnvironmentOutput) VipMappings() VirtualIPMappingResponseArrayOutput {
	return o.ApplyT(func(v *HostingEnvironment) VirtualIPMappingResponseArrayOutput { return v.VipMappings }).(VirtualIPMappingResponseArrayOutput)
}

func (o HostingEnvironmentOutput) VirtualNetwork() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) VirtualNetworkProfileResponsePtrOutput { return v.VirtualNetwork }).(VirtualNetworkProfileResponsePtrOutput)
}

func (o HostingEnvironmentOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) VnetResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.VnetResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) VnetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironment) pulumi.StringPtrOutput { return v.VnetSubnetName }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentOutput) WorkerPools() WorkerPoolResponseArrayOutput {
	return o.ApplyT(func(v *HostingEnvironment) WorkerPoolResponseArrayOutput { return v.WorkerPools }).(WorkerPoolResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(HostingEnvironmentOutput{})
}
