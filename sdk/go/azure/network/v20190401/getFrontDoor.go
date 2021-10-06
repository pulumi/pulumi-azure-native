


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFrontDoor(ctx *pulumi.Context, args *LookupFrontDoorArgs, opts ...pulumi.InvokeOption) (*LookupFrontDoorResult, error) {
	var rv LookupFrontDoorResult
	err := ctx.Invoke("azure-native:network/v20190401:getFrontDoor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFrontDoorArgs struct {
	FrontDoorName     string `pulumi:"frontDoorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFrontDoorResult struct {
	BackendPools          []BackendPoolResponse                `pulumi:"backendPools"`
	BackendPoolsSettings  *BackendPoolsSettingsResponse        `pulumi:"backendPoolsSettings"`
	Cname                 string                               `pulumi:"cname"`
	EnabledState          *string                              `pulumi:"enabledState"`
	FriendlyName          *string                              `pulumi:"friendlyName"`
	FrontendEndpoints     []FrontendEndpointResponse           `pulumi:"frontendEndpoints"`
	HealthProbeSettings   []HealthProbeSettingsModelResponse   `pulumi:"healthProbeSettings"`
	Id                    string                               `pulumi:"id"`
	LoadBalancingSettings []LoadBalancingSettingsModelResponse `pulumi:"loadBalancingSettings"`
	Location              *string                              `pulumi:"location"`
	Name                  string                               `pulumi:"name"`
	ProvisioningState     string                               `pulumi:"provisioningState"`
	ResourceState         string                               `pulumi:"resourceState"`
	RoutingRules          []RoutingRuleResponse                `pulumi:"routingRules"`
	Tags                  map[string]string                    `pulumi:"tags"`
	Type                  string                               `pulumi:"type"`
}
