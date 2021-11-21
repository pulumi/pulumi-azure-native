


package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceAuthorizationRule struct {
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


func NewNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, args *NamespaceAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:NamespaceAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceAuthorizationRule
	err := ctx.RegisterResource("azure-native:servicebus/v20140901:NamespaceAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	var resource NamespaceAuthorizationRule
	err := ctx.ReadResource("azure-native:servicebus/v20140901:NamespaceAuthorizationRule", name, id, state, &resource, opts...)
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
	AuthorizationRuleName *string        `pulumi:"authorizationRuleName"`
	ClaimType             *string        `pulumi:"claimType"`
	ClaimValue            *string        `pulumi:"claimValue"`
	KeyName               *string        `pulumi:"keyName"`
	Location              *string        `pulumi:"location"`
	Name                  *string        `pulumi:"name"`
	NamespaceName         string         `pulumi:"namespaceName"`
	PrimaryKey            *string        `pulumi:"primaryKey"`
	ResourceGroupName     string         `pulumi:"resourceGroupName"`
	Rights                []AccessRights `pulumi:"rights"`
	SecondaryKey          *string        `pulumi:"secondaryKey"`
}


type NamespaceAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	ClaimType             pulumi.StringPtrInput
	ClaimValue            pulumi.StringPtrInput
	KeyName               pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	NamespaceName         pulumi.StringInput
	PrimaryKey            pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Rights                AccessRightsArrayInput
	SecondaryKey          pulumi.StringPtrInput
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
	return reflect.TypeOf((*NamespaceAuthorizationRule)(nil))
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return i.ToNamespaceAuthorizationRuleOutputWithContext(context.Background())
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceAuthorizationRuleOutput)
}

type NamespaceAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (NamespaceAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceAuthorizationRule)(nil))
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return o
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceAuthorizationRuleOutput{})
}
