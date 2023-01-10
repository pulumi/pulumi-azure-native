


package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridRunbookWorker(ctx *pulumi.Context, args *LookupHybridRunbookWorkerArgs, opts ...pulumi.InvokeOption) (*LookupHybridRunbookWorkerResult, error) {
	var rv LookupHybridRunbookWorkerResult
	err := ctx.Invoke("azure-native:automation:getHybridRunbookWorker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridRunbookWorkerArgs struct {
	AutomationAccountName        string `pulumi:"automationAccountName"`
	HybridRunbookWorkerGroupName string `pulumi:"hybridRunbookWorkerGroupName"`
	HybridRunbookWorkerId        string `pulumi:"hybridRunbookWorkerId"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupHybridRunbookWorkerResult struct {
	Id                 string             `pulumi:"id"`
	Ip                 *string            `pulumi:"ip"`
	LastSeenDateTime   *string            `pulumi:"lastSeenDateTime"`
	Name               string             `pulumi:"name"`
	RegisteredDateTime *string            `pulumi:"registeredDateTime"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
	VmResourceId       *string            `pulumi:"vmResourceId"`
	WorkerName         *string            `pulumi:"workerName"`
	WorkerType         *string            `pulumi:"workerType"`
}

func LookupHybridRunbookWorkerOutput(ctx *pulumi.Context, args LookupHybridRunbookWorkerOutputArgs, opts ...pulumi.InvokeOption) LookupHybridRunbookWorkerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridRunbookWorkerResult, error) {
			args := v.(LookupHybridRunbookWorkerArgs)
			r, err := LookupHybridRunbookWorker(ctx, &args, opts...)
			var s LookupHybridRunbookWorkerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridRunbookWorkerResultOutput)
}

type LookupHybridRunbookWorkerOutputArgs struct {
	AutomationAccountName        pulumi.StringInput `pulumi:"automationAccountName"`
	HybridRunbookWorkerGroupName pulumi.StringInput `pulumi:"hybridRunbookWorkerGroupName"`
	HybridRunbookWorkerId        pulumi.StringInput `pulumi:"hybridRunbookWorkerId"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHybridRunbookWorkerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridRunbookWorkerArgs)(nil)).Elem()
}


type LookupHybridRunbookWorkerResultOutput struct{ *pulumi.OutputState }

func (LookupHybridRunbookWorkerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridRunbookWorkerResult)(nil)).Elem()
}

func (o LookupHybridRunbookWorkerResultOutput) ToLookupHybridRunbookWorkerResultOutput() LookupHybridRunbookWorkerResultOutput {
	return o
}

func (o LookupHybridRunbookWorkerResultOutput) ToLookupHybridRunbookWorkerResultOutputWithContext(ctx context.Context) LookupHybridRunbookWorkerResultOutput {
	return o
}

func (o LookupHybridRunbookWorkerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) LastSeenDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) *string { return v.LastSeenDateTime }).(pulumi.StringPtrOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) RegisteredDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) *string { return v.RegisteredDateTime }).(pulumi.StringPtrOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) *string { return v.VmResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) WorkerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) *string { return v.WorkerName }).(pulumi.StringPtrOutput)
}

func (o LookupHybridRunbookWorkerResultOutput) WorkerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerResult) *string { return v.WorkerType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridRunbookWorkerResultOutput{})
}
