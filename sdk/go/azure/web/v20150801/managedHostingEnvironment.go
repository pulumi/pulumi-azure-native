


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
	return reflect.TypeOf((*ManagedHostingEnvironment)(nil))
}

func (i *ManagedHostingEnvironment) ToManagedHostingEnvironmentOutput() ManagedHostingEnvironmentOutput {
	return i.ToManagedHostingEnvironmentOutputWithContext(context.Background())
}

func (i *ManagedHostingEnvironment) ToManagedHostingEnvironmentOutputWithContext(ctx context.Context) ManagedHostingEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHostingEnvironmentOutput)
}

type ManagedHostingEnvironmentOutput struct{ *pulumi.OutputState }

func (ManagedHostingEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHostingEnvironment)(nil))
}

func (o ManagedHostingEnvironmentOutput) ToManagedHostingEnvironmentOutput() ManagedHostingEnvironmentOutput {
	return o
}

func (o ManagedHostingEnvironmentOutput) ToManagedHostingEnvironmentOutputWithContext(ctx context.Context) ManagedHostingEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedHostingEnvironmentOutput{})
}
