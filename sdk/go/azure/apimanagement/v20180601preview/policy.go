


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Policy struct {
	pulumi.CustomResourceState

	ContentFormat pulumi.StringPtrOutput `pulumi:"contentFormat"`
	Name          pulumi.StringOutput    `pulumi:"name"`
	PolicyContent pulumi.StringOutput    `pulumi:"policyContent"`
	Type          pulumi.StringOutput    `pulumi:"type"`
}


func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyContent == nil {
		return nil, errors.New("invalid value for required argument 'PolicyContent'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.ContentFormat == nil {
		args.ContentFormat = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	ContentFormat     *string `pulumi:"contentFormat"`
	PolicyContent     string  `pulumi:"policyContent"`
	PolicyId          *string `pulumi:"policyId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type PolicyArgs struct {
	ContentFormat     pulumi.StringPtrInput
	PolicyContent     pulumi.StringInput
	PolicyId          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
}
