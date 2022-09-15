


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccountCredential(ctx *pulumi.Context, args *LookupStorageAccountCredentialArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountCredentialResult, error) {
	var rv LookupStorageAccountCredentialResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getStorageAccountCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountCredentialArgs struct {
	ManagerName                  string `pulumi:"managerName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
	StorageAccountCredentialName string `pulumi:"storageAccountCredentialName"`
}


type LookupStorageAccountCredentialResult struct {
	AccessKey    *AsymmetricEncryptedSecretResponse `pulumi:"accessKey"`
	EndPoint     string                             `pulumi:"endPoint"`
	Id           string                             `pulumi:"id"`
	Kind         *string                            `pulumi:"kind"`
	Name         string                             `pulumi:"name"`
	SslStatus    string                             `pulumi:"sslStatus"`
	Type         string                             `pulumi:"type"`
	VolumesCount int                                `pulumi:"volumesCount"`
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
	ManagerName                  pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageAccountCredentialName pulumi.StringInput `pulumi:"storageAccountCredentialName"`
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

func (o LookupStorageAccountCredentialResultOutput) EndPoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.EndPoint }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) SslStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.SslStatus }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupStorageAccountCredentialResultOutput) VolumesCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) int { return v.VolumesCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountCredentialResultOutput{})
}
