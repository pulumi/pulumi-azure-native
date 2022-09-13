


package v20190601preview

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
			Type: pulumi.String("azure-native:containerregistry:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:containerregistry/v20190601preview:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:containerregistry/v20190601preview:AgentPool", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (i *AgentPool) ToAgentPoolOutput() AgentPoolOutput {
	return i.ToAgentPoolOutputWithContext(context.Background())
}

func (i *AgentPool) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolOutput)
}

type AgentPoolOutput struct{ *pulumi.OutputState }

func (AgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (o AgentPoolOutput) ToAgentPoolOutput() AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Os }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AgentPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AgentPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AgentPoolOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Tier }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) VirtualNetworkSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VirtualNetworkSubnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
