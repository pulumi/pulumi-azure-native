


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceEndpointPolicyDefinition struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	Etag              pulumi.StringOutput      `pulumi:"etag"`
	Name              pulumi.StringPtrOutput   `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Service           pulumi.StringPtrOutput   `pulumi:"service"`
	ServiceResources  pulumi.StringArrayOutput `pulumi:"serviceResources"`
}


func NewServiceEndpointPolicyDefinition(ctx *pulumi.Context,
	name string, args *ServiceEndpointPolicyDefinitionArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicyDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceEndpointPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointPolicyName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:ServiceEndpointPolicyDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceEndpointPolicyDefinition
	err := ctx.RegisterResource("azure-native:network/v20210201:ServiceEndpointPolicyDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceEndpointPolicyDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointPolicyDefinitionState, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicyDefinition, error) {
	var resource ServiceEndpointPolicyDefinition
	err := ctx.ReadResource("azure-native:network/v20210201:ServiceEndpointPolicyDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceEndpointPolicyDefinitionState struct {
}

type ServiceEndpointPolicyDefinitionState struct {
}

func (ServiceEndpointPolicyDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyDefinitionState)(nil)).Elem()
}

type serviceEndpointPolicyDefinitionArgs struct {
	Description                         *string  `pulumi:"description"`
	Id                                  *string  `pulumi:"id"`
	Name                                *string  `pulumi:"name"`
	ResourceGroupName                   string   `pulumi:"resourceGroupName"`
	Service                             *string  `pulumi:"service"`
	ServiceEndpointPolicyDefinitionName *string  `pulumi:"serviceEndpointPolicyDefinitionName"`
	ServiceEndpointPolicyName           string   `pulumi:"serviceEndpointPolicyName"`
	ServiceResources                    []string `pulumi:"serviceResources"`
}


type ServiceEndpointPolicyDefinitionArgs struct {
	Description                         pulumi.StringPtrInput
	Id                                  pulumi.StringPtrInput
	Name                                pulumi.StringPtrInput
	ResourceGroupName                   pulumi.StringInput
	Service                             pulumi.StringPtrInput
	ServiceEndpointPolicyDefinitionName pulumi.StringPtrInput
	ServiceEndpointPolicyName           pulumi.StringInput
	ServiceResources                    pulumi.StringArrayInput
}

func (ServiceEndpointPolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyDefinitionArgs)(nil)).Elem()
}

type ServiceEndpointPolicyDefinitionInput interface {
	pulumi.Input

	ToServiceEndpointPolicyDefinitionOutput() ServiceEndpointPolicyDefinitionOutput
	ToServiceEndpointPolicyDefinitionOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionOutput
}

func (*ServiceEndpointPolicyDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPolicyDefinition)(nil))
}

func (i *ServiceEndpointPolicyDefinition) ToServiceEndpointPolicyDefinitionOutput() ServiceEndpointPolicyDefinitionOutput {
	return i.ToServiceEndpointPolicyDefinitionOutputWithContext(context.Background())
}

func (i *ServiceEndpointPolicyDefinition) ToServiceEndpointPolicyDefinitionOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPolicyDefinitionOutput)
}

type ServiceEndpointPolicyDefinitionOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPolicyDefinition)(nil))
}

func (o ServiceEndpointPolicyDefinitionOutput) ToServiceEndpointPolicyDefinitionOutput() ServiceEndpointPolicyDefinitionOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionOutput) ToServiceEndpointPolicyDefinitionOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointPolicyDefinitionOutput{})
}
