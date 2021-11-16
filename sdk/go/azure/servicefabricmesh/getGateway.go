


package servicefabricmesh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	var rv LookupGatewayResult
	err := ctx.Invoke("azure-native:servicefabricmesh:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayArgs struct {
	GatewayResourceName string `pulumi:"gatewayResourceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupGatewayResult struct {
	Description        *string              `pulumi:"description"`
	DestinationNetwork NetworkRefResponse   `pulumi:"destinationNetwork"`
	Http               []HttpConfigResponse `pulumi:"http"`
	Id                 string               `pulumi:"id"`
	IpAddress          string               `pulumi:"ipAddress"`
	Location           string               `pulumi:"location"`
	Name               string               `pulumi:"name"`
	ProvisioningState  string               `pulumi:"provisioningState"`
	SourceNetwork      NetworkRefResponse   `pulumi:"sourceNetwork"`
	Status             string               `pulumi:"status"`
	StatusDetails      string               `pulumi:"statusDetails"`
	Tags               map[string]string    `pulumi:"tags"`
	Tcp                []TcpConfigResponse  `pulumi:"tcp"`
	Type               string               `pulumi:"type"`
}
