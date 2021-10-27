


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	AssociatedWorkspaces          AssociatedWorkspaceResponseArrayOutput         `pulumi:"associatedWorkspaces"`
	BillingType                   pulumi.StringPtrOutput                         `pulumi:"billingType"`
	CapacityReservationProperties CapacityReservationPropertiesResponsePtrOutput `pulumi:"capacityReservationProperties"`
	ClusterId                     pulumi.StringOutput                            `pulumi:"clusterId"`
	CreatedDate                   pulumi.StringOutput                            `pulumi:"createdDate"`
	Identity                      IdentityResponsePtrOutput                      `pulumi:"identity"`
	IsAvailabilityZonesEnabled    pulumi.BoolPtrOutput                           `pulumi:"isAvailabilityZonesEnabled"`
	KeyVaultProperties            KeyVaultPropertiesResponsePtrOutput            `pulumi:"keyVaultProperties"`
	LastModifiedDate              pulumi.StringOutput                            `pulumi:"lastModifiedDate"`
	Location                      pulumi.StringOutput                            `pulumi:"location"`
	Name                          pulumi.StringOutput                            `pulumi:"name"`
	ProvisioningState             pulumi.StringOutput                            `pulumi:"provisioningState"`
	Sku                           ClusterSkuResponsePtrOutput                    `pulumi:"sku"`
	Tags                          pulumi.StringMapOutput                         `pulumi:"tags"`
	Type                          pulumi.StringOutput                            `pulumi:"type"`
}


func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:operationalinsights:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20190801preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200801:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20201001:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20201001:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20210601:Cluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20210601:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:operationalinsights:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:operationalinsights:Cluster", name, id, state, &resource, opts...)
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
	BillingType                *string             `pulumi:"billingType"`
	ClusterName                *string             `pulumi:"clusterName"`
	Identity                   *Identity           `pulumi:"identity"`
	IsAvailabilityZonesEnabled *bool               `pulumi:"isAvailabilityZonesEnabled"`
	IsDoubleEncryptionEnabled  *bool               `pulumi:"isDoubleEncryptionEnabled"`
	KeyVaultProperties         *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Location                   *string             `pulumi:"location"`
	ResourceGroupName          string              `pulumi:"resourceGroupName"`
	Sku                        *ClusterSku         `pulumi:"sku"`
	Tags                       map[string]string   `pulumi:"tags"`
}


type ClusterArgs struct {
	BillingType                pulumi.StringPtrInput
	ClusterName                pulumi.StringPtrInput
	Identity                   IdentityPtrInput
	IsAvailabilityZonesEnabled pulumi.BoolPtrInput
	IsDoubleEncryptionEnabled  pulumi.BoolPtrInput
	KeyVaultProperties         KeyVaultPropertiesPtrInput
	Location                   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	Sku                        ClusterSkuPtrInput
	Tags                       pulumi.StringMapInput
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
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterOutputType(ClusterOutput{})
}
