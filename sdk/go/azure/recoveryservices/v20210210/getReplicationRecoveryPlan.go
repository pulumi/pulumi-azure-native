


package v20210210

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationRecoveryPlan(ctx *pulumi.Context, args *LookupReplicationRecoveryPlanArgs, opts ...pulumi.InvokeOption) (*LookupReplicationRecoveryPlanResult, error) {
	var rv LookupReplicationRecoveryPlanResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210210:getReplicationRecoveryPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationRecoveryPlanArgs struct {
	RecoveryPlanName  string `pulumi:"recoveryPlanName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupReplicationRecoveryPlanResult struct {
	Id         string                         `pulumi:"id"`
	Location   *string                        `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties RecoveryPlanPropertiesResponse `pulumi:"properties"`
	Type       string                         `pulumi:"type"`
}

func LookupReplicationRecoveryPlanOutput(ctx *pulumi.Context, args LookupReplicationRecoveryPlanOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationRecoveryPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationRecoveryPlanResult, error) {
			args := v.(LookupReplicationRecoveryPlanArgs)
			r, err := LookupReplicationRecoveryPlan(ctx, &args, opts...)
			var s LookupReplicationRecoveryPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationRecoveryPlanResultOutput)
}

type LookupReplicationRecoveryPlanOutputArgs struct {
	RecoveryPlanName  pulumi.StringInput `pulumi:"recoveryPlanName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationRecoveryPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationRecoveryPlanArgs)(nil)).Elem()
}


type LookupReplicationRecoveryPlanResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationRecoveryPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationRecoveryPlanResult)(nil)).Elem()
}

func (o LookupReplicationRecoveryPlanResultOutput) ToLookupReplicationRecoveryPlanResultOutput() LookupReplicationRecoveryPlanResultOutput {
	return o
}

func (o LookupReplicationRecoveryPlanResultOutput) ToLookupReplicationRecoveryPlanResultOutputWithContext(ctx context.Context) LookupReplicationRecoveryPlanResultOutput {
	return o
}

func (o LookupReplicationRecoveryPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationRecoveryPlanResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationRecoveryPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationRecoveryPlanResultOutput) Properties() RecoveryPlanPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) RecoveryPlanPropertiesResponse { return v.Properties }).(RecoveryPlanPropertiesResponseOutput)
}

func (o LookupReplicationRecoveryPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationRecoveryPlanResultOutput{})
}
