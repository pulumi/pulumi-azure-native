


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutomation(ctx *pulumi.Context, args *LookupAutomationArgs, opts ...pulumi.InvokeOption) (*LookupAutomationResult, error) {
	var rv LookupAutomationResult
	err := ctx.Invoke("azure-native:security/v20190101preview:getAutomation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationArgs struct {
	AutomationName    string `pulumi:"automationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAutomationResult struct {
	Actions     []interface{}              `pulumi:"actions"`
	Description *string                    `pulumi:"description"`
	Etag        *string                    `pulumi:"etag"`
	Id          string                     `pulumi:"id"`
	IsEnabled   *bool                      `pulumi:"isEnabled"`
	Kind        *string                    `pulumi:"kind"`
	Location    *string                    `pulumi:"location"`
	Name        string                     `pulumi:"name"`
	Scopes      []AutomationScopeResponse  `pulumi:"scopes"`
	Sources     []AutomationSourceResponse `pulumi:"sources"`
	Tags        map[string]string          `pulumi:"tags"`
	Type        string                     `pulumi:"type"`
}

func LookupAutomationOutput(ctx *pulumi.Context, args LookupAutomationOutputArgs, opts ...pulumi.InvokeOption) LookupAutomationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutomationResult, error) {
			args := v.(LookupAutomationArgs)
			r, err := LookupAutomation(ctx, &args, opts...)
			var s LookupAutomationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutomationResultOutput)
}

type LookupAutomationOutputArgs struct {
	AutomationName    pulumi.StringInput `pulumi:"automationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAutomationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationArgs)(nil)).Elem()
}


type LookupAutomationResultOutput struct{ *pulumi.OutputState }

func (LookupAutomationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationResult)(nil)).Elem()
}

func (o LookupAutomationResultOutput) ToLookupAutomationResultOutput() LookupAutomationResultOutput {
	return o
}

func (o LookupAutomationResultOutput) ToLookupAutomationResultOutputWithContext(ctx context.Context) LookupAutomationResultOutput {
	return o
}

func (o LookupAutomationResultOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupAutomationResult) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o LookupAutomationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAutomationResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAutomationResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAutomationResultOutput) Scopes() AutomationScopeResponseArrayOutput {
	return o.ApplyT(func(v LookupAutomationResult) []AutomationScopeResponse { return v.Scopes }).(AutomationScopeResponseArrayOutput)
}

func (o LookupAutomationResultOutput) Sources() AutomationSourceResponseArrayOutput {
	return o.ApplyT(func(v LookupAutomationResult) []AutomationSourceResponse { return v.Sources }).(AutomationSourceResponseArrayOutput)
}

func (o LookupAutomationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAutomationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAutomationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutomationResultOutput{})
}
