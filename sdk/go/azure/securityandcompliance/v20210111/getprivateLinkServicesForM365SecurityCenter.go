


package v20210111

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context, args *GetprivateLinkServicesForM365SecurityCenterArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForM365SecurityCenterResult, error) {
	var rv GetprivateLinkServicesForM365SecurityCenterResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210111:getprivateLinkServicesForM365SecurityCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForM365SecurityCenterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForM365SecurityCenterResult struct {
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
