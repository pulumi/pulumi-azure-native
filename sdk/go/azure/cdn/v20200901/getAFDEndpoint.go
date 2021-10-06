


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAFDEndpoint(ctx *pulumi.Context, args *LookupAFDEndpointArgs, opts ...pulumi.InvokeOption) (*LookupAFDEndpointResult, error) {
	var rv LookupAFDEndpointResult
	err := ctx.Invoke("azure-native:cdn/v20200901:getAFDEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAFDEndpointResult struct {
	DeploymentStatus             string             `pulumi:"deploymentStatus"`
	EnabledState                 *string            `pulumi:"enabledState"`
	HostName                     string             `pulumi:"hostName"`
	Id                           string             `pulumi:"id"`
	Location                     string             `pulumi:"location"`
	Name                         string             `pulumi:"name"`
	OriginResponseTimeoutSeconds *int               `pulumi:"originResponseTimeoutSeconds"`
	ProvisioningState            string             `pulumi:"provisioningState"`
	SystemData                   SystemDataResponse `pulumi:"systemData"`
	Tags                         map[string]string  `pulumi:"tags"`
	Type                         string             `pulumi:"type"`
}
