


package v20200301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachineLearningCompute struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput    `pulumi:"location"`
	Name       pulumi.StringOutput       `pulumi:"name"`
	Properties pulumi.AnyOutput          `pulumi:"properties"`
	Sku        SkuResponsePtrOutput      `pulumi:"sku"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	Type       pulumi.StringOutput       `pulumi:"type"`
}


func NewMachineLearningCompute(ctx *pulumi.Context,
	name string, args *MachineLearningComputeArgs, opts ...pulumi.ResourceOption) (*MachineLearningCompute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200301:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20180301preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20180301preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20181119:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20181119:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190501:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20190501:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190601:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20190601:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20191101:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20191101:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200101:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200101:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200218preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200218preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200401:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200401:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200501preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200515preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200601:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200801:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200901preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210101:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210401:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210701:MachineLearningCompute"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineLearningCompute
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20200301:MachineLearningCompute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachineLearningCompute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineLearningComputeState, opts ...pulumi.ResourceOption) (*MachineLearningCompute, error) {
	var resource MachineLearningCompute
	err := ctx.ReadResource("azure-native:machinelearningservices/v20200301:MachineLearningCompute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineLearningComputeState struct {
}

type MachineLearningComputeState struct {
}

func (MachineLearningComputeState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningComputeState)(nil)).Elem()
}

type machineLearningComputeArgs struct {
	ComputeName       *string           `pulumi:"computeName"`
	Identity          *Identity         `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	WorkspaceName     string            `pulumi:"workspaceName"`
}


type MachineLearningComputeArgs struct {
	ComputeName       pulumi.StringPtrInput
	Identity          IdentityPtrInput
	Location          pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
}

func (MachineLearningComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningComputeArgs)(nil)).Elem()
}

type MachineLearningComputeInput interface {
	pulumi.Input

	ToMachineLearningComputeOutput() MachineLearningComputeOutput
	ToMachineLearningComputeOutputWithContext(ctx context.Context) MachineLearningComputeOutput
}

func (*MachineLearningCompute) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningCompute)(nil))
}

func (i *MachineLearningCompute) ToMachineLearningComputeOutput() MachineLearningComputeOutput {
	return i.ToMachineLearningComputeOutputWithContext(context.Background())
}

func (i *MachineLearningCompute) ToMachineLearningComputeOutputWithContext(ctx context.Context) MachineLearningComputeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningComputeOutput)
}

type MachineLearningComputeOutput struct{ *pulumi.OutputState }

func (MachineLearningComputeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningCompute)(nil))
}

func (o MachineLearningComputeOutput) ToMachineLearningComputeOutput() MachineLearningComputeOutput {
	return o
}

func (o MachineLearningComputeOutput) ToMachineLearningComputeOutputWithContext(ctx context.Context) MachineLearningComputeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MachineLearningComputeOutput{})
}
