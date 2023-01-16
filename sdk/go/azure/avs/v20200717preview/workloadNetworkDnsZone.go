


package v20200717preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadNetworkDnsZone struct {
	pulumi.CustomResourceState

	DisplayName       pulumi.StringPtrOutput   `pulumi:"displayName"`
	DnsServerIps      pulumi.StringArrayOutput `pulumi:"dnsServerIps"`
	DnsServices       pulumi.Float64PtrOutput  `pulumi:"dnsServices"`
	Domain            pulumi.StringArrayOutput `pulumi:"domain"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Revision          pulumi.Float64PtrOutput  `pulumi:"revision"`
	SourceIp          pulumi.StringPtrOutput   `pulumi:"sourceIp"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewWorkloadNetworkDnsZone(ctx *pulumi.Context,
	name string, args *WorkloadNetworkDnsZoneArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsZone, error) {
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
			Type: pulumi.String("azure-native:avs:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:WorkloadNetworkDnsZone"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkDnsZone
	err := ctx.RegisterResource("azure-native:avs/v20200717preview:WorkloadNetworkDnsZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkDnsZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDnsZoneState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsZone, error) {
	var resource WorkloadNetworkDnsZone
	err := ctx.ReadResource("azure-native:avs/v20200717preview:WorkloadNetworkDnsZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadNetworkDnsZoneState struct {
}

type WorkloadNetworkDnsZoneState struct {
}

func (WorkloadNetworkDnsZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsZoneState)(nil)).Elem()
}

type workloadNetworkDnsZoneArgs struct {
	DisplayName       *string  `pulumi:"displayName"`
	DnsServerIps      []string `pulumi:"dnsServerIps"`
	DnsServices       *float64 `pulumi:"dnsServices"`
	DnsZoneId         *string  `pulumi:"dnsZoneId"`
	Domain            []string `pulumi:"domain"`
	PrivateCloudName  string   `pulumi:"privateCloudName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Revision          *float64 `pulumi:"revision"`
	SourceIp          *string  `pulumi:"sourceIp"`
}


type WorkloadNetworkDnsZoneArgs struct {
	DisplayName       pulumi.StringPtrInput
	DnsServerIps      pulumi.StringArrayInput
	DnsServices       pulumi.Float64PtrInput
	DnsZoneId         pulumi.StringPtrInput
	Domain            pulumi.StringArrayInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Revision          pulumi.Float64PtrInput
	SourceIp          pulumi.StringPtrInput
}

func (WorkloadNetworkDnsZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsZoneArgs)(nil)).Elem()
}

type WorkloadNetworkDnsZoneInput interface {
	pulumi.Input

	ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput
	ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput
}

func (*WorkloadNetworkDnsZone) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDnsZone)(nil)).Elem()
}

func (i *WorkloadNetworkDnsZone) ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput {
	return i.ToWorkloadNetworkDnsZoneOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDnsZone) ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDnsZoneOutput)
}

type WorkloadNetworkDnsZoneOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDnsZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDnsZone)(nil)).Elem()
}

func (o WorkloadNetworkDnsZoneOutput) ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput {
	return o
}

func (o WorkloadNetworkDnsZoneOutput) ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput {
	return o
}

func (o WorkloadNetworkDnsZoneOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDnsZoneOutput) DnsServerIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringArrayOutput { return v.DnsServerIps }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkDnsZoneOutput) DnsServices() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.Float64PtrOutput { return v.DnsServices }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDnsZoneOutput) Domain() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringArrayOutput { return v.Domain }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkDnsZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDnsZoneOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDnsZoneOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDnsZoneOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringPtrOutput { return v.SourceIp }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDnsZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkDnsZoneOutput{})
}
