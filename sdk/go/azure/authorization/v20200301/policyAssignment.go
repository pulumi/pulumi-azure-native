


package v20200301

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
			Type: pulumi.String("azure-native:authorization/v20190901:PolicyAssignment"),
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
	err := ctx.RegisterResource("azure-native:authorization/v20200301:PolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAssignmentState, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	var resource PolicyAssignment
	err := ctx.ReadResource("azure-native:authorization/v20200301:PolicyAssignment", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*PolicyAssignment)(nil))
}

func (i *PolicyAssignment) ToPolicyAssignmentOutput() PolicyAssignmentOutput {
	return i.ToPolicyAssignmentOutputWithContext(context.Background())
}

func (i *PolicyAssignment) ToPolicyAssignmentOutputWithContext(ctx context.Context) PolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentOutput)
}

type PolicyAssignmentOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignment)(nil))
}

func (o PolicyAssignmentOutput) ToPolicyAssignmentOutput() PolicyAssignmentOutput {
	return o
}

func (o PolicyAssignmentOutput) ToPolicyAssignmentOutputWithContext(ctx context.Context) PolicyAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyAssignmentOutput{})
}
