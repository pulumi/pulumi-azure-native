


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationAccelerator struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                            `pulumi:"name"`
	Properties ApplicationAcceleratorPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput                           `pulumi:"sku"`
	SystemData SystemDataResponseOutput                       `pulumi:"systemData"`
	Type       pulumi.StringOutput                            `pulumi:"type"`
}


func NewApplicationAccelerator(ctx *pulumi.Context,
	name string, args *ApplicationAcceleratorArgs, opts ...pulumi.ResourceOption) (*ApplicationAccelerator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource ApplicationAccelerator
	err := ctx.RegisterResource("azure-native:appplatform/v20221101preview:ApplicationAccelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationAcceleratorState, opts ...pulumi.ResourceOption) (*ApplicationAccelerator, error) {
	var resource ApplicationAccelerator
	err := ctx.ReadResource("azure-native:appplatform/v20221101preview:ApplicationAccelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationAcceleratorState struct {
}

type ApplicationAcceleratorState struct {
}

func (ApplicationAcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationAcceleratorState)(nil)).Elem()
}

type applicationAcceleratorArgs struct {
	ApplicationAcceleratorName *string `pulumi:"applicationAcceleratorName"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	ServiceName                string  `pulumi:"serviceName"`
	Sku                        *Sku    `pulumi:"sku"`
}


type ApplicationAcceleratorArgs struct {
	ApplicationAcceleratorName pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	ServiceName                pulumi.StringInput
	Sku                        SkuPtrInput
}

func (ApplicationAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationAcceleratorArgs)(nil)).Elem()
}

type ApplicationAcceleratorInput interface {
	pulumi.Input

	ToApplicationAcceleratorOutput() ApplicationAcceleratorOutput
	ToApplicationAcceleratorOutputWithContext(ctx context.Context) ApplicationAcceleratorOutput
}

func (*ApplicationAccelerator) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationAccelerator)(nil)).Elem()
}

func (i *ApplicationAccelerator) ToApplicationAcceleratorOutput() ApplicationAcceleratorOutput {
	return i.ToApplicationAcceleratorOutputWithContext(context.Background())
}

func (i *ApplicationAccelerator) ToApplicationAcceleratorOutputWithContext(ctx context.Context) ApplicationAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAcceleratorOutput)
}

type ApplicationAcceleratorOutput struct{ *pulumi.OutputState }

func (ApplicationAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationAccelerator)(nil)).Elem()
}

func (o ApplicationAcceleratorOutput) ToApplicationAcceleratorOutput() ApplicationAcceleratorOutput {
	return o
}

func (o ApplicationAcceleratorOutput) ToApplicationAcceleratorOutputWithContext(ctx context.Context) ApplicationAcceleratorOutput {
	return o
}

func (o ApplicationAcceleratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationAcceleratorOutput) Properties() ApplicationAcceleratorPropertiesResponseOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) ApplicationAcceleratorPropertiesResponseOutput { return v.Properties }).(ApplicationAcceleratorPropertiesResponseOutput)
}

func (o ApplicationAcceleratorOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ApplicationAcceleratorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ApplicationAcceleratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationAcceleratorOutput{})
}
