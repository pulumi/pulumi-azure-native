


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomIPPrefix(ctx *pulumi.Context, args *LookupCustomIPPrefixArgs, opts ...pulumi.InvokeOption) (*LookupCustomIPPrefixResult, error) {
	var rv LookupCustomIPPrefixResult
	err := ctx.Invoke("azure-native:network/v20210301:getCustomIPPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomIPPrefixArgs struct {
	CustomIpPrefixName string  `pulumi:"customIpPrefixName"`
	Expand             *string `pulumi:"expand"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupCustomIPPrefixResult struct {
	AuthorizationMessage  *string                   `pulumi:"authorizationMessage"`
	ChildCustomIpPrefixes []CustomIpPrefixResponse  `pulumi:"childCustomIpPrefixes"`
	Cidr                  *string                   `pulumi:"cidr"`
	CommissionedState     *string                   `pulumi:"commissionedState"`
	CustomIpPrefixParent  *CustomIpPrefixResponse   `pulumi:"customIpPrefixParent"`
	Etag                  string                    `pulumi:"etag"`
	ExtendedLocation      *ExtendedLocationResponse `pulumi:"extendedLocation"`
	FailedReason          string                    `pulumi:"failedReason"`
	Id                    *string                   `pulumi:"id"`
	Location              *string                   `pulumi:"location"`
	Name                  string                    `pulumi:"name"`
	ProvisioningState     string                    `pulumi:"provisioningState"`
	PublicIpPrefixes      []SubResourceResponse     `pulumi:"publicIpPrefixes"`
	ResourceGuid          string                    `pulumi:"resourceGuid"`
	SignedMessage         *string                   `pulumi:"signedMessage"`
	Tags                  map[string]string         `pulumi:"tags"`
	Type                  string                    `pulumi:"type"`
	Zones                 []string                  `pulumi:"zones"`
}
