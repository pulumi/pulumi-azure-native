


package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartner(ctx *pulumi.Context, args *LookupPartnerArgs, opts ...pulumi.InvokeOption) (*LookupPartnerResult, error) {
	var rv LookupPartnerResult
	err := ctx.Invoke("azure-native:managementpartner/v20180201:getPartner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerArgs struct {
	PartnerId string `pulumi:"partnerId"`
}


type LookupPartnerResult struct {
	CreatedTime *string `pulumi:"createdTime"`
	Etag        *int    `pulumi:"etag"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	ObjectId    *string `pulumi:"objectId"`
	PartnerId   *string `pulumi:"partnerId"`
	PartnerName *string `pulumi:"partnerName"`
	TenantId    *string `pulumi:"tenantId"`
	Type        string  `pulumi:"type"`
	UpdatedTime *string `pulumi:"updatedTime"`
	Version     *int    `pulumi:"version"`
}
