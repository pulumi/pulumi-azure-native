


package v20180315preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringPtrOutput      `pulumi:"etag"`
	Kind              pulumi.StringPtrOutput      `pulumi:"kind"`
	Location          pulumi.StringOutput         `pulumi:"location"`
	Name              pulumi.StringOutput         `pulumi:"name"`
	ProvisioningState pulumi.StringOutput         `pulumi:"provisioningState"`
	PublicKey         pulumi.StringPtrOutput      `pulumi:"publicKey"`
	Sku               ServiceSkuResponsePtrOutput `pulumi:"sku"`
	Tags              pulumi.StringMapOutput      `pulumi:"tags"`
	Type              pulumi.StringOutput         `pulumi:"type"`
	VirtualSubnetId   pulumi.StringOutput         `pulumi:"virtualSubnetId"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.VirtualSubnetId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualSubnetId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20171115preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180331preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180419:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20210630:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:datamigration/v20180315preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:datamigration/v20180315preview:Service", name, id, state, &resource, opts...)
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
	Etag            *string           `pulumi:"etag"`
	GroupName       string            `pulumi:"groupName"`
	Kind            *string           `pulumi:"kind"`
	Location        *string           `pulumi:"location"`
	PublicKey       *string           `pulumi:"publicKey"`
	ServiceName     *string           `pulumi:"serviceName"`
	Sku             *ServiceSku       `pulumi:"sku"`
	Tags            map[string]string `pulumi:"tags"`
	VirtualSubnetId string            `pulumi:"virtualSubnetId"`
}


type ServiceArgs struct {
	Etag            pulumi.StringPtrInput
	GroupName       pulumi.StringInput
	Kind            pulumi.StringPtrInput
	Location        pulumi.StringPtrInput
	PublicKey       pulumi.StringPtrInput
	ServiceName     pulumi.StringPtrInput
	Sku             ServiceSkuPtrInput
	Tags            pulumi.StringMapInput
	VirtualSubnetId pulumi.StringInput
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
	pulumi.RegisterOutputType(ServiceOutput{})
}
