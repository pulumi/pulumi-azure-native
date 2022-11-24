


package v20190901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyAssignment struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput                `pulumi:"description"`
	DisplayName        pulumi.StringPtrOutput                `pulumi:"displayName"`
	EnforcementMode    pulumi.StringPtrOutput                `pulumi:"enforcementMode"`
	Identity           IdentityResponsePtrOutput             `pulumi:"identity"`
	Location           pulumi.StringPtrOutput                `pulumi:"location"`
	Metadata           pulumi.AnyOutput                      `pulumi:"metadata"`
	Name               pulumi.StringOutput                   `pulumi:"name"`
	NotScopes          pulumi.StringArrayOutput              `pulumi:"notScopes"`
	Parameters         ParameterValuesValueResponseMapOutput `pulumi:"parameters"`
	PolicyDefinitionId pulumi.StringPtrOutput                `pulumi:"policyDefinitionId"`
	Scope              pulumi.StringPtrOutput                `pulumi:"scope"`
	Sku                PolicySkuResponsePtrOutput            `pulumi:"sku"`
	Type               pulumi.StringOutput                   `pulumi:"type"`
}


func NewPolicyAssignment(ctx *pulumi.Context,
	name string, args *PolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20151001preview:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160401:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20161201:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20170601preview:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190601:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210601:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20220601:PolicyAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyAssignment
	err := ctx.RegisterResource("azure-native:authorization/v20190901:PolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAssignmentState, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	var resource PolicyAssignment
	err := ctx.ReadResource("azure-native:authorization/v20190901:PolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyAssignmentState struct {
}

type PolicyAssignmentState struct {
}

func (PolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentState)(nil)).Elem()
}

type policyAssignmentArgs struct {
	Description          *string                         `pulumi:"description"`
	DisplayName          *string                         `pulumi:"displayName"`
	EnforcementMode      *string                         `pulumi:"enforcementMode"`
	Identity             *Identity                       `pulumi:"identity"`
	Location             *string                         `pulumi:"location"`
	Metadata             interface{}                     `pulumi:"metadata"`
	NotScopes            []string                        `pulumi:"notScopes"`
	Parameters           map[string]ParameterValuesValue `pulumi:"parameters"`
	PolicyAssignmentName *string                         `pulumi:"policyAssignmentName"`
	PolicyDefinitionId   *string                         `pulumi:"policyDefinitionId"`
	Scope                string                          `pulumi:"scope"`
	Sku                  *PolicySku                      `pulumi:"sku"`
}


type PolicyAssignmentArgs struct {
	Description          pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	EnforcementMode      pulumi.StringPtrInput
	Identity             IdentityPtrInput
	Location             pulumi.StringPtrInput
	Metadata             pulumi.Input
	NotScopes            pulumi.StringArrayInput
	Parameters           ParameterValuesValueMapInput
	PolicyAssignmentName pulumi.StringPtrInput
	PolicyDefinitionId   pulumi.StringPtrInput
	Scope                pulumi.StringInput
	Sku                  PolicySkuPtrInput
}

func (PolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentArgs)(nil)).Elem()
}

type PolicyAssignmentInput interface {
	pulumi.Input

	ToPolicyAssignmentOutput() PolicyAssignmentOutput
	ToPolicyAssignmentOutputWithContext(ctx context.Context) PolicyAssignmentOutput
}

func (*PolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignment)(nil)).Elem()
}

func (i *PolicyAssignment) ToPolicyAssignmentOutput() PolicyAssignmentOutput {
	return i.ToPolicyAssignmentOutputWithContext(context.Background())
}

func (i *PolicyAssignment) ToPolicyAssignmentOutputWithContext(ctx context.Context) PolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentOutput)
}

type PolicyAssignmentOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignment)(nil)).Elem()
}

func (o PolicyAssignmentOutput) ToPolicyAssignmentOutput() PolicyAssignmentOutput {
	return o
}

func (o PolicyAssignmentOutput) ToPolicyAssignmentOutputWithContext(ctx context.Context) PolicyAssignmentOutput {
	return o
}

func (o PolicyAssignmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) EnforcementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.EnforcementMode }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o PolicyAssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o PolicyAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyAssignmentOutput) NotScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringArrayOutput { return v.NotScopes }).(pulumi.StringArrayOutput)
}

func (o PolicyAssignmentOutput) Parameters() ParameterValuesValueResponseMapOutput {
	return o.ApplyT(func(v *PolicyAssignment) ParameterValuesValueResponseMapOutput { return v.Parameters }).(ParameterValuesValueResponseMapOutput)
}

func (o PolicyAssignmentOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) Sku() PolicySkuResponsePtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) PolicySkuResponsePtrOutput { return v.Sku }).(PolicySkuResponsePtrOutput)
}

func (o PolicyAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyAssignmentOutput{})
}
