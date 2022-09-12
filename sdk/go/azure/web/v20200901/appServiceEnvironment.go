


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServiceEnvironment struct {
	pulumi.CustomResourceState

	AllowedMultiSizes          pulumi.StringOutput                          `pulumi:"allowedMultiSizes"`
	AllowedWorkerSizes         pulumi.StringOutput                          `pulumi:"allowedWorkerSizes"`
	ApiManagementAccountId     pulumi.StringPtrOutput                       `pulumi:"apiManagementAccountId"`
	ClusterSettings            NameValuePairResponseArrayOutput             `pulumi:"clusterSettings"`
	DatabaseEdition            pulumi.StringOutput                          `pulumi:"databaseEdition"`
	DatabaseServiceObjective   pulumi.StringOutput                          `pulumi:"databaseServiceObjective"`
	DefaultFrontEndScaleFactor pulumi.IntOutput                             `pulumi:"defaultFrontEndScaleFactor"`
	DnsSuffix                  pulumi.StringPtrOutput                       `pulumi:"dnsSuffix"`
	DynamicCacheEnabled        pulumi.BoolPtrOutput                         `pulumi:"dynamicCacheEnabled"`
	EnvironmentCapacities      StampCapacityResponseArrayOutput             `pulumi:"environmentCapacities"`
	EnvironmentIsHealthy       pulumi.BoolOutput                            `pulumi:"environmentIsHealthy"`
	EnvironmentStatus          pulumi.StringOutput                          `pulumi:"environmentStatus"`
	FrontEndScaleFactor        pulumi.IntPtrOutput                          `pulumi:"frontEndScaleFactor"`
	HasLinuxWorkers            pulumi.BoolPtrOutput                         `pulumi:"hasLinuxWorkers"`
	InternalLoadBalancingMode  pulumi.StringPtrOutput                       `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount          pulumi.IntPtrOutput                          `pulumi:"ipsslAddressCount"`
	Kind                       pulumi.StringPtrOutput                       `pulumi:"kind"`
	LastAction                 pulumi.StringOutput                          `pulumi:"lastAction"`
	LastActionResult           pulumi.StringOutput                          `pulumi:"lastActionResult"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	MaximumNumberOfMachines    pulumi.IntOutput                             `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount             pulumi.IntPtrOutput                          `pulumi:"multiRoleCount"`
	MultiSize                  pulumi.StringPtrOutput                       `pulumi:"multiSize"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	NetworkAccessControlList   NetworkAccessControlEntryResponseArrayOutput `pulumi:"networkAccessControlList"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	ResourceGroup              pulumi.StringOutput                          `pulumi:"resourceGroup"`
	SslCertKeyVaultId          pulumi.StringPtrOutput                       `pulumi:"sslCertKeyVaultId"`
	SslCertKeyVaultSecretName  pulumi.StringPtrOutput                       `pulumi:"sslCertKeyVaultSecretName"`
	Status                     pulumi.StringOutput                          `pulumi:"status"`
	SubscriptionId             pulumi.StringOutput                          `pulumi:"subscriptionId"`
	Suspended                  pulumi.BoolPtrOutput                         `pulumi:"suspended"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
	UpgradeDomains             pulumi.IntOutput                             `pulumi:"upgradeDomains"`
	UserWhitelistedIpRanges    pulumi.StringArrayOutput                     `pulumi:"userWhitelistedIpRanges"`
	VipMappings                VirtualIPMappingResponseArrayOutput          `pulumi:"vipMappings"`
	VirtualNetwork             VirtualNetworkProfileResponseOutput          `pulumi:"virtualNetwork"`
	VnetName                   pulumi.StringPtrOutput                       `pulumi:"vnetName"`
	VnetResourceGroupName      pulumi.StringPtrOutput                       `pulumi:"vnetResourceGroupName"`
	VnetSubnetName             pulumi.StringPtrOutput                       `pulumi:"vnetSubnetName"`
	WorkerPools                WorkerPoolResponseArrayOutput                `pulumi:"workerPools"`
}


