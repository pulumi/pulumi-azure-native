


package datalakeanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputePolicy struct {
	pulumi.CustomResourceState

	MaxDegreeOfParallelismPerJob pulumi.IntOutput    `pulumi:"maxDegreeOfParallelismPerJob"`
	MinPriorityPerJob            pulumi.IntOutput    `pulumi:"minPriorityPerJob"`
	Name                         pulumi.StringOutput `pulumi:"name"`
	ObjectId                     pulumi.StringOutput `pulumi:"objectId"`
	ObjectType                   pulumi.StringOutput `pulumi:"objectType"`
	Type                         pulumi.StringOutput `pulumi:"type"`
}


func NewComputePolicy(ctx *pulumi.Context,
	name string, args *ComputePolicyArgs, opts ...pulumi.ResourceOption) (*ComputePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ObjectId'")
	}
	if args.ObjectType == nil {
		return nil, errors.New("invalid value for required argument 'ObjectType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics:ComputePolicy"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20151001preview:ComputePolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics/v20151001preview:ComputePolicy"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20161101:ComputePolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics/v20161101:ComputePolicy"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20191101preview:ComputePolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics/v20191101preview:ComputePolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ComputePolicy
	err := ctx.RegisterResource("azure-native:datalakeanalytics:ComputePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetComputePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputePolicyState, opts ...pulumi.ResourceOption) (*ComputePolicy, error) {
	var resource ComputePolicy
	err := ctx.ReadResource("azure-native:datalakeanalytics:ComputePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type computePolicyState struct {
}

type ComputePolicyState struct {
}

func (ComputePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computePolicyState)(nil)).Elem()
}

type computePolicyArgs struct {
	AccountName                  string  `pulumi:"accountName"`
	ComputePolicyName            *string `pulumi:"computePolicyName"`
	MaxDegreeOfParallelismPerJob *int    `pulumi:"maxDegreeOfParallelismPerJob"`
	MinPriorityPerJob            *int    `pulumi:"minPriorityPerJob"`
	ObjectId                     string  `pulumi:"objectId"`
	ObjectType                   string  `pulumi:"objectType"`
	ResourceGroupName            string  `pulumi:"resourceGroupName"`
}


type ComputePolicyArgs struct {
	AccountName                  pulumi.StringInput
	ComputePolicyName            pulumi.StringPtrInput
	MaxDegreeOfParallelismPerJob pulumi.IntPtrInput
	MinPriorityPerJob            pulumi.IntPtrInput
	ObjectId                     pulumi.StringInput
	ObjectType                   pulumi.StringInput
	ResourceGroupName            pulumi.StringInput
}

func (ComputePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computePolicyArgs)(nil)).Elem()
}

type ComputePolicyInput interface {
	pulumi.Input

	ToComputePolicyOutput() ComputePolicyOutput
	ToComputePolicyOutputWithContext(ctx context.Context) ComputePolicyOutput
}

func (*ComputePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputePolicy)(nil))
}

func (i *ComputePolicy) ToComputePolicyOutput() ComputePolicyOutput {
	return i.ToComputePolicyOutputWithContext(context.Background())
}

func (i *ComputePolicy) ToComputePolicyOutputWithContext(ctx context.Context) ComputePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputePolicyOutput)
}

type ComputePolicyOutput struct{ *pulumi.OutputState }

func (ComputePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputePolicy)(nil))
}

func (o ComputePolicyOutput) ToComputePolicyOutput() ComputePolicyOutput {
	return o
}

func (o ComputePolicyOutput) ToComputePolicyOutputWithContext(ctx context.Context) ComputePolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputePolicyInput)(nil)).Elem(), &ComputePolicy{})
	pulumi.RegisterOutputType(ComputePolicyOutput{})
}
