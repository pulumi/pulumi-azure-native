module github.com/pulumi/pulumi-azure-nextgen/provider

go 1.13

require (
	github.com/Azure/go-autorest/autorest v0.10.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/gedex/inflector v0.0.0-20170307190818-16278e9db813
	github.com/go-openapi/jsonreference v0.19.2
	github.com/go-openapi/spec v0.19.7
	github.com/go-openapi/swag v0.19.9
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/go-azure-helpers v0.10.0
	github.com/hashicorp/hcl/v2 v2.6.0
	github.com/magiconair/properties v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi/pkg/v2 v2.10.2-0.20200920185652-988ebd8b8e6f
	github.com/pulumi/pulumi/sdk v1.14.1
	github.com/pulumi/pulumi/sdk/v2 v2.10.2-0.20200920185652-988ebd8b8e6f
	github.com/schollz/progressbar/v3 v3.5.0
	github.com/sourcegraph/jsonx v0.0.0-20200629203448-1a936bd500cf
	github.com/stretchr/testify v1.6.1
	github.com/zclconf/go-cty v1.3.1
)

replace (
	github.com/Azure/go-autorest => github.com/tombuildsstuff/go-autorest v14.0.1-0.20200416184303-d4e299a3c04a+incompatible
	github.com/Azure/go-autorest/autorest => github.com/tombuildsstuff/go-autorest/autorest v0.10.1-0.20200416184303-d4e299a3c04a
	github.com/Azure/go-autorest/autorest/azure/auth => github.com/tombuildsstuff/go-autorest/autorest/azure/auth v0.4.3-0.20200416184303-d4e299a3c04a
)
