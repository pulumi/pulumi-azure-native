


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUpdateRun(ctx *pulumi.Context, args *LookupUpdateRunArgs, opts ...pulumi.InvokeOption) (*LookupUpdateRunResult, error) {
	var rv LookupUpdateRunResult
	err := ctx.Invoke("azure-native:azurestackhci/v20221201:getUpdateRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUpdateRunArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	UpdateName        string `pulumi:"updateName"`
	UpdateRunName     string `pulumi:"updateRunName"`
}


type LookupUpdateRunResult struct {
	Description        *string            `pulumi:"description"`
	Duration           *string            `pulumi:"duration"`
	EndTimeUtc         *string            `pulumi:"endTimeUtc"`
	ErrorMessage       *string            `pulumi:"errorMessage"`
	Id                 string             `pulumi:"id"`
	LastUpdatedTime    *string            `pulumi:"lastUpdatedTime"`
	LastUpdatedTimeUtc *string            `pulumi:"lastUpdatedTimeUtc"`
	Location           *string            `pulumi:"location"`
	Name               string             `pulumi:"name"`
	ProvisioningState  string             `pulumi:"provisioningState"`
	StartTimeUtc       *string            `pulumi:"startTimeUtc"`
	State              *string            `pulumi:"state"`
	Status             *string            `pulumi:"status"`
	Steps              []StepResponse     `pulumi:"steps"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	TimeStarted        *string            `pulumi:"timeStarted"`
	Type               string             `pulumi:"type"`
}

func LookupUpdateRunOutput(ctx *pulumi.Context, args LookupUpdateRunOutputArgs, opts ...pulumi.InvokeOption) LookupUpdateRunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUpdateRunResult, error) {
			args := v.(LookupUpdateRunArgs)
			r, err := LookupUpdateRun(ctx, &args, opts...)
			var s LookupUpdateRunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUpdateRunResultOutput)
}

type LookupUpdateRunOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	UpdateName        pulumi.StringInput `pulumi:"updateName"`
	UpdateRunName     pulumi.StringInput `pulumi:"updateRunName"`
}

func (LookupUpdateRunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateRunArgs)(nil)).Elem()
}


type LookupUpdateRunResultOutput struct{ *pulumi.OutputState }

func (LookupUpdateRunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateRunResult)(nil)).Elem()
}

func (o LookupUpdateRunResultOutput) ToLookupUpdateRunResultOutput() LookupUpdateRunResultOutput {
	return o
}

func (o LookupUpdateRunResultOutput) ToLookupUpdateRunResultOutputWithContext(ctx context.Context) LookupUpdateRunResultOutput {
	return o
}

func (o LookupUpdateRunResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.EndTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUpdateRunResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.LastUpdatedTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUpdateRunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupUpdateRunResultOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.StartTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) Steps() StepResponseArrayOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) []StepResponse { return v.Steps }).(StepResponseArrayOutput)
}

func (o LookupUpdateRunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupUpdateRunResultOutput) TimeStarted() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.TimeStarted }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateRunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUpdateRunResultOutput{})
}
