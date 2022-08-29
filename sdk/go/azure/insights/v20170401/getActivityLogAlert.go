


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActivityLogAlert(ctx *pulumi.Context, args *LookupActivityLogAlertArgs, opts ...pulumi.InvokeOption) (*LookupActivityLogAlertResult, error) {
	var rv LookupActivityLogAlertResult
	err := ctx.Invoke("azure-native:insights/v20170401:getActivityLogAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupActivityLogAlertArgs struct {
	ActivityLogAlertName string `pulumi:"activityLogAlertName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupActivityLogAlertResult struct {
	Actions     ActivityLogAlertActionListResponse     `pulumi:"actions"`
	Condition   ActivityLogAlertAllOfConditionResponse `pulumi:"condition"`
	Description *string                                `pulumi:"description"`
	Enabled     *bool                                  `pulumi:"enabled"`
	Id          string                                 `pulumi:"id"`
	Location    string                                 `pulumi:"location"`
	Name        string                                 `pulumi:"name"`
	Scopes      []string                               `pulumi:"scopes"`
	Tags        map[string]string                      `pulumi:"tags"`
	Type        string                                 `pulumi:"type"`
}


func (val *LookupActivityLogAlertResult) Defaults() *LookupActivityLogAlertResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		enabled_ := true
		tmp.Enabled = &enabled_
	}
	return &tmp
}

func LookupActivityLogAlertOutput(ctx *pulumi.Context, args LookupActivityLogAlertOutputArgs, opts ...pulumi.InvokeOption) LookupActivityLogAlertResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActivityLogAlertResult, error) {
			args := v.(LookupActivityLogAlertArgs)
			r, err := LookupActivityLogAlert(ctx, &args, opts...)
			var s LookupActivityLogAlertResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActivityLogAlertResultOutput)
}

type LookupActivityLogAlertOutputArgs struct {
	ActivityLogAlertName pulumi.StringInput `pulumi:"activityLogAlertName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupActivityLogAlertOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivityLogAlertArgs)(nil)).Elem()
}


type LookupActivityLogAlertResultOutput struct{ *pulumi.OutputState }

func (LookupActivityLogAlertResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivityLogAlertResult)(nil)).Elem()
}

func (o LookupActivityLogAlertResultOutput) ToLookupActivityLogAlertResultOutput() LookupActivityLogAlertResultOutput {
	return o
}

func (o LookupActivityLogAlertResultOutput) ToLookupActivityLogAlertResultOutputWithContext(ctx context.Context) LookupActivityLogAlertResultOutput {
	return o
}

func (o LookupActivityLogAlertResultOutput) Actions() ActivityLogAlertActionListResponseOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) ActivityLogAlertActionListResponse { return v.Actions }).(ActivityLogAlertActionListResponseOutput)
}

func (o LookupActivityLogAlertResultOutput) Condition() ActivityLogAlertAllOfConditionResponseOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) ActivityLogAlertAllOfConditionResponse { return v.Condition }).(ActivityLogAlertAllOfConditionResponseOutput)
}

func (o LookupActivityLogAlertResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupActivityLogAlertResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupActivityLogAlertResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActivityLogAlertResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupActivityLogAlertResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupActivityLogAlertResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o LookupActivityLogAlertResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupActivityLogAlertResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityLogAlertResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActivityLogAlertResultOutput{})
}
