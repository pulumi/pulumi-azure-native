


package v20211101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVideoAnalyzer(ctx *pulumi.Context, args *LookupVideoAnalyzerArgs, opts ...pulumi.InvokeOption) (*LookupVideoAnalyzerResult, error) {
	var rv LookupVideoAnalyzerResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20211101preview:getVideoAnalyzer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVideoAnalyzerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVideoAnalyzerResult struct {
	Encryption                 *AccountEncryptionResponse          `pulumi:"encryption"`
	Endpoints                  []EndpointResponse                  `pulumi:"endpoints"`
	Id                         string                              `pulumi:"id"`
	Identity                   *VideoAnalyzerIdentityResponse      `pulumi:"identity"`
	IotHubs                    []IotHubResponse                    `pulumi:"iotHubs"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	NetworkAccessControl       *NetworkAccessControlResponse       `pulumi:"networkAccessControl"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	StorageAccounts            []StorageAccountResponse            `pulumi:"storageAccounts"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}
