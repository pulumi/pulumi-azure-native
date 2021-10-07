


package v20210325preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkServicesForEDMUpload(ctx *pulumi.Context, args *GetprivateLinkServicesForEDMUploadArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForEDMUploadResult, error) {
	var rv GetprivateLinkServicesForEDMUploadResult
	err := ctx.Invoke("azure-native:m365securityandcompliance/v20210325preview:getprivateLinkServicesForEDMUpload", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForEDMUploadArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetprivateLinkServicesForEDMUploadResult struct {
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
