


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessment(ctx *pulumi.Context, args *LookupAssessmentArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentResult, error) {
	var rv LookupAssessmentResult
	err := ctx.Invoke("azure-native:security/v20200101:getAssessment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentArgs struct {
	AssessmentName string  `pulumi:"assessmentName"`
	Expand         *string `pulumi:"expand"`
	ResourceId     string  `pulumi:"resourceId"`
}


type LookupAssessmentResult struct {
	AdditionalData  map[string]string                             `pulumi:"additionalData"`
	DisplayName     string                                        `pulumi:"displayName"`
	Id              string                                        `pulumi:"id"`
	Links           AssessmentLinksResponse                       `pulumi:"links"`
	Metadata        *SecurityAssessmentMetadataPropertiesResponse `pulumi:"metadata"`
	Name            string                                        `pulumi:"name"`
	PartnersData    *SecurityAssessmentPartnerDataResponse        `pulumi:"partnersData"`
	ResourceDetails interface{}                                   `pulumi:"resourceDetails"`
	Status          AssessmentStatusResponse                      `pulumi:"status"`
	Type            string                                        `pulumi:"type"`
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
	AssessmentName pulumi.StringInput    `pulumi:"assessmentName"`
	Expand         pulumi.StringPtrInput `pulumi:"expand"`
	ResourceId     pulumi.StringInput    `pulumi:"resourceId"`
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

func (o LookupAssessmentResultOutput) AdditionalData() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAssessmentResult) map[string]string { return v.AdditionalData }).(pulumi.StringMapOutput)
}

func (o LookupAssessmentResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) Links() AssessmentLinksResponseOutput {
	return o.ApplyT(func(v LookupAssessmentResult) AssessmentLinksResponse { return v.Links }).(AssessmentLinksResponseOutput)
}

func (o LookupAssessmentResultOutput) Metadata() SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupAssessmentResult) *SecurityAssessmentMetadataPropertiesResponse { return v.Metadata }).(SecurityAssessmentMetadataPropertiesResponsePtrOutput)
}

func (o LookupAssessmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) PartnersData() SecurityAssessmentPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v LookupAssessmentResult) *SecurityAssessmentPartnerDataResponse { return v.PartnersData }).(SecurityAssessmentPartnerDataResponsePtrOutput)
}

func (o LookupAssessmentResultOutput) ResourceDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAssessmentResult) interface{} { return v.ResourceDetails }).(pulumi.AnyOutput)
}

func (o LookupAssessmentResultOutput) Status() AssessmentStatusResponseOutput {
	return o.ApplyT(func(v LookupAssessmentResult) AssessmentStatusResponse { return v.Status }).(AssessmentStatusResponseOutput)
}

func (o LookupAssessmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssessmentResultOutput{})
}
