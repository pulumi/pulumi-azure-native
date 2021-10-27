


package delegatednetwork

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OrchestratorInstanceServiceDetails struct {
	pulumi.CustomResourceState

	ApiServerEndpoint     pulumi.StringPtrOutput                `pulumi:"apiServerEndpoint"`
	ClusterRootCA         pulumi.StringPtrOutput                `pulumi:"clusterRootCA"`
	ControllerDetails     ControllerDetailsResponseOutput       `pulumi:"controllerDetails"`
	Identity              OrchestratorIdentityResponsePtrOutput `pulumi:"identity"`
	Kind                  pulumi.StringOutput                   `pulumi:"kind"`
	Location              pulumi.StringPtrOutput                `pulumi:"location"`
	Name                  pulumi.StringOutput                   `pulumi:"name"`
	OrchestratorAppId     pulumi.StringPtrOutput                `pulumi:"orchestratorAppId"`
	OrchestratorTenantId  pulumi.StringPtrOutput                `pulumi:"orchestratorTenantId"`
	PrivateLinkResourceId pulumi.StringPtrOutput                `pulumi:"privateLinkResourceId"`
	ProvisioningState     pulumi.StringOutput                   `pulumi:"provisioningState"`
	ResourceGuid          pulumi.StringOutput                   `pulumi:"resourceGuid"`
	Tags                  pulumi.StringMapOutput                `pulumi:"tags"`
	Type                  pulumi.StringOutput                   `pulumi:"type"`
}


func NewOrchestratorInstanceServiceDetails(ctx *pulumi.Context,
	name string, args *OrchestratorInstanceServiceDetailsArgs, opts ...pulumi.ResourceOption) (*OrchestratorInstanceServiceDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControllerDetails == nil {
		return nil, errors.New("invalid value for required argument 'ControllerDetails'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:delegatednetwork:OrchestratorInstanceServiceDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20200808preview:OrchestratorInstanceServiceDetails"),
		},
		{
			Type: pulumi.String("azure-nextgen:delegatednetwork/v20200808preview:OrchestratorInstanceServiceDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20210315:OrchestratorInstanceServiceDetails"),
		},
		{
			Type: pulumi.String("azure-nextgen:delegatednetwork/v20210315:OrchestratorInstanceServiceDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource OrchestratorInstanceServiceDetails
	err := ctx.RegisterResource("azure-native:delegatednetwork:OrchestratorInstanceServiceDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrchestratorInstanceServiceDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrchestratorInstanceServiceDetailsState, opts ...pulumi.ResourceOption) (*OrchestratorInstanceServiceDetails, error) {
	var resource OrchestratorInstanceServiceDetails
	err := ctx.ReadResource("azure-native:delegatednetwork:OrchestratorInstanceServiceDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type orchestratorInstanceServiceDetailsState struct {
}

type OrchestratorInstanceServiceDetailsState struct {
}

func (OrchestratorInstanceServiceDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratorInstanceServiceDetailsState)(nil)).Elem()
}

type orchestratorInstanceServiceDetailsArgs struct {
	ApiServerEndpoint     *string               `pulumi:"apiServerEndpoint"`
	ClusterRootCA         *string               `pulumi:"clusterRootCA"`
	ControllerDetails     ControllerDetailsType `pulumi:"controllerDetails"`
	Identity              *OrchestratorIdentity `pulumi:"identity"`
	Kind                  string                `pulumi:"kind"`
	Location              *string               `pulumi:"location"`
	OrchestratorAppId     *string               `pulumi:"orchestratorAppId"`
	OrchestratorTenantId  *string               `pulumi:"orchestratorTenantId"`
	PrivateLinkResourceId *string               `pulumi:"privateLinkResourceId"`
	ResourceGroupName     string                `pulumi:"resourceGroupName"`
	ResourceName          *string               `pulumi:"resourceName"`
	Tags                  map[string]string     `pulumi:"tags"`
}


type OrchestratorInstanceServiceDetailsArgs struct {
	ApiServerEndpoint     pulumi.StringPtrInput
	ClusterRootCA         pulumi.StringPtrInput
	ControllerDetails     ControllerDetailsTypeInput
	Identity              OrchestratorIdentityPtrInput
	Kind                  pulumi.StringInput
	Location              pulumi.StringPtrInput
	OrchestratorAppId     pulumi.StringPtrInput
	OrchestratorTenantId  pulumi.StringPtrInput
	PrivateLinkResourceId pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ResourceName          pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
}

func (OrchestratorInstanceServiceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratorInstanceServiceDetailsArgs)(nil)).Elem()
}

type OrchestratorInstanceServiceDetailsInput interface {
	pulumi.Input

	ToOrchestratorInstanceServiceDetailsOutput() OrchestratorInstanceServiceDetailsOutput
	ToOrchestratorInstanceServiceDetailsOutputWithContext(ctx context.Context) OrchestratorInstanceServiceDetailsOutput
}

func (*OrchestratorInstanceServiceDetails) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorInstanceServiceDetails)(nil))
}

func (i *OrchestratorInstanceServiceDetails) ToOrchestratorInstanceServiceDetailsOutput() OrchestratorInstanceServiceDetailsOutput {
	return i.ToOrchestratorInstanceServiceDetailsOutputWithContext(context.Background())
}

func (i *OrchestratorInstanceServiceDetails) ToOrchestratorInstanceServiceDetailsOutputWithContext(ctx context.Context) OrchestratorInstanceServiceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratorInstanceServiceDetailsOutput)
}

type OrchestratorInstanceServiceDetailsOutput struct{ *pulumi.OutputState }

func (OrchestratorInstanceServiceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorInstanceServiceDetails)(nil))
}

func (o OrchestratorInstanceServiceDetailsOutput) ToOrchestratorInstanceServiceDetailsOutput() OrchestratorInstanceServiceDetailsOutput {
	return o
}

func (o OrchestratorInstanceServiceDetailsOutput) ToOrchestratorInstanceServiceDetailsOutputWithContext(ctx context.Context) OrchestratorInstanceServiceDetailsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrchestratorInstanceServiceDetailsOutput{})
}
