


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineLearningDataset(ctx *pulumi.Context, args *LookupMachineLearningDatasetArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningDatasetResult, error) {
	var rv LookupMachineLearningDatasetResult
	err := ctx.Invoke("azure-native:machinelearningservices:getMachineLearningDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineLearningDatasetArgs struct {
	DatasetName       string `pulumi:"datasetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMachineLearningDatasetResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties DatasetResponse   `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupMachineLearningDatasetOutput(ctx *pulumi.Context, args LookupMachineLearningDatasetOutputArgs, opts ...pulumi.InvokeOption) LookupMachineLearningDatasetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineLearningDatasetResult, error) {
			args := v.(LookupMachineLearningDatasetArgs)
			r, err := LookupMachineLearningDataset(ctx, &args, opts...)
			var s LookupMachineLearningDatasetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineLearningDatasetResultOutput)
}

type LookupMachineLearningDatasetOutputArgs struct {
	DatasetName       pulumi.StringInput `pulumi:"datasetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMachineLearningDatasetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningDatasetArgs)(nil)).Elem()
}


type LookupMachineLearningDatasetResultOutput struct{ *pulumi.OutputState }

func (LookupMachineLearningDatasetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningDatasetResult)(nil)).Elem()
}

func (o LookupMachineLearningDatasetResultOutput) ToLookupMachineLearningDatasetResultOutput() LookupMachineLearningDatasetResultOutput {
	return o
}

func (o LookupMachineLearningDatasetResultOutput) ToLookupMachineLearningDatasetResultOutputWithContext(ctx context.Context) LookupMachineLearningDatasetResultOutput {
	return o
}

func (o LookupMachineLearningDatasetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatasetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineLearningDatasetResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatasetResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupMachineLearningDatasetResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatasetResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMachineLearningDatasetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatasetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineLearningDatasetResultOutput) Properties() DatasetResponseOutput {
	return o.ApplyT(func(v LookupMachineLearningDatasetResult) DatasetResponse { return v.Properties }).(DatasetResponseOutput)
}

func (o LookupMachineLearningDatasetResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatasetResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupMachineLearningDatasetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineLearningDatasetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMachineLearningDatasetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatasetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineLearningDatasetResultOutput{})
}
