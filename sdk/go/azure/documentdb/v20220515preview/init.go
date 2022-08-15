


package v20220515preview

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
	case "azure-native:documentdb/v20220515preview:CassandraCluster":
		r = &CassandraCluster{}
	case "azure-native:documentdb/v20220515preview:CassandraDataCenter":
		r = &CassandraDataCenter{}
	case "azure-native:documentdb/v20220515preview:CassandraResourceCassandraKeyspace":
		r = &CassandraResourceCassandraKeyspace{}
	case "azure-native:documentdb/v20220515preview:CassandraResourceCassandraTable":
		r = &CassandraResourceCassandraTable{}
	case "azure-native:documentdb/v20220515preview:CassandraResourceCassandraView":
		r = &CassandraResourceCassandraView{}
	case "azure-native:documentdb/v20220515preview:DatabaseAccount":
		r = &DatabaseAccount{}
	case "azure-native:documentdb/v20220515preview:GraphResourceGraph":
		r = &GraphResourceGraph{}
	case "azure-native:documentdb/v20220515preview:GremlinResourceGremlinDatabase":
		r = &GremlinResourceGremlinDatabase{}
	case "azure-native:documentdb/v20220515preview:GremlinResourceGremlinGraph":
		r = &GremlinResourceGremlinGraph{}
	case "azure-native:documentdb/v20220515preview:MongoDBResourceMongoDBCollection":
		r = &MongoDBResourceMongoDBCollection{}
	case "azure-native:documentdb/v20220515preview:MongoDBResourceMongoDBDatabase":
		r = &MongoDBResourceMongoDBDatabase{}
	case "azure-native:documentdb/v20220515preview:MongoDBResourceMongoRoleDefinition":
		r = &MongoDBResourceMongoRoleDefinition{}
	case "azure-native:documentdb/v20220515preview:MongoDBResourceMongoUserDefinition":
		r = &MongoDBResourceMongoUserDefinition{}
	case "azure-native:documentdb/v20220515preview:NotebookWorkspace":
		r = &NotebookWorkspace{}
	case "azure-native:documentdb/v20220515preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:documentdb/v20220515preview:Service":
		r = &Service{}
	case "azure-native:documentdb/v20220515preview:SqlResourceSqlContainer":
		r = &SqlResourceSqlContainer{}
	case "azure-native:documentdb/v20220515preview:SqlResourceSqlDatabase":
		r = &SqlResourceSqlDatabase{}
	case "azure-native:documentdb/v20220515preview:SqlResourceSqlRoleAssignment":
		r = &SqlResourceSqlRoleAssignment{}
	case "azure-native:documentdb/v20220515preview:SqlResourceSqlRoleDefinition":
		r = &SqlResourceSqlRoleDefinition{}
	case "azure-native:documentdb/v20220515preview:SqlResourceSqlStoredProcedure":
		r = &SqlResourceSqlStoredProcedure{}
	case "azure-native:documentdb/v20220515preview:SqlResourceSqlTrigger":
		r = &SqlResourceSqlTrigger{}
	case "azure-native:documentdb/v20220515preview:SqlResourceSqlUserDefinedFunction":
		r = &SqlResourceSqlUserDefinedFunction{}
	case "azure-native:documentdb/v20220515preview:TableResourceTable":
		r = &TableResourceTable{}
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
		"documentdb/v20220515preview",
		&module{version},
	)
}
