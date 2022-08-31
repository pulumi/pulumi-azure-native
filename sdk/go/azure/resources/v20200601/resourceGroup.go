


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceGroup struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                   `pulumi:"location"`
	ManagedBy  pulumi.StringPtrOutput                `pulumi:"managedBy"`
	Name       pulumi.StringOutput                   `pulumi:"name"`
	Properties ResourceGroupPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                `pulumi:"tags"`
	Type       pulumi.StringOutput                   `pulumi:"type"`
}


func NewResourceGroup(ctx *pulumi.Context,
	name string, args *ResourceGroupArgs, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	if args == nil {
		args = &ResourceGroupArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20151101:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160201:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160701:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160901:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20170510:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180201:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180501:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190301:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:ResourceGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceGroup
	err := ctx.RegisterResource("azure-native:resources/v20200601:ResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupState, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	var resource ResourceGroup
	err := ctx.ReadResource("azure-native:resources/v20200601:ResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourceGroupState struct {
}

type ResourceGroupState struct {
}

func (ResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupState)(nil)).Elem()
}

type resourceGroupArgs struct {
	Location          *string           `pulumi:"location"`
	ManagedBy         *string           `pulumi:"managedBy"`
	ResourceGroupName *string           `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type ResourceGroupArgs struct {
	Location          pulumi.StringPtrInput
	ManagedBy         pulumi.StringPtrInput
	ResourceGroupName pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (ResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupArgs)(nil)).Elem()
}

type ResourceGroupInput interface {
	pulumi.Input

	ToResourceGroupOutput() ResourceGroupOutput
	ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput
}

func (*ResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroup)(nil)).Elem()
}

func (i *ResourceGroup) ToResourceGroupOutput() ResourceGroupOutput {
	return i.ToResourceGroupOutputWithContext(context.Background())
}

func (i *ResourceGroup) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupOutput)
}

type ResourceGroupOutput struct{ *pulumi.OutputState }

func (ResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroup)(nil)).Elem()
}

func (o ResourceGroupOutput) ToResourceGroupOutput() ResourceGroupOutput {
	return o
}

func (o ResourceGroupOutput) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return o
}

func (o ResourceGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ResourceGroupOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceGroupOutput) Properties() ResourceGroupPropertiesResponseOutput {
	return o.ApplyT(func(v *ResourceGroup) ResourceGroupPropertiesResponseOutput { return v.Properties }).(ResourceGroupPropertiesResponseOutput)
}

func (o ResourceGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ResourceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceGroupOutput{})
}
