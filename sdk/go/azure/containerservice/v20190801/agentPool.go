


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPool struct {
	pulumi.CustomResourceState

	AvailabilityZones      pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	Count                  pulumi.IntPtrOutput      `pulumi:"count"`
	EnableAutoScaling      pulumi.BoolPtrOutput     `pulumi:"enableAutoScaling"`
	EnableNodePublicIP     pulumi.BoolPtrOutput     `pulumi:"enableNodePublicIP"`
	MaxCount               pulumi.IntPtrOutput      `pulumi:"maxCount"`
	MaxPods                pulumi.IntPtrOutput      `pulumi:"maxPods"`
	MinCount               pulumi.IntPtrOutput      `pulumi:"minCount"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	NodeTaints             pulumi.StringArrayOutput `pulumi:"nodeTaints"`
	OrchestratorVersion    pulumi.StringPtrOutput   `pulumi:"orchestratorVersion"`
	OsDiskSizeGB           pulumi.IntPtrOutput      `pulumi:"osDiskSizeGB"`
	OsType                 pulumi.StringPtrOutput   `pulumi:"osType"`
	ProvisioningState      pulumi.StringOutput      `pulumi:"provisioningState"`
	ScaleSetEvictionPolicy pulumi.StringPtrOutput   `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority       pulumi.StringPtrOutput   `pulumi:"scaleSetPriority"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
	VmSize                 pulumi.StringPtrOutput   `pulumi:"vmSize"`
	VnetSubnetID           pulumi.StringPtrOutput   `pulumi:"vnetSubnetID"`
}


func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:containerservice/v20190801:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:containerservice/v20190801:AgentPool", name, id, state, &resource, opts...)
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
	AgentPoolName          *string  `pulumi:"agentPoolName"`
	AvailabilityZones      []string `pulumi:"availabilityZones"`
	Count                  *int     `pulumi:"count"`
	EnableAutoScaling      *bool    `pulumi:"enableAutoScaling"`
	EnableNodePublicIP     *bool    `pulumi:"enableNodePublicIP"`
	MaxCount               *int     `pulumi:"maxCount"`
	MaxPods                *int     `pulumi:"maxPods"`
	MinCount               *int     `pulumi:"minCount"`
	NodeTaints             []string `pulumi:"nodeTaints"`
	OrchestratorVersion    *string  `pulumi:"orchestratorVersion"`
	OsDiskSizeGB           *int     `pulumi:"osDiskSizeGB"`
	OsType                 *string  `pulumi:"osType"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
	ResourceName           string   `pulumi:"resourceName"`
	ScaleSetEvictionPolicy *string  `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority       *string  `pulumi:"scaleSetPriority"`
	Type                   *string  `pulumi:"type"`
	VmSize                 *string  `pulumi:"vmSize"`
	VnetSubnetID           *string  `pulumi:"vnetSubnetID"`
}


type AgentPoolArgs struct {
	AgentPoolName          pulumi.StringPtrInput
	AvailabilityZones      pulumi.StringArrayInput
	Count                  pulumi.IntPtrInput
	EnableAutoScaling      pulumi.BoolPtrInput
	EnableNodePublicIP     pulumi.BoolPtrInput
	MaxCount               pulumi.IntPtrInput
	MaxPods                pulumi.IntPtrInput
	MinCount               pulumi.IntPtrInput
	NodeTaints             pulumi.StringArrayInput
	OrchestratorVersion    pulumi.StringPtrInput
	OsDiskSizeGB           pulumi.IntPtrInput
	OsType                 pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	ResourceName           pulumi.StringInput
	ScaleSetEvictionPolicy pulumi.StringPtrInput
	ScaleSetPriority       pulumi.StringPtrInput
	Type                   pulumi.StringPtrInput
	VmSize                 pulumi.StringPtrInput
	VnetSubnetID           pulumi.StringPtrInput
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

func (o AgentPoolOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o AgentPoolOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o AgentPoolOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

func (o AgentPoolOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o AgentPoolOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
