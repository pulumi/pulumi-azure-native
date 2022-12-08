


package v20220701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type AdminRule struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput      `pulumi:"etag"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewAdminRule(ctx *pulumi.Context,
	name string, args *AdminRuleArgs, opts ...pulumi.ResourceOption) (*AdminRule, error) {
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
			Type: pulumi.String("azure-native:network:AdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:AdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:AdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:AdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:AdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:AdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:AdminRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AdminRule
	err := ctx.RegisterResource("azure-native:network/v20220701:AdminRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAdminRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdminRuleState, opts ...pulumi.ResourceOption) (*AdminRule, error) {
	var resource AdminRule
	err := ctx.ReadResource("azure-native:network/v20220701:AdminRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adminRuleState struct {
}

type AdminRuleState struct {
}

func (AdminRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*adminRuleState)(nil)).Elem()
}

type adminRuleArgs struct {
	ConfigurationName  string  `pulumi:"configurationName"`
	Kind               string  `pulumi:"kind"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	RuleCollectionName string  `pulumi:"ruleCollectionName"`
	RuleName           *string `pulumi:"ruleName"`
}


type AdminRuleArgs struct {
	ConfigurationName  pulumi.StringInput
	Kind               pulumi.StringInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	RuleCollectionName pulumi.StringInput
	RuleName           pulumi.StringPtrInput
}

func (AdminRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adminRuleArgs)(nil)).Elem()
}

type AdminRuleInput interface {
	pulumi.Input

	ToAdminRuleOutput() AdminRuleOutput
	ToAdminRuleOutputWithContext(ctx context.Context) AdminRuleOutput
}

func (*AdminRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminRule)(nil)).Elem()
}

func (i *AdminRule) ToAdminRuleOutput() AdminRuleOutput {
	return i.ToAdminRuleOutputWithContext(context.Background())
}

func (i *AdminRule) ToAdminRuleOutputWithContext(ctx context.Context) AdminRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminRuleOutput)
}

type AdminRuleOutput struct{ *pulumi.OutputState }

func (AdminRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminRule)(nil)).Elem()
}

func (o AdminRuleOutput) ToAdminRuleOutput() AdminRuleOutput {
	return o
}

func (o AdminRuleOutput) ToAdminRuleOutputWithContext(ctx context.Context) AdminRuleOutput {
	return o
}

func (o AdminRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AdminRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AdminRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AdminRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AdminRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AdminRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AdminRuleOutput{})
}
