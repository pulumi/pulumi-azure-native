


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGovernanceAssignment(ctx *pulumi.Context, args *LookupGovernanceAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGovernanceAssignmentResult, error) {
	var rv LookupGovernanceAssignmentResult
	err := ctx.Invoke("azure-native:security/v20220101preview:getGovernanceAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGovernanceAssignmentArgs struct {
	AssessmentName string `pulumi:"assessmentName"`
	AssignmentKey  string `pulumi:"assignmentKey"`
	Scope          string `pulumi:"scope"`
}


type LookupGovernanceAssignmentResult struct {
	AdditionalData              *GovernanceAssignmentAdditionalDataResponse `pulumi:"additionalData"`
	GovernanceEmailNotification *GovernanceEmailNotificationResponse        `pulumi:"governanceEmailNotification"`
	Id                          string                                      `pulumi:"id"`
	IsGracePeriod               *bool                                       `pulumi:"isGracePeriod"`
	Name                        string                                      `pulumi:"name"`
	Owner                       *string                                     `pulumi:"owner"`
	RemediationDueDate          string                                      `pulumi:"remediationDueDate"`
	RemediationEta              *RemediationEtaResponse                     `pulumi:"remediationEta"`
	Type                        string                                      `pulumi:"type"`
}

func LookupGovernanceAssignmentOutput(ctx *pulumi.Context, args LookupGovernanceAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupGovernanceAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGovernanceAssignmentResult, error) {
			args := v.(LookupGovernanceAssignmentArgs)
			r, err := LookupGovernanceAssignment(ctx, &args, opts...)
			var s LookupGovernanceAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGovernanceAssignmentResultOutput)
}

type LookupGovernanceAssignmentOutputArgs struct {
	AssessmentName pulumi.StringInput `pulumi:"assessmentName"`
	AssignmentKey  pulumi.StringInput `pulumi:"assignmentKey"`
	Scope          pulumi.StringInput `pulumi:"scope"`
}

func (LookupGovernanceAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGovernanceAssignmentArgs)(nil)).Elem()
}


type LookupGovernanceAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupGovernanceAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGovernanceAssignmentResult)(nil)).Elem()
}

func (o LookupGovernanceAssignmentResultOutput) ToLookupGovernanceAssignmentResultOutput() LookupGovernanceAssignmentResultOutput {
	return o
}

func (o LookupGovernanceAssignmentResultOutput) ToLookupGovernanceAssignmentResultOutputWithContext(ctx context.Context) LookupGovernanceAssignmentResultOutput {
	return o
}

func (o LookupGovernanceAssignmentResultOutput) AdditionalData() GovernanceAssignmentAdditionalDataResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *GovernanceAssignmentAdditionalDataResponse {
		return v.AdditionalData
	}).(GovernanceAssignmentAdditionalDataResponsePtrOutput)
}

func (o LookupGovernanceAssignmentResultOutput) GovernanceEmailNotification() GovernanceEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *GovernanceEmailNotificationResponse {
		return v.GovernanceEmailNotification
	}).(GovernanceEmailNotificationResponsePtrOutput)
}

func (o LookupGovernanceAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGovernanceAssignmentResultOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *bool { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o LookupGovernanceAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGovernanceAssignmentResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o LookupGovernanceAssignmentResultOutput) RemediationDueDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) string { return v.RemediationDueDate }).(pulumi.StringOutput)
}

func (o LookupGovernanceAssignmentResultOutput) RemediationEta() RemediationEtaResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *RemediationEtaResponse { return v.RemediationEta }).(RemediationEtaResponsePtrOutput)
}

func (o LookupGovernanceAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGovernanceAssignmentResultOutput{})
}
