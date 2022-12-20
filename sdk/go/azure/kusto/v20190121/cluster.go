


package v20190121

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	DataIngestionUri       pulumi.StringOutput                      `pulumi:"dataIngestionUri"`
	Location               pulumi.StringOutput                      `pulumi:"location"`
	Name                   pulumi.StringOutput                      `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput                      `pulumi:"provisioningState"`
	Sku                    AzureSkuResponseOutput                   `pulumi:"sku"`
	State                  pulumi.StringOutput                      `pulumi:"state"`
	Tags                   pulumi.StringMapOutput                   `pulumi:"tags"`
	TrustedExternalTenants TrustedExternalTenantResponseArrayOutput `pulumi:"trustedExternalTenants"`
	Type                   pulumi.StringOutput                      `pulumi:"type"`
	Uri                    pulumi.StringOutput                      `pulumi:"uri"`
}


func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:kusto/v20190121:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:kusto/v20190121:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	ClusterName            *string                 `pulumi:"clusterName"`
	Location               *string                 `pulumi:"location"`
	ResourceGroupName      string                  `pulumi:"resourceGroupName"`
	Sku                    AzureSku                `pulumi:"sku"`
	Tags                   map[string]string       `pulumi:"tags"`
	TrustedExternalTenants []TrustedExternalTenant `pulumi:"trustedExternalTenants"`
}


type ClusterArgs struct {
	ClusterName            pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    AzureSkuInput
	Tags                   pulumi.StringMapInput
	TrustedExternalTenants TrustedExternalTenantArrayInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) DataIngestionUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DataIngestionUri }).(pulumi.StringOutput)
}

func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ClusterOutput) Sku() AzureSkuResponseOutput {
	return o.ApplyT(func(v *Cluster) AzureSkuResponseOutput { return v.Sku }).(AzureSkuResponseOutput)
}

func (o ClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ClusterOutput) TrustedExternalTenants() TrustedExternalTenantResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) TrustedExternalTenantResponseArrayOutput { return v.TrustedExternalTenants }).(TrustedExternalTenantResponseArrayOutput)
}

func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ClusterOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
