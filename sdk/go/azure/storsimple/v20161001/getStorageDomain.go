


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupStorageDomain(ctx *pulumi.Context, args *LookupStorageDomainArgs, opts ...pulumi.InvokeOption) (*LookupStorageDomainResult, error) {
	var rv LookupStorageDomainResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getStorageDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageDomainArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageDomainName string `pulumi:"storageDomainName"`
}


type LookupStorageDomainResult struct {
	EncryptionKey               *AsymmetricEncryptedSecretResponse `pulumi:"encryptionKey"`
	EncryptionStatus            string                             `pulumi:"encryptionStatus"`
	Id                          string                             `pulumi:"id"`
	Name                        string                             `pulumi:"name"`
	StorageAccountCredentialIds []string                           `pulumi:"storageAccountCredentialIds"`
	Type                        string                             `pulumi:"type"`
}

func LookupStorageDomainOutput(ctx *pulumi.Context, args LookupStorageDomainOutputArgs, opts ...pulumi.InvokeOption) LookupStorageDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageDomainResult, error) {
			args := v.(LookupStorageDomainArgs)
			r, err := LookupStorageDomain(ctx, &args, opts...)
			var s LookupStorageDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageDomainResultOutput)
}

type LookupStorageDomainOutputArgs struct {
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageDomainName pulumi.StringInput `pulumi:"storageDomainName"`
}

func (LookupStorageDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageDomainArgs)(nil)).Elem()
}


type LookupStorageDomainResultOutput struct{ *pulumi.OutputState }

func (LookupStorageDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageDomainResult)(nil)).Elem()
}

func (o LookupStorageDomainResultOutput) ToLookupStorageDomainResultOutput() LookupStorageDomainResultOutput {
	return o
}

func (o LookupStorageDomainResultOutput) ToLookupStorageDomainResultOutputWithContext(ctx context.Context) LookupStorageDomainResultOutput {
	return o
}

func (o LookupStorageDomainResultOutput) EncryptionKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageDomainResult) *AsymmetricEncryptedSecretResponse { return v.EncryptionKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o LookupStorageDomainResultOutput) EncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageDomainResult) string { return v.EncryptionStatus }).(pulumi.StringOutput)
}

func (o LookupStorageDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageDomainResultOutput) StorageAccountCredentialIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStorageDomainResult) []string { return v.StorageAccountCredentialIds }).(pulumi.StringArrayOutput)
}

func (o LookupStorageDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageDomainResultOutput{})
}
