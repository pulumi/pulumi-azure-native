


package customproviders

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomResourceProvider struct {
	pulumi.CustomResourceState

	Actions           CustomRPActionRouteDefinitionResponseArrayOutput       `pulumi:"actions"`
	Location          pulumi.StringOutput                                    `pulumi:"location"`
	Name              pulumi.StringOutput                                    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                                    `pulumi:"provisioningState"`
	ResourceTypes     CustomRPResourceTypeRouteDefinitionResponseArrayOutput `pulumi:"resourceTypes"`
	Tags              pulumi.StringMapOutput                                 `pulumi:"tags"`
	Type              pulumi.StringOutput                                    `pulumi:"type"`
	Validations       CustomRPValidationsResponseArrayOutput                 `pulumi:"validations"`
}


func NewCustomResourceProvider(ctx *pulumi.Context,
	name string, args *CustomResourceProviderArgs, opts ...pulumi.ResourceOption) (*CustomResourceProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customproviders/v20180901preview:CustomResourceProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomResourceProvider
	err := ctx.RegisterResource("azure-native:customproviders:CustomResourceProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomResourceProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomResourceProviderState, opts ...pulumi.ResourceOption) (*CustomResourceProvider, error) {
	var resource CustomResourceProvider
	err := ctx.ReadResource("azure-native:customproviders:CustomResourceProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customResourceProviderState struct {
}

type CustomResourceProviderState struct {
}

func (CustomResourceProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceProviderState)(nil)).Elem()
}

type customResourceProviderArgs struct {
	Actions              []CustomRPActionRouteDefinition       `pulumi:"actions"`
	Location             *string                               `pulumi:"location"`
	ResourceGroupName    string                                `pulumi:"resourceGroupName"`
	ResourceProviderName *string                               `pulumi:"resourceProviderName"`
	ResourceTypes        []CustomRPResourceTypeRouteDefinition `pulumi:"resourceTypes"`
	Tags                 map[string]string                     `pulumi:"tags"`
	Validations          []CustomRPValidations                 `pulumi:"validations"`
}


type CustomResourceProviderArgs struct {
	Actions              CustomRPActionRouteDefinitionArrayInput
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	ResourceProviderName pulumi.StringPtrInput
	ResourceTypes        CustomRPResourceTypeRouteDefinitionArrayInput
	Tags                 pulumi.StringMapInput
	Validations          CustomRPValidationsArrayInput
}

func (CustomResourceProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceProviderArgs)(nil)).Elem()
}

type CustomResourceProviderInput interface {
	pulumi.Input

	ToCustomResourceProviderOutput() CustomResourceProviderOutput
	ToCustomResourceProviderOutputWithContext(ctx context.Context) CustomResourceProviderOutput
}

func (*CustomResourceProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceProvider)(nil)).Elem()
}

func (i *CustomResourceProvider) ToCustomResourceProviderOutput() CustomResourceProviderOutput {
	return i.ToCustomResourceProviderOutputWithContext(context.Background())
}

func (i *CustomResourceProvider) ToCustomResourceProviderOutputWithContext(ctx context.Context) CustomResourceProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceProviderOutput)
}

type CustomResourceProviderOutput struct{ *pulumi.OutputState }

func (CustomResourceProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceProvider)(nil)).Elem()
}

func (o CustomResourceProviderOutput) ToCustomResourceProviderOutput() CustomResourceProviderOutput {
	return o
}

func (o CustomResourceProviderOutput) ToCustomResourceProviderOutputWithContext(ctx context.Context) CustomResourceProviderOutput {
	return o
}

func (o CustomResourceProviderOutput) Actions() CustomRPActionRouteDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *CustomResourceProvider) CustomRPActionRouteDefinitionResponseArrayOutput { return v.Actions }).(CustomRPActionRouteDefinitionResponseArrayOutput)
}

func (o CustomResourceProviderOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceProvider) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CustomResourceProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CustomResourceProviderOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceProvider) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CustomResourceProviderOutput) ResourceTypes() CustomRPResourceTypeRouteDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *CustomResourceProvider) CustomRPResourceTypeRouteDefinitionResponseArrayOutput {
		return v.ResourceTypes
	}).(CustomRPResourceTypeRouteDefinitionResponseArrayOutput)
}

func (o CustomResourceProviderOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomResourceProvider) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CustomResourceProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CustomResourceProviderOutput) Validations() CustomRPValidationsResponseArrayOutput {
	return o.ApplyT(func(v *CustomResourceProvider) CustomRPValidationsResponseArrayOutput { return v.Validations }).(CustomRPValidationsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomResourceProviderOutput{})
}
