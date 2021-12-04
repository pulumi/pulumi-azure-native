


package securityandcompliance

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForSCCPowershell(ctx *pulumi.Context, args *GetprivateLinkServicesForSCCPowershellArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForSCCPowershellResult, error) {
	var rv GetprivateLinkServicesForSCCPowershellResult
	err := ctx.Invoke("azure-native:securityandcompliance:getprivateLinkServicesForSCCPowershell", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForSCCPowershellArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForSCCPowershellResult struct {
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
