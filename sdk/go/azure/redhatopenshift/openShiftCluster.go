


package redhatopenshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OpenShiftCluster struct {
	pulumi.CustomResourceState

	ApiserverProfile        APIServerProfileResponsePtrOutput        `pulumi:"apiserverProfile"`
	ClusterProfile          ClusterProfileResponsePtrOutput          `pulumi:"clusterProfile"`
	ConsoleProfile          ConsoleProfileResponsePtrOutput          `pulumi:"consoleProfile"`
	IngressProfiles         IngressProfileResponseArrayOutput        `pulumi:"ingressProfiles"`
	Location                pulumi.StringOutput                      `pulumi:"location"`
	MasterProfile           MasterProfileResponsePtrOutput           `pulumi:"masterProfile"`
	Name                    pulumi.StringOutput                      `pulumi:"name"`
	NetworkProfile          NetworkProfileResponsePtrOutput          `pulumi:"networkProfile"`
	ProvisioningState       pulumi.StringPtrOutput                   `pulumi:"provisioningState"`
	ServicePrincipalProfile ServicePrincipalProfileResponsePtrOutput `pulumi:"servicePrincipalProfile"`
	Tags                    pulumi.StringMapOutput                   `pulumi:"tags"`
	Type                    pulumi.StringOutput                      `pulumi:"type"`
	WorkerProfiles          WorkerProfileResponseArrayOutput         `pulumi:"workerProfiles"`
}


func NewOpenShiftCluster(ctx *pulumi.Context,
	name string, args *OpenShiftClusterArgs, opts ...pulumi.ResourceOption) (*OpenShiftCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20200430:OpenShiftCluster"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20210901preview:OpenShiftCluster"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20220401:OpenShiftCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource OpenShiftCluster
	err := ctx.RegisterResource("azure-native:redhatopenshift:OpenShiftCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOpenShiftCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenShiftClusterState, opts ...pulumi.ResourceOption) (*OpenShiftCluster, error) {
	var resource OpenShiftCluster
	err := ctx.ReadResource("azure-native:redhatopenshift:OpenShiftCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type openShiftClusterState struct {
}

type OpenShiftClusterState struct {
}

func (OpenShiftClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftClusterState)(nil)).Elem()
}

type openShiftClusterArgs struct {
	ApiserverProfile        *APIServerProfile        `pulumi:"apiserverProfile"`
	ClusterProfile          *ClusterProfile          `pulumi:"clusterProfile"`
	ConsoleProfile          *ConsoleProfile          `pulumi:"consoleProfile"`
	IngressProfiles         []IngressProfile         `pulumi:"ingressProfiles"`
	Location                *string                  `pulumi:"location"`
	MasterProfile           *MasterProfile           `pulumi:"masterProfile"`
	NetworkProfile          *NetworkProfile          `pulumi:"networkProfile"`
	ProvisioningState       *string                  `pulumi:"provisioningState"`
	ResourceGroupName       string                   `pulumi:"resourceGroupName"`
	ResourceName            *string                  `pulumi:"resourceName"`
	ServicePrincipalProfile *ServicePrincipalProfile `pulumi:"servicePrincipalProfile"`
	Tags                    map[string]string        `pulumi:"tags"`
	WorkerProfiles          []WorkerProfile          `pulumi:"workerProfiles"`
}


type OpenShiftClusterArgs struct {
	ApiserverProfile        APIServerProfilePtrInput
	ClusterProfile          ClusterProfilePtrInput
	ConsoleProfile          ConsoleProfilePtrInput
	IngressProfiles         IngressProfileArrayInput
	Location                pulumi.StringPtrInput
	MasterProfile           MasterProfilePtrInput
	NetworkProfile          NetworkProfilePtrInput
	ProvisioningState       pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ResourceName            pulumi.StringPtrInput
	ServicePrincipalProfile ServicePrincipalProfilePtrInput
	Tags                    pulumi.StringMapInput
	WorkerProfiles          WorkerProfileArrayInput
}

func (OpenShiftClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftClusterArgs)(nil)).Elem()
}

type OpenShiftClusterInput interface {
	pulumi.Input

	ToOpenShiftClusterOutput() OpenShiftClusterOutput
	ToOpenShiftClusterOutputWithContext(ctx context.Context) OpenShiftClusterOutput
}

func (*OpenShiftCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftCluster)(nil)).Elem()
}

func (i *OpenShiftCluster) ToOpenShiftClusterOutput() OpenShiftClusterOutput {
	return i.ToOpenShiftClusterOutputWithContext(context.Background())
}

func (i *OpenShiftCluster) ToOpenShiftClusterOutputWithContext(ctx context.Context) OpenShiftClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftClusterOutput)
}

type OpenShiftClusterOutput struct{ *pulumi.OutputState }

func (OpenShiftClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftCluster)(nil)).Elem()
}

func (o OpenShiftClusterOutput) ToOpenShiftClusterOutput() OpenShiftClusterOutput {
	return o
}

func (o OpenShiftClusterOutput) ToOpenShiftClusterOutputWithContext(ctx context.Context) OpenShiftClusterOutput {
	return o
}

func (o OpenShiftClusterOutput) ApiserverProfile() APIServerProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) APIServerProfileResponsePtrOutput { return v.ApiserverProfile }).(APIServerProfileResponsePtrOutput)
}

func (o OpenShiftClusterOutput) ClusterProfile() ClusterProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) ClusterProfileResponsePtrOutput { return v.ClusterProfile }).(ClusterProfileResponsePtrOutput)
}

func (o OpenShiftClusterOutput) ConsoleProfile() ConsoleProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) ConsoleProfileResponsePtrOutput { return v.ConsoleProfile }).(ConsoleProfileResponsePtrOutput)
}

func (o OpenShiftClusterOutput) IngressProfiles() IngressProfileResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftCluster) IngressProfileResponseArrayOutput { return v.IngressProfiles }).(IngressProfileResponseArrayOutput)
}

func (o OpenShiftClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o OpenShiftClusterOutput) MasterProfile() MasterProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) MasterProfileResponsePtrOutput { return v.MasterProfile }).(MasterProfileResponsePtrOutput)
}

func (o OpenShiftClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OpenShiftClusterOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o OpenShiftClusterOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o OpenShiftClusterOutput) ServicePrincipalProfile() ServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftCluster) ServicePrincipalProfileResponsePtrOutput { return v.ServicePrincipalProfile }).(ServicePrincipalProfileResponsePtrOutput)
}

func (o OpenShiftClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OpenShiftClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o OpenShiftClusterOutput) WorkerProfiles() WorkerProfileResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftCluster) WorkerProfileResponseArrayOutput { return v.WorkerProfiles }).(WorkerProfileResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(OpenShiftClusterOutput{})
}
