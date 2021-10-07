


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiskEncryptionSet(ctx *pulumi.Context, args *LookupDiskEncryptionSetArgs, opts ...pulumi.InvokeOption) (*LookupDiskEncryptionSetResult, error) {
	var rv LookupDiskEncryptionSetResult
	err := ctx.Invoke("azure-native:compute/v20210401:getDiskEncryptionSet", args, &rv, opts...)
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
	ActiveKey                         *KeyForDiskEncryptionSetResponse  `pulumi:"activeKey"`
	AutoKeyRotationError              ApiErrorResponse                  `pulumi:"autoKeyRotationError"`
	EncryptionType                    *string                           `pulumi:"encryptionType"`
	Id                                string                            `pulumi:"id"`
	Identity                          *EncryptionSetIdentityResponse    `pulumi:"identity"`
	LastKeyRotationTimestamp          string                            `pulumi:"lastKeyRotationTimestamp"`
	Location                          string                            `pulumi:"location"`
	Name                              string                            `pulumi:"name"`
	PreviousKeys                      []KeyForDiskEncryptionSetResponse `pulumi:"previousKeys"`
	ProvisioningState                 string                            `pulumi:"provisioningState"`
	RotationToLatestKeyVersionEnabled *bool                             `pulumi:"rotationToLatestKeyVersionEnabled"`
	Tags                              map[string]string                 `pulumi:"tags"`
	Type                              string                            `pulumi:"type"`
}
