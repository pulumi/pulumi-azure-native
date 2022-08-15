


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedCluster struct {
	pulumi.CustomResourceState

	AddonFeatures                pulumi.StringArrayOutput                      `pulumi:"addonFeatures"`
	AdminPassword                pulumi.StringPtrOutput                        `pulumi:"adminPassword"`
	AdminUserName                pulumi.StringOutput                           `pulumi:"adminUserName"`
	AzureActiveDirectory         AzureActiveDirectoryResponsePtrOutput         `pulumi:"azureActiveDirectory"`
	ClientConnectionPort         pulumi.IntPtrOutput                           `pulumi:"clientConnectionPort"`
	Clients                      ClientCertificateResponseArrayOutput          `pulumi:"clients"`
	ClusterCertificateThumbprint pulumi.StringOutput                           `pulumi:"clusterCertificateThumbprint"`
	ClusterCodeVersion           pulumi.StringPtrOutput                        `pulumi:"clusterCodeVersion"`
	ClusterId                    pulumi.StringOutput                           `pulumi:"clusterId"`
	ClusterState                 pulumi.StringOutput                           `pulumi:"clusterState"`
	DnsName                      pulumi.StringOutput                           `pulumi:"dnsName"`
	Etag                         pulumi.StringOutput                           `pulumi:"etag"`
	FabricSettings               SettingsSectionDescriptionResponseArrayOutput `pulumi:"fabricSettings"`
	Fqdn                         pulumi.StringOutput                           `pulumi:"fqdn"`
	HttpGatewayConnectionPort    pulumi.IntPtrOutput                           `pulumi:"httpGatewayConnectionPort"`
	LoadBalancingRules           LoadBalancingRuleResponseArrayOutput          `pulumi:"loadBalancingRules"`
	Location                     pulumi.StringOutput                           `pulumi:"location"`
	Name                         pulumi.StringOutput                           `pulumi:"name"`
	ProvisioningState            pulumi.StringOutput                           `pulumi:"provisioningState"`
	Sku                          SkuResponsePtrOutput                          `pulumi:"sku"`
	Tags                         pulumi.StringMapOutput                        `pulumi:"tags"`
	Type                         pulumi.StringOutput                           `pulumi:"type"`
}


func NewManagedCluster(ctx *pulumi.Context,
	name string, args *ManagedClusterArgs, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminUserName == nil {
		return nil, errors.New("invalid value for required argument 'AdminUserName'")
	}
	if args.DnsName == nil {
		return nil, errors.New("invalid value for required argument 'DnsName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.ClientConnectionPort) {
		args.ClientConnectionPort = pulumi.IntPtr(19000)
	}
	if isZero(args.HttpGatewayConnectionPort) {
		args.HttpGatewayConnectionPort = pulumi.IntPtr(19080)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210701preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:ManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedCluster
	err := ctx.RegisterResource("azure-native:servicefabric/v20200101preview:ManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterState, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	var resource ManagedCluster
	err := ctx.ReadResource("azure-native:servicefabric/v20200101preview:ManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedClusterState struct {
}

type ManagedClusterState struct {
}

func (ManagedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterState)(nil)).Elem()
}

type managedClusterArgs struct {
	AddonFeatures             []string                     `pulumi:"addonFeatures"`
	AdminPassword             *string                      `pulumi:"adminPassword"`
	AdminUserName             string                       `pulumi:"adminUserName"`
	AzureActiveDirectory      *AzureActiveDirectory        `pulumi:"azureActiveDirectory"`
	ClientConnectionPort      *int                         `pulumi:"clientConnectionPort"`
	Clients                   []ClientCertificate          `pulumi:"clients"`
	ClusterCodeVersion        *string                      `pulumi:"clusterCodeVersion"`
	ClusterName               *string                      `pulumi:"clusterName"`
	DnsName                   string                       `pulumi:"dnsName"`
	FabricSettings            []SettingsSectionDescription `pulumi:"fabricSettings"`
	HttpGatewayConnectionPort *int                         `pulumi:"httpGatewayConnectionPort"`
	LoadBalancingRules        []LoadBalancingRule          `pulumi:"loadBalancingRules"`
	Location                  *string                      `pulumi:"location"`
	ResourceGroupName         string                       `pulumi:"resourceGroupName"`
	Sku                       *Sku                         `pulumi:"sku"`
	Tags                      map[string]string            `pulumi:"tags"`
}


type ManagedClusterArgs struct {
	AddonFeatures             pulumi.StringArrayInput
	AdminPassword             pulumi.StringPtrInput
	AdminUserName             pulumi.StringInput
	AzureActiveDirectory      AzureActiveDirectoryPtrInput
	ClientConnectionPort      pulumi.IntPtrInput
	Clients                   ClientCertificateArrayInput
	ClusterCodeVersion        pulumi.StringPtrInput
	ClusterName               pulumi.StringPtrInput
	DnsName                   pulumi.StringInput
	FabricSettings            SettingsSectionDescriptionArrayInput
	HttpGatewayConnectionPort pulumi.IntPtrInput
	LoadBalancingRules        LoadBalancingRuleArrayInput
	Location                  pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Sku                       SkuPtrInput
	Tags                      pulumi.StringMapInput
}

func (ManagedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterArgs)(nil)).Elem()
}

type ManagedClusterInput interface {
	pulumi.Input

	ToManagedClusterOutput() ManagedClusterOutput
	ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput
}

func (*ManagedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCluster)(nil)).Elem()
}

func (i *ManagedCluster) ToManagedClusterOutput() ManagedClusterOutput {
	return i.ToManagedClusterOutputWithContext(context.Background())
}

func (i *ManagedCluster) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterOutput)
}

type ManagedClusterOutput struct{ *pulumi.OutputState }

func (ManagedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCluster)(nil)).Elem()
}

func (o ManagedClusterOutput) ToManagedClusterOutput() ManagedClusterOutput {
	return o
}

func (o ManagedClusterOutput) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return o
}

func (o ManagedClusterOutput) AddonFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringArrayOutput { return v.AddonFeatures }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.AdminUserName }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) AzureActiveDirectoryResponsePtrOutput { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

func (o ManagedClusterOutput) ClientConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.IntPtrOutput { return v.ClientConnectionPort }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterOutput) Clients() ClientCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) ClientCertificateResponseArrayOutput { return v.Clients }).(ClientCertificateResponseArrayOutput)
}

func (o ManagedClusterOutput) ClusterCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ClusterCertificateThumbprint }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) ClusterState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ClusterState }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) SettingsSectionDescriptionResponseArrayOutput { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

func (o ManagedClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) HttpGatewayConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.IntPtrOutput { return v.HttpGatewayConnectionPort }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) LoadBalancingRuleResponseArrayOutput { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
}

func (o ManagedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ManagedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterOutput{})
}
