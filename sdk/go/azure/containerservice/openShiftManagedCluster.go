


package containerservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OpenShiftManagedCluster struct {
	pulumi.CustomResourceState

	AgentPoolProfiles OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput `pulumi:"agentPoolProfiles"`
	AuthProfile       OpenShiftManagedClusterAuthProfileResponsePtrOutput        `pulumi:"authProfile"`
	ClusterVersion    pulumi.StringOutput                                        `pulumi:"clusterVersion"`
	Fqdn              pulumi.StringOutput                                        `pulumi:"fqdn"`
	Location          pulumi.StringOutput                                        `pulumi:"location"`
	MasterPoolProfile OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput  `pulumi:"masterPoolProfile"`
	Name              pulumi.StringOutput                                        `pulumi:"name"`
	NetworkProfile    NetworkProfileResponsePtrOutput                            `pulumi:"networkProfile"`
	OpenShiftVersion  pulumi.StringOutput                                        `pulumi:"openShiftVersion"`
	Plan              PurchasePlanResponsePtrOutput                              `pulumi:"plan"`
	ProvisioningState pulumi.StringOutput                                        `pulumi:"provisioningState"`
	PublicHostname    pulumi.StringOutput                                        `pulumi:"publicHostname"`
	RouterProfiles    OpenShiftRouterProfileResponseArrayOutput                  `pulumi:"routerProfiles"`
	Tags              pulumi.StringMapOutput                                     `pulumi:"tags"`
	Type              pulumi.StringOutput                                        `pulumi:"type"`
}


func NewOpenShiftManagedCluster(ctx *pulumi.Context,
	name string, args *OpenShiftManagedClusterArgs, opts ...pulumi.ResourceOption) (*OpenShiftManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OpenShiftVersion == nil {
		return nil, errors.New("invalid value for required argument 'OpenShiftVersion'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.NetworkProfile != nil {
		args.NetworkProfile = args.NetworkProfile.ToNetworkProfilePtrOutput().ApplyT(func(v *NetworkProfile) *NetworkProfile { return v.Defaults() }).(NetworkProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice/v20180930preview:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190430:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190930preview:OpenShiftManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191027preview:OpenShiftManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource OpenShiftManagedCluster
	err := ctx.RegisterResource("azure-native:containerservice:OpenShiftManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOpenShiftManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenShiftManagedClusterState, opts ...pulumi.ResourceOption) (*OpenShiftManagedCluster, error) {
	var resource OpenShiftManagedCluster
	err := ctx.ReadResource("azure-native:containerservice:OpenShiftManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type openShiftManagedClusterState struct {
}

type OpenShiftManagedClusterState struct {
}

func (OpenShiftManagedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftManagedClusterState)(nil)).Elem()
}

type openShiftManagedClusterArgs struct {
	AgentPoolProfiles []OpenShiftManagedClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	AuthProfile       *OpenShiftManagedClusterAuthProfile       `pulumi:"authProfile"`
	Location          *string                                   `pulumi:"location"`
	MasterPoolProfile *OpenShiftManagedClusterMasterPoolProfile `pulumi:"masterPoolProfile"`
	NetworkProfile    *NetworkProfile                           `pulumi:"networkProfile"`
	OpenShiftVersion  string                                    `pulumi:"openShiftVersion"`
	Plan              *PurchasePlan                             `pulumi:"plan"`
	ResourceGroupName string                                    `pulumi:"resourceGroupName"`
	ResourceName      *string                                   `pulumi:"resourceName"`
	RouterProfiles    []OpenShiftRouterProfile                  `pulumi:"routerProfiles"`
	Tags              map[string]string                         `pulumi:"tags"`
}


type OpenShiftManagedClusterArgs struct {
	AgentPoolProfiles OpenShiftManagedClusterAgentPoolProfileArrayInput
	AuthProfile       OpenShiftManagedClusterAuthProfilePtrInput
	Location          pulumi.StringPtrInput
	MasterPoolProfile OpenShiftManagedClusterMasterPoolProfilePtrInput
	NetworkProfile    NetworkProfilePtrInput
	OpenShiftVersion  pulumi.StringInput
	Plan              PurchasePlanPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	RouterProfiles    OpenShiftRouterProfileArrayInput
	Tags              pulumi.StringMapInput
}

func (OpenShiftManagedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openShiftManagedClusterArgs)(nil)).Elem()
}

type OpenShiftManagedClusterInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput
	ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput
}

func (*OpenShiftManagedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedCluster)(nil)).Elem()
}

func (i *OpenShiftManagedCluster) ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput {
	return i.ToOpenShiftManagedClusterOutputWithContext(context.Background())
}

func (i *OpenShiftManagedCluster) ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterOutput)
}

type OpenShiftManagedClusterOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedCluster)(nil)).Elem()
}

func (o OpenShiftManagedClusterOutput) ToOpenShiftManagedClusterOutput() OpenShiftManagedClusterOutput {
	return o
}

func (o OpenShiftManagedClusterOutput) ToOpenShiftManagedClusterOutputWithContext(ctx context.Context) OpenShiftManagedClusterOutput {
	return o
}

func (o OpenShiftManagedClusterOutput) AgentPoolProfiles() OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput {
		return v.AgentPoolProfiles
	}).(OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput)
}

func (o OpenShiftManagedClusterOutput) AuthProfile() OpenShiftManagedClusterAuthProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftManagedClusterAuthProfileResponsePtrOutput {
		return v.AuthProfile
	}).(OpenShiftManagedClusterAuthProfileResponsePtrOutput)
}

func (o OpenShiftManagedClusterOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.ClusterVersion }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterOutput) MasterPoolProfile() OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput {
		return v.MasterPoolProfile
	}).(OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput)
}

func (o OpenShiftManagedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o OpenShiftManagedClusterOutput) OpenShiftVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.OpenShiftVersion }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterOutput) Plan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) PurchasePlanResponsePtrOutput { return v.Plan }).(PurchasePlanResponsePtrOutput)
}

func (o OpenShiftManagedClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterOutput) PublicHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.PublicHostname }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterOutput) RouterProfiles() OpenShiftRouterProfileResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) OpenShiftRouterProfileResponseArrayOutput { return v.RouterProfiles }).(OpenShiftRouterProfileResponseArrayOutput)
}

func (o OpenShiftManagedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OpenShiftManagedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenShiftManagedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OpenShiftManagedClusterOutput{})
}
