


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiskEncryptionSet(ctx *pulumi.Context, args *LookupDiskEncryptionSetArgs, opts ...pulumi.InvokeOption) (*LookupDiskEncryptionSetResult, error) {
	var rv LookupDiskEncryptionSetResult
	err := ctx.Invoke("azure-native:compute/v20191101:getDiskEncryptionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskEncryptionSetArgs struct {
	DiskEncryptionSetName string `pulumi:"diskEncryptionSetName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDiskEncryptionSetResult struct {
	ActiveKey         *KeyVaultAndKeyReferenceResponse  `pulumi:"activeKey"`
	Id                string                            `pulumi:"id"`
	Identity          *EncryptionSetIdentityResponse    `pulumi:"identity"`
	Location          string                            `pulumi:"location"`
	Name              string                            `pulumi:"name"`
	PreviousKeys      []KeyVaultAndKeyReferenceResponse `pulumi:"previousKeys"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	Tags              map[string]string                 `pulumi:"tags"`
	Type              string                            `pulumi:"type"`
}
