


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	var rv LookupAppResult
	err := ctx.Invoke("azure-native:iotcentral/v20210601:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupAppResult struct {
	ApplicationId string                                 `pulumi:"applicationId"`
	DisplayName   *string                                `pulumi:"displayName"`
	Id            string                                 `pulumi:"id"`
	Identity      *SystemAssignedServiceIdentityResponse `pulumi:"identity"`
	Location      string                                 `pulumi:"location"`
	Name          string                                 `pulumi:"name"`
	Sku           AppSkuInfoResponse                     `pulumi:"sku"`
	State         string                                 `pulumi:"state"`
	Subdomain     *string                                `pulumi:"subdomain"`
	Tags          map[string]string                      `pulumi:"tags"`
	Template      *string                                `pulumi:"template"`
	Type          string                                 `pulumi:"type"`
}
