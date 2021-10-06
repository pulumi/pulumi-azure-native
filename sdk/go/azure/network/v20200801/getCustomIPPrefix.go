


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomIPPrefix(ctx *pulumi.Context, args *LookupCustomIPPrefixArgs, opts ...pulumi.InvokeOption) (*LookupCustomIPPrefixResult, error) {
	var rv LookupCustomIPPrefixResult
	err := ctx.Invoke("azure-native:network/v20200801:getCustomIPPrefix", args, &rv, opts...)
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
	Cidr              *string                   `pulumi:"cidr"`
	CommissionedState *string                   `pulumi:"commissionedState"`
	Etag              string                    `pulumi:"etag"`
	ExtendedLocation  *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                *string                   `pulumi:"id"`
	Location          *string                   `pulumi:"location"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	PublicIpPrefixes  []SubResourceResponse     `pulumi:"publicIpPrefixes"`
	ResourceGuid      string                    `pulumi:"resourceGuid"`
	Tags              map[string]string         `pulumi:"tags"`
	Type              string                    `pulumi:"type"`
	Zones             []string                  `pulumi:"zones"`
}
