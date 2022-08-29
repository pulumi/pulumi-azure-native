


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedHsm struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput             `pulumi:"location"`
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties ManagedHsmPropertiesResponseOutput `pulumi:"properties"`
	Sku        ManagedHsmSkuResponsePtrOutput     `pulumi:"sku"`
	SystemData SystemDataResponseOutput           `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput             `pulumi:"tags"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewManagedHsm(ctx *pulumi.Context,
	name string, args *ManagedHsmArgs, opts ...pulumi.ResourceOption) (*ManagedHsm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToManagedHsmPropertiesPtrOutput().ApplyT(func(v *ManagedHsmProperties) *ManagedHsmProperties { return v.Defaults() }).(ManagedHsmPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:keyvault:ManagedHsm"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20200401preview:ManagedHsm"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210401preview:ManagedHsm"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210601preview:ManagedHsm"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20211001:ManagedHsm"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20220701:ManagedHsm"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedHsm
	err := ctx.RegisterResource("azure-native:keyvault/v20211101preview:ManagedHsm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedHsm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedHsmState, opts ...pulumi.ResourceOption) (*ManagedHsm, error) {
	var resource ManagedHsm
	err := ctx.ReadResource("azure-native:keyvault/v20211101preview:ManagedHsm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedHsmState struct {
}

type ManagedHsmState struct {
}

func (ManagedHsmState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedHsmState)(nil)).Elem()
}

type managedHsmArgs struct {
	Location          *string               `pulumi:"location"`
	Name              *string               `pulumi:"name"`
	Properties        *ManagedHsmProperties `pulumi:"properties"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	Sku               *ManagedHsmSku        `pulumi:"sku"`
	Tags              map[string]string     `pulumi:"tags"`
}


type ManagedHsmArgs struct {
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Properties        ManagedHsmPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               ManagedHsmSkuPtrInput
	Tags              pulumi.StringMapInput
}

func (ManagedHsmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedHsmArgs)(nil)).Elem()
}

type ManagedHsmInput interface {
	pulumi.Input

	ToManagedHsmOutput() ManagedHsmOutput
	ToManagedHsmOutputWithContext(ctx context.Context) ManagedHsmOutput
}

func (*ManagedHsm) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsm)(nil)).Elem()
}

func (i *ManagedHsm) ToManagedHsmOutput() ManagedHsmOutput {
	return i.ToManagedHsmOutputWithContext(context.Background())
}

func (i *ManagedHsm) ToManagedHsmOutputWithContext(ctx context.Context) ManagedHsmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmOutput)
}

type ManagedHsmOutput struct{ *pulumi.OutputState }

func (ManagedHsmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsm)(nil)).Elem()
}

func (o ManagedHsmOutput) ToManagedHsmOutput() ManagedHsmOutput {
	return o
}

func (o ManagedHsmOutput) ToManagedHsmOutputWithContext(ctx context.Context) ManagedHsmOutput {
	return o
}

func (o ManagedHsmOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsm) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ManagedHsmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedHsm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedHsmOutput) Properties() ManagedHsmPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagedHsm) ManagedHsmPropertiesResponseOutput { return v.Properties }).(ManagedHsmPropertiesResponseOutput)
}

func (o ManagedHsmOutput) Sku() ManagedHsmSkuResponsePtrOutput {
	return o.ApplyT(func(v *ManagedHsm) ManagedHsmSkuResponsePtrOutput { return v.Sku }).(ManagedHsmSkuResponsePtrOutput)
}

func (o ManagedHsmOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedHsm) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ManagedHsmOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedHsm) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedHsmOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedHsm) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedHsmOutput{})
}
