


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSentinelOnboardingState(ctx *pulumi.Context, args *LookupSentinelOnboardingStateArgs, opts ...pulumi.InvokeOption) (*LookupSentinelOnboardingStateResult, error) {
	var rv LookupSentinelOnboardingStateResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101preview:getSentinelOnboardingState", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSentinelOnboardingStateArgs struct {
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	SentinelOnboardingStateName string `pulumi:"sentinelOnboardingStateName"`
	WorkspaceName               string `pulumi:"workspaceName"`
}


type LookupSentinelOnboardingStateResult struct {
	CustomerManagedKey *bool              `pulumi:"customerManagedKey"`
	Etag               *string            `pulumi:"etag"`
	Id                 string             `pulumi:"id"`
	Name               string             `pulumi:"name"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}

func LookupSentinelOnboardingStateOutput(ctx *pulumi.Context, args LookupSentinelOnboardingStateOutputArgs, opts ...pulumi.InvokeOption) LookupSentinelOnboardingStateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSentinelOnboardingStateResult, error) {
			args := v.(LookupSentinelOnboardingStateArgs)
			r, err := LookupSentinelOnboardingState(ctx, &args, opts...)
			var s LookupSentinelOnboardingStateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSentinelOnboardingStateResultOutput)
}

type LookupSentinelOnboardingStateOutputArgs struct {
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
	SentinelOnboardingStateName pulumi.StringInput `pulumi:"sentinelOnboardingStateName"`
	WorkspaceName               pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSentinelOnboardingStateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSentinelOnboardingStateArgs)(nil)).Elem()
}


type LookupSentinelOnboardingStateResultOutput struct{ *pulumi.OutputState }

func (LookupSentinelOnboardingStateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSentinelOnboardingStateResult)(nil)).Elem()
}

func (o LookupSentinelOnboardingStateResultOutput) ToLookupSentinelOnboardingStateResultOutput() LookupSentinelOnboardingStateResultOutput {
	return o
}

func (o LookupSentinelOnboardingStateResultOutput) ToLookupSentinelOnboardingStateResultOutputWithContext(ctx context.Context) LookupSentinelOnboardingStateResultOutput {
	return o
}

func (o LookupSentinelOnboardingStateResultOutput) CustomerManagedKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) *bool { return v.CustomerManagedKey }).(pulumi.BoolPtrOutput)
}

func (o LookupSentinelOnboardingStateResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSentinelOnboardingStateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSentinelOnboardingStateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSentinelOnboardingStateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSentinelOnboardingStateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentinelOnboardingStateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSentinelOnboardingStateResultOutput{})
}
