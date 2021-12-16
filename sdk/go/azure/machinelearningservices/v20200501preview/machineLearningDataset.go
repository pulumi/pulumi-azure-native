


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachineLearningDataset struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput    `pulumi:"location"`
	Name       pulumi.StringOutput       `pulumi:"name"`
	Properties DatasetResponseOutput     `pulumi:"properties"`
	Sku        SkuResponsePtrOutput      `pulumi:"sku"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	Type       pulumi.StringOutput       `pulumi:"type"`
}


func NewMachineLearningDataset(ctx *pulumi.Context,
	name string, args *MachineLearningDatasetArgs, opts ...pulumi.ResourceOption) (*MachineLearningDataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetType == nil {
		return nil, errors.New("invalid value for required argument 'DatasetType'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.Registration == nil {
		return nil, errors.New("invalid value for required argument 'Registration'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	parametersApplier := func(v DatasetCreateRequestParameters) *DatasetCreateRequestParameters { return v.Defaults() }
	args.Parameters = args.Parameters.ToDatasetCreateRequestParametersOutput().ApplyT(parametersApplier).(DatasetCreateRequestParametersPtrOutput).Elem()
	if isZero(args.SkipValidation) {
		args.SkipValidation = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:MachineLearningDataset"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineLearningDataset
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20200501preview:MachineLearningDataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachineLearningDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineLearningDatasetState, opts ...pulumi.ResourceOption) (*MachineLearningDataset, error) {
	var resource MachineLearningDataset
	err := ctx.ReadResource("azure-native:machinelearningservices/v20200501preview:MachineLearningDataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineLearningDatasetState struct {
}

type MachineLearningDatasetState struct {
}

func (MachineLearningDatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningDatasetState)(nil)).Elem()
}

type machineLearningDatasetArgs struct {
	DatasetName       *string                          `pulumi:"datasetName"`
	DatasetType       string                           `pulumi:"datasetType"`
	Parameters        DatasetCreateRequestParameters   `pulumi:"parameters"`
	Registration      DatasetCreateRequestRegistration `pulumi:"registration"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
	SkipValidation    *bool                            `pulumi:"skipValidation"`
	TimeSeries        *DatasetCreateRequestTimeSeries  `pulumi:"timeSeries"`
	WorkspaceName     string                           `pulumi:"workspaceName"`
}


type MachineLearningDatasetArgs struct {
	DatasetName       pulumi.StringPtrInput
	DatasetType       pulumi.StringInput
	Parameters        DatasetCreateRequestParametersInput
	Registration      DatasetCreateRequestRegistrationInput
	ResourceGroupName pulumi.StringInput
	SkipValidation    pulumi.BoolPtrInput
	TimeSeries        DatasetCreateRequestTimeSeriesPtrInput
	WorkspaceName     pulumi.StringInput
}

func (MachineLearningDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningDatasetArgs)(nil)).Elem()
}

type MachineLearningDatasetInput interface {
	pulumi.Input

	ToMachineLearningDatasetOutput() MachineLearningDatasetOutput
	ToMachineLearningDatasetOutputWithContext(ctx context.Context) MachineLearningDatasetOutput
}

func (*MachineLearningDataset) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningDataset)(nil)).Elem()
}

func (i *MachineLearningDataset) ToMachineLearningDatasetOutput() MachineLearningDatasetOutput {
	return i.ToMachineLearningDatasetOutputWithContext(context.Background())
}

func (i *MachineLearningDataset) ToMachineLearningDatasetOutputWithContext(ctx context.Context) MachineLearningDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningDatasetOutput)
}

type MachineLearningDatasetOutput struct{ *pulumi.OutputState }

func (MachineLearningDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningDataset)(nil)).Elem()
}

func (o MachineLearningDatasetOutput) ToMachineLearningDatasetOutput() MachineLearningDatasetOutput {
	return o
}

func (o MachineLearningDatasetOutput) ToMachineLearningDatasetOutputWithContext(ctx context.Context) MachineLearningDatasetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MachineLearningDatasetOutput{})
}
