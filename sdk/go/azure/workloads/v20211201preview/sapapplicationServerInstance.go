


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SAPApplicationServerInstance struct {
	pulumi.CustomResourceState

	Errors            SAPVirtualInstanceErrorResponseOutput `pulumi:"errors"`
	GatewayPort       pulumi.Float64Output                  `pulumi:"gatewayPort"`
	Health            pulumi.StringOutput                   `pulumi:"health"`
	Hostname          pulumi.StringOutput                   `pulumi:"hostname"`
	IcmHttpPort       pulumi.Float64Output                  `pulumi:"icmHttpPort"`
	IcmHttpsPort      pulumi.Float64Output                  `pulumi:"icmHttpsPort"`
	InstanceNo        pulumi.StringOutput                   `pulumi:"instanceNo"`
	IpAddress         pulumi.StringOutput                   `pulumi:"ipAddress"`
	KernelPatch       pulumi.StringOutput                   `pulumi:"kernelPatch"`
	KernelVersion     pulumi.StringOutput                   `pulumi:"kernelVersion"`
	Location          pulumi.StringOutput                   `pulumi:"location"`
	Name              pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                   `pulumi:"provisioningState"`
	Status            pulumi.StringOutput                   `pulumi:"status"`
	Subnet            pulumi.StringOutput                   `pulumi:"subnet"`
	SystemData        SystemDataResponseOutput              `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                `pulumi:"tags"`
	Type              pulumi.StringOutput                   `pulumi:"type"`
	VirtualMachineId  pulumi.StringOutput                   `pulumi:"virtualMachineId"`
}


func NewSAPApplicationServerInstance(ctx *pulumi.Context,
	name string, args *SAPApplicationServerInstanceArgs, opts ...pulumi.ResourceOption) (*SAPApplicationServerInstance, error) {
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
			Type: pulumi.String("azure-native:workloads:SAPApplicationServerInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource SAPApplicationServerInstance
	err := ctx.RegisterResource("azure-native:workloads/v20211201preview:SAPApplicationServerInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSAPApplicationServerInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAPApplicationServerInstanceState, opts ...pulumi.ResourceOption) (*SAPApplicationServerInstance, error) {
	var resource SAPApplicationServerInstance
	err := ctx.ReadResource("azure-native:workloads/v20211201preview:SAPApplicationServerInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sapapplicationServerInstanceState struct {
}

type SAPApplicationServerInstanceState struct {
}

func (SAPApplicationServerInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapapplicationServerInstanceState)(nil)).Elem()
}

type sapapplicationServerInstanceArgs struct {
	ApplicationInstanceName *string           `pulumi:"applicationInstanceName"`
	Location                *string           `pulumi:"location"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	SapVirtualInstanceName  string            `pulumi:"sapVirtualInstanceName"`
	Tags                    map[string]string `pulumi:"tags"`
}


type SAPApplicationServerInstanceArgs struct {
	ApplicationInstanceName pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	SapVirtualInstanceName  pulumi.StringInput
	Tags                    pulumi.StringMapInput
}

func (SAPApplicationServerInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapapplicationServerInstanceArgs)(nil)).Elem()
}

type SAPApplicationServerInstanceInput interface {
	pulumi.Input

	ToSAPApplicationServerInstanceOutput() SAPApplicationServerInstanceOutput
	ToSAPApplicationServerInstanceOutputWithContext(ctx context.Context) SAPApplicationServerInstanceOutput
}

func (*SAPApplicationServerInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPApplicationServerInstance)(nil)).Elem()
}

func (i *SAPApplicationServerInstance) ToSAPApplicationServerInstanceOutput() SAPApplicationServerInstanceOutput {
	return i.ToSAPApplicationServerInstanceOutputWithContext(context.Background())
}

func (i *SAPApplicationServerInstance) ToSAPApplicationServerInstanceOutputWithContext(ctx context.Context) SAPApplicationServerInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPApplicationServerInstanceOutput)
}

type SAPApplicationServerInstanceOutput struct{ *pulumi.OutputState }

func (SAPApplicationServerInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPApplicationServerInstance)(nil)).Elem()
}

func (o SAPApplicationServerInstanceOutput) ToSAPApplicationServerInstanceOutput() SAPApplicationServerInstanceOutput {
	return o
}

func (o SAPApplicationServerInstanceOutput) ToSAPApplicationServerInstanceOutputWithContext(ctx context.Context) SAPApplicationServerInstanceOutput {
	return o
}

func (o SAPApplicationServerInstanceOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) SAPVirtualInstanceErrorResponseOutput { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

func (o SAPApplicationServerInstanceOutput) GatewayPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.Float64Output { return v.GatewayPort }).(pulumi.Float64Output)
}

func (o SAPApplicationServerInstanceOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) IcmHttpPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.Float64Output { return v.IcmHttpPort }).(pulumi.Float64Output)
}

func (o SAPApplicationServerInstanceOutput) IcmHttpsPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.Float64Output { return v.IcmHttpsPort }).(pulumi.Float64Output)
}

func (o SAPApplicationServerInstanceOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.InstanceNo }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.KernelPatch }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.KernelVersion }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SAPApplicationServerInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SAPApplicationServerInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SAPApplicationServerInstanceOutput) VirtualMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.VirtualMachineId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SAPApplicationServerInstanceOutput{})
}
