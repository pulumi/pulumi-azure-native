


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDpsCertificate(ctx *pulumi.Context, args *LookupDpsCertificateArgs, opts ...pulumi.InvokeOption) (*LookupDpsCertificateResult, error) {
	var rv LookupDpsCertificateResult
	err := ctx.Invoke("azure-native:devices/v20200301:getDpsCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDpsCertificateArgs struct {
	CertificateName         string `pulumi:"certificateName"`
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupDpsCertificateResult struct {
	Etag       string                        `pulumi:"etag"`
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties CertificatePropertiesResponse `pulumi:"properties"`
	Type       string                        `pulumi:"type"`
}
