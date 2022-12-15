


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupConnectionMonitor(ctx *pulumi.Context, args *LookupConnectionMonitorArgs, opts ...pulumi.InvokeOption) (*LookupConnectionMonitorResult, error) {
	var rv LookupConnectionMonitorResult
	err := ctx.Invoke("azure-native:network/v20180201:getConnectionMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectionMonitorArgs struct {
	ConnectionMonitorName string `pulumi:"connectionMonitorName"`
	NetworkWatcherName    string `pulumi:"networkWatcherName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupConnectionMonitorResult struct {
	AutoStart                   *bool                                `pulumi:"autoStart"`
	Destination                 ConnectionMonitorDestinationResponse `pulumi:"destination"`
	Etag                        *string                              `pulumi:"etag"`
	Id                          string                               `pulumi:"id"`
	Location                    *string                              `pulumi:"location"`
	MonitoringIntervalInSeconds *int                                 `pulumi:"monitoringIntervalInSeconds"`
	MonitoringStatus            *string                              `pulumi:"monitoringStatus"`
	Name                        string                               `pulumi:"name"`
	ProvisioningState           *string                              `pulumi:"provisioningState"`
	Source                      ConnectionMonitorSourceResponse      `pulumi:"source"`
	StartTime                   *string                              `pulumi:"startTime"`
	Tags                        map[string]string                    `pulumi:"tags"`
	Type                        string                               `pulumi:"type"`
}


func (val *LookupConnectionMonitorResult) Defaults() *LookupConnectionMonitorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutoStart) {
		autoStart_ := true
		tmp.AutoStart = &autoStart_
	}
	if isZero(tmp.Etag) {
		etag_ := "A unique read-only string that changes whenever the resource is updated."
		tmp.Etag = &etag_
	}
	if isZero(tmp.MonitoringIntervalInSeconds) {
		monitoringIntervalInSeconds_ := 60
		tmp.MonitoringIntervalInSeconds = &monitoringIntervalInSeconds_
	}
	return &tmp
}

func LookupConnectionMonitorOutput(ctx *pulumi.Context, args LookupConnectionMonitorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionMonitorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionMonitorResult, error) {
			args := v.(LookupConnectionMonitorArgs)
			r, err := LookupConnectionMonitor(ctx, &args, opts...)
			var s LookupConnectionMonitorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionMonitorResultOutput)
}

type LookupConnectionMonitorOutputArgs struct {
	ConnectionMonitorName pulumi.StringInput `pulumi:"connectionMonitorName"`
	NetworkWatcherName    pulumi.StringInput `pulumi:"networkWatcherName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectionMonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionMonitorArgs)(nil)).Elem()
}


type LookupConnectionMonitorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionMonitorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionMonitorResult)(nil)).Elem()
}

func (o LookupConnectionMonitorResultOutput) ToLookupConnectionMonitorResultOutput() LookupConnectionMonitorResultOutput {
	return o
}

func (o LookupConnectionMonitorResultOutput) ToLookupConnectionMonitorResultOutputWithContext(ctx context.Context) LookupConnectionMonitorResultOutput {
	return o
}

func (o LookupConnectionMonitorResultOutput) AutoStart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *bool { return v.AutoStart }).(pulumi.BoolPtrOutput)
}

func (o LookupConnectionMonitorResultOutput) Destination() ConnectionMonitorDestinationResponseOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) ConnectionMonitorDestinationResponse { return v.Destination }).(ConnectionMonitorDestinationResponseOutput)
}

func (o LookupConnectionMonitorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionMonitorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectionMonitorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionMonitorResultOutput) MonitoringIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *int { return v.MonitoringIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupConnectionMonitorResultOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionMonitorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectionMonitorResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionMonitorResultOutput) Source() ConnectionMonitorSourceResponseOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) ConnectionMonitorSourceResponse { return v.Source }).(ConnectionMonitorSourceResponseOutput)
}

func (o LookupConnectionMonitorResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionMonitorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConnectionMonitorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionMonitorResultOutput{})
}
