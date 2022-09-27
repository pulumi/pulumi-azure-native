


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:containerservice/v20210501:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceConfigurationArgs struct {
	ConfigName        string `pulumi:"configName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMaintenanceConfigurationResult struct {
	Id             string               `pulumi:"id"`
	Name           string               `pulumi:"name"`
	NotAllowedTime []TimeSpanResponse   `pulumi:"notAllowedTime"`
	SystemData     SystemDataResponse   `pulumi:"systemData"`
	TimeInWeek     []TimeInWeekResponse `pulumi:"timeInWeek"`
	Type           string               `pulumi:"type"`
}

func LookupMaintenanceConfigurationOutput(ctx *pulumi.Context, args LookupMaintenanceConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupMaintenanceConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMaintenanceConfigurationResult, error) {
			args := v.(LookupMaintenanceConfigurationArgs)
			r, err := LookupMaintenanceConfiguration(ctx, &args, opts...)
			var s LookupMaintenanceConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMaintenanceConfigurationResultOutput)
}

type LookupMaintenanceConfigurationOutputArgs struct {
	ConfigName        pulumi.StringInput `pulumi:"configName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupMaintenanceConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceConfigurationArgs)(nil)).Elem()
}


type LookupMaintenanceConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupMaintenanceConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceConfigurationResult)(nil)).Elem()
}

func (o LookupMaintenanceConfigurationResultOutput) ToLookupMaintenanceConfigurationResultOutput() LookupMaintenanceConfigurationResultOutput {
	return o
}

func (o LookupMaintenanceConfigurationResultOutput) ToLookupMaintenanceConfigurationResultOutputWithContext(ctx context.Context) LookupMaintenanceConfigurationResultOutput {
	return o
}

func (o LookupMaintenanceConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) NotAllowedTime() TimeSpanResponseArrayOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) []TimeSpanResponse { return v.NotAllowedTime }).(TimeSpanResponseArrayOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) TimeInWeek() TimeInWeekResponseArrayOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) []TimeInWeekResponse { return v.TimeInWeek }).(TimeInWeekResponseArrayOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMaintenanceConfigurationResultOutput{})
}
