


package v20210901preview

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
	SystemData              SystemDataResponseOutput                 `pulumi:"systemData"`
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
			Type: pulumi.String("azure-native:redhatopenshift:OpenShiftCluster"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20200430:OpenShiftCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource OpenShiftCluster
	err := ctx.RegisterResource("azure-native:redhatopenshift/v20210901preview:OpenShiftCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOpenShiftCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenShiftClusterState, opts ...pulumi.ResourceOption) (*OpenShiftCluster, error) {
	var resource OpenShiftCluster
	err := ctx.ReadResource("azure-native:redhatopenshift/v20210901preview:OpenShiftCluster", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*OpenShiftCluster)(nil))
}

func (i *OpenShiftCluster) ToOpenShiftClusterOutput() OpenShiftClusterOutput {
	return i.ToOpenShiftClusterOutputWithContext(context.Background())
}

func (i *OpenShiftCluster) ToOpenShiftClusterOutputWithContext(ctx context.Context) OpenShiftClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftClusterOutput)
}

type OpenShiftClusterOutput struct{ *pulumi.OutputState }

func (OpenShiftClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftCluster)(nil))
}

func (o OpenShiftClusterOutput) ToOpenShiftClusterOutput() OpenShiftClusterOutput {
	return o
}

func (o OpenShiftClusterOutput) ToOpenShiftClusterOutputWithContext(ctx context.Context) OpenShiftClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OpenShiftClusterOutput{})
}
