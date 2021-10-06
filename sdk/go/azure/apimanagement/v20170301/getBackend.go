


package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackend(ctx *pulumi.Context, args *LookupBackendArgs, opts ...pulumi.InvokeOption) (*LookupBackendResult, error) {
	var rv LookupBackendResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getBackend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackendArgs struct {
	Backendid         string `pulumi:"backendid"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupBackendResult struct {
	Credentials *BackendCredentialsContractResponse `pulumi:"credentials"`
	Description *string                             `pulumi:"description"`
	Id          string                              `pulumi:"id"`
	Name        string                              `pulumi:"name"`
	Properties  BackendPropertiesResponse           `pulumi:"properties"`
	Protocol    string                              `pulumi:"protocol"`
	Proxy       *BackendProxyContractResponse       `pulumi:"proxy"`
	ResourceId  *string                             `pulumi:"resourceId"`
	Title       *string                             `pulumi:"title"`
	Tls         *BackendTlsPropertiesResponse       `pulumi:"tls"`
	Type        string                              `pulumi:"type"`
	Url         string                              `pulumi:"url"`
}
