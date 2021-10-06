


package v20161001

import (
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
