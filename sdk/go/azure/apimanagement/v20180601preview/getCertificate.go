


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	CertificateId     string `pulumi:"certificateId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupCertificateResult struct {
	ExpirationDate string `pulumi:"expirationDate"`
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	Subject        string `pulumi:"subject"`
	Thumbprint     string `pulumi:"thumbprint"`
	Type           string `pulumi:"type"`
}
