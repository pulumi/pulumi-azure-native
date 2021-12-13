


package v20150819

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	HostingMode       pulumi.StringPtrOutput    `pulumi:"hostingMode"`
	Identity          IdentityResponsePtrOutput `pulumi:"identity"`
	Location          pulumi.StringPtrOutput    `pulumi:"location"`
	Name              pulumi.StringOutput       `pulumi:"name"`
	PartitionCount    pulumi.IntPtrOutput       `pulumi:"partitionCount"`
	ProvisioningState pulumi.StringOutput       `pulumi:"provisioningState"`
	ReplicaCount      pulumi.IntPtrOutput       `pulumi:"replicaCount"`
	Sku               SkuResponsePtrOutput      `pulumi:"sku"`
	Status            pulumi.StringOutput       `pulumi:"status"`
	StatusDetails     pulumi.StringOutput       `pulumi:"statusDetails"`
	Tags              pulumi.StringMapOutput    `pulumi:"tags"`
	Type              pulumi.StringOutput       `pulumi:"type"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.HostingMode) {
		args.HostingMode = HostingMode("default")
	}
	if isZero(args.PartitionCount) {
		args.PartitionCount = pulumi.IntPtr(1)
	}
	if isZero(args.ReplicaCount) {
		args.ReplicaCount = pulumi.IntPtr(1)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:search:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20191001preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200313:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:search/v20210401preview:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:search/v20150819:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:search/v20150819:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	HostingMode       *HostingMode      `pulumi:"hostingMode"`
	Identity          *Identity         `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	PartitionCount    *int              `pulumi:"partitionCount"`
	ReplicaCount      *int              `pulumi:"replicaCount"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SearchServiceName *string           `pulumi:"searchServiceName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type ServiceArgs struct {
	HostingMode       HostingModePtrInput
	Identity          IdentityPtrInput
	Location          pulumi.StringPtrInput
	PartitionCount    pulumi.IntPtrInput
	ReplicaCount      pulumi.IntPtrInput
	ResourceGroupName pulumi.StringInput
	SearchServiceName pulumi.StringPtrInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
