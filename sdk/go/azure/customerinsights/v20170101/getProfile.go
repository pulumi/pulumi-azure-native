


package v20170101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	HubName           string  `pulumi:"hubName"`
	LocaleCode        *string `pulumi:"localeCode"`
	ProfileName       string  `pulumi:"profileName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupProfileResult struct {
	ApiEntitySetName    *string                      `pulumi:"apiEntitySetName"`
	Attributes          map[string][]string          `pulumi:"attributes"`
	Description         map[string]string            `pulumi:"description"`
	DisplayName         map[string]string            `pulumi:"displayName"`
	EntityType          *string                      `pulumi:"entityType"`
	Fields              []PropertyDefinitionResponse `pulumi:"fields"`
	Id                  string                       `pulumi:"id"`
	InstancesCount      *int                         `pulumi:"instancesCount"`
	LargeImage          *string                      `pulumi:"largeImage"`
	LastChangedUtc      string                       `pulumi:"lastChangedUtc"`
	LocalizedAttributes map[string]map[string]string `pulumi:"localizedAttributes"`
	MediumImage         *string                      `pulumi:"mediumImage"`
	Name                string                       `pulumi:"name"`
	ProvisioningState   string                       `pulumi:"provisioningState"`
	SchemaItemTypeLink  *string                      `pulumi:"schemaItemTypeLink"`
	SmallImage          *string                      `pulumi:"smallImage"`
	StrongIds           []StrongIdResponse           `pulumi:"strongIds"`
	TenantId            string                       `pulumi:"tenantId"`
	TimestampFieldName  *string                      `pulumi:"timestampFieldName"`
	Type                string                       `pulumi:"type"`
	TypeName            *string                      `pulumi:"typeName"`
}
