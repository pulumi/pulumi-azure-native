


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceFabric(ctx *pulumi.Context, args *LookupServiceFabricArgs, opts ...pulumi.InvokeOption) (*LookupServiceFabricResult, error) {
	var rv LookupServiceFabricResult
	err := ctx.Invoke("azure-native:devtestlab:getServiceFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServiceFabricArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupServiceFabricResult struct {
	ApplicableSchedule      ApplicableScheduleResponse `pulumi:"applicableSchedule"`
	EnvironmentId           *string                    `pulumi:"environmentId"`
	ExternalServiceFabricId *string                    `pulumi:"externalServiceFabricId"`
	Id                      string                     `pulumi:"id"`
	Location                *string                    `pulumi:"location"`
	Name                    string                     `pulumi:"name"`
	ProvisioningState       string                     `pulumi:"provisioningState"`
	Tags                    map[string]string          `pulumi:"tags"`
	Type                    string                     `pulumi:"type"`
	UniqueIdentifier        string                     `pulumi:"uniqueIdentifier"`
}


func (val *LookupServiceFabricResult) Defaults() *LookupServiceFabricResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ApplicableSchedule = *tmp.ApplicableSchedule.Defaults()

	return &tmp
}

func LookupServiceFabricOutput(ctx *pulumi.Context, args LookupServiceFabricOutputArgs, opts ...pulumi.InvokeOption) LookupServiceFabricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceFabricResult, error) {
			args := v.(LookupServiceFabricArgs)
			r, err := LookupServiceFabric(ctx, &args, opts...)
			var s LookupServiceFabricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceFabricResultOutput)
}

type LookupServiceFabricOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	UserName          pulumi.StringInput    `pulumi:"userName"`
}

func (LookupServiceFabricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceFabricArgs)(nil)).Elem()
}


type LookupServiceFabricResultOutput struct{ *pulumi.OutputState }

func (LookupServiceFabricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceFabricResult)(nil)).Elem()
}

func (o LookupServiceFabricResultOutput) ToLookupServiceFabricResultOutput() LookupServiceFabricResultOutput {
	return o
}

func (o LookupServiceFabricResultOutput) ToLookupServiceFabricResultOutputWithContext(ctx context.Context) LookupServiceFabricResultOutput {
	return o
}

func (o LookupServiceFabricResultOutput) ApplicableSchedule() ApplicableScheduleResponseOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) ApplicableScheduleResponse { return v.ApplicableSchedule }).(ApplicableScheduleResponseOutput)
}

func (o LookupServiceFabricResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupServiceFabricResultOutput) ExternalServiceFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) *string { return v.ExternalServiceFabricId }).(pulumi.StringPtrOutput)
}

func (o LookupServiceFabricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceFabricResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupServiceFabricResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceFabricResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServiceFabricResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceFabricResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServiceFabricResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceFabricResultOutput{})
}
