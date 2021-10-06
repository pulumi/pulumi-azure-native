


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:batch/v20210601:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	AccountName       string `pulumi:"accountName"`
	CertificateName   string `pulumi:"certificateName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCertificateResult struct {
	DeleteCertificateError                  DeleteCertificateErrorResponse `pulumi:"deleteCertificateError"`
	Etag                                    string                         `pulumi:"etag"`
	Format                                  *string                        `pulumi:"format"`
	Id                                      string                         `pulumi:"id"`
	Name                                    string                         `pulumi:"name"`
	PreviousProvisioningState               string                         `pulumi:"previousProvisioningState"`
	PreviousProvisioningStateTransitionTime string                         `pulumi:"previousProvisioningStateTransitionTime"`
	ProvisioningState                       string                         `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime         string                         `pulumi:"provisioningStateTransitionTime"`
	PublicData                              string                         `pulumi:"publicData"`
	Thumbprint                              *string                        `pulumi:"thumbprint"`
	ThumbprintAlgorithm                     *string                        `pulumi:"thumbprintAlgorithm"`
	Type                                    string                         `pulumi:"type"`
}
