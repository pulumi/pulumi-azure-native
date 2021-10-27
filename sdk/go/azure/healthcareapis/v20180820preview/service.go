


package v20180820preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput            `pulumi:"etag"`
	Identity   ResourceResponseIdentityPtrOutput `pulumi:"identity"`
	Kind       pulumi.StringOutput               `pulumi:"kind"`
	Location   pulumi.StringOutput               `pulumi:"location"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties ServicesPropertiesResponseOutput  `pulumi:"properties"`
	Tags       pulumi.StringMapOutput            `pulumi:"tags"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:healthcareapis/v20180820preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis:Service"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthcareapis:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20190916:Service"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthcareapis/v20190916:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20200315:Service"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthcareapis/v20200315:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20200330:Service"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthcareapis/v20200330:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210111:Service"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthcareapis/v20210111:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210601preview:Service"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthcareapis/v20210601preview:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:healthcareapis/v20180820preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:healthcareapis/v20180820preview:Service", name, id, state, &resource, opts...)
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
	Etag              *string             `pulumi:"etag"`
	Identity          *ResourceIdentity   `pulumi:"identity"`
	Kind              Kind                `pulumi:"kind"`
	Location          *string             `pulumi:"location"`
	Properties        *ServicesProperties `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	ResourceName      *string             `pulumi:"resourceName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type ServiceArgs struct {
	Etag              pulumi.StringPtrInput
	Identity          ResourceIdentityPtrInput
	Kind              KindInput
	Location          pulumi.StringPtrInput
	Properties        ServicesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
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
	return reflect.TypeOf((*Service)(nil))
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterOutputType(ServiceOutput{})
}
