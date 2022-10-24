


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectedCluster struct {
	pulumi.CustomResourceState

	AgentPublicKeyCertificate                pulumi.StringOutput                    `pulumi:"agentPublicKeyCertificate"`
	AgentVersion                             pulumi.StringOutput                    `pulumi:"agentVersion"`
	ConnectivityStatus                       pulumi.StringOutput                    `pulumi:"connectivityStatus"`
	Distribution                             pulumi.StringPtrOutput                 `pulumi:"distribution"`
	Identity                                 ConnectedClusterIdentityResponseOutput `pulumi:"identity"`
	Infrastructure                           pulumi.StringPtrOutput                 `pulumi:"infrastructure"`
	KubernetesVersion                        pulumi.StringOutput                    `pulumi:"kubernetesVersion"`
	LastConnectivityTime                     pulumi.StringOutput                    `pulumi:"lastConnectivityTime"`
	Location                                 pulumi.StringOutput                    `pulumi:"location"`
	ManagedIdentityCertificateExpirationTime pulumi.StringOutput                    `pulumi:"managedIdentityCertificateExpirationTime"`
	Name                                     pulumi.StringOutput                    `pulumi:"name"`
	Offering                                 pulumi.StringOutput                    `pulumi:"offering"`
	ProvisioningState                        pulumi.StringPtrOutput                 `pulumi:"provisioningState"`
	SystemData                               SystemDataResponseOutput               `pulumi:"systemData"`
	Tags                                     pulumi.StringMapOutput                 `pulumi:"tags"`
	TotalCoreCount                           pulumi.IntOutput                       `pulumi:"totalCoreCount"`
	TotalNodeCount                           pulumi.IntOutput                       `pulumi:"totalNodeCount"`
	Type                                     pulumi.StringOutput                    `pulumi:"type"`
}


func NewConnectedCluster(ctx *pulumi.Context,
	name string, args *ConnectedClusterArgs, opts ...pulumi.ResourceOption) (*ConnectedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentPublicKeyCertificate == nil {
		return nil, errors.New("invalid value for required argument 'AgentPublicKeyCertificate'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Identity = args.Identity.ToConnectedClusterIdentityOutput().ApplyT(func(v ConnectedClusterIdentity) ConnectedClusterIdentity { return *v.Defaults() }).(ConnectedClusterIdentityOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kubernetes:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20200101preview:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20210401preview:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20211001:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20220501preview:ConnectedCluster"),
		},
		{
			Type: pulumi.String("azure-native:kubernetes/v20221001preview:ConnectedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectedCluster
	err := ctx.RegisterResource("azure-native:kubernetes/v20210301:ConnectedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedClusterState, opts ...pulumi.ResourceOption) (*ConnectedCluster, error) {
	var resource ConnectedCluster
	err := ctx.ReadResource("azure-native:kubernetes/v20210301:ConnectedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectedClusterState struct {
}

type ConnectedClusterState struct {
}

func (ConnectedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedClusterState)(nil)).Elem()
}

type connectedClusterArgs struct {
	AgentPublicKeyCertificate string                   `pulumi:"agentPublicKeyCertificate"`
	ClusterName               *string                  `pulumi:"clusterName"`
	Distribution              *string                  `pulumi:"distribution"`
	Identity                  ConnectedClusterIdentity `pulumi:"identity"`
	Infrastructure            *string                  `pulumi:"infrastructure"`
	Location                  *string                  `pulumi:"location"`
	ProvisioningState         *string                  `pulumi:"provisioningState"`
	ResourceGroupName         string                   `pulumi:"resourceGroupName"`
	Tags                      map[string]string        `pulumi:"tags"`
}


type ConnectedClusterArgs struct {
	AgentPublicKeyCertificate pulumi.StringInput
	ClusterName               pulumi.StringPtrInput
	Distribution              pulumi.StringPtrInput
	Identity                  ConnectedClusterIdentityInput
	Infrastructure            pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	ProvisioningState         pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
}

func (ConnectedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedClusterArgs)(nil)).Elem()
}

type ConnectedClusterInput interface {
	pulumi.Input

	ToConnectedClusterOutput() ConnectedClusterOutput
	ToConnectedClusterOutputWithContext(ctx context.Context) ConnectedClusterOutput
}

func (*ConnectedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedCluster)(nil)).Elem()
}

func (i *ConnectedCluster) ToConnectedClusterOutput() ConnectedClusterOutput {
	return i.ToConnectedClusterOutputWithContext(context.Background())
}

func (i *ConnectedCluster) ToConnectedClusterOutputWithContext(ctx context.Context) ConnectedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterOutput)
}

type ConnectedClusterOutput struct{ *pulumi.OutputState }

func (ConnectedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedCluster)(nil)).Elem()
}

func (o ConnectedClusterOutput) ToConnectedClusterOutput() ConnectedClusterOutput {
	return o
}

func (o ConnectedClusterOutput) ToConnectedClusterOutputWithContext(ctx context.Context) ConnectedClusterOutput {
	return o
}

func (o ConnectedClusterOutput) AgentPublicKeyCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.AgentPublicKeyCertificate }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) ConnectivityStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.ConnectivityStatus }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) Distribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.Distribution }).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterOutput) Identity() ConnectedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ConnectedCluster) ConnectedClusterIdentityResponseOutput { return v.Identity }).(ConnectedClusterIdentityResponseOutput)
}

func (o ConnectedClusterOutput) Infrastructure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.Infrastructure }).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) LastConnectivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.LastConnectivityTime }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) ManagedIdentityCertificateExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.ManagedIdentityCertificateExpirationTime }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) Offering() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.Offering }).(pulumi.StringOutput)
}

func (o ConnectedClusterOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectedClusterOutput) TotalCoreCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.IntOutput { return v.TotalCoreCount }).(pulumi.IntOutput)
}

func (o ConnectedClusterOutput) TotalNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.IntOutput { return v.TotalNodeCount }).(pulumi.IntOutput)
}

func (o ConnectedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedClusterOutput{})
}
