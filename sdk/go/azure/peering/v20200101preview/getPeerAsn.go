


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPeerAsn(ctx *pulumi.Context, args *LookupPeerAsnArgs, opts ...pulumi.InvokeOption) (*LookupPeerAsnResult, error) {
	var rv LookupPeerAsnResult
	err := ctx.Invoke("azure-native:peering/v20200101preview:getPeerAsn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeerAsnArgs struct {
	PeerAsnName string `pulumi:"peerAsnName"`
}


type LookupPeerAsnResult struct {
	ErrorMessage      string                  `pulumi:"errorMessage"`
	Id                string                  `pulumi:"id"`
	Name              string                  `pulumi:"name"`
	PeerAsn           *int                    `pulumi:"peerAsn"`
	PeerContactDetail []ContactDetailResponse `pulumi:"peerContactDetail"`
	PeerName          *string                 `pulumi:"peerName"`
	Type              string                  `pulumi:"type"`
	ValidationState   *string                 `pulumi:"validationState"`
}
