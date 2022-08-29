


package deploymentmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRollout(ctx *pulumi.Context, args *LookupRolloutArgs, opts ...pulumi.InvokeOption) (*LookupRolloutResult, error) {
	var rv LookupRolloutResult
	err := ctx.Invoke("azure-native:deploymentmanager:getRollout", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRolloutArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RetryAttempt      *int   `pulumi:"retryAttempt"`
	RolloutName       string `pulumi:"rolloutName"`
}


type LookupRolloutResult struct {
	ArtifactSourceId        *string                      `pulumi:"artifactSourceId"`
	BuildVersion            string                       `pulumi:"buildVersion"`
	Id                      string                       `pulumi:"id"`
	Identity                *IdentityResponse            `pulumi:"identity"`
	Location                string                       `pulumi:"location"`
	Name                    string                       `pulumi:"name"`
	OperationInfo           RolloutOperationInfoResponse `pulumi:"operationInfo"`
	Services                []ServiceResponse            `pulumi:"services"`
	Status                  string                       `pulumi:"status"`
	StepGroups              []StepGroupResponse          `pulumi:"stepGroups"`
	Tags                    map[string]string            `pulumi:"tags"`
	TargetServiceTopologyId string                       `pulumi:"targetServiceTopologyId"`
	TotalRetryAttempts      int                          `pulumi:"totalRetryAttempts"`
	Type                    string                       `pulumi:"type"`
}

func LookupRolloutOutput(ctx *pulumi.Context, args LookupRolloutOutputArgs, opts ...pulumi.InvokeOption) LookupRolloutResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRolloutResult, error) {
			args := v.(LookupRolloutArgs)
			r, err := LookupRollout(ctx, &args, opts...)
			var s LookupRolloutResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRolloutResultOutput)
}

type LookupRolloutOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RetryAttempt      pulumi.IntPtrInput `pulumi:"retryAttempt"`
	RolloutName       pulumi.StringInput `pulumi:"rolloutName"`
}

func (LookupRolloutOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRolloutArgs)(nil)).Elem()
}


type LookupRolloutResultOutput struct{ *pulumi.OutputState }

func (LookupRolloutResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRolloutResult)(nil)).Elem()
}

func (o LookupRolloutResultOutput) ToLookupRolloutResultOutput() LookupRolloutResultOutput {
	return o
}

func (o LookupRolloutResultOutput) ToLookupRolloutResultOutputWithContext(ctx context.Context) LookupRolloutResultOutput {
	return o
}

func (o LookupRolloutResultOutput) ArtifactSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRolloutResult) *string { return v.ArtifactSourceId }).(pulumi.StringPtrOutput)
}

func (o LookupRolloutResultOutput) BuildVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.BuildVersion }).(pulumi.StringOutput)
}

func (o LookupRolloutResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRolloutResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupRolloutResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupRolloutResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRolloutResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRolloutResultOutput) OperationInfo() RolloutOperationInfoResponseOutput {
	return o.ApplyT(func(v LookupRolloutResult) RolloutOperationInfoResponse { return v.OperationInfo }).(RolloutOperationInfoResponseOutput)
}

func (o LookupRolloutResultOutput) Services() ServiceResponseArrayOutput {
	return o.ApplyT(func(v LookupRolloutResult) []ServiceResponse { return v.Services }).(ServiceResponseArrayOutput)
}

func (o LookupRolloutResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupRolloutResultOutput) StepGroups() StepGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupRolloutResult) []StepGroupResponse { return v.StepGroups }).(StepGroupResponseArrayOutput)
}

func (o LookupRolloutResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRolloutResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRolloutResultOutput) TargetServiceTopologyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.TargetServiceTopologyId }).(pulumi.StringOutput)
}

func (o LookupRolloutResultOutput) TotalRetryAttempts() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRolloutResult) int { return v.TotalRetryAttempts }).(pulumi.IntOutput)
}

func (o LookupRolloutResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolloutResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRolloutResultOutput{})
}
