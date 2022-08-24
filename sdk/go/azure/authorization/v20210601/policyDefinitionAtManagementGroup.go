


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyDefinitionAtManagementGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput                     `pulumi:"description"`
	DisplayName pulumi.StringPtrOutput                     `pulumi:"displayName"`
	Metadata    pulumi.AnyOutput                           `pulumi:"metadata"`
	Mode        pulumi.StringPtrOutput                     `pulumi:"mode"`
	Name        pulumi.StringOutput                        `pulumi:"name"`
	Parameters  ParameterDefinitionsValueResponseMapOutput `pulumi:"parameters"`
	PolicyRule  pulumi.AnyOutput                           `pulumi:"policyRule"`
	PolicyType  pulumi.StringPtrOutput                     `pulumi:"policyType"`
	SystemData  SystemDataResponseOutput                   `pulumi:"systemData"`
	Type        pulumi.StringOutput                        `pulumi:"type"`
}


func NewPolicyDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, args *PolicyDefinitionAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*PolicyDefinitionAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	if isZero(args.Mode) {
		args.Mode = pulumi.StringPtr("Indexed")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20161201:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190601:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190901:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicyDefinitionAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyDefinitionAtManagementGroup
	err := ctx.RegisterResource("azure-native:authorization/v20210601:PolicyDefinitionAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyDefinitionAtManagementGroupState, opts ...pulumi.ResourceOption) (*PolicyDefinitionAtManagementGroup, error) {
	var resource PolicyDefinitionAtManagementGroup
	err := ctx.ReadResource("azure-native:authorization/v20210601:PolicyDefinitionAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyDefinitionAtManagementGroupState struct {
}

type PolicyDefinitionAtManagementGroupState struct {
}

func (PolicyDefinitionAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionAtManagementGroupState)(nil)).Elem()
}

type policyDefinitionAtManagementGroupArgs struct {
	Description          *string                              `pulumi:"description"`
	DisplayName          *string                              `pulumi:"displayName"`
	ManagementGroupId    string                               `pulumi:"managementGroupId"`
	Metadata             interface{}                          `pulumi:"metadata"`
	Mode                 *string                              `pulumi:"mode"`
	Parameters           map[string]ParameterDefinitionsValue `pulumi:"parameters"`
	PolicyDefinitionName *string                              `pulumi:"policyDefinitionName"`
	PolicyRule           interface{}                          `pulumi:"policyRule"`
	PolicyType           *string                              `pulumi:"policyType"`
}


type PolicyDefinitionAtManagementGroupArgs struct {
	Description          pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	ManagementGroupId    pulumi.StringInput
	Metadata             pulumi.Input
	Mode                 pulumi.StringPtrInput
	Parameters           ParameterDefinitionsValueMapInput
	PolicyDefinitionName pulumi.StringPtrInput
	PolicyRule           pulumi.Input
	PolicyType           pulumi.StringPtrInput
}

func (PolicyDefinitionAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionAtManagementGroupArgs)(nil)).Elem()
}

type PolicyDefinitionAtManagementGroupInput interface {
	pulumi.Input

	ToPolicyDefinitionAtManagementGroupOutput() PolicyDefinitionAtManagementGroupOutput
	ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicyDefinitionAtManagementGroupOutput
}

func (*PolicyDefinitionAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinitionAtManagementGroup)(nil)).Elem()
}

func (i *PolicyDefinitionAtManagementGroup) ToPolicyDefinitionAtManagementGroupOutput() PolicyDefinitionAtManagementGroupOutput {
	return i.ToPolicyDefinitionAtManagementGroupOutputWithContext(context.Background())
}

func (i *PolicyDefinitionAtManagementGroup) ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicyDefinitionAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionAtManagementGroupOutput)
}

type PolicyDefinitionAtManagementGroupOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinitionAtManagementGroup)(nil)).Elem()
}

func (o PolicyDefinitionAtManagementGroupOutput) ToPolicyDefinitionAtManagementGroupOutput() PolicyDefinitionAtManagementGroupOutput {
	return o
}

func (o PolicyDefinitionAtManagementGroupOutput) ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicyDefinitionAtManagementGroupOutput {
	return o
}

func (o PolicyDefinitionAtManagementGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) ParameterDefinitionsValueResponseMapOutput {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.AnyOutput { return v.PolicyRule }).(pulumi.AnyOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PolicyDefinitionAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyDefinitionAtManagementGroupOutput{})
}