func NewAppServiceEnvironment(ctx *pulumi.Context,
	name string, args *AppServiceEnvironmentArgs, opts ...pulumi.ResourceOption) (*AppServiceEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetwork == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetwork'")
	}
	if args.WorkerPools == nil {
		return nil, errors.New("invalid value for required argument 'WorkerPools'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:AppServiceEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceEnvironment
	err := ctx.RegisterResource("azure-native:web/v20200901:AppServiceEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServiceEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceEnvironmentState, opts ...pulumi.ResourceOption) (*AppServiceEnvironment, error) {
	var resource AppServiceEnvironment
	err := ctx.ReadResource("azure-native:web/v20200901:AppServiceEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServiceEnvironmentState struct {
}

type AppServiceEnvironmentState struct {
}

func (AppServiceEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentState)(nil)).Elem()
}

type appServiceEnvironmentArgs struct {
	ApiManagementAccountId    *string                     `pulumi:"apiManagementAccountId"`
	ClusterSettings           []NameValuePair             `pulumi:"clusterSettings"`
	DnsSuffix                 *string                     `pulumi:"dnsSuffix"`
	DynamicCacheEnabled       *bool                       `pulumi:"dynamicCacheEnabled"`
	FrontEndScaleFactor       *int                        `pulumi:"frontEndScaleFactor"`
	HasLinuxWorkers           *bool                       `pulumi:"hasLinuxWorkers"`
	InternalLoadBalancingMode *string                     `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount         *int                        `pulumi:"ipsslAddressCount"`
	Kind                      *string                     `pulumi:"kind"`
	Location                  *string                     `pulumi:"location"`
	MultiRoleCount            *int                        `pulumi:"multiRoleCount"`
	MultiSize                 *string                     `pulumi:"multiSize"`
	Name                      *string                     `pulumi:"name"`
	NetworkAccessControlList  []NetworkAccessControlEntry `pulumi:"networkAccessControlList"`
	ResourceGroupName         string                      `pulumi:"resourceGroupName"`
	SslCertKeyVaultId         *string                     `pulumi:"sslCertKeyVaultId"`
	SslCertKeyVaultSecretName *string                     `pulumi:"sslCertKeyVaultSecretName"`
	Suspended                 *bool                       `pulumi:"suspended"`
	Tags                      map[string]string           `pulumi:"tags"`
	UserWhitelistedIpRanges   []string                    `pulumi:"userWhitelistedIpRanges"`
	VirtualNetwork            VirtualNetworkProfile       `pulumi:"virtualNetwork"`
	VnetName                  *string                     `pulumi:"vnetName"`
	VnetResourceGroupName     *string                     `pulumi:"vnetResourceGroupName"`
	VnetSubnetName            *string                     `pulumi:"vnetSubnetName"`
	WorkerPools               []WorkerPool                `pulumi:"workerPools"`
}


type AppServiceEnvironmentArgs struct {
	ApiManagementAccountId    pulumi.StringPtrInput
	ClusterSettings           NameValuePairArrayInput
	DnsSuffix                 pulumi.StringPtrInput
	DynamicCacheEnabled       pulumi.BoolPtrInput
	FrontEndScaleFactor       pulumi.IntPtrInput
	HasLinuxWorkers           pulumi.BoolPtrInput
	InternalLoadBalancingMode pulumi.StringPtrInput
	IpsslAddressCount         pulumi.IntPtrInput
	Kind                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	MultiRoleCount            pulumi.IntPtrInput
	MultiSize                 pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	NetworkAccessControlList  NetworkAccessControlEntryArrayInput
	ResourceGroupName         pulumi.StringInput
	SslCertKeyVaultId         pulumi.StringPtrInput
	SslCertKeyVaultSecretName pulumi.StringPtrInput
	Suspended                 pulumi.BoolPtrInput
	Tags                      pulumi.StringMapInput
	UserWhitelistedIpRanges   pulumi.StringArrayInput
	VirtualNetwork            VirtualNetworkProfileInput
	VnetName                  pulumi.StringPtrInput
	VnetResourceGroupName     pulumi.StringPtrInput
	VnetSubnetName            pulumi.StringPtrInput
	WorkerPools               WorkerPoolArrayInput
}

func (AppServiceEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentArgs)(nil)).Elem()
}

type AppServiceEnvironmentInput interface {
	pulumi.Input

	ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput
	ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput
}

func (*AppServiceEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironment)(nil)).Elem()
}

func (i *AppServiceEnvironment) ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput {
	return i.ToAppServiceEnvironmentOutputWithContext(context.Background())
}

func (i *AppServiceEnvironment) ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceEnvironmentOutput)
}

type AppServiceEnvironmentOutput struct{ *pulumi.OutputState }

func (AppServiceEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironment)(nil)).Elem()
}

func (o AppServiceEnvironmentOutput) ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput {
	return o
}

func (o AppServiceEnvironmentOutput) ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput {
	return o
}

func (o AppServiceEnvironmentOutput) AllowedMultiSizes() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.AllowedMultiSizes }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) AllowedWorkerSizes() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.AllowedWorkerSizes }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) ApiManagementAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.ApiManagementAccountId }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) ClusterSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) NameValuePairResponseArrayOutput { return v.ClusterSettings }).(NameValuePairResponseArrayOutput)
}

func (o AppServiceEnvironmentOutput) DatabaseEdition() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.DatabaseEdition }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) DatabaseServiceObjective() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.DatabaseServiceObjective }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) DefaultFrontEndScaleFactor() pulumi.IntOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntOutput { return v.DefaultFrontEndScaleFactor }).(pulumi.IntOutput)
}

func (o AppServiceEnvironmentOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) DynamicCacheEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.BoolPtrOutput { return v.DynamicCacheEnabled }).(pulumi.BoolPtrOutput)
}

func (o AppServiceEnvironmentOutput) EnvironmentCapacities() StampCapacityResponseArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) StampCapacityResponseArrayOutput { return v.EnvironmentCapacities }).(StampCapacityResponseArrayOutput)
}

func (o AppServiceEnvironmentOutput) EnvironmentIsHealthy() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.BoolOutput { return v.EnvironmentIsHealthy }).(pulumi.BoolOutput)
}

func (o AppServiceEnvironmentOutput) EnvironmentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.EnvironmentStatus }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) FrontEndScaleFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntPtrOutput { return v.FrontEndScaleFactor }).(pulumi.IntPtrOutput)
}

func (o AppServiceEnvironmentOutput) HasLinuxWorkers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.BoolPtrOutput { return v.HasLinuxWorkers }).(pulumi.BoolPtrOutput)
}

func (o AppServiceEnvironmentOutput) InternalLoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.InternalLoadBalancingMode }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) IpsslAddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntPtrOutput { return v.IpsslAddressCount }).(pulumi.IntPtrOutput)
}

func (o AppServiceEnvironmentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) LastAction() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.LastAction }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) LastActionResult() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.LastActionResult }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) MaximumNumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntOutput { return v.MaximumNumberOfMachines }).(pulumi.IntOutput)
}

func (o AppServiceEnvironmentOutput) MultiRoleCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntPtrOutput { return v.MultiRoleCount }).(pulumi.IntPtrOutput)
}

func (o AppServiceEnvironmentOutput) MultiSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.MultiSize }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) NetworkAccessControlList() NetworkAccessControlEntryResponseArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) NetworkAccessControlEntryResponseArrayOutput {
		return v.NetworkAccessControlList
	}).(NetworkAccessControlEntryResponseArrayOutput)
}

func (o AppServiceEnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) SslCertKeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.SslCertKeyVaultId }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) SslCertKeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.SslCertKeyVaultSecretName }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) Suspended() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.BoolPtrOutput { return v.Suspended }).(pulumi.BoolPtrOutput)
}

func (o AppServiceEnvironmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AppServiceEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AppServiceEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentOutput) UpgradeDomains() pulumi.IntOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.IntOutput { return v.UpgradeDomains }).(pulumi.IntOutput)
}

func (o AppServiceEnvironmentOutput) UserWhitelistedIpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringArrayOutput { return v.UserWhitelistedIpRanges }).(pulumi.StringArrayOutput)
}

func (o AppServiceEnvironmentOutput) VipMappings() VirtualIPMappingResponseArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) VirtualIPMappingResponseArrayOutput { return v.VipMappings }).(VirtualIPMappingResponseArrayOutput)
}

func (o AppServiceEnvironmentOutput) VirtualNetwork() VirtualNetworkProfileResponseOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) VirtualNetworkProfileResponseOutput { return v.VirtualNetwork }).(VirtualNetworkProfileResponseOutput)
}

func (o AppServiceEnvironmentOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) VnetResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.VnetResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) VnetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) pulumi.StringPtrOutput { return v.VnetSubnetName }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentOutput) WorkerPools() WorkerPoolResponseArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironment) WorkerPoolResponseArrayOutput { return v.WorkerPools }).(WorkerPoolResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServiceEnvironmentOutput{})
}
