


package v20210901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobServiceProperties(ctx *pulumi.Context, args *LookupBlobServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupBlobServicePropertiesResult, error) {
	var rv LookupBlobServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20210901:getBlobServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	BlobServicesName  string `pulumi:"blobServicesName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBlobServicePropertiesResult struct {
	AutomaticSnapshotPolicyEnabled *bool                                 `pulumi:"automaticSnapshotPolicyEnabled"`
	ChangeFeed                     *ChangeFeedResponse                   `pulumi:"changeFeed"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicyResponse        `pulumi:"containerDeleteRetentionPolicy"`
	Cors                           *CorsRulesResponse                    `pulumi:"cors"`
	DefaultServiceVersion          *string                               `pulumi:"defaultServiceVersion"`
	DeleteRetentionPolicy          *DeleteRetentionPolicyResponse        `pulumi:"deleteRetentionPolicy"`
	Id                             string                                `pulumi:"id"`
	IsVersioningEnabled            *bool                                 `pulumi:"isVersioningEnabled"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicyResponse `pulumi:"lastAccessTimeTrackingPolicy"`
	Name                           string                                `pulumi:"name"`
	RestorePolicy                  *RestorePolicyPropertiesResponse      `pulumi:"restorePolicy"`
	Sku                            SkuResponse                           `pulumi:"sku"`
	Type                           string                                `pulumi:"type"`
}

func LookupBlobServicePropertiesOutput(ctx *pulumi.Context, args LookupBlobServicePropertiesOutputArgs, opts ...pulumi.InvokeOption) LookupBlobServicePropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobServicePropertiesResult, error) {
			args := v.(LookupBlobServicePropertiesArgs)
			r, err := LookupBlobServiceProperties(ctx, &args, opts...)
			var s LookupBlobServicePropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobServicePropertiesResultOutput)
}

type LookupBlobServicePropertiesOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	BlobServicesName  pulumi.StringInput `pulumi:"blobServicesName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobServicePropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobServicePropertiesArgs)(nil)).Elem()
}


type LookupBlobServicePropertiesResultOutput struct{ *pulumi.OutputState }

func (LookupBlobServicePropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobServicePropertiesResult)(nil)).Elem()
}

func (o LookupBlobServicePropertiesResultOutput) ToLookupBlobServicePropertiesResultOutput() LookupBlobServicePropertiesResultOutput {
	return o
}

func (o LookupBlobServicePropertiesResultOutput) ToLookupBlobServicePropertiesResultOutputWithContext(ctx context.Context) LookupBlobServicePropertiesResultOutput {
	return o
}

func (o LookupBlobServicePropertiesResultOutput) AutomaticSnapshotPolicyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *bool { return v.AutomaticSnapshotPolicyEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) ChangeFeed() ChangeFeedResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *ChangeFeedResponse { return v.ChangeFeed }).(ChangeFeedResponsePtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) ContainerDeleteRetentionPolicy() DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *DeleteRetentionPolicyResponse {
		return v.ContainerDeleteRetentionPolicy
	}).(DeleteRetentionPolicyResponsePtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) DefaultServiceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *string { return v.DefaultServiceVersion }).(pulumi.StringPtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) DeleteRetentionPolicy() DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *DeleteRetentionPolicyResponse {
		return v.DeleteRetentionPolicy
	}).(DeleteRetentionPolicyResponsePtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobServicePropertiesResultOutput) IsVersioningEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *bool { return v.IsVersioningEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) LastAccessTimeTrackingPolicy() LastAccessTimeTrackingPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *LastAccessTimeTrackingPolicyResponse {
		return v.LastAccessTimeTrackingPolicy
	}).(LastAccessTimeTrackingPolicyResponsePtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobServicePropertiesResultOutput) RestorePolicy() RestorePolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *RestorePolicyPropertiesResponse { return v.RestorePolicy }).(RestorePolicyPropertiesResponsePtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupBlobServicePropertiesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobServicePropertiesResultOutput{})
}
