


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackend(ctx *pulumi.Context, args *LookupBackendArgs, opts ...pulumi.InvokeOption) (*LookupBackendResult, error) {
	var rv LookupBackendResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:getBackend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBackendArgs struct {
	BackendId         string `pulumi:"backendId"`
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


func (val *LookupBackendResult) Defaults() *LookupBackendResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Tls = tmp.Tls.Defaults()

	return &tmp
}
