


package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AnalyticsItem struct {
	pulumi.CustomResourceState

	Content      pulumi.StringPtrOutput                                            `pulumi:"content"`
	Name         pulumi.StringPtrOutput                                            `pulumi:"name"`
	Properties   ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput `pulumi:"properties"`
	Scope        pulumi.StringPtrOutput                                            `pulumi:"scope"`
	TimeCreated  pulumi.StringOutput                                               `pulumi:"timeCreated"`
	TimeModified pulumi.StringOutput                                               `pulumi:"timeModified"`
	Type         pulumi.StringPtrOutput                                            `pulumi:"type"`
	Version      pulumi.StringOutput                                               `pulumi:"version"`
}


func NewAnalyticsItem(ctx *pulumi.Context,
	name string, args *AnalyticsItemArgs, opts ...pulumi.ResourceOption) (*AnalyticsItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.ScopePath == nil {
		return nil, errors.New("invalid value for required argument 'ScopePath'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20150501:AnalyticsItem"),
		},
	})
	opts = append(opts, aliases)
	var resource AnalyticsItem
	err := ctx.RegisterResource("azure-native:insights:AnalyticsItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAnalyticsItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsItemState, opts ...pulumi.ResourceOption) (*AnalyticsItem, error) {
	var resource AnalyticsItem
	err := ctx.ReadResource("azure-native:insights:AnalyticsItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type analyticsItemState struct {
}

type AnalyticsItemState struct {
}

func (AnalyticsItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsItemState)(nil)).Elem()
}

type analyticsItemArgs struct {
	Content           *string                                              `pulumi:"content"`
	Id                *string                                              `pulumi:"id"`
	Name              *string                                              `pulumi:"name"`
	OverrideItem      *bool                                                `pulumi:"overrideItem"`
	Properties        *ApplicationInsightsComponentAnalyticsItemProperties `pulumi:"properties"`
	ResourceGroupName string                                               `pulumi:"resourceGroupName"`
	ResourceName      string                                               `pulumi:"resourceName"`
	Scope             *string                                              `pulumi:"scope"`
	ScopePath         string                                               `pulumi:"scopePath"`
	Type              *string                                              `pulumi:"type"`
}


type AnalyticsItemArgs struct {
	Content           pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	OverrideItem      pulumi.BoolPtrInput
	Properties        ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
	Scope             pulumi.StringPtrInput
	ScopePath         pulumi.StringInput
	Type              pulumi.StringPtrInput
}

func (AnalyticsItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsItemArgs)(nil)).Elem()
}

type AnalyticsItemInput interface {
	pulumi.Input

	ToAnalyticsItemOutput() AnalyticsItemOutput
	ToAnalyticsItemOutputWithContext(ctx context.Context) AnalyticsItemOutput
}

func (*AnalyticsItem) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsItem)(nil)).Elem()
}

func (i *AnalyticsItem) ToAnalyticsItemOutput() AnalyticsItemOutput {
	return i.ToAnalyticsItemOutputWithContext(context.Background())
}

func (i *AnalyticsItem) ToAnalyticsItemOutputWithContext(ctx context.Context) AnalyticsItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsItemOutput)
}

type AnalyticsItemOutput struct{ *pulumi.OutputState }

func (AnalyticsItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsItem)(nil)).Elem()
}

func (o AnalyticsItemOutput) ToAnalyticsItemOutput() AnalyticsItemOutput {
	return o
}

func (o AnalyticsItemOutput) ToAnalyticsItemOutputWithContext(ctx context.Context) AnalyticsItemOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AnalyticsItemOutput{})
}
