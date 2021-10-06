


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPublicCertificate(ctx *pulumi.Context, args *LookupWebAppPublicCertificateArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPublicCertificateResult, error) {
	var rv LookupWebAppPublicCertificateResult
	err := ctx.Invoke("azure-native:web/v20201001:getWebAppPublicCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPublicCertificateArgs struct {
	Name                  string `pulumi:"name"`
	PublicCertificateName string `pulumi:"publicCertificateName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupWebAppPublicCertificateResult struct {
	Blob                      *string            `pulumi:"blob"`
	Id                        string             `pulumi:"id"`
	Kind                      *string            `pulumi:"kind"`
	Name                      string             `pulumi:"name"`
	PublicCertificateLocation *string            `pulumi:"publicCertificateLocation"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	Thumbprint                string             `pulumi:"thumbprint"`
	Type                      string             `pulumi:"type"`
}
