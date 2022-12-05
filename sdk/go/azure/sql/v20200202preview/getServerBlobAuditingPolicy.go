


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerBlobAuditingPolicy(ctx *pulumi.Context, args *LookupServerBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServerBlobAuditingPolicyResult, error) {
	var rv LookupServerBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql/v20200202preview:getServerBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerBlobAuditingPolicyArgs struct {
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
}


type LookupServerBlobAuditingPolicyResult struct {
	AuditActionsAndGroups        []string `pulumi:"auditActionsAndGroups"`
	Id                           string   `pulumi:"id"`
	IsAzureMonitorTargetEnabled  *bool    `pulumi:"isAzureMonitorTargetEnabled"`
	IsDevopsAuditEnabled         *bool    `pulumi:"isDevopsAuditEnabled"`
	IsStorageSecondaryKeyInUse   *bool    `pulumi:"isStorageSecondaryKeyInUse"`
	Name                         string   `pulumi:"name"`
	QueueDelayMs                 *int     `pulumi:"queueDelayMs"`
	RetentionDays                *int     `pulumi:"retentionDays"`
	State                        string   `pulumi:"state"`
	StorageAccountSubscriptionId *string  `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              *string  `pulumi:"storageEndpoint"`
	Type                         string   `pulumi:"type"`
}

func LookupServerBlobAuditingPolicyOutput(ctx *pulumi.Context, args LookupServerBlobAuditingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupServerBlobAuditingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerBlobAuditingPolicyResult, error) {
			args := v.(LookupServerBlobAuditingPolicyArgs)
			r, err := LookupServerBlobAuditingPolicy(ctx, &args, opts...)
			var s LookupServerBlobAuditingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerBlobAuditingPolicyResultOutput)
}

type LookupServerBlobAuditingPolicyOutputArgs struct {
	BlobAuditingPolicyName pulumi.StringInput `pulumi:"blobAuditingPolicyName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName             pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerBlobAuditingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerBlobAuditingPolicyArgs)(nil)).Elem()
}


type LookupServerBlobAuditingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupServerBlobAuditingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerBlobAuditingPolicyResult)(nil)).Elem()
}

func (o LookupServerBlobAuditingPolicyResultOutput) ToLookupServerBlobAuditingPolicyResultOutput() LookupServerBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupServerBlobAuditingPolicyResultOutput) ToLookupServerBlobAuditingPolicyResultOutputWithContext(ctx context.Context) LookupServerBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupServerBlobAuditingPolicyResultOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) []string { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *bool { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) IsDevopsAuditEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *bool { return v.IsDevopsAuditEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *bool { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *int { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *string { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupServerBlobAuditingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerBlobAuditingPolicyResultOutput{})
}
