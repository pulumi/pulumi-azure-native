


package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionGroup struct {
	pulumi.CustomResourceState

	AutomationRunbookReceivers AutomationRunbookReceiverResponseArrayOutput `pulumi:"automationRunbookReceivers"`
	AzureAppPushReceivers      AzureAppPushReceiverResponseArrayOutput      `pulumi:"azureAppPushReceivers"`
	EmailReceivers             EmailReceiverResponseArrayOutput             `pulumi:"emailReceivers"`
	Enabled                    pulumi.BoolOutput                            `pulumi:"enabled"`
	GroupShortName             pulumi.StringOutput                          `pulumi:"groupShortName"`
	ItsmReceivers              ItsmReceiverResponseArrayOutput              `pulumi:"itsmReceivers"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	SmsReceivers               SmsReceiverResponseArrayOutput               `pulumi:"smsReceivers"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
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
	if isZero(args.Enabled) {
		args.Enabled = pulumi.Bool(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:ActionGroup"),
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
			Type: pulumi.String("azure-native:insights/v20190601:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210901:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20220401:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20220601:ActionGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ActionGroup
	err := ctx.RegisterResource("azure-native:insights/v20170401:ActionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetActionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionGroupState, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	var resource ActionGroup
	err := ctx.ReadResource("azure-native:insights/v20170401:ActionGroup", name, id, state, &resource, opts...)
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
	AutomationRunbookReceivers []AutomationRunbookReceiver `pulumi:"automationRunbookReceivers"`
	AzureAppPushReceivers      []AzureAppPushReceiver      `pulumi:"azureAppPushReceivers"`
	EmailReceivers             []EmailReceiver             `pulumi:"emailReceivers"`
	Enabled                    bool                        `pulumi:"enabled"`
	GroupShortName             string                      `pulumi:"groupShortName"`
	ItsmReceivers              []ItsmReceiver              `pulumi:"itsmReceivers"`
	Location                   *string                     `pulumi:"location"`
	ResourceGroupName          string                      `pulumi:"resourceGroupName"`
	SmsReceivers               []SmsReceiver               `pulumi:"smsReceivers"`
	Tags                       map[string]string           `pulumi:"tags"`
	WebhookReceivers           []WebhookReceiver           `pulumi:"webhookReceivers"`
}


type ActionGroupArgs struct {
	ActionGroupName            pulumi.StringPtrInput
	AutomationRunbookReceivers AutomationRunbookReceiverArrayInput
	AzureAppPushReceivers      AzureAppPushReceiverArrayInput
	EmailReceivers             EmailReceiverArrayInput
	Enabled                    pulumi.BoolInput
	GroupShortName             pulumi.StringInput
	ItsmReceivers              ItsmReceiverArrayInput
	Location                   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	SmsReceivers               SmsReceiverArrayInput
	Tags                       pulumi.StringMapInput
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
	return reflect.TypeOf((**ActionGroup)(nil)).Elem()
}

func (i *ActionGroup) ToActionGroupOutput() ActionGroupOutput {
	return i.ToActionGroupOutputWithContext(context.Background())
}

func (i *ActionGroup) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupOutput)
}

type ActionGroupOutput struct{ *pulumi.OutputState }

func (ActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroup)(nil)).Elem()
}

func (o ActionGroupOutput) ToActionGroupOutput() ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) AutomationRunbookReceivers() AutomationRunbookReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) AutomationRunbookReceiverResponseArrayOutput { return v.AutomationRunbookReceivers }).(AutomationRunbookReceiverResponseArrayOutput)
}

func (o ActionGroupOutput) AzureAppPushReceivers() AzureAppPushReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) AzureAppPushReceiverResponseArrayOutput { return v.AzureAppPushReceivers }).(AzureAppPushReceiverResponseArrayOutput)
}

func (o ActionGroupOutput) EmailReceivers() EmailReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) EmailReceiverResponseArrayOutput { return v.EmailReceivers }).(EmailReceiverResponseArrayOutput)
}

func (o ActionGroupOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ActionGroupOutput) GroupShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.GroupShortName }).(pulumi.StringOutput)
}

func (o ActionGroupOutput) ItsmReceivers() ItsmReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) ItsmReceiverResponseArrayOutput { return v.ItsmReceivers }).(ItsmReceiverResponseArrayOutput)
}

func (o ActionGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ActionGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ActionGroupOutput) SmsReceivers() SmsReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) SmsReceiverResponseArrayOutput { return v.SmsReceivers }).(SmsReceiverResponseArrayOutput)
}

func (o ActionGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ActionGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ActionGroupOutput) WebhookReceivers() WebhookReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) WebhookReceiverResponseArrayOutput { return v.WebhookReceivers }).(WebhookReceiverResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionGroupOutput{})
}
