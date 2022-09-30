


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupStorageAccountCredential(ctx *pulumi.Context, args *LookupStorageAccountCredentialArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountCredentialResult, error) {
	var rv LookupStorageAccountCredentialResult
	err := ctx.Invoke("azure-native:databoxedge/v20190801:getStorageAccountCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountCredentialArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStorageAccountCredentialResult struct {
	AccountKey       *AsymmetricEncryptedSecretResponse `pulumi:"accountKey"`
	AccountType      string                             `pulumi:"accountType"`
	Alias            string                             `pulumi:"alias"`
	BlobDomainName   *string                            `pulumi:"blobDomainName"`
	ConnectionString *string                            `pulumi:"connectionString"`
	Id               string                             `pulumi:"id"`
	Name             string                             `pulumi:"name"`
	SslStatus        string                             `pulumi:"sslStatus"`
	StorageAccountId *string                            `pulumi:"storageAccountId"`
	Type             string                             `pulumi:"type"`
	UserName         *string                            `pulumi:"userName"`
}

func LookupStorageAccountCredentialOutput(ctx *pulumi.Context, args LookupStorageAccountCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupStorageAccountCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageAccountCredentialResult, error) {
			args := v.(LookupStorageAccountCredentialArgs)
			r, err := LookupStorageAccountCredential(ctx, &args, opts...)
			var s LookupStorageAccountCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageAccountCredentialResultOutput)
}

type LookupStorageAccountCredentialOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStorageAccountCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountCredentialArgs)(nil)).Elem()
}


type LookupStorageAccountCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupStorageAccountCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountCredentialResult)(nil)).Elem()
}

func (o LookupStorageAccountCredentialResultOutput) ToLookupStorageAccountCredentialResultOutput() LookupStorageAccountCredentialResultOutput {
	return o
}

func (o LookupStorageAccountCredentialResultOutput) ToLookupStorageAccountCredentialResultOutputWithContext(ctx context.Context) LookupStorageAccountCredentialResultOutput {
	return o
}

func (o LookupStorageAccountCredentialResultOutput) AccountKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *AsymmetricEncryptedSecretResponse { return v.AccountKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o LookupStorageAccountCredentialResultOutput) AccountType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.AccountType }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Alias }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) BlobDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.BlobDomainName }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountCredentialResultOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) SslStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.SslStatus }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountCredentialResultOutput{})
}
