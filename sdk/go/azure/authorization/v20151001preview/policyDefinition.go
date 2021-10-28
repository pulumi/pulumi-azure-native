


package v20151001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyDefinition struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	Name        pulumi.StringPtrOutput `pulumi:"name"`
	PolicyRule  pulumi.AnyOutput       `pulumi:"policyRule"`
	PolicyType  pulumi.StringPtrOutput `pulumi:"policyType"`
}


func NewPolicyDefinition(ctx *pulumi.Context,
	name string, args *PolicyDefinitionArgs, opts ...pulumi.ResourceOption) (*PolicyDefinition, error) {
	if args == nil {
		args = &PolicyDefinitionArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:authorization/v20151001preview:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160401:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20160401:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20161201:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20161201:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20180301:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20180501:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20190101:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190601:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20190601:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190901:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20190901:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20200301:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20200901:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210601:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20210601:PolicyDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyDefinition
	err := ctx.RegisterResource("azure-native:authorization/v20151001preview:PolicyDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyDefinitionState, opts ...pulumi.ResourceOption) (*PolicyDefinition, error) {
	var resource PolicyDefinition
	err := ctx.ReadResource("azure-native:authorization/v20151001preview:PolicyDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyDefinitionState struct {
}

type PolicyDefinitionState struct {
}

func (PolicyDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionState)(nil)).Elem()
}

type policyDefinitionArgs struct {
	Description          *string     `pulumi:"description"`
	DisplayName          *string     `pulumi:"displayName"`
	Name                 *string     `pulumi:"name"`
	PolicyDefinitionName *string     `pulumi:"policyDefinitionName"`
	PolicyRule           interface{} `pulumi:"policyRule"`
	PolicyType           *string     `pulumi:"policyType"`
}


type PolicyDefinitionArgs struct {
	Description          pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	PolicyDefinitionName pulumi.StringPtrInput
	PolicyRule           pulumi.Input
	PolicyType           pulumi.StringPtrInput
}

func (PolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionArgs)(nil)).Elem()
}

type PolicyDefinitionInput interface {
	pulumi.Input

	ToPolicyDefinitionOutput() PolicyDefinitionOutput
	ToPolicyDefinitionOutputWithContext(ctx context.Context) PolicyDefinitionOutput
}

func (*PolicyDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinition)(nil))
}

func (i *PolicyDefinition) ToPolicyDefinitionOutput() PolicyDefinitionOutput {
	return i.ToPolicyDefinitionOutputWithContext(context.Background())
}

func (i *PolicyDefinition) ToPolicyDefinitionOutputWithContext(ctx context.Context) PolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionOutput)
}

type PolicyDefinitionOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinition)(nil))
}

func (o PolicyDefinitionOutput) ToPolicyDefinitionOutput() PolicyDefinitionOutput {
	return o
}

func (o PolicyDefinitionOutput) ToPolicyDefinitionOutputWithContext(ctx context.Context) PolicyDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyDefinitionInput)(nil)).Elem(), &PolicyDefinition{})
	pulumi.RegisterOutputType(PolicyDefinitionOutput{})
}
