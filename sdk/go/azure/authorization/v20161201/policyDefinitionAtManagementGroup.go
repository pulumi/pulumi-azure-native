


package v20161201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyDefinitionAtManagementGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	Metadata    pulumi.AnyOutput       `pulumi:"metadata"`
	Mode        pulumi.StringPtrOutput `pulumi:"mode"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Parameters  pulumi.AnyOutput       `pulumi:"parameters"`
	PolicyRule  pulumi.AnyOutput       `pulumi:"policyRule"`
	PolicyType  pulumi.StringPtrOutput `pulumi:"policyType"`
}


func NewPolicyDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, args *PolicyDefinitionAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*PolicyDefinitionAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:authorization/v20161201:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20180301:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20180501:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20190101:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190601:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20190601:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190901:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20190901:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20200301:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20200901:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210601:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20210601:PolicyDefinitionAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyDefinitionAtManagementGroup
	err := ctx.RegisterResource("azure-native:authorization/v20161201:PolicyDefinitionAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyDefinitionAtManagementGroupState, opts ...pulumi.ResourceOption) (*PolicyDefinitionAtManagementGroup, error) {
	var resource PolicyDefinitionAtManagementGroup
	err := ctx.ReadResource("azure-native:authorization/v20161201:PolicyDefinitionAtManagementGroup", name, id, state, &resource, opts...)
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
	Description          *string     `pulumi:"description"`
	DisplayName          *string     `pulumi:"displayName"`
	ManagementGroupId    string      `pulumi:"managementGroupId"`
	Metadata             interface{} `pulumi:"metadata"`
	Mode                 *string     `pulumi:"mode"`
	Parameters           interface{} `pulumi:"parameters"`
	PolicyDefinitionName *string     `pulumi:"policyDefinitionName"`
	PolicyRule           interface{} `pulumi:"policyRule"`
	PolicyType           *string     `pulumi:"policyType"`
}


type PolicyDefinitionAtManagementGroupArgs struct {
	Description          pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	ManagementGroupId    pulumi.StringInput
	Metadata             pulumi.Input
	Mode                 pulumi.StringPtrInput
	Parameters           pulumi.Input
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
	return reflect.TypeOf((*PolicyDefinitionAtManagementGroup)(nil))
}

func (i *PolicyDefinitionAtManagementGroup) ToPolicyDefinitionAtManagementGroupOutput() PolicyDefinitionAtManagementGroupOutput {
	return i.ToPolicyDefinitionAtManagementGroupOutputWithContext(context.Background())
}

func (i *PolicyDefinitionAtManagementGroup) ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicyDefinitionAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionAtManagementGroupOutput)
}

type PolicyDefinitionAtManagementGroupOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionAtManagementGroup)(nil))
}

func (o PolicyDefinitionAtManagementGroupOutput) ToPolicyDefinitionAtManagementGroupOutput() PolicyDefinitionAtManagementGroupOutput {
	return o
}

func (o PolicyDefinitionAtManagementGroupOutput) ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicyDefinitionAtManagementGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyDefinitionAtManagementGroupOutput{})
}
