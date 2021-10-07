


package v20190501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationServiceEnvironment struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                                `pulumi:"location"`
	Name       pulumi.StringOutput                                   `pulumi:"name"`
	Properties IntegrationServiceEnvironmentPropertiesResponseOutput `pulumi:"properties"`
	Sku        IntegrationServiceEnvironmentSkuResponsePtrOutput     `pulumi:"sku"`
	Tags       pulumi.StringMapOutput                                `pulumi:"tags"`
	Type       pulumi.StringOutput                                   `pulumi:"type"`
}


func NewIntegrationServiceEnvironment(ctx *pulumi.Context,
	name string, args *IntegrationServiceEnvironmentArgs, opts ...pulumi.ResourceOption) (*IntegrationServiceEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:IntegrationServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:logic:IntegrationServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:IntegrationServiceEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationServiceEnvironment
	err := ctx.RegisterResource("azure-native:logic/v20190501:IntegrationServiceEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationServiceEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationServiceEnvironmentState, opts ...pulumi.ResourceOption) (*IntegrationServiceEnvironment, error) {
	var resource IntegrationServiceEnvironment
	err := ctx.ReadResource("azure-native:logic/v20190501:IntegrationServiceEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationServiceEnvironmentState struct {
}

type IntegrationServiceEnvironmentState struct {
}

func (IntegrationServiceEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationServiceEnvironmentState)(nil)).Elem()
}

type integrationServiceEnvironmentArgs struct {
	IntegrationServiceEnvironmentName *string                                  `pulumi:"integrationServiceEnvironmentName"`
	Location                          *string                                  `pulumi:"location"`
	Properties                        *IntegrationServiceEnvironmentProperties `pulumi:"properties"`
	ResourceGroup                     string                                   `pulumi:"resourceGroup"`
	Sku                               *IntegrationServiceEnvironmentSku        `pulumi:"sku"`
	Tags                              map[string]string                        `pulumi:"tags"`
}


type IntegrationServiceEnvironmentArgs struct {
	IntegrationServiceEnvironmentName pulumi.StringPtrInput
	Location                          pulumi.StringPtrInput
	Properties                        IntegrationServiceEnvironmentPropertiesPtrInput
	ResourceGroup                     pulumi.StringInput
	Sku                               IntegrationServiceEnvironmentSkuPtrInput
	Tags                              pulumi.StringMapInput
}

func (IntegrationServiceEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationServiceEnvironmentArgs)(nil)).Elem()
}

type IntegrationServiceEnvironmentInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentOutput() IntegrationServiceEnvironmentOutput
	ToIntegrationServiceEnvironmentOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentOutput
}

func (*IntegrationServiceEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironment)(nil))
}

func (i *IntegrationServiceEnvironment) ToIntegrationServiceEnvironmentOutput() IntegrationServiceEnvironmentOutput {
	return i.ToIntegrationServiceEnvironmentOutputWithContext(context.Background())
}

func (i *IntegrationServiceEnvironment) ToIntegrationServiceEnvironmentOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentOutput)
}

type IntegrationServiceEnvironmentOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironment)(nil))
}

func (o IntegrationServiceEnvironmentOutput) ToIntegrationServiceEnvironmentOutput() IntegrationServiceEnvironmentOutput {
	return o
}

func (o IntegrationServiceEnvironmentOutput) ToIntegrationServiceEnvironmentOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentOutput{})
}
