


package v20210114preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScalingPlan struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput                                       `pulumi:"description"`
	Etag               pulumi.StringOutput                                          `pulumi:"etag"`
	ExclusionTag       pulumi.StringPtrOutput                                       `pulumi:"exclusionTag"`
	FriendlyName       pulumi.StringPtrOutput                                       `pulumi:"friendlyName"`
	HostPoolReferences ScalingHostPoolReferenceResponseArrayOutput                  `pulumi:"hostPoolReferences"`
	HostPoolType       pulumi.StringPtrOutput                                       `pulumi:"hostPoolType"`
	Identity           ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput `pulumi:"identity"`
	Kind               pulumi.StringPtrOutput                                       `pulumi:"kind"`
	Location           pulumi.StringPtrOutput                                       `pulumi:"location"`
	ManagedBy          pulumi.StringPtrOutput                                       `pulumi:"managedBy"`
	Name               pulumi.StringOutput                                          `pulumi:"name"`
	ObjectId           pulumi.StringOutput                                          `pulumi:"objectId"`
	Plan               ResourceModelWithAllowedPropertySetResponsePlanPtrOutput     `pulumi:"plan"`
	Ring               pulumi.IntPtrOutput                                          `pulumi:"ring"`
	Schedules          ScalingScheduleResponseArrayOutput                           `pulumi:"schedules"`
	Sku                ResourceModelWithAllowedPropertySetResponseSkuPtrOutput      `pulumi:"sku"`
	Tags               pulumi.StringMapOutput                                       `pulumi:"tags"`
	TimeZone           pulumi.StringPtrOutput                                       `pulumi:"timeZone"`
	Type               pulumi.StringOutput                                          `pulumi:"type"`
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
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:ScalingPlan"),
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
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:ScalingPlan"),
		},
	})
	opts = append(opts, aliases)
	var resource ScalingPlan
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20210114preview:ScalingPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScalingPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPlanState, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	var resource ScalingPlan
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20210114preview:ScalingPlan", name, id, state, &resource, opts...)
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
	Description        *string                                      `pulumi:"description"`
	ExclusionTag       *string                                      `pulumi:"exclusionTag"`
	FriendlyName       *string                                      `pulumi:"friendlyName"`
	HostPoolReferences []ScalingHostPoolReference                   `pulumi:"hostPoolReferences"`
	HostPoolType       *string                                      `pulumi:"hostPoolType"`
	Identity           *ResourceModelWithAllowedPropertySetIdentity `pulumi:"identity"`
	Kind               *string                                      `pulumi:"kind"`
	Location           *string                                      `pulumi:"location"`
	ManagedBy          *string                                      `pulumi:"managedBy"`
	Plan               *ResourceModelWithAllowedPropertySetPlan     `pulumi:"plan"`
	ResourceGroupName  string                                       `pulumi:"resourceGroupName"`
	Ring               *int                                         `pulumi:"ring"`
	ScalingPlanName    *string                                      `pulumi:"scalingPlanName"`
	Schedules          []ScalingSchedule                            `pulumi:"schedules"`
	Sku                *ResourceModelWithAllowedPropertySetSku      `pulumi:"sku"`
	Tags               map[string]string                            `pulumi:"tags"`
	TimeZone           *string                                      `pulumi:"timeZone"`
}


type ScalingPlanArgs struct {
	Description        pulumi.StringPtrInput
	ExclusionTag       pulumi.StringPtrInput
	FriendlyName       pulumi.StringPtrInput
	HostPoolReferences ScalingHostPoolReferenceArrayInput
	HostPoolType       pulumi.StringPtrInput
	Identity           ResourceModelWithAllowedPropertySetIdentityPtrInput
	Kind               pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	ManagedBy          pulumi.StringPtrInput
	Plan               ResourceModelWithAllowedPropertySetPlanPtrInput
	ResourceGroupName  pulumi.StringInput
	Ring               pulumi.IntPtrInput
	ScalingPlanName    pulumi.StringPtrInput
	Schedules          ScalingScheduleArrayInput
	Sku                ResourceModelWithAllowedPropertySetSkuPtrInput
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
	return reflect.TypeOf((**ScalingPlan)(nil)).Elem()
}

func (i *ScalingPlan) ToScalingPlanOutput() ScalingPlanOutput {
	return i.ToScalingPlanOutputWithContext(context.Background())
}

func (i *ScalingPlan) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanOutput)
}

type ScalingPlanOutput struct{ *pulumi.OutputState }

func (ScalingPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlan)(nil)).Elem()
}

func (o ScalingPlanOutput) ToScalingPlanOutput() ScalingPlanOutput {
	return o
}

func (o ScalingPlanOutput) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return o
}

func (o ScalingPlanOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ScalingPlanOutput) ExclusionTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.ExclusionTag }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) HostPoolReferences() ScalingHostPoolReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ScalingPlan) ScalingHostPoolReferenceResponseArrayOutput { return v.HostPoolReferences }).(ScalingHostPoolReferenceResponseArrayOutput)
}

func (o ScalingPlanOutput) HostPoolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.HostPoolType }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

func (o ScalingPlanOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScalingPlanOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

func (o ScalingPlanOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

func (o ScalingPlanOutput) Ring() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.IntPtrOutput { return v.Ring }).(pulumi.IntPtrOutput)
}

func (o ScalingPlanOutput) Schedules() ScalingScheduleResponseArrayOutput {
	return o.ApplyT(func(v *ScalingPlan) ScalingScheduleResponseArrayOutput { return v.Schedules }).(ScalingScheduleResponseArrayOutput)
}

func (o ScalingPlanOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

func (o ScalingPlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ScalingPlanOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalingPlanOutput{})
}
