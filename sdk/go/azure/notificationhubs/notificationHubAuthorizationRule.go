


package notificationhubs

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
			Type: pulumi.String("azure-nextgen:notificationhubs:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:notificationhubs/v20160301:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20170401:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:notificationhubs/v20170401:NotificationHubAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationHubAuthorizationRule
	err := ctx.RegisterResource("azure-native:notificationhubs:NotificationHubAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationHubAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationHubAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NotificationHubAuthorizationRule, error) {
	var resource NotificationHubAuthorizationRule
	err := ctx.ReadResource("azure-native:notificationhubs:NotificationHubAuthorizationRule", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*NotificationHubAuthorizationRule)(nil))
}

func (i *NotificationHubAuthorizationRule) ToNotificationHubAuthorizationRuleOutput() NotificationHubAuthorizationRuleOutput {
	return i.ToNotificationHubAuthorizationRuleOutputWithContext(context.Background())
}

func (i *NotificationHubAuthorizationRule) ToNotificationHubAuthorizationRuleOutputWithContext(ctx context.Context) NotificationHubAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubAuthorizationRuleOutput)
}

type NotificationHubAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (NotificationHubAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationHubAuthorizationRule)(nil))
}

func (o NotificationHubAuthorizationRuleOutput) ToNotificationHubAuthorizationRuleOutput() NotificationHubAuthorizationRuleOutput {
	return o
}

func (o NotificationHubAuthorizationRuleOutput) ToNotificationHubAuthorizationRuleOutputWithContext(ctx context.Context) NotificationHubAuthorizationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NotificationHubAuthorizationRuleOutput{})
}
