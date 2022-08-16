


package v20190501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationServiceEnvironment struct {
	pulumi.CustomResourceState

	Identity   ManagedServiceIdentityResponsePtrOutput               `pulumi:"identity"`
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
			Type: pulumi.String("azure-native:logic:IntegrationServiceEnvironment"),
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
	Identity                          *ManagedServiceIdentity                  `pulumi:"identity"`
	IntegrationServiceEnvironmentName *string                                  `pulumi:"integrationServiceEnvironmentName"`
	Location                          *string                                  `pulumi:"location"`
	Properties                        *IntegrationServiceEnvironmentProperties `pulumi:"properties"`
	ResourceGroup                     string                                   `pulumi:"resourceGroup"`
	Sku                               *IntegrationServiceEnvironmentSku        `pulumi:"sku"`
	Tags                              map[string]string                        `pulumi:"tags"`
}


type IntegrationServiceEnvironmentArgs struct {
	Identity                          ManagedServiceIdentityPtrInput
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
	return reflect.TypeOf((**IntegrationServiceEnvironment)(nil)).Elem()
}

func (i *IntegrationServiceEnvironment) ToIntegrationServiceEnvironmentOutput() IntegrationServiceEnvironmentOutput {
	return i.ToIntegrationServiceEnvironmentOutputWithContext(context.Background())
}

func (i *IntegrationServiceEnvironment) ToIntegrationServiceEnvironmentOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentOutput)
}

type IntegrationServiceEnvironmentOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironment)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentOutput) ToIntegrationServiceEnvironmentOutput() IntegrationServiceEnvironmentOutput {
	return o
}

func (o IntegrationServiceEnvironmentOutput) ToIntegrationServiceEnvironmentOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentOutput {
	return o
}

func (o IntegrationServiceEnvironmentOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o IntegrationServiceEnvironmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IntegrationServiceEnvironmentOutput) Properties() IntegrationServiceEnvironmentPropertiesResponseOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) IntegrationServiceEnvironmentPropertiesResponseOutput {
		return v.Properties
	}).(IntegrationServiceEnvironmentPropertiesResponseOutput)
}

func (o IntegrationServiceEnvironmentOutput) Sku() IntegrationServiceEnvironmentSkuResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) IntegrationServiceEnvironmentSkuResponsePtrOutput { return v.Sku }).(IntegrationServiceEnvironmentSkuResponsePtrOutput)
}

func (o IntegrationServiceEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IntegrationServiceEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentOutput{})
}
