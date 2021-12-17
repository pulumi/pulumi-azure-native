


package v20170419

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Suppression struct {
	pulumi.CustomResourceState

	Name          pulumi.StringOutput    `pulumi:"name"`
	SuppressionId pulumi.StringPtrOutput `pulumi:"suppressionId"`
	Ttl           pulumi.StringPtrOutput `pulumi:"ttl"`
	Type          pulumi.StringOutput    `pulumi:"type"`
}


func NewSuppression(ctx *pulumi.Context,
	name string, args *SuppressionArgs, opts ...pulumi.ResourceOption) (*Suppression, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecommendationId == nil {
		return nil, errors.New("invalid value for required argument 'RecommendationId'")
	}
	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:advisor:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20160712preview:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20170331:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20200101:Suppression"),
		},
	})
	opts = append(opts, aliases)
	var resource Suppression
	err := ctx.RegisterResource("azure-native:advisor/v20170419:Suppression", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSuppression(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SuppressionState, opts ...pulumi.ResourceOption) (*Suppression, error) {
	var resource Suppression
	err := ctx.ReadResource("azure-native:advisor/v20170419:Suppression", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type suppressionState struct {
}

type SuppressionState struct {
}

func (SuppressionState) ElementType() reflect.Type {
	return reflect.TypeOf((*suppressionState)(nil)).Elem()
}

type suppressionArgs struct {
	Name             *string `pulumi:"name"`
	RecommendationId string  `pulumi:"recommendationId"`
	ResourceUri      string  `pulumi:"resourceUri"`
	SuppressionId    *string `pulumi:"suppressionId"`
	Ttl              *string `pulumi:"ttl"`
}


type SuppressionArgs struct {
	Name             pulumi.StringPtrInput
	RecommendationId pulumi.StringInput
	ResourceUri      pulumi.StringInput
	SuppressionId    pulumi.StringPtrInput
	Ttl              pulumi.StringPtrInput
}

func (SuppressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*suppressionArgs)(nil)).Elem()
}

type SuppressionInput interface {
	pulumi.Input

	ToSuppressionOutput() SuppressionOutput
	ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput
}

func (*Suppression) ElementType() reflect.Type {
	return reflect.TypeOf((**Suppression)(nil)).Elem()
}

func (i *Suppression) ToSuppressionOutput() SuppressionOutput {
	return i.ToSuppressionOutputWithContext(context.Background())
}

func (i *Suppression) ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionOutput)
}

type SuppressionOutput struct{ *pulumi.OutputState }

func (SuppressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Suppression)(nil)).Elem()
}

func (o SuppressionOutput) ToSuppressionOutput() SuppressionOutput {
	return o
}

func (o SuppressionOutput) ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SuppressionOutput{})
}
