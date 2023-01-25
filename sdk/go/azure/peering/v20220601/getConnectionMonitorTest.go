


package v20220601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectionMonitorTest(ctx *pulumi.Context, args *LookupConnectionMonitorTestArgs, opts ...pulumi.InvokeOption) (*LookupConnectionMonitorTestResult, error) {
	var rv LookupConnectionMonitorTestResult
	err := ctx.Invoke("azure-native:peering/v20220601:getConnectionMonitorTest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionMonitorTestArgs struct {
	ConnectionMonitorTestName string `pulumi:"connectionMonitorTestName"`
	PeeringServiceName        string `pulumi:"peeringServiceName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupConnectionMonitorTestResult struct {
	Destination        *string  `pulumi:"destination"`
	DestinationPort    *int     `pulumi:"destinationPort"`
	Id                 string   `pulumi:"id"`
	IsTestSuccessful   bool     `pulumi:"isTestSuccessful"`
	Name               string   `pulumi:"name"`
	Path               []string `pulumi:"path"`
	ProvisioningState  string   `pulumi:"provisioningState"`
	SourceAgent        *string  `pulumi:"sourceAgent"`
	TestFrequencyInSec *int     `pulumi:"testFrequencyInSec"`
	Type               string   `pulumi:"type"`
}

func LookupConnectionMonitorTestOutput(ctx *pulumi.Context, args LookupConnectionMonitorTestOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionMonitorTestResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionMonitorTestResult, error) {
			args := v.(LookupConnectionMonitorTestArgs)
			r, err := LookupConnectionMonitorTest(ctx, &args, opts...)
			var s LookupConnectionMonitorTestResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionMonitorTestResultOutput)
}

type LookupConnectionMonitorTestOutputArgs struct {
	ConnectionMonitorTestName pulumi.StringInput `pulumi:"connectionMonitorTestName"`
	PeeringServiceName        pulumi.StringInput `pulumi:"peeringServiceName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectionMonitorTestOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionMonitorTestArgs)(nil)).Elem()
}


type LookupConnectionMonitorTestResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionMonitorTestResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionMonitorTestResult)(nil)).Elem()
}

func (o LookupConnectionMonitorTestResultOutput) ToLookupConnectionMonitorTestResultOutput() LookupConnectionMonitorTestResultOutput {
	return o
}

func (o LookupConnectionMonitorTestResultOutput) ToLookupConnectionMonitorTestResultOutputWithContext(ctx context.Context) LookupConnectionMonitorTestResultOutput {
	return o
}

func (o LookupConnectionMonitorTestResultOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionMonitorTestResultOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o LookupConnectionMonitorTestResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectionMonitorTestResultOutput) IsTestSuccessful() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) bool { return v.IsTestSuccessful }).(pulumi.BoolOutput)
}

func (o LookupConnectionMonitorTestResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectionMonitorTestResultOutput) Path() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) []string { return v.Path }).(pulumi.StringArrayOutput)
}

func (o LookupConnectionMonitorTestResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupConnectionMonitorTestResultOutput) SourceAgent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) *string { return v.SourceAgent }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionMonitorTestResultOutput) TestFrequencyInSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) *int { return v.TestFrequencyInSec }).(pulumi.IntPtrOutput)
}

func (o LookupConnectionMonitorTestResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorTestResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionMonitorTestResultOutput{})
}
