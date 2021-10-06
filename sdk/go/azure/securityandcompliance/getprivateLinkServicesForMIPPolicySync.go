


package securityandcompliance

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForMIPPolicySync(ctx *pulumi.Context, args *GetprivateLinkServicesForMIPPolicySyncArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForMIPPolicySyncResult, error) {
	var rv GetprivateLinkServicesForMIPPolicySyncResult
	err := ctx.Invoke("azure-native:securityandcompliance:getprivateLinkServicesForMIPPolicySync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForMIPPolicySyncArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForMIPPolicySyncResult struct {
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
