


package workloads

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SAPCentralInstance struct {
	pulumi.CustomResourceState

	EnqueueReplicationServerProperties EnqueueReplicationServerPropertiesResponsePtrOutput `pulumi:"enqueueReplicationServerProperties"`
	EnqueueServerProperties            EnqueueServerPropertiesResponsePtrOutput            `pulumi:"enqueueServerProperties"`
	Errors                             SAPVirtualInstanceErrorResponseOutput               `pulumi:"errors"`
	GatewayServerProperties            GatewayServerPropertiesResponsePtrOutput            `pulumi:"gatewayServerProperties"`
	Health                             pulumi.StringOutput                                 `pulumi:"health"`
	InstanceNo                         pulumi.StringOutput                                 `pulumi:"instanceNo"`
	KernelPatch                        pulumi.StringOutput                                 `pulumi:"kernelPatch"`
	KernelVersion                      pulumi.StringOutput                                 `pulumi:"kernelVersion"`
	LoadBalancerDetails                LoadBalancerDetailsResponseOutput                   `pulumi:"loadBalancerDetails"`
	Location                           pulumi.StringOutput                                 `pulumi:"location"`
	MessageServerProperties            MessageServerPropertiesResponsePtrOutput            `pulumi:"messageServerProperties"`
	Name                               pulumi.StringOutput                                 `pulumi:"name"`
	ProvisioningState                  pulumi.StringOutput                                 `pulumi:"provisioningState"`
	Status                             pulumi.StringOutput                                 `pulumi:"status"`
	Subnet                             pulumi.StringOutput                                 `pulumi:"subnet"`
	SystemData                         SystemDataResponseOutput                            `pulumi:"systemData"`
	Tags                               pulumi.StringMapOutput                              `pulumi:"tags"`
	Type                               pulumi.StringOutput                                 `pulumi:"type"`
	VmDetails                          CentralServerVmDetailsResponseArrayOutput           `pulumi:"vmDetails"`
}


func NewSAPCentralInstance(ctx *pulumi.Context,
	name string, args *SAPCentralInstanceArgs, opts ...pulumi.ResourceOption) (*SAPCentralInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SapVirtualInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SapVirtualInstanceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:SAPCentralInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource SAPCentralInstance
	err := ctx.RegisterResource("azure-native:workloads:SAPCentralInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSAPCentralInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAPCentralInstanceState, opts ...pulumi.ResourceOption) (*SAPCentralInstance, error) {
	var resource SAPCentralInstance
	err := ctx.ReadResource("azure-native:workloads:SAPCentralInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sapcentralInstanceState struct {
}

type SAPCentralInstanceState struct {
}

func (SAPCentralInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapcentralInstanceState)(nil)).Elem()
}

type sapcentralInstanceArgs struct {
	CentralInstanceName    *string           `pulumi:"centralInstanceName"`
	Location               *string           `pulumi:"location"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SapVirtualInstanceName string            `pulumi:"sapVirtualInstanceName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type SAPCentralInstanceArgs struct {
	CentralInstanceName    pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	SapVirtualInstanceName pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (SAPCentralInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapcentralInstanceArgs)(nil)).Elem()
}

type SAPCentralInstanceInput interface {
	pulumi.Input

	ToSAPCentralInstanceOutput() SAPCentralInstanceOutput
	ToSAPCentralInstanceOutputWithContext(ctx context.Context) SAPCentralInstanceOutput
}

func (*SAPCentralInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPCentralInstance)(nil)).Elem()
}

func (i *SAPCentralInstance) ToSAPCentralInstanceOutput() SAPCentralInstanceOutput {
	return i.ToSAPCentralInstanceOutputWithContext(context.Background())
}

func (i *SAPCentralInstance) ToSAPCentralInstanceOutputWithContext(ctx context.Context) SAPCentralInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPCentralInstanceOutput)
}

type SAPCentralInstanceOutput struct{ *pulumi.OutputState }

func (SAPCentralInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPCentralInstance)(nil)).Elem()
}

func (o SAPCentralInstanceOutput) ToSAPCentralInstanceOutput() SAPCentralInstanceOutput {
	return o
}

func (o SAPCentralInstanceOutput) ToSAPCentralInstanceOutputWithContext(ctx context.Context) SAPCentralInstanceOutput {
	return o
}

func (o SAPCentralInstanceOutput) EnqueueReplicationServerProperties() EnqueueReplicationServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SAPCentralInstance) EnqueueReplicationServerPropertiesResponsePtrOutput {
		return v.EnqueueReplicationServerProperties
	}).(EnqueueReplicationServerPropertiesResponsePtrOutput)
}

func (o SAPCentralInstanceOutput) EnqueueServerProperties() EnqueueServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SAPCentralInstance) EnqueueServerPropertiesResponsePtrOutput { return v.EnqueueServerProperties }).(EnqueueServerPropertiesResponsePtrOutput)
}

func (o SAPCentralInstanceOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v *SAPCentralInstance) SAPVirtualInstanceErrorResponseOutput { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

func (o SAPCentralInstanceOutput) GatewayServerProperties() GatewayServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SAPCentralInstance) GatewayServerPropertiesResponsePtrOutput { return v.GatewayServerProperties }).(GatewayServerPropertiesResponsePtrOutput)
}

func (o SAPCentralInstanceOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.InstanceNo }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.KernelPatch }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.KernelVersion }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) LoadBalancerDetails() LoadBalancerDetailsResponseOutput {
	return o.ApplyT(func(v *SAPCentralInstance) LoadBalancerDetailsResponseOutput { return v.LoadBalancerDetails }).(LoadBalancerDetailsResponseOutput)
}

func (o SAPCentralInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) MessageServerProperties() MessageServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SAPCentralInstance) MessageServerPropertiesResponsePtrOutput { return v.MessageServerProperties }).(MessageServerPropertiesResponsePtrOutput)
}

func (o SAPCentralInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SAPCentralInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SAPCentralInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SAPCentralInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SAPCentralInstanceOutput) VmDetails() CentralServerVmDetailsResponseArrayOutput {
	return o.ApplyT(func(v *SAPCentralInstance) CentralServerVmDetailsResponseArrayOutput { return v.VmDetails }).(CentralServerVmDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SAPCentralInstanceOutput{})
}
