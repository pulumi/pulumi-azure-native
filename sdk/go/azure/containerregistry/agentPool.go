


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPool struct {
	pulumi.CustomResourceState

	Count                          pulumi.IntPtrOutput      `pulumi:"count"`
	Location                       pulumi.StringOutput      `pulumi:"location"`
	Name                           pulumi.StringOutput      `pulumi:"name"`
	Os                             pulumi.StringPtrOutput   `pulumi:"os"`
	ProvisioningState              pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData                     SystemDataResponseOutput `pulumi:"systemData"`
	Tags                           pulumi.StringMapOutput   `pulumi:"tags"`
	Tier                           pulumi.StringPtrOutput   `pulumi:"tier"`
	Type                           pulumi.StringOutput      `pulumi:"type"`
	VirtualNetworkSubnetResourceId pulumi.StringPtrOutput   `pulumi:"virtualNetworkSubnetResourceId"`
}


func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry/v20190601preview:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:containerregistry:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:containerregistry:AgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type agentPoolState struct {
}

type AgentPoolState struct {
}

func (AgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolState)(nil)).Elem()
}

type agentPoolArgs struct {
	AgentPoolName                  *string           `pulumi:"agentPoolName"`
	Count                          *int              `pulumi:"count"`
	Location                       *string           `pulumi:"location"`
	Os                             *string           `pulumi:"os"`
	RegistryName                   string            `pulumi:"registryName"`
	ResourceGroupName              string            `pulumi:"resourceGroupName"`
	Tags                           map[string]string `pulumi:"tags"`
	Tier                           *string           `pulumi:"tier"`
	VirtualNetworkSubnetResourceId *string           `pulumi:"virtualNetworkSubnetResourceId"`
}


type AgentPoolArgs struct {
	AgentPoolName                  pulumi.StringPtrInput
	Count                          pulumi.IntPtrInput
	Location                       pulumi.StringPtrInput
	Os                             pulumi.StringPtrInput
	RegistryName                   pulumi.StringInput
	ResourceGroupName              pulumi.StringInput
	Tags                           pulumi.StringMapInput
	Tier                           pulumi.StringPtrInput
	VirtualNetworkSubnetResourceId pulumi.StringPtrInput
}

func (AgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolArgs)(nil)).Elem()
}

type AgentPoolInput interface {
	pulumi.Input

	ToAgentPoolOutput() AgentPoolOutput
	ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput
}

func (*AgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPool)(nil))
}

func (i *AgentPool) ToAgentPoolOutput() AgentPoolOutput {
	return i.ToAgentPoolOutputWithContext(context.Background())
}

func (i *AgentPool) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolOutput)
}

type AgentPoolOutput struct{ *pulumi.OutputState }

func (AgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPool)(nil))
}

func (o AgentPoolOutput) ToAgentPoolOutput() AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
