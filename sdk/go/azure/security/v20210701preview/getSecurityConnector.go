


package v20210701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityConnector(ctx *pulumi.Context, args *LookupSecurityConnectorArgs, opts ...pulumi.InvokeOption) (*LookupSecurityConnectorResult, error) {
	var rv LookupSecurityConnectorResult
	err := ctx.Invoke("azure-native:security/v20210701preview:getSecurityConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityConnectorArgs struct {
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}


type LookupSecurityConnectorResult struct {
	CloudName           *string                                                `pulumi:"cloudName"`
	Etag                *string                                                `pulumi:"etag"`
	HierarchyIdentifier *string                                                `pulumi:"hierarchyIdentifier"`
	Id                  string                                                 `pulumi:"id"`
	Kind                *string                                                `pulumi:"kind"`
	Location            *string                                                `pulumi:"location"`
	Name                string                                                 `pulumi:"name"`
	Offerings           []interface{}                                          `pulumi:"offerings"`
	OrganizationalData  *SecurityConnectorPropertiesResponseOrganizationalData `pulumi:"organizationalData"`
	SystemData          SystemDataResponse                                     `pulumi:"systemData"`
	Tags                map[string]string                                      `pulumi:"tags"`
	Type                string                                                 `pulumi:"type"`
}
