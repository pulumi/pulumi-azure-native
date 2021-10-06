


package attestation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttestationProvider(ctx *pulumi.Context, args *LookupAttestationProviderArgs, opts ...pulumi.InvokeOption) (*LookupAttestationProviderResult, error) {
	var rv LookupAttestationProviderResult
	err := ctx.Invoke("azure-native:attestation:getAttestationProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationProviderArgs struct {
	ProviderName      string `pulumi:"providerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAttestationProviderResult struct {
	AttestUri                  *string                             `pulumi:"attestUri"`
	Id                         string                              `pulumi:"id"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	Status                     *string                             `pulumi:"status"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	TrustModel                 *string                             `pulumi:"trustModel"`
	Type                       string                              `pulumi:"type"`
}
