


package v20200717preview

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
			Type: pulumi.String("azure-nextgen:avs/v20200717preview:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210101preview:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210601:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20211201:WorkloadNetworkDnsService"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkDnsService
	err := ctx.RegisterResource("azure-native:avs/v20200717preview:WorkloadNetworkDnsService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkDnsService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDnsServiceState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsService, error) {
	var resource WorkloadNetworkDnsService
	err := ctx.ReadResource("azure-native:avs/v20200717preview:WorkloadNetworkDnsService", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*WorkloadNetworkDnsService)(nil))
}

func (i *WorkloadNetworkDnsService) ToWorkloadNetworkDnsServiceOutput() WorkloadNetworkDnsServiceOutput {
	return i.ToWorkloadNetworkDnsServiceOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDnsService) ToWorkloadNetworkDnsServiceOutputWithContext(ctx context.Context) WorkloadNetworkDnsServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDnsServiceOutput)
}

type WorkloadNetworkDnsServiceOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDnsServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDnsService)(nil))
}

func (o WorkloadNetworkDnsServiceOutput) ToWorkloadNetworkDnsServiceOutput() WorkloadNetworkDnsServiceOutput {
	return o
}

func (o WorkloadNetworkDnsServiceOutput) ToWorkloadNetworkDnsServiceOutputWithContext(ctx context.Context) WorkloadNetworkDnsServiceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadNetworkDnsServiceInput)(nil)).Elem(), &WorkloadNetworkDnsService{})
	pulumi.RegisterOutputType(WorkloadNetworkDnsServiceOutput{})
}
