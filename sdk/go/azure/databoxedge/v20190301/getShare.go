


package v20190301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupShare(ctx *pulumi.Context, args *LookupShareArgs, opts ...pulumi.InvokeOption) (*LookupShareResult, error) {
	var rv LookupShareResult
	err := ctx.Invoke("azure-native:databoxedge/v20190301:getShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupShareArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupShareResult struct {
	AccessProtocol     string                      `pulumi:"accessProtocol"`
	AzureContainerInfo *AzureContainerInfoResponse `pulumi:"azureContainerInfo"`
	ClientAccessRights []ClientAccessRightResponse `pulumi:"clientAccessRights"`
	DataPolicy         *string                     `pulumi:"dataPolicy"`
	Description        *string                     `pulumi:"description"`
	Id                 string                      `pulumi:"id"`
	MonitoringStatus   string                      `pulumi:"monitoringStatus"`
	Name               string                      `pulumi:"name"`
	RefreshDetails     *RefreshDetailsResponse     `pulumi:"refreshDetails"`
	ShareMappings      []MountPointMapResponse     `pulumi:"shareMappings"`
	ShareStatus        string                      `pulumi:"shareStatus"`
	Type               string                      `pulumi:"type"`
	UserAccessRights   []UserAccessRightResponse   `pulumi:"userAccessRights"`
}
