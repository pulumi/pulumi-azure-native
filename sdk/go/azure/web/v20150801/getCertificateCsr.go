


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificateCsr(ctx *pulumi.Context, args *LookupCertificateCsrArgs, opts ...pulumi.InvokeOption) (*LookupCertificateCsrResult, error) {
	var rv LookupCertificateCsrResult
	err := ctx.Invoke("azure-native:web/v20150801:getCertificateCsr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateCsrArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCertificateCsrResult struct {
	CsrString          *string           `pulumi:"csrString"`
	DistinguishedName  *string           `pulumi:"distinguishedName"`
	HostingEnvironment *string           `pulumi:"hostingEnvironment"`
	Id                 *string           `pulumi:"id"`
	Kind               *string           `pulumi:"kind"`
	Location           string            `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	Password           *string           `pulumi:"password"`
	PfxBlob            *string           `pulumi:"pfxBlob"`
	PublicKeyHash      *string           `pulumi:"publicKeyHash"`
	Tags               map[string]string `pulumi:"tags"`
	Type               *string           `pulumi:"type"`
}
