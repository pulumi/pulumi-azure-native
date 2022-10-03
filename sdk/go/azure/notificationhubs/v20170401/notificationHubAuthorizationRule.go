


package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationHubAuthorizationRule struct {
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


func NewNotificationHubAuthorizationRule(ctx *pulumi.Context,
	name string, args *NotificationHubAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NotificationHubAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.NotificationHubName == nil {
		return nil, errors.New("invalid value for required argument 'NotificationHubName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:notificationhubs:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:NotificationHubAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationHubAuthorizationRule
	err := ctx.RegisterResource("azure-native:notificationhubs/v20170401:NotificationHubAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationHubAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationHubAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NotificationHubAuthorizationRule, error) {
	var resource NotificationHubAuthorizationRule
	err := ctx.ReadResource("azure-native:notificationhubs/v20170401:NotificationHubAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notificationHubAuthorizationRuleState struct {
}

type NotificationHubAuthorizationRuleState struct {
}

func (NotificationHubAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubAuthorizationRuleState)(nil)).Elem()
}

type notificationHubAuthorizationRuleArgs struct {
	AuthorizationRuleName *string                                 `pulumi:"authorizationRuleName"`
	NamespaceName         string                                  `pulumi:"namespaceName"`
	NotificationHubName   string                                  `pulumi:"notificationHubName"`
	Properties            SharedAccessAuthorizationRuleProperties `pulumi:"properties"`
	ResourceGroupName     string                                  `pulumi:"resourceGroupName"`
}


type NotificationHubAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	NamespaceName         pulumi.StringInput
	NotificationHubName   pulumi.StringInput
	Properties            SharedAccessAuthorizationRulePropertiesInput
	ResourceGroupName     pulumi.StringInput
}

func (NotificationHubAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubAuthorizationRuleArgs)(nil)).Elem()
}

type NotificationHubAuthorizationRuleInput interface {
	pulumi.Input

	ToNotificationHubAuthorizationRuleOutput() NotificationHubAuthorizationRuleOutput
	ToNotificationHubAuthorizationRuleOutputWithContext(ctx context.Context) NotificationHubAuthorizationRuleOutput
}

func (*NotificationHubAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHubAuthorizationRule)(nil)).Elem()
}

func (i *NotificationHubAuthorizationRule) ToNotificationHubAuthorizationRuleOutput() NotificationHubAuthorizationRuleOutput {
	return i.ToNotificationHubAuthorizationRuleOutputWithContext(context.Background())
}

func (i *NotificationHubAuthorizationRule) ToNotificationHubAuthorizationRuleOutputWithContext(ctx context.Context) NotificationHubAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubAuthorizationRuleOutput)
}

type NotificationHubAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (NotificationHubAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHubAuthorizationRule)(nil)).Elem()
}

func (o NotificationHubAuthorizationRuleOutput) ToNotificationHubAuthorizationRuleOutput() NotificationHubAuthorizationRuleOutput {
	return o
}

func (o NotificationHubAuthorizationRuleOutput) ToNotificationHubAuthorizationRuleOutputWithContext(ctx context.Context) NotificationHubAuthorizationRuleOutput {
	return o
}

func (o NotificationHubAuthorizationRuleOutput) ClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.ClaimType }).(pulumi.StringOutput)
}

func (o NotificationHubAuthorizationRuleOutput) ClaimValue() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.ClaimValue }).(pulumi.StringOutput)
}

func (o NotificationHubAuthorizationRuleOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o NotificationHubAuthorizationRuleOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

func (o NotificationHubAuthorizationRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NotificationHubAuthorizationRuleOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o NotificationHubAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NotificationHubAuthorizationRuleOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o NotificationHubAuthorizationRuleOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.IntOutput { return v.Revision }).(pulumi.IntOutput)
}

func (o NotificationHubAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o NotificationHubAuthorizationRuleOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.SecondaryKey }).(pulumi.StringOutput)
}

func (o NotificationHubAuthorizationRuleOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o NotificationHubAuthorizationRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NotificationHubAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationHubAuthorizationRuleOutput{})
}
