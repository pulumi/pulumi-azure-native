


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRestorePoint(ctx *pulumi.Context, args *LookupRestorePointArgs, opts ...pulumi.InvokeOption) (*LookupRestorePointResult, error) {
	var rv LookupRestorePointResult
	err := ctx.Invoke("azure-native:compute/v20220301:getRestorePoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestorePointArgs struct {
	Expand                     *string `pulumi:"expand"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	RestorePointCollectionName string  `pulumi:"restorePointCollectionName"`
	RestorePointName           string  `pulumi:"restorePointName"`
}


type LookupRestorePointResult struct {
	ConsistencyMode    *string                            `pulumi:"consistencyMode"`
	ExcludeDisks       []ApiEntityReferenceResponse       `pulumi:"excludeDisks"`
	Id                 string                             `pulumi:"id"`
	InstanceView       RestorePointInstanceViewResponse   `pulumi:"instanceView"`
	Name               string                             `pulumi:"name"`
	ProvisioningState  string                             `pulumi:"provisioningState"`
	SourceMetadata     RestorePointSourceMetadataResponse `pulumi:"sourceMetadata"`
	SourceRestorePoint *ApiEntityReferenceResponse        `pulumi:"sourceRestorePoint"`
	TimeCreated        *string                            `pulumi:"timeCreated"`
	Type               string                             `pulumi:"type"`
}

func LookupRestorePointOutput(ctx *pulumi.Context, args LookupRestorePointOutputArgs, opts ...pulumi.InvokeOption) LookupRestorePointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRestorePointResult, error) {
			args := v.(LookupRestorePointArgs)
			r, err := LookupRestorePoint(ctx, &args, opts...)
			var s LookupRestorePointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRestorePointResultOutput)
}

type LookupRestorePointOutputArgs struct {
	Expand                     pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName          pulumi.StringInput    `pulumi:"resourceGroupName"`
	RestorePointCollectionName pulumi.StringInput    `pulumi:"restorePointCollectionName"`
	RestorePointName           pulumi.StringInput    `pulumi:"restorePointName"`
}

func (LookupRestorePointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointArgs)(nil)).Elem()
}


type LookupRestorePointResultOutput struct{ *pulumi.OutputState }

func (LookupRestorePointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointResult)(nil)).Elem()
}

func (o LookupRestorePointResultOutput) ToLookupRestorePointResultOutput() LookupRestorePointResultOutput {
	return o
}

func (o LookupRestorePointResultOutput) ToLookupRestorePointResultOutputWithContext(ctx context.Context) LookupRestorePointResultOutput {
	return o
}

func (o LookupRestorePointResultOutput) ConsistencyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestorePointResult) *string { return v.ConsistencyMode }).(pulumi.StringPtrOutput)
}

func (o LookupRestorePointResultOutput) ExcludeDisks() ApiEntityReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRestorePointResult) []ApiEntityReferenceResponse { return v.ExcludeDisks }).(ApiEntityReferenceResponseArrayOutput)
}

func (o LookupRestorePointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRestorePointResultOutput) InstanceView() RestorePointInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupRestorePointResult) RestorePointInstanceViewResponse { return v.InstanceView }).(RestorePointInstanceViewResponseOutput)
}

func (o LookupRestorePointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRestorePointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRestorePointResultOutput) SourceMetadata() RestorePointSourceMetadataResponseOutput {
	return o.ApplyT(func(v LookupRestorePointResult) RestorePointSourceMetadataResponse { return v.SourceMetadata }).(RestorePointSourceMetadataResponseOutput)
}

func (o LookupRestorePointResultOutput) SourceRestorePoint() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupRestorePointResult) *ApiEntityReferenceResponse { return v.SourceRestorePoint }).(ApiEntityReferenceResponsePtrOutput)
}

func (o LookupRestorePointResultOutput) TimeCreated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestorePointResult) *string { return v.TimeCreated }).(pulumi.StringPtrOutput)
}

func (o LookupRestorePointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestorePointResultOutput{})
}
