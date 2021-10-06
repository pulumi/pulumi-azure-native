


package securityandcompliance

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForO365ManagementActivityAPI(ctx *pulumi.Context, args *GetprivateLinkServicesForO365ManagementActivityAPIArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForO365ManagementActivityAPIResult, error) {
	var rv GetprivateLinkServicesForO365ManagementActivityAPIResult
	err := ctx.Invoke("azure-native:securityandcompliance:getprivateLinkServicesForO365ManagementActivityAPI", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForO365ManagementActivityAPIArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForO365ManagementActivityAPIResult struct {
	Etag       *string                           `pulumi:"etag"`
	Id         string                            `pulumi:"id"`
	Identity   *ServicesResourceResponseIdentity `pulumi:"identity"`
	Kind       string                            `pulumi:"kind"`
	Location   string                            `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties ServicesPropertiesResponse        `pulumi:"properties"`
	SystemData SystemDataResponse                `pulumi:"systemData"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}
