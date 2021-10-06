


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:web/v20150801:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCertificateResult struct {
	CerBlob                   *string                            `pulumi:"cerBlob"`
	ExpirationDate            *string                            `pulumi:"expirationDate"`
	FriendlyName              *string                            `pulumi:"friendlyName"`
	HostNames                 []string                           `pulumi:"hostNames"`
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	Id                        *string                            `pulumi:"id"`
	IssueDate                 *string                            `pulumi:"issueDate"`
	Issuer                    *string                            `pulumi:"issuer"`
	Kind                      *string                            `pulumi:"kind"`
	Location                  string                             `pulumi:"location"`
	Name                      *string                            `pulumi:"name"`
	Password                  *string                            `pulumi:"password"`
	PfxBlob                   *string                            `pulumi:"pfxBlob"`
	PublicKeyHash             *string                            `pulumi:"publicKeyHash"`
	SelfLink                  *string                            `pulumi:"selfLink"`
	SiteName                  *string                            `pulumi:"siteName"`
	SubjectName               *string                            `pulumi:"subjectName"`
	Tags                      map[string]string                  `pulumi:"tags"`
	Thumbprint                *string                            `pulumi:"thumbprint"`
	Type                      *string                            `pulumi:"type"`
	Valid                     *bool                              `pulumi:"valid"`
}
