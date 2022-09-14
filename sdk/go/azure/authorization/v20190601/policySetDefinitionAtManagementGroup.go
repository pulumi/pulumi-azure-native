


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicySetDefinitionAtManagementGroup struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput                       `pulumi:"description"`
	DisplayName       pulumi.StringPtrOutput                       `pulumi:"displayName"`
	Metadata          pulumi.AnyOutput                             `pulumi:"metadata"`
	Name              pulumi.StringOutput                          `pulumi:"name"`
	Parameters        pulumi.AnyOutput                             `pulumi:"parameters"`
	PolicyDefinitions PolicyDefinitionReferenceResponseArrayOutput `pulumi:"policyDefinitions"`
	PolicyType        pulumi.StringPtrOutput                       `pulumi:"policyType"`
	Type              pulumi.StringOutput                          `pulumi:"type"`
}


func NewPolicySetDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, args *PolicySetDefinitionAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*PolicySetDefinitionAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	if args.PolicyDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDefinitions'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20170601preview:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190901:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicySetDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210601:PolicySetDefinitionAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicySetDefinitionAtManagementGroup
	err := ctx.RegisterResource("azure-native:authorization/v20190601:PolicySetDefinitionAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicySetDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicySetDefinitionAtManagementGroupState, opts ...pulumi.ResourceOption) (*PolicySetDefinitionAtManagementGroup, error) {
	var resource PolicySetDefinitionAtManagementGroup
	err := ctx.ReadResource("azure-native:authorization/v20190601:PolicySetDefinitionAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policySetDefinitionAtManagementGroupState struct {
}

type PolicySetDefinitionAtManagementGroupState struct {
}

func (PolicySetDefinitionAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionAtManagementGroupState)(nil)).Elem()
}

type policySetDefinitionAtManagementGroupArgs struct {
	Description             *string                     `pulumi:"description"`
	DisplayName             *string                     `pulumi:"displayName"`
	ManagementGroupId       string                      `pulumi:"managementGroupId"`
	Metadata                interface{}                 `pulumi:"metadata"`
	Parameters              interface{}                 `pulumi:"parameters"`
	PolicyDefinitions       []PolicyDefinitionReference `pulumi:"policyDefinitions"`
	PolicySetDefinitionName *string                     `pulumi:"policySetDefinitionName"`
	PolicyType              *string                     `pulumi:"policyType"`
}


type PolicySetDefinitionAtManagementGroupArgs struct {
	Description             pulumi.StringPtrInput
	DisplayName             pulumi.StringPtrInput
	ManagementGroupId       pulumi.StringInput
	Metadata                pulumi.Input
	Parameters              pulumi.Input
	PolicyDefinitions       PolicyDefinitionReferenceArrayInput
	PolicySetDefinitionName pulumi.StringPtrInput
	PolicyType              pulumi.StringPtrInput
}

func (PolicySetDefinitionAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionAtManagementGroupArgs)(nil)).Elem()
}

type PolicySetDefinitionAtManagementGroupInput interface {
	pulumi.Input

	ToPolicySetDefinitionAtManagementGroupOutput() PolicySetDefinitionAtManagementGroupOutput
	ToPolicySetDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicySetDefinitionAtManagementGroupOutput
}

func (*PolicySetDefinitionAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySetDefinitionAtManagementGroup)(nil)).Elem()
}

func (i *PolicySetDefinitionAtManagementGroup) ToPolicySetDefinitionAtManagementGroupOutput() PolicySetDefinitionAtManagementGroupOutput {
	return i.ToPolicySetDefinitionAtManagementGroupOutputWithContext(context.Background())
}

func (i *PolicySetDefinitionAtManagementGroup) ToPolicySetDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicySetDefinitionAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySetDefinitionAtManagementGroupOutput)
}

type PolicySetDefinitionAtManagementGroupOutput struct{ *pulumi.OutputState }

func (PolicySetDefinitionAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySetDefinitionAtManagementGroup)(nil)).Elem()
}

func (o PolicySetDefinitionAtManagementGroupOutput) ToPolicySetDefinitionAtManagementGroupOutput() PolicySetDefinitionAtManagementGroupOutput {
	return o
}

func (o PolicySetDefinitionAtManagementGroupOutput) ToPolicySetDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicySetDefinitionAtManagementGroupOutput {
	return o
}

func (o PolicySetDefinitionAtManagementGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicySetDefinitionAtManagementGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicySetDefinitionAtManagementGroupOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicySetDefinitionAtManagementGroup) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o PolicySetDefinitionAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySetDefinitionAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicySetDefinitionAtManagementGroupOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicySetDefinitionAtManagementGroup) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

func (o PolicySetDefinitionAtManagementGroupOutput) PolicyDefinitions() PolicyDefinitionReferenceResponseArrayOutput {
	return o.ApplyT(func(v *PolicySetDefinitionAtManagementGroup) PolicyDefinitionReferenceResponseArrayOutput {
		return v.PolicyDefinitions
	}).(PolicyDefinitionReferenceResponseArrayOutput)
}

func (o PolicySetDefinitionAtManagementGroupOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o PolicySetDefinitionAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySetDefinitionAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicySetDefinitionAtManagementGroupOutput{})
}
