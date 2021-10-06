


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:security/v20200101preview:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	ConnectorName string `pulumi:"connectorName"`
}


type LookupConnectorResult struct {
	AuthenticationDetails interface{}                              `pulumi:"authenticationDetails"`
	HybridComputeSettings *HybridComputeSettingsPropertiesResponse `pulumi:"hybridComputeSettings"`
	Id                    string                                   `pulumi:"id"`
	Name                  string                                   `pulumi:"name"`
	Type                  string                                   `pulumi:"type"`
}
