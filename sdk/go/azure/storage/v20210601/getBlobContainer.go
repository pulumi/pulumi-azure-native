


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobContainer(ctx *pulumi.Context, args *LookupBlobContainerArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerResult, error) {
	var rv LookupBlobContainerResult
	err := ctx.Invoke("azure-native:storage/v20210601:getBlobContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ContainerName     string `pulumi:"containerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBlobContainerResult struct {
	DefaultEncryptionScope         *string                                 `pulumi:"defaultEncryptionScope"`
	Deleted                        bool                                    `pulumi:"deleted"`
	DeletedTime                    string                                  `pulumi:"deletedTime"`
	DenyEncryptionScopeOverride    *bool                                   `pulumi:"denyEncryptionScopeOverride"`
	EnableNfsV3AllSquash           *bool                                   `pulumi:"enableNfsV3AllSquash"`
	EnableNfsV3RootSquash          *bool                                   `pulumi:"enableNfsV3RootSquash"`
	Etag                           string                                  `pulumi:"etag"`
	HasImmutabilityPolicy          bool                                    `pulumi:"hasImmutabilityPolicy"`
	HasLegalHold                   bool                                    `pulumi:"hasLegalHold"`
	Id                             string                                  `pulumi:"id"`
	ImmutabilityPolicy             ImmutabilityPolicyPropertiesResponse    `pulumi:"immutabilityPolicy"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioningResponse `pulumi:"immutableStorageWithVersioning"`
	LastModifiedTime               string                                  `pulumi:"lastModifiedTime"`
	LeaseDuration                  string                                  `pulumi:"leaseDuration"`
	LeaseState                     string                                  `pulumi:"leaseState"`
	LeaseStatus                    string                                  `pulumi:"leaseStatus"`
	LegalHold                      LegalHoldPropertiesResponse             `pulumi:"legalHold"`
	Metadata                       map[string]string                       `pulumi:"metadata"`
	Name                           string                                  `pulumi:"name"`
	PublicAccess                   *string                                 `pulumi:"publicAccess"`
	RemainingRetentionDays         int                                     `pulumi:"remainingRetentionDays"`
	Type                           string                                  `pulumi:"type"`
	Version                        string                                  `pulumi:"version"`
}

func LookupBlobContainerOutput(ctx *pulumi.Context, args LookupBlobContainerOutputArgs, opts ...pulumi.InvokeOption) LookupBlobContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobContainerResult, error) {
			args := v.(LookupBlobContainerArgs)
			r, err := LookupBlobContainer(ctx, &args, opts...)
			var s LookupBlobContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobContainerResultOutput)
}

type LookupBlobContainerOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ContainerName     pulumi.StringInput `pulumi:"containerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerArgs)(nil)).Elem()
}


type LookupBlobContainerResultOutput struct{ *pulumi.OutputState }

func (LookupBlobContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerResult)(nil)).Elem()
}

func (o LookupBlobContainerResultOutput) ToLookupBlobContainerResultOutput() LookupBlobContainerResultOutput {
	return o
}

func (o LookupBlobContainerResultOutput) ToLookupBlobContainerResultOutputWithContext(ctx context.Context) LookupBlobContainerResultOutput {
	return o
}

func (o LookupBlobContainerResultOutput) DefaultEncryptionScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *string { return v.DefaultEncryptionScope }).(pulumi.StringPtrOutput)
}

func (o LookupBlobContainerResultOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) bool { return v.Deleted }).(pulumi.BoolOutput)
}

func (o LookupBlobContainerResultOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.DeletedTime }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) DenyEncryptionScopeOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *bool { return v.DenyEncryptionScopeOverride }).(pulumi.BoolPtrOutput)
}

func (o LookupBlobContainerResultOutput) EnableNfsV3AllSquash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *bool { return v.EnableNfsV3AllSquash }).(pulumi.BoolPtrOutput)
}

func (o LookupBlobContainerResultOutput) EnableNfsV3RootSquash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *bool { return v.EnableNfsV3RootSquash }).(pulumi.BoolPtrOutput)
}

func (o LookupBlobContainerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) HasImmutabilityPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) bool { return v.HasImmutabilityPolicy }).(pulumi.BoolOutput)
}

func (o LookupBlobContainerResultOutput) HasLegalHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) bool { return v.HasLegalHold }).(pulumi.BoolOutput)
}

func (o LookupBlobContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) ImmutabilityPolicy() ImmutabilityPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) ImmutabilityPolicyPropertiesResponse { return v.ImmutabilityPolicy }).(ImmutabilityPolicyPropertiesResponseOutput)
}

func (o LookupBlobContainerResultOutput) ImmutableStorageWithVersioning() ImmutableStorageWithVersioningResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *ImmutableStorageWithVersioningResponse {
		return v.ImmutableStorageWithVersioning
	}).(ImmutableStorageWithVersioningResponsePtrOutput)
}

func (o LookupBlobContainerResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) LeaseDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.LeaseDuration }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) LeaseState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.LeaseState }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) LeaseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.LeaseStatus }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) LegalHold() LegalHoldPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) LegalHoldPropertiesResponse { return v.LegalHold }).(LegalHoldPropertiesResponseOutput)
}

func (o LookupBlobContainerResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupBlobContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) PublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) *string { return v.PublicAccess }).(pulumi.StringPtrOutput)
}

func (o LookupBlobContainerResultOutput) RemainingRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) int { return v.RemainingRetentionDays }).(pulumi.IntOutput)
}

func (o LookupBlobContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBlobContainerResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobContainerResultOutput{})
}
