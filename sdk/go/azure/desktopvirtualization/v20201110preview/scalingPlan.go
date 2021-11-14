


package v20201110preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScalingPlan struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput                      `pulumi:"description"`
	ExclusionTag       pulumi.StringPtrOutput                      `pulumi:"exclusionTag"`
	FriendlyName       pulumi.StringPtrOutput                      `pulumi:"friendlyName"`
	HostPoolReferences ScalingHostPoolReferenceResponseArrayOutput `pulumi:"hostPoolReferences"`
	HostPoolType       pulumi.StringPtrOutput                      `pulumi:"hostPoolType"`
	Location           pulumi.StringOutput                         `pulumi:"location"`
	Name               pulumi.StringOutput                         `pulumi:"name"`
	Schedules          ScalingScheduleResponseArrayOutput          `pulumi:"schedules"`
	Tags               pulumi.StringMapOutput                      `pulumi:"tags"`
	TimeZone           pulumi.StringPtrOutput                      `pulumi:"timeZone"`
	Type               pulumi.StringOutput                         `pulumi:"type"`
}


func NewScalingPlan(ctx *pulumi.Context,
	name string, args *ScalingPlanArgs, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:ScalingPlan"),
		},
	})
	opts = append(opts, aliases)
	var resource ScalingPlan
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20201110preview:ScalingPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScalingPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPlanState, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	var resource ScalingPlan
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20201110preview:ScalingPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scalingPlanState struct {
}

type ScalingPlanState struct {
}

func (ScalingPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanState)(nil)).Elem()
}

type scalingPlanArgs struct {
	Description        *string                    `pulumi:"description"`
	ExclusionTag       *string                    `pulumi:"exclusionTag"`
	FriendlyName       *string                    `pulumi:"friendlyName"`
	HostPoolReferences []ScalingHostPoolReference `pulumi:"hostPoolReferences"`
	HostPoolType       *string                    `pulumi:"hostPoolType"`
	Location           *string                    `pulumi:"location"`
	ResourceGroupName  string                     `pulumi:"resourceGroupName"`
	ScalingPlanName    *string                    `pulumi:"scalingPlanName"`
	Schedules          []ScalingSchedule          `pulumi:"schedules"`
	Tags               map[string]string          `pulumi:"tags"`
	TimeZone           *string                    `pulumi:"timeZone"`
}


type ScalingPlanArgs struct {
	Description        pulumi.StringPtrInput
	ExclusionTag       pulumi.StringPtrInput
	FriendlyName       pulumi.StringPtrInput
	HostPoolReferences ScalingHostPoolReferenceArrayInput
	HostPoolType       pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	ScalingPlanName    pulumi.StringPtrInput
	Schedules          ScalingScheduleArrayInput
	Tags               pulumi.StringMapInput
	TimeZone           pulumi.StringPtrInput
}

func (ScalingPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanArgs)(nil)).Elem()
}

type ScalingPlanInput interface {
	pulumi.Input

	ToScalingPlanOutput() ScalingPlanOutput
	ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput
}

func (*ScalingPlan) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPlan)(nil))
}

func (i *ScalingPlan) ToScalingPlanOutput() ScalingPlanOutput {
	return i.ToScalingPlanOutputWithContext(context.Background())
}

func (i *ScalingPlan) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanOutput)
}

type ScalingPlanOutput struct{ *pulumi.OutputState }

func (ScalingPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPlan)(nil))
}

func (o ScalingPlanOutput) ToScalingPlanOutput() ScalingPlanOutput {
	return o
}

func (o ScalingPlanOutput) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScalingPlanOutput{})
}
