


package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAFDOriginGroup(ctx *pulumi.Context, args *LookupAFDOriginGroupArgs, opts ...pulumi.InvokeOption) (*LookupAFDOriginGroupResult, error) {
	var rv LookupAFDOriginGroupResult
	err := ctx.Invoke("azure-native:cdn:getAFDOriginGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDOriginGroupArgs struct {
	OriginGroupName   string `pulumi:"originGroupName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAFDOriginGroupResult struct {
	DeploymentStatus                                      string                                               `pulumi:"deploymentStatus"`
	HealthProbeSettings                                   *HealthProbeParametersResponse                       `pulumi:"healthProbeSettings"`
	Id                                                    string                                               `pulumi:"id"`
	LoadBalancingSettings                                 *LoadBalancingSettingsParametersResponse             `pulumi:"loadBalancingSettings"`
	Name                                                  string                                               `pulumi:"name"`
	ProvisioningState                                     string                                               `pulumi:"provisioningState"`
	ResponseBasedAfdOriginErrorDetectionSettings          *ResponseBasedOriginErrorDetectionParametersResponse `pulumi:"responseBasedAfdOriginErrorDetectionSettings"`
	SessionAffinityState                                  *string                                              `pulumi:"sessionAffinityState"`
	SystemData                                            SystemDataResponse                                   `pulumi:"systemData"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                                 `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	Type                                                  string                                               `pulumi:"type"`
}
