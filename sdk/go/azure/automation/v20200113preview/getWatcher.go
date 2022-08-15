


package v20200113preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWatcher(ctx *pulumi.Context, args *LookupWatcherArgs, opts ...pulumi.InvokeOption) (*LookupWatcherResult, error) {
	var rv LookupWatcherResult
	err := ctx.Invoke("azure-native:automation/v20200113preview:getWatcher", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatcherArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	WatcherName           string `pulumi:"watcherName"`
}


type LookupWatcherResult struct {
	CreationTime                string            `pulumi:"creationTime"`
	Description                 *string           `pulumi:"description"`
	Etag                        *string           `pulumi:"etag"`
	ExecutionFrequencyInSeconds *float64          `pulumi:"executionFrequencyInSeconds"`
	Id                          string            `pulumi:"id"`
	LastModifiedBy              string            `pulumi:"lastModifiedBy"`
	LastModifiedTime            string            `pulumi:"lastModifiedTime"`
	Location                    *string           `pulumi:"location"`
	Name                        string            `pulumi:"name"`
	ScriptName                  *string           `pulumi:"scriptName"`
	ScriptParameters            map[string]string `pulumi:"scriptParameters"`
	ScriptRunOn                 *string           `pulumi:"scriptRunOn"`
	Status                      string            `pulumi:"status"`
	Tags                        map[string]string `pulumi:"tags"`
	Type                        string            `pulumi:"type"`
}

func LookupWatcherOutput(ctx *pulumi.Context, args LookupWatcherOutputArgs, opts ...pulumi.InvokeOption) LookupWatcherResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWatcherResult, error) {
			args := v.(LookupWatcherArgs)
			r, err := LookupWatcher(ctx, &args, opts...)
			var s LookupWatcherResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWatcherResultOutput)
}

type LookupWatcherOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	WatcherName           pulumi.StringInput `pulumi:"watcherName"`
}

func (LookupWatcherOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatcherArgs)(nil)).Elem()
}


type LookupWatcherResultOutput struct{ *pulumi.OutputState }

func (LookupWatcherResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatcherResult)(nil)).Elem()
}

func (o LookupWatcherResultOutput) ToLookupWatcherResultOutput() LookupWatcherResultOutput {
	return o
}

func (o LookupWatcherResultOutput) ToLookupWatcherResultOutputWithContext(ctx context.Context) LookupWatcherResultOutput {
	return o
}

func (o LookupWatcherResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupWatcherResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWatcherResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupWatcherResultOutput) ExecutionFrequencyInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *float64 { return v.ExecutionFrequencyInSeconds }).(pulumi.Float64PtrOutput)
}

func (o LookupWatcherResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWatcherResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o LookupWatcherResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupWatcherResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWatcherResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWatcherResultOutput) ScriptName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.ScriptName }).(pulumi.StringPtrOutput)
}

func (o LookupWatcherResultOutput) ScriptParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWatcherResult) map[string]string { return v.ScriptParameters }).(pulumi.StringMapOutput)
}

func (o LookupWatcherResultOutput) ScriptRunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.ScriptRunOn }).(pulumi.StringPtrOutput)
}

func (o LookupWatcherResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupWatcherResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWatcherResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWatcherResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWatcherResultOutput{})
}
