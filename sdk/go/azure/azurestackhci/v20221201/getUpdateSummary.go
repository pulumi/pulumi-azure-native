


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUpdateSummary(ctx *pulumi.Context, args *LookupUpdateSummaryArgs, opts ...pulumi.InvokeOption) (*LookupUpdateSummaryResult, error) {
	var rv LookupUpdateSummaryResult
	err := ctx.Invoke("azure-native:azurestackhci/v20221201:getUpdateSummary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUpdateSummaryArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupUpdateSummaryResult struct {
	CurrentVersion    *string            `pulumi:"currentVersion"`
	HardwareModel     *string            `pulumi:"hardwareModel"`
	HealthCheckDate   *string            `pulumi:"healthCheckDate"`
	Id                string             `pulumi:"id"`
	LastChecked       *string            `pulumi:"lastChecked"`
	LastUpdated       *string            `pulumi:"lastUpdated"`
	Location          *string            `pulumi:"location"`
	Name              string             `pulumi:"name"`
	OemFamily         *string            `pulumi:"oemFamily"`
	ProvisioningState string             `pulumi:"provisioningState"`
	State             *string            `pulumi:"state"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupUpdateSummaryOutput(ctx *pulumi.Context, args LookupUpdateSummaryOutputArgs, opts ...pulumi.InvokeOption) LookupUpdateSummaryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUpdateSummaryResult, error) {
			args := v.(LookupUpdateSummaryArgs)
			r, err := LookupUpdateSummary(ctx, &args, opts...)
			var s LookupUpdateSummaryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUpdateSummaryResultOutput)
}

type LookupUpdateSummaryOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupUpdateSummaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateSummaryArgs)(nil)).Elem()
}


type LookupUpdateSummaryResultOutput struct{ *pulumi.OutputState }

func (LookupUpdateSummaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateSummaryResult)(nil)).Elem()
}

func (o LookupUpdateSummaryResultOutput) ToLookupUpdateSummaryResultOutput() LookupUpdateSummaryResultOutput {
	return o
}

func (o LookupUpdateSummaryResultOutput) ToLookupUpdateSummaryResultOutputWithContext(ctx context.Context) LookupUpdateSummaryResultOutput {
	return o
}

func (o LookupUpdateSummaryResultOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateSummaryResultOutput) HardwareModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.HardwareModel }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateSummaryResultOutput) HealthCheckDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.HealthCheckDate }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateSummaryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUpdateSummaryResultOutput) LastChecked() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.LastChecked }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateSummaryResultOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.LastUpdated }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateSummaryResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateSummaryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUpdateSummaryResultOutput) OemFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.OemFamily }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateSummaryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupUpdateSummaryResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateSummaryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupUpdateSummaryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUpdateSummaryResultOutput{})
}
