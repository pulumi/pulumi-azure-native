


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerApp(ctx *pulumi.Context, args *LookupContainerAppArgs, opts ...pulumi.InvokeOption) (*LookupContainerAppResult, error) {
	var rv LookupContainerAppResult
	err := ctx.Invoke("azure-native:web/v20210301:getContainerApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerAppArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupContainerAppResult struct {
	Configuration      *ConfigurationResponse `pulumi:"configuration"`
	Id                 string                 `pulumi:"id"`
	Kind               *string                `pulumi:"kind"`
	KubeEnvironmentId  *string                `pulumi:"kubeEnvironmentId"`
	LatestRevisionFqdn string                 `pulumi:"latestRevisionFqdn"`
	LatestRevisionName string                 `pulumi:"latestRevisionName"`
	Location           string                 `pulumi:"location"`
	Name               string                 `pulumi:"name"`
	ProvisioningState  string                 `pulumi:"provisioningState"`
	Tags               map[string]string      `pulumi:"tags"`
	Template           *TemplateResponse      `pulumi:"template"`
	Type               string                 `pulumi:"type"`
}
