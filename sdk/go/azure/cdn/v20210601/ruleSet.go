


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RuleSet struct {
	pulumi.CustomResourceState

	DeploymentStatus  pulumi.StringOutput      `pulumi:"deploymentStatus"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProfileName       pulumi.StringOutput      `pulumi:"profileName"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewRuleSet(ctx *pulumi.Context,
	name string, args *RuleSetArgs, opts ...pulumi.ResourceOption) (*RuleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:RuleSet"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:RuleSet"),
		},
	})
	opts = append(opts, aliases)
	var resource RuleSet
	err := ctx.RegisterResource("azure-native:cdn/v20210601:RuleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRuleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleSetState, opts ...pulumi.ResourceOption) (*RuleSet, error) {
	var resource RuleSet
	err := ctx.ReadResource("azure-native:cdn/v20210601:RuleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ruleSetState struct {
}

type RuleSetState struct {
}

func (RuleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleSetState)(nil)).Elem()
}

type ruleSetArgs struct {
	ProfileName       string  `pulumi:"profileName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RuleSetName       *string `pulumi:"ruleSetName"`
}


type RuleSetArgs struct {
	ProfileName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RuleSetName       pulumi.StringPtrInput
}

func (RuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleSetArgs)(nil)).Elem()
}

type RuleSetInput interface {
	pulumi.Input

	ToRuleSetOutput() RuleSetOutput
	ToRuleSetOutputWithContext(ctx context.Context) RuleSetOutput
}

func (*RuleSet) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleSet)(nil)).Elem()
}

func (i *RuleSet) ToRuleSetOutput() RuleSetOutput {
	return i.ToRuleSetOutputWithContext(context.Background())
}

func (i *RuleSet) ToRuleSetOutputWithContext(ctx context.Context) RuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleSetOutput)
}

type RuleSetOutput struct{ *pulumi.OutputState }

func (RuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleSet)(nil)).Elem()
}

func (o RuleSetOutput) ToRuleSetOutput() RuleSetOutput {
	return o
}

func (o RuleSetOutput) ToRuleSetOutputWithContext(ctx context.Context) RuleSetOutput {
	return o
}

func (o RuleSetOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o RuleSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RuleSetOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

func (o RuleSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RuleSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RuleSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RuleSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RuleSetOutput{})
}
