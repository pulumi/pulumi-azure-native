


package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Webhook struct {
	pulumi.CustomResourceState

	CreationTime     pulumi.StringPtrOutput                      `pulumi:"creationTime"`
	Description      pulumi.StringPtrOutput                      `pulumi:"description"`
	ExpiryTime       pulumi.StringPtrOutput                      `pulumi:"expiryTime"`
	IsEnabled        pulumi.BoolPtrOutput                        `pulumi:"isEnabled"`
	LastInvokedTime  pulumi.StringPtrOutput                      `pulumi:"lastInvokedTime"`
	LastModifiedBy   pulumi.StringPtrOutput                      `pulumi:"lastModifiedBy"`
	LastModifiedTime pulumi.StringPtrOutput                      `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput                         `pulumi:"name"`
	Parameters       pulumi.StringMapOutput                      `pulumi:"parameters"`
	RunOn            pulumi.StringPtrOutput                      `pulumi:"runOn"`
	Runbook          RunbookAssociationPropertyResponsePtrOutput `pulumi:"runbook"`
	Type             pulumi.StringOutput                         `pulumi:"type"`
	Uri              pulumi.StringPtrOutput                      `pulumi:"uri"`
}


func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Webhook"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20151031:Webhook"),
		},
	})
	opts = append(opts, aliases)
	var resource Webhook
	err := ctx.RegisterResource("azure-native:automation:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("azure-native:automation:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webhookState struct {
}

type WebhookState struct {
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	AutomationAccountName string                      `pulumi:"automationAccountName"`
	ExpiryTime            *string                     `pulumi:"expiryTime"`
	IsEnabled             *bool                       `pulumi:"isEnabled"`
	Name                  string                      `pulumi:"name"`
	Parameters            map[string]string           `pulumi:"parameters"`
	ResourceGroupName     string                      `pulumi:"resourceGroupName"`
	RunOn                 *string                     `pulumi:"runOn"`
	Runbook               *RunbookAssociationProperty `pulumi:"runbook"`
	Uri                   *string                     `pulumi:"uri"`
	WebhookName           *string                     `pulumi:"webhookName"`
}


type WebhookArgs struct {
	AutomationAccountName pulumi.StringInput
	ExpiryTime            pulumi.StringPtrInput
	IsEnabled             pulumi.BoolPtrInput
	Name                  pulumi.StringInput
	Parameters            pulumi.StringMapInput
	ResourceGroupName     pulumi.StringInput
	RunOn                 pulumi.StringPtrInput
	Runbook               RunbookAssociationPropertyPtrInput
	Uri                   pulumi.StringPtrInput
	WebhookName           pulumi.StringPtrInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil))
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

type WebhookOutput struct{ *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil))
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebhookOutput{})
}
