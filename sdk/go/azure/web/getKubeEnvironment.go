


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKubeEnvironment(ctx *pulumi.Context, args *LookupKubeEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupKubeEnvironmentResult, error) {
	var rv LookupKubeEnvironmentResult
	err := ctx.Invoke("azure-native:web:getKubeEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKubeEnvironmentArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupKubeEnvironmentResult struct {
	AksResourceID               *string                       `pulumi:"aksResourceID"`
	AppLogsConfiguration        *AppLogsConfigurationResponse `pulumi:"appLogsConfiguration"`
	ArcConfiguration            *ArcConfigurationResponse     `pulumi:"arcConfiguration"`
	DefaultDomain               string                        `pulumi:"defaultDomain"`
	DeploymentErrors            string                        `pulumi:"deploymentErrors"`
	ExtendedLocation            *ExtendedLocationResponse     `pulumi:"extendedLocation"`
	Id                          string                        `pulumi:"id"`
	InternalLoadBalancerEnabled *bool                         `pulumi:"internalLoadBalancerEnabled"`
	Kind                        *string                       `pulumi:"kind"`
	Location                    string                        `pulumi:"location"`
	Name                        string                        `pulumi:"name"`
	ProvisioningState           string                        `pulumi:"provisioningState"`
	StaticIp                    *string                       `pulumi:"staticIp"`
	Tags                        map[string]string             `pulumi:"tags"`
	Type                        string                        `pulumi:"type"`
}
