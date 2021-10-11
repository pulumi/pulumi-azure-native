


package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPublicCertificateSlot(ctx *pulumi.Context, args *LookupWebAppPublicCertificateSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPublicCertificateSlotResult, error) {
	var rv LookupWebAppPublicCertificateSlotResult
	err := ctx.Invoke("azure-native:web/v20160801:getWebAppPublicCertificateSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPublicCertificateSlotArgs struct {
	Name                  string `pulumi:"name"`
	PublicCertificateName string `pulumi:"publicCertificateName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	Slot                  string `pulumi:"slot"`
}


type LookupWebAppPublicCertificateSlotResult struct {
	Blob                      *string `pulumi:"blob"`
	Id                        string  `pulumi:"id"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	PublicCertificateLocation *string `pulumi:"publicCertificateLocation"`
	Thumbprint                string  `pulumi:"thumbprint"`
	Type                      string  `pulumi:"type"`
}
