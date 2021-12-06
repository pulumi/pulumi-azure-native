


package v20150801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	var rv LookupConnectionResult
	err := ctx.Invoke("azure-native:web/v20150801preview:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectionResult struct {
	Api                      *ExpandedParentApiEntityResponse                     `pulumi:"api"`
	ChangedTime              *string                                              `pulumi:"changedTime"`
	CreatedTime              *string                                              `pulumi:"createdTime"`
	CustomParameterValues    map[string]ParameterCustomLoginSettingValuesResponse `pulumi:"customParameterValues"`
	DisplayName              *string                                              `pulumi:"displayName"`
	FirstExpirationTime      *string                                              `pulumi:"firstExpirationTime"`
	Id                       *string                                              `pulumi:"id"`
	Keywords                 []string                                             `pulumi:"keywords"`
	Kind                     *string                                              `pulumi:"kind"`
	Location                 string                                               `pulumi:"location"`
	Metadata                 interface{}                                          `pulumi:"metadata"`
	Name                     *string                                              `pulumi:"name"`
	NonSecretParameterValues map[string]interface{}                               `pulumi:"nonSecretParameterValues"`
	ParameterValues          map[string]interface{}                               `pulumi:"parameterValues"`
	Statuses                 []ConnectionStatusResponse                           `pulumi:"statuses"`
	Tags                     map[string]string                                    `pulumi:"tags"`
	TenantId                 *string                                              `pulumi:"tenantId"`
	Type                     *string                                              `pulumi:"type"`
}
