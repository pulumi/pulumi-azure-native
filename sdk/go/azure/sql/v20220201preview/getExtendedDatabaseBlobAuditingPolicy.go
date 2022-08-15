


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtendedDatabaseBlobAuditingPolicy(ctx *pulumi.Context, args *LookupExtendedDatabaseBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupExtendedDatabaseBlobAuditingPolicyResult, error) {
	var rv LookupExtendedDatabaseBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql/v20220201preview:getExtendedDatabaseBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtendedDatabaseBlobAuditingPolicyArgs struct {
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	DatabaseName           string `pulumi:"databaseName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
}


type LookupExtendedDatabaseBlobAuditingPolicyResult struct {
	AuditActionsAndGroups        []string `pulumi:"auditActionsAndGroups"`
	Id                           string   `pulumi:"id"`
	IsAzureMonitorTargetEnabled  *bool    `pulumi:"isAzureMonitorTargetEnabled"`
	IsManagedIdentityInUse       *bool    `pulumi:"isManagedIdentityInUse"`
	IsStorageSecondaryKeyInUse   *bool    `pulumi:"isStorageSecondaryKeyInUse"`
	Name                         string   `pulumi:"name"`
	PredicateExpression          *string  `pulumi:"predicateExpression"`
	QueueDelayMs                 *int     `pulumi:"queueDelayMs"`
	RetentionDays                *int     `pulumi:"retentionDays"`
	State                        string   `pulumi:"state"`
	StorageAccountSubscriptionId *string  `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              *string  `pulumi:"storageEndpoint"`
	Type                         string   `pulumi:"type"`
}

func LookupExtendedDatabaseBlobAuditingPolicyOutput(ctx *pulumi.Context, args LookupExtendedDatabaseBlobAuditingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupExtendedDatabaseBlobAuditingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtendedDatabaseBlobAuditingPolicyResult, error) {
			args := v.(LookupExtendedDatabaseBlobAuditingPolicyArgs)
			r, err := LookupExtendedDatabaseBlobAuditingPolicy(ctx, &args, opts...)
			var s LookupExtendedDatabaseBlobAuditingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtendedDatabaseBlobAuditingPolicyResultOutput)
}

type LookupExtendedDatabaseBlobAuditingPolicyOutputArgs struct {
	BlobAuditingPolicyName pulumi.StringInput `pulumi:"blobAuditingPolicyName"`
	DatabaseName           pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName             pulumi.StringInput `pulumi:"serverName"`
}

func (LookupExtendedDatabaseBlobAuditingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtendedDatabaseBlobAuditingPolicyArgs)(nil)).Elem()
}


type LookupExtendedDatabaseBlobAuditingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupExtendedDatabaseBlobAuditingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtendedDatabaseBlobAuditingPolicyResult)(nil)).Elem()
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) ToLookupExtendedDatabaseBlobAuditingPolicyResultOutput() LookupExtendedDatabaseBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) ToLookupExtendedDatabaseBlobAuditingPolicyResultOutputWithContext(ctx context.Context) LookupExtendedDatabaseBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) []string { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) *bool { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) IsManagedIdentityInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) *bool { return v.IsManagedIdentityInUse }).(pulumi.BoolPtrOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) *bool { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) PredicateExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) *string { return v.PredicateExpression }).(pulumi.StringPtrOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) *int { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) *string { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupExtendedDatabaseBlobAuditingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedDatabaseBlobAuditingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtendedDatabaseBlobAuditingPolicyResultOutput{})
}
