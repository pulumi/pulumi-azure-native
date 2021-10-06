


package v20180301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SupportPlanType struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewSupportPlanType(ctx *pulumi.Context,
	name string, args *SupportPlanTypeArgs, opts ...pulumi.ResourceOption) (*SupportPlanType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:addons/v20180301:SupportPlanType"),
		},
		{
			Type: pulumi.String("azure-native:addons:SupportPlanType"),
		},
		{
			Type: pulumi.String("azure-nextgen:addons:SupportPlanType"),
		},
		{
			Type: pulumi.String("azure-native:addons/v20170515:SupportPlanType"),
		},
		{
			Type: pulumi.String("azure-nextgen:addons/v20170515:SupportPlanType"),
		},
	})
	opts = append(opts, aliases)
	var resource SupportPlanType
	err := ctx.RegisterResource("azure-native:addons/v20180301:SupportPlanType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSupportPlanType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SupportPlanTypeState, opts ...pulumi.ResourceOption) (*SupportPlanType, error) {
	var resource SupportPlanType
	err := ctx.ReadResource("azure-native:addons/v20180301:SupportPlanType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type supportPlanTypeState struct {
}

type SupportPlanTypeState struct {
}

func (SupportPlanTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*supportPlanTypeState)(nil)).Elem()
}

type supportPlanTypeArgs struct {
	PlanTypeName *string `pulumi:"planTypeName"`
	ProviderName string  `pulumi:"providerName"`
}


type SupportPlanTypeArgs struct {
	PlanTypeName pulumi.StringPtrInput
	ProviderName pulumi.StringInput
}

func (SupportPlanTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*supportPlanTypeArgs)(nil)).Elem()
}

type SupportPlanTypeInput interface {
	pulumi.Input

	ToSupportPlanTypeOutput() SupportPlanTypeOutput
	ToSupportPlanTypeOutputWithContext(ctx context.Context) SupportPlanTypeOutput
}

func (*SupportPlanType) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportPlanType)(nil))
}

func (i *SupportPlanType) ToSupportPlanTypeOutput() SupportPlanTypeOutput {
	return i.ToSupportPlanTypeOutputWithContext(context.Background())
}

func (i *SupportPlanType) ToSupportPlanTypeOutputWithContext(ctx context.Context) SupportPlanTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportPlanTypeOutput)
}

type SupportPlanTypeOutput struct{ *pulumi.OutputState }

func (SupportPlanTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportPlanType)(nil))
}

func (o SupportPlanTypeOutput) ToSupportPlanTypeOutput() SupportPlanTypeOutput {
	return o
}

func (o SupportPlanTypeOutput) ToSupportPlanTypeOutputWithContext(ctx context.Context) SupportPlanTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SupportPlanTypeOutput{})
}
