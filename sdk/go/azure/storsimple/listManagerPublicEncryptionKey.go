


package storsimple

import (
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
