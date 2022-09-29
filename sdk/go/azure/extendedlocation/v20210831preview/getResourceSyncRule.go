


package v20210831preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceSyncRule(ctx *pulumi.Context, args *LookupResourceSyncRuleArgs, opts ...pulumi.InvokeOption) (*LookupResourceSyncRuleResult, error) {
	var rv LookupResourceSyncRuleResult
	err := ctx.Invoke("azure-native:extendedlocation/v20210831preview:getResourceSyncRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceSyncRuleArgs struct {
	ChildResourceName string `pulumi:"childResourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupResourceSyncRuleResult struct {
	Id                  string                                      `pulumi:"id"`
	Location            string                                      `pulumi:"location"`
	Name                string                                      `pulumi:"name"`
	Priority            *int                                        `pulumi:"priority"`
	ProvisioningState   string                                      `pulumi:"provisioningState"`
	Selector            *ResourceSyncRulePropertiesResponseSelector `pulumi:"selector"`
	SystemData          SystemDataResponse                          `pulumi:"systemData"`
	Tags                map[string]string                           `pulumi:"tags"`
	TargetResourceGroup *string                                     `pulumi:"targetResourceGroup"`
	Type                string                                      `pulumi:"type"`
}

func LookupResourceSyncRuleOutput(ctx *pulumi.Context, args LookupResourceSyncRuleOutputArgs, opts ...pulumi.InvokeOption) LookupResourceSyncRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceSyncRuleResult, error) {
			args := v.(LookupResourceSyncRuleArgs)
			r, err := LookupResourceSyncRule(ctx, &args, opts...)
			var s LookupResourceSyncRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceSyncRuleResultOutput)
}

type LookupResourceSyncRuleOutputArgs struct {
	ChildResourceName pulumi.StringInput `pulumi:"childResourceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupResourceSyncRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceSyncRuleArgs)(nil)).Elem()
}


type LookupResourceSyncRuleResultOutput struct{ *pulumi.OutputState }

func (LookupResourceSyncRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceSyncRuleResult)(nil)).Elem()
}

func (o LookupResourceSyncRuleResultOutput) ToLookupResourceSyncRuleResultOutput() LookupResourceSyncRuleResultOutput {
	return o
}

func (o LookupResourceSyncRuleResultOutput) ToLookupResourceSyncRuleResultOutputWithContext(ctx context.Context) LookupResourceSyncRuleResultOutput {
	return o
}

func (o LookupResourceSyncRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceSyncRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupResourceSyncRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourceSyncRuleResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupResourceSyncRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupResourceSyncRuleResultOutput) Selector() ResourceSyncRulePropertiesResponseSelectorPtrOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) *ResourceSyncRulePropertiesResponseSelector { return v.Selector }).(ResourceSyncRulePropertiesResponseSelectorPtrOutput)
}

func (o LookupResourceSyncRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupResourceSyncRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupResourceSyncRuleResultOutput) TargetResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) *string { return v.TargetResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupResourceSyncRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceSyncRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceSyncRuleResultOutput{})
}
