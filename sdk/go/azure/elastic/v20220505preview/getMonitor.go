


package v20220505preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMonitor(ctx *pulumi.Context, args *LookupMonitorArgs, opts ...pulumi.InvokeOption) (*LookupMonitorResult, error) {
	var rv LookupMonitorResult
	err := ctx.Invoke("azure-native:elastic/v20220505preview:getMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitorArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMonitorResult struct {
	Id         string                      `pulumi:"id"`
	Identity   *IdentityPropertiesResponse `pulumi:"identity"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties MonitorPropertiesResponse   `pulumi:"properties"`
	Sku        *ResourceSkuResponse        `pulumi:"sku"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}

func LookupMonitorOutput(ctx *pulumi.Context, args LookupMonitorOutputArgs, opts ...pulumi.InvokeOption) LookupMonitorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitorResult, error) {
			args := v.(LookupMonitorArgs)
			r, err := LookupMonitor(ctx, &args, opts...)
			var s LookupMonitorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitorResultOutput)
}

type LookupMonitorOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitorArgs)(nil)).Elem()
}


type LookupMonitorResultOutput struct{ *pulumi.OutputState }

func (LookupMonitorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitorResult)(nil)).Elem()
}

func (o LookupMonitorResultOutput) ToLookupMonitorResultOutput() LookupMonitorResultOutput {
	return o
}

func (o LookupMonitorResultOutput) ToLookupMonitorResultOutputWithContext(ctx context.Context) LookupMonitorResultOutput {
	return o
}

func (o LookupMonitorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMonitorResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupMonitorResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupMonitorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMonitorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMonitorResultOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMonitorResult) MonitorPropertiesResponse { return v.Properties }).(MonitorPropertiesResponseOutput)
}

func (o LookupMonitorResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupMonitorResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o LookupMonitorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMonitorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMonitorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMonitorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMonitorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitorResultOutput{})
}
