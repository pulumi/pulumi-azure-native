


package v20150501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComponentCurrentBillingFeature struct {
	pulumi.CustomResourceState

	CurrentBillingFeatures pulumi.StringArrayOutput                                   `pulumi:"currentBillingFeatures"`
	DataVolumeCap          ApplicationInsightsComponentDataVolumeCapResponsePtrOutput `pulumi:"dataVolumeCap"`
}


func NewComponentCurrentBillingFeature(ctx *pulumi.Context,
	name string, args *ComponentCurrentBillingFeatureArgs, opts ...pulumi.ResourceOption) (*ComponentCurrentBillingFeature, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20150501:ComponentCurrentBillingFeature"),
		},
		{
			Type: pulumi.String("azure-native:insights:ComponentCurrentBillingFeature"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:ComponentCurrentBillingFeature"),
		},
	})
	opts = append(opts, aliases)
	var resource ComponentCurrentBillingFeature
	err := ctx.RegisterResource("azure-native:insights/v20150501:ComponentCurrentBillingFeature", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetComponentCurrentBillingFeature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentCurrentBillingFeatureState, opts ...pulumi.ResourceOption) (*ComponentCurrentBillingFeature, error) {
	var resource ComponentCurrentBillingFeature
	err := ctx.ReadResource("azure-native:insights/v20150501:ComponentCurrentBillingFeature", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type componentCurrentBillingFeatureState struct {
}

type ComponentCurrentBillingFeatureState struct {
}

func (ComponentCurrentBillingFeatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentCurrentBillingFeatureState)(nil)).Elem()
}

type componentCurrentBillingFeatureArgs struct {
	CurrentBillingFeatures []string                                   `pulumi:"currentBillingFeatures"`
	DataVolumeCap          *ApplicationInsightsComponentDataVolumeCap `pulumi:"dataVolumeCap"`
	ResourceGroupName      string                                     `pulumi:"resourceGroupName"`
	ResourceName           string                                     `pulumi:"resourceName"`
}


type ComponentCurrentBillingFeatureArgs struct {
	CurrentBillingFeatures pulumi.StringArrayInput
	DataVolumeCap          ApplicationInsightsComponentDataVolumeCapPtrInput
	ResourceGroupName      pulumi.StringInput
	ResourceName           pulumi.StringInput
}

func (ComponentCurrentBillingFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentCurrentBillingFeatureArgs)(nil)).Elem()
}

type ComponentCurrentBillingFeatureInput interface {
	pulumi.Input

	ToComponentCurrentBillingFeatureOutput() ComponentCurrentBillingFeatureOutput
	ToComponentCurrentBillingFeatureOutputWithContext(ctx context.Context) ComponentCurrentBillingFeatureOutput
}

func (*ComponentCurrentBillingFeature) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentCurrentBillingFeature)(nil))
}

func (i *ComponentCurrentBillingFeature) ToComponentCurrentBillingFeatureOutput() ComponentCurrentBillingFeatureOutput {
	return i.ToComponentCurrentBillingFeatureOutputWithContext(context.Background())
}

func (i *ComponentCurrentBillingFeature) ToComponentCurrentBillingFeatureOutputWithContext(ctx context.Context) ComponentCurrentBillingFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentCurrentBillingFeatureOutput)
}

type ComponentCurrentBillingFeatureOutput struct{ *pulumi.OutputState }

func (ComponentCurrentBillingFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentCurrentBillingFeature)(nil))
}

func (o ComponentCurrentBillingFeatureOutput) ToComponentCurrentBillingFeatureOutput() ComponentCurrentBillingFeatureOutput {
	return o
}

func (o ComponentCurrentBillingFeatureOutput) ToComponentCurrentBillingFeatureOutputWithContext(ctx context.Context) ComponentCurrentBillingFeatureOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ComponentCurrentBillingFeatureOutput{})
}
