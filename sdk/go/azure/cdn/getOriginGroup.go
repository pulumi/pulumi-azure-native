


package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOriginGroup(ctx *pulumi.Context, args *LookupOriginGroupArgs, opts ...pulumi.InvokeOption) (*LookupOriginGroupResult, error) {
	var rv LookupOriginGroupResult
	err := ctx.Invoke("azure-native:cdn:getOriginGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOriginGroupArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	OriginGroupName   string `pulumi:"originGroupName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOriginGroupResult struct {
	HealthProbeSettings                                   *HealthProbeParametersResponse                       `pulumi:"healthProbeSettings"`
	Id                                                    string                                               `pulumi:"id"`
	Name                                                  string                                               `pulumi:"name"`
	Origins                                               []ResourceReferenceResponse                          `pulumi:"origins"`
	ProvisioningState                                     string                                               `pulumi:"provisioningState"`
	ResourceState                                         string                                               `pulumi:"resourceState"`
	ResponseBasedOriginErrorDetectionSettings             *ResponseBasedOriginErrorDetectionParametersResponse `pulumi:"responseBasedOriginErrorDetectionSettings"`
	SystemData                                            SystemDataResponse                                   `pulumi:"systemData"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                                 `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	Type                                                  string                                               `pulumi:"type"`
}
