


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupStorageAccountCredential(ctx *pulumi.Context, args *LookupStorageAccountCredentialArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountCredentialResult, error) {
	var rv LookupStorageAccountCredentialResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getStorageAccountCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountCredentialArgs struct {
	CredentialName    string `pulumi:"credentialName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStorageAccountCredentialResult struct {
	AccessKey *AsymmetricEncryptedSecretResponse `pulumi:"accessKey"`
	CloudType string                             `pulumi:"cloudType"`
	EnableSSL string                             `pulumi:"enableSSL"`
	EndPoint  string                             `pulumi:"endPoint"`
	Id        string                             `pulumi:"id"`
	Location  *string                            `pulumi:"location"`
	Login     string                             `pulumi:"login"`
	Name      string                             `pulumi:"name"`
	Type      string                             `pulumi:"type"`
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
	CredentialName    pulumi.StringInput `pulumi:"credentialName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
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

func (o LookupStorageAccountCredentialResultOutput) AccessKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *AsymmetricEncryptedSecretResponse { return v.AccessKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o LookupStorageAccountCredentialResultOutput) CloudType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.CloudType }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) EnableSSL() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.EnableSSL }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) EndPoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.EndPoint }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Login }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountCredentialResultOutput{})
}
