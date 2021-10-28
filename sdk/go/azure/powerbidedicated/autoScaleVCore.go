


package powerbidedicated

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoScaleVCore struct {
	pulumi.CustomResourceState

	CapacityLimit     pulumi.IntPtrOutput             `pulumi:"capacityLimit"`
	CapacityObjectId  pulumi.StringPtrOutput          `pulumi:"capacityObjectId"`
	Location          pulumi.StringOutput             `pulumi:"location"`
	Name              pulumi.StringOutput             `pulumi:"name"`
	ProvisioningState pulumi.StringOutput             `pulumi:"provisioningState"`
	Sku               AutoScaleVCoreSkuResponseOutput `pulumi:"sku"`
	SystemData        SystemDataResponsePtrOutput     `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput          `pulumi:"tags"`
	Type              pulumi.StringOutput             `pulumi:"type"`
}


func NewAutoScaleVCore(ctx *pulumi.Context,
	name string, args *AutoScaleVCoreArgs, opts ...pulumi.ResourceOption) (*AutoScaleVCore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:powerbidedicated:AutoScaleVCore"),
		},
		{
			Type: pulumi.String("azure-native:powerbidedicated/v20210101:AutoScaleVCore"),
		},
		{
			Type: pulumi.String("azure-nextgen:powerbidedicated/v20210101:AutoScaleVCore"),
		},
	})
	opts = append(opts, aliases)
	var resource AutoScaleVCore
	err := ctx.RegisterResource("azure-native:powerbidedicated:AutoScaleVCore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAutoScaleVCore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoScaleVCoreState, opts ...pulumi.ResourceOption) (*AutoScaleVCore, error) {
	var resource AutoScaleVCore
	err := ctx.ReadResource("azure-native:powerbidedicated:AutoScaleVCore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type autoScaleVCoreState struct {
}

type AutoScaleVCoreState struct {
}

func (AutoScaleVCoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoScaleVCoreState)(nil)).Elem()
}

type autoScaleVCoreArgs struct {
	CapacityLimit     *int              `pulumi:"capacityLimit"`
	CapacityObjectId  *string           `pulumi:"capacityObjectId"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               AutoScaleVCoreSku `pulumi:"sku"`
	SystemData        *SystemData       `pulumi:"systemData"`
	Tags              map[string]string `pulumi:"tags"`
	VcoreName         *string           `pulumi:"vcoreName"`
}


type AutoScaleVCoreArgs struct {
	CapacityLimit     pulumi.IntPtrInput
	CapacityObjectId  pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               AutoScaleVCoreSkuInput
	SystemData        SystemDataPtrInput
	Tags              pulumi.StringMapInput
	VcoreName         pulumi.StringPtrInput
}

func (AutoScaleVCoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoScaleVCoreArgs)(nil)).Elem()
}

type AutoScaleVCoreInput interface {
	pulumi.Input

	ToAutoScaleVCoreOutput() AutoScaleVCoreOutput
	ToAutoScaleVCoreOutputWithContext(ctx context.Context) AutoScaleVCoreOutput
}

func (*AutoScaleVCore) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCore)(nil))
}

func (i *AutoScaleVCore) ToAutoScaleVCoreOutput() AutoScaleVCoreOutput {
	return i.ToAutoScaleVCoreOutputWithContext(context.Background())
}

func (i *AutoScaleVCore) ToAutoScaleVCoreOutputWithContext(ctx context.Context) AutoScaleVCoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreOutput)
}

type AutoScaleVCoreOutput struct{ *pulumi.OutputState }

func (AutoScaleVCoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCore)(nil))
}

func (o AutoScaleVCoreOutput) ToAutoScaleVCoreOutput() AutoScaleVCoreOutput {
	return o
}

func (o AutoScaleVCoreOutput) ToAutoScaleVCoreOutputWithContext(ctx context.Context) AutoScaleVCoreOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AutoScaleVCoreOutput{})
}
