


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Action struct {
	pulumi.CustomResourceState

	Etag               pulumi.StringPtrOutput   `pulumi:"etag"`
	LogicAppResourceId pulumi.StringOutput      `pulumi:"logicAppResourceId"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
	WorkflowId         pulumi.StringPtrOutput   `pulumi:"workflowId"`
}


func NewAction(ctx *pulumi.Context,
	name string, args *ActionArgs, opts ...pulumi.ResourceOption) (*Action, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogicAppResourceId == nil {
		return nil, errors.New("invalid value for required argument 'LogicAppResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	if args.TriggerUri == nil {
		return nil, errors.New("invalid value for required argument 'TriggerUri'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Action"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Action"),
		},
	})
	opts = append(opts, aliases)
	var resource Action
	err := ctx.RegisterResource("azure-native:securityinsights/v20220801preview:Action", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionState, opts ...pulumi.ResourceOption) (*Action, error) {
	var resource Action
	err := ctx.ReadResource("azure-native:securityinsights/v20220801preview:Action", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type actionState struct {
}

type ActionState struct {
}

func (ActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionState)(nil)).Elem()
}

type actionArgs struct {
	ActionId           *string `pulumi:"actionId"`
	LogicAppResourceId string  `pulumi:"logicAppResourceId"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	RuleId             string  `pulumi:"ruleId"`
	TriggerUri         string  `pulumi:"triggerUri"`
	WorkspaceName      string  `pulumi:"workspaceName"`
}


type ActionArgs struct {
	ActionId           pulumi.StringPtrInput
	LogicAppResourceId pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	RuleId             pulumi.StringInput
	TriggerUri         pulumi.StringInput
	WorkspaceName      pulumi.StringInput
}

func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionArgs)(nil)).Elem()
}

type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(ctx context.Context) ActionOutput
}

func (*Action) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (i *Action) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i *Action) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Action) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ActionOutput) LogicAppResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.LogicAppResourceId }).(pulumi.StringOutput)
}

func (o ActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ActionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Action) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ActionOutput) WorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Action) pulumi.StringPtrOutput { return v.WorkflowId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
}
