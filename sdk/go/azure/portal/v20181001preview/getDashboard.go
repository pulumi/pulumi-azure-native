


package v20181001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDashboard(ctx *pulumi.Context, args *LookupDashboardArgs, opts ...pulumi.InvokeOption) (*LookupDashboardResult, error) {
	var rv LookupDashboardResult
	err := ctx.Invoke("azure-native:portal/v20181001preview:getDashboard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDashboardArgs struct {
	DashboardName     string `pulumi:"dashboardName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDashboardResult struct {
	Id       string                           `pulumi:"id"`
	Lenses   map[string]DashboardLensResponse `pulumi:"lenses"`
	Location string                           `pulumi:"location"`
	Metadata map[string]interface{}           `pulumi:"metadata"`
	Name     string                           `pulumi:"name"`
	Tags     map[string]string                `pulumi:"tags"`
	Type     string                           `pulumi:"type"`
}

func LookupDashboardOutput(ctx *pulumi.Context, args LookupDashboardOutputArgs, opts ...pulumi.InvokeOption) LookupDashboardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDashboardResult, error) {
			args := v.(LookupDashboardArgs)
			r, err := LookupDashboard(ctx, &args, opts...)
			var s LookupDashboardResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDashboardResultOutput)
}

type LookupDashboardOutputArgs struct {
	DashboardName     pulumi.StringInput `pulumi:"dashboardName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDashboardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDashboardArgs)(nil)).Elem()
}


type LookupDashboardResultOutput struct{ *pulumi.OutputState }

func (LookupDashboardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDashboardResult)(nil)).Elem()
}

func (o LookupDashboardResultOutput) ToLookupDashboardResultOutput() LookupDashboardResultOutput {
	return o
}

func (o LookupDashboardResultOutput) ToLookupDashboardResultOutputWithContext(ctx context.Context) LookupDashboardResultOutput {
	return o
}

func (o LookupDashboardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDashboardResultOutput) Lenses() DashboardLensResponseMapOutput {
	return o.ApplyT(func(v LookupDashboardResult) map[string]DashboardLensResponse { return v.Lenses }).(DashboardLensResponseMapOutput)
}

func (o LookupDashboardResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDashboardResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupDashboardResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o LookupDashboardResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDashboardResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDashboardResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDashboardResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDashboardResultOutput{})
}
