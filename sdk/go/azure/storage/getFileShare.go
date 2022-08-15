


package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileShare(ctx *pulumi.Context, args *LookupFileShareArgs, opts ...pulumi.InvokeOption) (*LookupFileShareResult, error) {
	var rv LookupFileShareResult
	err := ctx.Invoke("azure-native:storage:getFileShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileShareArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ShareName         string  `pulumi:"shareName"`
}


type LookupFileShareResult struct {
	AccessTier             *string           `pulumi:"accessTier"`
	AccessTierChangeTime   string            `pulumi:"accessTierChangeTime"`
	AccessTierStatus       string            `pulumi:"accessTierStatus"`
	Deleted                bool              `pulumi:"deleted"`
	DeletedTime            string            `pulumi:"deletedTime"`
	EnabledProtocols       *string           `pulumi:"enabledProtocols"`
	Etag                   string            `pulumi:"etag"`
	Id                     string            `pulumi:"id"`
	LastModifiedTime       string            `pulumi:"lastModifiedTime"`
	Metadata               map[string]string `pulumi:"metadata"`
	Name                   string            `pulumi:"name"`
	RemainingRetentionDays int               `pulumi:"remainingRetentionDays"`
	RootSquash             *string           `pulumi:"rootSquash"`
	ShareQuota             *int              `pulumi:"shareQuota"`
	ShareUsageBytes        float64           `pulumi:"shareUsageBytes"`
	SnapshotTime           string            `pulumi:"snapshotTime"`
	Type                   string            `pulumi:"type"`
	Version                string            `pulumi:"version"`
}

func LookupFileShareOutput(ctx *pulumi.Context, args LookupFileShareOutputArgs, opts ...pulumi.InvokeOption) LookupFileShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileShareResult, error) {
			args := v.(LookupFileShareArgs)
			r, err := LookupFileShare(ctx, &args, opts...)
			var s LookupFileShareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileShareResultOutput)
}

type LookupFileShareOutputArgs struct {
	AccountName       pulumi.StringInput    `pulumi:"accountName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput    `pulumi:"shareName"`
}

func (LookupFileShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileShareArgs)(nil)).Elem()
}


type LookupFileShareResultOutput struct{ *pulumi.OutputState }

func (LookupFileShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileShareResult)(nil)).Elem()
}

func (o LookupFileShareResultOutput) ToLookupFileShareResultOutput() LookupFileShareResultOutput {
	return o
}

func (o LookupFileShareResultOutput) ToLookupFileShareResultOutputWithContext(ctx context.Context) LookupFileShareResultOutput {
	return o
}

func (o LookupFileShareResultOutput) AccessTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *string { return v.AccessTier }).(pulumi.StringPtrOutput)
}

func (o LookupFileShareResultOutput) AccessTierChangeTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.AccessTierChangeTime }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) AccessTierStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.AccessTierStatus }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFileShareResult) bool { return v.Deleted }).(pulumi.BoolOutput)
}

func (o LookupFileShareResultOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.DeletedTime }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) EnabledProtocols() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *string { return v.EnabledProtocols }).(pulumi.StringPtrOutput)
}

func (o LookupFileShareResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFileShareResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupFileShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) RemainingRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFileShareResult) int { return v.RemainingRetentionDays }).(pulumi.IntOutput)
}

func (o LookupFileShareResultOutput) RootSquash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *string { return v.RootSquash }).(pulumi.StringPtrOutput)
}

func (o LookupFileShareResultOutput) ShareQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *int { return v.ShareQuota }).(pulumi.IntPtrOutput)
}

func (o LookupFileShareResultOutput) ShareUsageBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupFileShareResult) float64 { return v.ShareUsageBytes }).(pulumi.Float64Output)
}

func (o LookupFileShareResultOutput) SnapshotTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.SnapshotTime }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileShareResultOutput{})
}
