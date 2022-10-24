


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetManagerEncryptionKey(ctx *pulumi.Context, args *GetManagerEncryptionKeyArgs, opts ...pulumi.InvokeOption) (*GetManagerEncryptionKeyResult, error) {
	var rv GetManagerEncryptionKeyResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getManagerEncryptionKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetManagerEncryptionKeyArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetManagerEncryptionKeyResult struct {
	EncryptionAlgorithm        string  `pulumi:"encryptionAlgorithm"`
	Value                      string  `pulumi:"value"`
	ValueCertificateThumbprint *string `pulumi:"valueCertificateThumbprint"`
}

func GetManagerEncryptionKeyOutput(ctx *pulumi.Context, args GetManagerEncryptionKeyOutputArgs, opts ...pulumi.InvokeOption) GetManagerEncryptionKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetManagerEncryptionKeyResult, error) {
			args := v.(GetManagerEncryptionKeyArgs)
			r, err := GetManagerEncryptionKey(ctx, &args, opts...)
			var s GetManagerEncryptionKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetManagerEncryptionKeyResultOutput)
}

type GetManagerEncryptionKeyOutputArgs struct {
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetManagerEncryptionKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagerEncryptionKeyArgs)(nil)).Elem()
}


type GetManagerEncryptionKeyResultOutput struct{ *pulumi.OutputState }

func (GetManagerEncryptionKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagerEncryptionKeyResult)(nil)).Elem()
}

func (o GetManagerEncryptionKeyResultOutput) ToGetManagerEncryptionKeyResultOutput() GetManagerEncryptionKeyResultOutput {
	return o
}

func (o GetManagerEncryptionKeyResultOutput) ToGetManagerEncryptionKeyResultOutputWithContext(ctx context.Context) GetManagerEncryptionKeyResultOutput {
	return o
}

func (o GetManagerEncryptionKeyResultOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagerEncryptionKeyResult) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

func (o GetManagerEncryptionKeyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagerEncryptionKeyResult) string { return v.Value }).(pulumi.StringOutput)
}

func (o GetManagerEncryptionKeyResultOutput) ValueCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetManagerEncryptionKeyResult) *string { return v.ValueCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetManagerEncryptionKeyResultOutput{})
}
