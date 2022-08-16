


package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagerPublicEncryptionKey(ctx *pulumi.Context, args *ListManagerPublicEncryptionKeyArgs, opts ...pulumi.InvokeOption) (*ListManagerPublicEncryptionKeyResult, error) {
	var rv ListManagerPublicEncryptionKeyResult
	err := ctx.Invoke("azure-native:storsimple:listManagerPublicEncryptionKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagerPublicEncryptionKeyArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListManagerPublicEncryptionKeyResult struct {
	EncryptionAlgorithm        string  `pulumi:"encryptionAlgorithm"`
	Value                      string  `pulumi:"value"`
	ValueCertificateThumbprint *string `pulumi:"valueCertificateThumbprint"`
}

func ListManagerPublicEncryptionKeyOutput(ctx *pulumi.Context, args ListManagerPublicEncryptionKeyOutputArgs, opts ...pulumi.InvokeOption) ListManagerPublicEncryptionKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagerPublicEncryptionKeyResult, error) {
			args := v.(ListManagerPublicEncryptionKeyArgs)
			r, err := ListManagerPublicEncryptionKey(ctx, &args, opts...)
			var s ListManagerPublicEncryptionKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagerPublicEncryptionKeyResultOutput)
}

type ListManagerPublicEncryptionKeyOutputArgs struct {
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListManagerPublicEncryptionKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerPublicEncryptionKeyArgs)(nil)).Elem()
}


type ListManagerPublicEncryptionKeyResultOutput struct{ *pulumi.OutputState }

func (ListManagerPublicEncryptionKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerPublicEncryptionKeyResult)(nil)).Elem()
}

func (o ListManagerPublicEncryptionKeyResultOutput) ToListManagerPublicEncryptionKeyResultOutput() ListManagerPublicEncryptionKeyResultOutput {
	return o
}

func (o ListManagerPublicEncryptionKeyResultOutput) ToListManagerPublicEncryptionKeyResultOutputWithContext(ctx context.Context) ListManagerPublicEncryptionKeyResultOutput {
	return o
}

func (o ListManagerPublicEncryptionKeyResultOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagerPublicEncryptionKeyResult) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

func (o ListManagerPublicEncryptionKeyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagerPublicEncryptionKeyResult) string { return v.Value }).(pulumi.StringOutput)
}

func (o ListManagerPublicEncryptionKeyResultOutput) ValueCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListManagerPublicEncryptionKeyResult) *string { return v.ValueCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagerPublicEncryptionKeyResultOutput{})
}
