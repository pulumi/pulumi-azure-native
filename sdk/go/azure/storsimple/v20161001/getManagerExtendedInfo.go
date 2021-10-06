


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagerExtendedInfo(ctx *pulumi.Context, args *LookupManagerExtendedInfoArgs, opts ...pulumi.InvokeOption) (*LookupManagerExtendedInfoResult, error) {
	var rv LookupManagerExtendedInfoResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getManagerExtendedInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerExtendedInfoArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagerExtendedInfoResult struct {
	Algorithm                   string  `pulumi:"algorithm"`
	EncryptionKey               *string `pulumi:"encryptionKey"`
	EncryptionKeyThumbprint     *string `pulumi:"encryptionKeyThumbprint"`
	Etag                        *string `pulumi:"etag"`
	Id                          string  `pulumi:"id"`
	IntegrityKey                string  `pulumi:"integrityKey"`
	Name                        string  `pulumi:"name"`
	PortalCertificateThumbprint *string `pulumi:"portalCertificateThumbprint"`
	Type                        string  `pulumi:"type"`
	Version                     *string `pulumi:"version"`
}
