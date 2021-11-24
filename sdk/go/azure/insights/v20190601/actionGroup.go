


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionGroup struct {
	pulumi.CustomResourceState

	ArmRoleReceivers           ArmRoleReceiverResponseArrayOutput           `pulumi:"armRoleReceivers"`
	AutomationRunbookReceivers AutomationRunbookReceiverResponseArrayOutput `pulumi:"automationRunbookReceivers"`
	AzureAppPushReceivers      AzureAppPushReceiverResponseArrayOutput      `pulumi:"azureAppPushReceivers"`
	AzureFunctionReceivers     AzureFunctionReceiverResponseArrayOutput     `pulumi:"azureFunctionReceivers"`
	EmailReceivers             EmailReceiverResponseArrayOutput             `pulumi:"emailReceivers"`
	Enabled                    pulumi.BoolOutput                            `pulumi:"enabled"`
	GroupShortName             pulumi.StringOutput                          `pulumi:"groupShortName"`
	Identity                   pulumi.StringOutput                          `pulumi:"identity"`
	ItsmReceivers              ItsmReceiverResponseArrayOutput              `pulumi:"itsmReceivers"`
	Kind                       pulumi.StringOutput                          `pulumi:"kind"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	LogicAppReceivers          LogicAppReceiverResponseArrayOutput          `pulumi:"logicAppReceivers"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	SmsReceivers               SmsReceiverResponseArrayOutput               `pulumi:"smsReceivers"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
	VoiceReceivers             VoiceReceiverResponseArrayOutput             `pulumi:"voiceReceivers"`
	WebhookReceivers           WebhookReceiverResponseArrayOutput           `pulumi:"webhookReceivers"`
}


func NewActionGroup(ctx *pulumi.Context,
	name string, args *ActionGroupArgs, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupShortName == nil {
		return nil, errors.New("invalid value for required argument 'GroupShortName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Enabled == nil {
		args.Enabled = pulumi.Bool(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20170401:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180301:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180901:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20190301:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210901:ActionGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ActionGroup
	err := ctx.RegisterResource("azure-native:insights/v20190601:ActionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetActionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionGroupState, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	var resource ActionGroup
	err := ctx.ReadResource("azure-native:insights/v20190601:ActionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type actionGroupState struct {
}

type ActionGroupState struct {
}

func (ActionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupState)(nil)).Elem()
}

type actionGroupArgs struct {
	ActionGroupName            *string                     `pulumi:"actionGroupName"`
	ArmRoleReceivers           []ArmRoleReceiver           `pulumi:"armRoleReceivers"`
	AutomationRunbookReceivers []AutomationRunbookReceiver `pulumi:"automationRunbookReceivers"`
	AzureAppPushReceivers      []AzureAppPushReceiver      `pulumi:"azureAppPushReceivers"`
	AzureFunctionReceivers     []AzureFunctionReceiver     `pulumi:"azureFunctionReceivers"`
	EmailReceivers             []EmailReceiver             `pulumi:"emailReceivers"`
	Enabled                    bool                        `pulumi:"enabled"`
	GroupShortName             string                      `pulumi:"groupShortName"`
	ItsmReceivers              []ItsmReceiver              `pulumi:"itsmReceivers"`
	Location                   *string                     `pulumi:"location"`
	LogicAppReceivers          []LogicAppReceiver          `pulumi:"logicAppReceivers"`
	ResourceGroupName          string                      `pulumi:"resourceGroupName"`
	SmsReceivers               []SmsReceiver               `pulumi:"smsReceivers"`
	Tags                       map[string]string           `pulumi:"tags"`
	VoiceReceivers             []VoiceReceiver             `pulumi:"voiceReceivers"`
	WebhookReceivers           []WebhookReceiver           `pulumi:"webhookReceivers"`
}


type ActionGroupArgs struct {
	ActionGroupName            pulumi.StringPtrInput
	ArmRoleReceivers           ArmRoleReceiverArrayInput
	AutomationRunbookReceivers AutomationRunbookReceiverArrayInput
	AzureAppPushReceivers      AzureAppPushReceiverArrayInput
	AzureFunctionReceivers     AzureFunctionReceiverArrayInput
	EmailReceivers             EmailReceiverArrayInput
	Enabled                    pulumi.BoolInput
	GroupShortName             pulumi.StringInput
	ItsmReceivers              ItsmReceiverArrayInput
	Location                   pulumi.StringPtrInput
	LogicAppReceivers          LogicAppReceiverArrayInput
	ResourceGroupName          pulumi.StringInput
	SmsReceivers               SmsReceiverArrayInput
	Tags                       pulumi.StringMapInput
	VoiceReceivers             VoiceReceiverArrayInput
	WebhookReceivers           WebhookReceiverArrayInput
}

func (ActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupArgs)(nil)).Elem()
}

type ActionGroupInput interface {
	pulumi.Input

	ToActionGroupOutput() ActionGroupOutput
	ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput
}

func (*ActionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroup)(nil))
}

func (i *ActionGroup) ToActionGroupOutput() ActionGroupOutput {
	return i.ToActionGroupOutputWithContext(context.Background())
}

func (i *ActionGroup) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupOutput)
}

type ActionGroupOutput struct{ *pulumi.OutputState }

func (ActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroup)(nil))
}

func (o ActionGroupOutput) ToActionGroupOutput() ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ActionGroupOutput{})
}
