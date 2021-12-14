


package v20210401preview

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:apimanagement/v20210401preview:Api":
		r = &Api{}
	case "azure-native:apimanagement/v20210401preview:ApiDiagnostic":
		r = &ApiDiagnostic{}
	case "azure-native:apimanagement/v20210401preview:ApiIssue":
		r = &ApiIssue{}
	case "azure-native:apimanagement/v20210401preview:ApiIssueAttachment":
		r = &ApiIssueAttachment{}
	case "azure-native:apimanagement/v20210401preview:ApiIssueComment":
		r = &ApiIssueComment{}
	case "azure-native:apimanagement/v20210401preview:ApiManagementService":
		r = &ApiManagementService{}
	case "azure-native:apimanagement/v20210401preview:ApiOperation":
		r = &ApiOperation{}
	case "azure-native:apimanagement/v20210401preview:ApiOperationPolicy":
		r = &ApiOperationPolicy{}
	case "azure-native:apimanagement/v20210401preview:ApiPolicy":
		r = &ApiPolicy{}
	case "azure-native:apimanagement/v20210401preview:ApiRelease":
		r = &ApiRelease{}
	case "azure-native:apimanagement/v20210401preview:ApiSchema":
		r = &ApiSchema{}
	case "azure-native:apimanagement/v20210401preview:ApiTagDescription":
		r = &ApiTagDescription{}
	case "azure-native:apimanagement/v20210401preview:ApiVersionSet":
		r = &ApiVersionSet{}
	case "azure-native:apimanagement/v20210401preview:AuthorizationServer":
		r = &AuthorizationServer{}
	case "azure-native:apimanagement/v20210401preview:Backend":
		r = &Backend{}
	case "azure-native:apimanagement/v20210401preview:Cache":
		r = &Cache{}
	case "azure-native:apimanagement/v20210401preview:Certificate":
		r = &Certificate{}
	case "azure-native:apimanagement/v20210401preview:ContentItem":
		r = &ContentItem{}
	case "azure-native:apimanagement/v20210401preview:ContentType":
		r = &ContentType{}
	case "azure-native:apimanagement/v20210401preview:Diagnostic":
		r = &Diagnostic{}
	case "azure-native:apimanagement/v20210401preview:EmailTemplate":
		r = &EmailTemplate{}
	case "azure-native:apimanagement/v20210401preview:Gateway":
		r = &Gateway{}
	case "azure-native:apimanagement/v20210401preview:GatewayApiEntityTag":
		r = &GatewayApiEntityTag{}
	case "azure-native:apimanagement/v20210401preview:GatewayCertificateAuthority":
		r = &GatewayCertificateAuthority{}
	case "azure-native:apimanagement/v20210401preview:GatewayHostnameConfiguration":
		r = &GatewayHostnameConfiguration{}
	case "azure-native:apimanagement/v20210401preview:Group":
		r = &Group{}
	case "azure-native:apimanagement/v20210401preview:GroupUser":
		r = &GroupUser{}
	case "azure-native:apimanagement/v20210401preview:IdentityProvider":
		r = &IdentityProvider{}
	case "azure-native:apimanagement/v20210401preview:Logger":
		r = &Logger{}
	case "azure-native:apimanagement/v20210401preview:NamedValue":
		r = &NamedValue{}
	case "azure-native:apimanagement/v20210401preview:NotificationRecipientEmail":
		r = &NotificationRecipientEmail{}
	case "azure-native:apimanagement/v20210401preview:NotificationRecipientUser":
		r = &NotificationRecipientUser{}
	case "azure-native:apimanagement/v20210401preview:OpenIdConnectProvider":
		r = &OpenIdConnectProvider{}
	case "azure-native:apimanagement/v20210401preview:Policy":
		r = &Policy{}
	case "azure-native:apimanagement/v20210401preview:PrivateEndpointConnectionByName":
		r = &PrivateEndpointConnectionByName{}
	case "azure-native:apimanagement/v20210401preview:Product":
		r = &Product{}
	case "azure-native:apimanagement/v20210401preview:ProductApi":
		r = &ProductApi{}
	case "azure-native:apimanagement/v20210401preview:ProductGroup":
		r = &ProductGroup{}
	case "azure-native:apimanagement/v20210401preview:ProductPolicy":
		r = &ProductPolicy{}
	case "azure-native:apimanagement/v20210401preview:Schema":
		r = &Schema{}
	case "azure-native:apimanagement/v20210401preview:Subscription":
		r = &Subscription{}
	case "azure-native:apimanagement/v20210401preview:Tag":
		r = &Tag{}
	case "azure-native:apimanagement/v20210401preview:TagByApi":
		r = &TagByApi{}
	case "azure-native:apimanagement/v20210401preview:TagByOperation":
		r = &TagByOperation{}
	case "azure-native:apimanagement/v20210401preview:TagByProduct":
		r = &TagByProduct{}
	case "azure-native:apimanagement/v20210401preview:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"apimanagement/v20210401preview",
		&module{version},
	)
}
