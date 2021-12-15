


package v20200701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	var rv LookupDeploymentResult
	err := ctx.Invoke("azure-native:appplatform/v20200701:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDeploymentArgs struct {
	AppName           string `pulumi:"appName"`
	DeploymentName    string `pulumi:"deploymentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupDeploymentResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentResourcePropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                         `pulumi:"sku"`
	Type       string                               `pulumi:"type"`
}


func (val *LookupDeploymentResult) Defaults() *LookupDeploymentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
