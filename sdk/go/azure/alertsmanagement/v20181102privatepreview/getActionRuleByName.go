


package v20181102privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupActionRuleByName(ctx *pulumi.Context, args *LookupActionRuleByNameArgs, opts ...pulumi.InvokeOption) (*LookupActionRuleByNameResult, error) {
	var rv LookupActionRuleByNameResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20181102privatepreview:getActionRuleByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActionRuleByNameArgs struct {
	ActionRuleName string `pulumi:"actionRuleName"`
	ResourceGroup  string `pulumi:"resourceGroup"`
}


type LookupActionRuleByNameResult struct {
	Id         string                       `pulumi:"id"`
	Location   string                       `pulumi:"location"`
	Name       string                       `pulumi:"name"`
	Properties ActionRulePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string            `pulumi:"tags"`
	Type       string                       `pulumi:"type"`
}

func LookupActionRuleByNameOutput(ctx *pulumi.Context, args LookupActionRuleByNameOutputArgs, opts ...pulumi.InvokeOption) LookupActionRuleByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActionRuleByNameResult, error) {
			args := v.(LookupActionRuleByNameArgs)
			r, err := LookupActionRuleByName(ctx, &args, opts...)
			var s LookupActionRuleByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActionRuleByNameResultOutput)
}

type LookupActionRuleByNameOutputArgs struct {
	ActionRuleName pulumi.StringInput `pulumi:"actionRuleName"`
	ResourceGroup  pulumi.StringInput `pulumi:"resourceGroup"`
}

func (LookupActionRuleByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionRuleByNameArgs)(nil)).Elem()
}


type LookupActionRuleByNameResultOutput struct{ *pulumi.OutputState }

func (LookupActionRuleByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionRuleByNameResult)(nil)).Elem()
}

func (o LookupActionRuleByNameResultOutput) ToLookupActionRuleByNameResultOutput() LookupActionRuleByNameResultOutput {
	return o
}

func (o LookupActionRuleByNameResultOutput) ToLookupActionRuleByNameResultOutputWithContext(ctx context.Context) LookupActionRuleByNameResultOutput {
	return o
}

func (o LookupActionRuleByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionRuleByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActionRuleByNameResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionRuleByNameResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupActionRuleByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionRuleByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupActionRuleByNameResultOutput) Properties() ActionRulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupActionRuleByNameResult) ActionRulePropertiesResponse { return v.Properties }).(ActionRulePropertiesResponseOutput)
}

func (o LookupActionRuleByNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupActionRuleByNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupActionRuleByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionRuleByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActionRuleByNameResultOutput{})
}
