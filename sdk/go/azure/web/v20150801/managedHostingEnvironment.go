


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ManagedHostingEnvironment struct {
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


func NewManagedHostingEnvironment(ctx *pulumi.Context,
	name string, args *ManagedHostingEnvironmentArgs, opts ...pulumi.ResourceOption) (*ManagedHostingEnvironment, error) {
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
			Type: pulumi.String("azure-native:web:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:ManagedHostingEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:ManagedHostingEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedHostingEnvironment
	err := ctx.RegisterResource("azure-native:web/v20150801:ManagedHostingEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedHostingEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedHostingEnvironmentState, opts ...pulumi.ResourceOption) (*ManagedHostingEnvironment, error) {
	var resource ManagedHostingEnvironment
	err := ctx.ReadResource("azure-native:web/v20150801:ManagedHostingEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedHostingEnvironmentState struct {
}

type ManagedHostingEnvironmentState struct {
}

func (ManagedHostingEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedHostingEnvironmentState)(nil)).Elem()
}

type managedHostingEnvironmentArgs struct {
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


type ManagedHostingEnvironmentArgs struct {
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

func (ManagedHostingEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedHostingEnvironmentArgs)(nil)).Elem()
}

type ManagedHostingEnvironmentInput interface {
	pulumi.Input

	ToManagedHostingEnvironmentOutput() ManagedHostingEnvironmentOutput
	ToManagedHostingEnvironmentOutputWithContext(ctx context.Context) ManagedHostingEnvironmentOutput
}

func (*ManagedHostingEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHostingEnvironment)(nil)).Elem()
}

func (i *ManagedHostingEnvironment) ToManagedHostingEnvironmentOutput() ManagedHostingEnvironmentOutput {
	return i.ToManagedHostingEnvironmentOutputWithContext(context.Background())
}

func (i *ManagedHostingEnvironment) ToManagedHostingEnvironmentOutputWithContext(ctx context.Context) ManagedHostingEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHostingEnvironmentOutput)
}

type ManagedHostingEnvironmentOutput struct{ *pulumi.OutputState }

func (ManagedHostingEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHostingEnvironment)(nil)).Elem()
}

func (o ManagedHostingEnvironmentOutput) ToManagedHostingEnvironmentOutput() ManagedHostingEnvironmentOutput {
	return o
}

func (o ManagedHostingEnvironmentOutput) ToManagedHostingEnvironmentOutputWithContext(ctx context.Context) ManagedHostingEnvironmentOutput {
	return o
}

func (o ManagedHostingEnvironmentOutput) AllowedMultiSizes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.AllowedMultiSizes }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) AllowedWorkerSizes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.AllowedWorkerSizes }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) ApiManagementAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.ApiManagementAccountId }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) ClusterSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) NameValuePairResponseArrayOutput { return v.ClusterSettings }).(NameValuePairResponseArrayOutput)
}

func (o ManagedHostingEnvironmentOutput) DatabaseEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.DatabaseEdition }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) DatabaseServiceObjective() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.DatabaseServiceObjective }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) EnvironmentCapacities() StampCapacityResponseArrayOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) StampCapacityResponseArrayOutput { return v.EnvironmentCapacities }).(StampCapacityResponseArrayOutput)
}

func (o ManagedHostingEnvironmentOutput) EnvironmentIsHealthy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.BoolPtrOutput { return v.EnvironmentIsHealthy }).(pulumi.BoolPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) EnvironmentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.EnvironmentStatus }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) InternalLoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.InternalLoadBalancingMode }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) IpsslAddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.IntPtrOutput { return v.IpsslAddressCount }).(pulumi.IntPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) LastAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.LastAction }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) LastActionResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.LastActionResult }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedHostingEnvironmentOutput) MaximumNumberOfMachines() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.IntPtrOutput { return v.MaximumNumberOfMachines }).(pulumi.IntPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) MultiRoleCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.IntPtrOutput { return v.MultiRoleCount }).(pulumi.IntPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) MultiSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.MultiSize }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) NetworkAccessControlList() NetworkAccessControlEntryResponseArrayOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) NetworkAccessControlEntryResponseArrayOutput {
		return v.NetworkAccessControlList
	}).(NetworkAccessControlEntryResponseArrayOutput)
}

func (o ManagedHostingEnvironmentOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ManagedHostingEnvironmentOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) Suspended() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.BoolPtrOutput { return v.Suspended }).(pulumi.BoolPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedHostingEnvironmentOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) UpgradeDomains() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.IntPtrOutput { return v.UpgradeDomains }).(pulumi.IntPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) VipMappings() VirtualIPMappingResponseArrayOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) VirtualIPMappingResponseArrayOutput { return v.VipMappings }).(VirtualIPMappingResponseArrayOutput)
}

func (o ManagedHostingEnvironmentOutput) VirtualNetwork() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) VirtualNetworkProfileResponsePtrOutput { return v.VirtualNetwork }).(VirtualNetworkProfileResponsePtrOutput)
}

func (o ManagedHostingEnvironmentOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) VnetResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.VnetResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) VnetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) pulumi.StringPtrOutput { return v.VnetSubnetName }).(pulumi.StringPtrOutput)
}

func (o ManagedHostingEnvironmentOutput) WorkerPools() WorkerPoolResponseArrayOutput {
	return o.ApplyT(func(v *ManagedHostingEnvironment) WorkerPoolResponseArrayOutput { return v.WorkerPools }).(WorkerPoolResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedHostingEnvironmentOutput{})
}
