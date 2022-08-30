


package v20211201preview

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
	case "azure-native:apimanagement/v20211201preview:Api":
		r = &Api{}
	case "azure-native:apimanagement/v20211201preview:ApiDiagnostic":
		r = &ApiDiagnostic{}
	case "azure-native:apimanagement/v20211201preview:ApiIssue":
		r = &ApiIssue{}
	case "azure-native:apimanagement/v20211201preview:ApiIssueAttachment":
		r = &ApiIssueAttachment{}
	case "azure-native:apimanagement/v20211201preview:ApiIssueComment":
		r = &ApiIssueComment{}
	case "azure-native:apimanagement/v20211201preview:ApiManagementService":
		r = &ApiManagementService{}
	case "azure-native:apimanagement/v20211201preview:ApiOperation":
		r = &ApiOperation{}
	case "azure-native:apimanagement/v20211201preview:ApiOperationPolicy":
		r = &ApiOperationPolicy{}
	case "azure-native:apimanagement/v20211201preview:ApiPolicy":
		r = &ApiPolicy{}
	case "azure-native:apimanagement/v20211201preview:ApiRelease":
		r = &ApiRelease{}
	case "azure-native:apimanagement/v20211201preview:ApiSchema":
		r = &ApiSchema{}
	case "azure-native:apimanagement/v20211201preview:ApiTagDescription":
		r = &ApiTagDescription{}
	case "azure-native:apimanagement/v20211201preview:ApiVersionSet":
		r = &ApiVersionSet{}
	case "azure-native:apimanagement/v20211201preview:AuthorizationServer":
		r = &AuthorizationServer{}
	case "azure-native:apimanagement/v20211201preview:Backend":
		r = &Backend{}
	case "azure-native:apimanagement/v20211201preview:Cache":
		r = &Cache{}
	case "azure-native:apimanagement/v20211201preview:Certificate":
		r = &Certificate{}
	case "azure-native:apimanagement/v20211201preview:ContentItem":
		r = &ContentItem{}
	case "azure-native:apimanagement/v20211201preview:ContentType":
		r = &ContentType{}
	case "azure-native:apimanagement/v20211201preview:Diagnostic":
		r = &Diagnostic{}
	case "azure-native:apimanagement/v20211201preview:EmailTemplate":
		r = &EmailTemplate{}
	case "azure-native:apimanagement/v20211201preview:Gateway":
		r = &Gateway{}
	case "azure-native:apimanagement/v20211201preview:GatewayApiEntityTag":
		r = &GatewayApiEntityTag{}
	case "azure-native:apimanagement/v20211201preview:GatewayCertificateAuthority":
		r = &GatewayCertificateAuthority{}
	case "azure-native:apimanagement/v20211201preview:GatewayHostnameConfiguration":
		r = &GatewayHostnameConfiguration{}
	case "azure-native:apimanagement/v20211201preview:GlobalSchema":
		r = &GlobalSchema{}
	case "azure-native:apimanagement/v20211201preview:Group":
		r = &Group{}
	case "azure-native:apimanagement/v20211201preview:GroupUser":
		r = &GroupUser{}
	case "azure-native:apimanagement/v20211201preview:IdentityProvider":
		r = &IdentityProvider{}
	case "azure-native:apimanagement/v20211201preview:Logger":
		r = &Logger{}
	case "azure-native:apimanagement/v20211201preview:NamedValue":
		r = &NamedValue{}
	case "azure-native:apimanagement/v20211201preview:NotificationRecipientEmail":
		r = &NotificationRecipientEmail{}
	case "azure-native:apimanagement/v20211201preview:NotificationRecipientUser":
		r = &NotificationRecipientUser{}
	case "azure-native:apimanagement/v20211201preview:OpenIdConnectProvider":
		r = &OpenIdConnectProvider{}
	case "azure-native:apimanagement/v20211201preview:Policy":
		r = &Policy{}
	case "azure-native:apimanagement/v20211201preview:PolicyFragment":
		r = &PolicyFragment{}
	case "azure-native:apimanagement/v20211201preview:PrivateEndpointConnectionByName":
		r = &PrivateEndpointConnectionByName{}
	case "azure-native:apimanagement/v20211201preview:Product":
		r = &Product{}
	case "azure-native:apimanagement/v20211201preview:ProductApi":
		r = &ProductApi{}
	case "azure-native:apimanagement/v20211201preview:ProductGroup":
		r = &ProductGroup{}
	case "azure-native:apimanagement/v20211201preview:ProductPolicy":
		r = &ProductPolicy{}
	case "azure-native:apimanagement/v20211201preview:Subscription":
		r = &Subscription{}
	case "azure-native:apimanagement/v20211201preview:Tag":
		r = &Tag{}
	case "azure-native:apimanagement/v20211201preview:TagByApi":
		r = &TagByApi{}
	case "azure-native:apimanagement/v20211201preview:TagByOperation":
		r = &TagByOperation{}
	case "azure-native:apimanagement/v20211201preview:TagByProduct":
		r = &TagByProduct{}
	case "azure-native:apimanagement/v20211201preview:User":
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
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"apimanagement/v20211201preview",
		&module{version},
	)
}
