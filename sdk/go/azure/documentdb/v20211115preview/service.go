


package v20211115preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:documentdb/v20211115preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:documentdb/v20211115preview:Service", name, id, state, &resource, opts...)
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
	AccountName       string  `pulumi:"accountName"`
	InstanceCount     *int    `pulumi:"instanceCount"`
	InstanceSize      *string `pulumi:"instanceSize"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       *string `pulumi:"serviceName"`
	ServiceType       *string `pulumi:"serviceType"`
}


type ServiceArgs struct {
	AccountName       pulumi.StringInput
	InstanceCount     pulumi.IntPtrInput
	InstanceSize      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringPtrInput
	ServiceType       pulumi.StringPtrInput
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

func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Service) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
