


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerTrustCertificate(ctx *pulumi.Context, args *LookupServerTrustCertificateArgs, opts ...pulumi.InvokeOption) (*LookupServerTrustCertificateResult, error) {
	var rv LookupServerTrustCertificateResult
	err := ctx.Invoke("azure-native:sql:getServerTrustCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerTrustCertificateArgs struct {
	CertificateName     string `pulumi:"certificateName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupServerTrustCertificateResult struct {
	CertificateName string  `pulumi:"certificateName"`
	Id              string  `pulumi:"id"`
	Name            string  `pulumi:"name"`
	PublicBlob      *string `pulumi:"publicBlob"`
	Thumbprint      string  `pulumi:"thumbprint"`
	Type            string  `pulumi:"type"`
}
