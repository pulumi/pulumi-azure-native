


package v20220201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceEndpoint struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                             `pulumi:"location"`
	Name       pulumi.StringOutput                             `pulumi:"name"`
	Properties ServiceEndpointResourceResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                        `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                          `pulumi:"tags"`
	Type       pulumi.StringOutput                             `pulumi:"type"`
}


func NewServiceEndpoint(ctx *pulumi.Context,
	name string, args *ServiceEndpointArgs, opts ...pulumi.ResourceOption) (*ServiceEndpoint, error) {
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
			Type: pulumi.String("azure-native:recommendationsservice:ServiceEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceEndpoint
	err := ctx.RegisterResource("azure-native:recommendationsservice/v20220201:ServiceEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointState, opts ...pulumi.ResourceOption) (*ServiceEndpoint, error) {
	var resource ServiceEndpoint
	err := ctx.ReadResource("azure-native:recommendationsservice/v20220201:ServiceEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceEndpointState struct {
}

type ServiceEndpointState struct {
}

func (ServiceEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointState)(nil)).Elem()
}

type serviceEndpointArgs struct {
	AccountName         string                             `pulumi:"accountName"`
	Location            *string                            `pulumi:"location"`
	Properties          *ServiceEndpointResourceProperties `pulumi:"properties"`
	ResourceGroupName   string                             `pulumi:"resourceGroupName"`
	ServiceEndpointName *string                            `pulumi:"serviceEndpointName"`
	Tags                map[string]string                  `pulumi:"tags"`
}


type ServiceEndpointArgs struct {
	AccountName         pulumi.StringInput
	Location            pulumi.StringPtrInput
	Properties          ServiceEndpointResourcePropertiesPtrInput
	ResourceGroupName   pulumi.StringInput
	ServiceEndpointName pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
}

func (ServiceEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointArgs)(nil)).Elem()
}

type ServiceEndpointInput interface {
	pulumi.Input

	ToServiceEndpointOutput() ServiceEndpointOutput
	ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput
}

func (*ServiceEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpoint)(nil)).Elem()
}

func (i *ServiceEndpoint) ToServiceEndpointOutput() ServiceEndpointOutput {
	return i.ToServiceEndpointOutputWithContext(context.Background())
}

func (i *ServiceEndpoint) ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointOutput)
}

type ServiceEndpointOutput struct{ *pulumi.OutputState }

func (ServiceEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpoint)(nil)).Elem()
}

func (o ServiceEndpointOutput) ToServiceEndpointOutput() ServiceEndpointOutput {
	return o
}

func (o ServiceEndpointOutput) ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput {
	return o
}

func (o ServiceEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServiceEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceEndpointOutput) Properties() ServiceEndpointResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *ServiceEndpoint) ServiceEndpointResourceResponsePropertiesOutput { return v.Properties }).(ServiceEndpointResourceResponsePropertiesOutput)
}

func (o ServiceEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServiceEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ServiceEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointOutput{})
}
