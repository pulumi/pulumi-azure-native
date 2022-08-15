


package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceAuthorizationRule struct {
	pulumi.CustomResourceState

	ClaimType    pulumi.StringOutput      `pulumi:"claimType"`
	ClaimValue   pulumi.StringOutput      `pulumi:"claimValue"`
	CreatedTime  pulumi.StringOutput      `pulumi:"createdTime"`
	KeyName      pulumi.StringOutput      `pulumi:"keyName"`
	Location     pulumi.StringPtrOutput   `pulumi:"location"`
	ModifiedTime pulumi.StringOutput      `pulumi:"modifiedTime"`
	Name         pulumi.StringOutput      `pulumi:"name"`
	PrimaryKey   pulumi.StringOutput      `pulumi:"primaryKey"`
	Revision     pulumi.IntOutput         `pulumi:"revision"`
	Rights       pulumi.StringArrayOutput `pulumi:"rights"`
	SecondaryKey pulumi.StringOutput      `pulumi:"secondaryKey"`
	Sku          SkuResponsePtrOutput     `pulumi:"sku"`
	Tags         pulumi.StringMapOutput   `pulumi:"tags"`
	Type         pulumi.StringOutput      `pulumi:"type"`
}


func NewNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, args *NamespaceAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:notificationhubs:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:NamespaceAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceAuthorizationRule
	err := ctx.RegisterResource("azure-native:notificationhubs/v20170401:NamespaceAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	var resource NamespaceAuthorizationRule
	err := ctx.ReadResource("azure-native:notificationhubs/v20170401:NamespaceAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type namespaceAuthorizationRuleState struct {
}

type NamespaceAuthorizationRuleState struct {
}

func (NamespaceAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleState)(nil)).Elem()
}

type namespaceAuthorizationRuleArgs struct {
	AuthorizationRuleName *string                                 `pulumi:"authorizationRuleName"`
	NamespaceName         string                                  `pulumi:"namespaceName"`
	Properties            SharedAccessAuthorizationRuleProperties `pulumi:"properties"`
	ResourceGroupName     string                                  `pulumi:"resourceGroupName"`
}


type NamespaceAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	NamespaceName         pulumi.StringInput
	Properties            SharedAccessAuthorizationRulePropertiesInput
	ResourceGroupName     pulumi.StringInput
}

func (NamespaceAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleArgs)(nil)).Elem()
}

type NamespaceAuthorizationRuleInput interface {
	pulumi.Input

	ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput
	ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput
}

func (*NamespaceAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceAuthorizationRule)(nil)).Elem()
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return i.ToNamespaceAuthorizationRuleOutputWithContext(context.Background())
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceAuthorizationRuleOutput)
}

type NamespaceAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (NamespaceAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceAuthorizationRule)(nil)).Elem()
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return o
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return o
}

func (o NamespaceAuthorizationRuleOutput) ClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.ClaimType }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) ClaimValue() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.ClaimValue }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NamespaceAuthorizationRuleOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.IntOutput { return v.Revision }).(pulumi.IntOutput)
}

func (o NamespaceAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o NamespaceAuthorizationRuleOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.SecondaryKey }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o NamespaceAuthorizationRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NamespaceAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceAuthorizationRuleOutput{})
}
