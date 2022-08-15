


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTopLevelDomainAgreements(ctx *pulumi.Context, args *ListTopLevelDomainAgreementsArgs, opts ...pulumi.InvokeOption) (*ListTopLevelDomainAgreementsResult, error) {
	var rv ListTopLevelDomainAgreementsResult
	err := ctx.Invoke("azure-native:domainregistration/v20220301:listTopLevelDomainAgreements", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTopLevelDomainAgreementsArgs struct {
	ForTransfer    *bool  `pulumi:"forTransfer"`
	IncludePrivacy *bool  `pulumi:"includePrivacy"`
	Name           string `pulumi:"name"`
}


type ListTopLevelDomainAgreementsResult struct {
	NextLink string                      `pulumi:"nextLink"`
	Value    []TldLegalAgreementResponse `pulumi:"value"`
}

func ListTopLevelDomainAgreementsOutput(ctx *pulumi.Context, args ListTopLevelDomainAgreementsOutputArgs, opts ...pulumi.InvokeOption) ListTopLevelDomainAgreementsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTopLevelDomainAgreementsResult, error) {
			args := v.(ListTopLevelDomainAgreementsArgs)
			r, err := ListTopLevelDomainAgreements(ctx, &args, opts...)
			var s ListTopLevelDomainAgreementsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTopLevelDomainAgreementsResultOutput)
}

type ListTopLevelDomainAgreementsOutputArgs struct {
	ForTransfer    pulumi.BoolPtrInput `pulumi:"forTransfer"`
	IncludePrivacy pulumi.BoolPtrInput `pulumi:"includePrivacy"`
	Name           pulumi.StringInput  `pulumi:"name"`
}

func (ListTopLevelDomainAgreementsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopLevelDomainAgreementsArgs)(nil)).Elem()
}


type ListTopLevelDomainAgreementsResultOutput struct{ *pulumi.OutputState }

func (ListTopLevelDomainAgreementsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopLevelDomainAgreementsResult)(nil)).Elem()
}

func (o ListTopLevelDomainAgreementsResultOutput) ToListTopLevelDomainAgreementsResultOutput() ListTopLevelDomainAgreementsResultOutput {
	return o
}

func (o ListTopLevelDomainAgreementsResultOutput) ToListTopLevelDomainAgreementsResultOutputWithContext(ctx context.Context) ListTopLevelDomainAgreementsResultOutput {
	return o
}

func (o ListTopLevelDomainAgreementsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopLevelDomainAgreementsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListTopLevelDomainAgreementsResultOutput) Value() TldLegalAgreementResponseArrayOutput {
	return o.ApplyT(func(v ListTopLevelDomainAgreementsResult) []TldLegalAgreementResponse { return v.Value }).(TldLegalAgreementResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTopLevelDomainAgreementsResultOutput{})
}
