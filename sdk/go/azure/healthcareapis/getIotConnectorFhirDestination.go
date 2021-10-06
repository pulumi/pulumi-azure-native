


package healthcareapis

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotConnectorFhirDestination(ctx *pulumi.Context, args *LookupIotConnectorFhirDestinationArgs, opts ...pulumi.InvokeOption) (*LookupIotConnectorFhirDestinationResult, error) {
	var rv LookupIotConnectorFhirDestinationResult
	err := ctx.Invoke("azure-native:healthcareapis:getIotConnectorFhirDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotConnectorFhirDestinationArgs struct {
	FhirDestinationName string `pulumi:"fhirDestinationName"`
	IotConnectorName    string `pulumi:"iotConnectorName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	WorkspaceName       string `pulumi:"workspaceName"`
}


type LookupIotConnectorFhirDestinationResult struct {
	Etag                           *string                      `pulumi:"etag"`
	FhirMapping                    IotMappingPropertiesResponse `pulumi:"fhirMapping"`
	FhirServiceResourceId          string                       `pulumi:"fhirServiceResourceId"`
	Id                             string                       `pulumi:"id"`
	Location                       *string                      `pulumi:"location"`
	Name                           string                       `pulumi:"name"`
	ResourceIdentityResolutionType string                       `pulumi:"resourceIdentityResolutionType"`
	SystemData                     SystemDataResponse           `pulumi:"systemData"`
	Type                           string                       `pulumi:"type"`
}
