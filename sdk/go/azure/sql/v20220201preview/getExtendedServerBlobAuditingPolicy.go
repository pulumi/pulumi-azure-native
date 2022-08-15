


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtendedServerBlobAuditingPolicy(ctx *pulumi.Context, args *LookupExtendedServerBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupExtendedServerBlobAuditingPolicyResult, error) {
	var rv LookupExtendedServerBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql/v20220201preview:getExtendedServerBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtendedServerBlobAuditingPolicyArgs struct {
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
}


type LookupExtendedServerBlobAuditingPolicyResult struct {
	AuditActionsAndGroups        []string `pulumi:"auditActionsAndGroups"`
	Id                           string   `pulumi:"id"`
	IsAzureMonitorTargetEnabled  *bool    `pulumi:"isAzureMonitorTargetEnabled"`
	IsDevopsAuditEnabled         *bool    `pulumi:"isDevopsAuditEnabled"`
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

func LookupExtendedServerBlobAuditingPolicyOutput(ctx *pulumi.Context, args LookupExtendedServerBlobAuditingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupExtendedServerBlobAuditingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtendedServerBlobAuditingPolicyResult, error) {
			args := v.(LookupExtendedServerBlobAuditingPolicyArgs)
			r, err := LookupExtendedServerBlobAuditingPolicy(ctx, &args, opts...)
			var s LookupExtendedServerBlobAuditingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtendedServerBlobAuditingPolicyResultOutput)
}

type LookupExtendedServerBlobAuditingPolicyOutputArgs struct {
	BlobAuditingPolicyName pulumi.StringInput `pulumi:"blobAuditingPolicyName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName             pulumi.StringInput `pulumi:"serverName"`
}

func (LookupExtendedServerBlobAuditingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtendedServerBlobAuditingPolicyArgs)(nil)).Elem()
}


type LookupExtendedServerBlobAuditingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupExtendedServerBlobAuditingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtendedServerBlobAuditingPolicyResult)(nil)).Elem()
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) ToLookupExtendedServerBlobAuditingPolicyResultOutput() LookupExtendedServerBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) ToLookupExtendedServerBlobAuditingPolicyResultOutputWithContext(ctx context.Context) LookupExtendedServerBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) []string { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *bool { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) IsDevopsAuditEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *bool { return v.IsDevopsAuditEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) IsManagedIdentityInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *bool { return v.IsManagedIdentityInUse }).(pulumi.BoolPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *bool { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) PredicateExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *string { return v.PredicateExpression }).(pulumi.StringPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *int { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *string { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtendedServerBlobAuditingPolicyResultOutput{})
}
