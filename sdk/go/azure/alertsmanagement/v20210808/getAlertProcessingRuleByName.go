


package v20210808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlertProcessingRuleByName(ctx *pulumi.Context, args *LookupAlertProcessingRuleByNameArgs, opts ...pulumi.InvokeOption) (*LookupAlertProcessingRuleByNameResult, error) {
	var rv LookupAlertProcessingRuleByNameResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20210808:getAlertProcessingRuleByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAlertProcessingRuleByNameArgs struct {
	AlertProcessingRuleName string `pulumi:"alertProcessingRuleName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupAlertProcessingRuleByNameResult struct {
	Id         string                                `pulumi:"id"`
	Location   string                                `pulumi:"location"`
	Name       string                                `pulumi:"name"`
	Properties AlertProcessingRulePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                    `pulumi:"systemData"`
	Tags       map[string]string                     `pulumi:"tags"`
	Type       string                                `pulumi:"type"`
}


func (val *LookupAlertProcessingRuleByNameResult) Defaults() *LookupAlertProcessingRuleByNameResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupAlertProcessingRuleByNameOutput(ctx *pulumi.Context, args LookupAlertProcessingRuleByNameOutputArgs, opts ...pulumi.InvokeOption) LookupAlertProcessingRuleByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlertProcessingRuleByNameResult, error) {
			args := v.(LookupAlertProcessingRuleByNameArgs)
			r, err := LookupAlertProcessingRuleByName(ctx, &args, opts...)
			var s LookupAlertProcessingRuleByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlertProcessingRuleByNameResultOutput)
}

type LookupAlertProcessingRuleByNameOutputArgs struct {
	AlertProcessingRuleName pulumi.StringInput `pulumi:"alertProcessingRuleName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAlertProcessingRuleByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertProcessingRuleByNameArgs)(nil)).Elem()
}


type LookupAlertProcessingRuleByNameResultOutput struct{ *pulumi.OutputState }

func (LookupAlertProcessingRuleByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertProcessingRuleByNameResult)(nil)).Elem()
}

func (o LookupAlertProcessingRuleByNameResultOutput) ToLookupAlertProcessingRuleByNameResultOutput() LookupAlertProcessingRuleByNameResultOutput {
	return o
}

func (o LookupAlertProcessingRuleByNameResultOutput) ToLookupAlertProcessingRuleByNameResultOutputWithContext(ctx context.Context) LookupAlertProcessingRuleByNameResultOutput {
	return o
}

func (o LookupAlertProcessingRuleByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlertProcessingRuleByNameResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAlertProcessingRuleByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAlertProcessingRuleByNameResultOutput) Properties() AlertProcessingRulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) AlertProcessingRulePropertiesResponse {
		return v.Properties
	}).(AlertProcessingRulePropertiesResponseOutput)
}

func (o LookupAlertProcessingRuleByNameResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAlertProcessingRuleByNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAlertProcessingRuleByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlertProcessingRuleByNameResultOutput{})
}
