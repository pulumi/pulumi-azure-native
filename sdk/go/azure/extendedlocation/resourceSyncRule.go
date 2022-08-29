


package extendedlocation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceSyncRule struct {
	pulumi.CustomResourceState

	Location            pulumi.StringOutput                                 `pulumi:"location"`
	Name                pulumi.StringOutput                                 `pulumi:"name"`
	Priority            pulumi.IntPtrOutput                                 `pulumi:"priority"`
	ProvisioningState   pulumi.StringOutput                                 `pulumi:"provisioningState"`
	Selector            ResourceSyncRulePropertiesResponseSelectorPtrOutput `pulumi:"selector"`
	SystemData          SystemDataResponseOutput                            `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput                              `pulumi:"tags"`
	TargetResourceGroup pulumi.StringPtrOutput                              `pulumi:"targetResourceGroup"`
	Type                pulumi.StringOutput                                 `pulumi:"type"`
}


func NewResourceSyncRule(ctx *pulumi.Context,
	name string, args *ResourceSyncRuleArgs, opts ...pulumi.ResourceOption) (*ResourceSyncRule, error) {
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
			Type: pulumi.String("azure-native:extendedlocation/v20210831preview:ResourceSyncRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceSyncRule
	err := ctx.RegisterResource("azure-native:extendedlocation:ResourceSyncRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResourceSyncRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceSyncRuleState, opts ...pulumi.ResourceOption) (*ResourceSyncRule, error) {
	var resource ResourceSyncRule
	err := ctx.ReadResource("azure-native:extendedlocation:ResourceSyncRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourceSyncRuleState struct {
}

type ResourceSyncRuleState struct {
}

func (ResourceSyncRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceSyncRuleState)(nil)).Elem()
}

type resourceSyncRuleArgs struct {
	ChildResourceName   *string                             `pulumi:"childResourceName"`
	Location            *string                             `pulumi:"location"`
	Priority            *int                                `pulumi:"priority"`
	ResourceGroupName   string                              `pulumi:"resourceGroupName"`
	ResourceName        string                              `pulumi:"resourceName"`
	Selector            *ResourceSyncRulePropertiesSelector `pulumi:"selector"`
	Tags                map[string]string                   `pulumi:"tags"`
	TargetResourceGroup *string                             `pulumi:"targetResourceGroup"`
}


type ResourceSyncRuleArgs struct {
	ChildResourceName   pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	Priority            pulumi.IntPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringInput
	Selector            ResourceSyncRulePropertiesSelectorPtrInput
	Tags                pulumi.StringMapInput
	TargetResourceGroup pulumi.StringPtrInput
}

func (ResourceSyncRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceSyncRuleArgs)(nil)).Elem()
}

type ResourceSyncRuleInput interface {
	pulumi.Input

	ToResourceSyncRuleOutput() ResourceSyncRuleOutput
	ToResourceSyncRuleOutputWithContext(ctx context.Context) ResourceSyncRuleOutput
}

func (*ResourceSyncRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRule)(nil)).Elem()
}

func (i *ResourceSyncRule) ToResourceSyncRuleOutput() ResourceSyncRuleOutput {
	return i.ToResourceSyncRuleOutputWithContext(context.Background())
}

func (i *ResourceSyncRule) ToResourceSyncRuleOutputWithContext(ctx context.Context) ResourceSyncRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSyncRuleOutput)
}

type ResourceSyncRuleOutput struct{ *pulumi.OutputState }

func (ResourceSyncRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRule)(nil)).Elem()
}

func (o ResourceSyncRuleOutput) ToResourceSyncRuleOutput() ResourceSyncRuleOutput {
	return o
}

func (o ResourceSyncRuleOutput) ToResourceSyncRuleOutputWithContext(ctx context.Context) ResourceSyncRuleOutput {
	return o
}

func (o ResourceSyncRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ResourceSyncRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSyncRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o ResourceSyncRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ResourceSyncRuleOutput) Selector() ResourceSyncRulePropertiesResponseSelectorPtrOutput {
	return o.ApplyT(func(v *ResourceSyncRule) ResourceSyncRulePropertiesResponseSelectorPtrOutput { return v.Selector }).(ResourceSyncRulePropertiesResponseSelectorPtrOutput)
}

func (o ResourceSyncRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ResourceSyncRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ResourceSyncRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ResourceSyncRuleOutput) TargetResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringPtrOutput { return v.TargetResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ResourceSyncRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceSyncRuleOutput{})
}
