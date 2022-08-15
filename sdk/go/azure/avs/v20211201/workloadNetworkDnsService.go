


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadNetworkDnsService struct {
	pulumi.CustomResourceState

	DefaultDnsZone    pulumi.StringPtrOutput   `pulumi:"defaultDnsZone"`
	DisplayName       pulumi.StringPtrOutput   `pulumi:"displayName"`
	DnsServiceIp      pulumi.StringPtrOutput   `pulumi:"dnsServiceIp"`
	FqdnZones         pulumi.StringArrayOutput `pulumi:"fqdnZones"`
	LogLevel          pulumi.StringPtrOutput   `pulumi:"logLevel"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Revision          pulumi.Float64PtrOutput  `pulumi:"revision"`
	Status            pulumi.StringOutput      `pulumi:"status"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewWorkloadNetworkDnsService(ctx *pulumi.Context,
	name string, args *WorkloadNetworkDnsServiceArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkDnsService"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkDnsService
	err := ctx.RegisterResource("azure-native:avs/v20211201:WorkloadNetworkDnsService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkDnsService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDnsServiceState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsService, error) {
	var resource WorkloadNetworkDnsService
	err := ctx.ReadResource("azure-native:avs/v20211201:WorkloadNetworkDnsService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadNetworkDnsServiceState struct {
}

type WorkloadNetworkDnsServiceState struct {
}

func (WorkloadNetworkDnsServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsServiceState)(nil)).Elem()
}

type workloadNetworkDnsServiceArgs struct {
	DefaultDnsZone    *string  `pulumi:"defaultDnsZone"`
	DisplayName       *string  `pulumi:"displayName"`
	DnsServiceId      *string  `pulumi:"dnsServiceId"`
	DnsServiceIp      *string  `pulumi:"dnsServiceIp"`
	FqdnZones         []string `pulumi:"fqdnZones"`
	LogLevel          *string  `pulumi:"logLevel"`
	PrivateCloudName  string   `pulumi:"privateCloudName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Revision          *float64 `pulumi:"revision"`
}


type WorkloadNetworkDnsServiceArgs struct {
	DefaultDnsZone    pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	DnsServiceId      pulumi.StringPtrInput
	DnsServiceIp      pulumi.StringPtrInput
	FqdnZones         pulumi.StringArrayInput
	LogLevel          pulumi.StringPtrInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Revision          pulumi.Float64PtrInput
}

func (WorkloadNetworkDnsServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsServiceArgs)(nil)).Elem()
}

type WorkloadNetworkDnsServiceInput interface {
	pulumi.Input

	ToWorkloadNetworkDnsServiceOutput() WorkloadNetworkDnsServiceOutput
	ToWorkloadNetworkDnsServiceOutputWithContext(ctx context.Context) WorkloadNetworkDnsServiceOutput
}

func (*WorkloadNetworkDnsService) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDnsService)(nil)).Elem()
}

func (i *WorkloadNetworkDnsService) ToWorkloadNetworkDnsServiceOutput() WorkloadNetworkDnsServiceOutput {
	return i.ToWorkloadNetworkDnsServiceOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDnsService) ToWorkloadNetworkDnsServiceOutputWithContext(ctx context.Context) WorkloadNetworkDnsServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDnsServiceOutput)
}

type WorkloadNetworkDnsServiceOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDnsServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDnsService)(nil)).Elem()
}

func (o WorkloadNetworkDnsServiceOutput) ToWorkloadNetworkDnsServiceOutput() WorkloadNetworkDnsServiceOutput {
	return o
}

func (o WorkloadNetworkDnsServiceOutput) ToWorkloadNetworkDnsServiceOutputWithContext(ctx context.Context) WorkloadNetworkDnsServiceOutput {
	return o
}

func (o WorkloadNetworkDnsServiceOutput) DefaultDnsZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringPtrOutput { return v.DefaultDnsZone }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDnsServiceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDnsServiceOutput) DnsServiceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringPtrOutput { return v.DnsServiceIp }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDnsServiceOutput) FqdnZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringArrayOutput { return v.FqdnZones }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkDnsServiceOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringPtrOutput { return v.LogLevel }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDnsServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDnsServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDnsServiceOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDnsServiceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDnsServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkDnsServiceOutput{})
}
