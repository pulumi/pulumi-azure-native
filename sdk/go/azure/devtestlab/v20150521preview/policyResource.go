


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyResource struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput `pulumi:"description"`
	EvaluatorType     pulumi.StringPtrOutput `pulumi:"evaluatorType"`
	FactData          pulumi.StringPtrOutput `pulumi:"factData"`
	FactName          pulumi.StringPtrOutput `pulumi:"factName"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringPtrOutput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	Status            pulumi.StringPtrOutput `pulumi:"status"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Threshold         pulumi.StringPtrOutput `pulumi:"threshold"`
	Type              pulumi.StringPtrOutput `pulumi:"type"`
}


func NewPolicyResource(ctx *pulumi.Context,
	name string, args *PolicyResourceArgs, opts ...pulumi.ResourceOption) (*PolicyResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.PolicySetName == nil {
		return nil, errors.New("invalid value for required argument 'PolicySetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:PolicyResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:PolicyResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:PolicyResource"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:PolicyResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyResourceState, opts ...pulumi.ResourceOption) (*PolicyResource, error) {
	var resource PolicyResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:PolicyResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyResourceState struct {
}

type PolicyResourceState struct {
}

func (PolicyResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyResourceState)(nil)).Elem()
}

type policyResourceArgs struct {
	Description       *string           `pulumi:"description"`
	EvaluatorType     *string           `pulumi:"evaluatorType"`
	FactData          *string           `pulumi:"factData"`
	FactName          *string           `pulumi:"factName"`
	Id                *string           `pulumi:"id"`
	LabName           string            `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	PolicySetName     string            `pulumi:"policySetName"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Threshold         *string           `pulumi:"threshold"`
	Type              *string           `pulumi:"type"`
}


type PolicyResourceArgs struct {
	Description       pulumi.StringPtrInput
	EvaluatorType     pulumi.StringPtrInput
	FactData          pulumi.StringPtrInput
	FactName          pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	PolicySetName     pulumi.StringInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Status            pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Threshold         pulumi.StringPtrInput
	Type              pulumi.StringPtrInput
}

func (PolicyResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyResourceArgs)(nil)).Elem()
}

type PolicyResourceInput interface {
	pulumi.Input

	ToPolicyResourceOutput() PolicyResourceOutput
	ToPolicyResourceOutputWithContext(ctx context.Context) PolicyResourceOutput
}

func (*PolicyResource) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyResource)(nil)).Elem()
}

func (i *PolicyResource) ToPolicyResourceOutput() PolicyResourceOutput {
	return i.ToPolicyResourceOutputWithContext(context.Background())
}

func (i *PolicyResource) ToPolicyResourceOutputWithContext(ctx context.Context) PolicyResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyResourceOutput)
}

type PolicyResourceOutput struct{ *pulumi.OutputState }

func (PolicyResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyResource)(nil)).Elem()
}

func (o PolicyResourceOutput) ToPolicyResourceOutput() PolicyResourceOutput {
	return o
}

func (o PolicyResourceOutput) ToPolicyResourceOutputWithContext(ctx context.Context) PolicyResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyResourceOutput{})
}
