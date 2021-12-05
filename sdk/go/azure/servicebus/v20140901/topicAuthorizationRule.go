


package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TopicAuthorizationRule struct {
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


func NewTopicAuthorizationRule(ctx *pulumi.Context,
	name string, args *TopicAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
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
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:TopicAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:TopicAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:TopicAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:TopicAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:TopicAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:TopicAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:TopicAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource TopicAuthorizationRule
	err := ctx.RegisterResource("azure-native:servicebus/v20140901:TopicAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTopicAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicAuthorizationRuleState, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
	var resource TopicAuthorizationRule
	err := ctx.ReadResource("azure-native:servicebus/v20140901:TopicAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type topicAuthorizationRuleState struct {
}

type TopicAuthorizationRuleState struct {
}

func (TopicAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicAuthorizationRuleState)(nil)).Elem()
}

type topicAuthorizationRuleArgs struct {
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
	TopicName             string         `pulumi:"topicName"`
}


type TopicAuthorizationRuleArgs struct {
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
	TopicName             pulumi.StringInput
}

func (TopicAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicAuthorizationRuleArgs)(nil)).Elem()
}

type TopicAuthorizationRuleInput interface {
	pulumi.Input

	ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput
	ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput
}

func (*TopicAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicAuthorizationRule)(nil))
}

func (i *TopicAuthorizationRule) ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput {
	return i.ToTopicAuthorizationRuleOutputWithContext(context.Background())
}

func (i *TopicAuthorizationRule) ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicAuthorizationRuleOutput)
}

type TopicAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (TopicAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicAuthorizationRule)(nil))
}

func (o TopicAuthorizationRuleOutput) ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput {
	return o
}

func (o TopicAuthorizationRuleOutput) ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TopicAuthorizationRuleOutput{})
}
