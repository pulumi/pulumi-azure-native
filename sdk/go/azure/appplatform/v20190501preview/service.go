


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                  `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties ClusterResourcePropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput                    `pulumi:"sku"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	skuApplier := func(v Sku) *Sku { return v.Defaults() }
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().Elem().ApplyT(skuApplier).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:appplatform/v20190501preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:appplatform/v20190501preview:Service", name, id, state, &resource, opts...)
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
	Location          *string                    `pulumi:"location"`
	Properties        *ClusterResourceProperties `pulumi:"properties"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	ServiceName       *string                    `pulumi:"serviceName"`
	Sku               *Sku                       `pulumi:"sku"`
	Tags              map[string]string          `pulumi:"tags"`
}


type ServiceArgs struct {
	Location          pulumi.StringPtrInput
	Properties        ClusterResourcePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringPtrInput
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
