


package importexport

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBitLockerKey(ctx *pulumi.Context, args *ListBitLockerKeyArgs, opts ...pulumi.InvokeOption) (*ListBitLockerKeyResult, error) {
	var rv ListBitLockerKeyResult
	err := ctx.Invoke("azure-native:importexport:listBitLockerKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBitLockerKeyArgs struct {
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListBitLockerKeyResult struct {
	Value []DriveBitLockerKeyResponse `pulumi:"value"`
}
