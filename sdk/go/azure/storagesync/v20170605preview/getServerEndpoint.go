


package v20170605preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupServerEndpoint(ctx *pulumi.Context, args *LookupServerEndpointArgs, opts ...pulumi.InvokeOption) (*LookupServerEndpointResult, error) {
	var rv LookupServerEndpointResult
	err := ctx.Invoke("azure-native:storagesync/v20170605preview:getServerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerEndpointArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerEndpointName     string `pulumi:"serverEndpointName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	SyncGroupName          string `pulumi:"syncGroupName"`
}


type LookupServerEndpointResult struct {
	ByteProgress            *int    `pulumi:"byteProgress"`
	ByteTotal               *int    `pulumi:"byteTotal"`
	CloudTiering            *string `pulumi:"cloudTiering"`
	CurrentProgressType     *string `pulumi:"currentProgressType"`
	FriendlyName            *string `pulumi:"friendlyName"`
	Id                      string  `pulumi:"id"`
	ItemDownloadErrorCount  *int    `pulumi:"itemDownloadErrorCount"`
	ItemProgressCount       *int    `pulumi:"itemProgressCount"`
	ItemTotalCount          *int    `pulumi:"itemTotalCount"`
	ItemUploadErrorCount    *int    `pulumi:"itemUploadErrorCount"`
	LastSyncSuccess         *string `pulumi:"lastSyncSuccess"`
	LastWorkflowId          *string `pulumi:"lastWorkflowId"`
	Name                    string  `pulumi:"name"`
	ProvisioningState       *string `pulumi:"provisioningState"`
	ServerLocalPath         *string `pulumi:"serverLocalPath"`
	ServerResourceId        *string `pulumi:"serverResourceId"`
	SyncErrorContext        *string `pulumi:"syncErrorContext"`
	SyncErrorDirection      *string `pulumi:"syncErrorDirection"`
	SyncErrorState          *string `pulumi:"syncErrorState"`
	SyncErrorStateTimestamp *string `pulumi:"syncErrorStateTimestamp"`
	TotalProgress           *int    `pulumi:"totalProgress"`
	Type                    string  `pulumi:"type"`
	VolumeFreeSpacePercent  *int    `pulumi:"volumeFreeSpacePercent"`
}

func LookupServerEndpointOutput(ctx *pulumi.Context, args LookupServerEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupServerEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerEndpointResult, error) {
			args := v.(LookupServerEndpointArgs)
			r, err := LookupServerEndpoint(ctx, &args, opts...)
			var s LookupServerEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerEndpointResultOutput)
}

type LookupServerEndpointOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerEndpointName     pulumi.StringInput `pulumi:"serverEndpointName"`
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
	SyncGroupName          pulumi.StringInput `pulumi:"syncGroupName"`
}

func (LookupServerEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerEndpointArgs)(nil)).Elem()
}


type LookupServerEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupServerEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerEndpointResult)(nil)).Elem()
}

func (o LookupServerEndpointResultOutput) ToLookupServerEndpointResultOutput() LookupServerEndpointResultOutput {
	return o
}

func (o LookupServerEndpointResultOutput) ToLookupServerEndpointResultOutputWithContext(ctx context.Context) LookupServerEndpointResultOutput {
	return o
}

func (o LookupServerEndpointResultOutput) ByteProgress() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.ByteProgress }).(pulumi.IntPtrOutput)
}

func (o LookupServerEndpointResultOutput) ByteTotal() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.ByteTotal }).(pulumi.IntPtrOutput)
}

func (o LookupServerEndpointResultOutput) CloudTiering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.CloudTiering }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) CurrentProgressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.CurrentProgressType }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) ItemDownloadErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.ItemDownloadErrorCount }).(pulumi.IntPtrOutput)
}

func (o LookupServerEndpointResultOutput) ItemProgressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.ItemProgressCount }).(pulumi.IntPtrOutput)
}

func (o LookupServerEndpointResultOutput) ItemTotalCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.ItemTotalCount }).(pulumi.IntPtrOutput)
}

func (o LookupServerEndpointResultOutput) ItemUploadErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.ItemUploadErrorCount }).(pulumi.IntPtrOutput)
}

func (o LookupServerEndpointResultOutput) LastSyncSuccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.LastSyncSuccess }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) LastWorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.LastWorkflowId }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) ServerLocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.ServerLocalPath }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) ServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.ServerResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) SyncErrorContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.SyncErrorContext }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) SyncErrorDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.SyncErrorDirection }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) SyncErrorState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.SyncErrorState }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) SyncErrorStateTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.SyncErrorStateTimestamp }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) TotalProgress() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.TotalProgress }).(pulumi.IntPtrOutput)
}

func (o LookupServerEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) VolumeFreeSpacePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.VolumeFreeSpacePercent }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerEndpointResultOutput{})
}
