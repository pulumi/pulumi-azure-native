


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceCertificateOrderCertificate(ctx *pulumi.Context, args *LookupAppServiceCertificateOrderCertificateArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceCertificateOrderCertificateResult, error) {
	var rv LookupAppServiceCertificateOrderCertificateResult
	err := ctx.Invoke("azure-native:certificateregistration/v20150801:getAppServiceCertificateOrderCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceCertificateOrderCertificateArgs struct {
	CertificateOrderName string `pulumi:"certificateOrderName"`
	Name                 string `pulumi:"name"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupAppServiceCertificateOrderCertificateResult struct {
	Id                 string            `pulumi:"id"`
	KeyVaultId         *string           `pulumi:"keyVaultId"`
	KeyVaultSecretName *string           `pulumi:"keyVaultSecretName"`
	Kind               *string           `pulumi:"kind"`
	Location           string            `pulumi:"location"`
	Name               string            `pulumi:"name"`
	ProvisioningState  string            `pulumi:"provisioningState"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
}
