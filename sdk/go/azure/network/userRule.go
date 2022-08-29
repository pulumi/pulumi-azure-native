


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type UserRule struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput      `pulumi:"etag"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewUserRule(ctx *pulumi.Context,
	name string, args *UserRuleArgs, opts ...pulumi.ResourceOption) (*UserRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'RuleCollectionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:UserRule"),
		},
	})
	opts = append(opts, aliases)
	var resource UserRule
	err := ctx.RegisterResource("azure-native:network:UserRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUserRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRuleState, opts ...pulumi.ResourceOption) (*UserRule, error) {
	var resource UserRule
	err := ctx.ReadResource("azure-native:network:UserRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type userRuleState struct {
}

type UserRuleState struct {
}

func (UserRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRuleState)(nil)).Elem()
}

type userRuleArgs struct {
	ConfigurationName  string  `pulumi:"configurationName"`
	Kind               string  `pulumi:"kind"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	RuleCollectionName string  `pulumi:"ruleCollectionName"`
	RuleName           *string `pulumi:"ruleName"`
}


type UserRuleArgs struct {
	ConfigurationName  pulumi.StringInput
	Kind               pulumi.StringInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	RuleCollectionName pulumi.StringInput
	RuleName           pulumi.StringPtrInput
}

func (UserRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRuleArgs)(nil)).Elem()
}

type UserRuleInput interface {
	pulumi.Input

	ToUserRuleOutput() UserRuleOutput
	ToUserRuleOutputWithContext(ctx context.Context) UserRuleOutput
}

func (*UserRule) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRule)(nil)).Elem()
}

func (i *UserRule) ToUserRuleOutput() UserRuleOutput {
	return i.ToUserRuleOutputWithContext(context.Background())
}

func (i *UserRule) ToUserRuleOutputWithContext(ctx context.Context) UserRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRuleOutput)
}

type UserRuleOutput struct{ *pulumi.OutputState }

func (UserRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRule)(nil)).Elem()
}

func (o UserRuleOutput) ToUserRuleOutput() UserRuleOutput {
	return o
}

func (o UserRuleOutput) ToUserRuleOutputWithContext(ctx context.Context) UserRuleOutput {
	return o
}

func (o UserRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o UserRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o UserRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UserRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o UserRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserRuleOutput{})
}
