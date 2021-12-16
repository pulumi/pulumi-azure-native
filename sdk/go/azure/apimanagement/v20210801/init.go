


package v20210801

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
	case "azure-native:apimanagement/v20210801:Api":
		r = &Api{}
	case "azure-native:apimanagement/v20210801:ApiDiagnostic":
		r = &ApiDiagnostic{}
	case "azure-native:apimanagement/v20210801:ApiIssue":
		r = &ApiIssue{}
	case "azure-native:apimanagement/v20210801:ApiIssueAttachment":
		r = &ApiIssueAttachment{}
	case "azure-native:apimanagement/v20210801:ApiIssueComment":
		r = &ApiIssueComment{}
	case "azure-native:apimanagement/v20210801:ApiManagementService":
		r = &ApiManagementService{}
	case "azure-native:apimanagement/v20210801:ApiOperation":
		r = &ApiOperation{}
	case "azure-native:apimanagement/v20210801:ApiOperationPolicy":
		r = &ApiOperationPolicy{}
	case "azure-native:apimanagement/v20210801:ApiPolicy":
		r = &ApiPolicy{}
	case "azure-native:apimanagement/v20210801:ApiRelease":
		r = &ApiRelease{}
	case "azure-native:apimanagement/v20210801:ApiSchema":
		r = &ApiSchema{}
	case "azure-native:apimanagement/v20210801:ApiTagDescription":
		r = &ApiTagDescription{}
	case "azure-native:apimanagement/v20210801:ApiVersionSet":
		r = &ApiVersionSet{}
	case "azure-native:apimanagement/v20210801:AuthorizationServer":
		r = &AuthorizationServer{}
	case "azure-native:apimanagement/v20210801:Backend":
		r = &Backend{}
	case "azure-native:apimanagement/v20210801:Cache":
		r = &Cache{}
	case "azure-native:apimanagement/v20210801:Certificate":
		r = &Certificate{}
	case "azure-native:apimanagement/v20210801:ContentItem":
		r = &ContentItem{}
	case "azure-native:apimanagement/v20210801:ContentType":
		r = &ContentType{}
	case "azure-native:apimanagement/v20210801:Diagnostic":
		r = &Diagnostic{}
	case "azure-native:apimanagement/v20210801:EmailTemplate":
		r = &EmailTemplate{}
	case "azure-native:apimanagement/v20210801:Gateway":
		r = &Gateway{}
	case "azure-native:apimanagement/v20210801:GatewayApiEntityTag":
		r = &GatewayApiEntityTag{}
	case "azure-native:apimanagement/v20210801:GatewayCertificateAuthority":
		r = &GatewayCertificateAuthority{}
	case "azure-native:apimanagement/v20210801:GatewayHostnameConfiguration":
		r = &GatewayHostnameConfiguration{}
	case "azure-native:apimanagement/v20210801:GlobalSchema":
		r = &GlobalSchema{}
	case "azure-native:apimanagement/v20210801:Group":
		r = &Group{}
	case "azure-native:apimanagement/v20210801:GroupUser":
		r = &GroupUser{}
	case "azure-native:apimanagement/v20210801:IdentityProvider":
		r = &IdentityProvider{}
	case "azure-native:apimanagement/v20210801:Logger":
		r = &Logger{}
	case "azure-native:apimanagement/v20210801:NamedValue":
		r = &NamedValue{}
	case "azure-native:apimanagement/v20210801:NotificationRecipientEmail":
		r = &NotificationRecipientEmail{}
	case "azure-native:apimanagement/v20210801:NotificationRecipientUser":
		r = &NotificationRecipientUser{}
	case "azure-native:apimanagement/v20210801:OpenIdConnectProvider":
		r = &OpenIdConnectProvider{}
	case "azure-native:apimanagement/v20210801:Policy":
		r = &Policy{}
	case "azure-native:apimanagement/v20210801:PrivateEndpointConnectionByName":
		r = &PrivateEndpointConnectionByName{}
	case "azure-native:apimanagement/v20210801:Product":
		r = &Product{}
	case "azure-native:apimanagement/v20210801:ProductApi":
		r = &ProductApi{}
	case "azure-native:apimanagement/v20210801:ProductGroup":
		r = &ProductGroup{}
	case "azure-native:apimanagement/v20210801:ProductPolicy":
		r = &ProductPolicy{}
	case "azure-native:apimanagement/v20210801:Subscription":
		r = &Subscription{}
	case "azure-native:apimanagement/v20210801:Tag":
		r = &Tag{}
	case "azure-native:apimanagement/v20210801:TagByApi":
		r = &TagByApi{}
	case "azure-native:apimanagement/v20210801:TagByOperation":
		r = &TagByOperation{}
	case "azure-native:apimanagement/v20210801:TagByProduct":
		r = &TagByProduct{}
	case "azure-native:apimanagement/v20210801:User":
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
		"apimanagement/v20210801",
		&module{version},
	)
}
