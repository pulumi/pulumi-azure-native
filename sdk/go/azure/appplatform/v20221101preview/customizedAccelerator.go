


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomizedAccelerator struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                           `pulumi:"name"`
	Properties CustomizedAcceleratorPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput                          `pulumi:"sku"`
	SystemData SystemDataResponseOutput                      `pulumi:"systemData"`
	Type       pulumi.StringOutput                           `pulumi:"type"`
}


func NewCustomizedAccelerator(ctx *pulumi.Context,
	name string, args *CustomizedAcceleratorArgs, opts ...pulumi.ResourceOption) (*CustomizedAccelerator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationAcceleratorName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationAcceleratorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	var resource CustomizedAccelerator
	err := ctx.RegisterResource("azure-native:appplatform/v20221101preview:CustomizedAccelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomizedAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomizedAcceleratorState, opts ...pulumi.ResourceOption) (*CustomizedAccelerator, error) {
	var resource CustomizedAccelerator
	err := ctx.ReadResource("azure-native:appplatform/v20221101preview:CustomizedAccelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customizedAcceleratorState struct {
}

type CustomizedAcceleratorState struct {
}

func (CustomizedAcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*customizedAcceleratorState)(nil)).Elem()
}

type customizedAcceleratorArgs struct {
	ApplicationAcceleratorName string                           `pulumi:"applicationAcceleratorName"`
	CustomizedAcceleratorName  *string                          `pulumi:"customizedAcceleratorName"`
	Properties                 *CustomizedAcceleratorProperties `pulumi:"properties"`
	ResourceGroupName          string                           `pulumi:"resourceGroupName"`
	ServiceName                string                           `pulumi:"serviceName"`
	Sku                        *Sku                             `pulumi:"sku"`
}


type CustomizedAcceleratorArgs struct {
	ApplicationAcceleratorName pulumi.StringInput
	CustomizedAcceleratorName  pulumi.StringPtrInput
	Properties                 CustomizedAcceleratorPropertiesPtrInput
	ResourceGroupName          pulumi.StringInput
	ServiceName                pulumi.StringInput
	Sku                        SkuPtrInput
}

func (CustomizedAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customizedAcceleratorArgs)(nil)).Elem()
}

type CustomizedAcceleratorInput interface {
	pulumi.Input

	ToCustomizedAcceleratorOutput() CustomizedAcceleratorOutput
	ToCustomizedAcceleratorOutputWithContext(ctx context.Context) CustomizedAcceleratorOutput
}

func (*CustomizedAccelerator) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizedAccelerator)(nil)).Elem()
}

func (i *CustomizedAccelerator) ToCustomizedAcceleratorOutput() CustomizedAcceleratorOutput {
	return i.ToCustomizedAcceleratorOutputWithContext(context.Background())
}

func (i *CustomizedAccelerator) ToCustomizedAcceleratorOutputWithContext(ctx context.Context) CustomizedAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizedAcceleratorOutput)
}

type CustomizedAcceleratorOutput struct{ *pulumi.OutputState }

func (CustomizedAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizedAccelerator)(nil)).Elem()
}

func (o CustomizedAcceleratorOutput) ToCustomizedAcceleratorOutput() CustomizedAcceleratorOutput {
	return o
}

func (o CustomizedAcceleratorOutput) ToCustomizedAcceleratorOutputWithContext(ctx context.Context) CustomizedAcceleratorOutput {
	return o
}

func (o CustomizedAcceleratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CustomizedAcceleratorOutput) Properties() CustomizedAcceleratorPropertiesResponseOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) CustomizedAcceleratorPropertiesResponseOutput { return v.Properties }).(CustomizedAcceleratorPropertiesResponseOutput)
}

func (o CustomizedAcceleratorOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o CustomizedAcceleratorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CustomizedAcceleratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomizedAcceleratorOutput{})
}
