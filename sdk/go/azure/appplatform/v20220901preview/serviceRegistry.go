


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceRegistry struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties ServiceRegistryPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewServiceRegistry(ctx *pulumi.Context,
	name string, args *ServiceRegistryArgs, opts ...pulumi.ResourceOption) (*ServiceRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:ServiceRegistry"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceRegistry
	err := ctx.RegisterResource("azure-native:appplatform/v20220901preview:ServiceRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceRegistryState, opts ...pulumi.ResourceOption) (*ServiceRegistry, error) {
	var resource ServiceRegistry
	err := ctx.ReadResource("azure-native:appplatform/v20220901preview:ServiceRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceRegistryState struct {
}

type ServiceRegistryState struct {
}

func (ServiceRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceRegistryState)(nil)).Elem()
}

type serviceRegistryArgs struct {
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	ServiceName         string  `pulumi:"serviceName"`
	ServiceRegistryName *string `pulumi:"serviceRegistryName"`
}


type ServiceRegistryArgs struct {
	ResourceGroupName   pulumi.StringInput
	ServiceName         pulumi.StringInput
	ServiceRegistryName pulumi.StringPtrInput
}

func (ServiceRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceRegistryArgs)(nil)).Elem()
}

type ServiceRegistryInput interface {
	pulumi.Input

	ToServiceRegistryOutput() ServiceRegistryOutput
	ToServiceRegistryOutputWithContext(ctx context.Context) ServiceRegistryOutput
}

func (*ServiceRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceRegistry)(nil)).Elem()
}

func (i *ServiceRegistry) ToServiceRegistryOutput() ServiceRegistryOutput {
	return i.ToServiceRegistryOutputWithContext(context.Background())
}

func (i *ServiceRegistry) ToServiceRegistryOutputWithContext(ctx context.Context) ServiceRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceRegistryOutput)
}

type ServiceRegistryOutput struct{ *pulumi.OutputState }

func (ServiceRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceRegistry)(nil)).Elem()
}

func (o ServiceRegistryOutput) ToServiceRegistryOutput() ServiceRegistryOutput {
	return o
}

func (o ServiceRegistryOutput) ToServiceRegistryOutputWithContext(ctx context.Context) ServiceRegistryOutput {
	return o
}

func (o ServiceRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceRegistryOutput) Properties() ServiceRegistryPropertiesResponseOutput {
	return o.ApplyT(func(v *ServiceRegistry) ServiceRegistryPropertiesResponseOutput { return v.Properties }).(ServiceRegistryPropertiesResponseOutput)
}

func (o ServiceRegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServiceRegistry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ServiceRegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRegistry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceRegistryOutput{})
}
