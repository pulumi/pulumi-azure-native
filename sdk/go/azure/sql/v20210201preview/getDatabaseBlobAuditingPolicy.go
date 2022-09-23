


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseBlobAuditingPolicy(ctx *pulumi.Context, args *LookupDatabaseBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseBlobAuditingPolicyResult, error) {
	var rv LookupDatabaseBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getDatabaseBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseBlobAuditingPolicyArgs struct {
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	DatabaseName           string `pulumi:"databaseName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
}


type LookupDatabaseBlobAuditingPolicyResult struct {
	AuditActionsAndGroups        []string `pulumi:"auditActionsAndGroups"`
	Id                           string   `pulumi:"id"`
	IsAzureMonitorTargetEnabled  *bool    `pulumi:"isAzureMonitorTargetEnabled"`
	IsStorageSecondaryKeyInUse   *bool    `pulumi:"isStorageSecondaryKeyInUse"`
	Kind                         string   `pulumi:"kind"`
	Name                         string   `pulumi:"name"`
	QueueDelayMs                 *int     `pulumi:"queueDelayMs"`
	RetentionDays                *int     `pulumi:"retentionDays"`
	State                        string   `pulumi:"state"`
	StorageAccountSubscriptionId *string  `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              *string  `pulumi:"storageEndpoint"`
	Type                         string   `pulumi:"type"`
}

func LookupDatabaseBlobAuditingPolicyOutput(ctx *pulumi.Context, args LookupDatabaseBlobAuditingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseBlobAuditingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseBlobAuditingPolicyResult, error) {
			args := v.(LookupDatabaseBlobAuditingPolicyArgs)
			r, err := LookupDatabaseBlobAuditingPolicy(ctx, &args, opts...)
			var s LookupDatabaseBlobAuditingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseBlobAuditingPolicyResultOutput)
}

type LookupDatabaseBlobAuditingPolicyOutputArgs struct {
	BlobAuditingPolicyName pulumi.StringInput `pulumi:"blobAuditingPolicyName"`
	DatabaseName           pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName             pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseBlobAuditingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseBlobAuditingPolicyArgs)(nil)).Elem()
}


type LookupDatabaseBlobAuditingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseBlobAuditingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseBlobAuditingPolicyResult)(nil)).Elem()
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) ToLookupDatabaseBlobAuditingPolicyResultOutput() LookupDatabaseBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) ToLookupDatabaseBlobAuditingPolicyResultOutputWithContext(ctx context.Context) LookupDatabaseBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) []string { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) *bool { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) *bool { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) *int { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) *string { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseBlobAuditingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseBlobAuditingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseBlobAuditingPolicyResultOutput{})
}
