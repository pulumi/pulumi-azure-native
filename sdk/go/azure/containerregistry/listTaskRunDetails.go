


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTaskRunDetails(ctx *pulumi.Context, args *ListTaskRunDetailsArgs, opts ...pulumi.InvokeOption) (*ListTaskRunDetailsResult, error) {
	var rv ListTaskRunDetailsResult
	err := ctx.Invoke("azure-native:containerregistry:listTaskRunDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListTaskRunDetailsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TaskRunName       string `pulumi:"taskRunName"`
}



type ListTaskRunDetailsResult struct {
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


func (val *ListTaskRunDetailsResult) Defaults() *ListTaskRunDetailsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.RunResult = *tmp.RunResult.Defaults()

	return &tmp
}

func ListTaskRunDetailsOutput(ctx *pulumi.Context, args ListTaskRunDetailsOutputArgs, opts ...pulumi.InvokeOption) ListTaskRunDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTaskRunDetailsResult, error) {
			args := v.(ListTaskRunDetailsArgs)
			r, err := ListTaskRunDetails(ctx, &args, opts...)
			var s ListTaskRunDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTaskRunDetailsResultOutput)
}

type ListTaskRunDetailsOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TaskRunName       pulumi.StringInput `pulumi:"taskRunName"`
}

func (ListTaskRunDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTaskRunDetailsArgs)(nil)).Elem()
}



type ListTaskRunDetailsResultOutput struct{ *pulumi.OutputState }

func (ListTaskRunDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTaskRunDetailsResult)(nil)).Elem()
}

func (o ListTaskRunDetailsResultOutput) ToListTaskRunDetailsResultOutput() ListTaskRunDetailsResultOutput {
	return o
}

func (o ListTaskRunDetailsResultOutput) ToListTaskRunDetailsResultOutputWithContext(ctx context.Context) ListTaskRunDetailsResultOutput {
	return o
}

func (o ListTaskRunDetailsResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o ListTaskRunDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListTaskRunDetailsResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o ListTaskRunDetailsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ListTaskRunDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListTaskRunDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ListTaskRunDetailsResultOutput) RunRequest() pulumi.AnyOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) interface{} { return v.RunRequest }).(pulumi.AnyOutput)
}

func (o ListTaskRunDetailsResultOutput) RunResult() RunResponseOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) RunResponse { return v.RunResult }).(RunResponseOutput)
}

func (o ListTaskRunDetailsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ListTaskRunDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskRunDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTaskRunDetailsResultOutput{})
}
