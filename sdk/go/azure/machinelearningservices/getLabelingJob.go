


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLabelingJob(ctx *pulumi.Context, args *LookupLabelingJobArgs, opts ...pulumi.InvokeOption) (*LookupLabelingJobResult, error) {
	var rv LookupLabelingJobResult
	err := ctx.Invoke("azure-native:machinelearningservices:getLabelingJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabelingJobArgs struct {
	IncludeJobInstructions *bool  `pulumi:"includeJobInstructions"`
	IncludeLabelCategories *bool  `pulumi:"includeLabelCategories"`
	LabelingJobId          string `pulumi:"labelingJobId"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type LookupLabelingJobResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties LabelingJobPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}

func LookupLabelingJobOutput(ctx *pulumi.Context, args LookupLabelingJobOutputArgs, opts ...pulumi.InvokeOption) LookupLabelingJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabelingJobResult, error) {
			args := v.(LookupLabelingJobArgs)
			r, err := LookupLabelingJob(ctx, &args, opts...)
			var s LookupLabelingJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabelingJobResultOutput)
}

type LookupLabelingJobOutputArgs struct {
	IncludeJobInstructions pulumi.BoolPtrInput `pulumi:"includeJobInstructions"`
	IncludeLabelCategories pulumi.BoolPtrInput `pulumi:"includeLabelCategories"`
	LabelingJobId          pulumi.StringInput  `pulumi:"labelingJobId"`
	ResourceGroupName      pulumi.StringInput  `pulumi:"resourceGroupName"`
	WorkspaceName          pulumi.StringInput  `pulumi:"workspaceName"`
}

func (LookupLabelingJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabelingJobArgs)(nil)).Elem()
}


type LookupLabelingJobResultOutput struct{ *pulumi.OutputState }

func (LookupLabelingJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabelingJobResult)(nil)).Elem()
}

func (o LookupLabelingJobResultOutput) ToLookupLabelingJobResultOutput() LookupLabelingJobResultOutput {
	return o
}

func (o LookupLabelingJobResultOutput) ToLookupLabelingJobResultOutputWithContext(ctx context.Context) LookupLabelingJobResultOutput {
	return o
}

func (o LookupLabelingJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabelingJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabelingJobResultOutput) Properties() LabelingJobPropertiesResponseOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) LabelingJobPropertiesResponse { return v.Properties }).(LabelingJobPropertiesResponseOutput)
}

func (o LookupLabelingJobResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLabelingJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabelingJobResultOutput{})
}
