


package v20151001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type PolicyAssignment struct {
	pulumi.CustomResourceState

	DisplayName        pulumi.StringPtrOutput `pulumi:"displayName"`
	Name               pulumi.StringPtrOutput `pulumi:"name"`
	PolicyDefinitionId pulumi.StringPtrOutput `pulumi:"policyDefinitionId"`
	Scope              pulumi.StringPtrOutput `pulumi:"scope"`
	Type               pulumi.StringPtrOutput `pulumi:"type"`
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
			Type: pulumi.String("azure-native:authorization/v20190901:PolicyAssignment"),
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
	})
	opts = append(opts, aliases)
	var resource PolicyAssignment
	err := ctx.RegisterResource("azure-native:authorization/v20151001preview:PolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAssignmentState, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	var resource PolicyAssignment
	err := ctx.ReadResource("azure-native:authorization/v20151001preview:PolicyAssignment", name, id, state, &resource, opts...)
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
	DisplayName          *string `pulumi:"displayName"`
	Id                   *string `pulumi:"id"`
	Name                 *string `pulumi:"name"`
	PolicyAssignmentName *string `pulumi:"policyAssignmentName"`
	PolicyDefinitionId   *string `pulumi:"policyDefinitionId"`
	Scope                string  `pulumi:"scope"`
	Type                 *string `pulumi:"type"`
}


type PolicyAssignmentArgs struct {
	DisplayName          pulumi.StringPtrInput
	Id                   pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	PolicyAssignmentName pulumi.StringPtrInput
	PolicyDefinitionId   pulumi.StringPtrInput
	Scope                pulumi.StringInput
	Type                 pulumi.StringPtrInput
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

func (o PolicyAssignmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyAssignmentOutput{})
}
