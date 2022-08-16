


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	AdminDomainName   pulumi.StringPtrOutput  `pulumi:"adminDomainName"`
	BillingDomainName pulumi.StringPtrOutput  `pulumi:"billingDomainName"`
	Etag              pulumi.StringPtrOutput  `pulumi:"etag"`
	Location          pulumi.StringPtrOutput  `pulumi:"location"`
	Name              pulumi.StringOutput     `pulumi:"name"`
	Notes             pulumi.StringPtrOutput  `pulumi:"notes"`
	Quantity          pulumi.Float64PtrOutput `pulumi:"quantity"`
	StartDate         pulumi.StringOutput     `pulumi:"startDate"`
	Tags              pulumi.StringMapOutput  `pulumi:"tags"`
	Type              pulumi.StringOutput     `pulumi:"type"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:windowsiot:Service"),
		},
		{
			Type: pulumi.String("azure-native:windowsiot/v20180216preview:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:windowsiot/v20190601:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:windowsiot/v20190601:Service", name, id, state, &resource, opts...)
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
	AdminDomainName   *string           `pulumi:"adminDomainName"`
	BillingDomainName *string           `pulumi:"billingDomainName"`
	DeviceName        *string           `pulumi:"deviceName"`
	Location          *string           `pulumi:"location"`
	Notes             *string           `pulumi:"notes"`
	Quantity          *float64          `pulumi:"quantity"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type ServiceArgs struct {
	AdminDomainName   pulumi.StringPtrInput
	BillingDomainName pulumi.StringPtrInput
	DeviceName        pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Notes             pulumi.StringPtrInput
	Quantity          pulumi.Float64PtrInput
	ResourceGroupName pulumi.StringInput
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

func (o ServiceOutput) AdminDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.AdminDomainName }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) BillingDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.BillingDomainName }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Quantity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Service) pulumi.Float64PtrOutput { return v.Quantity }).(pulumi.Float64PtrOutput)
}

func (o ServiceOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
