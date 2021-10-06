


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupController(ctx *pulumi.Context, args *LookupControllerArgs, opts ...pulumi.InvokeOption) (*LookupControllerResult, error) {
	var rv LookupControllerResult
	err := ctx.Invoke("azure-native:devspaces/v20190401:getController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupControllerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupControllerResult struct {
	DataPlaneFqdn                    string            `pulumi:"dataPlaneFqdn"`
	HostSuffix                       string            `pulumi:"hostSuffix"`
	Id                               string            `pulumi:"id"`
	Location                         string            `pulumi:"location"`
	Name                             string            `pulumi:"name"`
	ProvisioningState                string            `pulumi:"provisioningState"`
	Sku                              SkuResponse       `pulumi:"sku"`
	Tags                             map[string]string `pulumi:"tags"`
	TargetContainerHostApiServerFqdn string            `pulumi:"targetContainerHostApiServerFqdn"`
	TargetContainerHostResourceId    string            `pulumi:"targetContainerHostResourceId"`
	Type                             string            `pulumi:"type"`
}
