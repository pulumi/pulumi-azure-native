


package v20190201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPool struct {
	pulumi.CustomResourceState

	AvailabilityZones   pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	Count               pulumi.IntOutput         `pulumi:"count"`
	EnableAutoScaling   pulumi.BoolPtrOutput     `pulumi:"enableAutoScaling"`
	MaxCount            pulumi.IntPtrOutput      `pulumi:"maxCount"`
	MaxPods             pulumi.IntPtrOutput      `pulumi:"maxPods"`
	MinCount            pulumi.IntPtrOutput      `pulumi:"minCount"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	OrchestratorVersion pulumi.StringPtrOutput   `pulumi:"orchestratorVersion"`
	OsDiskSizeGB        pulumi.IntPtrOutput      `pulumi:"osDiskSizeGB"`
	OsType              pulumi.StringPtrOutput   `pulumi:"osType"`
	ProvisioningState   pulumi.StringOutput      `pulumi:"provisioningState"`
	Type                pulumi.StringOutput      `pulumi:"type"`
	VmSize              pulumi.StringOutput      `pulumi:"vmSize"`
	VnetSubnetID        pulumi.StringPtrOutput   `pulumi:"vnetSubnetID"`
}


func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmSize == nil {
		return nil, errors.New("invalid value for required argument 'VmSize'")
	}
	if args.Count == nil {
		args.Count = pulumi.Int(1)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20191001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20201101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20201201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210901:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:containerservice/v20190201:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:containerservice/v20190201:AgentPool", name, id, state, &resource, opts...)
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
	AgentPoolName       *string  `pulumi:"agentPoolName"`
	AvailabilityZones   []string `pulumi:"availabilityZones"`
	Count               int      `pulumi:"count"`
	EnableAutoScaling   *bool    `pulumi:"enableAutoScaling"`
	ManagedClusterName  string   `pulumi:"managedClusterName"`
	MaxCount            *int     `pulumi:"maxCount"`
	MaxPods             *int     `pulumi:"maxPods"`
	MinCount            *int     `pulumi:"minCount"`
	OrchestratorVersion *string  `pulumi:"orchestratorVersion"`
	OsDiskSizeGB        *int     `pulumi:"osDiskSizeGB"`
	OsType              *string  `pulumi:"osType"`
	ResourceGroupName   string   `pulumi:"resourceGroupName"`
	Type                *string  `pulumi:"type"`
	VmSize              string   `pulumi:"vmSize"`
	VnetSubnetID        *string  `pulumi:"vnetSubnetID"`
}


type AgentPoolArgs struct {
	AgentPoolName       pulumi.StringPtrInput
	AvailabilityZones   pulumi.StringArrayInput
	Count               pulumi.IntInput
	EnableAutoScaling   pulumi.BoolPtrInput
	ManagedClusterName  pulumi.StringInput
	MaxCount            pulumi.IntPtrInput
	MaxPods             pulumi.IntPtrInput
	MinCount            pulumi.IntPtrInput
	OrchestratorVersion pulumi.StringPtrInput
	OsDiskSizeGB        pulumi.IntPtrInput
	OsType              pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Type                pulumi.StringPtrInput
	VmSize              pulumi.StringInput
	VnetSubnetID        pulumi.StringPtrInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPoolInput)(nil)).Elem(), &AgentPool{})
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
