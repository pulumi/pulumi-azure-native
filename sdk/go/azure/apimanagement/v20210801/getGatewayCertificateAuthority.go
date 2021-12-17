


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayCertificateAuthority(ctx *pulumi.Context, args *LookupGatewayCertificateAuthorityArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCertificateAuthorityResult, error) {
	var rv LookupGatewayCertificateAuthorityResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:getGatewayCertificateAuthority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayCertificateAuthorityArgs struct {
	CertificateId     string `pulumi:"certificateId"`
	GatewayId         string `pulumi:"gatewayId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGatewayCertificateAuthorityResult struct {
	Id        string `pulumi:"id"`
	IsTrusted *bool  `pulumi:"isTrusted"`
	Name      string `pulumi:"name"`
	Type      string `pulumi:"type"`
}
