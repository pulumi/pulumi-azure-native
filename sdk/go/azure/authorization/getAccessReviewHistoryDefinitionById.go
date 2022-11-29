


package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessReviewHistoryDefinitionById(ctx *pulumi.Context, args *LookupAccessReviewHistoryDefinitionByIdArgs, opts ...pulumi.InvokeOption) (*LookupAccessReviewHistoryDefinitionByIdResult, error) {
	var rv LookupAccessReviewHistoryDefinitionByIdResult
	err := ctx.Invoke("azure-native:authorization:getAccessReviewHistoryDefinitionById", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessReviewHistoryDefinitionByIdArgs struct {
	HistoryDefinitionId string `pulumi:"historyDefinitionId"`
}


type LookupAccessReviewHistoryDefinitionByIdResult struct {
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

func LookupAccessReviewHistoryDefinitionByIdOutput(ctx *pulumi.Context, args LookupAccessReviewHistoryDefinitionByIdOutputArgs, opts ...pulumi.InvokeOption) LookupAccessReviewHistoryDefinitionByIdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessReviewHistoryDefinitionByIdResult, error) {
			args := v.(LookupAccessReviewHistoryDefinitionByIdArgs)
			r, err := LookupAccessReviewHistoryDefinitionById(ctx, &args, opts...)
			var s LookupAccessReviewHistoryDefinitionByIdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessReviewHistoryDefinitionByIdResultOutput)
}

type LookupAccessReviewHistoryDefinitionByIdOutputArgs struct {
	HistoryDefinitionId pulumi.StringInput `pulumi:"historyDefinitionId"`
}

func (LookupAccessReviewHistoryDefinitionByIdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessReviewHistoryDefinitionByIdArgs)(nil)).Elem()
}


type LookupAccessReviewHistoryDefinitionByIdResultOutput struct{ *pulumi.OutputState }

func (LookupAccessReviewHistoryDefinitionByIdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessReviewHistoryDefinitionByIdResult)(nil)).Elem()
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) ToLookupAccessReviewHistoryDefinitionByIdResultOutput() LookupAccessReviewHistoryDefinitionByIdResultOutput {
	return o
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) ToLookupAccessReviewHistoryDefinitionByIdResultOutputWithContext(ctx context.Context) LookupAccessReviewHistoryDefinitionByIdResultOutput {
	return o
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) Decisions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) []string { return v.Decisions }).(pulumi.StringArrayOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) Instances() AccessReviewHistoryInstanceResponseArrayOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) []AccessReviewHistoryInstanceResponse {
		return v.Instances
	}).(AccessReviewHistoryInstanceResponseArrayOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) *int { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) ReviewHistoryPeriodEndDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.ReviewHistoryPeriodEndDateTime }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) ReviewHistoryPeriodStartDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string {
		return v.ReviewHistoryPeriodStartDateTime
	}).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) Scopes() AccessReviewScopeResponseArrayOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) []AccessReviewScopeResponse { return v.Scopes }).(AccessReviewScopeResponseArrayOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAccessReviewHistoryDefinitionByIdResultOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewHistoryDefinitionByIdResult) string { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessReviewHistoryDefinitionByIdResultOutput{})
}
