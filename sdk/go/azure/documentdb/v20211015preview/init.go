


package v20211015preview

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
	case "azure-native:documentdb/v20211015preview:CassandraCluster":
		r = &CassandraCluster{}
	case "azure-native:documentdb/v20211015preview:CassandraDataCenter":
		r = &CassandraDataCenter{}
	case "azure-native:documentdb/v20211015preview:CassandraResourceCassandraKeyspace":
		r = &CassandraResourceCassandraKeyspace{}
	case "azure-native:documentdb/v20211015preview:CassandraResourceCassandraTable":
		r = &CassandraResourceCassandraTable{}
	case "azure-native:documentdb/v20211015preview:CassandraResourceCassandraView":
		r = &CassandraResourceCassandraView{}
	case "azure-native:documentdb/v20211015preview:DatabaseAccount":
		r = &DatabaseAccount{}
	case "azure-native:documentdb/v20211015preview:GraphResourceGraph":
		r = &GraphResourceGraph{}
	case "azure-native:documentdb/v20211015preview:GremlinResourceGremlinDatabase":
		r = &GremlinResourceGremlinDatabase{}
	case "azure-native:documentdb/v20211015preview:GremlinResourceGremlinGraph":
		r = &GremlinResourceGremlinGraph{}
	case "azure-native:documentdb/v20211015preview:MongoDBResourceMongoDBCollection":
		r = &MongoDBResourceMongoDBCollection{}
	case "azure-native:documentdb/v20211015preview:MongoDBResourceMongoDBDatabase":
		r = &MongoDBResourceMongoDBDatabase{}
	case "azure-native:documentdb/v20211015preview:MongoDBResourceMongoRoleDefinition":
		r = &MongoDBResourceMongoRoleDefinition{}
	case "azure-native:documentdb/v20211015preview:MongoDBResourceMongoUserDefinition":
		r = &MongoDBResourceMongoUserDefinition{}
	case "azure-native:documentdb/v20211015preview:NotebookWorkspace":
		r = &NotebookWorkspace{}
	case "azure-native:documentdb/v20211015preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:documentdb/v20211015preview:Service":
		r = &Service{}
	case "azure-native:documentdb/v20211015preview:SqlResourceSqlContainer":
		r = &SqlResourceSqlContainer{}
	case "azure-native:documentdb/v20211015preview:SqlResourceSqlDatabase":
		r = &SqlResourceSqlDatabase{}
	case "azure-native:documentdb/v20211015preview:SqlResourceSqlRoleAssignment":
		r = &SqlResourceSqlRoleAssignment{}
	case "azure-native:documentdb/v20211015preview:SqlResourceSqlRoleDefinition":
		r = &SqlResourceSqlRoleDefinition{}
	case "azure-native:documentdb/v20211015preview:SqlResourceSqlStoredProcedure":
		r = &SqlResourceSqlStoredProcedure{}
	case "azure-native:documentdb/v20211015preview:SqlResourceSqlTrigger":
		r = &SqlResourceSqlTrigger{}
	case "azure-native:documentdb/v20211015preview:SqlResourceSqlUserDefinedFunction":
		r = &SqlResourceSqlUserDefinedFunction{}
	case "azure-native:documentdb/v20211015preview:TableResourceTable":
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
		"documentdb/v20211015preview",
		&module{version},
	)
}
