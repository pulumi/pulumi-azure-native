


package v20171115preview

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
			Type: pulumi.String("azure-native:datamigration/v20180315preview:Service"),
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
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:datamigration/v20171115preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:datamigration/v20171115preview:Service", name, id, state, &resource, opts...)
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

func (o ServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Sku() ServiceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Service) ServiceSkuResponsePtrOutput { return v.Sku }).(ServiceSkuResponsePtrOutput)
}

func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ServiceOutput) VirtualSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.VirtualSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
