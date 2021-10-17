


package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:automation/v20190601:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	CertificateName       string `pulumi:"certificateName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupCertificateResult struct {
	CreationTime     string  `pulumi:"creationTime"`
	Description      *string `pulumi:"description"`
	ExpiryTime       string  `pulumi:"expiryTime"`
	Id               string  `pulumi:"id"`
	IsExportable     bool    `pulumi:"isExportable"`
	LastModifiedTime string  `pulumi:"lastModifiedTime"`
	Name             string  `pulumi:"name"`
	Thumbprint       string  `pulumi:"thumbprint"`
	Type             string  `pulumi:"type"`
}
