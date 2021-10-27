


package v20210601

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
			Type: pulumi.String("azure-nextgen:avs/v20210601:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20200717preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210101preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20211201:WorkloadNetworkDnsZone"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkDnsZone
	err := ctx.RegisterResource("azure-native:avs/v20210601:WorkloadNetworkDnsZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkDnsZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDnsZoneState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsZone, error) {
	var resource WorkloadNetworkDnsZone
	err := ctx.ReadResource("azure-native:avs/v20210601:WorkloadNetworkDnsZone", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*WorkloadNetworkDnsZone)(nil))
}

func (i *WorkloadNetworkDnsZone) ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput {
	return i.ToWorkloadNetworkDnsZoneOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDnsZone) ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDnsZoneOutput)
}

type WorkloadNetworkDnsZoneOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDnsZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDnsZone)(nil))
}

func (o WorkloadNetworkDnsZoneOutput) ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput {
	return o
}

func (o WorkloadNetworkDnsZoneOutput) ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadNetworkDnsZoneInput)(nil)).Elem(), &WorkloadNetworkDnsZone{})
	pulumi.RegisterOutputType(WorkloadNetworkDnsZoneOutput{})
}
