


package v20190301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSyncGroup(ctx *pulumi.Context, args *LookupSyncGroupArgs, opts ...pulumi.InvokeOption) (*LookupSyncGroupResult, error) {
	var rv LookupSyncGroupResult
	err := ctx.Invoke("azure-native:storagesync/v20190301:getSyncGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncGroupArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	SyncGroupName          string `pulumi:"syncGroupName"`
}


type LookupSyncGroupResult struct {
	Id              string  `pulumi:"id"`
	Name            string  `pulumi:"name"`
	SyncGroupStatus string  `pulumi:"syncGroupStatus"`
	Type            string  `pulumi:"type"`
	UniqueId        *string `pulumi:"uniqueId"`
}

func LookupSyncGroupOutput(ctx *pulumi.Context, args LookupSyncGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSyncGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSyncGroupResult, error) {
			args := v.(LookupSyncGroupArgs)
			r, err := LookupSyncGroup(ctx, &args, opts...)
			var s LookupSyncGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSyncGroupResultOutput)
}

type LookupSyncGroupOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
	SyncGroupName          pulumi.StringInput `pulumi:"syncGroupName"`
}

func (LookupSyncGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncGroupArgs)(nil)).Elem()
}


type LookupSyncGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSyncGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncGroupResult)(nil)).Elem()
}

func (o LookupSyncGroupResultOutput) ToLookupSyncGroupResultOutput() LookupSyncGroupResultOutput {
	return o
}

func (o LookupSyncGroupResultOutput) ToLookupSyncGroupResultOutputWithContext(ctx context.Context) LookupSyncGroupResultOutput {
	return o
}

func (o LookupSyncGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSyncGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSyncGroupResultOutput) SyncGroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.SyncGroupStatus }).(pulumi.StringOutput)
}

func (o LookupSyncGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSyncGroupResultOutput) UniqueId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *string { return v.UniqueId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncGroupResultOutput{})
}
