


package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessment(ctx *pulumi.Context, args *LookupAssessmentArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentResult, error) {
	var rv LookupAssessmentResult
	err := ctx.Invoke("azure-native:migrate:getAssessment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentArgs struct {
	AssessmentName    string `pulumi:"assessmentName"`
	GroupName         string `pulumi:"groupName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAssessmentResult struct {
	ETag       *string                      `pulumi:"eTag"`
	Id         string                       `pulumi:"id"`
	Name       string                       `pulumi:"name"`
	Properties AssessmentPropertiesResponse `pulumi:"properties"`
	Type       string                       `pulumi:"type"`
}

func LookupAssessmentOutput(ctx *pulumi.Context, args LookupAssessmentOutputArgs, opts ...pulumi.InvokeOption) LookupAssessmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssessmentResult, error) {
			args := v.(LookupAssessmentArgs)
			r, err := LookupAssessment(ctx, &args, opts...)
			var s LookupAssessmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssessmentResultOutput)
}

type LookupAssessmentOutputArgs struct {
	AssessmentName    pulumi.StringInput `pulumi:"assessmentName"`
	GroupName         pulumi.StringInput `pulumi:"groupName"`
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAssessmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentArgs)(nil)).Elem()
}


type LookupAssessmentResultOutput struct{ *pulumi.OutputState }

func (LookupAssessmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentResult)(nil)).Elem()
}

func (o LookupAssessmentResultOutput) ToLookupAssessmentResultOutput() LookupAssessmentResultOutput {
	return o
}

func (o LookupAssessmentResultOutput) ToLookupAssessmentResultOutputWithContext(ctx context.Context) LookupAssessmentResultOutput {
	return o
}

func (o LookupAssessmentResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupAssessmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) Properties() AssessmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAssessmentResult) AssessmentPropertiesResponse { return v.Properties }).(AssessmentPropertiesResponseOutput)
}

func (o LookupAssessmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssessmentResultOutput{})
}
