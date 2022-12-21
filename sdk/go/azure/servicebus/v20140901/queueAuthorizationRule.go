


package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type QueueAuthorizationRule struct {
	pulumi.CustomResourceState

	ClaimType    pulumi.StringPtrOutput   `pulumi:"claimType"`
	ClaimValue   pulumi.StringPtrOutput   `pulumi:"claimValue"`
	CreatedTime  pulumi.StringOutput      `pulumi:"createdTime"`
	KeyName      pulumi.StringPtrOutput   `pulumi:"keyName"`
	Location     pulumi.StringPtrOutput   `pulumi:"location"`
	ModifiedTime pulumi.StringOutput      `pulumi:"modifiedTime"`
	Name         pulumi.StringOutput      `pulumi:"name"`
	PrimaryKey   pulumi.StringPtrOutput   `pulumi:"primaryKey"`
	Rights       pulumi.StringArrayOutput `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrOutput   `pulumi:"secondaryKey"`
	Type         pulumi.StringOutput      `pulumi:"type"`
}


func NewQueueAuthorizationRule(ctx *pulumi.Context,
	name string, args *QueueAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.QueueName == nil {
		return nil, errors.New("invalid value for required argument 'QueueName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20221001preview:QueueAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource QueueAuthorizationRule
	err := ctx.RegisterResource("azure-native:servicebus/v20140901:QueueAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetQueueAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueAuthorizationRuleState, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	var resource QueueAuthorizationRule
	err := ctx.ReadResource("azure-native:servicebus/v20140901:QueueAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type queueAuthorizationRuleState struct {
}

type QueueAuthorizationRuleState struct {
}

func (QueueAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleState)(nil)).Elem()
}

type queueAuthorizationRuleArgs struct {
	AuthorizationRuleName *string        `pulumi:"authorizationRuleName"`
	ClaimType             *string        `pulumi:"claimType"`
	ClaimValue            *string        `pulumi:"claimValue"`
	KeyName               *string        `pulumi:"keyName"`
	Location              *string        `pulumi:"location"`
	Name                  *string        `pulumi:"name"`
	NamespaceName         string         `pulumi:"namespaceName"`
	PrimaryKey            *string        `pulumi:"primaryKey"`
	QueueName             string         `pulumi:"queueName"`
	ResourceGroupName     string         `pulumi:"resourceGroupName"`
	Rights                []AccessRights `pulumi:"rights"`
	SecondaryKey          *string        `pulumi:"secondaryKey"`
}


type QueueAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	ClaimType             pulumi.StringPtrInput
	ClaimValue            pulumi.StringPtrInput
	KeyName               pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	NamespaceName         pulumi.StringInput
	PrimaryKey            pulumi.StringPtrInput
	QueueName             pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Rights                AccessRightsArrayInput
	SecondaryKey          pulumi.StringPtrInput
}

func (QueueAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleArgs)(nil)).Elem()
}

type QueueAuthorizationRuleInput interface {
	pulumi.Input

	ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput
	ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput
}

func (*QueueAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAuthorizationRule)(nil)).Elem()
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return i.ToQueueAuthorizationRuleOutputWithContext(context.Background())
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRuleOutput)
}

type QueueAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (QueueAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAuthorizationRule)(nil)).Elem()
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return o
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return o
}

func (o QueueAuthorizationRuleOutput) ClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.ClaimType }).(pulumi.StringPtrOutput)
}

func (o QueueAuthorizationRuleOutput) ClaimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.ClaimValue }).(pulumi.StringPtrOutput)
}

func (o QueueAuthorizationRuleOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o QueueAuthorizationRuleOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o QueueAuthorizationRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o QueueAuthorizationRuleOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o QueueAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o QueueAuthorizationRuleOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o QueueAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o QueueAuthorizationRuleOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func (o QueueAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueAuthorizationRuleOutput{})
}
