


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncGroup(ctx *pulumi.Context, args *LookupSyncGroupArgs, opts ...pulumi.InvokeOption) (*LookupSyncGroupResult, error) {
	var rv LookupSyncGroupResult
	err := ctx.Invoke("azure-native:sql/v20150501preview:getSyncGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncGroupArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	SyncGroupName     string `pulumi:"syncGroupName"`
}


type LookupSyncGroupResult struct {
	ConflictResolutionPolicy *string                  `pulumi:"conflictResolutionPolicy"`
	HubDatabaseUserName      *string                  `pulumi:"hubDatabaseUserName"`
	Id                       string                   `pulumi:"id"`
	Interval                 *int                     `pulumi:"interval"`
	LastSyncTime             string                   `pulumi:"lastSyncTime"`
	Name                     string                   `pulumi:"name"`
	Schema                   *SyncGroupSchemaResponse `pulumi:"schema"`
	SyncDatabaseId           *string                  `pulumi:"syncDatabaseId"`
	SyncState                string                   `pulumi:"syncState"`
	Type                     string                   `pulumi:"type"`
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
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
	SyncGroupName     pulumi.StringInput `pulumi:"syncGroupName"`
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

func (o LookupSyncGroupResultOutput) ConflictResolutionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *string { return v.ConflictResolutionPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupSyncGroupResultOutput) HubDatabaseUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *string { return v.HubDatabaseUserName }).(pulumi.StringPtrOutput)
}

func (o LookupSyncGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSyncGroupResultOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o LookupSyncGroupResultOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

func (o LookupSyncGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSyncGroupResultOutput) Schema() SyncGroupSchemaResponsePtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *SyncGroupSchemaResponse { return v.Schema }).(SyncGroupSchemaResponsePtrOutput)
}

func (o LookupSyncGroupResultOutput) SyncDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *string { return v.SyncDatabaseId }).(pulumi.StringPtrOutput)
}

func (o LookupSyncGroupResultOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.SyncState }).(pulumi.StringOutput)
}

func (o LookupSyncGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncGroupResultOutput{})
}
