


package v20180330preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAsset(ctx *pulumi.Context, args *LookupAssetArgs, opts ...pulumi.InvokeOption) (*LookupAssetResult, error) {
	var rv LookupAssetResult
	err := ctx.Invoke("azure-native:media/v20180330preview:getAsset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAssetResult struct {
	AlternateId             *string `pulumi:"alternateId"`
	AssetId                 string  `pulumi:"assetId"`
	Container               *string `pulumi:"container"`
	Created                 string  `pulumi:"created"`
	Description             *string `pulumi:"description"`
	Id                      string  `pulumi:"id"`
	LastModified            string  `pulumi:"lastModified"`
	Name                    string  `pulumi:"name"`
	StorageAccountName      *string `pulumi:"storageAccountName"`
	StorageEncryptionFormat string  `pulumi:"storageEncryptionFormat"`
	Type                    string  `pulumi:"type"`
}

func LookupAssetOutput(ctx *pulumi.Context, args LookupAssetOutputArgs, opts ...pulumi.InvokeOption) LookupAssetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssetResult, error) {
			args := v.(LookupAssetArgs)
			r, err := LookupAsset(ctx, &args, opts...)
			var s LookupAssetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssetResultOutput)
}

type LookupAssetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	AssetName         pulumi.StringInput `pulumi:"assetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAssetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetArgs)(nil)).Elem()
}


type LookupAssetResultOutput struct{ *pulumi.OutputState }

func (LookupAssetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetResult)(nil)).Elem()
}

func (o LookupAssetResultOutput) ToLookupAssetResultOutput() LookupAssetResultOutput {
	return o
}

func (o LookupAssetResultOutput) ToLookupAssetResultOutputWithContext(ctx context.Context) LookupAssetResultOutput {
	return o
}

func (o LookupAssetResultOutput) AlternateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.AlternateId }).(pulumi.StringPtrOutput)
}

func (o LookupAssetResultOutput) AssetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetResult) string { return v.AssetId }).(pulumi.StringOutput)
}

func (o LookupAssetResultOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o LookupAssetResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupAssetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAssetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssetResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupAssetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssetResultOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

func (o LookupAssetResultOutput) StorageEncryptionFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetResult) string { return v.StorageEncryptionFormat }).(pulumi.StringOutput)
}

func (o LookupAssetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssetResultOutput{})
}
