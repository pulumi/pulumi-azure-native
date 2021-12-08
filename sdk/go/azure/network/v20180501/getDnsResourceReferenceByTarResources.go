


package v20180501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDnsResourceReferenceByTarResources(ctx *pulumi.Context, args *GetDnsResourceReferenceByTarResourcesArgs, opts ...pulumi.InvokeOption) (*GetDnsResourceReferenceByTarResourcesResult, error) {
	var rv GetDnsResourceReferenceByTarResourcesResult
	err := ctx.Invoke("azure-native:network/v20180501:getDnsResourceReferenceByTarResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDnsResourceReferenceByTarResourcesArgs struct {
	TargetResources []SubResource `pulumi:"targetResources"`
}


type GetDnsResourceReferenceByTarResourcesResult struct {
	DnsResourceReferences []DnsResourceReferenceResponse `pulumi:"dnsResourceReferences"`
}
