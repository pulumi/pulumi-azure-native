


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTaskRun(ctx *pulumi.Context, args *LookupTaskRunArgs, opts ...pulumi.InvokeOption) (*LookupTaskRunResult, error) {
	var rv LookupTaskRunResult
	err := ctx.Invoke("azure-native:containerregistry/v20190601preview:getTaskRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTaskRunArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TaskRunName       string `pulumi:"taskRunName"`
}



type LookupTaskRunResult struct {
	ForceUpdateTag    *string                     `pulumi:"forceUpdateTag"`
	Id                string                      `pulumi:"id"`
	Identity          *IdentityPropertiesResponse `pulumi:"identity"`
	Location          *string                     `pulumi:"location"`
	Name              string                      `pulumi:"name"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	RunRequest        interface{}                 `pulumi:"runRequest"`
	RunResult         RunResponse                 `pulumi:"runResult"`
	SystemData        SystemDataResponse          `pulumi:"systemData"`
	Type              string                      `pulumi:"type"`
}


func (val *LookupTaskRunResult) Defaults() *LookupTaskRunResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.RunResult = *tmp.RunResult.Defaults()

	return &tmp
}

func LookupTaskRunOutput(ctx *pulumi.Context, args LookupTaskRunOutputArgs, opts ...pulumi.InvokeOption) LookupTaskRunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTaskRunResult, error) {
			args := v.(LookupTaskRunArgs)
			r, err := LookupTaskRun(ctx, &args, opts...)
			var s LookupTaskRunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTaskRunResultOutput)
}

type LookupTaskRunOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TaskRunName       pulumi.StringInput `pulumi:"taskRunName"`
}

func (LookupTaskRunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskRunArgs)(nil)).Elem()
}



type LookupTaskRunResultOutput struct{ *pulumi.OutputState }

func (LookupTaskRunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskRunResult)(nil)).Elem()
}

func (o LookupTaskRunResultOutput) ToLookupTaskRunResultOutput() LookupTaskRunResultOutput {
	return o
}

func (o LookupTaskRunResultOutput) ToLookupTaskRunResultOutputWithContext(ctx context.Context) LookupTaskRunResultOutput {
	return o
}

func (o LookupTaskRunResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskRunResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupTaskRunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskRunResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTaskRunResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupTaskRunResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupTaskRunResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskRunResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupTaskRunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskRunResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTaskRunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskRunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTaskRunResultOutput) RunRequest() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTaskRunResult) interface{} { return v.RunRequest }).(pulumi.AnyOutput)
}

func (o LookupTaskRunResultOutput) RunResult() RunResponseOutput {
	return o.ApplyT(func(v LookupTaskRunResult) RunResponse { return v.RunResult }).(RunResponseOutput)
}

func (o LookupTaskRunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTaskRunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTaskRunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskRunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskRunResultOutput{})
}
