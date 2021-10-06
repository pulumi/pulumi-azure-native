


package v20201101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApplicationGatewayBackendHealthOnDemand(ctx *pulumi.Context, args *GetApplicationGatewayBackendHealthOnDemandArgs, opts ...pulumi.InvokeOption) (*GetApplicationGatewayBackendHealthOnDemandResult, error) {
	var rv GetApplicationGatewayBackendHealthOnDemandResult
	err := ctx.Invoke("azure-native:network/v20201101:getApplicationGatewayBackendHealthOnDemand", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApplicationGatewayBackendHealthOnDemandArgs struct {
	ApplicationGatewayName              string                                      `pulumi:"applicationGatewayName"`
	BackendAddressPool                  *SubResource                                `pulumi:"backendAddressPool"`
	BackendHttpSettings                 *SubResource                                `pulumi:"backendHttpSettings"`
	Expand                              *string                                     `pulumi:"expand"`
	Host                                *string                                     `pulumi:"host"`
	Match                               *ApplicationGatewayProbeHealthResponseMatch `pulumi:"match"`
	Path                                *string                                     `pulumi:"path"`
	PickHostNameFromBackendHttpSettings *bool                                       `pulumi:"pickHostNameFromBackendHttpSettings"`
	Protocol                            *string                                     `pulumi:"protocol"`
	ResourceGroupName                   string                                      `pulumi:"resourceGroupName"`
	Timeout                             *int                                        `pulumi:"timeout"`
}


type GetApplicationGatewayBackendHealthOnDemandResult struct {
	BackendAddressPool        *ApplicationGatewayBackendAddressPoolResponse        `pulumi:"backendAddressPool"`
	BackendHealthHttpSettings *ApplicationGatewayBackendHealthHttpSettingsResponse `pulumi:"backendHealthHttpSettings"`
}
