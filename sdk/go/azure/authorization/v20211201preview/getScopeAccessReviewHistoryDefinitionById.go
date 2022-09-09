


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScopeAccessReviewHistoryDefinitionById(ctx *pulumi.Context, args *LookupScopeAccessReviewHistoryDefinitionByIdArgs, opts ...pulumi.InvokeOption) (*LookupScopeAccessReviewHistoryDefinitionByIdResult, error) {
	var rv LookupScopeAccessReviewHistoryDefinitionByIdResult
	err := ctx.Invoke("azure-native:authorization/v20211201preview:getScopeAccessReviewHistoryDefinitionById", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeAccessReviewHistoryDefinitionByIdArgs struct {
	HistoryDefinitionId string `pulumi:"historyDefinitionId"`
	Scope               string `pulumi:"scope"`
}


type LookupScopeAccessReviewHistoryDefinitionByIdResult struct {
	CreatedDateTime                  string                                `pulumi:"createdDateTime"`
	Decisions                        []string                              `pulumi:"decisions"`
	DisplayName                      *string                               `pulumi:"displayName"`
	EndDate                          *string                               `pulumi:"endDate"`
	Id                               string                                `pulumi:"id"`
	Instances                        []AccessReviewHistoryInstanceResponse `pulumi:"instances"`
	Interval                         *int                                  `pulumi:"interval"`
	Name                             string                                `pulumi:"name"`
	NumberOfOccurrences              *int                                  `pulumi:"numberOfOccurrences"`
	PrincipalId                      string                                `pulumi:"principalId"`
	PrincipalName                    string                                `pulumi:"principalName"`
	PrincipalType                    string                                `pulumi:"principalType"`
	ReviewHistoryPeriodEndDateTime   string                                `pulumi:"reviewHistoryPeriodEndDateTime"`
	ReviewHistoryPeriodStartDateTime string                                `pulumi:"reviewHistoryPeriodStartDateTime"`
	Scopes                           []AccessReviewScopeResponse           `pulumi:"scopes"`
	StartDate                        *string                               `pulumi:"startDate"`
	Status                           string                                `pulumi:"status"`
	Type                             string                                `pulumi:"type"`
	UserPrincipalName                string                                `pulumi:"userPrincipalName"`
}

func LookupScopeAccessReviewHistoryDefinitionByIdOutput(ctx *pulumi.Context, args LookupScopeAccessReviewHistoryDefinitionByIdOutputArgs, opts ...pulumi.InvokeOption) LookupScopeAccessReviewHistoryDefinitionByIdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScopeAccessReviewHistoryDefinitionByIdResult, error) {
			args := v.(LookupScopeAccessReviewHistoryDefinitionByIdArgs)
			r, err := LookupScopeAccessReviewHistoryDefinitionById(ctx, &args, opts...)
			var s LookupScopeAccessReviewHistoryDefinitionByIdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScopeAccessReviewHistoryDefinitionByIdResultOutput)
}

type LookupScopeAccessReviewHistoryDefinitionByIdOutputArgs struct {
	HistoryDefinitionId pulumi.StringInput `pulumi:"historyDefinitionId"`
	Scope               pulumi.StringInput `pulumi:"scope"`
}

func (LookupScopeAccessReviewHistoryDefinitionByIdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeAccessReviewHistoryDefinitionByIdArgs)(nil)).Elem()
}


type LookupScopeAccessReviewHistoryDefinitionByIdResultOutput struct{ *pulumi.OutputState }

func (LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeAccessReviewHistoryDefinitionByIdResult)(nil)).Elem()
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) ToLookupScopeAccessReviewHistoryDefinitionByIdResultOutput() LookupScopeAccessReviewHistoryDefinitionByIdResultOutput {
	return o
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) ToLookupScopeAccessReviewHistoryDefinitionByIdResultOutputWithContext(ctx context.Context) LookupScopeAccessReviewHistoryDefinitionByIdResultOutput {
	return o
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) Decisions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) []string { return v.Decisions }).(pulumi.StringArrayOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) Instances() AccessReviewHistoryInstanceResponseArrayOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) []AccessReviewHistoryInstanceResponse {
		return v.Instances
	}).(AccessReviewHistoryInstanceResponseArrayOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) *int { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) ReviewHistoryPeriodEndDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string {
		return v.ReviewHistoryPeriodEndDateTime
	}).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) ReviewHistoryPeriodStartDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string {
		return v.ReviewHistoryPeriodStartDateTime
	}).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) Scopes() AccessReviewScopeResponseArrayOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) []AccessReviewScopeResponse {
		return v.Scopes
	}).(AccessReviewScopeResponseArrayOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewHistoryDefinitionByIdResultOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewHistoryDefinitionByIdResult) string { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScopeAccessReviewHistoryDefinitionByIdResultOutput{})
}
