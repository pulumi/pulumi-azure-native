


package avs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PlacementPolicy struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewPlacementPolicy(ctx *pulumi.Context,
	name string, args *PlacementPolicyArgs, opts ...pulumi.ResourceOption) (*PlacementPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:avs:PlacementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:PlacementPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20211201:PlacementPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource PlacementPolicy
	err := ctx.RegisterResource("azure-native:avs:PlacementPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPlacementPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlacementPolicyState, opts ...pulumi.ResourceOption) (*PlacementPolicy, error) {
	var resource PlacementPolicy
	err := ctx.ReadResource("azure-native:avs:PlacementPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type placementPolicyState struct {
}

type PlacementPolicyState struct {
}

func (PlacementPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*placementPolicyState)(nil)).Elem()
}

type placementPolicyArgs struct {
	ClusterName         string      `pulumi:"clusterName"`
	PlacementPolicyName *string     `pulumi:"placementPolicyName"`
	PrivateCloudName    string      `pulumi:"privateCloudName"`
	Properties          interface{} `pulumi:"properties"`
	ResourceGroupName   string      `pulumi:"resourceGroupName"`
}


type PlacementPolicyArgs struct {
	ClusterName         pulumi.StringInput
	PlacementPolicyName pulumi.StringPtrInput
	PrivateCloudName    pulumi.StringInput
	Properties          pulumi.Input
	ResourceGroupName   pulumi.StringInput
}

func (PlacementPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placementPolicyArgs)(nil)).Elem()
}

type PlacementPolicyInput interface {
	pulumi.Input

	ToPlacementPolicyOutput() PlacementPolicyOutput
	ToPlacementPolicyOutputWithContext(ctx context.Context) PlacementPolicyOutput
}

func (*PlacementPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementPolicy)(nil))
}

func (i *PlacementPolicy) ToPlacementPolicyOutput() PlacementPolicyOutput {
	return i.ToPlacementPolicyOutputWithContext(context.Background())
}

func (i *PlacementPolicy) ToPlacementPolicyOutputWithContext(ctx context.Context) PlacementPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementPolicyOutput)
}

type PlacementPolicyOutput struct{ *pulumi.OutputState }

func (PlacementPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementPolicy)(nil))
}

func (o PlacementPolicyOutput) ToPlacementPolicyOutput() PlacementPolicyOutput {
	return o
}

func (o PlacementPolicyOutput) ToPlacementPolicyOutputWithContext(ctx context.Context) PlacementPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlacementPolicyInput)(nil)).Elem(), &PlacementPolicy{})
	pulumi.RegisterOutputType(PlacementPolicyOutput{})
}
