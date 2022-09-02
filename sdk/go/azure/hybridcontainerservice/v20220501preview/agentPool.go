


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPool struct {
	pulumi.CustomResourceState

	AvailabilityZones    pulumi.StringArrayOutput                           `pulumi:"availabilityZones"`
	CloudProviderProfile CloudProviderProfileResponsePtrOutput              `pulumi:"cloudProviderProfile"`
	Count                pulumi.IntPtrOutput                                `pulumi:"count"`
	ExtendedLocation     AgentPoolResponseExtendedLocationPtrOutput         `pulumi:"extendedLocation"`
	Location             pulumi.StringPtrOutput                             `pulumi:"location"`
	MaxCount             pulumi.IntPtrOutput                                `pulumi:"maxCount"`
	MaxPods              pulumi.IntPtrOutput                                `pulumi:"maxPods"`
	MinCount             pulumi.IntPtrOutput                                `pulumi:"minCount"`
	Mode                 pulumi.StringPtrOutput                             `pulumi:"mode"`
	Name                 pulumi.StringOutput                                `pulumi:"name"`
	NodeImageVersion     pulumi.StringPtrOutput                             `pulumi:"nodeImageVersion"`
	NodeLabels           pulumi.StringMapOutput                             `pulumi:"nodeLabels"`
	NodeTaints           pulumi.StringArrayOutput                           `pulumi:"nodeTaints"`
	OsType               pulumi.StringPtrOutput                             `pulumi:"osType"`
	ProvisioningState    pulumi.StringOutput                                `pulumi:"provisioningState"`
	Status               AgentPoolProvisioningStatusResponseStatusPtrOutput `pulumi:"status"`
	SystemData           SystemDataResponseOutput                           `pulumi:"systemData"`
	Tags                 pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                 pulumi.StringOutput                                `pulumi:"type"`
	VmSize               pulumi.StringPtrOutput                             `pulumi:"vmSize"`
}


func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProvisionedClustersName == nil {
		return nil, errors.New("invalid value for required argument 'ProvisionedClustersName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.Count) {
		args.Count = pulumi.IntPtr(1)
	}
	if isZero(args.Mode) {
		args.Mode = pulumi.StringPtr("User")
	}
	if isZero(args.OsType) {
		args.OsType = pulumi.StringPtr("Linux")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice:agentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:hybridcontainerservice/v20220501preview:agentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:hybridcontainerservice/v20220501preview:agentPool", name, id, state, &resource, opts...)
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
	AgentPoolName           *string                            `pulumi:"agentPoolName"`
	AvailabilityZones       []string                           `pulumi:"availabilityZones"`
	CloudProviderProfile    *CloudProviderProfile              `pulumi:"cloudProviderProfile"`
	Count                   *int                               `pulumi:"count"`
	ExtendedLocation        *AgentPoolExtendedLocation         `pulumi:"extendedLocation"`
	Location                *string                            `pulumi:"location"`
	MaxCount                *int                               `pulumi:"maxCount"`
	MaxPods                 *int                               `pulumi:"maxPods"`
	MinCount                *int                               `pulumi:"minCount"`
	Mode                    *string                            `pulumi:"mode"`
	NodeImageVersion        *string                            `pulumi:"nodeImageVersion"`
	NodeLabels              map[string]string                  `pulumi:"nodeLabels"`
	NodeTaints              []string                           `pulumi:"nodeTaints"`
	OsType                  *string                            `pulumi:"osType"`
	ProvisionedClustersName string                             `pulumi:"provisionedClustersName"`
	ResourceGroupName       string                             `pulumi:"resourceGroupName"`
	Status                  *AgentPoolProvisioningStatusStatus `pulumi:"status"`
	Tags                    map[string]string                  `pulumi:"tags"`
	VmSize                  *string                            `pulumi:"vmSize"`
}


type AgentPoolArgs struct {
	AgentPoolName           pulumi.StringPtrInput
	AvailabilityZones       pulumi.StringArrayInput
	CloudProviderProfile    CloudProviderProfilePtrInput
	Count                   pulumi.IntPtrInput
	ExtendedLocation        AgentPoolExtendedLocationPtrInput
	Location                pulumi.StringPtrInput
	MaxCount                pulumi.IntPtrInput
	MaxPods                 pulumi.IntPtrInput
	MinCount                pulumi.IntPtrInput
	Mode                    pulumi.StringPtrInput
	NodeImageVersion        pulumi.StringPtrInput
	NodeLabels              pulumi.StringMapInput
	NodeTaints              pulumi.StringArrayInput
	OsType                  pulumi.StringPtrInput
	ProvisionedClustersName pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	Status                  AgentPoolProvisioningStatusStatusPtrInput
	Tags                    pulumi.StringMapInput
	VmSize                  pulumi.StringPtrInput
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

func (o AgentPoolOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) CloudProviderProfileResponsePtrOutput { return v.CloudProviderProfile }).(CloudProviderProfileResponsePtrOutput)
}

func (o AgentPoolOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) ExtendedLocation() AgentPoolResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentPoolResponseExtendedLocationPtrOutput { return v.ExtendedLocation }).(AgentPoolResponseExtendedLocationPtrOutput)
}

func (o AgentPoolOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
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

func (o AgentPoolOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o AgentPoolOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o AgentPoolOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) Status() AgentPoolProvisioningStatusResponseStatusPtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentPoolProvisioningStatusResponseStatusPtrOutput { return v.Status }).(AgentPoolProvisioningStatusResponseStatusPtrOutput)
}

func (o AgentPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AgentPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AgentPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
