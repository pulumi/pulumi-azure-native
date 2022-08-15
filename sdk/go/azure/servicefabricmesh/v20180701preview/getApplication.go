


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:servicefabricmesh/v20180701preview:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	DebugParams         *string                              `pulumi:"debugParams"`
	Description         *string                              `pulumi:"description"`
	Diagnostics         *DiagnosticsDescriptionResponse      `pulumi:"diagnostics"`
	HealthState         string                               `pulumi:"healthState"`
	Id                  string                               `pulumi:"id"`
	Location            string                               `pulumi:"location"`
	Name                string                               `pulumi:"name"`
	ProvisioningState   string                               `pulumi:"provisioningState"`
	ServiceNames        []string                             `pulumi:"serviceNames"`
	Services            []ServiceResourceDescriptionResponse `pulumi:"services"`
	Status              string                               `pulumi:"status"`
	StatusDetails       string                               `pulumi:"statusDetails"`
	Tags                map[string]string                    `pulumi:"tags"`
	Type                string                               `pulumi:"type"`
	UnhealthyEvaluation string                               `pulumi:"unhealthyEvaluation"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}


type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) DebugParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.DebugParams }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Diagnostics() DiagnosticsDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *DiagnosticsDescriptionResponse { return v.Diagnostics }).(DiagnosticsDescriptionResponsePtrOutput)
}

func (o LookupApplicationResultOutput) HealthState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.HealthState }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) ServiceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []string { return v.ServiceNames }).(pulumi.StringArrayOutput)
}

func (o LookupApplicationResultOutput) Services() ServiceResourceDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ServiceResourceDescriptionResponse { return v.Services }).(ServiceResourceDescriptionResponseArrayOutput)
}

func (o LookupApplicationResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.StatusDetails }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) UnhealthyEvaluation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.UnhealthyEvaluation }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
