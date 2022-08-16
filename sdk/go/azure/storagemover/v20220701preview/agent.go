


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Agent struct {
	pulumi.CustomResourceState

	AgentStatus       pulumi.StringOutput                       `pulumi:"agentStatus"`
	AgentVersion      pulumi.StringOutput                       `pulumi:"agentVersion"`
	ArcResourceId     pulumi.StringOutput                       `pulumi:"arcResourceId"`
	ArcVmUuid         pulumi.StringOutput                       `pulumi:"arcVmUuid"`
	Description       pulumi.StringPtrOutput                    `pulumi:"description"`
	ErrorDetails      AgentPropertiesResponseErrorDetailsOutput `pulumi:"errorDetails"`
	LastStatusUpdate  pulumi.StringOutput                       `pulumi:"lastStatusUpdate"`
	LocalIPAddress    pulumi.StringOutput                       `pulumi:"localIPAddress"`
	MemoryInMB        pulumi.Float64Output                      `pulumi:"memoryInMB"`
	Name              pulumi.StringOutput                       `pulumi:"name"`
	NumberOfCores     pulumi.Float64Output                      `pulumi:"numberOfCores"`
	ProvisioningState pulumi.StringOutput                       `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                  `pulumi:"systemData"`
	Type              pulumi.StringOutput                       `pulumi:"type"`
	UptimeInSeconds   pulumi.Float64Output                      `pulumi:"uptimeInSeconds"`
}


func NewAgent(ctx *pulumi.Context,
	name string, args *AgentArgs, opts ...pulumi.ResourceOption) (*Agent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArcResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ArcResourceId'")
	}
	if args.ArcVmUuid == nil {
		return nil, errors.New("invalid value for required argument 'ArcVmUuid'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageMoverName == nil {
		return nil, errors.New("invalid value for required argument 'StorageMoverName'")
	}
	var resource Agent
	err := ctx.RegisterResource("azure-native:storagemover/v20220701preview:Agent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentState, opts ...pulumi.ResourceOption) (*Agent, error) {
	var resource Agent
	err := ctx.ReadResource("azure-native:storagemover/v20220701preview:Agent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type agentState struct {
}

type AgentState struct {
}

func (AgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentState)(nil)).Elem()
}

type agentArgs struct {
	AgentName         *string `pulumi:"agentName"`
	ArcResourceId     string  `pulumi:"arcResourceId"`
	ArcVmUuid         string  `pulumi:"arcVmUuid"`
	Description       *string `pulumi:"description"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	StorageMoverName  string  `pulumi:"storageMoverName"`
}


type AgentArgs struct {
	AgentName         pulumi.StringPtrInput
	ArcResourceId     pulumi.StringInput
	ArcVmUuid         pulumi.StringInput
	Description       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	StorageMoverName  pulumi.StringInput
}

func (AgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentArgs)(nil)).Elem()
}

type AgentInput interface {
	pulumi.Input

	ToAgentOutput() AgentOutput
	ToAgentOutputWithContext(ctx context.Context) AgentOutput
}

func (*Agent) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil)).Elem()
}

func (i *Agent) ToAgentOutput() AgentOutput {
	return i.ToAgentOutputWithContext(context.Background())
}

func (i *Agent) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentOutput)
}

type AgentOutput struct{ *pulumi.OutputState }

func (AgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil)).Elem()
}

func (o AgentOutput) ToAgentOutput() AgentOutput {
	return o
}

func (o AgentOutput) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return o
}

func (o AgentOutput) AgentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.AgentStatus }).(pulumi.StringOutput)
}

func (o AgentOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o AgentOutput) ArcResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.ArcResourceId }).(pulumi.StringOutput)
}

func (o AgentOutput) ArcVmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.ArcVmUuid }).(pulumi.StringOutput)
}

func (o AgentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AgentOutput) ErrorDetails() AgentPropertiesResponseErrorDetailsOutput {
	return o.ApplyT(func(v *Agent) AgentPropertiesResponseErrorDetailsOutput { return v.ErrorDetails }).(AgentPropertiesResponseErrorDetailsOutput)
}

func (o AgentOutput) LastStatusUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.LastStatusUpdate }).(pulumi.StringOutput)
}

func (o AgentOutput) LocalIPAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.LocalIPAddress }).(pulumi.StringOutput)
}

func (o AgentOutput) MemoryInMB() pulumi.Float64Output {
	return o.ApplyT(func(v *Agent) pulumi.Float64Output { return v.MemoryInMB }).(pulumi.Float64Output)
}

func (o AgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AgentOutput) NumberOfCores() pulumi.Float64Output {
	return o.ApplyT(func(v *Agent) pulumi.Float64Output { return v.NumberOfCores }).(pulumi.Float64Output)
}

func (o AgentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AgentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Agent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AgentOutput) UptimeInSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v *Agent) pulumi.Float64Output { return v.UptimeInSeconds }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(AgentOutput{})
}
